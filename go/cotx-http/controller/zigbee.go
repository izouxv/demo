package controller

import (
	log "github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"cotx-http/pb"
	"cotx-http/result"
	"cotx-http/rpcClient"
	"fmt"
	"cotx-http/utils"
)

func GetZigbee(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	gatewayid := param.ByName("gwid")

	reqzigbee := &pb.ReqGetZigbee{UserId:userinfo.Uid,GatewayId:gatewayid}

	//rpc调用
	reszigbee,err := rpcClient.GetZigbee(reqzigbee)
	if reszigbee.ErrCode != 10000 || err != nil {
		log.Error("error:",err)
		switch reszigbee.ErrCode {
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
	result.JsonReply("Successful",reszigbee.Zigbee,res)

}

func UpdateZigbee(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	gatewayid := param.ByName("gwid")

	reqzigbee := &pb.ReqUpdateZigbee{UserId:userinfo.Uid,GatewayId:gatewayid}
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", &reqzigbee)
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
	reszigbee,err := rpcClient.UpdateZigbee(reqzigbee)
	if reszigbee.ErrCode != 10000 || err != nil {
		log.Error("error:",err)
		switch reszigbee.ErrCode {
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

	result.JsonReply("Successful",reszigbee,res)

}