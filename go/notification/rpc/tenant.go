package rpc

import (
	"notification/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"sync"
	log "github.com/cihub/seelog"
	"google.golang.org/grpc/grpclog"

)


const Open_TLS = false
var (
	tenantOnce         sync.Once
	tenantRpcClient    api.TenantServerClient
	tenantConn      *grpc.ClientConn
)

func TenantClient() api.TenantServerClient {
	return tenantRpcClient
}

func NewTenantClient(hostname string) {
	tenantOnce.Do(func() {

			var opts []grpc.DialOption

			opts = append(opts, grpc.WithInsecure())

			// 使用自定义认证
			opts = append(opts, grpc.WithPerRPCCredentials(new(customCredentials)))

		    tenantConn, err = grpc.Dial(hostname, opts...)

			if err != nil {
				grpclog.Fatalln(err)
			}
		tenantRpcClient = api.NewTenantServerClient(tenantConn)
	})
}

// customCredential 自定义认证
type customCredentials struct{}

// GetRequestMetadata 实现自定义认证接口
func (c customCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	ctxmap := map[string]string{}
	if ctx.Value("source") != nil {
		ctxmap["source"] = ctx.Value("source").(string)
	}
	return ctxmap, nil
}

// RequireTransportSecurity 自定义认证是否开启TLS
func (c customCredentials) RequireTransportSecurity() bool {
	return Open_TLS
}


// 根据tid查找did
func  GetDidByTid(req *api.GetDidByTidRequest)(response *api.GetDidByTidResponse, err error)  {
	log.Infof("Start rpc...")
	response, err = TenantClient().GetDidByTid(context.Background(),req)
	return
}

func TenantRpcClientClose() {
	if tenantConn != nil {
		tenantConn.Close()
	}
}
