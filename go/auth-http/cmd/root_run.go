package cmd

import (
	"context"
	"github.com/spf13/cobra"
	log "github.com/cihub/seelog"
	"time"
	"auth-http/config"
	"os"
	"os/signal"
	"syscall"
	"net"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"auth-http/controller"
	"auth-http/rpc"
)

func run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	tasks := []func() error{
		setLogLevel,
		printStartMessage,
		setServerClient,
		setServeHTTP,
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

func setServeHTTP () error{
	log.Info("start auth-http server")
	router := httprouter.New()
	controller.RouterMethod(router)
	log.Info("111:",config.C.AuthHttp.Host,config.C.AuthHttp.Port)
	log.Info("222:",config.C.Auth.HostName)
	address := net.JoinHostPort(config.C.AuthHttp.Host,config.C.AuthHttp.Port)
	http := &http.Server{Addr: address, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second,
		IdleTimeout: 10 * time.Second, Handler: router,            //TLSConfig:tls.Config{},
	}
	log.Info(http.ListenAndServe())
	return nil
}

func setServerClient() error {
	rpc.NewActionLogRpcClient(config.C.Auth.HostName)
	rpc.NewAuthRpcClient(config.C.Auth.HostName)
	rpc.NewDomainRpcClient(config.C.Auth.HostName)
	rpc.NewPaymentRpcClient(config.C.Auth.HostName)
	rpc.NewPolicyRpcClient(config.C.Auth.HostName)
	rpc.NewRoleClient(config.C.Auth.HostName)
	rpc.NewServiceRpcClient(config.C.Auth.HostName)
	rpc.NewTenantRpcClient(config.C.Auth.HostName)
	rpc.NewTradingRpcClient(config.C.Auth.HostName)
	rpc.NewUserRpcClient(config.C.Auth.HostName)
	return nil
}

