package rpcClient

import (
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"cotx-http/pb"
	"sync"
	"time"
)

var (
	accountOnce      sync.Once
	accountRpcClient pb.AccountClient
)

const (
	//accountServer = "rpc.account.radacat.com:8003"
	accountServer = "localhost:7007"
)

func NewAccountRpcClient() {
	accountOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		conn, err := grpc.DialContext(ctx, accountServer, grpc.WithInsecure())
		if err != nil {
			log.Error("AccountRpcInit Error:", err)
			return
		}
		accountRpcClient = pb.NewAccountClient(conn)
	})
}

func GetAccountRpcClient() pb.AccountClient {
	return accountRpcClient
}

//获取用户信息
func GetUserInfoAll(account *pb.AccountRequest) *pb.AccountReply {
	log.Info("GetUserInfoRpc-account:",account)
	accountReply, _ := GetAccountRpcClient().GetUserInfoAll(context.Background(), account)
	log.Info("GetUserInfoRpc-accountReply:",accountReply)
	return accountReply
}

//修改用户信息
func UpdateUserInfo(account *pb.AccountRequest) *pb.AccountReply {
	log.Info("UpdateUserInfoRpc-account:",account)
	accountReply, _ := GetAccountRpcClient().UpdateAccountInfo(context.Background(), account)
	log.Info("UpdateUserInfoRpc-accountReply:",accountReply)
	return accountReply
}

