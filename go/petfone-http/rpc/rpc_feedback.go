package rpc

import (
	"context"
	log "github.com/cihub/seelog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"petfone-http/pb/feedback"
	"sync"
)

var (
	fdOnce   sync.Once
	feedBackClient feedback.FeedBackClient
	fdConn *grpc.ClientConn
)

//FeedBackRpc
func FeedBackRpcInit(address string) {
	fdOnce.Do(func() {
		fdConn, err = grpc.DialContext(otherCtx, address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Error("FeedBackRpcInit err:", err)
			return
		}
		feedBackClient = feedback.NewFeedBackClient(fdConn)
	})
}

//结束Rpc
func FeedBackRpcClose() {
	if fdConn != nil {
		fdConn.Close()
	}
}

//FeedBackRpc
func FeedBackRpc(feedBack *feedback.AddFeedbackRequest, method string) *feedback.AddFeedbackReply {
	log.Info("FeedBackRpc-feedBack:", feedBack)
	var err error
	var reply *feedback.AddFeedbackReply
	if feedBackClient == nil {
		log.Error("FeedBackRpc feedBackClient nil")
		reply.ErrorCode = 10001
		return reply
	}
	switch method {
	//上报
	case "AddFeedback":
		reply, err = feedBackClient.AddFeedback(context.Background(), feedBack)
	default:
		err = errors.New("没有该RPC")
	}
	if err != nil {
		log.Error("FeedBackRpc-Error", err)
		reply.ErrorCode = 10001
	}
	log.Info("FeedBackRpc-reply:", reply)
	return reply
}
