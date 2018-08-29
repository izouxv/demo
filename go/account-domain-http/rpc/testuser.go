package rpc

import (
	pb "account-domain-http/api/user"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"sync"
	"time"
)

var (
	testUserOnce       sync.Once
	testUserRpcClient  pb.TestUserServerClient
	testUserConn      *grpc.ClientConn
)

//rpc-role
func TestUserRpcClient() pb.TestUserServerClient {
	return testUserRpcClient
}
func NewTeetUserRpcClient() {
	testUserOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		testUserConn, err = grpc.DialContext(ctx, RpcServer, grpc.WithInsecure())
		if err != nil {
			log.Error(err)
			return
		}
		testUserRpcClient = pb.NewTestUserServerClient(testUserConn)
	})
}

func TestUserRpcClientClose() {
	if testUserConn != nil {
		testUserConn.Close()
	}
}

