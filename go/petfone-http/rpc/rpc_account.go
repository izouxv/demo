package rpc

import (
	"context"
	log "github.com/cihub/seelog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"petfone-http/pb"
	"time"
	"sync"
)

var (
	accountOnce      sync.Once
	accountClient pb.AccountClient
	accountConn *grpc.ClientConn
	err error
	ctx,cancel = context.WithTimeout(context.Background(), 30*time.Second)
	otherCtx, _ = context.WithTimeout(context.Background(), 5*time.Second)
)

//初始化ssoRpc
func AccountRpcInit(address string) {
	accountOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		accountConn, err = grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Error("AccountRpcInit-err:", err)
			panic(err)
		}
		accountClient = pb.NewAccountClient(accountConn)
	})
}

//结束Rpc
func AccountRpcClose() {
	if accountConn != nil {
		accountConn.Close()
	}
}

//调用account-rpc
func AccountRpc(account *pb.AccountRequest, method string) *pb.AccountReply {
	log.Info("AccountRpc-account:", account)
	var err error
	var accountReply *pb.AccountReply
	switch method {
	//获取用户所有信息
	case "GetAccountInfo":
		accountReply, err = accountClient.GetAccountInfo(context.Background(), account)
		//获取用户信息
	case "Show":
		accountReply, err = accountClient.Show(context.Background(), account)
		//修改用户信息全部
	case "UpdateAccountInfo":
		accountReply, err = accountClient.UpdateAccountInfo(context.Background(), account)
	default:
		err = errors.New("没有该RPC")
	}
	if err != nil {
		log.Error("AccountRpc-Error", err)
		return &pb.AccountReply{Code:10001}
	}
	log.Info("AccountRpc-accountReply:", accountReply)
	return accountReply
}
