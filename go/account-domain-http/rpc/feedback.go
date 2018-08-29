package rpc

import (
	pb "account-domain-http/api/feedback"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"sync"
	"time"
)

var (
	feedbackOnce      sync.Once
	feedbackRpcClient pb.FeedBackClient
	feedbackConn      *grpc.ClientConn
)

//rpc-role
func FeedbackRpcClient() pb.FeedBackClient {
	return feedbackRpcClient
}
func NewFeedbackRpcClient() {
	feedbackOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		feedbackConn, err = grpc.DialContext(ctx, RpcServer, grpc.WithInsecure())
		if err != nil {
			log.Error(err)
			return
		}
		feedbackRpcClient = pb.NewFeedBackClient(feedbackConn)
	})
}

func FeedbackRpcClientClose() {
	if feedbackConn != nil {
		feedbackConn.Close()
	}
}
