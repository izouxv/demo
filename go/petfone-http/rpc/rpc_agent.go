package rpc

import (
	"context"
	log "github.com/cihub/seelog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"petfone-http/pb/api"
	"sync"
)

var (
	agentOnce   sync.Once
	agentClient api.TwinsAgentServerClient
	agentConn *grpc.ClientConn
)

//AgentRpcInit
func AgentRpcInit(address string) {
	agentOnce.Do(func() {
		agentConn, err = grpc.DialContext(otherCtx, address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Error("AgentRpcInit error:", err)
			return
		}
		agentClient = api.NewTwinsAgentServerClient(agentConn)
	})
}

//结束Rpc
func AgentRpcClose() {
	if agentConn != nil {
		agentConn.Close()
	}
}

//AdverRpc
func TwinsRpc(twins *api.AddTwinsAgentRequest, method string) {
	log.Info("TwinsRpc-twins:", twins)
	if agentClient == nil {
		log.Error("TwinsRpc twinsClient nil")
		return
	}
	var err error
	var reply *api.AddTwinsAgentResponse
	switch method {
	//上报
	case "AddTwinsAgent":
		reply, err = agentClient.AddTwinsAgent(context.Background(), twins)
	default:
		err = errors.New("没有该RPC")
	}
	if err != nil {
		log.Error("TwinsRpc-Error", err)
	}
	log.Info("TwinsRpc-reply:", reply)
}
