package utils

import (
	log "github.com/cihub/seelog"
	"cotx-http/pb"
	"time"
	"math/rand"
	"cotx-http/rpcClient"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"bufio"
	"os"
	"golang.org/x/net/context"
	"github.com/pkg/errors"
	"fmt"
)
//处理HTTP请求中的数据
func GetHttpData(req *http.Request, types string,t interface{}) int32 {
	bodyType := req.Header.Get("Content-Type")
	log.Info("GetHttp-Content-Type:",bodyType)
	//index := strings.Contains(types, bodyType)
	if types != bodyType {
		return 404
	}
	//获取body
	body, bodyErr := ioutil.ReadAll(bufio.NewReader(req.Body))
	if bodyErr != nil {
		log.Error("GetHttp-ReadAll-bodyErr:",bodyErr)
		return 21002
	}
	bodyStr := string(body)
	//bodyStr = strings.Replace(strings.Replace(bodyStr," ","",-1),"\n","",-1)
	log.Info("GetHttp-bodyStr:",bodyStr)
	if len(body) < 8 {
		return 21001
	}
	if ssoJsonerr := json.Unmarshal(body, &t); ssoJsonerr !=nil {
		log.Error("GetHttp-Unmarshal-ssoJsonerr:",ssoJsonerr)
		return 21002
	}
	return 10000
}

func GetHttpDataGwSet(req *http.Request, types string,t *pb.ReqInstruction) int32 {
	bodyType := req.Header.Get("Content-Type")
	log.Info("GetHttp-Content-Type:",bodyType)
	//index := strings.Contains(types, bodyType)
	if types != bodyType {
		return 404
	}
	//获取body
	body, bodyErr := ioutil.ReadAll(bufio.NewReader(req.Body))
	if bodyErr != nil {
		log.Error("GetHttp-ReadAll-bodyErr:",bodyErr)
		return 20001
	}
	bodyStr := string(body)
	//bodyStr = strings.Replace(strings.Replace(bodyStr," ","",-1),"\n","",-1)
	log.Info("GetHttp-bodyStr:",bodyStr)
	if len(body) < 8 {
		return 20001
	}
	if ssoJsonerr := json.Unmarshal(body, &t); ssoJsonerr !=nil {
		log.Error("GetHttp-Unmarshal-ssoJsonerr:",ssoJsonerr)
		return 20001
	}
	pbType,err := ValidationInstrustion(t.Instruction,body,t.GatewayId,t.UserId)
	log.Info("rpcRes ==:",pbType)
	if err!=nil {
		return 30022
	}
	return pbType
}
func UnMarshalJsonBody(body []byte,t interface{}) int32 {
	if ssoJsonerr := json.Unmarshal(body, &t); ssoJsonerr !=nil {
		log.Error("GetHttp-Unmarshal-ssoJsonerr:",ssoJsonerr)
		return 21002
	}
	return 10000
}
// 对参数进行int32判断与转换
func VerifyParamsUInt32(params ...int32) bool {
	lens := len(params)
	for i := 0; i < lens; i++{
		if params[i] < 0 {
			return true
		}
	}
	return false
}

// 对参数进行int64判断与转换
func VerifyParamsUInt64(params ...int64) bool {
	lens := len(params)
	for i := 0; i < lens; i++{
		if params[i] < 0 {
			return true
		}
	}
	return false
}

//HTTP获取上下文
func GetContext(req *http.Request) (*pb.SsoRequest, bool) {
	userInfo, ok := req.Context().Value("userInfo").(*pb.SsoRequest)
	return userInfo, ok
}

// 对参数进行非空判断
func VerifyParamsEmpty(params ...string) bool {
	for i := 0; i < len(params); i++{
		if params[i] == "" {
			return true
		}
	}
	return false
}

//对用户名进行验证
func VerifyUsername(username string) bool {
	if username == "" {
		return false
	}
	if REGEXP_MAIL.MatchString(username) {
		return false
	}
	if REGEXP_MOBILE.MatchString(username) {
		return false
	}
	return true
}

//对密码进行验证
func VerifyPassword(password string) bool {
	if password == "" {
		return false
	}
	if REGEXP_PWD.MatchString(password) {
		return false
	}
	return true
}

//对昵称进行正则验证
func VerifyNickname(nickname string) bool {
	if nickname == "" {
		return false
	}
	if REGEXP_NICKNAME.MatchString(nickname) {
		return false
	}
	return true
}

//生成随机字节
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

