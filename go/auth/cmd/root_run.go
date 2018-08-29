package cmd

import (
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	log "github.com/cihub/seelog"
	"time"
	"auth/config"
	"auth/module"
	"os"
	"os/signal"
	"syscall"
	"net"
	"google.golang.org/grpc"
	"auth/common"
	"auth/internal/api"
	"auth/filter"
	"auth/rpc"
	pb "auth/api"
	"auth/go-drbac/drbac"
)

func run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	tasks := []func() error{
		setLogLevel,
		printStartMessage,
		setMysqlSQLConnection,
		setRedisPool,
		startClientAPI(ctx),
	}
	for _, t := range tasks {
		if err := t(); err != nil {
			panic(err)
		}
	}

	sigChan := make(chan os.Signal)
	exitChan := make(chan struct{})
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	log.Info("signal", <-sigChan, "signal received")
	go func() {
		log.Warn("stopping auth")
		// todo: handle graceful shutdown?
		exitChan <- struct{}{}
	}()
	select {
	case <-exitChan:
	case s := <-sigChan:
		log.Info("signal", s, "signal received, stopping immediately")
	}

	return nil
}

func setLogLevel() error {
	logger, err := log.LoggerFromConfigAsString(config.LogConfig)
	if err != nil {
		log.Debugf("Deprecation warning! no log configuration file found, falling back on environment variables. Update your log configuration,%s",err)
		return err
	}
	if err = log.ReplaceLogger(logger);err != nil {
		return err
	}
	return nil
}

func printStartMessage() error {
	log.Infof("starting Auth Server, version:%s",version)
	return nil
}

func setMysqlSQLConnection() error {
	log.Info("connecting to mysql")
	dsn := config.C.MySQL.User+":"+config.C.MySQL.Password+"@tcp("+config.C.MySQL.Host+":"+config.C.MySQL.Port+")/"+config.C.MySQL.Database+"?charset=utf8&loc=Asia%2FShanghai&parseTime=true"
	db, err := module.OpenDatabase(dsn)
	if err != nil {
		return errors.Wrap(err, "database connection error")
	}
	config.C.MySQL.DB = db
	return nil
}

func setRedisPool() error {
	// setup redis pool
	log.Info("setup redis connection pool")
	config.C.Redis.Pool = module.NewRedisPool(config.C.Redis.URL)
	return config.C.Redis.Pool.TestOnBorrow(config.C.Redis.Pool.Get(),time.Now())
}

func startClientAPI(ctx context.Context) func() error {
	return func() error {
		var err error
		log.Info("注册rpc的服务")
		var opts []grpc.ServerOption
		opts = append(opts,grpc.UnaryInterceptor(
			filter.ChainUnaryServer(filter.PrometheusInterceptor(),filter.AuthInterceptor())),)
		clientAPIHandler := grpc.NewServer(opts...)
		log.Infof("启动go-drbac框架服务")
		server, err := drbac.NewDrbacServer(config.C.MySQL.DB,config.C.Redis.Pool)
		if err != nil {
			log.Errorf("加载权限控制框架失败")
			panic(err)
		}
		common.Drbac = server
		log.Info("开始注册服务")
		pb.RegisterRoleServerServer(clientAPIHandler, &api.RoleServer{DrbacServer:server})
		pb.RegisterAuthServerServer(clientAPIHandler,&api.AuthServer{DrbacServer:server})
		pb.RegisterUserServerServer(clientAPIHandler,&api.UserServer{DrbacServer:server})
		pb.RegisterTenantServerServer(clientAPIHandler,&api.TenantServer{DrbacServer:server})
		pb.RegisterDomainServerServer(clientAPIHandler,&api.DomainServer{DrbacServer:server})
		pb.RegisterActionLogServer(clientAPIHandler,&rpc.ActionLogServer{})
		pb.RegisterServiceServerServer(clientAPIHandler,&api.ServiceServer{DrbacServer:server})
		pb.RegisterPolicyServerServer(clientAPIHandler,&api.PolicyServer{DrbacServer:server})
		pb.RegisterPaymentServerServer(clientAPIHandler,&api.PaymentServer{DrbacServer:server})
		pb.RegisterTradingServerServer(clientAPIHandler,&api.TradingServer{DrbacServer:server})

		/*启动auth的rpc服务端*/
		AuthHostname := config.C.Auth.Host + ":" + config.C.Auth.Port
		log.Info("启动rpc服务监听地址:", AuthHostname)
		ln, err := net.Listen("tcp", AuthHostname)
		if err != nil {
			log.Error("启动rpc服务监听端口异常:", err)
			return err
		}
		go clientAPIHandler.Serve(ln)
		return nil
	}
}




