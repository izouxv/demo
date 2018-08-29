package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	pb "auth-http/api"
	"auth-http/rpc"
	. "auth-http/util"
	log "github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc/status"
	"golang.org/x/net/context"
	"strconv"
	"strings"
)

type User struct{
	Uid         int32  `json:"uid"`
	Username    string `json:"username"`
	Password    string `json:"password,omitempty"`
	Nickname    string `json:"nickname"`
	NewPassword string `json:"new_password,omitempty"`
	LoginState  int32  `json:"login_state"`
	Token       string `json:"token,omitempty"`
	State  		int32  `json:"state,omitempty"`
	CreateTime  int64  `json:"create_time,omitempty"`
}

type UserState struct {
	Uid         int32  `json:"uid"`
	State  		int32  `json:"state"`
}

type FindPasswordReq struct {
	Username    string `json:"username"`
	Tid  		int32  `json:"tid"`
	Did  		int32  `json:"did"`
}


/*UpdateUserInfo
修改用户信息
*/
func UpdateUserInfoAndPassword(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	log.Info(" Start  UpdateUserInfo")
	r.ParseForm()
	var user User
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	if err:=json.Unmarshal(body, &user);err !=nil {
		log.Errorf("Json_is_error %s ", err)
		JsonReply("JsonError", nil, w)
		return
	}
	if user.Nickname == ""{
		JsonReply("NicknameIsIncorrectOrEmpty", nil, w)
		return
	}
	//todo ctx中存入信息
	token := GetCookie(r, TokenCookieKey)
	log.Infof(" get token is (%s)", token)
	if token == "" {
		JsonReply("TokenIsInvalid", nil, w)
		return
	}
	opt := r.Method
	url := strings.Split(r.RequestURI,"?")[0]
	log.Infof("url is (%s),split url is %s",r.RequestURI,url)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "url", url)
	ctx = context.WithValue(ctx, "opt", opt)
	ctx = context.WithValue(ctx, "token", token)
	log.Info("url:", ctx.Value("url"))

	log.Info("1111111111111111111111UserInfo:",user)
	if user.Password == "" && user.NewPassword == "" {
		_,err := rpc.UserRpcClient().UpdateUserInfo(ctx,&pb.UpdateUserInfoRequest{Nickname:user.Nickname})
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
					JsonReply(k, nil, w)
					return
				}
			}
			JsonReply("SystemError", nil, w)
			return
		}
	}else {
		if user.Password == "" {
			log.Error("PasswordIsIncorrectOrEmpty")
			JsonReply("PasswordIsIncorrectOrEmpty", nil, w)
			return
		}
		if user.NewPassword == "" {
			log.Error("NewPasswordIsIncorrectOrEmpty")
			JsonReply("NewPasswordIsIncorrectOrEmpty", nil, w)
			return
		}
		_,err := rpc.UserRpcClient().UpdateNicknameAndPassword(ctx,&pb.UpdateNicknameAndPasswordRequest{Nickname:user.Nickname,Password:user.Password,NewPassword:user.NewPassword})
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
					JsonReply(k, nil, w)
					return
				}
			}
			JsonReply("SystemError", nil, w)
			return
		}
		log.Infof("修改密码成功username is %s,password is %s,new password is %s",user.Username,user.Password,user.NewPassword)
		JsonReply("Successful", nil, w)
		return
	}
}
/*UpdatePassword
修改密码
*/
func UpdatePassword(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	log.Info("-------UpdatePassword--------")
	//todo ctx中存入信息
	token := GetCookie(r, TokenCookieKey)
	log.Infof(" get token is (%s)", token)
	if token == "" {
		JsonReply("TokenIsInvalid", nil, w)
		return
	}
	opt := r.Method
	url := strings.Split(r.RequestURI,"?")[0]
	log.Infof("url is (%s),split url is %s",r.RequestURI,url)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "url", url)
	ctx = context.WithValue(ctx, "opt", opt)
	ctx = context.WithValue(ctx, "token", token)
	log.Info("url:", ctx.Value("url"))

	r.ParseForm()
	var user User
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	if err:=json.Unmarshal(body, &user);err !=nil {
		log.Errorf("Json_is_error %s ", err)
		JsonReply("JsonError", nil, w)
		return
	}
	if user.Password == "" {
		log.Error("PasswordIsIncorrectOrEmpty")
		JsonReply("PasswordIsIncorrectOrEmpty", nil, w)
		return
	}
	if user.NewPassword == "" {
		log.Error("NewPasswordIsIncorrectOrEmpty")
		JsonReply("NewPasswordIsIncorrectOrEmpty", nil, w)
		return
	}
	_,err := rpc.UserRpcClient().UpdatePassword(ctx,&pb.UpdatePasswordRequest{Password:user.Password,NewPassword:user.NewPassword})
	//todo 判断错误码
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
				JsonReply(k, nil, w)
				return
			}
		}
		JsonReply("SystemError", nil, w)
		return
	}
	log.Infof("修改密码成功username is %s,password is %s,new password is %s",user.Username,user.Password,user.NewPassword)
	JsonReply("Successful", nil, w)
	return
}
/*FindPassword
邮箱找回密码
*/
func FindPassword(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	log.Info("-------FindPassword--------")
	r.ParseForm()
	var req FindPasswordReq
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	if err:=json.Unmarshal(body, &req);err !=nil {
		log.Errorf("Json_is_error %s ", err)
		JsonReply("Json_is_error", nil, w)
		return
	}
	if !UsernameRegexp.MatchString(req.Username) {
		log.Info("用户名为空或格式不正确")
		JsonReply("UsernameIsIncorrectOrEmpty", nil, w)
		return
	}
	log.Info("FindPasswordReq:",req)
	opt := r.Method
	url := strings.Split(r.RequestURI,"?")[0]
	log.Infof("url is (%s),split url is %s",r.RequestURI,url)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "url", url)
	ctx = context.WithValue(ctx, "opt", opt)
	log.Info("url:", ctx.Value("url"))

	_,err := rpc.UserRpcClient().FindPassword(ctx,&pb.FindPasswordRequest{Username:req.Username,Tid:req.Tid, Did:req.Did})
	//todo 判断错误码
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
				JsonReply(k, nil, w)
				return
			}
		}
		JsonReply("SystemError", nil, w)
		return
	}
	log.Infof("找回密码发送邮件成功，username is %s",req.Username)
	JsonReply("Successful", nil, w)
	return
}
/*ResetPassword
重置密码
*/
func ResetPassword(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("------ResetPassword--------")
	token := p.ByName("token")
	if token == "" {
		log.Error("Get token Failed , ")
		JsonReply("PasswordTokenIsIncorrectOrEmpty", nil, w)
		return
	}
	r.ParseForm()
	var req User
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	if err:=json.Unmarshal(body, &req);err !=nil {
		log.Errorf("Json_is_error %s ", err)
		JsonReply("JsonError", nil, w)
		return
	}
	opt := r.Method
	url := strings.Split(r.RequestURI,"?")[0]
	log.Infof("url is (%s),split url is %s",r.RequestURI,url)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "url", url)
	ctx = context.WithValue(ctx, "opt", opt)
	log.Info("url:", ctx.Value("url"))

	_,err := rpc.UserRpcClient().ResetPassword(ctx,&pb.ResetPasswordRequest{Password:req.Password, Token:token})
	//todo 判断错误码
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
				JsonReply(k, nil, w)
				return
			}
		}
		JsonReply("SystemError", nil, w)
		return
	}
	log.Infof("重置密码成功，username is %s",req.Username)
	JsonReply("Successful", nil, w)
	return
}

