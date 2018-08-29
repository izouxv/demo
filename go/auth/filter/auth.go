package filter

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"golang.org/x/net/context"
	"auth/common"
	. "auth/util"
	log "github.com/cihub/seelog"
	"strconv"
	"time"
)

func AuthInterceptor() grpc.UnaryServerInterceptor {
	log.Info("start Auth Interceptor")
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx, err = auth(ctx)
		if err != nil {
			return
		}
		// 继续处理请求
		return handler(ctx, req)
	}
}

// auth 验证Token
func auth(ctx context.Context) (context.Context,error) {
	md, ok := metadata.FromContext(ctx)
	log.Info("1111111111111111111md:", md)

	if !ok {
		return ctx, SystemError
	}
	var (
		token     string
		url       string
		opt       string
		tid       int32
		did       int32
		source	  string
	)
	var val []string
	if val, ok = md["source"]; ok {
		source = val[0]
	}
	log.Info("Source:",source)
	if source == "dms" {
		return ctx,nil
	}
	if source == "notification" {
		return ctx,nil
	}
	if source == "account-domain-http" {
		return ctx,nil
	}
	if val, ok = md["token"]; ok {
		token = val[0]
	}
	if val, ok = md["url"]; ok {
		url = val[0]
		log.Info("url:", url)
	}
	if val, ok = md["opt"]; ok {
		opt = val[0]
	}
	if val, ok = md["tid"]; ok && val[0] != ""{
		tidInt, _ := strconv.Atoi(val[0])
		tid = int32(tidInt)
	}
	log.Info("auth tid:",tid)
	if val, ok = md["did"]; ok && val[0] != ""{
		didInt, _ := strconv.Atoi(val[0])
		did = int32(didInt)
	}
	log.Info("auth did:",did)
	//todo 校验url是否在白名单内
	if common.Drbac.CheckWhitelist(url, opt) {
		if token != "" {
			userToken := common.Drbac.GetAuthorizationInfo(token)
			/*判断用户的登录状态*/
			if userToken == nil {
				log.Errorf("授权失败，用户未登录，token失效 token is %s", token)
				return ctx,TokenIsInvalid
			}
			ctx = context.WithValue(ctx, "userToken", userToken)
		}
		return ctx,nil
	}
	log.Info("url不在白名单，开始校验")
	if token == "" || url == "" || opt == "" {
		log.Infof("校验参数为空, token:",token, "url:",url, "opt:",opt)
		return ctx,InvalidArgument
	}
	userToken := common.Drbac.GetAuthorizationInfo(token)
	if userToken == nil {
		log.Errorf("授权失败，用户未登录，token失效 token is %s", token)
		return ctx,TokenIsInvalid
	}
	start := time.Now()
	if tid == 0 && did == 0{
		log.Error("Tid或Did值有误，权限校验失败")
		return ctx,InvalidArgument
	}
	if did != 0 {
		//todo 域权限校验
		userToken, auth := common.Drbac.AuthorizationDomain(did, url, opt, token)
		if !auth {
			/*判断用户的登录状态*/
			if userToken == nil {
				log.Errorf("授权失败，用户未登录，token失效 token is %s", token)
				return ctx,TokenIsInvalid
			}
			/*用户没有权限访问*/
			log.Errorf("授权失败，用户没有权限，url is (%s),opt is (%s),token is (%s)", url, opt, token)
			return ctx,PermissionDenied
		}
		log.Infof("授权成功,username is %#v", userToken)
		// 用戶信息存上下文
		ctx = context.WithValue(ctx, "userToken", userToken)
	return ctx, nil
	}else {
		//todo 租户权限校验
		userToken, auth := common.Drbac.AuthorizationTenant(tid, url, opt, token)
		log.Info("调用时间:", time.Now().Sub(start))
		if !auth {
			/*判断用户的登录状态*/
			if userToken == nil {
				log.Errorf("授权失败，用户未登录，token失效 token is %s", token)
				return ctx,TokenIsInvalid
			}
			/*用户没有权限访问*/
			log.Errorf("授权失败，用户没有权限，url is (%s),opt is (%s),token is (%s)", url, opt, token)
			return ctx,PermissionDenied
		}
	}
	log.Infof("授权成功,username is %#v", userToken)
	// 用戶信息存上下文
	ctx = context.WithValue(ctx, "userToken", userToken)
	return ctx, nil
}
