package controller

import (
	pb "auth-http/api"
	"auth-http/rpc"
	. "auth-http/util"
	log "github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"strings"
)


type AddDomainUserReq struct {
	Username string  `json:"username,omitempty"`
	Nickname string  `json:"nickname,omitempty"`
	Rids     []int32 `json:"rids,omitempty"`
}

type EnterDomainReq struct {
	Token    string `json:"token,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserRoles struct {
	User  User    `json:"user"`
	Roles []*Role `json:"roles"`
}

type UserRolesTotalCount struct {
	UserRoles []*UserRoles  `json:"user_roles"`
	TotalCount	int32		`json:"total_count"`
}

type Rids struct {
	Rids  []int32    `json:"rids"`
}



//todo 对租户内用户的操作
/*AddTenantUser
邀请用户进入租户
*/
func AddDomainUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------AddTenantUser----------")
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	addUser := AddDomainUserReq{}
	errCode := GetHttpData(r, "application/json;charset=UTF-8", &addUser)
	if errCode != 10000 {
		if errCode == 404 {
			ResCode(http.StatusNotFound, w)
			return
		}
		JsonReply("BodyIsIncorrectOrEmpty", nil, w)
		return
	}
	log.Info("AddUserInfo:", addUser)

	//todo ctx中存入信息
	token := GetCookie(r, TokenCookieKey)
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
	ctx = context.WithValue(ctx, "token", token)
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:", did)
	//todo 调用rpc
	didInt,_:= strconv.Atoi(did)
	did32 := int32(didInt)
	_, err := rpc.DomainRpcClient().AddUserInDomain(ctx, &pb.AddUserInDomainRequest{Did: did32, AddUserUsername: addUser.Username, AddUserNickname: addUser.Nickname, AddUserRids: addUser.Rids})
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

/*EnterTenant
同意进入租户
*/
func EnterDomain(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------EnterDomain----------")
	reqInfo := EnterDomainReq{}
	errCode := GetHttpData(r, "application/json;charset=UTF-8", &reqInfo)
	if errCode != 10000 {
		if errCode == 404 {
			ResCode(http.StatusNotFound, w)
			return
		}
		JsonReply("BodyIsIncorrectOrEmpty", nil, w)
		return
	}
	log.Info("ReqInfo:", reqInfo)

	//todo ctx中存入信息
	opt := r.Method
	url := strings.Split(r.RequestURI, "?")[0]
	log.Infof("url is (%s),split url is %s", r.RequestURI, url)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "url", url)
	ctx = context.WithValue(ctx, "opt", opt)
	log.Info("url:", ctx.Value("url"))

	//todo 调用rpc
	_, err := rpc.DomainRpcClient().EnterDomain(ctx, &pb.EnterDomainRequest{Token: reqInfo.Token, Password: reqInfo.Password})
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

/*UpdateTenantUserRole
修改租户用户角色
*/
func UpdateDomainUserRole(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------UpdateTenantUserRole----------")
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	uidInt, err := strconv.Atoi(p.ByName("uid"))
	if uidInt == 0 || err != nil {
		log.Error("Get tid Failed , ")
		JsonReply("UidIsIncorrectOrEmpty", nil, w)
		return
	}
	updateUserId := int32(uidInt)
	rids := Rids{}
	errCode := GetHttpData(r, "application/json;charset=UTF-8", &rids)
	if errCode != 10000 {
		if errCode == 404 {
			ResCode(http.StatusNotFound, w)
			return
		}
		JsonReply("BodyIsIncorrectOrEmpty", nil, w)
		return
	}
	log.Info("rids:", rids)

	//todo ctx中存入信息
	token := GetCookie(r, TokenCookieKey)
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
	ctx = context.WithValue(ctx, "token", token)
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:", did)
	//todo 调用rpc
	_, err = rpc.DomainRpcClient().UpdateUserRoleInDomain(ctx, &pb.UpdateUserRoleInDomainRequest{UpdateUserID:updateUserId, UpdateUserRids:rids.Rids})
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

/*UpdateTenantUserState
修改租户用户状态
*/
func UpdateDomainUserState(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------UpdateTenantUserState----------")
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	uidInt, err := strconv.Atoi(p.ByName("uid"))
	if uidInt == 0 || err != nil {
		log.Error("Get tid Failed , ")
		JsonReply("UidIsIncorrectOrEmpty", nil, w)
		return
	}
	updateUserId := int32(uidInt)
	user := User{}
	errCode := GetHttpData(r, "application/json;charset=UTF-8", &user)
	if errCode != 10000 {
		if errCode == 404 {
			ResCode(http.StatusNotFound, w)
			return
		}
		JsonReply("BodyIsIncorrectOrEmpty", nil, w)
		return
	}
	log.Info("state:", user.State)

	//todo ctx中存入信息
	token := GetCookie(r, TokenCookieKey)
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
	ctx = context.WithValue(ctx, "token", token)
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:", did)
	//todo 调用rpc
	_, err = rpc.UserRpcClient().UpdateUserState(ctx, &pb.UpdateUserStateRequest{UpdateUid:updateUserId,UpdateState:user.State})
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

/*GetTenantUsers
获取租户用户列表
*/
func GetDomainUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------GetTenantUsers----------")
	did := p.ByName("did")
	if did == ""{
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	count := r.FormValue("per_page")
	page := r.FormValue("page")
	req := &pb.GetUserInfoInDomainRequest{}
	if page != "" {
		p, err := strconv.Atoi(page)
		if err != nil || p == 0 {
			log.Info("strconv.Atoi(page) Failed,", err)
			JsonReply("PageIsIncorrectOrEmpty", nil, w)
			return
		}
		req.Page = int32(p)
	}
	if count != "" {
		c, err := strconv.Atoi(count)
		if err != nil || c == 0 {
			log.Info("strconv.Atoi(count) Failed,", err)
			JsonReply("PerpageIsIncorrectOrEmpty", nil, w)
			return
		}
		req.Count = int32(c)
	}

	//todo ctx中存入信息
	token := GetCookie(r, TokenCookieKey)
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
	ctx = context.WithValue(ctx, "token", token)
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:", did)
	//todo 调用rpc
	resp, err := rpc.DomainRpcClient().GetUserInfoInDomain(ctx, req)
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
	var reply []*UserRoles
	for _, v := range resp.UserRoles {
		var roles []*Role
		for _, v2 := range v.Roles {
			roles = append(roles, &Role{Rid: v2.Rid, RoleName: v2.RoleName})
		}
		reply = append(reply, &UserRoles{
			User: User{
				Uid:        v.User.Uid,
				Username:   v.User.Username,
				Nickname:   v.User.Nickname,
				State:      v.User.State,
				CreateTime: v.User.CreateTime,
			},
			Roles: roles,
		})
	}
	log.Infof("调用rpc成功%v+", reply)
	JsonReply("Successful", UserRolesTotalCount{UserRoles:reply,  TotalCount:resp.TotalCount}, w)
	return
}

/*DeleteTenantUser
删除租户用户
*/
func DeleteDomainUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------DeleteTenantUser----------")
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	uidInt, err := strconv.Atoi(p.ByName("uid"))
	if uidInt == 0 || err != nil {
		log.Error("Get tid Failed , ")
		JsonReply("UidIsIncorrectOrEmpty", nil, w)
		return
	}
	UserId := int32(uidInt)
	//todo ctx中存入信息
	token := GetCookie(r, TokenCookieKey)
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
	ctx = context.WithValue(ctx, "token", token)
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:", did)
	//todo 调用rpc
	_, err = rpc.DomainRpcClient().DeleteUserInDomain(ctx, &pb.DeleteUserInDomainRequest{DeleteUserID:UserId})
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