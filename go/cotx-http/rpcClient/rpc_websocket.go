package rpcClient

import (
	"time"
	"cotx-http/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	log "github.com/cihub/seelog"
)

var (
	WebsocketRpcClient pb.WebsocketClient
)

func NewWebscoketRpcClient() {
	ctx,_ :=context.WithTimeout(context.Background(),10*time.Second)
	conn,err:= grpc.DialContext(ctx,address,grpc.WithInsecure())
	if err != nil {
		log.Error("gwuserRpcInit Error:", err)
		return
	}
	WebsocketRpcClient = pb.NewWebsocketClient(conn)
}

func GetWebsocketRpcClient()pb.WebsocketClient {
	return WebsocketRpcClient
}

func GetPushMessage(req *pb.ReqShadow) (*pb.ResShadow,error)  {
	log.Info("Websocket:获取告警信息")
	res,err := GetWebsocketRpcClient().GetPushMessage(context.Background(),req)
	return res,err
}