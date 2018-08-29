package rpc

import (
	"sync"
	"auth-http/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

)

var (
	actionLogOnce            sync.Once
	actionLogRpcClient       api.ActionLogClient
)

func ActionLogRpcClient() api.ActionLogClient {
	return actionLogRpcClient
}
func NewActionLogRpcClient(hostname string) {
	actionLogOnce.Do(func() {
		var err error
		var opts []grpc.DialOption

		opts = append(opts, grpc.WithInsecure())

		// 使用自定义认证
		opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

		conn, err := grpc.Dial(hostname, opts...)

		if err != nil {
			grpclog.Fatalln(err)
		}
		actionLogRpcClient = api.NewActionLogClient(conn)
	})
}
