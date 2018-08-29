package rpcClient

import (
	"cotx-http/pb"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	log "github.com/cihub/seelog"
)

const (
	Address = "127.0.0.1:7007"
)
var (
	GatewayRpcClient pb.GatewayClient
)

func NewGatewayRpcClient() {
	ctx,_ :=context.WithTimeout(context.Background(),10*time.Second)
	conn,err:= grpc.DialContext(ctx,address,grpc.WithInsecure())
	if err != nil {
		log.Error("gwuserRpcInit Error:", err)
		return
	}
	GatewayRpcClient = pb.NewGatewayClient(conn)
}
func GetGatewayRpcClient() pb.GatewayClient {
	return GatewayRpcClient
}
//app获取单个网关的网络状态信息
func GetGatewayNetState(req *pb.ReqGateway)(*pb.ResGwNetState ,error) {
	log.Info("rpc/start get gateway net state")
	resgwnetstate,err := GetGatewayRpcClient().GetGatewayNetState(context.Background(),req)
	return resgwnetstate,err
}
//app获取网关的状态信息
func GetGatewayState(req * pb.ReqGateway)(*pb.ResGwState,error) {
	log.Info("rpc/start get gateway state")
	resgwstate ,err:= GetGatewayRpcClient().GetGatewayState(context.Background(),req)
	return resgwstate,err
}
//获取网关的文件统计信息
func GetGatewayFileStat(req *pb.ReqGateway)(*pb.ResGwFile,error) {
	log.Info("rpc/start get gateway file stat ")
	resgwfile ,err:= GetGatewayRpcClient().GetGatewayFileStat(context.Background(),req)
	return resgwfile,err
}
//获取网关的视频信息
func GetGatewayVideos(req *pb.ReqGateway)(*pb.ResGwVideos,error){
	log.Info("rpc/start get gateway videos")
	resgwvideos,err := GetGatewayRpcClient().GetGatewayVideos(context.Background(),req)
	return resgwvideos,err
}
//获取网关的图片信息
func GetGatewayPhotos(req *pb.ReqGateway)(*pb.ResGwPhotos,error){
	log.Info("rpc/start get gateway photos")
	resgwphotos,err := GetGatewayRpcClient().GetGatewayPhotos(context.Background(),req)
	return resgwphotos,err
}
//获取网关的usb统计信息
func GetGatewayUSBStat(req *pb.ReqGateway)(*pb.ResGwUSBStat,error){
	log.Info("rpc/start get gateway usb stat")
	resgwusb,err:= GetGatewayRpcClient().GetGatewayUSBStat(context.Background(),req)
	return resgwusb,err
}
//获取wifi的扫描信息
func GetGatewayWifiScans(req *pb.ReqGateway)(*pb.ResGwWifiScans ,error) {
	log.Info("rpc/start get gateway wifi scnas")
	resgwwifiscans ,err:= GetGatewayRpcClient().GetGatewayWifiScans(context.Background(),req)
	return resgwwifiscans,err
}
//获取wifi的ipv4信息
func GetGatewayWifiAddress(req *pb.ReqGateway)(*pb.ResGwWifiAddress,error){
	log.Info("rpc/start get gateway wifi address")
	resgwwifiaddress,err:=GetGatewayRpcClient().GetGatewayWifiAddress(context.Background(),req)
	return resgwwifiaddress,err
}
//获取wifi的dns
func GetGatewayWifiDNS(req *pb.ReqGateway)(*pb.ResGwWifiDNS,error){
	log.Info("rpc/start get gateway wifi dns")
	resgwwifidns,err:= GetGatewayRpcClient().GetGatewayWifiDNS(context.Background(),req)
	return resgwwifidns,err
}
//获取有线额ipv4地址
func GetGatewayCableAddress(req *pb.ReqGateway)(*pb.ResGwCableAddress,error){
	log.Info("rpc/start get gateway cable address")
	resgwcableaddress,err:=GetGatewayRpcClient().GetGatewayCableAddress(context.Background(),req)
	return resgwcableaddress,err
}
//获取有线的dns
func GetGatewayCAbleDNS(req *pb.ReqGateway)(*pb.ResGwCableDNS,error){
	log.Info("rpc/start get gateway Cable DNS")
	resgwcabledns,err:=GetGatewayRpcClient().GetGatewayCableDNS(context.Background(),req)
	return resgwcabledns,err
}
//获取网关的属性信息
func GetGatewayMessage(req *pb.ReqGateway)(*pb.ResGwMessage ,error) {
	log.Info("rpc/start get gateway message ")
	resgwmessage,err := GetGatewayRpcClient().GetGatewayMessage(context.Background(),req)
	return resgwmessage,err
}
//获取网关的lora信息
func GetGatewayLora(req *pb.ReqGateway)(*pb.ResGwLora,error){
	log.Info("rpc/Start get gateway lora message")
	resgwlora ,err := GetGatewayRpcClient().GetGatewayLora(context.Background(),req)
	return resgwlora,err
}
//
func GetGatewayUsbWifiAddress(req *pb.ReqGateway)(*pb.ResGwUsbWifiAddress,error)  {
	log.Info("rpc/start get gateway usb  wifiaddress")
	res,err:= GetGatewayRpcClient().GetGatewayUsbWifiAddress(context.Background(),req)
	return res,err
}
//
func GetGatewayUsbWifiDNS(req *pb.ReqGateway)( *pb.ResGwUsbWifiDNS ,error){
	log.Info("rpc/start get gateway usb wifidns")
	res,err:= GetGatewayRpcClient().GetGatewayUsbWifiDNS(context.Background(),req)
	return res,err
}
// 云平台的信息
func GetGatewayServerIots(req  *pb.ReqGateway)(*pb.ResServerIots ,error) {
	log.Info("rpc/start get gateway serveriots ")
	res ,err  := GetGatewayRpcClient().GetGatewayServerIOts(context.Background(),req)
	return res,err
}
// 网关的开关状态信息
func GetGatewaySwitch(req *pb.ReqGateway)(*pb.ResGatewaySwitch,error)  {
	log.Info("rpc/start get gateway switch ")
	res,err:= GetGatewayRpcClient().GetGatewaySwitch(context.Background(),req)
	return res,err
}
//网关自带wifi的热点信息
func GetHotSpot(req *pb.ReqGateway)(*pb.ResHotSpot,error)  {
	log.Info("rpc/ start get gateway hotspot ")
	res,err := GetGatewayRpcClient().GetHotSpot(context.Background(),req)
	return res,err
}
//网关的外接usb wifi 热点记录信息
func GetUsbHotSpot(req *pb.ReqGateway)(*pb.ResUsbHotSpot,error)  {
	log.Info("rpc/ start get gateway usb hotspot")
	res ,err:= GetGatewayRpcClient().GetUsbHotSpot(context.Background(),req)
	return res,err
}
//
func GetUsbWifistat(req *pb.ReqGateway)(*pb.ResUsbWifiStat,error)  {
	log.Info("rpc/start get gateway usb wifi stat")
	res,err:= GetGatewayRpcClient().GetUsbWifiStat(context.Background(),req)
	return res,err
}
//
func GetUsbGCardStat(req *pb.ReqGateway)(*pb.ResUsbGCardStat,error) {
	log.Info("rpc/start get gateway usb gcard stat")
	res ,err:= GetGatewayRpcClient().GetUsbGCardStat(context.Background(),req)
	return res,err
}
//
func GetWifi (req *pb.ReqGateway)(*pb.ResWifiConnected,error){
	log.Info("rpc/start get gateway wifi connecting")
	res ,err := GetGatewayRpcClient().GetUsbWifiConnecting(context.Background(),req)
	return res,err
}
//usb接口的扫描数据
func GetUsbWifiScan(req *pb.ReqGateway)(*pb.ResGwUsbWifiScans,error)  {
	log.Info("rpc/start get gateway usb wifiscan")
	res,err := GetGatewayRpcClient().GetUsbWifiScans(context.Background(),req)
	return res,err
}

