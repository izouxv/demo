package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/cihub/seelog"
	"cotx-http/result"
	"cotx-http/utils"
	"cotx-http/pb"
	"cotx-http/rpcClient"
	"cotx-http/po"
	"fmt"
	"strconv"
	"os"
	"io"
)

//登陆
func Login(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("Start Func Login: ",params)
	ssoreq := pb.SsoRequest{}
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", &ssoreq)
	fmt.Println("errCode = ", errCode)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	//获取source
	ssoreq.Source = req.Header.Get("source")
	log.Info("Login-ssoJson:",ssoreq)
	fmt.Println("ssoreq = ", ssoreq)
	//非空校验
	if utils.VerifyParamsEmpty(ssoreq.Password) {
		result.ResCode(20015, res)
		return
	}

	// 参数校验
	if utils.VerifyUsername(ssoreq.Username) {
		result.JsonReply("Username_is_incorrect_or_empty",nil, res)
		return
	}
	if utils.VerifyPassword(ssoreq.Password) {
		result.JsonReply("Password_is_incorrect_or_empty",nil, res)
		return
	}
	if ssoreq.Source == "" {
		log.Error("Get source Failed")
		result.JsonReply("Source_is_incorrect_or_empty", nil, res)
		return
	}
	//调用sso-rpc Login
	ssoReply := rpcClient.Login(&ssoreq)
	log.Info("ssoReply:",ssoReply)
	fmt.Println("ssoReply.ErrorCode ", ssoReply.ErrorCode)
	if ssoReply.ErrorCode != 10000 {
		switch ssoReply.ErrorCode {
		case 33002:
			result.JsonReply("Account_does_not_existed",nil, res)
			return
		case 33005:
			result.JsonReply("Token_is_incorrect_or_lose_efficacy",nil, res)
			return
		case 33006:
			result.JsonReply("System_error",nil, res)
			return
		case 33007:
			result.JsonReply("Account_does_not_activated",nil, res)
			return
		case 33009:
			result.JsonReply("Password_mismatched",nil, res)
			return
		default :
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	//调用account-rpc GetUserInfoAll
	accountReply := rpcClient.GetUserInfoAll(&pb.AccountRequest{Source:ssoreq.Source, Uid:ssoReply.Uid})
	if accountReply.ErrorCode != 10000 {
		switch accountReply.ErrorCode {
		case 10001:
		case 33001:
		case 33010:
			result.JsonReply("System_error",nil, res)
			return
		case 33002:
			result.JsonReply("Account_does_not_existed",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	//返回JSON数据
	loginPo := po.LoginPo{}
	loginPo.SetSsoReplyIntoLoginPo(ssoReply)
	loginPo.SetAccountReplyIntoLoginPo(accountReply)
	result.JsonReply("Successful", loginPo, res)
	return
}

//获取用户信息
func GetUserInfo(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	//从context中读取userinfo信息
	log.Info("user.go/GetUserInfo: began ...")
	req.ParseForm()
	source := req.Header.Get("source")
	token  := req.Header.Get("token")
	//非空校验
	if source == ""{
		log.Info("user.go/GetUserInfo: source is empty")
		result.JsonReply("Source_is_incorrect_or_empty",nil, res)
		return
	}
	userInfo, ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Error("user.go/GetUserInfo: context get userinfo assert error!")
		result.JsonReply("System_error", nil, res)
		return
	}
	log.Info("user.go/GetUserInfo: context userInfo:", userInfo)
	//开始调用rpc
	ssoR := pb.SsoRequest{SessionName:token, Source:source}
	log.Info("user.go/GetUserInfo:  ssoR = ", ssoR)
	ssoReply := rpcClient.GetUserInfo(&ssoR)
	//accountReply := rpcClient.GetUserInfoAll(&pb.AccountRequest{Source:source, Uid:accountReply.Uid})
	//调用rpc进行token校验
	//reply := rpcClient.GetUserInfo(&pb.SsoRequest{Source:source, SessionName:token})
	log.Info("user.go/GetUserInfo:  after rpc ssoR = ", ssoR)
	if ssoReply.ErrorCode != 10000 {
		log.Error("user.go/GetUserInfo:	rpc GetUserInfo err")
		result.JsonReply("System_error", nil, res)
		return
	}
	//调用account-rpc GetUserInfoAll
	accountReply := rpcClient.GetUserInfoAll(&pb.AccountRequest{Source:source, Uid:ssoReply.Uid})
	if accountReply.ErrorCode != 10000 {
		switch accountReply.ErrorCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 33010:
			result.JsonReply("Source_is_incorrect_or_empty",nil, res)
			return
		case 33002:
			result.JsonReply("Account_does_not_existed",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	//返回JSON数据
	loginPo := po.LoginPo{}
	loginPo.SetSsoReplyIntoLoginPo(ssoReply)
	loginPo.SetAccountReplyIntoLoginPo(accountReply)
	loginPo.Token = token
	result.JsonReply("Successful", loginPo, res)
	return
}

//注册
func Register(res http.ResponseWriter,req *http.Request, param httprouter.Params) {
	log.Info("user.go/Register: Register......")
	ssoreq := pb.SsoRequest{CodeType:1}
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", &ssoreq)
	if errCode != 10000 {
		if errCode == 404 {
			log.Error("Register_errCode == 404")
			result.ResCode(http.StatusNotFound,res)
			return
		}
		log.Error("Register_errCode == ", errCode)
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	ssoreq.Source = req.Header.Get("source")
	ssoreq.State  = 3
	ssoreq.Salt	  = string(utils.Krand(6, utils.KC_RAND_KIND_ALL))
	log.Info("Register-ssoJson:",ssoreq)
	//参数校验
	if ssoreq.Source == "" {
		log.Error("Get source Failed")
		result.JsonReply("Source_is_incorrect_or_empty", nil, res)
		return
	}
	if utils.VerifyUsername(ssoreq.Username) {
		result.JsonReply("Username_is_incorrect_or_empty",nil, res)
		return
	}
	if utils.VerifyPassword(ssoreq.Password) {
		result.JsonReply("Password_is_incorrect_or_empty",nil, res)
		return
	}
	if utils.VerifyNickname(ssoreq.Nickname) {
		result.JsonReply("Nickname_is_incorrect_or_empty",nil, res)
		return
	}

	//手机号注册校验验证码
	if utils.REGEXP_MOBILE.MatchString(ssoreq.Username){
		if ssoreq.Code == "" {
			log.Error("Get MobileCode Failed")
			result.JsonReply("MobileCode_is_incorrect_or_empty", nil, res)
			return
		}
		//调用sso-rpc CheckCode
		ssoReplyC := rpcClient.CheckCode(&ssoreq)
		log.Info("CheckCode ssoReply:",ssoReplyC)
		if ssoReplyC.ErrorCode != 10000 {
			switch  ssoReplyC.ErrorCode {
			case 33010:
				result.JsonReply("System_error", nil, res)
				return
			case 33005:
				result.JsonReply("MobileCode_is_incorrect_or_lose_efficacy", nil, res)
				return
			case 33008:
				result.JsonReply("Account_already_existed", nil, res)
				return
			default:
				result.JsonReply("Unknown_error", nil, res)
				return
			}
		}
	}

	//调用sso-rpc Register
	ssoReply := rpcClient.Register(&ssoreq)
	log.Info("Register ssoReply:",ssoReply)
	if ssoReply.ErrorCode != 10000 {
		switch ssoReply.ErrorCode {
		case 10001:
			result.JsonReply("System_error", nil, res)
			return
		case 33003:
			result.JsonReply("System_error", nil, res)
			return
		case 33010:
			result.JsonReply("System_error", nil, res)
			return
		case 33008:
			result.JsonReply("Account_already_existed", nil, res)
			return
		default:
			result.JsonReply("Unknown_error", nil, res)
			return
		}
	}
	result.JsonReply("Successful",nil, res)
	return
}

//判断用户名是否占用
func JudgeUsername(res http.ResponseWriter,req *http.Request, param httprouter.Params) {
	log.Info("user.go/JudgeUsername:  began ...")
	ssoreq := pb.SsoRequest{}
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", &ssoreq)
	if errCode != 10000 {
		if errCode == 404 {
			log.Error("Register_errCode == 404")
			result.ResCode(http.StatusNotFound,res)
			return
		}
		log.Error("JudgeUsername_errCode == ", errCode)
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	ssoreq.Source = req.Header.Get("source")
	log.Info("Register-ssoJson:",ssoreq)
	//参数校验
	if utils.VerifyUsername(ssoreq.Username) {
		result.JsonReply("Username_is_incorrect_or_empty",nil, res)
		return
	}
	if ssoreq.Source == "" {
		log.Error("Get source Failed")
		result.JsonReply("Source_is_incorrect_or_empty", nil, res)
		return
	}

	//调用sso-rpc JudgeUsername
	ssoReply := rpcClient.JudgeUsername(&ssoreq)
	log.Info("Register ssoReply:",ssoReply)
	switch ssoReply.ErrorCode {
	case 33002:
		result.JsonReply("Account_does_not_existed", nil, res)
		return
	case 33008:
		result.JsonReply("Account_already_existed", nil, res)
		return
	default:
		result.JsonReply("Unknown_error", nil, res)
		return
	}
}

//修改用户信息
func UpdateUserInfo(res http.ResponseWriter,req *http.Request, param httprouter.Params) {
	log.Info("Start UpdateUserInfo")
	accountreq := pb.AccountRequest{}
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", &accountreq)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	userInfo, ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Error("context get userInfo assert error!")
		result.JsonReply("System_error", nil, res)
		return
	}
	accountreq.Uid = userInfo.Uid
	accountreq.Source = req.Header.Get("source")
	log.Info("Login-ssoJson:",accountreq)

	//参数校验
	if accountreq.Nickname != "" {
		if !(utils.REGEXP_NICKNAME.MatchString(accountreq.Nickname)) {
			log.Debug("UpdateUserInfo, Nickname_is_incorrect_or_empty")
			result.JsonReply("Nickname_is_incorrect_or_empty",nil, res)
			return
		}
	}
	//调用account-rpc
	accountReply := rpcClient.UpdateUserInfo(&accountreq)
	if accountReply.ErrorCode != 10000 {
		switch accountReply.ErrorCode {
		case 10001:
			result.JsonReply("System_error",nil, res)
			return
		case 33003:
			result.JsonReply("System_error",nil, res)
			return
		case 33010:
			result.JsonReply("System_error",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	//result.JsonReply("Successful",nil ,res)
	//return
	//用户信息更新后，更新redis的内容
	//调用account-rpc
	source := req.Header.Get("source")
	token := req.Header.Get("token")
	ssoR := pb.SsoRequest{Token:token, Source:source, Uid:accountreq.Uid, Username:accountreq.Username, Nickname:accountreq.Nickname,State:accountreq.State}
	ssoReply := rpcClient.UpdateRedis(&ssoR)
	if ssoReply.ErrorCode != 10000 {
		switch ssoReply.ErrorCode {
		case 33010:
			result.JsonReply("System_error",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	//
	result.JsonReply("Successful",nil ,res)
	return
}

//登出
func Logout(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	log.Info("Start Logout")

	userInfo, ok := req.Context().Value("userInfo").(*pb.SsoReply)
	fmt.Println("userinfo = ", userInfo)
	fmt.Println("ok = ", ok)
	log.Info("Context UserInfo:",userInfo," ok:",ok)
	if !ok {
		log.Error("context get userInfo error!")
		result.JsonReply("System_error", nil, res)
		return
	}
	ssoR := pb.SsoRequest{SessionName:userInfo.SessionName,Source:req.Header.Get("source")}
	fmt.Println("ssoR: ", ssoR)
	//调用rpc
	ssoReply := rpcClient.Logout(&ssoR)
	log.Info("ssoReply:",ssoReply)
	if ssoReply.ErrorCode != 10000 {
		switch ssoReply.ErrorCode {
		case 10001:
		case 33001:
		case 33006:
			result.JsonReply("System_error", nil, res)
			return
		default:
			result.JsonReply("Unknown_error", nil, res)
			return
		}
	}
	result.JsonReply("Successful",nil, res)
	return
}

/*
//发送验证码
func SendCode(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	log.Info("Start SendCode")
	ssoR := pb.SsoRequest{}
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", &ssoR)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	ssoR.Source = req.Header.Get("source")
	//参数校验
	if utils.VerifyUsername(ssoR.Username) {
		result.JsonReply("Username_is_incorrect_or_empty",nil, res)
		return
	}
	if ssoR.Source == "" {
		log.Error("Get source Failed")
		result.JsonReply("Source_is_incorrect_or_empty", nil, res)
		return
	}
	if ssoR.CodeType == 0 {
		log.Error("Get CodeType Failed")
		result.JsonReply("CodeType_is_incorrect_or_empty", nil, res)
		return
	}
	//调用rpc
	ssoReply := rpcClient.SendMobileCode(&ssoR)
	if ssoReply.ErrorCode != 10000 {
		switch ssoReply.ErrorCode {
		case 10001:
		case 33001:
		case 33005:
		case 33006:
		case 33010:
			result.JsonReply("System_error", nil, res)
			return
		case 33002:
			result.JsonReply("Account_does_not_existed", nil, res)
			return
		case 33008:
			result.JsonReply("Account_already_existed", nil, res)
			return
		default:
			result.JsonReply("Unknown_error", nil, res)
			return
		}
	}
	result.JsonReply("Successful",nil, res)
	return
}

*/
//发送验证码(注册、重置密码)
func SendCode(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	log.Info("user.go/SendCode:  began ......")
	// todo 判断是注册or重置密码
	action := param.ByName("action")
	ssoR, flag := req.Context().Value("contextSso").(*pb.SsoRequest)
	log.Info("user.go/SendCode:  ssoR: ", ssoR, " flag: ",flag)
	if !flag || ssoR == nil {
		log.Error("user.go/SendCode:  context get userInfo assert error!")
		result.JsonReply("System_error", nil, res)
		return
	}
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", &ssoR)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	log.Info("user.go/SendCode:  bodyJson = ", ssoR)
	if utils.VerifyParamsEmpty(ssoR.Username) {
		result.JsonReply("Password_is_incorrect_or_empty", nil, res)
		return
	}
	//验证username(手机号或邮箱)
	log.Info("user.go/SendCode:  ssoR.Username =  ", ssoR.Username)
	mobileRegexp := !utils.REGEXP_MOBILE.MatchString(ssoR.Username)
	mailRegexp := !utils.REGEXP_MAIL.MatchString(ssoR.Username)
	if mobileRegexp && mailRegexp {
		log.Error("user.go/SendCode: ssoR.Username is_incorrect_or_empty ")
		result.JsonReply("Body_is_incorrect_or_empty", nil, res)
		return
	}
	//1：注册    2:发送验证码   3：校验发送验证码
	if "send" == action {
		ssoR.CodeType = 2  //rpc层有进行验证CodeType是否为： ""
		//调用rpc
		ssoReply := rpcClient.SendMobileCode(ssoR)
		if ssoReply.ErrorCode != 10000 {
			switch ssoReply.ErrorCode {
			case 33005:
				result.JsonReply("System_error",nil, res)
				return
			default:
				result.JsonReply("Unknown_error",nil, res)
				return
			}
		}
		result.JsonReply("Successful",nil ,res)
		return
	}
	if "check" == action {
		ssoR.CodeType = 3  //rpc层有进行验证CodeType是否为： ""
		//调用rpc
		ssoReply := rpcClient.SendMobileCode(ssoR)
		if ssoReply.ErrorCode != 10000 {
			switch ssoReply.ErrorCode {
			case 33004:
				result.JsonReply("MobileCode_is_incorrect_or_empty",nil, res)
				return
			default:
				result.JsonReply("Unknown_error",nil, res)
				return
			}
		}
		result.JsonReply("Successful",nil ,res)
		return
	}

/*
	//1：注册    2:手机注册发送验证码   3：重置密码发送验证码 	4：校验验证码
	if "send" == action {
		ssoR.CodeType = 2  //rpc层有进行验证CodeType是否为： ""

	} else if "check" == action {
		ssoR.CodeType = 3  //rpc层有进行验证CodeType是否为： ""
	} else {
		result.ResCode(404, res)
		return
	}
	//调用rpc
	ssoReply := rpcClient.SendMobileCode(ssoR)
	if ssoReply.ErrorCode != 10000 {
		switch ssoReply.ErrorCode {
		case 33002:
			result.JsonReply("Account_does_not_existed",nil, res)
			return
		case 33008:
			result.JsonReply("Account_already_existed",nil, res)
			return
		case 33010:
			result.JsonReply("System_error",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",nil ,res)
	return
*/

/*
	if "check" == action {
		fmt.Println("0000    ssoR = ", ssoR)
		//1是注册  2是发送验证码  3是验证验证码
		ssoR.CodeType = 3
		ssoReply := rpcClient.CheckCode(ssoR)
		if ssoReply.ErrorCode != 10000 {
			result.JsonReply("Unknown_error", nil, res)
			return
		}
		return
	}
*/

}

//发送邮箱校验与密码重置（未登录）
func SendMail(res http.ResponseWriter,req *http.Request, param httprouter.Params)  {
	log.Info("user.go/SendMail:  began ......")
	action := param.ByName("action")
	ssoR, flag := req.Context().Value("contextSso").(*pb.SsoRequest)
	log.Info("user.go/SendMail:  ssoR: ", ssoR, " flag: ",flag)
	if !flag || ssoR == nil {
		log.Error("user.go/SendMail:  context get userInfo assert error!")
		result.JsonReply("System_error", nil, res)
		return
	}
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", &ssoR)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	log.Info("user.go/SendMail:  bodyJson = ", ssoR)
	if utils.VerifyParamsEmpty(ssoR.Username) {
		result.JsonReply("Username_is_incorrect_or_empty", nil, res)
		return
	}
	//验证username
	log.Info("user.go/SendMail:  ssoR.Username =  ", ssoR.Username)
	//mobileRegexp := !utils.REGEXP_MOBILE.MatchString(ssoR.Username)
	mailRegexp := !utils.REGEXP_MAIL.MatchString(ssoR.Username)
	if mailRegexp {
		log.Error("user.go/SendMail: ssoR.Username is_incorrect_or_empty ")
		result.JsonReply("Body_is_incorrect_or_empty", nil, res)
		return
	}
	//1：注册  2:手机注册发送验证码  3：重置密码发送验证码
	ssoReply := new(pb.SsoReply)
	if "send" == action {
		//调用rpc
		ssoReply = rpcClient.FindPasswordByMail(ssoR)
	} else if "check" == action {
		ssoR.Salt = string(utils.Krand(6, utils.KC_RAND_KIND_ALL))
		ssoReply = rpcClient.ResetPassword(ssoR)
	} else {
		result.ResCode(404, res)
		return
	}
	//调用rpc
	if ssoReply.ErrorCode != 10000 {
		switch ssoReply.ErrorCode {
		case 33002:
			result.JsonReply("Account_does_not_existed",nil, res)
			return
		case 33008:
			result.JsonReply("Account_already_existed",nil, res)
			return
		case 33010:
			result.JsonReply("System_error",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",nil ,res)
	return
}

//修改密码
func UpdatePassword(res http.ResponseWriter,req *http.Request, param httprouter.Params) {
	log.Info("Start UpdatePassword")
	ssoR := pb.SsoRequest{}
	//获取source
	ssoR.Source = req.Header.Get("source")
	uidStr := param.ByName("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil{
		result.JsonReply("Parameter_format_error",nil, res)
		return
	}
	//ssoreq := pb.SsoRequest{}
	userPwd := po.UserPwd{}
	//获取HTTP请求体里的数据到userPwd
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", &userPwd)
	//获取uid
	ssoR.Uid = int32(uid)
	ssoR.SessionName = req.Header.Get("token")
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	sso, ok := req.Context().Value("userInfo").(*pb.SsoReply)

	// sso, ok := utils.GetContext(req)
	//req.Context().Value("userInfo").(*pb.SsoReply)

	fmt.Println("sso = ", sso)
	if !ok {
		log.Error("context get userInfo assert error!")
		result.JsonReply("System_error", nil, res)
		return
	}
	log.Info("UpdatePassword-ssoJson:", userPwd)
	//sso.Password = userPwd.Password
	//调用ssoRpc-CheckPassword
	//原密码非空校验
	if utils.VerifyParamsEmpty(userPwd.Password) {
		// result.ResCode(20015, res)
		result.JsonReply("Password_is_incorrect_or_empty", nil, res)
		return
	}
	//原密码校验一致性
	ssoR.Password = userPwd.Password
	fmt.Println("Request ssoR: ", ssoR)
	ssoReply := rpcClient.CheckPassword(&ssoR)
	if ssoReply.ErrorCode != 10000 {
		result.JsonReply("Password_mismatched", nil, res)
		return
	}
	//新密码非空校验
	if utils.VerifyParamsEmpty(userPwd.NewPassword) {
		result.JsonReply("Password_is_incorrect_or_empty", nil, res)
		return
	}
	//验证新密码的格式
	if utils.VerifyPassword(userPwd.NewPassword) {
		result.JsonReply("Password_is_incorrect_or_empty",nil, res)
		return
	}
	//新密码替换原密码
	ssoR.Password = userPwd.NewPassword
	ssoR.Salt = string(utils.Krand(6, utils.KC_RAND_KIND_ALL))
	ssoReply = rpcClient.UpdatePassword(&ssoR)
	if ssoReply.ErrorCode != 10000 {
		result.JsonReply("Unknown_error", nil, res)
		return
	}
	result.JsonReply("Successful", nil, res)
}

//手机重置密码
func ResetPassword(res http.ResponseWriter,req *http.Request, param httprouter.Params) {
	log.Info("user.go/ResetPassword:  began  ....")
	//var userPwd = new(po.UserPwd)
	//var resetPwd = new(po.ResetPwd)
	ssoR := pb.SsoRequest{}
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", &ssoR)
	ssoR.Source = req.Header.Get("source")
	log.Info("ssoR: ",ssoR )
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Unknown_error", nil, res)
		return
	}
	//log.Info("ResetPassword-updatepwdJson:",userPwd)
	//验证pwd， username， code 是否为空
	if utils.VerifyParamsEmpty(ssoR.Password, ssoR.Code, ssoR.Username) {
		result.JsonReply("Parameter_format_error", nil, res)
		return
	}
	if utils.VerifyPassword(ssoR.Password) {
		result.JsonReply("Parameter_format_error", nil, res)
		return
	}
	if utils.VerifyUsername(ssoR.Username) {
		result.JsonReply("Parameter_format_error", nil, res)
	}
	//数据库里查询用户信息，若没查到，提示注册，查到了进行密码更新
	//调用rpc-ResetPasswordByPhone
	ssoReply := rpcClient.ResetPasswordByPhone(&ssoR)
	if ssoReply.ErrorCode != 10000 {
		switch ssoReply.ErrorCode {
		case 33001:
			result.JsonReply("Parameter_format_error", nil, res)
			return
		case 33002:
			result.JsonReply("Account_does_not_existed",nil, res)
			return
		case 33004:
			result.JsonReply("MobileCode_is_incorrect_or_lose_efficacy",nil, res)
			return
		case 33010:
			result.JsonReply("Source_is_incorrect_or_empty",nil, res)
			return
		case 33000:
			result.JsonReply("System_error",nil, res)
			return
		default :
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful", nil, res)

	/*
		sso, flag := utils.GetContext(req)
		if !utils.REGEXP_MOBILE.MatchString(ssoR.Username) {
			ssoR := rpcClient.SsoRpc(sso, "FindPasswordByMail")
			if ssoR.ErrorCode != 10000 {
				result.ResCode(20015, res)
				return
			}
			result.ResCode(20015, res)
			return
		}

		log.Info("UserPassword-sso:",sso, "flag:",flag)
		if !flag || sso == nil {
			result.JsonReply("Unknown_error", nil, res)
			return
		}
		sso.Username = userPwd.Username
		sso.Code = userPwd.MCode
		sso.CodeType = 2
		sso.Password = userPwd.Password
	update:
	//调用ssoRpc-CheckCode
		funcs := "CheckCode"
		if sso.CodeType == 3 {
			funcs = "UpdatePasswordByName"
		}
		ssoR := rpc.SsoRpc(sso, funcs)
		if ssoR.ErrorCode != 10000 {
			result.JsonReply("Unknown_error", nil, res)
			return
		}
		if sso.CodeType == 2 {
			sso.CodeType = 3
			goto update
		}
		result.JsonReply("Successful", nil, res)
	*/
}

//投诉与建议接口
func Feedback(res http.ResponseWriter,req *http.Request, param httprouter.Params) {
	log.Info("user.go/Feedback:  began ...")
	//从HTTP中获取文件信息
		err := req.ParseMultipartForm(32 << 20)
		if err != nil {
		log.Info("获取HTTP信息失败",err)
		result.JsonReply("System_error", nil, res)
		return
	}
	types := req.FormValue("types")
	description := req.FormValue("description")
	contact := req.FormValue("contact")
	file, fileHeader, fileErr := req.FormFile("file")
	log.Info("Feedback-fileHeader:",fileHeader.Filename,",fileErr:",fileErr,",types:",types,",description:",description,",contact:",contact)
	if fileErr != nil {
		log.Info("user.go/Feedback:  从HTTP中获取字节流失败")
		result.JsonReply("System_error", nil, res)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			file.Close()
			log.Info("user.go/Feedback:  file end fail")
			result.JsonReply("System_error", nil, res)
			return
		}
	}()
	//获取文件名
	fileName := fileHeader.Filename
	if fileName == "" {
		log.Info("文件名为空")
		result.JsonReply("System_error", nil, res)
		return
	}
	//判断path是否存在，不存在则创建
	path := "E:/project/go"+"/"+types+"/"
	isExists := utils.ExistsPath(path)
	if isExists {
		log.Info("user.go/Feedback:  路径不存在并且创建失败")
		result.JsonReply("System_error", nil, res)
		return
	}
	//再保存位置创建文件
	f, err := os.OpenFile(path+utils.TimeToStr_()+"_"+fileHeader.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		log.Info("user.go/Feedback:  保存文件到路径失败")
		return
	}
	defer f.Close()
	io.Copy(f, file)
	//返回JSON数据
	result.JsonReply("Successful", nil, res)
}