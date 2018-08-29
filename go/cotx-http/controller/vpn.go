package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	log "github.com/cihub/seelog"
	"cotx-http/result"
	"cotx-http/pb"
	"cotx-http/rpcClient"
	"fmt"
	"cotx-http/utils"
)

func GetVpn(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	gatewayid := param.ByName("gwid")
	reqvpn := &pb.ReqGetVpn{UserId:userinfo.Uid,GatewayId:gatewayid}

	//rpc调用
	resvpn,err := rpcClient.GetVpn(reqvpn)
	if resvpn.ErrCode != 10000 || err != nil {
		log.Error("error:",err)
		switch resvpn.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
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
	result.JsonReply("Successful",resvpn.Vpn,res)

}

func UpdateVpn(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	gatewayid := param.ByName("gwid")
	reqvpn := &pb.ReqUpdateVpn{UserId:userinfo.Uid,GatewayId:gatewayid}
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", &reqvpn)
	fmt.Println("errCode = ", errCode)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	//rpc调用
	resvpn,err := rpcClient.UpdateVpn(reqvpn)
	if resvpn.ErrCode != 10000 || err != nil {
		log.Error("error:",err)
		switch resvpn.ErrCode {
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
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
	result.JsonReply("Successful",resvpn,res)
}

func PingGateway(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	gatewayid := param.ByName("gwid")
	ping := &pb.ReqPingGateway{UserId:userinfo.Uid,GatewayId:gatewayid}
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", &ping)
	fmt.Println("errCode = ", errCode)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	//rpc调用
	resping,err := rpcClient.PingGateway(ping)
	if resping.ErrCode != 10000 || err != nil {
		log.Error("error:",err)
		switch resping.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
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
	result.JsonReply("Successful",resping,res)

}