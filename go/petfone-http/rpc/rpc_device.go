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
	devicesOnce      sync.Once
	devicesRpcClient pb.DevicesClient
	deviceConn *grpc.ClientConn
)

//初始化DevicesRpcInit
func DevicesRpcInit(address string) {
	devicesOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		deviceConn, err = grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Error("DevicesRpcInit-err:", err)
			panic(err)
		}
		devicesRpcClient = pb.NewDevicesClient(deviceConn)
	})
}

//结束Rpc
func DeviceRpcClose() {
	if deviceConn != nil {
		deviceConn.Close()
	}
}

//调用devicesRpc
func DevicesRpc(device *pb.DeviceRequest, method string) *pb.DeviceReply {
	log.Info("DeviceRpc-device:", device)
	var err error
	var reply *pb.DeviceReply
	switch method {
	//校验设备
	case "VerificationDeviceBySn":
		reply, err = devicesRpcClient.VerificationDeviceBySn(context.Background(), device)
		//添加设备
	case "SetDeviceBySn":
		reply, err = devicesRpcClient.SetDeviceBySn(context.Background(), device)
		//解绑设备
	case "DeleteDeviceByDid":
		reply, err = devicesRpcClient.DeleteDeviceByDid(context.Background(), device)
		//修改设备指令
	case "UpdateDeviceByDid":
		reply, err = devicesRpcClient.UpdateDeviceByDid(context.Background(), device)
		//获取设备sn
	case "GetDeviceSn":
		reply, err = devicesRpcClient.GetDeviceSn(context.Background(), device)
	case "GetDevicesByDid":
		reply, err = devicesRpcClient.GetDevicesByDid(context.Background(), device)
	default:
		err = errors.New("DeviceRpc-没有该RPC:" + method)
	}
	if err != nil {
		log.Error("DeviceRpc-Error", err)
		reply.Code = 10001
	}
	log.Info("DeviceRpc-reply:", reply)
	return reply
}

//调用deviceRpc
func DeviceMapRpc(device *pb.DeviceRequest, method string) *pb.BatchDeviceRe {
	log.Info("DevicesRpc-device:", device)
	var deviceErr error
	var deviceReply *pb.BatchDeviceRe
	switch method {
	//获取用户设备
	case "GetDevicesByUid":
		deviceReply, deviceErr = devicesRpcClient.GetDevicesByUid(context.Background(), device)
	default:
		deviceErr = errors.New("DevicesRpc-没有该RPC:" + method)
	}
	if deviceErr != nil {
		log.Error("DevicesRpc-Error", deviceErr)
		deviceReply.Code = 10001
	}
	return deviceReply
}
