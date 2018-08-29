package controller

import (
	log "github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"petfone-http/pb/setting"
	adver "petfone-http/pb/adv"
	"petfone-http/pb"
	"petfone-http/po"
	"petfone-http/result"
	"petfone-http/rpc"
	"petfone-http/util"
	"strconv"
	"fmt"
	"strings"
	"petfone-http/pb/feedback"
	"petfone-http/core"
	"petfone-http/db"
	"github.com/garyburd/redigo/redis"
	"petfone-http/thirdPartyServer"
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"bufio"
)

/**
不需登录
swagger.yaml generate spec -o ./swagger.yaml
 */

//主路径
func Index(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	log.Info("Index......")
	http.Redirect(res,req,"http://127.0.0.1:7006/petfone/v1.0/sessions",http.StatusPermanentRedirect)
}

//注册接口
func Register(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	log.Info("Register......")
	var sso = new(pb.SsoRequest)
	err := util.GetHttpData(req, util.ReqMethodJson, &sso)
	if err {
		result.RESC(21002, res)
		return
	}
	log.Info("Register-ssoJson:", sso)
	if util.VerifyParamsStr(sso.Username, sso.Password, sso.Nickname) {
		result.RESC(21001, res)
		return
	}
	if util.VerifyUsername(sso.Username) {
		result.RESC(20002, res)
		return
	}
	if util.VerifyPassword(sso.Password) {
		result.RESC(20004, res)
		return
	}
	if util.VerifyNickname(sso.Nickname) {
		result.RESC(20016, res)
		return
	}
	conSso := util.GetContext(req)
	log.Info("Register-conSso:", conSso)
	if conSso == nil {
		result.RESC(10001, res)
		return
	}
	sso.Source = conSso.Source
	if util.CheckUsername(sso) {
		result.RESC(20008, res)
		return
	}
	//todo 验证手机验证码
	if util.REGEXP_MOBILE.MatchString(sso.Username) {
		sso.CodeType = 1
		ssoR := rpc.SsoRpc(sso, "CheckCode")
		if ssoR.Code != 10000 {
			result.RESC(ssoR.Code, res)
			return
		}
	}
	//todo 验证邮箱验证码-----
	//todo 调用rpc创建用户
	sso.State = 3
	sso.Salt = string(util.Krand(6, util.KC_NUMBERS_LETTERS))
	ip := strings.Split(req.RemoteAddr,":")[0]
	sso.AgentInfo = &pb.AgentInfo{Ip:ip,DevInfo:req.Header.Get("User-Agent")}
	log.Info("sso-add:", sso)
	//调用ssoRpc-add
	ssoR := rpc.SsoRpc(sso, "Add")
	if ssoR.Code != 10000 {
		result.RESC(ssoR.Code, res)
		return
	}
	//todo APP注册时添加登录
	if sso.GetSource() == "AgIDAA==" {
		CreateSession(sso, req, res)
		return
	}
	result.RESC(10000, res)

}

//登录接口
func Login(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	log.Info("Login......")
	sso := &pb.SsoRequest{}
	err := util.GetHttpData(req, util.ReqMethodJson, &sso)
	if err {
		result.RESC(21002, res)
		return
	}
	conSso := util.GetContext(req)
	if conSso == nil {
		result.RESC(10001, res)
		return
	}
	sso.Source = conSso.Source
	log.Info("Login-ssoJson:", sso)
	if util.VerifyParamsStr(sso.Username, sso.Password) {
		result.RESC(20001, res)
		return
	}
	if util.VerifyUsername(sso.Username) {
		result.RESC(20002, res)
		return
	}
	if util.VerifyPassword(sso.Password) {
		result.RESC(20004, res)
		return
	}
	//调用ssoRpc-login
	CreateSession(sso, req, res)
}

//创建token
func CreateSession(sso *pb.SsoRequest, req *http.Request, res http.ResponseWriter)  {
	ssoR := rpc.SsoRpc(sso, "Login")
	if ssoR.Code != 10000 {
		result.RESC(ssoR.Code, res)
		return
	}
	login := new(po.LoginPo)
	util.LoginTssoR(login, ssoR)
	//todo 获取用户基本信息
	account := &pb.AccountRequest{Source:sso.Source, Uid:ssoR.Uid}
	accountR := rpc.AccountRpc(account, "GetAccountInfo")
	if accountR.Code == 10000 {
		util.LoginTaccountR(login, accountR)
	}
	//todo 返回JSON数据
	req.Header.Add("uid", util.Int32ToStr(ssoR.Uid))
	req.Header.Add("token", ssoR.Token)
	result.REST(login, res)
}

