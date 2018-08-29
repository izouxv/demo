package rpc

import (
	"auth-http/api"

	"google.golang.org/grpc"
	"sync"
	"google.golang.org/grpc/grpclog"
)

var (
	serviceOnce         sync.Once
	serviceRpcClient    api.ServiceServerClient
)

func ServiceRpcClient() api.ServiceServerClient {
	return serviceRpcClient
}
func NewServiceRpcClient(hostname string) {
	serviceOnce.Do(func() {
		var err error
		var opts []grpc.DialOption

		opts = append(opts, grpc.WithInsecure())

		// 使用自定义认证
		opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

		conn, err := grpc.Dial(hostname, opts...)

		if err != nil {
			grpclog.Fatalln(err)
		}
		serviceRpcClient = api.NewServiceServerClient(conn)
	})
}