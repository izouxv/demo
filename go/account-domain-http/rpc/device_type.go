package  rpc

import (
	pb "account-domain-http/api"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"sync"
	"time"
)


var (
	dtOnce      sync.Once
	dtRpcClient pb.DeviceTypeClient
	dtConn      *grpc.ClientConn
)

func DevTypeRpcClient() pb.DeviceTypeClient {
	return dtRpcClient
}
func NewDevTypeRpcClient() {
	dtOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		dtConn, err = grpc.DialContext(ctx, RpcServer, grpc.WithInsecure())
		if err != nil {
			log.Error("DevTypeRpcInit Error:", err.Error())
			return
		}
		dtRpcClient = pb.NewDeviceTypeClient(dtConn)
	})
}

func DevTypeRpcClientClose() {
	if dtConn != nil {
		dtConn.Close()
	}
}