//HTTP接口地址
const (
	//项目路径
	serverName		= "/petfone"
	version_1		= "/v1.0"
	category_user	= "/user"
	category_pwd	= "/pwd"
	//接口路径
	SendA			= serverName + version_1 + "/send"
	RegA			= serverName + version_1 + "/reg"
	LoginA			= serverName + version_1 + "/login"
	LogoutA			= serverName + version_1 + "/logout"
	UserinfoA		= serverName + version_1 + "/userinfo"
	UpdatePwdA		= serverName + version_1 + category_pwd + "/change/:action"
	ResetPwdA		= serverName + version_1 + category_pwd + "/reset/:action"
	FeedbackA		= serverName + version_1 + "/feedback"
)

//半公开接口地址
var halfOpenAddress = [...]string{
	//发送验证码
	"/cotx/v1.0/code/send",
	"/cotx/v1.0/code/check",

	"/cotx/v1.0/register",

	"/cotx/v1.0/sendcode",
	"/cotx/v1.0/session",
	"/cotx/v1.0/reset/pwd/:uid",
	"/cotx/v1.0/reset/mobile",
	//邮箱校验
	"/cotx/v1.0/mailbox/send",
	"/cotx/v1.0/mailbox/check",
	//判断用户名是否存在
	"/cotx/v1.0/user/judge",
}

//对半公开接口判断
func VerifyHalfOpenAddress(url string, method string) bool {
	log.Info("VerifyHalfOpenAddress-url:",url)
	for _, v := range halfOpenAddress {
		if url == "/cotx/v1.0/session" && method != "POST" {
			return false
		}
		if url == v  {
			return true
		}
	}
	return false

}

//对sso/ssoR进行属性处理
func SsoTssoR(sso *pb.SsoRequest, ssoR *pb.SsoReply)  {
	sso.Uid			= ssoR.Uid
	sso.Username	= ssoR.Username
	sso.Nickname	= ssoR.Nickname
	sso.Token		= ssoR.SessionName
	sso.State		= ssoR.State
}

//时间戳格式化为字符串
func TimeToStr_() string {
	return time.Now().Format("2006-01-02-15-04-05")
}

//时间戳格式化为字符串
func TimeToStr() string {
	return time.Now().Format("2006年01月02 15:04:05")
}

//判断指定的路径是否存在
func ExistsPath(path string) bool {
	_,err := os.Stat(path)
	if err == nil {
		return false
	}
	if os.IsExist(err) {
		return true
	}
	mkErr := os.Mkdir(path, os.ModePerm)
	if mkErr != nil {
		return true
	}
	return false
}

