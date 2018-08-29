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
	petInfoOnce      sync.Once
	petInfoRpcClient pb.PetInfoClient
	petInfoConn *grpc.ClientConn
)

//初始化PetfoneRpc
func PetInfoRpcInit(address string) {
	petInfoOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		petInfoConn, err = grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Error("PetInfoRpcInit:", err)
			panic(err)
		}
		petInfoRpcClient = pb.NewPetInfoClient(petInfoConn)
	})
}

//结束Rpc
func PetInfoRpcClose() {
	if petInfoConn != nil {
		petInfoConn.Close()
	}
}

//调用PetinfoRpc
func PetinfoRpc(petinfo *pb.PetInfoRequest, method string) *pb.PetInfoReply {
	log.Info("PetinfoRpc-petinfo:", petinfo)
	var err error
	var reply *pb.PetInfoReply
	switch method {
	//添加宠物信息
	case "SetPetInfo":
		reply, err = petInfoRpcClient.SetPetInfo(context.Background(), petinfo)
		//删除宠物信息
	case "DeletePetInfoByPid":
		reply, err = petInfoRpcClient.DeletePetInfoByPid(context.Background(), petinfo)
		//修改宠物信息
	case "UpdatePetInfoByPid":
		reply, err = petInfoRpcClient.UpdatePetInfoByPid(context.Background(), petinfo)
		//查询宠物信息
	case "GetPetInfoByPid":
		reply, err = petInfoRpcClient.GetPetInfoByPid(context.Background(), petinfo)
		//关联设备
	case "SetDevicePet":
		reply, err = petInfoRpcClient.SetDevicePet(context.Background(), petinfo)
		//取消关联
	case "DeleteDevicePet":
		reply, err = petInfoRpcClient.DeleteDevicePet(context.Background(), petinfo)
	default:
		err = errors.New("没有该RPC")
	}
	if err != nil {
		log.Error("PetinfoRpc-err", err)
		reply.Code = 10001
	}
	log.Info("PetinfoRpc-reply:", reply)
	return reply
}

//调用PetinfosRpc
func PetinfosRpc(petinfo *pb.PetInfoRequest, method string) *pb.PetInfoMapReply {
	log.Info("PetinfoRpc-petinfo:", petinfo)
	var err error
	var reply *pb.PetInfoMapReply
	switch method {
	//查询用户的宠物信息
	case "GetPetInfoByUid":
		reply, err = petInfoRpcClient.GetPetInfoByUid(context.Background(), petinfo)
		break
	default:
		err = errors.New("没有该RPC")
	}
	if err != nil {
		log.Error("PetinfoRpc-Error", err)
		reply.Code = 10001
	}
	log.Info("PetinfoRpc-reply-len:", len(reply.Petinfos))
	return reply
}

//调用trainRpc
func TrainRpc(train *pb.PetTrainRequest, method string) *pb.PetTrainReply {
	log.Info("TrainRpc-train:", train)
	var err error
	var reply *pb.PetTrainReply
	switch method {
		//修改训练信息
	case "UpdatePetTrainByPid":
		reply, err = petInfoRpcClient.UpdatePetTrainByPid(context.Background(), train)
		break
		//次数计数
	case "CounterPetTrainByPid":
		reply, err = petInfoRpcClient.CounterPetTrainByPid(context.Background(), train)
		break
	default:
		err = errors.New("没有该RPC")
	}
	if err != nil {
		log.Error("TrainRpc-Error", err)
		reply.Code = 10001
	}
	log.Info("TrainRpc-reply:", reply)
	return reply
}

//调用trainsRpc
func TrainsRpc(train *pb.PetTrainRequest, method string) *pb.PetSliceTrainsReply {
	log.Info("TrainsRpc-train:", train)
	var err error
	var reply *pb.PetSliceTrainsReply
	switch method {
	//获取训练信息
	case "GetPetTrainByPid":
		reply, err = petInfoRpcClient.GetPetTrainByPid(context.Background(), train)
		break
	default:
		err = errors.New("没有该RPC")
	}
	if err != nil {
		log.Error("TrainsRpc-Error", err)
		reply.Code = 10001
	}
	log.Info("TrainsRpc-reply:", len(reply.SliceTrains))
	return reply
}

//更新宠端设备训练录音
func UpdateDeviceTrainByDid( in *pb.DeviceTrainRequest) *pb.DeviceTrainReply {
	log.Info("UpdateDeviceTrainByDid...")
	reply, err := petInfoRpcClient.UpdateDeviceTrainByDid(context.Background(),in)
	if err != nil {
		log.Error("UpdateDeviceTrainByDid-Error", err)
		reply.Code = 10001
	}
	log.Info("UpdateDeviceTrainByDid-reply:", reply)
	return reply
}

//获取宠物信息和宠端设备录音
func GetPetInfoBydid( in *pb.PetInfoRequest) *pb.PetInfoReply {
	log.Info("GetPetInfoBydid...")
	reply, err := petInfoRpcClient.GetPetInfoBydid(context.Background(),in)
	if err != nil {
		log.Error("UpdateDeviceTrainByDid-Error", err)
		reply.Code = 10001
	}
	log.Info("UpdateDeviceTrainByDid-reply:", reply)
	return reply
}

//获取批量宠物信息和设备录音
func GetPetInfosBydid( in *pb.PetInfoRequest) *pb.PetInfoMapReply {
	log.Info("GetPetInfosBydid...")
	reply, err := petInfoRpcClient.GetPetInfosBydid(context.Background(),in)
	if err != nil {
		log.Error("GetPetInfosBydid-Error", err)
		reply.Code = 10001
	}
	log.Info("GetPetInfosBydid-reply:", reply)
	return reply
}




