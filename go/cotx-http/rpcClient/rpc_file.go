package rpcClient

import (
	"cotx-http/pb"
	"context"
	"time"
	"google.golang.org/grpc"
	log "github.com/cihub/seelog"
)

var fileClient pb.FileServiceClient

func NewFileClient()  {
	ctx,_ := context.WithTimeout(context.Background(),10*time.Second)
	conn,err := grpc.DialContext(ctx,address,grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return
	}
	fileClient = pb.NewFileServiceClient(conn)
}

func GetFileClient() pb.FileServiceClient  {
	return fileClient
}