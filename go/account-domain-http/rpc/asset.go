package rpc

import (
	pb  "account-domain-http/api"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"sync"
	"time"
)

const (
	RpcAssetServer = "rpc.asset.com:7012"
)

var (
	assetOnce           sync.Once
	assetRpcClient      pb.AssetServerClient
	assetConn          *grpc.ClientConn
)

func AssetRpcClient() pb.AssetServerClient {
	return assetRpcClient
}
func NewAssetRpcClient() {
	assetOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		assetConn, err = grpc.DialContext(ctx, RpcAssetServer, grpc.WithInsecure())
		if err != nil {
			log.Error("AssetRpcInit Error:", err.Error())
			return
		}
		assetRpcClient = pb.NewAssetServerClient(assetConn)
	})
}

func AssetRpcClientClose() {
	if assetConn != nil {
		assetConn.Close()
	}
}


// 基于租户模糊获取资产列表
func GetAssetsForKeywordMoreBaseTenant(req *pb.GetAssetsForKeywordMoreBaseTenantRequest) (response *pb.GetAssetsForKeywordMoreBaseTenantResponse, err error) {
	log.Infof("Start rpc_asset GetAssetsForKeywordMoreBaseTenant: req(%#v)",req)
	response, err = AssetRpcClient().GetAssetsForKeywordMoreBaseTenant(context.Background(),req)
	return
}


// 基于租户模糊获取资产列表
func GetAssetsForKeywordBaseTenant(req *pb.GetAssetsForKeywordBaseTenantRequest) (response *pb.GetAssetsForKeywordBaseTenantResponse, err error) {
	log.Infof("Start rpc_asset GetAssetsForKeywordBaseTenant: req(%#v)",req)
	response, err = AssetRpcClient().GetAssetsForKeywordBaseTenant(context.Background(),req)
	return
}



// 基于租户获取资产列表
func GetAssetBaseTenant(req *pb.GetAssetBaseTenantRequest) (response *pb.GetAssetBaseTenantResponse, err error)  {
	log.Infof("Start rpc_asset GetAssetBaseTenant: req(%#v)",req)
	response, err = AssetRpcClient().GetAssetBaseTenant(context.Background(),req)
	return
}


