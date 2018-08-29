package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"cotx-http/result"
	"fmt"
	log "github.com/cihub/seelog"
	"cotx-http/pb"
	"cotx-http/rpcClient"
	"strconv"
	"context"
)

//获取网关的网络状态信息
func GetGatewayNetState(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")

	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid}
	//rpc 调用
	resgateway ,err:=rpcClient.GetGatewayNetState(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}

	fmt.Println("++++++++++++++++++CODE=:",resgateway.ErrCode)
	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway,res)
}
func GetGatewayState(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid}
	//rpc 调用
	resgateway ,err:=rpcClient.GetGatewayState(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}

	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway,res)
}
func GetGatewayFileState(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid}
	//rpc 调用
	resgateway ,err:=rpcClient.GetGatewayFileStat(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}

	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway,res)
}
func GetGatewayVideos(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid}
	//rpc 调用
	resgateway,err :=rpcClient.GetGatewayVideos(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}

	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway.RGV,res)
}
func GetGatewayPhotos(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {

	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid}
	//rpc 调用
	resgateway,err :=rpcClient.GetGatewayPhotos(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}

	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway.RGP,res)
}
func GetGatewayUSBStat(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid}
	//rpc 调用
	resgateway ,err:=rpcClient.GetGatewayUSBStat(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}

	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway,res)
}
func GetGatewayWifiScans(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid}
	//rpc 调用
	resgateway ,err:=rpcClient.GetGatewayWifiScans(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}

	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway.RGWS,res)
}
func GetGatewayWifiAddress(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid}
	//rpc 调用
	resgateway ,err:=rpcClient.GetGatewayWifiAddress(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30025:
			result.JsonReply("Value_IS_Error",nil,res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway,res)
}
func GetGatewayCableAddress(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid}
	//rpc 调用
	resgateway ,err:=rpcClient.GetGatewayCableAddress(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}

	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30025:
			result.JsonReply("Value_IS_Error",nil,res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway,res)
}
func GetGatewayWifiDNS(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {

	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid}
	//rpc 调用
	resgateway ,err :=rpcClient.GetGatewayWifiDNS(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30025:
			result.JsonReply("Value_IS_Error",nil,res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway,res)
}
func GetGatewayCableDNS(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid}
	//rpc 调用
	resgateway ,err :=rpcClient.GetGatewayCAbleDNS(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30025:
			result.JsonReply("Value_IS_Error",nil,res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway,res)
}
func GetGatewayMessage(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid}
	//rpc 调用
	resgateway ,err :=rpcClient.GetGatewayMessage(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway,res)
}
func GetGatewayLora(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid}
	//rpc 调用
	resgateway ,err :=rpcClient.GetGatewayLora(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway.GatewayLora,res)
}
func GetGatewayUsbWifiDNS(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	usbnumString    := param.ByName("usb")
	keyString       := req.FormValue("key")
	key,_ :=strconv.ParseInt(keyString,0,32)
	usbnum ,_:= strconv.ParseInt(usbnumString,10,32)
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid,UsbNum:int32(usbnum),Key:int32(key)}
	//rpc 调用
	resgateway,err  :=rpcClient.GetGatewayUsbWifiDNS(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30025:
			result.JsonReply("Value_IS_Error",nil,res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway,res)
}
func GetGatewayUsbWifiAddress(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	usbnumString    := param.ByName("usb")
	usbnum ,_:= strconv.ParseInt(usbnumString,10,32)
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid,UsbNum:int32(usbnum)}
	//rpc 调用
	resgateway ,err:=rpcClient.GetGatewayUsbWifiAddress(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30025:
			result.JsonReply("Value_IS_Error",nil,res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway,res)
}
func GetGatewayServerIots(res http.ResponseWriter,req *http.Request,param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid}
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	resgateway ,err:=rpcClient.GetGatewayServerIots(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}

	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway.IOTs,res)
}
func GetGatewaySwitch(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid}
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	resgateway,err :=rpcClient.GetGatewaySwitch(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}

	if resgateway.Errcode != 10000 {
		switch resgateway.Errcode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway.SwitchState,res)
}
func GetGatewayHotSpot(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid}
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	resgateway ,err:=rpcClient.GetHotSpot(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}

	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway.HotSpot,res)
}
func GetGatewayUsbWifiStat(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid}
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	resgateway ,err:=rpcClient.GetUsbWifistat(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}

	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway.WifiStat,res)
}
func GetGatewayUsbHotSpot(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	usbnumString    := param.ByName("usb")
	usbnum ,_:= strconv.ParseInt(usbnumString,10,32)
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid,UsbNum:int32(usbnum)}
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	resgateway ,err:=rpcClient.GetUsbHotSpot(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}

	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway.UsbHotSpot,res)
}
func GetGatewayUsbGCardStat(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid}
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	resgateway ,err:=rpcClient.GetUsbGCardStat(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}

	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway.GCardStat,res)
}
func GetGatewayWifi(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid}
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	resgateway ,err:=rpcClient.GetWifi(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}

	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway.Wifi ,res)
}
func GetGatewayUsbWifiScans(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return

	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	usbnumString    := param.ByName("usb")
	usbnum ,_:= strconv.ParseInt(usbnumString,10,32)
	reqgateway := &pb.ReqGateway{GatewayID:gatewayid,UserID:userinfo.Uid,UsbNum:int32(usbnum)}
	//rpc 调用
	resgateway ,err :=rpcClient.GetUsbWifiScan(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}

	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway.RGWS,res)
}

