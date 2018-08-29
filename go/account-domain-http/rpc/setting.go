package rpc

import (
	pb "account-domain-http/api/setting"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"sync"
	"time"
)

var (
	radacatVersionOnce      sync.Once
	radacatVersionRpcClient pb.RadacatVersionClient
	radacatVersionConn      *grpc.ClientConn

)

func RadacatVersionRpcClient() pb.RadacatVersionClient {
	return radacatVersionRpcClient
}
func NewRadacatVersionRpcClient() {
	radacatVersionOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		radacatVersionConn, err = grpc.DialContext(ctx, RpcServer, grpc.WithInsecure())
		if err != nil {
			log.Error(err)
			return
		}
		radacatVersionRpcClient = pb.NewRadacatVersionClient(radacatVersionConn)
	})
}

func RadacatVersionRpcClientClose() {
	if radacatVersionConn != nil {
		radacatVersionConn.Close()
	}
}
