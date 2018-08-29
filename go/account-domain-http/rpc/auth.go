package rpc

import (
	api "account-domain-http/api/auth/api"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"sync"
	"google.golang.org/grpc/grpclog"
)

var  (
	ipServer = "120.76.54.242:7010"
	OpenTLS = false
)

var (
	authOnce           sync.Once
	authRpcClient      api.AuthServerClient
	authConn          *grpc.ClientConn
)

func AuthRpcClient() api.AuthServerClient {
	return authRpcClient
}

func NewAuthRpcClient() {
	authOnce.Do(func() {

		var err error
		var opts []grpc.DialOption

		opts = append(opts, grpc.WithInsecure())

		// 使用自定义认证
		opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

		authConn, err = grpc.Dial(ipServer, opts...)

		if err != nil {
			grpclog.Fatalln(err)
		}

		/*ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		conn, err := grpc.DialContext(ctx, ipServer, grpc.WithInsecure())
		if err != nil {
			log.Error(err)
			return
		}*/
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
	response,err = AuthRpcClient().Authorization(context.Background(),req)
	return
}


func AuthRpcClientClose() {
	if authConn != nil {
		authConn.Close()
	}
}

