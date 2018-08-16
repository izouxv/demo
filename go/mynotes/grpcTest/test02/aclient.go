package main

import (
	"time"
	"mynotes/grpcTest/api" // 引入proto包
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"fmt"
	"google.golang.org/grpc/credentials"
)

const (
	// Address gRPC服务地址
	CAddress = "127.0.0.1:50052"

	// OpenTLS 是否开启TLS认证
	OpenTLS = false
)

// customCredential 自定义认证
type customCredential struct{}

// GetRequestMetadata 实现自定义认证接口
func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
	}, nil
}

// RequireTransportSecurity 自定义认证是否开启TLS
func (c customCredential) RequireTransportSecurity() bool {
	return OpenTLS
}

func main() {
	var err error
	var opts []grpc.DialOption
	if false {
		// TLS连接
		creds, err := credentials.NewClientTLSFromFile("../../keys/server.pem", "server name")
		if err != nil {
			fmt.Printf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	// 指定自定义认证
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))
	// 指定客户端interceptor
	opts = append(opts, grpc.WithUnaryInterceptor(cinterceptor))
	conn, err := grpc.Dial(CAddress, opts...)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	// 初始化客户端
	c := api.NewHelloClient(conn)
	// 调用方法
	req := &api.HelloRequest{Name: "gRPC"}
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Message)
}

// interceptor 客户端拦截器
func cinterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	fmt.Println("cinterceptor:",method)
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Printf("method=%s req=%v rep=%v duration=%s error=%v\n", method, req, reply, time.Since(start), err)
	return err
}