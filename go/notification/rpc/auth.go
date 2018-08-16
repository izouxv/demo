package rpc

import (
	"notification/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"sync"
	"google.golang.org/grpc/grpclog"
	log "github.com/cihub/seelog"
)

const (
	ipServer = "rpc.dms.cotxnetworks.com:7010"
)
const OpenTLS = false

var (
	authOnce         sync.Once
	authRpcClient    api.AuthServerClient
	authConn      *grpc.ClientConn
	err                   error
)

func AuthClient() api.AuthServerClient {
	return authRpcClient
}

func NewAuthClient(hostname string) {
	authOnce.Do(func() {
		var err error
		var opts []grpc.DialOption

		opts = append(opts, grpc.WithInsecure())

		// 使用自定义认证
		opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

		authConn, err = grpc.Dial(hostname, opts...)

		if err != nil {
			grpclog.Fatalln(err)
		}
		authRpcClient = api.NewAuthServerClient(authConn)
	})
}

// customCredential 自定义认证
type customCredential struct{}

// GetRequestMetadata 实现自定义认证接口
func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	ctxmap := map[string]string{}
	if ctx.Value("source") != nil {
		ctxmap["source"] = ctx.Value("source").(string)
	}
	return ctxmap, nil
}

// RequireTransportSecurity 自定义认证是否开启TLS
func (c customCredential) RequireTransportSecurity() bool {
	return OpenTLS
}

func Authorization(req *api.AuthorizationRequest) (response *api.AuthorizationResponse, err error)  {
	log.Info("Start rpc_Authorization:",req)
	response,err = AuthClient().Authorization(context.Background(),req)
	return
}
func AuthRpcClientClose() {
	if authConn != nil {
		authConn.Close()

	}
}
