package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	log "github.com/cihub/seelog"
	"cotx-http/pb"
	"cotx-http/result"
	"cotx-http/rpcClient"
	"fmt"
	"cotx-http/utils"
	"strconv"
)

func GetNewGatewayVersion(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	gatewayid := param.ByName("gwid")
	versionType := param.ByName("type")
	vTpye,_ :=strconv.ParseInt(versionType,0,32)
	reqgateway := &pb.ReqGetNewGatewayVersion{GatewayId:gatewayid,UserId:userinfo.Uid,VersionType:int32(vTpye)}

	//rpc调用
	resgateway,err := rpcClient.GetNewGatewayVersion(reqgateway)
	if resgateway.ErrCode != 10000 || err != nil {
		log.Error("error:",err)
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30032:
			result.JsonReply("No_Update_Device",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway.NewGatewayVersion,res)
}

func UpdateGatewayVersion(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	gatewayid := param.ByName("gwid")

	reqgateway := pb.ReqUpdateGatewayVersion{GatewayId:gatewayid,UserId:userinfo.Uid}
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", &reqgateway)
	fmt.Println("errCode = ", errCode)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}

	// rpc 调用
	resgateway,err := rpcClient.UpdateGatewayVersion(&reqgateway)

	if resgateway.ErrCode != 10000 || err != nil {
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
	result.JsonReply("Successful",nil,res)
}

func GetGatewayVersionState(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	gatewayid := param.ByName("gwid")
	versionType := param.ByName("type")
	vTpye,_ :=strconv.ParseInt(versionType,0,32)

	reqgateway := &pb.ReqGetGatewayVersionState{GatewayId:gatewayid,UserId:userinfo.Uid,VersionType:int32(vTpye)}

	//rpc调用
	resgateway,err := rpcClient.GetGatewayVersionState(reqgateway)
	if resgateway.ErrCode != 10000 || err != nil {
		log.Error("error:",err)
		switch resgateway.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30032:
			result.JsonReply("No_Update_Device",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",resgateway.VersionSate,res)
}

