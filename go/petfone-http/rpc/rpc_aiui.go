package rpc

import (
	log "github.com/cihub/seelog"
	"petfone-http/pb"
	"sync"
	"time"
	"google.golang.org/grpc"
	"context"
)

var (
	aiuiOnce      sync.Once
	aiuiClient 	  pb.AiuiServerClient
	aiuiConn      *grpc.ClientConn
)

func AiuiRpcInit(address string) {
	aiuiOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		aiuiConn, err = grpc.DialContext(ctx, address, grpc.WithInsecure())	  //, grpc.WithBlock()
		if err != nil {
			log.Error("AiuiRpcInit-err:", err)
			panic(err)
		}
		aiuiClient = pb.NewAiuiServerClient(aiuiConn)
	})
}
func AiuiRpcClose() {
	if aiuiConn != nil {
		log.Info("AiuiRpcClose:",aiuiConn.Close())
	}
}
func AiuiClient() pb.AiuiServerClient {
	log.Debugf("AiuiClient(%#v)", aiuiClient)
	return aiuiClient
}

//GetTextSemantics
func GetTextSemantics(in *pb.GetTextSemanticsRequest) (*pb.GetTextSemanticsResponse, error) {
	log.Info("GetTextSemantics...", )
	reply,err := AiuiClient().GetTextSemantics(context.Background(),in)
	if err != nil {
		log.Errorf("讯飞aiui 服务报错！err(%#v）",err)
	}
	return reply,err
}