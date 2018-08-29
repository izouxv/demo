package rpc

import (
	"context"
	log "github.com/cihub/seelog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"petfone-http/pb"
	"sync"
	"time"
)

var (
	filesOnce      sync.Once
	filesRpcClient pb.FilesClient
	filesConn *grpc.ClientConn
)

//初始化filesRpc
func FilesRpcInit(address string) {
	filesOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		filesConn, err = grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Error("PetInfoRpcInit:", err)
			panic(err)
		}
		filesRpcClient = pb.NewFilesClient(filesConn)
	})
}

//结束Rpc
func FilesRpcClose() {
	if filesConn != nil {
		filesConn.Close()
	}
}

//调用FileRpc
func FileRpc(file *pb.FilesRequest, method string) *pb.FilesReply {
	log.Info("FilesRpc-file:", file)
	var filesErr error
	var filesReply *pb.FilesReply
	switch method {
	//添加信息
	case "SetFile":
		filesReply, filesErr = filesRpcClient.SetFile(context.Background(), file)
		//修改信息
	case "GetFile":
		filesReply, filesErr = filesRpcClient.GetFile(context.Background(), file)
	default:
		filesErr = errors.New("没有该RPC")
	}
	if filesErr != nil {
		log.Error("FilesRpc-Error", filesErr)
		filesReply.Code = 10001
	}
	log.Info("FilesRpc-filesReply:", filesReply)
	return filesReply
}

//调用FilesRpc
func FilesRpc(file *pb.FilesRequest, method string) *pb.FilesMapReply {
	log.Info("FilesRpc-file:", file)
	var filesErr error
	var filesReply *pb.FilesMapReply
	switch method {
	//获取信息
	case "GetBreeds":
		filesReply, filesErr = filesRpcClient.GetBreeds(context.Background(), file)
	default:
		filesErr = errors.New("没有该RPC")
	}
	if filesErr != nil {
		log.Error("FilesRpc-Error", filesErr)
		filesReply.Code = 10001
	}
	return filesReply
}