//手机重置密码接口
func ResetPassword(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	log.Info("ResetPassword......")
	sso := util.GetContext(req)
	log.Info("UserPassword-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	err := util.GetHttpData(req, util.ReqMethodJson, &sso)
	if err {
		result.RESC(21002, res)
		return
	}
	log.Info("ResetPassword-ssoJson:", sso)
	if util.VerifyParamsStr(sso.Username) {
		result.RESC(20015, res)
		return
	}
	if !util.REGEXP_MOBILE.MatchString(sso.Username) {
		result.RESC(20002, res)
		return
	}
	if util.VerifyParamsStr(sso.Password) {
		result.RESC(20004, res)
		return
	}
	if util.VerifyParamsStr(sso.Code) {
		result.RESC(20012, res)
		return
	}
	if util.VerifyPassword(sso.Password) {
		result.RESC(21004, res)
		return
	}
	sso.CodeType = 2
update:
//调用ssoRpc-CheckCode
	funcs := "CheckCode"
	if sso.CodeType == 3 {
		funcs = "UpdatePasswordByName"
		sso.Salt = string(util.Krand(6, util.KC_NUMBERS_LETTERS))
	}
	ssoR := rpc.SsoRpc(sso, funcs)
	if ssoR.Code != 10000 {
		result.RESC(ssoR.Code, res)
		return
	}
	if sso.CodeType == 2 {
		sso.CodeType = 3
		goto update
	}
	result.RESC(10000, res)
}

//邮件重置密码接口
func MailResetPassword(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("MailResetPassword......")
	action := params.ByName("action")
	var sso = new(pb.SsoRequest)
	err := util.GetHttpData(req, util.ReqMethodJson, &sso)
	if err {
		result.RESC(21002, res)
		return
	}
	ssoCon := util.GetContext(req)
	log.Info("MailResetPassword-sso:", sso)
	if ssoCon == nil {
		result.RESC(10001, res)
		return
	}
	sso.Source = ssoCon.GetSource()
	if "send" == action {
		log.Info("ResetPassword-FindPasswordByMail:", sso)
		if util.VerifyParamsStr(sso.GetUsername()) {
			result.RESC(21001, res)
			return
		}
		if !util.REGEXP_MAIL.MatchString(sso.Username) {
			result.RESC(20002, res)
			return
		}
		ssoR := rpc.SsoRpc(sso, "FindPasswordByMail")
		if ssoR.Code != 10000 {
			result.RESC(ssoR.Code, res)
			return
		}
		result.RESC(10000, res)
		return
	}
	sso.Username = req.FormValue("username")
	sso.Token = req.FormValue("token")
	log.Info("MailResetPassword-ssoJson:", sso)
	if util.VerifyParamsStr(sso.GetUsername(), sso.GetToken()) || util.VerifyParamsStr(sso.GetPassword()) {
		result.RESC(21001, res)
		return
	}
	if !util.REGEXP_MAIL.MatchString(sso.Username) {
		result.RESC(20002, res)
		return
	}
	if util.VerifyPassword(sso.GetPassword()) {
		result.RESC(20004, res)
		return
	}
	sso.Salt = string(util.Krand(6, util.KC_NUMBERS_LETTERS))
	log.Info("MailResetPassword-ResetPassword:", sso)
	ssoR := rpc.SsoRpc(sso, "ResetPassword")
	if ssoR.Code != 10000 {
		result.RESC(ssoR.Code, res)
		return
	}
	result.RESC(10000, res)
}

//手机验证码接口
func MobileCode(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	log.Info("MobileCode......")
	action := param.ByName("action")
	sso := &pb.SsoRequest{}
	nofund := util.GetHttpData(req, util.ReqMethodJson, &sso)
	if nofund {
		result.RESC(21002, res)
		return
	}
	log.Info("SendCode-ssoJson:", sso,",action:",action)
	if util.VerifyParamsStr(sso.Username) {
		result.RESC(20002, res)
		return
	}
	if sso.CodeType != 1 && sso.CodeType != 2 {
		result.RESC(21002, res)
		return
	}
	if !util.REGEXP_MOBILE.MatchString(sso.Username) {
		result.RESC(20002, res)
		return
	}
	conSso := util.GetContext(req)
	if conSso == nil {
		result.RESC(10001, res)
		return
	}
	sso.Source = conSso.Source
	if "send" == action {
		ssoR := rpc.SsoRpc(sso,"SendMobileCode")
		if ssoR.Code != 33004 && ssoR.Code != 10000 {
			result.RESC(ssoR.Code, res)
			return
		}
		result.RESC(10000, res)
		return
	}
	if "check" == action {
		if util.VerifyParamsStr(sso.Code) {
			result.RESC(20012, res)
			return
		}
		ssoR := rpc.SsoRpc(sso, "CheckCode")
		if ssoR.Code != 10000 {
			result.RESC(ssoR.Code, res)
			return
		}
		result.RESC(10000, res)
		return
	}
	result.ResCode(404, res)
}

//意见反馈
func SetFeedback(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	log.Info("SetFeedback......")
	feedBack := &feedback.AddFeedbackRequest{}
	isOk := util.GetHttpData(req, util.ReqMethodJson, &feedBack)
	if isOk {
		result.RESC(21002, res)
		return
	}
	log.Info("SetFeedback-feedBack:", feedBack)
	if util.VerifyParamsStr(feedBack.Description,feedBack.MobileInfo,feedBack.AppInfo) {
		result.RESC(21001, res)
		return
	}
	sso := util.GetContext(req)
	log.Info("SetFeedback-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	feedBack.Source = sso.Source
	//todo 调用RPC
	feedBackRe := rpc.FeedBackRpc(feedBack,"AddFeedback")
	if feedBackRe.GetErrorCode() != 10000 {
		result.RESC(feedBackRe.ErrorCode, res)
		return
	}
	go func() {
		//todo 读取邮箱地址
		client := core.RedisClient(6379)
		defer client.Close()
		emailAddr, err := redis.String(client.Do("get", db.FeedBackMail))
		if err != nil || emailAddr == "" {
			log.Info("读取意见反馈邮箱地址失败：",err)
		}
		var fileStr string
		for _,v := range feedBack.Files {
			fileStr += ` `+v+` `
		}
		content :=util.StrAdd(
			`<pre>联系方式：`,	util.PrintStr(feedBack.Contact),	`</pre>`,
			`<pre>描述信息：`,	util.PrintStr(feedBack.Description),`</pre>`,
			`<pre>手机信息：`,	util.PrintStr(feedBack.MobileInfo),	`</pre>`,
			`<pre>应用信息：`,	util.PrintStr(feedBack.AppInfo),	`</pre>`,
			`<pre>设备信息：`,	util.PrintStr(feedBack.DeviceInfo),	`</pre>`,
			`<pre>用户信息：`,	util.PrintStr(feedBack.UserInfo),	`</pre>`,
			`<pre>文件信息：`,	util.PrintStr(fileStr),				`</pre>`,
			`<pre>扩展信息：`,	util.PrintStr(feedBack.ExtendInfo),	`</pre>`, )
		thirdPartyServer.FeedbackMail(emailAddr,util.Int32ToStr(feedBackRe.Id),content)
	}()
	result.RESC(10000, res)
}

//版本更新
func GetVersion(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	log.Info("GetVersion......")
	name := param.ByName("name")
	codeStr := param.ByName("code")
	log.Info("name:", name, ",codeStr:", codeStr)
	if util.VerifyParamsStr(name,codeStr) {
		result.RESC(21002, res)
		return
	}
	sso := util.GetContext(req)
	log.Info("GetVersion-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	version := &setting.GetLatestVersionRequest{Source: sso.Source, Device:name}
	//调用id-rpc
	log.Info("GetVersion-version:", version)
	versionRe := rpc.VersionRpc(version, "GetLatestVersion")
	if versionRe.GetErrorCode() != 10000 {
		result.RESC(versionRe.ErrorCode, res)
		return
	}
	result.REST(versionRe, res)
}
func GetVersion1(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	log.Info("GetVersion1......")
	name := param.ByName("name")
	didStr := req.FormValue("did")
	log.Info("GetVersion1 name:", name, ",didStr:", didStr)
	if util.VerifyParamsStr(name,didStr) {
		result.RESC(21002, res)
		return
	}
	sso := util.GetContext(req)
	log.Info("GetVersion1-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	//did, err := util.StrToInt32(didStr)
	//if err != nil {
	//	result.RESC(21001, res)
	//	return
	//}
	//device := &pb.DeviceRequest{Source:sso.Source,Uid:sso.Uid,Did:did}
	//deviceRe := rpc.DevicesRpc(device,"GetDevicesByDid")
	//todo 调用版本信息
	version := &setting.GetLatestVersionRequest{Source: sso.Source, Device:name, Username:sso.Username}
	log.Info("GetVersion1-version:", version)
	versionRe := rpc.VersionRpc(version, "GetLatestVersion")
	if versionRe.GetErrorCode() != 10000 {
		result.RESC(versionRe.ErrorCode, res)
		return
	}
	if versionRe.Version == nil {
		result.RESC(33013, res)
		return
	}
	result.REST(versionRe, res)
}

//获取广告
func GetAdver(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	log.Info("GetAdver......")
	sso := util.GetContext(req)
	log.Info("GetAdver-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	adve := &adver.AdvertisementRequest{Source: sso.Source}
	//调用rpc
	log.Info("GetAdver-adver:", adve)
	adverRe := rpc.AdvertisementRpc(adve, "GetAdvertisement")
	if adverRe.GetErrorCode() != 10000 {
		result.RESC(adverRe.ErrorCode, res)
		return
	}
	result.REST(adverRe, res)
}

func Test(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	body, err := ioutil.ReadAll(bufio.NewReader(req.Body))
	if err != nil {
		fmt.Println("err:",err)
		res.Write([]byte("test"))
		return
	}
	fmt.Println("len:",len(body))
	buf := new(bytes.Buffer)
	gzipw := gzip.NewWriter(buf)
	leng, err := gzipw.Write(body)
	if err != nil || leng == 0 {
		return
	}
	err = gzipw.Flush()
	if err != nil {
		return
	}
	err = gzipw.Close()
	if err != nil {
		return
	}
	res.Write(buf.Bytes())
	//io.Copy(res,req.Body)
	fmt.Println("aaaaaaaaaa")
}

/**
小助手
 */

//常见问题
func FaqCommonInfo(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	log.Info("FaqCommonInfo......")
	idStr := req.FormValue("id")
	name := req.FormValue("name")
	sso := util.GetContext(req)
	log.Info("FaqCommonInfo-sso:", sso, ",idStr:", idStr)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	faqCommonPo := &pb.FaqCommonRequest{Source: sso.Source, Uid: sso.Uid}
	log.Info("FaqCommonInfo-faqCommonPo:", faqCommonPo)
	if util.VerifyParamsStr(idStr) {
		var faqCRes *pb.FaqCommonsReply
		//调用rpc
		if util.VerifyParamsStr(name) {
			faqCRes = rpc.FaqCommonsRpc(faqCommonPo, "GetFaqCommons")
		} else {
			faqCommonPo.NameCn = name
			faqCRes = rpc.FaqCommonsRpc(faqCommonPo, "GetFaqCommonByKeyword")
		}
		log.Info("FaqCommonInfo-faqCRes:", faqCRes)
		if faqCRes.Code != 10000 {
			result.RESC(faqCRes.Code, res)
			return
		}
		//var FaqCommonPos []*po.FaqCommonPo
		//for _, v := range faqCRes.Faqcs {
		//	FaqCommonPos = append(FaqCommonPos, &po.FaqCommonPo{Id: v.Id, Name: v.NameCn, Parent: v.Parent, Info: v.InfoCn})
		//}
		result.REST(faqCRes.Faqcs, res)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		result.RESC(21002, res)
		return
	}
	faqCommonPo.Id = int32(id)
	//调用id-rpc
	log.Info("FaqCommonInfo-faqCommonPo:", faqCommonPo)
	faqRe := rpc.FaqCommonRpc(faqCommonPo, "GetFaqCommonById")
	if faqRe.Code != 10000 {
		result.RESC(faqRe.Code, res)
		return
	}
	result.REST(faqRe, res)
}

//问题例子
func PetChatFaqCommonInfo(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	log.Info("FaqCommonInfo......")
	sso := util.GetContext(req)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	petChatRequest := &pb.PetChatRequest{Source: sso.Source, Uid: sso.Uid}
	log.Info("FaqCommonInfo-petChatRequest:", petChatRequest)
	petChatFaqRes := rpc.PetChatFaqRpc(petChatRequest, "GetPetChatKey")
	log.Info("FaqCommonInfo-petChatFaqRes:", len(petChatFaqRes.ChatKeys))
	if petChatFaqRes.Code != 10000 {
		result.RESC(petChatFaqRes.Code, res)
		return
	}
	result.REST(petChatFaqRes.ChatKeys, res)
}


/**
需登录
 */

//获取用户信息接口
func GetUserInfo(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	log.Info("GetUserInfo......")
	var sso = new(pb.SsoRequest)
	conSso := util.GetContext(req)
	if conSso == nil {
		result.RESC(10001, res)
		return
	}
	sso.Source = conSso.Source
	sso.SessionName = conSso.GetSessionName()
	//调用ssoRpc-GetUserInfo
	log.Info("GetUserInfo-ssoJson:", sso)
	ssoRe := rpc.SsoRpc(sso, "GetUserInfo")
	if ssoRe.Code != 10000 {
		result.RESC(ssoRe.Code, res)
		return
	}
	login := new(po.LoginPo)
	util.LoginTssoR(login, ssoRe)
	//调用accountRpc-GetUserInfoAll
	account := new(pb.AccountRequest)
	account.Source = conSso.Source
	account.Uid = conSso.Uid
	account.Token = conSso.Token
	accountRe := rpc.AccountRpc(account, "GetAccountInfo")
	if accountRe.Code != 10000 {
		result.RESC(accountRe.Code, res)
		return
	}
	login.Nickname = accountRe.Nickname
	util.LoginTaccountR(login, accountRe)
	//返回JSON数据
	result.REST(login, res)
}

//用户名查询接口
func GetAccountByUserName(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	log.Info("GetAccountByUserName......")
	usernameStr := param.ByName("param")
	log.Info("usernameStr:", usernameStr)
	if util.VerifyParamsStr(usernameStr) {
		result.RESC(20001, res)
		return
	}
	if util.VerifyUsername(usernameStr) {
		result.RESC(20002, res)
		return
	}
	sso := util.GetContext(req)
	sso.Username = usernameStr
	ssoRe := rpc.SsoRpc(sso, "GetUserByName")
	if ssoRe.Code != 10000 {
		result.RESC(ssoRe.Code, res)
		return
	}
	friendPo := &po.AccountFriendPo{Username: usernameStr, Nickname: ssoRe.Nickname}
	//todo 获取用户基本信息
	account := new(pb.AccountRequest)
	account.Source = sso.Source
	account.Uid = ssoRe.Uid
	accountRe := rpc.AccountRpc(account, "GetAccountInfo")
	if accountRe.Code != 10000 {
		result.RESC(accountRe.Code, res)
		return
	}
	friendPo.Uid = accountRe.Uid
	friendPo.Phone = accountRe.Phone
	friendPo.Email = accountRe.Email
	friendPo.Avatar = accountRe.Avatar
	friendPo.Gender = accountRe.Gender
	//todo 返回JSON数据
	result.REST(friendPo, res)
}

//用户修改资料接口
func UpdateUserInfo(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	log.Info("UpdateUserInfo......")
	var account = &pb.AccountRequest{}
	if util.GetHttpData(req, util.ReqMethodJson, &account) {
		result.RESC(21002, res)
		return
	}
	log.Info("UpdateUserInfo-account:", account)
	if len(account.Nickname) > 0 {
		if util.VerifyNickname(account.Nickname) {
			result.RESC(21002, res)
			return
		}
	}
	if len(account.Signature) > 0 {
		if util.VerifyLenParams(account.Signature) {
			result.RESC(21002, res)
			return
		}
	}
	if len(account.Address) > 250 {
		result.RESC(21002, res)
		return
	}
	if len(account.Avatar) > 250 {
		result.RESC(21002, res)
		return
	}
	if len(account.Email) > 0 {
		if !util.REGEXP_MAIL.MatchString(account.Email) {
			result.RESC(21002, res)
			return
		}
	}
	if len(account.Phone) > 0 {
		if !util.REGEXP_MOBILE.MatchString(account.Phone) {
			result.RESC(21002, res)
			return
		}
	}
	if util.VerifyGender(account.Gender) {
		result.RESC(21002, res)
		return
	}
	conSso := util.GetContext(req)
	log.Info("UpdateUserInfo-conSso:", conSso)
	if conSso == nil {
		result.RESC(10001, res)
		return
	}
	account.Uid = conSso.Uid
	account.Source = conSso.Source
	account.Token = conSso.Token
	accountR := rpc.AccountRpc(account, "UpdateAccountInfo")
	if accountR.Code != 10000 {
		log.Info("UpdateUserInfo-AccountRpc Error:", accountR.Code)
		result.RESC(accountR.Code, res)
		return
	}
	//返回JSON数据
	result.REST(accountR, res)
}

//用户修改密码接口
func UserPassword(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	action := param.ByName("action")
	log.Info("UserPassword...... param:", action)
	var userPwd = new(po.UserPwd)
	err := util.GetHttpData(req, util.ReqMethodJson, &userPwd)
	if err {
		result.RESC(21002, res)
		return
	}
	log.Info("UserPassword-updatepwdJson:", userPwd)
	sso := util.GetContext(req)
	log.Info("UserPassword-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	if "check" == action || "update" == action {
		if util.VerifyParamsStr(userPwd.Password) {
			result.RESC(20004, res)
			return
		}
		if len(userPwd.Password) < 32 {
			result.RESC(20004, res)
			return
		}
		sso.Password = userPwd.Password
		//调用ssoRpc-CheckPassword
		ssoR := rpc.SsoRpc(sso, "CheckPassword")
		if ssoR.Code != 10000 {
			result.RESC(ssoR.Code, res)
			return
		}
		if "update" == action {
			goto update
		}
		result.RESC(10000, res)
		return
	}
update:
	if "update" == action {
		if util.VerifyParamsStr(userPwd.NewPassword) {
			result.RESC(20015, res)
			return
		}
		sso.Password = userPwd.NewPassword
		sso.Salt = string(util.Krand(6, util.KC_NUMBERS_LETTERS))
		//调用ssoRpc-UpdatePassword
		ssoR := rpc.SsoRpc(sso, "UpdatePassword")
		if ssoR.Code != 10000 {
			result.RESC(ssoR.Code, res)
			return
		}
		result.RESC(10000, res)
		return
	}
	result.ResCode(404, res)
}

//退出接口
func Logout(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
	log.Info("Logout......")
	sso := util.GetContext(req)
	log.Info("Logout-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	ssoR := rpc.SsoRpc(sso, "Logout")
	if ssoR.Code != 10000 {
		result.RESC(ssoR.Code, res)
		return
	}
	result.RESC(10000, res)
}

/**
用户业务
 */

//添加围栏 - 停用
func SetPetfoneFence(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("SetPetfoneFence......")
	result.RESC(10003, res)
}

//修改围栏
func UpdatePetfoneFence(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("UpdatePetfoneFence......")
	sso := util.GetContext(req)
	log.Info("UpdatePetfoneFence-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	petfoneReq := new(pb.PetfoneRequest)
	nofund := util.GetHttpData(req, util.ReqMethodJson, &petfoneReq)
	if nofund {
		result.RESC(21002, res)
		return
	}
	log.Info("UpdatePetfoneFence-petfoneReqJson:", petfoneReq)
	if util.VerifyParamsUInt32(1) {
		result.RESC(20015, res)
		return
	}
	petfoneReq.Source = sso.GetSource()
	petfoneReq.Uid = sso.GetUid()
	//调用rpc
	petfoneRep := rpc.PetfoneRpc(petfoneReq, "UpdatePetfoneByUid")
	if petfoneRep.Code != 10000 {
		result.RESC(petfoneRep.Code, res)
		return
	}
	result.RESC(10000, res)
}

//获取围栏
func GetPetfoneFence(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("GetPetfoneFence......")
	sso := util.GetContext(req)
	log.Info("GetPetfoneFence-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	petfoneReq := new(pb.PetfoneRequest)
	petfoneReq.Source = sso.GetSource()
	petfoneReq.Uid = sso.GetUid()
	//调用rpc
	petfoneRep := rpc.PetfoneRpc(petfoneReq, "GetPetfoneByUid")
	if petfoneRep.Code != 10000 {
		result.RESC(petfoneRep.Code, res)
		return
	}
	result.REST(petfoneRep, res)
}

//添加共享
func SetShare(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("SetShare......")
	sso := util.GetContext(req)
	log.Info("SetShare-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	shareReq := &pb.ShareRequest{}
	nofund := util.GetHttpData(req, util.ReqMethodJson, &shareReq)
	if nofund {
		result.RESC(21002, res)
		return
	}
	log.Info("shareReq:", shareReq)
	if len(shareReq.Pids) == 0 || shareReq.MemberUid == 0 || shareReq.OwnerUid != sso.Uid {
		result.RESC(21002, res)
		return
	}
	shareReq.Source = sso.Source
	shareReq.OwnerUid = sso.Uid
	//调用rpc
	shareRe := rpc.ShareRpc(shareReq, "SetShare")
	result.RESC(shareRe.Code, res)
}

//删除共享
func DeleteShare(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("DeleteShare......")
	sso := util.GetContext(req)
	log.Info("DeleteShare-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	shareReq := &pb.ShareRequest{}
	nofund := util.GetHttpData(req, util.ReqMethodJson, &shareReq)
	if nofund {
		result.RESC(21002, res)
		return
	}
	log.Info("shareReq:", shareReq)
	if len(shareReq.Pids) == 0 || shareReq.OwnerUid == 0 || shareReq.MemberUid == 0 {
		result.RESC(21001, res)
	}
	if shareReq.OwnerUid != sso.Uid && shareReq.MemberUid != sso.Uid {
		result.RESC(21002, res)
	}
	shareReq.Source = sso.Source
	//调用rpc
	shareRe := rpc.ShareRpc(shareReq, "DeleteShare")
	result.RESC(shareRe.Code, res)
}

//获取共享
func GetShare(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("GetShare......")
	sso := util.GetContext(req)
	log.Info("GetShare-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	shareReq := &pb.ShareRequest{Source: sso.GetSource(), OwnerUid: sso.GetUid()}
	log.Info("shareReq:", shareReq)
	//调用rpc
	shareRe := rpc.SharesRpc(shareReq, "GetShare")
	if shareRe.Code != 10000 {
		result.RESC(shareRe.Code, res)
		return
	}
	var sharePos []*pb.ShareReply
	for _, vs := range shareRe.Shares {
		if len(vs.Members) == 0 {
			continue
		}
		if vs.OwnerInfo.Uid == sso.Uid {
			vs.OwnerInfo = &pb.AccountReply{}
		} else {
			for _, vm := range vs.Members {
				vm.MemberInfo = &pb.AccountReply{}
				//	if sso.Uid != vm.MemberInfo.Uid {
				//		vs.Members = append(vs.Members[:k], vs.Members[k+1:]...)
				//	}
			}
		}
		sharePos = append(sharePos, vs)
	}
	result.REST(sharePos, res)
}

//------------    分割线   --------------------------------

func testfile(res http.ResponseWriter, req *http.Request, param httprouter.Params)  {
	//其实这里的 O_RDWR应该是 O_RDWR|O_CREATE，也就是文件不存在的情况下就建一个空文件，
	// 但是因为windows下还有BUG，如果使用这个O_CREATE，就会直接清空文件，
	// 所以这里就不用了这个标志，你自己事先建立好文件。
	file, err := os.Open("E:/共享智能空间.pptx")
	if err != nil {
		fmt.Println(err)
		log.Info("读取文件到路径失败")
		return
	}
	defer file.Close()
	fileName := path.Base("E:/共享智能空间.pptx")
	log.Info("--", fileName)
	//防止中文乱码
	fileName = url.QueryEscape(fileName)
	log.Info("--", fileName)
	res.Header().Add("Content-Type", "application/octet-stream")
	res.Header().Add("content-disposition", "attachment; filename=\""+fileName+"\"")
	io.Copy(res, file)
	//http.ServeFile(res, req, "E:/aaa.txt")
	//result.RESC(10000, res)
}

//文件大小
func cat(f *os.File) []byte {
	var payload []byte
	fatat, _ := f.Stat()
	log.Info("fatat:", fatat.Size())
	for {
		buf := make([]byte, 1024)
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			log.Infof("%s ; cat: error reading: %s\n", os.Stderr, err.Error())
			os.Exit(1)
		case nr == 0: // EOF
			return payload
		case nr > 0:
			payload = append(payload, buf...)
		}
	}

}
