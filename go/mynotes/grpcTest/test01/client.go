package main


import (
	"google.golang.org/grpc"
	"mynotes/grpcTest/api"
	"fmt"
	"sync"
	"github.com/silenceper/pool"
	"time"
	"io"
	"context"
)

const (
	address = "127.0.0.1:10023"
)

var wg = &sync.WaitGroup{}

func main(){
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %v", err)
	}
	defer conn.Close()

	// 创建连接
	factory := func() (interface{}, error) {
		return api.NewDataClient(conn),nil
	}

	//初始化链接池
	p,err := InitThread(5,30,factory,func(v interface{}) error { return conn.Close()})
	if err != nil{
		fmt.Println("init error")
		return
	}
	//todo string
	wg.Add(1)
	go func(){
		defer wg.Done()
		//获取连接
		v,_ := p.Get()
		rpcClient := v.(api.DataClient)
		info := &api.Request{Test:"test"}
		Test01(rpcClient,info)
		//归还链接
		p.Put(v)
	}()
	wg.Wait()

	//todo stream
	wg.Add(1)
	go func(){
		defer wg.Done()
		//获取连接
		v,_ := p.Get()
		rpcClient := v.(api.DataClient)
		Test02(rpcClient)
		//归还链接
		p.Put(v)
	}()
	wg.Wait()
	//获取链接池大小
	current := p.Len()
	fmt.Println("len=", current)
}


//简单模式
func Test01(client api.DataClient, req *api.Request)  {
	res, err := client.Test01(context.Background(),req)
	if err != nil {
		fmt.Printf("Could not create Customer: %v", err)
	}
	fmt.Println("Test01:",res.Test)
}

//双向流模式
func Test02(client api.DataClient){
	notes := []*api.Response{{Test:"jim"}, {Test:"Tom"}}
	stream, err := client.Test02(context.Background())
	if err != nil {
		fmt.Printf("%v.RouteChat(_) = _, %v", client, err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				// read done.
				fmt.Println("read done ")
				close(waitc)
				return
			}
			if err != nil {
				fmt.Printf("Failed to receive a note : %v", err)
			}
			fmt.Printf("Test02: %v",res.Test)
		}
	}()
	fmt.Println("notes",notes)
	for _, note := range notes {
		if err := stream.Send(note); err != nil {
			fmt.Printf("Failed to send a note: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
}

/*
    初始化
    min // 最小链接数
    max // 最大链接数
    factory func() (interface{}, error) //创建链接的方法
    close func(v interface{}) error //关闭链接的方法
*/
func InitThread(min,max int,factory func()(interface{}, error), close func(v interface{}) error)(pool.Pool,error){
	poolConfig := &pool.PoolConfig{
		InitialCap: min,
		MaxCap:     max,
		Factory:    factory,
		Close:      close,
		//链接最大空闲时间，超过该时间的链接 将会关闭，可避免空闲时链接EOF，自动失效的问题
		IdleTimeout: 15 * time.Second,
	}
	p, err := pool.NewChannelPool(poolConfig)
	if err != nil {
		fmt.Println("Init err=", err)
		return nil,err
	}
	return p,nil
}