package rpc

import (
	"context"
	log "github.com/cihub/seelog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"petfone-http/pb/adv"
	"sync"
)

var (
	advOnce   sync.Once
	advertisementClient adv.AdvertisementClient
	advConn *grpc.ClientConn
)

//AdverRpcInit
func AdverRpcInit(address string) {
	advOnce.Do(func() {
		advConn, err = grpc.DialContext(otherCtx, address, grpc.WithInsecure())
		if err != nil {
			log.Error("AdverRpcInit err:", err)
			return
		}
		advertisementClient = adv.NewAdvertisementClient(advConn)
	})
}

//结束Rpc
func AdverRpcClose() {
	if advConn != nil {
		advConn.Close()
	}
}

//AdverRpc
func AdvertisementRpc(advertisementRpc *adv.AdvertisementRequest, method string) *adv.AdvertisementReply {
	log.Info("AdvertisementRpc-advertisementRpc:", advertisementRpc)
	var err error
	var reply *adv.AdvertisementReply
	if agentClient == nil {
		log.Error("AdvertisementRpc advertisementClient nil")
		reply.ErrorCode = 10001
		return reply
	}
	switch method {
	//获取
	case "GetAdvertisement":
		reply, err = advertisementClient.GetAdvertisement(context.Background(), advertisementRpc)
	default:
		err = errors.New("没有该RPC")
	}
	if err != nil {
		log.Error("AdvertisementRpc-Error", err)
		reply.ErrorCode = 10001
	}
	log.Info("AdvertisementRpc-reply:", reply)
	return reply
}
