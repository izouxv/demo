package routers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"petfone-http/db"
	"petfone-http/result"
	"petfone-http/rpc"
	"time"
	"os"
	"runtime/debug"
	"strings"
	"fmt"
	. "petfone-http/util"
	"petfone-http/pb"
	"petfone-http/core"
	log "github.com/cihub/seelog"
	"context"
	"net"
	"flag"
	"os/signal"
	"syscall"
)

var (
	httpServer = &http.Server{
		ReadTimeout: 60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout: 60 * time.Second,
		//TLSConfig		: tls.Config{},
	}
	Router = httprouter.New()
	cfgFile = flag.String("c", Conf, "config path")
	seeLog  = flag.String("log", SeeLog, "log config path")
	halfOpenAddress = map[string]byte{
		serverName + version1 + "/services/faq":	3,
		serverName + version1 + "/chat/faq":		3,
		serverName + version1 + "/mobile/":      	3,
		serverName + version1 + "/mail/":        	3,
		serverName + version1 + "/reg":          	3,
		serverName + version1 + "/feedbacks":    	3,
		serverName + version1 + "/version/":     	3,
		serverName + version1 + "/advertisement": 	3,
		serverName + version1 + "/rpwd":          	3,
		serverName + version1 + "/rpwdm/":        	3,
		serverName + version1 + "/breeds/":       	3,
		serverName + version11 + "/breeds/":      	3,
		serverName + version1 + "/sessions":      	4,
	}
)

func init() {
	flag.Parse()
	log.Info("pet server start")
	//runtime.GOMAXPROCS(runtime.NumCPU())
	var signalChan = make(chan os.Signal, 1)
	signal.Notify(signalChan)
	go func() {
		select {
		case s := <-signalChan:
			if s == syscall.SIGPIPE {
				break
			}
			log.Info("close conn...")
			closeConn()
			log.Info("Go signal:", s)
		}
	}()
	go Runpprof()
	setLogger()
	core.ContextInit(*cfgFile)
	accountRouter(Router)
	settingRouter(Router)
	dataRouter(Router)
	deviceAndPetRouter(Router)
	otherRouter(Router)
	addr := net.JoinHostPort(core.GetContext().Server.Host, IntToStr(core.GetContext().Server.Port))
	httpServer.Addr = addr
	httpServer.Handler = Router
	log.Info("init error:", 	httpServer.ListenAndServe())
}

//关闭连接
func closeConn() {
	core.CloseRedis()
	core.CloseRpc()
	httpServer.Close()
}

//加载日志配置
func setLogger() {
	logger, err := log.LoggerFromConfigAsFile(*seeLog)
	if err != nil {
		log.Error("err parsing config log file", err)
		return
	}
	log.ReplaceLogger(logger)
}

//拦截器
func authInterceptor(handle httprouter.Handle) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		entryTime := GetNowTime()
		url, method := req.URL.Path, req.Method
		source, contentType := req.Header.Get("source"), req.Header.Get("Content-Type")
		log.Infof("authInterceptor-URL:%v, Method:%v, Content-Type:%v, Source:%v", url, method, contentType, source)
		length := req.ContentLength
		if length > 524288000 {
			log.Error("authInterceptor ContentLength:", length)
			result.ResCode(404, res)
			return
		}
		//todo context
		if db.VerifySource(source) {
			result.RESC(33010, res)
			return
		}
		//对接口判断
		sso := &pb.SsoRequest{Source: source}
		if verifyHalfOpenAddress(url, method) {
			token := req.Header.Get("token")
			if VerifyParamsStr(token) {
				sessionName, err := req.Cookie("token")
				if err != nil || sessionName.Value == "" {
					result.RESC(23019, res)
					return
				}
				token = sessionName.Value
			}
			//调用RPC
			sso.SessionName = token
			sso.Token = token
			ssoR := rpc.SsoRpc(sso, "GetUserInfo")
			if ssoR.Code != 10000 {
				result.RESC(ssoR.Code, res)
				return
			}
			SsoTssoR(sso, ssoR)
			sso.Uid = ssoR.Uid
			req.Header.Add("uid", Int32ToStr(ssoR.Uid))
		}
		req = req.WithContext(context.WithValue(req.Context(), "contextSso", sso))
		handle(res, req, params)
		defer func() {
			duration := time.Now().Sub(entryTime)
			if p := recover(); p != nil {
				log.Error("panic --- recover:", p)
				log.Error("进程名称:", os.Args[0])
				log.Error("进程ID:", os.Getpid())
				log.Error("堆栈信息:", string(debug.Stack()))
				result.RESC(10001, res)
			}
			log.Infof("authInterceptor - URL: %v, Method: %v, Consuming:%v", url, method, duration)
			go func(req *http.Request,res http.ResponseWriter,url,method string,duration time.Duration,entryTime time.Time) {
				uid, _ := StrToInt32(req.Header.Get("uid"))
				ip :=  strings.Split(req.RemoteAddr, ":")[0]
				rpc.PetfoneRpcActionLog(&pb.AgentInfo{Uid:uid, Token:req.Header.Get("token"), Path:url, Method:method,
					Duration:duration.String(),Ip:ip, DevInfo:req.UserAgent(),
					Code:res.Header().Get("code"),CreateTime:entryTime.Unix()})
			}(req,res,url,method,duration,entryTime)
		}()
	}
}
func interceptor(handle httprouter.Handle) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		fmt.Println("qqqqqqqqqqqqqq")
		handle(res, req, params)
	}
}

