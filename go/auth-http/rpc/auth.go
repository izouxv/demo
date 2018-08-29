package rpc

import (
	"auth-http/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"sync"
	"google.golang.org/grpc/grpclog"
)

const (
	// OpenTLS 是否开启TLS认证
	OpenTLS = false
)

var (
	authOnce         sync.Once
	authRpcClient    api.AuthServerClient
)

func AuthRpcClient() api.AuthServerClient {
	return authRpcClient
}

func NewAuthRpcClient(hostname string) {
	authOnce.Do(func() {
		var err error
		var opts []grpc.DialOption

		opts = append(opts, grpc.WithInsecure())

		// 使用自定义认证
		opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

		conn, err := grpc.Dial(hostname, opts...)

		if err != nil {
			grpclog.Fatalln(err)
		}
		authRpcClient = api.NewAuthServerClient(conn)
	})
}


// customCredential 自定义认证
type customCredential struct{}

// GetRequestMetadata 实现自定义认证接口
func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	ctxmap := map[string]string{}
	if ctx.Value("opt") != nil {
		ctxmap["opt"] = ctx.Value("opt").(string)
	}
	if ctx.Value("token") != nil {
		ctxmap["token"] = ctx.Value("token").(string)
	}
	if ctx.Value("url") != nil {
		ctxmap["url"] = ctx.Value("url").(string)
	}
	if ctx.Value("tid") != nil {
		ctxmap["tid"] = ctx.Value("tid").(string)
	}
	if ctx.Value("did") != nil {
		ctxmap["did"] = ctx.Value("did").(string)
	}
	return ctxmap, nil
}

// RequireTransportSecurity 自定义认证是否开启TLS
func (c customCredential) RequireTransportSecurity() bool {
	return OpenTLS
}