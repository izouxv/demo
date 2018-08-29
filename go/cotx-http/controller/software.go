package controller

import (
	"github.com/julienschmidt/httprouter"
	"cotx-http/result"
	log "github.com/cihub/seelog"
	"net/http"
	"cotx-http/pb"
	"strconv"
	"cotx-http/rpcClient"
	"context"
	"cotx-http/utils"
)

func GetSoftWareByversionType(res http.ResponseWriter,req *http.Request,param httprouter.Params)   {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	countValue := req.FormValue("count")
	pageValue  := req.FormValue("page")
	order_by   := req.FormValue("order_by")
	gatewayid := param.ByName("gwid")
	count,_ := strconv.Atoi(countValue)
	page,_ := strconv.Atoi(pageValue)
	version_typeValue := param.ByName("type")
	version_type ,_ :=strconv.Atoi(version_typeValue)
	var getSoftwareByVersionTypeRequest  = new(pb.GetSoftwareByVersionTypeRequest)
	getSoftwareByVersionTypeRequest.Count = int32(count)
	getSoftwareByVersionTypeRequest.Page = int32(page)
	getSoftwareByVersionTypeRequest.OrderBy = order_by
	getSoftwareByVersionTypeRequest.SoftwareType =  int32(version_type)
	getSoftwareByVersionTypeRequest.UserId = userinfo.Uid
	getSoftwareByVersionTypeRequest.GatewayId = gatewayid
	getSoftwareByversionTypeResponse,_ := rpcClient.GetSoftwareRpcClient().GetSoftwareByVersionType(context.Background(),getSoftwareByVersionTypeRequest)
	if getSoftwareByversionTypeResponse.ErrCode != 10000 {
		switch getSoftwareByversionTypeResponse.ErrCode {
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
	result.JsonReply("Successful",getSoftwareByversionTypeResponse,res)
}
func InstallSoftware(res http.ResponseWriter,req *http.Request,param httprouter.Params)   {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	var installSoftware = new(pb.InstallSoftwareRequest)
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", &installSoftware)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	var gatewayid = param.ByName("gwid")
	installSoftware.UserId= userinfo.Uid
	installSoftware.GatewayId = gatewayid
	installSoftwareResponse,_ := rpcClient.GetSoftwareRpcClient().InstallSoftware(context.Background(),installSoftware)
	if installSoftwareResponse.ErrCode != 10000 {
		switch installSoftwareResponse.ErrCode {
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
	result.JsonReply("Successful",installSoftwareResponse.Software,res)
}

func GetSoftwareById(res http.ResponseWriter,req *http.Request,param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	var getSoftwareByIdRequest = new(pb.GetSoftwareByIdRequest)
	idString := param.ByName("id")
	id,_ := strconv.Atoi(idString)
	gatewayId:= param.ByName("gwid")
	getSoftwareByIdRequest.UserId = userinfo.Uid
	getSoftwareByIdRequest.Id = int64(id)
	getSoftwareByIdRequest.GatewayId = gatewayId
	getSoftwareByIdResponse,_ := rpcClient.GetSoftwareRpcClient().GetSoftwareById(context.Background(),getSoftwareByIdRequest)
	if getSoftwareByIdResponse.ErrCode != 10000 {
		switch getSoftwareByIdResponse.ErrCode {
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
	result.JsonReply("Successful",getSoftwareByIdResponse.Software,res)
}

