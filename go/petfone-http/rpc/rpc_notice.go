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
	noticeOnce      sync.Once
	noticeRpcClient pb.NoticeClient
	noticeConn *grpc.ClientConn
)

//初始化noticeRpc
func NoticeRpcInit(address string) {
	noticeOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		noticeConn, err = grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Error("PetInfoRpcInit:", err)
			panic(err)
		}
		noticeRpcClient = pb.NewNoticeClient(noticeConn)
	})
}

//结束Rpc
func NoticeRpcClose() {
	if noticeConn != nil {
		noticeConn.Close()
	}
}

//调用NoticeRpc
func NoticeRpc(petinfo *pb.NoticeRequest, method string) *pb.NoticeReply {
	log.Info("PetinfoRpc-petinfo:", petinfo)
	var noticeErr error
	var noticeReply *pb.NoticeReply
	switch method {
	//添加信息
	case "SetNotice":
		noticeReply, noticeErr = noticeRpcClient.SetNotice(context.Background(), petinfo)
		//删除信息
	case "DeleteNotice":
		noticeReply, noticeErr = noticeRpcClient.DeleteNotice(context.Background(), petinfo)
		//修改信息
	case "UpdateNotice":
		noticeReply, noticeErr = noticeRpcClient.UpdateNotice(context.Background(), petinfo)
	default:
		noticeErr = errors.New("没有该RPC")
	}
	if noticeErr != nil {
		log.Error("PetinfoRpc-Error", noticeErr)
		noticeReply.Code = 10001
	}
	log.Info("PetinfoRpc-noticeReply:", noticeReply)
	return noticeReply
}

//调用PetinfoRpc
func NoticesRpc(notice *pb.NoticeRequest, method string) *pb.NoticeMapReply {
	log.Info("NoticesRpc-notice:", notice)
	var err error
	var reply *pb.NoticeMapReply
	switch method {
	//查询信息
	case "GetNotice":
		reply, err = noticeRpcClient.GetNotice(context.Background(), notice)
		break
	default:
		err = errors.New("没有该RPC")
	}
	if err != nil {
		log.Error("NoticesRpc-Error", err)
		reply.Code = 10001
	}
	log.Info("NoticesRpc-reply:", reply)
	return reply
}