/*UpdateUserState
修改用户状态
*/
func UpdateUserState(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	log.Info(" Start  UpdateUserState")
	userState := UserState{}
	errCode := GetHttpData(r, "application/json;charset=UTF-8", &userState)
	if errCode != 10000 {
		if errCode == 404 {
			ResCode(http.StatusNotFound, w)
			return
		}
		JsonReply("BodyIsIncorrectOrEmpty",nil, w)
		return
	}
	uid := p.ByName("uid")
	if uid != "" {
		uidInt,err := strconv.Atoi(uid)
		Updateuid := int32(uidInt)
		if err != nil || Updateuid == 0 {
			log.Error("strconv.Atoi(Rid) Failed,", err)
			JsonReply("InvalidArgument", nil, w)
			return
		}
		userState.Uid = Updateuid
	}
	log.Info("userStateInfo:",userState)

	//todo ctx中存入信息
	token := GetCookie(r, TokenCookieKey)
	log.Infof(" get token is (%s)", token)
	if token == "" {
		JsonReply("TokenIsInvalid", nil, w)
		return
	}
	opt := r.Method
	url := strings.Split(r.RequestURI,"?")[0]
	log.Infof("url is (%s),split url is %s",r.RequestURI,url)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "url", url)
	ctx = context.WithValue(ctx, "opt", opt)
	ctx = context.WithValue(ctx, "token", token)
	log.Info("url:", ctx.Value("url"))

	_,err := rpc.UserRpcClient().UpdateUserState(ctx,&pb.UpdateUserStateRequest{UpdateUid:userState.Uid,UpdateState:userState.State})
	//todo 判断错误码
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
				JsonReply(k, nil, w)
				return
			}
		}
		JsonReply("SystemError", nil, w)
		return
	}
	JsonReply("Successful", nil, w)
	return
}