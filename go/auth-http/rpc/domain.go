package rpc

import (
	"auth-http/api"

	"google.golang.org/grpc"
	"sync"
	"google.golang.org/grpc/grpclog"
)

var (
	domainOnce         sync.Once
	domainRpcClient    api.DomainServerClient
)

func DomainRpcClient() api.DomainServerClient {
	return domainRpcClient
}
func NewDomainRpcClient(hostname string) {
	domainOnce.Do(func() {
		var err error
		var opts []grpc.DialOption

		opts = append(opts, grpc.WithInsecure())

		// 使用自定义认证
		opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

		conn, err := grpc.Dial(hostname, opts...)

		if err != nil {
			grpclog.Fatalln(err)
		}
		domainRpcClient = api.NewDomainServerClient(conn)
	})
}