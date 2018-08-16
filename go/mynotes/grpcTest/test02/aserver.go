package main

import (
	"fmt"
	"net"
	"mynotes/grpcTest/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"       // grpc 响应状态码
	"google.golang.org/grpc/metadata" // grpc metadata包
	"errors"
	"time"
)

const (
	// Address gRPC服务地址
	SAddress = "127.0.0.1:50052"
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService Hello服务
var HelloService = helloService{}

// SayHello 实现Hello服务接口
func (h helloService) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloResponse, error) {
	resp := new(api.HelloResponse)
	fmt.Println("SayHello:",ctx)
	fmt.Println("SayHello:",in)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", SAddress)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	// TLS认证
	//creds, err := credentials.NewServerTLSFromFile("../../keys/server.pem", "../../keys/server.key")
	//if err != nil {
	//	fmt.Printf("Failed to generate credentials %v", err)
	//}
	//opts = append(opts, grpc.Creds(creds))
	// 注册interceptor
	opts = append(opts, grpc.UnaryInterceptor(interceptor))
	// 实例化grpc Server
	s := grpc.NewServer(opts...)
	// 注册HelloService
	api.RegisterHelloServer(s, HelloService)
	fmt.Println("Listen on " + SAddress + " with TLS + Token + Interceptor")
	s.Serve(listen)
}

// auth 验证Token
func auth(ctx context.Context) error {
	fmt.Println("ctx---:",ctx)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New(codes.Unauthenticated.String()+ "无Token认证信息")
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
	time.Sleep(time.Second*1)
	if appid != "101010" || appkey != "i am key" {
		return errors.New(codes.Unauthenticated.String()+ "Token认证信息无效:"+ appid+ appkey)
	}
	return nil
}

// interceptor 拦截器
func interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Println("ctx:----")
	err := auth(ctx)
	if err != nil {
		return nil, err
	}
	// 继续处理请求
	return handler(ctx, req)
}