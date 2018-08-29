package rpcClient

import (
	"cotx-http/pb"
	"context"
	"time"
	"google.golang.org/grpc"
	"github.com/prometheus/common/log"
)

var feedbackClient  pb.FeedbackServiceClient

func NewFeedbackClient()  {
	ctx,_ := context.WithTimeout(context.Background(),10*time.Second)
	conn,err := grpc.DialContext(ctx,address,grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return 
	}
	feedbackClient = pb.NewFeedbackServiceClient(conn)
}

func GetFeedbackClient()pb.FeedbackServiceClient  {
	return feedbackClient
}

