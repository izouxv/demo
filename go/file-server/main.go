package main

import (
	"context"
	pb "file-server/api"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"file-server/core"
	"file-server/internal/api"
	. "file-server/router"
	log "github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"strings"

	"file-server/common"
)

var version string

func run(c *cli.Context) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	fmt.Println("running ......")
	tasks := []func(*cli.Context) error{
		setLogLevel,
		setConfig,
		setHttpServer,
		setUploadDir,
		setServerAddress,
		startApplicationServerAPI,
	}

	for _, t := range tasks {
		if err := t(c); err != nil {
			panic(err)
		}
	}

	sigChan := make(chan os.Signal)
	exitChan := make(chan struct{})
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	log.Info("signal", <-sigChan, "signal received")
	go func() {
		log.Warn("stopping ......")
		exitChan <- struct{}{}
	}()
	select {
	case <-exitChan:
	case s := <-sigChan:
		log.Info("signal", s, "signal received, stopping immediately")
	}

	return nil
}

func setLogLevel(c *cli.Context) error {
	logger, err := log.LoggerFromConfigAsFile(c.String("seeLog"))
	if err != nil {
		log.Critical("初始化日志错误", err)
	}
	log.ReplaceLogger(logger)
	return nil
}

func setConfig(c *cli.Context) error {
	core.ContextInit(c.String("config"))
	return nil
}

func setHttpServer(c *cli.Context) error {
	bindParts := strings.SplitN(c.String("http-bind"), ":", 2)
	log.Infof("http bind port is %s", bindParts[1])
	if len(bindParts) != 2 {
		return errors.New("http-bind is error e.g. [0.0.0.0:8080]")
	}
	router := httprouter.New()
	NewRouter(router)
	interceptor := make(Interceptor)

	address := net.JoinHostPort(bindParts[0], bindParts[1])
	interceptor[address] = router
	go func() {
		err := http.ListenAndServe(address, router)
		if err != nil {
			log.Errorf("start http error %s", err)
		}
	}()
	return nil
}

type Interceptor map[string]http.Handler

func (hs Interceptor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler := hs[r.Host]; handler != nil {
		fmt.Println("r.RequestURI :", r.RequestURI)
	} else {
		http.Error(w, "Forbidden", 403)
	}
}

func setUploadDir(c *cli.Context) error {
	common.UploadDir = c.String("dir")
	return nil
}

func setServerAddress(c *cli.Context) error {
	common.ServerAddress = "http://" + c.String("domain-name") + ":88"
	return nil
}

func startApplicationServerAPI(c *cli.Context) error {
	apiServer := mustGetAPIServer()
	log.Infof("启动rpc服务监听地址 %s", c.String("bind"))
	ln, err := net.Listen("tcp", c.String("bind"))
	if err != nil {
		log.Errorf("启动rpc服务监听端口异常 : %s", err)
		return err
	}
	if err == nil {
		/*启动rpc服务*/
		log.Info("启动rpc服务")
		go apiServer.Serve(ln)
	}
	return nil
}

func mustGetAPIServer() *grpc.Server {
	var opts []grpc.ServerOption
	gs := grpc.NewServer(opts...)
	log.Info("注册服务")
	pb.RegisterFileServerServer(gs, &api.FileServer{})
	return gs
}
func main() {
	app := cli.NewApp()
	app.Name = "file-server"
	app.Usage = "file-server for O-Cloud"
	app.Version = version
	app.Copyright = ""
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config",
			Usage:  "config",
			Value:  "config/conf.yaml",
			EnvVar: "config",
		},
		cli.StringFlag{
			Name:   "seeLog",
			Usage:  "log",
			Value:  "config/seeLog.xml",
			EnvVar: "seeLog",
		},
		cli.StringFlag{
			Name:   "http-bind",
			Usage:  "ip:port to bind the (user facing) http server to (web-interface and REST / gRPC api)",
			Value:  "0.0.0.0:8080",
			EnvVar: "HTTP_BIND",
		},
		cli.StringFlag{
			Name:   "dir",
			Usage:  "dir save file",
			Value:  "./upload/",
			EnvVar: "DIR",
		},
		cli.StringFlag{
			Name:   "domain-name",
			Usage:  "domain-name e.g. upload.radacat.com",
			Value:  "file.radacat.com",
			EnvVar: "DOMAIN_NAME",
		},
		cli.StringFlag{
			Name:   "bind",
			Usage:  "ip:port to bind the api server",
			Value:  "0.0.0.0:7018",
			EnvVar: "BIND",
		},
	}
	app.Run(os.Args)
}
