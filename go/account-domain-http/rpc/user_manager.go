package rpc

import (
	"sync"
	"time"
	user "account-domain-http/api/user/pb"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const userServer = "rpc.account.radacat.com:8003"
//const userServer = "120.77.66.96:8003"
//const userServer = "192.168.1.178:8003"

var (
	mssoOnce       sync.Once
	muserRpcClient user.MSsoClient
	muserConn      *grpc.ClientConn
)

func MUserRpcClient() user.MSsoClient {
	return muserRpcClient
}
func NewMUserRpcClient() {
	mssoOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		muserConn, err = grpc.DialContext(ctx, userServer, grpc.WithInsecure())
		if err != nil || muserConn == nil {
			log.Error(err)
			return
		}
		muserRpcClient = user.NewMSsoClient(muserConn)
	})
}

func MuserRpcClientClose() {
	if muserConn != nil {
		muserConn.Close()
	}
}


