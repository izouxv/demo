package rpcClient

import (
	"cotx-http/pb"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	log "github.com/cihub/seelog"
)

var (
	GatewaySetRpcClient pb.GatewatSetClient
)
func NewGatewaySetRpcClient() {
	ctx,_ :=context.WithTimeout(context.Background(),10*time.Second)
	conn,err:= grpc.DialContext(ctx,address,grpc.WithInsecure())
	if err != nil {
		log.Error("gateway set RpcInit Error:", err)
		return
	}
	GatewaySetRpcClient = pb.NewGatewatSetClient(conn)
}
func GetGatewaySetRpcClient() pb.GatewatSetClient {
	return GatewaySetRpcClient
}
func SendInstruction (req *pb.ReqInstruction)(*pb.ResCode,error){
	 log.Info("rpc/start send instruction")
	res,err:= GetGatewaySetRpcClient().SendInstruction(context.Background(),req)
	return res,err
}
/*设置网关开关的状态*/
func SendSwitch(req *pb.ReqSwitch)(*pb.ResCode,error){
	log.Info("rpc/start send switch")
	res,err:= GetGatewaySetRpcClient().SendSwitch(context.Background(),req)
	return res,err
}
/*设置网关的节电模式*/
func SendSetPower(req *pb.ReqGwPower)(*pb.ResCode,error)  {
	log.Info("rpc/start send power model")
	res,err := GetGatewaySetRpcClient().SendSetPower(context.Background(),req)
	return res,err
}

/*设置网关lora频段参数*/
func SendLora(req *pb.ReqLora)(*pb.ResCode,error)  {
	log.Info("rpc/start send set lora")
	res,err:=GetGatewaySetRpcClient().SendLora(context.Background(),req)
	return res,err
}

/*设置摄像*/
func SendVideo(req *pb.ReqVideo)(*pb.ResCode,error) {
	log.Info("rpc/start send set Video")
	res,err:=GetGatewaySetRpcClient().SendVideo(context.Background(),req)
	return res,err
}

/*设置拍照智能实现*/
func SendPhoto(req *pb.ReqPhoto)(*pb.ResCode,error) {
	log.Info("rpc/start send set photo")
	res,err:=GetGatewaySetRpcClient().SendPhoto(context.Background(),req)
	return res,err
}
func SendMusic(req *pb.ReqMusic)(*pb.ResCode,error) {
	log.Info("rpc/start send set music")
	res,err:=GetGatewaySetRpcClient().SendMusic(context.Background(),req)
	return res,err
}
func SendDeletFile(req *pb.ReqDeletFile)(*pb.ResCode,error)  {
	log.Info("rpc/start send set Delet file")
	res,err:=GetGatewaySetRpcClient().SendDeletFile(context.Background(),req)
	return res,err
}
func SendUplog(req *pb.ReqUpLog)(*pb.ResCode,error)  {
	log.Info("rpc/start send set uplog")
	res,err:=GetGatewaySetRpcClient().SendUpLog(context.Background(),req)
	return res,err
}
/**/
func SendSetIP(req *pb.ReqSetIP)(*pb.ResCode,error)  {
	log.Info("rpc/start send set ip")
	res,err:=GetGatewaySetRpcClient().SendSetIp(context.Background(),req)
	return res,err
}
func SendSetDNS(req *pb.ReqSetDNS)(*pb.ResCode,error)  {
	log.Info("rpc/start send set dns")
	res,err:=GetGatewaySetRpcClient().SendSetDNS(context.Background(),req)
	return res,err
}
func SendSetHotSpot(req *pb.ReqSetHotSpot)(*pb.ResCode,error)  {
	log.Info("rpc/start send set hotspot")
	res,err:=GetGatewaySetRpcClient().SendSetHotSpot(context.Background(),req)
	return res,err
}
func SendName(req *pb.ReqSetName)(*pb.ResCode ,error) {
	log.Info("rpc/start send set Name")
	res,err:=GetGatewaySetRpcClient().SendSetName(context.Background(),req)
	return res,err
}
func SendConnectWifi (req *pb.ReqConnectWifi) (*pb.ResCode ,error){
	log.Info("rpc/start send connect wifi")
	res,err:=GetGatewaySetRpcClient().SendConnectWifi(context.Background(),req)
	return res,err
}
func SendUsbConnectionWifi(req *pb.ReqConnectUsbWifi)(*pb.ResCode ,error){
	log.Info("rpc/start send usb connection wifi")
	res,err:= GetGatewaySetRpcClient().SendUsbConnectWifi(context.Background(),req)
	return res,err
}
func SendSetUsbIp(req *pb.ReqSetUsbIP) (*pb.ResCode ,error){
	log.Info("rpc/set start send set usb ip")
	res,err:= GetGatewaySetRpcClient().SendSetUsbIp(context.Background(),req)
	return res,err
}
func SendSetUsbDns(req *pb.ReqSetUsbDNS)(*pb.ResCode ,error)  {
	log.Info("rpc/set start send set sub dns")
	res,err:= GetGatewaySetRpcClient().SendSetUsbDNS(context.Background(),req)
	return res,err
}
func SendUsbSwitch(req *pb.ReqUsbSwitch)(*pb.ResCode ,error) {
	log.Info("rpc/ set start send set usb switch")
	res,err := GetGatewaySetRpcClient().SendUsbSwitch(context.Background(),req)
	return res,err
}
func SendUsbSetHotSpot(req *pb.ReqSetUsbHotSpot)(*pb.ResCode ,error)  {
	log.Info("rpc/set start send set usb hotspot")
	res,err := GetGatewaySetRpcClient().SendSetUsbHotSpot(context.Background(),req)
	return res,err
}

/*设置第三方云平台信息*/
func SetIOTServer(req *pb.ReqIOTServer) (*pb.ResCode ,error) {
	log.Info("rpc_gatewayset/ start set IOT Server ...")
	res,err := GetGatewaySetRpcClient().SetIOTServer(context.Background(),req)
	return res,err
}

/*设置手动拍照*/
func SetGatewayPhoto(req *pb.ReqInstruction) (*pb.ResPhotoCode ,error) {
	log.Info("rpc_gatewayset /start set take photo")
	res,err := GetGatewaySetRpcClient().SetGatewayPhoto(context.Background(),req)
	return res,err

}

/*设置手动摄像*/
func SetGatewayVideo(req *pb.ReqSwitch)(*pb.ResVideoCode ,error) {
	log.Info("rpc_gatewayset /start set take video")
	res,err:= GetGatewaySetRpcClient().SetGatewayVideo(context.Background(),req)
	return res,err
}