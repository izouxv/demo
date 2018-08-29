package rpc

import (
	"context"
	log "github.com/cihub/seelog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"petfone-http/pb"
	"sync"
	"time"
)

var (
	shareOnce   sync.Once
	shareClient pb.ShareManageClient
	shareConn *grpc.ClientConn
)

//初始化ShareManageRpc
func ShareRpcInit(address string) {
	shareOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		shareConn, err = grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Error("ShareRpcInit:", err)
			panic(err)
		}
		shareClient = pb.NewShareManageClient(shareConn)
	})
}

//结束Rpc
func ShareRpcClose() {
	if shareConn != nil {
		shareConn.Close()
	}
}

//ShareManage
func ShareRpc(share *pb.ShareRequest, method string) *pb.ShareReply {
	log.Info("ShareRpc-share:", share)
	var err error
	var reply *pb.ShareReply
	switch method {
	//添加
	case "SetShare":
		reply, err = shareClient.SetShare(context.Background(), share)
		//删除
	case "DeleteShare":
		reply, err = shareClient.DeleteShare(context.Background(), share)
	default:
		err = errors.New("没有该RPC")
	}
	if err != nil {
		log.Error("ShareRpc-Error", err)
		reply.Code = 10001
	}
	log.Info("ShareRpc-reply:", reply)
	return reply
}

//ShareManages
func SharesRpc(share *pb.ShareRequest, method string) *pb.ShareMapReply {
	log.Info("ShareRpc-share:", share)
	var err error
	var reply *pb.ShareMapReply
	switch method {
	//获取
	case "GetShare":
		reply, err = shareClient.GetShare(context.Background(), share)
	default:
		err = errors.New("没有该RPC")
	}
	if err != nil {
		log.Error("PetinfoRpc错误日志:", err)
		reply.Code = 10001
	}
	log.Info("PetinfoRpc-shareMapReply:", len(reply.Shares))
	return reply
}
