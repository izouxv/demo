package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"mynotes/grpcTest/api"
	"fmt"
)

var (
	// Address gRPC服务地址
	CAddress = "127.0.0.1:50052"

	// OpenTLS 是否开启TLS认证
	OpenTLS1 = true
)

// customCredential 自定义认证
type customCredential struct{}

// GetRequestMetadata 实现自定义认证接口
func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	fmt.Println(uri)
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
	}, nil
}

//RequireTransportSecurity 自定义认证是否开启TLS
func (c customCredential) RequireTransportSecurity() bool {
	return false
}

func main() {
	var err error
	var opts []grpc.DialOption
	if false {
		// TLS连接
		creds, err := credentials.NewClientTLSFromFile("../../keys/server.pem", "server name")
		if err != nil {
			fmt.Println("Failed to create TLS credentials:", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	// 使用自定义认证
	fmt.Println("------")
	aa := new(customCredential)
	fmt.Println("aa:",aa)
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))
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