//接口地址
const (
	//项目路径
	serverName = "/petfone"
	version1   = "/v1.0"
	version11  = "/v1.1"
	version12  = "/v1.2"

	//todo v1.0 半公接口路径
	uriMobile          	= serverName + version1 + "/mobile/:action"
	uriReg             	= serverName + version1 + "/reg"
	uriFaq             	= serverName + version1 + "/services/faq"
	uriChatFaq			= serverName + version1 + "/chat/faq"
	uriVersion         	= serverName + version1 + "/version/:name/:code"
	uriVersion1         = serverName + version11 + "/version/:name"
	uriAdvertisement	= serverName + version1 + "/advertisement"
	uriFeedback       	= serverName + version1 + "/feedbacks"
	uriResetPwd       	= serverName + version1 + "/rpwd"
	uriResetMail      	= serverName + version1 + "/rpwdm/:action"
	uriFileBreeds     	= serverName + version1 + "/breeds/:types"
	uriFileBreeds11 	= serverName + version11 + "/breeds/:types"

	//todo v1.1
	uriMobile11      = serverName + version11 + "/mobile"
	uriMail11        = serverName + version11 + "/mail"
	uriAccounts11    = serverName + version11 + "/accounts"
	uriAccountsPwd11 = serverName + version11 + "/accounts/pwd"

	/*
		混用接口路径
	*/
	uriSessions = serverName + version1 + "/sessions"
	/*
		封闭接口路径
	*/
	//个人信息
	uriUserinfo = serverName + version1 + "/info"
	//用户信息
	uriAccounts = serverName + version1 + "/accounts/:param"
	//密码修改
	uriUpdatePwd = serverName + version1 + "/upwd/:action"
	//设备围栏
	uriFenceDy = serverName + version1 + "/fence/dynamic"
	//文件
	uriFiles = serverName + version1 + "/files/:use"
	uriImages = serverName + version1 + "/images/:fid"
	uriFiles1 = serverName + version11 + "/files/train"
	//设备
	uriDevices    = serverName + version1 + "/devices"
	uriDeviceDid  = serverName + version1 + "/devices/:did"
	uriStatistics = serverName + version1 + "/statistics/:dev"
	//宠物资料
	uriPetPid = serverName + version1 + "/pets/:pid"
	uriPets   = serverName + version1 + "/pets"
	uriPetPid1 = serverName + version11 + "/pets/:pid"
	uriPets1  = serverName + version11 + "/pets"
	//宠聊
	uriChat = serverName + version1 + "/chat/pets/:pid"
	uriChat11 = serverName + version11 + "/chat/pets/:pid"
	uriChat12 = serverName + version12 + "/chat/pets/:pid"
	//宠物与设备
	uriPetpiddev = serverName + version1 + "/pets/:pid/devices/:did"
	//宠物训练
	uriPetstrainid = serverName + version1 + "/pets/:pid/trains/:id"
	uriPetstrainid1 = serverName + version11 + "/device/:did/trains/:id"
	uriPetstrain   = serverName + version1 + "/pets/:pid/trains"
	//运动数据
	uriPetData = serverName + version1 + "/pets/:pid/datas"
	uriDataPet = serverName + version11 + "/data/pet"
	//分享功能
	uriShare = serverName + version1 + "/share"
	//用户通知
	uriNotices = serverName + version1 + "/services/notice"
	uriNotice  = serverName + version1 + "/services/notice/:id"
	//用户设置
	uriSet     = serverName + version1 + "/setting"

)

//接口判断
func verifyHalfOpenAddress(url, method string) bool {
	f := true
	for k, v := range halfOpenAddress {
		if (strings.HasPrefix(url, k) && v == 3) || (strings.HasPrefix(url, k) && "POST" == method) {
			f = false
		}
	}
	return f
}
