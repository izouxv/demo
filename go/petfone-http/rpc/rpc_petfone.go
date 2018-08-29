package rpc

import (
	"context"
	log "github.com/cihub/seelog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"petfone-http/pb"
	"sync"
	"time"
	"io"
)

var (
	petfoneOnce      sync.Once
	petfoneRpcClient pb.PetfoneClient
	petfoneConn *grpc.ClientConn
)

//初始化PetfoneRpc
func PetfoneRpcInit(address string) {
	petfoneOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		petfoneConn, err = grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Error("PetfoneRpcInit:", err)
			panic(err)
		}
		petfoneRpcClient = pb.NewPetfoneClient(petfoneConn)
	})
}

//结束Rpc
func PetfoneRpcClose() {
	if petfoneConn != nil {
		petfoneConn.Close()
	}
}

//调用PetfoneRpc
func PetfoneRpc(petfone *pb.PetfoneRequest, method string) *pb.PetfoneReply {
	log.Info("PetfoneRpc-petfone:", petfone)
	var err error
	var reply *pb.PetfoneReply
	switch method {
	//获取围栏信息
	case "GetPetfoneByUid":
		reply, err = petfoneRpcClient.GetPetfoneByUid(context.Background(), petfone)
		break
		//修改围栏信息
	case "UpdatePetfoneByUid":
		reply, err = petfoneRpcClient.UpdatePetfoneByUid(context.Background(), petfone)
		break
		//设置围栏信息
	case "SetPetfoneByUid":
		reply, err = petfoneRpcClient.SetPetfoneByUid(context.Background(), petfone)
		break
	default:
		err = errors.New("没有该RPC")
	}
	if err != nil {
		log.Error("PetfoneRpc-Error", err)
		reply.Code = 10001
	}
	log.Info("PetfoneRpc-reply:", reply)
	return reply
}

//PetChatRpc
func PetChatRpc(petChat *pb.PetChatRequest, method string) *pb.PetChatReply {
	log.Info("PetfoneRpc-petChat:", petChat)
	var err error
	var reply *pb.PetChatReply
	switch method {
	//获取信息
	case "GetPetChatByPid":
		reply, err = petfoneRpcClient.GetPetChatByPid(context.Background(), petChat)
	default:
		err = errors.New("没有该RPC")
	}
	if err != nil {
		log.Error("PetfoneRpc-err", err)
		reply.Code = 10001
	}
	return reply
}

//PetChatFaqRpc
func PetChatFaqRpc(petChat *pb.PetChatRequest, method string) *pb.PetChatKeysReply {
	log.Info("PetChatFaqRpc-petChat:", petChat)
	var err error
	var reply *pb.PetChatKeysReply
	switch method {
	//获取宠聊常见问题信息
	case "GetPetChatKey":
		reply, err = petfoneRpcClient.GetPetChatKey(context.Background(), petChat)
	default:
		err = errors.New("没有该RPC")
	}
	if err != nil {
		log.Error("PetChatFaqRpc-err", err)
		reply.Code = 10001
	}
	return reply
}

//PetfoneRpcActionLog
func PetfoneRpcActionLog(agent *pb.AgentInfo) {
	log.Info("ActionLog-agents:", agent.Ip)
	stream, err := petfoneRpcClient.SetActionLog(context.Background())
	if err != nil {
		log.Error("ActionLog SetActionLog err:", err)
		return
	}
	err = stream.Send(agent)
	if err != nil {
		log.Error("ActionLog Send err:", err)
		return
	}
	stream.CloseSend()
	res, err := stream.Recv()
	if err != nil {
		if err == io.EOF {
			log.Info("ActionLog Recv:",err)
			return
		}
		log.Error("ActionLog Recv:", err)
		return
	}
	log.Info("ActionLog Recv:",res.String())
}

//限制ip访问次数
func CheckPetfoneIp(res *pb.CheckPetfoneIpRequest) (reply *pb.CheckPetfoneIpResponse,err error ){
	log.Infof("CheckPetfoneIp...res(%#v)", res)
	reply, err = petfoneRpcClient.CheckPetfoneIp(context.Background(), res)
	if reply.Code != 10000 {
		log.Error("CheckPetfoneIp rpc err")
	}
	return
}

