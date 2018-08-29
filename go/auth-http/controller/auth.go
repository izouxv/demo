package controller

import (
	pb "auth-http/api"
	"auth-http/rpc"
	. "auth-http/util"
	"context"
	"encoding/json"
	log "github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"strconv"
)

type LoginReq struct {
	Tid      int32  `json:"tid,omitempty"`
	Did      int32  `json:"did,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Tenant struct {
	Tid        int32  `json:"tid"`
	TenantName string `json:"tenant_name"`
	CreateTime int64  `json:"create_time"`
}

type LoginResponse struct {
	User    *User     `json:"user"`
	Tenants []*Tenant `json:"tenants,omitempty"`
}

/*Login
登录租户
*/
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("-----Login------")
	var user LoginReq
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Errorf("Json_is_error %s ", err)
		JsonReply("JsonError", nil, w)
		return
	}
	log.Infof("username is %s,password is %s", user.Username, user.Password)
	username := user.Username
	password := user.Password
	if !UsernameRegexp.MatchString(username) {
		log.Infof("username is err (%s)", username)
		JsonReply("UsernameIsIncorrectOrEmpty", nil, w)
		return
	}
	if password == "" {
		log.Info("password is nil")
		JsonReply("PasswordIsIncorrectOrEmpty", nil, w)
		return
	}
	if user.Did != 0 {
		didString :=strconv.Itoa(int(user.Did))
			//todo 登录域
		log.Info("开始登录域")
		opt := r.Method
		url := strings.Split(r.RequestURI, "?")[0]
		log.Infof("url is (%s),split url is %s", r.RequestURI, url)
		ctx := context.Background()
		ctx = context.WithValue(ctx, "url", url)
		ctx = context.WithValue(ctx, "opt", opt)
		ctx = context.WithValue(ctx, "did", didString)
		log.Info("url:", ctx.Value("url"))
		log.Info("did:", didString)
		/*调用rpc的服务*/
		start := time.Now()
		resp, err := rpc.AuthRpcClient().AuthenticationWithDid(ctx, &pb.AuthenticationRequest{Username: username, Password: password})
		log.Info("服务调用时间:", time.Now().Sub(start))
		//todo 添加操作日志
		domainActionLog := &pb.ActionLogInfo{
			ActionUsername: username,
			ActionTime:     time.Now().Unix(),
			ActionType:     1,
			ActionName:     "登录",
			ActionObject:	"成功",
			Tid:user.Tid,
			Did:user.Did,
		}
		if err != nil {
			log.Errorf("调用rpc错误码,err is (%s)", err)
			s, ok := status.FromError(err)
			if !ok {
				log.Errorf("系统异常,", err)
				JsonReply("SystemError", nil, w)
				return
			}
			code := s.Proto().Code
			for k, v := range CodeMap {
				if code == v.Code {
					log.Info("ErrorInfo:", k)
					if code != 10000 {
						domainActionLog.ActionType = 2
						domainActionLog.ActionObject = v.Msg
					}
					go rpc.ActionLogRpcClient().AddActionLog(ctx,&pb.AddActionLogRequest{ActionLog:domainActionLog})
					JsonReply(k, nil, w)
					return
				}
			}
			JsonReply("SystemError", nil, w)
			return
		}
		returnUser := &User{
			Uid:        resp.UserToken.User.Uid,
			Username:   resp.UserToken.User.Username,
			Nickname:   resp.UserToken.User.Nickname,
			LoginState: resp.UserToken.User.LoginState,
			Token:      resp.UserToken.Token,
		}
		go rpc.ActionLogRpcClient().AddActionLog(ctx,&pb.AddActionLogRequest{ActionLog:domainActionLog})
		JsonReply("Successful", &LoginResponse{User: returnUser}, w)
		return
	} else {
		log.Info("开始登录租户")
		if user.Tid == 0 {
			log.Errorf("tid为空,", err)
			JsonReply("TidIsIncorrectOrEmpty", nil, w)
			return
		}
		tidString :=strconv.Itoa(int(user.Tid))
		opt := r.Method
		url := strings.Split(r.RequestURI, "?")[0]
		log.Infof("url is (%s),split url is %s", r.RequestURI, url)
		ctx := context.Background()
		ctx = context.WithValue(ctx, "url", url)
		ctx = context.WithValue(ctx, "opt", opt)
		ctx = context.WithValue(ctx, "tid", tidString)
		log.Info("url:", ctx.Value("url"))
		log.Info("tid:", tidString)

		/*调用rpc的服务*/
		start := time.Now()
		resp, err := rpc.AuthRpcClient().AuthenticationWithTid(ctx, &pb.AuthenticationRequest{Username: username, Password: password})
		log.Info("服务调用时间:", time.Now().Sub(start))
		//todo 添加操作日志
		tenantActionLog := &pb.ActionLogInfo{
			ActionUsername: username,
			ActionTime:     time.Now().Unix(),
			ActionType:     1,
			ActionName:     "登录",
			ActionObject:	"成功",
			Tid:user.Tid,
			Did:user.Did,
		}
		if err != nil {
			log.Errorf("调用rpc错误码,err is (%s)", err)
			s, ok := status.FromError(err)
			if !ok {
				log.Errorf("系统异常,", err)
				JsonReply("SystemError", nil, w)
				return
			}
			code := s.Proto().Code
			for k, v := range CodeMap {
				if code == v.Code {
					log.Info("ErrorInfo:", k)
					if code != 10000 {
						tenantActionLog.ActionType = 2
						tenantActionLog.ActionObject = v.Msg
					}
					go rpc.ActionLogRpcClient().AddActionLog(ctx,&pb.AddActionLogRequest{ActionLog:tenantActionLog})
					JsonReply(k, nil, w)
					return
				}
			}
			JsonReply("SystemError", nil, w)
			return
		}
		returnUser := &User{
			Uid:        resp.UserToken.User.Uid,
			Username:   resp.UserToken.User.Username,
			Nickname:   resp.UserToken.User.Nickname,
			LoginState: resp.UserToken.User.LoginState,
			Token:      resp.UserToken.Token,
		}
		var tenants []*Tenant
		for _, v := range resp.UserToken.TenantRoleTree {
			tenant := &Tenant{
				Tid:        v.Tenant.Tid,
				TenantName: v.Tenant.TenantName,
				CreateTime: v.Tenant.CreateTime,
			}
			tenants = append(tenants, tenant)
		}
		go rpc.ActionLogRpcClient().AddActionLog(ctx,&pb.AddActionLogRequest{ActionLog:tenantActionLog})
		JsonReply("Successful", &LoginResponse{User: returnUser, Tenants: tenants}, w)
		return
	}
}

/*Login
登录域
*/
func LoginWithDid(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("-----LoginWithoutTid------")
	var user User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Errorf("Json_is_error %s ", err)
		JsonReply("JsonError", nil, w)
		return
	}
	log.Infof("username is %s,password is %s", user.Username, user.Password)
	username := user.Username
	password := user.Password
	if !UsernameRegexp.MatchString(username) {
		log.Infof("username is err (%s)", username)
		JsonReply("UsernameIsIncorrectOrEmpty", nil, w)
		return
	}
	if password == "" {
		log.Info("password is nil")
		JsonReply("PasswordIsIncorrectOrEmpty", nil, w)
		return
	}
	opt := r.Method
	url := strings.Split(r.RequestURI, "?")[0]
	log.Infof("url is (%s),split url is %s", r.RequestURI, url)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "url", url)
	ctx = context.WithValue(ctx, "opt", opt)
	log.Info("url:", ctx.Value("url"))

	/*调用rpc的服务*/
	start := time.Now()
	resp, err := rpc.AuthRpcClient().AuthenticationWithDid(ctx, &pb.AuthenticationRequest{Username: username, Password: password})
	log.Info("服务调用时间:", time.Now().Sub(start))
	if err != nil {
		log.Errorf("调用rpc错误码,err is (%s)", err)
		s, ok := status.FromError(err)
		if !ok {
			log.Errorf("系统异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
		switch s.Proto().Code {
		//系统级错误
		case CodeMap["Successful"].Code:
			log.Info("成功,", err)
			JsonReply("Successful", nil, w)
			return
		case CodeMap["SystemError"].Code:
			log.Errorf("系统异常,", err)
			JsonReply("SystemError", nil, w)
			return
		case CodeMap["InvalidArgument"].Code:
			log.Errorf("参数异常,", err)
			JsonReply("InvalidArgument", nil, w)
			return
		case CodeMap["TokenIsInvalid"].Code:
			log.Info("Token失效或未登录，", err)
			JsonReply("TokenIsInvalid", nil, w)
			return
		case CodeMap["PermissionDenied"].Code:
			log.Info("权限不足，拒绝访问，", err)
			JsonReply("PermissionDenied", nil, w)
			return
		case CodeMap["UserKickedOut"].Code:
			log.Info("账号被踢出，", err)
			JsonReply("UserKickedOut", nil, w)
			return
			//账号错误
		case CodeMap["UserDoesNotExist"].Code:
			log.Info("用户不存在，", err)
			JsonReply("UserDoesNotExist", nil, w)
			return
		case CodeMap["AccountNotActive"].Code:
			log.Info("账号未激活，", err)
			JsonReply("AccountNotActive", nil, w)
			return
		case CodeMap["AccountDisableToUse"].Code:
			log.Info("账号被禁用，", err)
			JsonReply("AccountDisableToUse", nil, w)
			return
		case CodeMap["AccountException"].Code:
			log.Info("账号异常，", err)
			JsonReply("AccountException", nil, w)
			return
		case CodeMap["UsernameAndPasswordError"].Code:
			log.Info("用户名密码错误，", err)
			JsonReply("UsernameAndPasswordError", nil, w)
			return
		default:
			log.Errorf("未知异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
	}
	returnUser := &User{
		Uid:        resp.UserToken.User.Uid,
		Username:   resp.UserToken.User.Username,
		Nickname:   resp.UserToken.User.Nickname,
		LoginState: resp.UserToken.User.LoginState,
		Token:      resp.UserToken.Token,
	}
	var tenants []*Tenant
	for _, v := range resp.UserToken.TenantRoleTree {
		tenant := &Tenant{
			Tid:        v.Tenant.Tid,
			TenantName: v.Tenant.TenantName,
			CreateTime: v.Tenant.CreateTime,
		}
		tenants = append(tenants, tenant)
	}
	JsonReply("Successful", &LoginResponse{User: returnUser, Tenants: tenants}, w)
	return
}

/*
获取用户信息
*/
func GetUserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("-------start GetUserInfo--------")
	token := GetCookie(r, "token")
	log.Infof(" get token is (%s)", token)
	if token == "" {
		JsonReply("TokenIsInvalid", nil, w)
		return
	}
	opt := r.Method
	url := strings.Split(r.RequestURI, "?")[0]
	log.Infof("url is (%s),split url is %s", r.RequestURI, url)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "url", url)
	ctx = context.WithValue(ctx, "opt", opt)
	log.Info("url:", ctx.Value("url"))
	resp, err := rpc.AuthRpcClient().GetAuthorizationInfo(ctx, &pb.GetAuthorizationInfoRequest{Token: token})
	if err != nil {
		log.Errorf("调用rpc错误码,err is (%s)", err)
		s, ok := status.FromError(err)
		if !ok {
			log.Errorf("系统异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
		switch s.Proto().Code {
		//系统级错误
		case CodeMap["Successful"].Code:
			log.Info("成功,", err)
			JsonReply("Successful", nil, w)
			return
		case CodeMap["SystemError"].Code:
			log.Errorf("系统异常,", err)
			JsonReply("SystemError", nil, w)
			return
		case CodeMap["InvalidArgument"].Code:
			log.Errorf("参数异常,", err)
			JsonReply("InvalidArgument", nil, w)
			return
		case CodeMap["UserDoesNotExist"].Code:
			log.Info("账号不存在，", err)
			JsonReply("UserDoesNotExist", nil, w)
			return
		case CodeMap["UserKickedOut"].Code:
			log.Info("账号被踢出，", err)
			JsonReply("UserKickedOut", nil, w)
			return
		case CodeMap["PermissionDenied"].Code:
			log.Info("权限不足，拒绝访问，", err)
			JsonReply("PermissionDenied", nil, w)
			return
		case CodeMap["TokenIsInvalid"].Code:
			log.Info("Token失效或未登录，", err)
			JsonReply("TokenIsInvalid", nil, w)
			return
		default:
			log.Errorf("未知异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
	}
	returnUser := &User{
		Uid:        resp.UserToken.User.Uid,
		Username:   resp.UserToken.User.Username,
		Nickname:   resp.UserToken.User.Nickname,
		LoginState: resp.UserToken.User.LoginState,
		Token:      resp.UserToken.Token,
	}
	var tenants []*Tenant
	for _, v := range resp.UserToken.TenantRoleTree {
		tenant := &Tenant{
			Tid:        v.Tenant.Tid,
			TenantName: v.Tenant.TenantName,
			CreateTime: v.Tenant.CreateTime,
		}
		tenants = append(tenants, tenant)
	}
	JsonReply("Successful", &LoginResponse{User: returnUser,Tenants:tenants}, w)
	return
}

/*LoginOut
退出登录
*/
func LoginOut(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start LoginOut")
	token := GetCookie(r, "token")
	log.Infof(" get token is (%s)", token)
	if token == "" {
		JsonReply("TokenIsInvalid", nil, w)
		return
	}
	opt := r.Method
	url := strings.Split(r.RequestURI, "?")[0]
	log.Infof("url is (%s),split url is %s", r.RequestURI, url)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "url", url)
	ctx = context.WithValue(ctx, "opt", opt)
	log.Info("url:", ctx.Value("url"))
	start := time.Now()
	_, err := rpc.AuthRpcClient().Logout(ctx, &pb.LogoutRequest{Token: token})
	log.Info("服务调用时间:", time.Now().Sub(start))
	log.Errorf("get rpc-logout error is (%s) ：", err)
	log.Infof("login out successful")
	if err != nil {
		log.Errorf("调用rpc错误码,err is (%s)", err)
		s, ok := status.FromError(err)
		if !ok {
			log.Errorf("系统异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
		switch s.Proto().Code {
		//系统级错误
		case CodeMap["Successful"].Code:
			log.Info("成功,", err)
			JsonReply("Successful", nil, w)
			return
		case CodeMap["SystemError"].Code:
			log.Errorf("系统异常,", err)
			JsonReply("SystemError", nil, w)
			return
		case CodeMap["InvalidArgument"].Code:
			log.Errorf("参数异常,", err)
			JsonReply("InvalidArgument", nil, w)
			return
		case CodeMap["TokenIsInvalid"].Code:
			log.Info("Token失效或未登录，", err)
			JsonReply("TokenIsInvalid", nil, w)
			return
		case CodeMap["PermissionDenied"].Code:
			log.Info("权限不足，拒绝访问，", err)
			JsonReply("PermissionDenied", nil, w)
			return
		case CodeMap["UserKickedOut"].Code:
			log.Info("账号被踢出，", err)
			JsonReply("UserKickedOut", nil, w)
			return
		default:
			log.Errorf("未知异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
	}
	JsonReply("Successful", nil, w)
	return
}
