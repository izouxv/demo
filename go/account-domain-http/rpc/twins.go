package rpc


import (
	pb "account-domain-http/api"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"sync"
	"time"
)

const (
	RpcTwinsServer = "rpc.twins.com:7011"
)

var (
	twinsOnce           sync.Once
	twinsRpcClient      pb.TwinsServerClient
	twinsConn          *grpc.ClientConn
)

func TwinsRpcClient() pb.TwinsServerClient {
	return twinsRpcClient
}
func NewTwinsRpcClient() {
	twinsOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		twinsConn, err = grpc.DialContext(ctx, RpcTwinsServer, grpc.WithInsecure())
		if err != nil {
			log.Error("AssetRpcInit Error:", err.Error())
			return
		}
		twinsRpcClient = pb.NewTwinsServerClient(twinsConn)
	})
}

func TwinsRpcClientClose() {
	if twinsConn != nil {
		twinsConn.Close()
	}
}


// 基于租户模糊获取资产列表
func GetTwinssBaseTenant(req *pb.GetTwinsBaseTenantRequest) (response *pb.GetTwinsBaseTenantResponse, err error) {
	log.Infof("Start rpc_twins GetTwinssBaseTenant: req(%#v)",req)
	response, err = TwinsRpcClient().GetTwinsBaseTenant(context.Background(),req)
	return
}
