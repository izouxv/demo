package rpcClient

import (
	"cotx-http/pb"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	log "github.com/cihub/seelog"
)

var rpcClientDevice pb.DeviceClient

func NewRpcClinetDevice()  {
	ctx,_ :=context.WithTimeout(context.Background(),10*time.Second)
	conn,err:= grpc.DialContext(ctx,address,grpc.WithInsecure())
	if err != nil {
		log.Error("gwuserRpcInit Error:", err)
		return
	}
	rpcClientDevice = pb.NewDeviceClient(conn)
}
func GetDeviceRpcClient()pb.DeviceClient  {
	return rpcClientDevice
}
func GetNodeDateUpByNid(req *pb.ReqNodeMessageByNid)(*pb.ResNodeDateUp,error)  {
	log.Info("rpc/start get node dateup by nodeid")
	res,err:= GetDeviceRpcClient().GetNodeDateUpByNid(context.Background(),req)
	return res,err
}

func GetDeviceEuis(req *pb.ReqNodeMessageByNid)(*pb.ResGetDeviceScan ,error) {
	log.Info("rpc/start get nodes' deveuis ")
	res,err := GetDeviceRpcClient().GetDeviceEuis(context.Background(),req)
	return res,err
}
//
func RegistDevices(req *pb.ReqregistDevice)(*pb.ResRegistDevice ,error) {
	log.Info("rpc/start regist devices")
	res,err:= GetDeviceRpcClient().RegistDevices(context.Background(),req)
	return res,err
}
//
func  RegistNode(req *pb.ReqRegistNode)(*pb.ResRegistNode ,error) {
	log.Info("rpc/start regist node")
	res,err:= GetDeviceRpcClient().RegistNode(context.Background(),req)
	return res,err
}
func AddDevice(req *pb.PutDeviceRequest)(*pb.PutDeviceResponse)  {
	res,err := GetDeviceRpcClient().AddDevice(context.Background(),req)
	if err != nil {
		log.Errorf("添加终端失败 error %s",err)
		return nil
	}
	return res
}

func GetDevice(req *pb.GetDeviceRequest)(*pb.GetDeviceResponse)  {
	res,err := GetDeviceRpcClient().GetDevice(context.Background(),req)
	if err != nil {
		log.Errorf("获取终端信息失败 error %s",err)
		return nil
	}
	return res
}

func UpdateDevice(req *pb.PostDeviceRequest)(*pb.PostDeviceResponse)  {
	res,err := GetDeviceRpcClient().UpdateDevice(context.Background(),req)
	if err != nil {
		log.Errorf("更新终端的参数失败 error %s",err)
		return nil
	}
	return res
}

func GetDeviceListByDeveuis(req *pb.GetDeviceListRequest)(*pb.GetDeviceListResponse)  {
	res,err := GetDeviceRpcClient().GetDeviceListByDeveuis(context.Background(),req)
	if err != nil {
		log.Errorf("获取终端列表失败 error %s",err)
	}
	return res
}

func DeleteDeviceByDeveui(req *pb.DeleteDeviceRequest) (*pb.DeleteDeviceResponse)  {
	res, err := GetDeviceRpcClient().DeleteDeviceByDeveui(context.Background(),req)
	if err != nil {
		log.Error(err)
	}
	return res
}

func AddDevices(req *pb.AddDevicesRequest) (*pb.AddDevicesResponse)  {
	res,err := GetDeviceRpcClient().AddDevices(context.Background(),req)
	if err != nil {
		log.Error(err)
	}
	return res
}

