package rpcClient

import (
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"cotx-http/pb"
	"time"
)

var (
	ssoRpcClient pb.SsoClient
)

const (
	//ssoServer = "rpc.sso.radacat.com:8003"
	ssoServer = "localhost:7007"

)
func NewSsoRpcClient() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(ctx, ssoServer, grpc.WithInsecure())
	if err != nil {
		log.Error("SsoRpcInit Error:", err)
		return
	}
	ssoRpcClient = pb.NewSsoClient(conn)
}

func GetSsoRpcClient() pb.SsoClient {
	return ssoRpcClient
}

//通过用户名查询用户信息（验证账号重复）
func UpdateRedis(req *pb.SsoRequest) *pb.SsoReply {
	log.Info("Start Invoking SsoRpc_GetUserByName, SsoRequest:",req)
	ssoReply, _ := GetSsoRpcClient().UpdateRedis(context.Background(), req)
	log.Info("Finish Invoking SsoRpc_GetUserByName, SsoReply:",ssoReply)
	return ssoReply
}

//通过用户名查询用户信息（验证账号重复）
func JudgeUsername(req *pb.SsoRequest) *pb.SsoReply {
	log.Info("Start Invoking SsoRpc_GetUserByName, SsoRequest:",req)
	ssoReply, _ := GetSsoRpcClient().JudgeUsername(context.Background(), req)
	log.Info("Finish Invoking SsoRpc_GetUserByName, SsoReply:",ssoReply)
	return ssoReply
}

//登录
func Login(req *pb.SsoRequest) *pb.SsoReply {
	log.Info("Start Invoking SsoRpc_Login, SsoRequest:",req)
	ssoReply, _ := GetSsoRpcClient().Login(context.Background(), req)
	log.Info("Finish Invoking SsoRpc_Login, SsoReply:",ssoReply)
	return ssoReply
}

//获取用户基本信息
func GetUserInfo(req *pb.SsoRequest) *pb.SsoReply {
	log.Info("Start Invoking SsoRpc_GetUserInfo, SsoRequest:",req)
	ssoReply, _ := GetSsoRpcClient().GetUserInfo(context.Background(), req)
	log.Info("Finish Invoking SsoRpc_GetUserInfo, SsoReply:",ssoReply)
	return ssoReply
}

//注册用户信息
func Register(req *pb.SsoRequest) *pb.SsoReply {
	log.Info("Start Invoking SsoRpc_Add, SsoRequest:",req)
	ssoReply, _ := GetSsoRpcClient().Add(context.Background(), req)
	log.Info("Finish Invoking SsoRpc_Add, SsoReply:",ssoReply)
	return ssoReply
}

//验证密码
func CheckPassword(req *pb.SsoRequest) *pb.SsoReply {
	log.Info("Start SsoRpc_CheckPassword, SsoRequest:", req)
	ssoReply, _ := GetSsoRpcClient().CheckPassword(context.Background(), req)
	log.Info("finish SsoRpc_CheckPassword, SsoReply:", ssoReply)
	return ssoReply
}

//修改密码
func UpdatePassword(req *pb.SsoRequest) *pb.SsoReply {
	log.Info("Start Invoking SsoRpc_UpdatePassword, SsoRequest:",req)
	ssoReply, _ := GetSsoRpcClient().UpdatePassword(context.Background(), req)
	log.Info("Finish Invoking SsoRpc_UpdatePassword, SsoReply:",ssoReply)
	return ssoReply
}

//登出
func Logout(req *pb.SsoRequest) *pb.SsoReply {
	log.Info("Start Invoking SsoRpc_UpdatePassword, SsoRequest:",req)
	ssoReply, _ := GetSsoRpcClient().Logout(context.Background(), req)
	log.Info("Finish Invoking SsoRpc_UpdatePassword, SsoReply:",ssoReply)
	return ssoReply
}

//发送验证码
func SendMobileCode(req *pb.SsoRequest) *pb.SsoReply {
	log.Info("Start Invoking SsoRpc_SendMobileCode, SsoRequest:",req)
	ssoReply, _ := GetSsoRpcClient().SendMobileCode(context.Background(), req)
	log.Info("Finish Invoking SsoRpc_SendMobileCode, SsoReply:",ssoReply)
	return ssoReply
}

//手机校验验证码
func CheckCode(req *pb.SsoRequest) *pb.SsoReply {
	log.Info("rpc_sso.go/CheckCode:   SsoRequest: ",req)
	ssoReply, _ := GetSsoRpcClient().CheckCode(context.Background(), req)
	log.Info("rpc_sso.go/CheckCode:   SsoReply:",ssoReply)
	return ssoReply
}

//手机重置密码
func ResetPasswordByPhone(rep *pb.SsoRequest) *pb.SsoReply {
	log.Info("rpc_sso.go/ResetPasswordByPhone: SsoRequest = ", rep)
	ssoReply, _ := GetSsoRpcClient().ResetPasswordByPhone(context.Background(), rep)
	log.Info("rpc_sso.go/ResetPasswordByPhone:   SsoReply:",ssoReply)
	return ssoReply
}
//邮箱发送校验链接（未登入状态）
func FindPasswordByMail(rep *pb.SsoRequest) *pb.SsoReply {
	log.Info("rpc_sso.go/FindPasswordByMail: SsoRequest = ", rep)
	ssoReply, _ := GetSsoRpcClient().FindPasswordByMail(context.Background(), rep)
	log.Info("rpc_sso.go/FindPasswordByMail:   SsoReply:",ssoReply)
	return ssoReply
}

//邮箱重置密码
func ResetPassword(rep *pb.SsoRequest) *pb.SsoReply {
	log.Info("rpc_sso.go/ResetPassword:   SsoRequest = ", rep)
	ssoReply, _ := GetSsoRpcClient().ResetPassword(context.Background(), rep)
	log.Info("rpc_sso.go/ResetPassword:   SsoReply:",ssoReply)
	return ssoReply
}