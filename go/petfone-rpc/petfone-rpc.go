package main

import (
	"flag"
	"os"
	"os/signal"
	"net"
	"strconv"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	log "github.com/cihub/seelog"
	"petfone-rpc/core"
	"petfone-rpc/db"
	"petfone-rpc/filter"
	"petfone-rpc/pb"
	"petfone-rpc/service"
	"petfone-rpc/util"
	"github.com/jinzhu/gorm"
	"runtime"
)

var (
	cfgFile = flag.String("config", util.Conf, "config path")
	seeLog  = flag.String("seelog", util.SeeLog, "log config path")
	httpApiLog  = flag.String("api", util.HttpApiLogFilePrefix, "api log file")
)

func initLoger() {
	logger, err := log.LoggerFromConfigAsFile(*seeLog)
	if err != nil {
		log.Critical("err parsing config log file", err)
		return
	}
	log.ReplaceLogger(logger)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//signal.notify方法用来监听
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan)
	flag.Parse()
	initLoger()
	log.Info("petfone-rpc")
	util.HttpApiLogFilePrefix = *httpApiLog
	core.ContextInit(*cfgFile)
	//if core.GetContext().Metrics.Enable == true {
	//	log.Info("run metrics")
	//	metrics.Run()
	//}
	host := core.GetContext().Server.Host
	port :=strconv.Itoa(core.GetContext().Server.Port)
	log.Info("service:", host,":", port)
	lis, err := net.Listen("tcp", net.JoinHostPort(host, port))
	if err != nil {
		log.Error(err)
		panic(err)
	}
	server := grpc.NewServer(grpc.UnaryInterceptor(
		filter.ChainUnaryServer())) //filter.PrometheusInterceptor()
	//注册rpc服务
	pb.RegisterSsoServer(server, &service.Rpc_sso{})
	pb.RegisterMSsoServer(server, &service.Rpc_msso{})
	pb.RegisterAccountServer(server, &service.Rpc_account{})
	pb.RegisterDevicesServer(server, &service.DeviceRpc{})
	pb.RegisterPetfoneServer(server, &service.PetfoneRpc{})
	pb.RegisterPetInfoServer(server, &service.PetinfoRpc{})
	pb.RegisterShareManageServer(server, &service.ShareRpc{})
	pb.RegisterFaqCommonServer(server, &service.FaqCommonRpc{})
	pb.RegisterExerciseDataServer(server, &service.MotionDataRpc{})
	pb.RegisterNoticeServer(server, &service.NoticeRpc{})
	pb.RegisterFilesServer(server, &service.FileRpc{})
	reflection.Register(server)
	go Task()
	go func() {
		for s := range signalChan {
			log.Info("close conn...")
			core.MysqlClose()
			core.RedisClose()
			server.Stop()
			util.LogWriter.Close()
			log.Info("Go signal ", s)
		}
	}()
	log.Info("main-error :", server.Serve(lis))
}

//执行初始化与定期任务
func Task() {
	//todo 获取宠聊匹配信息
	petChatPo := db.PetChatPo{DataState:1}
	log.Info("读取宠聊模版")
	petChatPos,err := petChatPo.GetPetChatsDB()
	if err != nil || len(petChatPos) == 0 {
		log.Error("Task GetPetChat-读取宠聊模版失败 err:",err)
	}
	for _,v := range petChatPos {
		util.TagCnKeys = append(util.TagCnKeys,v.NameCn)
		util.TagCnValues = append(util.TagCnValues,v.InfoCn)
		util.TagEnKeys = append(util.TagEnKeys,v.NameEn)
		util.TagEnValues = append(util.TagEnValues,v.InfoEn)
	}

	//todo 0点定时执行
	for {
		startTime := time.Now()
		next := startTime.Add(time.Hour*24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
		t:= time.NewTimer(next.Sub(startTime))
		util.LogWriter, util.LogErr = log.NewFileWriter(util.HttpApiLogFilePrefix+startTime.Format("2006-01-02"))
		if util.LogErr != nil {
			log.Error("Task NewFileWriter err:",util.LogErr)
			util.LogErr = nil
		}
		<-t.C
		nowTime := util.GetNowTime()
		timerTask(nowTime)
	}
}

//创建新的并清除过期的
func timerTask(nowTime time.Time) string {
	log.Info("timerTask ---")
	util.LogWriter.Close()
	util.LogWriter, util.LogErr = log.NewFileWriter(util.HttpApiLogFilePrefix+nowTime.Format("2006-01-02"))
	if util.LogErr != nil {
		log.Error("timerTask NewFileWriter err:",util.LogErr)
		util.LogErr = nil
	}
	end := util.GetZeroTime(nowTime.Unix())
	start := util.GetZeroTime(nowTime.Unix()-86400)
	expiration := util.GetZeroTime(nowTime.Unix()-86400*14)
	err := db.Transaction(core.MysqlClient, func(tx *gorm.DB) error {
		db.SetClearPetTrainDB(core.MysqlClient,start,end,expiration)
		db.ClearRelationDB(core.MysqlClient,expiration)
		db.ClearActionDB(core.MysqlClient,expiration)
		return nil
	})
	if err != nil {
		log.Error("timerTask err:", err)
	}
	return time.Now().Sub(nowTime).String()
}
