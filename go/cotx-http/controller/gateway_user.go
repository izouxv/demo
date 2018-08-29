
package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	log "github.com/cihub/seelog"
	"cotx-http/result"
	"fmt"
	"cotx-http/utils"
	"cotx-http/rpcClient"
	"cotx-http/po"
	"cotx-http/pb"
)

//用户绑定网关
func RegistrationGateway(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	log.Info("gateway_user.go/registration gateway with user")
    userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
        log.Info("没有找到帐号")
        result.JsonReply("Account_abnormality",nil,res)
	}
	fmt.Println("userinfo= :",userinfo)
	reqgwuser := &pb.ReqGwUser{}
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", reqgwuser)
	fmt.Println("errCode=:",errCode)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	fmt.Println("gateway_mac==",reqgwuser.MAC)
    reqgwuser.UserID = userinfo.Uid
	//rpc调用 绑定网关方法
	resgwuser ,err:= rpcClient.BindGateway(reqgwuser)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	log.Info("res gateway user :=",resgwuser.ErrCode)
	if resgwuser.ErrCode != 10000 {
		switch resgwuser.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30018:
			result.JsonReply("Gateway_havebeen_binding",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	gatewaypo := po.UserGateway{
		GatewayId:resgwuser.GatewayID,
		AppId:resgwuser.AppEUI,
	}
	result.JsonReply("Successful",gatewaypo,res)
}
//用户授权网关给其他帐号
func BingAccountWithGw (res http.ResponseWriter,req *http.Request, param httprouter.Params){
	log.Info("start bing account with gateway")
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	reqgwuser := &pb.ReqGwUser{UserID:userinfo.Uid,}
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", reqgwuser)
	fmt.Println("errCode=:",errCode)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	fmt.Println("reqgwuser =",reqgwuser)
	//rpc 调用进行账户的授权
	resgwuser ,err:= rpcClient.BingAccount(reqgwuser)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	log.Info("=======================+%d",resgwuser.ErrCode)
	if resgwuser.ErrCode != 10000 {
		switch resgwuser.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30019:
			result.JsonReply("Account_Auth",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		case 30034:
			result.JsonReply("Insufficient_permissions",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",nil,res)
}
//删除授权帐号
func DeletAuthoriseAccount(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	log.Info("start delet account with gateway")
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	reqgwuser := &pb.ReqGwUser{UserID:userinfo.Uid,}
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", reqgwuser)
	fmt.Println("errCode=:",errCode)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	fmt.Println("reqgwuser =",reqgwuser)
	//调用rpc完成 删除授权用户
	resgwuser ,err:= rpcClient.DeletAuthoriseAccount(reqgwuser)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	if resgwuser.ErrCode != 10000 {
		switch resgwuser.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
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
//展示授权帐号
func ShowAuthoriseAccount(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	log.Info("start show authorise account with gateway")
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	gatewayid := param.ByName("gwid")
	reqgwuser := &pb.ReqGwUser{UserID:userinfo.Uid,GatewayID:gatewayid}
	//rpc调用
	resgwuser ,err:= rpcClient.ShowAuthoiseAccount(reqgwuser)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	if resgwuser.ErrCode != 10000 {
		switch resgwuser.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	AuthAccounts := []po.AuthoriseAccount{}
	for _,account:= range resgwuser.Ra {
		authaccount := po.AuthoriseAccount{
			Uid:account.UserID,
			UserName:account.UserName,
			Avatar:account.Acatar,
			NickName:account.NickName,
		}
		AuthAccounts = append(AuthAccounts,authaccount)
	}
	result.JsonReply("Successful",AuthAccounts,res)
}
//用户解绑网关
func UnwoundGateway(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	log.Info("start unwound gateway")
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	reqgwuser := &pb.ReqGwUser{UserID:userinfo.Uid,}
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", reqgwuser)
	fmt.Println("errCode=:",errCode)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	fmt.Println("reqgwuser =",reqgwuser)
	//rpc 调用
	resgwuser ,err:= rpcClient.UnwoundBindGateway(reqgwuser)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	if resgwuser.ErrCode != 10000 {
		switch resgwuser.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
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
//验证网关的是否已经被绑定
func ValidationGateway (res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	log.Info("start validation gateway")
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	reqgwuser := &pb.ReqGwUser{UserID:userinfo.Uid,}
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", reqgwuser)
	fmt.Println("errCode=:",errCode)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	fmt.Println("reqgwuser =",reqgwuser)
	//rpc 调用
	resgwuser,err := rpcClient.ValidationGateway(reqgwuser)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	if resgwuser.ErrCode != 10000 {
		switch resgwuser.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30018:
			result.JsonReply("Gateway_havebeen_binding",nil, res)
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
func ValidationGatewayAccount(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	log.Info("start validation gateway account")
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	reqgwuser := &pb.ReqGwUser{UserID:userinfo.Uid,}
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", reqgwuser)
	fmt.Println("errCode=:",errCode)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	fmt.Println("reqgwuser =",reqgwuser)
	//rpc 调用
	resgwuser ,err:= rpcClient.ValidationGatewayAccount(reqgwuser)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	if resgwuser.ErrCode != 10000 {
		switch resgwuser.ErrCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 30019:
			result.JsonReply("Account_Auth",nil, res)
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
//获取用户绑定的所有网关的位置信息
func ShowAllUserGws(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	log.Info("start validation gateway account")
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	resusergw :=&pb.ReqGwUser{UserID:userinfo.Uid}
	//cotx-rpc 调用
	resallgws ,err:= rpcClient.ShowAllUserGateway(resusergw)
	if err != nil {
		result.JsonReply("SysTem_ERROR",nil, res)
		return
	}
	if resallgws.ErrCode != 10000 {
		switch resallgws.ErrCode {
		case 10001:
			result.JsonReply("System_error", nil, res)
			return
		case 30020:
			result.JsonReply("System_error", nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		default:
			result.JsonReply("Unknown_error", nil, res)
			return
		}
	}
		result.JsonReply("Successful",resallgws.RUGS,res)
}
func ShowNodesByGateway(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	log.Info("start get nodes by gayewayid")
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	gatewayid := param.ByName("gwid")
	resusergw :=&pb.ReqGwUser{UserID:userinfo.Uid,GatewayID:gatewayid}
	//cotx-rpc 调用
	resgateway ,err:= rpcClient.ShowNodesByGateway(resusergw)
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
	result.JsonReply("Successful",resgateway.AllNodes,res)
}
func ShowAllNodesPos(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	log.Info("start get all nodes postion")
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	fmt.Println("userinfo= :",userinfo)
	resusergw :=&pb.ReqGwUser{UserID:userinfo.Uid}
	//cotx-rpc 调用
	resgateway ,err:= rpcClient.ShowAllNodesPos(resusergw)
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
	result.JsonReply("Successful",resgateway.AllNodesPos,res)
}
func ShowAllGws(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	log.Info("start get All gateways")
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	resusergw :=&pb.ReqGwUser{UserID:userinfo.Uid,}
	//cotx-rpc 调用
	resgateway ,err:= rpcClient.ShowAllGws(resusergw)
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
	result.JsonReply("Successful",resgateway.ShowAllGws,res)
}