/*获取网关媒体文件下载地址*/
func GetGatewayFile(req *pb.ReqGetGatewayFile)(*pb.ResGetGatewayFile,error)  {
	log.Info("rpc/start get gateway file")
	res,err := GetGatewayRpcClient().GetGatewayFile(context.Background(),req)
	return res,err
}

func GetPowerModelSet(req *pb.ReqGwAddtional)(*pb.ResPowerModel,error){
	log.Info("rpc/start get power model")
	res,err:= GetGatewayRpcClient().GetPowerModelSet(context.Background(),req)
	return res,err
}

func GetGatewayMedia(req *pb.ReqGwAddtional)(*pb.ResMedia ,error){
	log.Info("rpc/start Get gateway media")
	res,err:= GetGatewayRpcClient().GetGatewayMedia(context.Background(),req)
	return res,err
}

func GetAppEui(req *pb.ReqGwAddtional)(*pb.ResAppEui,error) {
	log.Info("rpc/start Get Appeui")
	res ,err:= GetGatewayRpcClient().GetAppEui(context.Background(),req)
	return res,err
}

/*网关蓝牙扫描设备信息*/
func GetBleScans(req *pb.ReqGwAddtional)(*pb.ResBleScans,error) {
	log.Info("rpc_gatewayaddtional/ start get gateway ble scans")
	res,err := GetGatewayRpcClient().GetBleScans(context.Background(),req)
	return res,err
}

