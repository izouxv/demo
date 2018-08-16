package cmd

import (
	log "github.com/cihub/seelog"
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"time"
	"net"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"notification/config"
	"notification/core"
	"notification/controller"
	"os"
	"os/signal"
	"notification/handler/receive"
	"notification/handler/mqtthandler"
	"notification/handler/multihandler"
	redis "notification/redis_handler"
	"notification/rpc"
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
		setWechatNotice,
		setMqttClient,
		handleReceivedPayloadsEmail,
		setServerClient,
		setServeHTTP,
	}
	for _, t := range tasks {
		if err := t(); err != nil {
			log.Info(err)
		}
	}
	sigChan := make(chan os.Signal)
	exitChan := make(chan struct{})
	signal.Notify(sigChan)
	log.Info("signal ", <-sigChan, "signal received")
	go func() {
		log.Info("stopping domain rpc -server")
		exitChan <- struct{}{}
	}()
	select {
	case <-exitChan:
	case s := <-sigChan:
		log.Info("signal received, stopping immediately signal ", s)
	}

	return nil
}

func setLogLevel() error {
	logger, err := log.LoggerFromConfigAsString(config.LogConfig)
	if err != nil {
		log.Errorf("err parsing config seelog file :%s", err)
		return err
	}
	log.ReplaceLogger(logger)
	return nil
}

func printStartMessage() error {
	log.Infof("starting Notification Server version:%s",version)
	return nil
}

func setMysqlSQLConnection() error {
	log.Info("connecting to mysql")
	dsn := config.C.MySQL.User+":"+config.C.MySQL.Password+"@tcp("+config.C.MySQL.Host+":"+config.C.MySQL.Port+")/"+config.C.MySQL.Database+"?charset=utf8&loc=Asia%2FShanghai&parseTime=true"
	db, err := core.OpenDatabase(dsn)
	if err != nil {
		return errors.Wrap(err, "database connection error")
	}
	config.C.MySQL.DB = db
	return nil
}

func setRedisPool() error {
	log.Info("setup redis connection pool")
	config.C.Redis.Pool = core.NewRedisPool(config.C.Redis.URL)
	return config.C.Redis.Pool.TestOnBorrow(config.C.Redis.Pool.Get(),time.Now())
}

func setWechatNotice() error {
	log.Info("启动微信通知")
	go redis.ReceiveRedisMess()
	return nil
}


func setServeHTTP () error{
	log.Info("notification server http ")
	router := httprouter.New()
	controller.RouterMethod(router)
	address := net.JoinHostPort(config.C.Notification.Host,config.C.Notification.Port)
	http := &http.Server{Addr: address, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second,
		IdleTimeout: 10 * time.Second, Handler: router,            //TLSConfig:tls.Config{},
	}
	log.Info(http.ListenAndServe())
	return nil
}

func setMqttClient()  error{
	log.Infof("StartMqttClient.....")
	h, err := mqtthandler.NewHandler(config.C.MQTT.Server,
		config.C.MQTT.Username, config.C.MQTT.Password,
		config.C.MQTT.CACert)
	if err != nil {
		log.Errorf("启动Mq服务异常 %s", err)
		return errors.Wrap(err, "setup mqtt handler error")
	}
	mqtthandler.MqttHandler = h
	setHandler("mqtt")
	return nil
}

func setHandler(h string) error {
	err := multihandler.NewMultiHandler(h)
	if err != nil {
		return errors.Wrap(err, "setup handler error")
	}
	return nil
}

func handleReceivedPayloadsEmail() error {
	log.Infof("Start Receive Mqtt notice/email ")
	go  received.HandleReceivedPayloadsEmail()
	return nil
}

func setServerClient() error {
	rpc.NewAuthClient(config.C.Auth.Hostname)
	rpc.NewTenantClient(config.C.Auth.Hostname)
	return nil
}


