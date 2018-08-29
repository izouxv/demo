package rpc

import (
	"context"
	log "github.com/cihub/seelog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"petfone-http/pb/setting"
	"sync"
)

var (
	vsOnce   sync.Once
	versionClient setting.RadacatVersionClient
	vsConn *grpc.ClientConn
)

//AdverRpcInit
func VersionRpcInit(address string) {
	vsOnce.Do(func() {
		vsConn, err = grpc.DialContext(otherCtx, address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Error("VersionRpcInit err:", err)
			return
		}
		versionClient = setting.NewRadacatVersionClient(vsConn)
	})
}

//结束Rpc
func VersionRpcClose() {
	if vsConn != nil {
		vsConn.Close()
	}
}

//VersionRpc
func VersionRpc(version *setting.GetLatestVersionRequest, method string) *setting.GetLatestVersionResponse {
	log.Info("VersionRpc-version:", version)
	var err error
	var reply *setting.GetLatestVersionResponse
	if versionClient == nil {
		log.Error("VersionRpc versionClient nil")
		reply.ErrorCode = 10001
		return reply
	}
	switch method {
	//获取
	case "GetLatestVersion":
		reply, err = versionClient.GetLatestVersion(context.Background(), version)
	default:
		err = errors.New("没有该RPC")
	}
	if err != nil {
		log.Error("VersionRpc err:", err)
		reply.ErrorCode = 10001
	}
	log.Info("VersionRpc-reply:", reply)
	return reply
}