//验证用户名是否存在   true表示用户存在,false表示用户不存在
func CheckUsername(sso *pb.SsoRequest) (bool, error) {
	log.Info("CheckUsername:",sso)
	ssoRC, ssoRErr := rpcClient.GetSsoRpcClient().GetUserInfo(context.Background(),sso)
	if ssoRErr != nil {
		log.Error("Register-CheckUserRPC Error",ssoRErr)
		return true,ssoRErr
	}
	log.Info("Register-CheckUserRPC Warn:", ssoRC.ErrorCode)
	if ssoRC.ErrorCode == 10000 {
		return true,nil
	}
	if ssoRC.ErrorCode != 33002 {
		return true,nil
	}
	return false,nil
}
func ValidationInstrustion (ins int32,body []byte,gwid string ,uid int32)(int32,error){
	switch ins {
	case GatewayUpdata,GatewayReset,DisConnect_Net,CloseGateway ,RevertBackUp,RevertGateway,BackUp,deleteBackUpFile,UpdateGatewayING:
		rins :=&pb.ReqInstruction{GatewayId:gwid,UserId:uid}
		errCode := UnMarshalJsonBody(body,rins)
		fmt.Println("errCode=:",errCode)
		if errCode != 10000 {
			return errCode,nil
		}
		//rpc调用
		res ,err := rpcClient.SendInstruction(rins)
		if err != nil {
			return 30033,nil
		}
		return res.ErrCode,nil
	case setBleScan,WifiCard,WifiCardScan,NBIOTSet,NFCSet,TakeMusic,SetSSH,HotSpot,PowerSwitch,UnConnectWifi,setGatewayNetWarn,setGatewaySysWarn:
		rcw :=&pb.ReqSwitch{GatewayId:gwid,UserId:uid}
		errCode := UnMarshalJsonBody(body,rcw)
		fmt.Println("errCode=:",errCode)
		if errCode != 10000 {
			return errCode,nil
		}
		//rpc调用
		resss ,err := rpcClient.SendSwitch(rcw)
		if err != nil {
			return 30033,nil
		}
		return resss.ErrCode,nil
	case PowerModel:
		rgp :=&pb.ReqGwPower{GatewayId:gwid,UserId:uid}
		errCode := UnMarshalJsonBody(body,rgp)
		fmt.Println("errCode=:",errCode)
		if errCode != 10000 {
			return errCode,nil
		}
		//rpc调用
		resssp ,err:= rpcClient.SendSetPower(rgp)
		if err != nil {
			return 30033,nil
		}

		return resssp.ErrCode ,nil
	case LoraSet,Lora1301:
		rl :=&pb.ReqLora{GatewayId:gwid,UserId:uid}
		errCode := UnMarshalJsonBody(body,rl)
		fmt.Println("errCode=:",errCode)
		if errCode != 10000 {
			return errCode,nil
		}
		//rpc调用
		resrl,err:= rpcClient.SendLora(rl)
		if err != nil {
			return 30033,nil
		}
		return resrl.ErrCode,nil
	case AutoVideo:
		rv :=&pb.ReqVideo{GatewayId:gwid,UserId:uid}
		errCode := UnMarshalJsonBody(body,rv)
		fmt.Println("errCode=:",errCode)
		if errCode != 10000 {
			return errCode,nil
		}
		//rpc调用
		ressv ,err:= rpcClient.SendVideo(rv)
		if err != nil {
			return 30033,nil
		}
		return ressv.ErrCode,nil
	case AutoPhoto:
		rp :=&pb.ReqPhoto{GatewayId:gwid,UserId:uid}
		errCode := UnMarshalJsonBody(body,rp)
		fmt.Println("errCode=:",errCode)
		if errCode != 10000 {
			return errCode,nil
		}
		//rpc调用
		rsssp ,err:= rpcClient.SendPhoto(rp)
		if err != nil {
			return 30033,nil
		}
		return rsssp.ErrCode,nil
	case AutoMusic:
		rm :=&pb.ReqMusic{GatewayId:gwid,UserId:uid}
		errCode := UnMarshalJsonBody(body,rm)
		fmt.Println("errCode=:",errCode)
		if errCode != 10000 {
			return errCode,nil
		}
		//rpc调用
		ressm ,err:= rpcClient.SendMusic(rm)
		if err != nil {
			return 30033,nil
		}

		return ressm.ErrCode,nil
	case DeletFile:
		rdf :=&pb.ReqDeletFile{GatewayId:gwid,UserId:uid}
		errCode := UnMarshalJsonBody(body,rdf)
		fmt.Println("errCode=:",errCode)
		if errCode != 10000 {
			return errCode,nil
		}
		//rpc调用
		ressdf ,err:= rpcClient.SendDeletFile(rdf)
		if err != nil {
			return 30033,nil
		}
		return ressdf.ErrCode,nil
	case UpLog:
		rul :=&pb.ReqUpLog{GatewayId:gwid,UserId:uid}
		errCode := UnMarshalJsonBody(body,rul)
		fmt.Println("errCode=:",errCode)
		if errCode != 10000 {
			return errCode,nil
		}
		//rpc调用
		ressul ,err:= rpcClient.SendUplog(rul)
		if err != nil {
			return 30033,nil
		}
		return ressul.ErrCode,nil
	case SetIP,SetwifiIp:
		rip :=&pb.ReqSetIP{GatewayId:gwid,UserId:uid}
		errCode := UnMarshalJsonBody(body,rip)
		fmt.Println("errCode=:",errCode)
		if errCode != 10000 {
			return errCode,nil
		}
		//rpc调用
		ressip ,err:= rpcClient.SendSetIP(rip)
		if err != nil {
			return 30033,nil
		}
		return ressip.ErrCode,nil
	case SetDNS,SetWifiDns:
		rsd :=&pb.ReqSetDNS{GatewayId:gwid,UserId:uid}
		errCode := UnMarshalJsonBody(body,rsd)
		fmt.Println("errCode=:",errCode)
		if errCode != 10000 {
			return errCode,nil
		}
		//rpc调用
		resdns ,err:= rpcClient.SendSetDNS(rsd)
		if err != nil {
			return 30033,nil
		}
		return resdns.ErrCode,nil
	case SetHotSpot:
		rshs :=&pb.ReqSetHotSpot{GatewayId:gwid,UserId:uid}
		errCode := UnMarshalJsonBody(body,rshs)
		fmt.Println("errCode=:",errCode)
		if errCode != 10000 {
			return errCode,nil
		}
		//rpc调用
		resshs ,err:= rpcClient.SendSetHotSpot(rshs)
		if err != nil {
			return 30033,nil
		}
		return resshs.ErrCode,nil
	case GatewayName:
		rsn :=&pb.ReqSetName{GatewayId:gwid,UserId:uid}
		errCode := UnMarshalJsonBody(body,rsn)
		fmt.Println("errCode=:",errCode)
		if errCode != 10000 {
			return errCode,nil
		}
		//rpc调用
		ressn ,err:= rpcClient.SendName(rsn)
		if err != nil {
			return 30033,nil
		}
		return ressn.ErrCode,nil
	case ConnectWifi:
		rcw :=&pb.ReqConnectWifi{GatewayId:gwid,UserId:uid}
		errCode := UnMarshalJsonBody(body,rcw)
		fmt.Println("errCode=:",errCode)
		if errCode != 10000 {
			return errCode,nil
		}
		//rpc调用
		resscw ,err:= rpcClient.SendConnectWifi(rcw)
		if err != nil {
			return 30033,nil
		}
		return resscw.ErrCode,nil
	case SetUsbIp:
		rcw :=&pb.ReqSetUsbIP{GatewayId:gwid,UserId:uid}
		errCode := UnMarshalJsonBody(body,rcw)
		fmt.Println("errCode=:",errCode)
		if errCode != 10000 {
			return errCode,nil
		}
		//rpc调用
		resscw,err := rpcClient.SendSetUsbIp(rcw)
		if err != nil {
			return 30033,nil
		}
		return resscw.ErrCode,nil
	case SetUsbDns:
		rcw :=&pb.ReqSetUsbDNS{GatewayId:gwid,UserId:uid}
		errCode := UnMarshalJsonBody(body,rcw)
		fmt.Println("errCode=:",errCode)
		if errCode != 10000 {
			return errCode,nil
		}
		//rpc调用
		resscw ,err:= rpcClient.SendSetUsbDns(rcw)
		if err != nil {
			return 30033,nil
		}
		return resscw.ErrCode,nil
	case UsbconnectionWifi:
		rcw :=&pb.ReqConnectUsbWifi{GatewayId:gwid,UserId:uid}
		errCode := UnMarshalJsonBody(body,rcw)
		fmt.Println("errCode=:",errCode)
		if errCode != 10000 {
			return errCode,nil
		}
		//rpc调用
		resscw ,err:= rpcClient.SendUsbConnectionWifi(rcw)
		if err != nil {
			return 30033,nil
		}
		return resscw.ErrCode,nil
	case Usbwifi,Usbwifiscan,UsbGCard,UsbHotSpot,UnConnectUsbWifi:
		log.Info("userid === : ",uid)
		rcw :=&pb.ReqUsbSwitch{GatewayId:gwid,UserId:uid}
		errCode := UnMarshalJsonBody(body,rcw)
		fmt.Println("errCode=:",errCode)
		if errCode != 10000 {
			return errCode,nil
		}
		//rpc调用
		resscw,err := rpcClient.SendUsbSwitch( rcw)
		if err != nil {
			return 30033,nil
		}
		return resscw.ErrCode,nil
	case UsbSetHotSpot:
		rcw := &pb.ReqSetUsbHotSpot{GatewayId:gwid,UserId:uid}
		errCode := UnMarshalJsonBody(body,rcw)
		if errCode != 10000 {
			 return errCode,nil
		}
		res ,err:= rpcClient.SendUsbSetHotSpot(rcw)
		if err != nil {
			return 30033,nil
		}
		return res.ErrCode,nil
	default:
		return 30022,errors.New("instruction error")
	}
	return 30022,errors.New("validation instruction error")
}
////发送邮件或短信
//func SendCode(sso *pb.SsoRequest, res http.ResponseWriter) bool {
//	if REGEXP_MOBILE.MatchString(sso.Username) {
//		sso.CodeType = 1
//		ssoR,ssoErr := rpc.SsoRpc(sso, "SendMobileCode")
//		if ssoErr != nil {
//			log.Error("SendCode- Error",ssoErr)
//			result.RESC(10001, res)
//			return true
//		}
//		if ssoR.ErrorCode != 10000 {
//			result.RESC(ssoR.ErrorCode, res)
//			return true
//		}
//		result.RESC(10000, res)
//		return false
//	}
//	if REGEXP_MAIL.MatchString(sso.Username) {
//		ssoR,ssoErr := rpc.SsoRpc(sso, "FindPasswordByMail")
//		if ssoErr != nil {
//			log.Error("SendCode- Error",ssoErr)
//			result.RESC(10001, res)
//			return true
//		}
//		if ssoR.ErrorCode != 10000 {
//			result.RESC(ssoR.ErrorCode, res)
//			return true
//		}
//		result.RESC(10000, res)
//		return false
//	}
//	result.RESC(20002, res)
//	return true
//}
//









