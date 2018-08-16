package main


import (
	"fmt"
	"net"
	"mynotes/grpcTest/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata" // 引入grpc meta包
	"errors"
)

const (
	// Address gRPC服务地址
	SAddress = "127.0.0.1:50052"
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService
var HelloService = helloService{}

//实现接口
func (h helloService) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloResponse, error) {
	// 解析metada中的信息并验证
	fmt.Println("ctx:",ctx)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New(codes.Unauthenticated.String()+ "无Token认证信息")
	}
	var (
		appid  string
		appkey string
	)
	if val, ok := md["appid"]; ok {
		appid = val[0]
	}
	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}
	if appid != "101010" || appkey != "i am key" {
		return nil, errors.New(codes.Unauthenticated.String()+ "Token认证信息无效:"+ appid+appkey)
	}
	resp := new(api.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.\nToken info: appid=%s,appkey=%s", in.Name, appid, appkey)
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", SAddress)
	if err != nil {
		errors.New("failed to listen: "+ err.Error())
	}
	// TLS认证
	//creds, err := credentials.NewServerTLSFromFile("../../keys/server.pem", "../../keys/server.key")
	//if err != nil {
	//	grpclog.Fatalf("Failed to generate credentials %v", err)
	//}grpc.Creds(creds)

	// 实例化grpc Server, --并开启TLS认证
	s := grpc.NewServer()

	// 注册HelloService
	api.RegisterHelloServer(s, HelloService)

	fmt.Println("Listen on " + SAddress + " with TLS + Token")

	s.Serve(listen)
}