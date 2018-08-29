package rpcClient

import (
	"time"
	"cotx-http/pb"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	log "github.com/cihub/seelog"
)

var rpcClientSoftware pb.SoftwareServiceClient

func NewRpcClinetSoftware()  {
	ctx,_ :=context.WithTimeout(context.Background(),10*time.Second)
	conn,err:= grpc.DialContext(ctx,address,grpc.WithInsecure())
	if err != nil {
		log.Error("gwuserRpcInit Error:", err)
		return
	}
	rpcClientSoftware = pb.NewSoftwareServiceClient(conn)
}
func GetSoftwareRpcClient() pb.SoftwareServiceClient  {
	return rpcClientSoftware
}