package controller

import (
	"github.com/julienschmidt/httprouter"
	"fmt"
	"net/http"
	"cotx-http/pb"
	log "github.com/cihub/seelog"
	"cotx-http/rpcClient"
	"cotx-http/result"
	"strconv"
)

func GetUsbWifiWlan(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
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
	reqgateway := &pb.ReqGatewayUsb{GatewayId:gatewayid,UserId:userinfo.Uid,UsbId:int32(usbnum)}
	//rpc 调用
	resgateway ,err:=rpcClient.GetUsbWifiWlan(reqgateway)
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
	result.JsonReply("Successful",resgateway.UsbWifiWlan,res)
}
func GetUsbHotSpotUser(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
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
	reqgateway := &pb.ReqGatewayUsb{GatewayId:gatewayid,UserId:userinfo.Uid,UsbId:int32(usbnum)}
	//rpc 调用
	resgateway ,err:=rpcClient.GetUsbHotsptUser(reqgateway)
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
	result.JsonReply("Successful",resgateway.HotSpotUser,res)
}
func GetWifiHotSpotUser(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqgateway := &pb.ReqGatewayUsb{GatewayId:gatewayid,UserId:userinfo.Uid}
	//rpc 调用
	resgateway ,err:=rpcClient.GetWifiHotSpotUser(reqgateway)
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
	result.JsonReply("Successful",resgateway.HotSpotUser,res)
}
func GetUsbNumWifiScan(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	reqgateway := &pb.ReqGatewayUsb{GatewayId:gatewayid,UserId:userinfo.Uid}
	//rpc 调用
	resgateway ,err:=rpcClient.GetUsbNumWifiScan(reqgateway)
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
	result.JsonReply("Successful",resgateway.UsbWifiScan,res)
}