/*获取网关媒体文件的下载地址*/
func GetGatewayFile(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	log.Info("http_gateway/start get gateway file")

	//获取上下文账户信息
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return

	}

	//获取网关id
	gatewayid := param.ByName("gwid")

	//获取文件名称
	filename := param.ByName("name")
	reqgateway := &pb.ReqGetGatewayFile{GatewayId:gatewayid,UserId:userinfo.Uid,FileName:filename}

	//rpc调用
     resgateway ,err := rpcClient.GetGatewayFile(reqgateway)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}

     //结果返回
	if resgateway.ErrCode != 10000 {
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30030:
             result.JsonReply("File_IsNot_Exit",nil,res)
             return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		case 30027:
			result.JsonReply("Gateway_UnResponse",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway.GetGatewayFile,res)
}
func GetPowerModelSet(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	//================分割线==============
	gwreq := &pb.ReqGwAddtional{GatewayID:gatewayid,UserID:userinfo.Uid}
	gwres,err := rpcClient.GetPowerModelSet(gwreq)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	//================分割线=============
	log.Info("rpc-Res == :",gwres.ErrCode)
	if gwres.ErrCode != 10000 {
		switch gwres.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",gwres,res)
}

func GetGatewayMedia(res http.ResponseWriter,req *http.Request, param httprouter.Params){
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	//================分割线==============
	gwreq := &pb.ReqGwAddtional{GatewayID:gatewayid,UserID:userinfo.Uid}
	gwres ,err:= rpcClient.GetGatewayMedia(gwreq)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	//================分割线=============
	log.Info("rpc-Res == :",gwres.ErrCode)
	if gwres.ErrCode != 10000 {
		switch gwres.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",gwres,res)
}
func GetAppEui(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	//================分割线==============
	gwreq := &pb.ReqGwAddtional{GatewayID:gatewayid,UserID:userinfo.Uid}
	gwres ,err:= rpcClient.GetAppEui(gwreq)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	//================分割线=============
	log.Info("rpc-Res == :",gwres.ErrCode)
	if gwres.ErrCode != 10000 {
		switch gwres.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",gwres,res)
}
func GetBleScans(res http.ResponseWriter,req *http.Request,param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")

	//获取请求参数
	gwreq := &pb.ReqGwAddtional{GatewayID:gatewayid,UserID:userinfo.Uid}

	//rpc调用
	gwres ,err:= rpcClient.GetBleScans(gwreq)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	//响应结果返回客户端
	log.Info("rpc-Res == :",gwres.ErrCode)
	if gwres.ErrCode != 10000 {
		switch gwres.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",gwres.BleScans,res)
}

/*基于网关的id获取网关的告警统计*/
func GetWarnStat(res http.ResponseWriter,req *http.Request,param httprouter.Params){
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	gatewayid := param.ByName("gwid")
	var getWarnStatRequest = new(pb.GetWarnStatRequest)
	getWarnStatRequest.UserId = userinfo.Uid
	getWarnStatRequest.GatewayId = gatewayid
	getWarnStatResponse,err := rpcClient.GetGatewayRpcClient().GetWarnStat(context.Background(),getWarnStatRequest)
	if err != nil || getWarnStatResponse.ErrCode != 10000 {
		switch getWarnStatResponse.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",getWarnStatResponse.WarnStat,res)
}

/*基于网关的id 获取网关的流量统计(package/bytes)*/
func GetGatewayTrafficStatByGatewayId(res http.ResponseWriter,req *http.Request,param httprouter.Params){
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	gwid := param.ByName("gwid")
	var getTraficStatRequest = new(pb.GetTrafficStatByGatewayIdRequest)
	getTraficStatRequest.UserId = userinfo.Uid
	getTraficStatRequest.GatewayId = gwid
	getTrafficStatResponse,err := rpcClient.GetGatewayRpcClient().GetGatewayTrafficStatByGatewayId(context.Background(),getTraficStatRequest)
	if getTrafficStatResponse.ErrCode != 10000 || err != nil {
		switch getTrafficStatResponse.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",getTrafficStatResponse.TrafficStat,res)
}