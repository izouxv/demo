package main

import (
	ctx1 "mynotes/blog/initContext"
	"mynotes/blog/logger"
	"mynotes/blog/router"

	"fmt"
	"os"
	"os/signal"
	"syscall"
	"runtime/debug"
)

func main() {
	fmt.Println("aaaaaa")
	defer PanicError()
	go waitSignal()
	ctx1.Run()
	logger.Info("main ListenAndServe error:",router.HttpServer.ListenAndServe())
}

func waitSignal()  {
	var signalChan = make(chan os.Signal, 1)
	signal.Notify(signalChan)
	select {
	case s := <-signalChan:
		if s == syscall.SIGPIPE {
			break
		}
		logger.Info("close conn...")
		ctx1.CloseConn()
		logger.Info("Go signal:", s)
	}
}

func PanicError() {
	if err := recover(); err != nil {
		fmt.Println("recover:", err)
		fmt.Println("Process name:", os.Args[0])
		fmt.Println("Process id:", os.Getpid())
		fmt.Println("Stack info:", string(debug.Stack()))
	}
}
