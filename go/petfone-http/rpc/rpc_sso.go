package rpc

import (
	"context"
	log "github.com/cihub/seelog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"petfone-http/pb"
	"sync"
	"time"
)

var (
	ssoRpcOnce   sync.Once
	ssoRpcClient pb.SsoClient
	ssoConn *grpc.ClientConn
)

//初始化ssoRpc
func SsoRpcInit(address string) {
	ssoRpcOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		ssoConn, err = grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Error("ShareRpcInit:", err)
			panic(err)
		}
		ssoRpcClient = pb.NewSsoClient(ssoConn)
	})
}

//结束Rpc
func SsoRpcClose() {
	if ssoConn != nil {
		ssoConn.Close()
	}
}

//用户信息rpc
func SsoRpc(sso *pb.SsoRequest, method string) *pb.SsoReply {
	log.Info("SsoRpc-sso:", sso)
	var ssoErr error
	var ssoReply *pb.SsoReply
	switch method {
	//添加用户
	case "Add":
		ssoReply, ssoErr = ssoRpcClient.Add(context.Background(), sso)
		//登录
	case "Login":
		ssoReply, ssoErr = ssoRpcClient.Login(context.Background(), sso)
		//获取用户信息
	case "GetUserInfo":
		ssoReply, ssoErr = ssoRpcClient.GetUserInfo(context.Background(), sso)
		//查询用户名
	case "GetUserByName":
		ssoReply, ssoErr = ssoRpcClient.GetUserByName(context.Background(), sso)
		break
		//校验密码
	case "CheckPassword":
		ssoReply, ssoErr = ssoRpcClient.CheckPassword(context.Background(), sso)
		//修改密码
	case "UpdatePassword":
		ssoReply, ssoErr = ssoRpcClient.UpdatePassword(context.Background(), sso)
		//重置密码
	case "ResetPassword":
		ssoReply, ssoErr = ssoRpcClient.ResetPassword(context.Background(), sso)
		break
		//校验手机code
	case "CheckCode":
		ssoReply, ssoErr = ssoRpcClient.CheckCode(context.Background(), sso)
		//根据用户名修改密码
	case "UpdatePasswordByName":
		ssoReply, ssoErr = ssoRpcClient.UpdatePasswordByName(context.Background(), sso)
		//根据用户名修改密码
	case "Logout":
		ssoReply, ssoErr = ssoRpcClient.Logout(context.Background(), sso)
		//向邮箱发送邮件
	case "FindPasswordByMail":
		ssoReply, ssoErr = ssoRpcClient.FindPasswordByMail(context.Background(), sso)
	case "SendMobileCode":
		ssoReply, ssoErr = ssoRpcClient.SendMobileCode(context.Background(), sso)
	default:
		ssoErr = errors.New("没有该RPC")
	}
	if ssoErr != nil {
		log.Error("SsoRpc错误日志:", ssoErr)
		return &pb.SsoReply{Code:10001}
	}
	log.Info("ssoRPC-ssoReply:", ssoReply)
	return ssoReply
}

