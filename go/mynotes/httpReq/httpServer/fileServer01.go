package main

import (
	"os"
	"os/signal"
	"github.com/julienschmidt/httprouter"
	"github.com/coreos/etcd/clientv3"
	"net"
	"net/http"
	"time"
	"fmt"
	"io"
	"context"
)

var fileDir = http.FileServer(http.Dir("xx/"))

func main() {
	//signal.notify方法用来监听所有收到的信号
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan)
	router := httprouter.New()
	router.GET("/test", Test)
	router.ServeFiles("/files/*filepath", http.Dir("xx/"))
	address := net.JoinHostPort("", "7007")
	htp := &http.Server{Addr:address, ReadTimeout:  10 * time.Second, WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,Handler: router,
	}
	go func() {
		for s := range signalChan {
			htp.Close()
			fmt.Println("s:",s)
		}
	}()
	go func() {
		cli,err := clientv3.New(
			clientv3.Config{
				Endpoints:[]string{"localhost:2379"},
				DialTimeout:5*time.Second,
			})
		if err != nil {
			fmt.Println("connect failed, err:", err)
			return
		}
		fmt.Println("connect succ")
		defer cli.Close()
		//设置1秒超时，访问etcd有超时控制
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		//操作etcd
		_, err = cli.Put(ctx, "/logagent/conf/", "sample_value")
		//操作完毕，取消etcd
		cancel()
		if err != nil {
			fmt.Println("put failed, err:", err)
			return
		}
		//取值，设置超时为1秒
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		resp, err := cli.Get(ctx, "/logagent/conf/")
		cancel()
		if err != nil {
			fmt.Println("get failed, err:", err)
			return
		}
		for _, ev := range resp.Kvs {
			fmt.Printf("%s : %s\n", ev.Key, ev.Value)
		}
	}()
	htp.ListenAndServe()
}

func Test01(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	fmt.Println("Test01-------")
	http.StripPrefix("/dir/", fileDir).ServeHTTP(res, req)
}

func Test(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	fmt.Println("Test-----",req.URL.Path)
	io.Copy(res,req.Body)
}

