package rpc

import (
	pb "account-domain-http/api/adv"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"sync"
	"time"
)

const (
	RpcServer = "rpc.domain.com:7002"
)

var (
	advOnce      sync.Once
	advRpcClient pb.AdvertisementClient
	advConn      *grpc.ClientConn
)

func AdvRpcClient() pb.AdvertisementClient {
	return advRpcClient
}
func NewAdvRpcClient() {
	advOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		advConn, err = grpc.DialContext(ctx, RpcServer, grpc.WithInsecure())
		if err != nil {
			log.Error("AdvRpcInit Error:", err.Error())
			return
		}
		advRpcClient = pb.NewAdvertisementClient(advConn)
	})
}

func AdvRpcClientClose() {
	if advConn != nil {
		advConn.Close()
	}
}
