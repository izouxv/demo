package rpc

import (
	"auth-http/api"
	"google.golang.org/grpc"
	"sync"
	"google.golang.org/grpc/grpclog"
)

var (
	tradingOnce         sync.Once
	tradingRpcClient    api.TradingServerClient
)

func TradingRpcClient() api.TradingServerClient {
	return tradingRpcClient
}
func NewTradingRpcClient(hostname string) {
	tradingOnce.Do(func() {
		var err error
		var opts []grpc.DialOption

		opts = append(opts, grpc.WithInsecure())

		// 使用自定义认证
		opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

		conn, err := grpc.Dial(hostname, opts...)

		if err != nil {
			grpclog.Fatalln(err)
		}
		tradingRpcClient = api.NewTradingServerClient(conn)
	})
}