package rpc

import (

	"google.golang.org/grpc"
	"sync"
	"auth-http/api"
	"google.golang.org/grpc/grpclog"
)

var (
	roleOnce            sync.Once
	roleRpcClient       api.RoleServerClient
)
//rpc-role
func RoleRpcClient() api.RoleServerClient {
	return roleRpcClient
}
func NewRoleClient(hostname string) {
	roleOnce.Do(func() {
		var err error
		var opts []grpc.DialOption

		opts = append(opts, grpc.WithInsecure())

		// 使用自定义认证
		opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

		conn, err := grpc.Dial(hostname, opts...)

		if err != nil {
			grpclog.Fatalln(err)
		}
		roleRpcClient = api.NewRoleServerClient(conn)
	})
}
