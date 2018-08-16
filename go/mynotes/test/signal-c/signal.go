package main

import (
	"os"
	"sync"
	"fmt"
	"time"
	"os/signal"
	"github.com/julienschmidt/httprouter"
	"net/http"
	log "github.com/cihub/seelog"
	"net"
)

var c chan os.Signal
var msgQueue chan *string
var wg sync.WaitGroup


//监听程序运行信号,ctrl+*
func main(){
	test1()
	test2()
}

func test1()  {
	c = make(chan os.Signal, 1)
	msgQueue = make(chan *string, 10)
	signal.Notify(c, os.Interrupt, os.Kill)
	//pruducer
	wg.Add(1)
	go Producer()
	//consumer
	wg.Add(1)
	go Consumer()
	wg.Wait()
}

func test2(){
	signalChan := make(chan os.Signal, 1)
	//cleanupDone := make(chan bool)
	signal.Notify(signalChan)
	fmt.Println("11111111")
	router := httprouter.New()
	router.POST("/", test)
	router.ServeFiles("/swagger/*filepath", http.Dir("e:/var"))
	//服务器监听主机地址与端口号
	address := net.JoinHostPort("127.0.0.1", "8080")
	htp := &http.Server{Addr:address, ReadTimeout:  10 * time.Second, WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,Handler: router, //TLSConfig		: tls.Config{},
	}
	log.Info("22222")
	go func() {
		log.Info("222333222:")
		err := htp.ListenAndServe()
		log.Info("3333333:",err)
		//for _ = range signalChan {
		//	fmt.Println("收到终端信号，停止服务... ")
		//	cleanup()
		//	cleanupDone <- true
		//
		//}
	}()
	log.Info("4444444")
	s := <-signalChan
	fmt.Println("Go signal:", s)
}

func cleanup() {
	fmt.Println("清理...")
}

func test(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	go func() {
		fmt.Println("test....")
	}()
	fmt.Fprint(res, "test")
}

func Producer(){
	i := 0
LOOP:
	for{
		select {
		case s := <-c:
			fmt.Println()
			fmt.Println("Producer | get", s)
			break LOOP
		}

		i ++
		s := fmt.Sprintf("work-%d", i)
		fmt.Println("Producer | produce", s)
		msgQueue <- &s
		time.Sleep(500 * time.Millisecond)
	}

	close(msgQueue)
	fmt.Println("Producer | close channel, exit")
	wg.Done()
}

func Consumer(){
	for m := range msgQueue{
		if m != nil{
			fmt.Println("Consumer | consume", *m)
		}else{
			fmt.Println("Consumer | channel closed")
			break
		}
	}
	fmt.Println("Consumer | exit")
	wg.Done()
}