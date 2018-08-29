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
	faqcommonOnce      sync.Once
	faqcommonRpcClient pb.FaqCommonClient
	faqConn *grpc.ClientConn
)

//初始化FaqCommonRpc
func FaqCommonRpcInit(address string) {
	faqcommonOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		faqConn, err = grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Error("FaqCommonRpcInit:", err)
			panic(err)
		}
		faqcommonRpcClient = pb.NewFaqCommonClient(faqConn)
	})
}

//结束Rpc
func FaqRpcClose() {
	if faqConn != nil {
		faqConn.Close()
	}
}

//调用FaqCommonRpc
func FaqCommonRpc(faqcommon *pb.FaqCommonRequest, method string) *pb.FaqCommonReply {
	log.Info("FaqCommonRpc-faqcommon:", faqcommon)
	var err error
	var reply *pb.FaqCommonReply
	switch method {
	//id查询常见问题
	case "GetFaqCommonById":
		reply, err = faqcommonRpcClient.GetFaqCommonById(context.Background(), faqcommon)
	default:
		err = errors.New("没有该RPC")
	}
	if err != nil {
		log.Error("FaqCommonRpc-err", err)
		reply.Code = 10001
	}
	log.Info("FaqCommonRpc-reply:", reply)
	return reply
}

//调用FaqCommonRpc
func FaqCommonsRpc(faqcommon *pb.FaqCommonRequest, method string) *pb.FaqCommonsReply {
	log.Info("FaqCommonsRpc-faqcommon:", faqcommon)
	var err error
	var faqCommonsReply *pb.FaqCommonsReply
	switch method {
	//Keyword查询常见问题
	case "GetFaqCommonByKeyword":
		faqCommonsReply, err = faqcommonRpcClient.GetFaqCommonByKeyword(context.Background(), faqcommon)
	//批量查询常见问题
	case "GetFaqCommons":
		faqCommonsReply, err = faqcommonRpcClient.GetFaqCommons(context.Background(), faqcommon)
	default:
		err = errors.New("没有该RPC")
	}
	if err != nil {
		log.Error("FaqCommonsRpc-Error", err)
		faqCommonsReply.Code = 10001
	}
	log.Info("FaqCommonsRpc-faqCommonsReply:", len(faqCommonsReply.Faqcs))
	return faqCommonsReply
}
