package rpcClient

import (
	"cotx-http/pb"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	log "github.com/cihub/seelog"
)

var  rpcClientGatewayUsb  pb.GatewayUsbClient
func NewGatewayUsbRpcClient() {
	ctx,_ :=context.WithTimeout(context.Background(),10*time.Second)
	conn,err:= grpc.DialContext(ctx,address,grpc.WithInsecure())
	if err != nil {
		log.Error("gateway set RpcInit Error:", err)
		return
	}
	rpcClientGatewayUsb = pb.NewGatewayUsbClient(conn)
}
func GetGatewayUsbRpcClient() pb.GatewayUsbClient {
	return rpcClientGatewayUsb
}
func GetUsbWifiWlan(req *pb.ReqGatewayUsb)(*pb.ResUsbWifiWlan ,error) {
	log.Info("rpc/start get gateway usb wifi wlan")
	res,err:= GetGatewayUsbRpcClient().GetUsbWifiWlan(context.Background(),req)
	return res,err
}
func GetUsbHotsptUser(req *pb.ReqGatewayUsb)(*pb.ResUsbWifiHotSpotUser,error) {
	log.Info("rpc/start get gateway usb ")
	res,err:= GetGatewayUsbRpcClient().GetUsbHotSpotUser(context.Background(),req)
	return res,err
}

func GetWifiHotSpotUser(req *pb.ReqGatewayUsb) (*pb.ResWifiHotSpotUser,error) {
	log.Info("rpc/start get gateway hotspot user")
	res,err := GetGatewayUsbRpcClient().GetWifiHotSpotUser(context.Background(),req)
	return res,err
}
func GetUsbNumWifiScan(req *pb.ReqGatewayUsb)(*pb.ResUsbWifiScan ,error){
	log.Info("rpc/start get gateway usb wifi scan")
	res,err := GetGatewayUsbRpcClient().GetUsbWifiScan(context.Background(),req)
	return res,err
}
