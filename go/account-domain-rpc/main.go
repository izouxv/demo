package main

import (
	"flag"
	"net"
	"strconv"
	api "account-domain-rpc/api"
	"account-domain-rpc/api/adv"
	"account-domain-rpc/api/feedback"
	"account-domain-rpc/api/setting"
	"account-domain-rpc/core"
	"account-domain-rpc/filter"
	"account-domain-rpc/metrics"
	"account-domain-rpc/rpc"
	log "github.com/cihub/seelog"
	"google.golang.org/grpc"
	"account-domain-rpc/api/user"
)

var (
	cfgFile *string = flag.String("config", "config/conf.yaml", "config path")
	seelog  *string = flag.String("seelog", "config/seelog.xml", "log config path")
)

func initLoger() {
	logger, err := log.LoggerFromConfigAsFile(*seelog)
	if err != nil {
		log.Errorf("err parsing config log file :%s", err)
		return
	}
	log.ReplaceLogger(logger)
}

func main() {
	flag.Parse()
	initLoger()
	log.Infof("Start account-domain-rpc server ")
	core.ContextInit(*cfgFile)

	lis, err := net.Listen("tcp", net.JoinHostPort(core.GetContext().Server.Host, strconv.Itoa(core.GetContext().Server.Port)))
	if err != nil {
		log.Error(err)
		return
	}

	if core.GetContext().Metrics.Enable == true {
		metrics.Run()
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			filter.ChainUnaryServer(filter.PrometheusInterceptor())),
	)
	RegisterRpcServer(s)

	if err = s.Serve(lis); err != nil {
		log.Error(err)
		return
	}
}

func RegisterRpcServer(s *grpc.Server) {
	log.Info("-----RegisterRpcServer-----")
	setting.RegisterRadacatVersionServer(s, &rpc.RadacatVersionServer{})
	adv.RegisterAdvertisementServer(s, &rpc.AdvertisementServer{})
	feedback.RegisterFeedBackServer(s, &rpc.FeedBackServer{})
	api.RegisterDeviceTypeServer(s,&rpc.DeviceTypeServer{})
	user.RegisterTestUserServerServer(s,&rpc.TestUserServer{})
}

