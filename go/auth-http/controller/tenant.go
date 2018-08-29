package controller

import (
	pb "auth-http/api"
	"auth-http/rpc"
	. "auth-http/util"
	"encoding/json"
	log "github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type AddUserTenantACLReq struct {
	Username 	string  	`json:"username,omitempty"`
	Tids  		[]int32  	`json:"tids"`
}

type AddTenantByDomainReq struct {
	TenantName  string `json:"tenant_name"`
	TenantURL   string `json:"tenant_url"`
	Description string `json:"description"`
	Contacts    string `json:"contacts"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Icon		string `json:"icon"`
	Logo		string `json:"logo"`
}

type TenantInfo struct {
	Tid  		int32  `json:"tid"`
	TenantName  string `json:"tenant_name"`
	TenantURL   string `json:"tenant_url"`
	Description string `json:"description"`
	State 		int32 `json:"state"`
	Contacts    string `json:"contacts"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	CreateTime  int64  `json:"create_time,omitempty"`
	Icon  		string  `json:"icon,omitempty"`
	Logo  		string  `json:"logo,omitempty"`
}

type GetTidByUrlResponse struct {
	Tid  		int32   `json:"tid"`
	Icon  		string  `json:"icon,omitempty"`
	Logo  		string  `json:"logo,omitempty"`
}

type GetDidByTidResponse struct {
	Did  		int32   `json:"did"`
}

type UpdateTenantStateReq struct {
	Tid         int64 `json:"tid"`
	TenantState int32 `json:"tenantState"`
}

type AddTenantUserReq struct {
	Username string  `json:"username,omitempty"`
	Nickname string  `json:"nickname,omitempty"`
	Rids     []int32 `json:"rids,omitempty"`
}

type EnterTenantReq struct {
	Token    string `json:"token,omitempty"`
	Password string `json:"password,omitempty"`
}



//todo 对租户的操作
/*AddTenantByDomain
添加租户
*/
func AddTenantByDomain(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------AddTenant----------")
	r.ParseForm()
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var tenant AddTenantByDomainReq
	if err := json.Unmarshal(body, &tenant); err != nil {
		log.Errorf("JsonError %s ", err)
		JsonReply("JsonError", nil, w)
		return
	}
	log.Infof("TenantName is %s", tenant.TenantName)
	log.Info("TenantContacts :",tenant.Contacts)
	if tenant.Contacts == "" {
		log.Info("tenant.Contacts is nil")
		JsonReply("BodyIsIncorrectOrEmpty", nil, w)
		return
	}

	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
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
	url := strings.Split(r.RequestURI, "?")[0]
	log.Infof("url is (%s),split url is %s", r.RequestURI, url)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "url", url)
	ctx = context.WithValue(ctx, "opt", opt)
	ctx = context.WithValue(ctx, "token", token)
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:", did)

	didInt,_ := strconv.Atoi(did)
	//todo 调用rpc
	req := &pb.AddTenantRequest{
		TenantName: tenant.TenantName,
		Pid: 0,
		Did: int32(didInt),
		TenantURL:tenant.TenantURL,
		Description:tenant.Description,
		Email:tenant.Email,
		Phone:tenant.Phone,
		Contacts:tenant.Contacts,

	}
	if tenant.Icon != "" {
		req.Icon = "http://file.radacat.com:88/v1.0/file/"+tenant.Icon
	}
	if tenant.Logo != "" {
		req.Logo = "http://file.radacat.com:88/v1.0/file/"+tenant.Logo
	}
	_, err := rpc.TenantRpcClient().AddTenant(ctx, req)
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

/*AddTenantByDomain
再次邀请租户
*/
func InviteUnactivatedTenant(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------InviteUnactivatedTenant----------")
	r.ParseForm()
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var tenant TenantInfo
	if err := json.Unmarshal(body, &tenant); err != nil {
		log.Errorf("JsonError %s ", err)
		JsonReply("JsonError", nil, w)
		return
	}
	log.Infof("TenantName is %s", tenant.TenantName)
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
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
	req := &pb.InviteUnactivatedTenantRequest{
		Tid:tenant.Tid,
		Username:tenant.Email,
	}

	_, err := rpc.TenantRpcClient().InviteUnactivatedTenant(ctx, req)
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


/*UpdateTenant
修改租户信息
*/
func UpdateTenant(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------UpdateTenant----------")
	r.ParseForm()
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var tenant AddTenantByDomainReq
	if err := json.Unmarshal(body, &tenant); err != nil {
		log.Errorf("JsonError %s ", err)
		JsonReply("JsonError", nil, w)
		return
	}
	log.Infof("TenantName is %s", tenant.TenantName)
	if tenant.TenantName == "" {
		log.Error("tenant.TenantName is Empty")
		JsonReply("InvalidArgument", nil, w)
		return
	}
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	tidInt, err := strconv.Atoi(p.ByName("tid"))
	if tidInt == 0 || err != nil {
		log.Error("Get tid Failed , ", err)
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	tid := int32(tidInt)
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
	req := &pb.TenantInfo{
		Tid:tid,
		TenantName: tenant.TenantName,
		TenantURL:tenant.TenantURL,
		Description:tenant.Description,
		Email:tenant.Email,
		Phone:tenant.Phone,
		Contacts:tenant.Contacts,
	}
	if tenant.Icon != "" {
		req.Icon = "http://file.radacat.com:88/v1.0/file/"+tenant.Icon
	}
	if tenant.Logo != "" {
		req.Logo = "http://file.radacat.com:88/v1.0/file/"+tenant.Logo
	}

	_, err = rpc.TenantRpcClient().UpdateTenant(ctx, &pb.UpdateTenantRequest{TenantInfo:req})
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

/*UpdateTenant
修改租户信息
*/
func UpdateTenantState(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------UpdateTenant----------")
	r.ParseForm()
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var tenant TenantInfo
	if err := json.Unmarshal(body, &tenant); err != nil {
		log.Errorf("JsonError %s ", err)
		JsonReply("JsonError", nil, w)
		return
	}
	log.Infof("Tid is %s", tenant.Tid)
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	tidInt, err := strconv.Atoi(p.ByName("tid"))
	if tidInt == 0 || err != nil {
		log.Error("Get tid Failed , ", err)
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	tid := int32(tidInt)
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
	_, err = rpc.TenantRpcClient().UpdateTenantState(ctx, &pb.UpdateTenantStateRequest{Tid:tid,State:tenant.State})
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


/*GetTidByUrl
通过URL获取Tid
*/
func GetTidByUrl(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------GetTidByUrl----------")
	r.ParseForm()
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var tenant AddTenantByDomainReq
	if err := json.Unmarshal(body, &tenant); err != nil {
		log.Errorf("JsonError %s ", err)
		JsonReply("JsonError", nil, w)
		return
	}
	log.Infof("TenantURL is %s", tenant.TenantURL)

	//todo ctx中存入信息
	opt := r.Method
	url := strings.Split(r.RequestURI, "?")[0]
	log.Infof("url is (%s),split url is %s", r.RequestURI, url)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "url", url)
	ctx = context.WithValue(ctx, "opt", opt)
	log.Info("url:", ctx.Value("url"))

	//todo 调用rpc
	resp, err := rpc.TenantRpcClient().GetTidByUrl(ctx, &pb.GetTidByUrlRequest{Url:tenant.TenantURL})
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

	JsonReply("Successful", &GetTidByUrlResponse{Tid:resp.Tid,Icon:resp.Icon,Logo:resp.Logo}, w)
	return
}

/*GetDidByTid
通过Tid获取Did
*/
func GetDidByTid(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------GetDidByTid----------")
	tidInt, err := strconv.Atoi(p.ByName("tid"))
	if tidInt == 0 || err != nil {
		log.Error("Get tid Failed , ", err)
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	tid := int32(tidInt)
	log.Infof("Tid is %s", tid)
	//todo ctx中存入信息
	opt := r.Method
	url := strings.Split(r.RequestURI, "?")[0]
	log.Infof("url is (%s),split url is %s", r.RequestURI, url)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "url", url)
	ctx = context.WithValue(ctx, "opt", opt)
	log.Info("url:", ctx.Value("url"))

	//todo 调用rpc
	resp, err := rpc.TenantRpcClient().GetDidByTid(ctx, &pb.GetDidByTidRequest{Tid:tid})
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

	JsonReply("Successful", &GetDidByTidResponse{Did:resp.Did}, w)
	return
}

/*GetTenants
获取域下租户列表
*/
func GetTenants(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------GetTenants----------")
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
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
	resp, err := rpc.TenantRpcClient().GetTenants(ctx, &pb.GetTenantsRequest{})
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
	var tenants []*TenantInfo
	for _,v := range resp.Tenants {
		tenant := &TenantInfo{
			Tid:v.Tid,
			TenantName:v.TenantName,
			TenantURL:v.TenantURL,
			Description:v.Description,
			State:v.State,
			Contacts:v.Contacts,
			Phone:v.Phone,
			Email:v.Email,
			CreateTime:v.CreateTime,
			Icon:v.Icon,
			Logo:v.Logo,
		}
		tenants = append(tenants, tenant)
	}
	JsonReply("Successful", tenants, w)
	return
}

/*DeleteTenant
删除指定租户
*/
func DeleteTenant(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------DeleteTenant----------")
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	tidInt, err := strconv.Atoi(p.ByName("tid"))
	if tidInt == 0 || err != nil {
		log.Error("Get tid Failed , ", err)
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	tid := int32(tidInt)
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
	_, err = rpc.TenantRpcClient().DeleteTenant(ctx, &pb.DeleteTenantRequest{Tid:tid})
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

/*AddTenantByFather
添加子租户
*/
func AddTenantByFather(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

//todo 对租户内用户的操作
/*AddTenantUser
邀请用户进入租户
*/
func AddTenantUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------AddTenantUser----------")
	tid := p.ByName("tid")
	if tid == "" {
		log.Error("Get tid Failed , ")
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	addUser := AddTenantUserReq{}
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
	ctx = context.WithValue(ctx, "tid", tid)
	log.Info("url:", ctx.Value("url"))
	log.Info("tid:", tid)
	//todo 调用rpc
	tidInt,_ := strconv.Atoi(tid)
	_, err := rpc.TenantRpcClient().AddUserInTenant(ctx, &pb.AddUserInTenantRequest{Tid: int32(tidInt), AddUserUsername: addUser.Username, AddUserNickname: addUser.Nickname, AddUserRids: addUser.Rids})
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
func EnterTenant(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------EnterTenant----------")

	reqInfo := EnterTenantReq{}
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
	_, err := rpc.TenantRpcClient().EnterTenant(ctx, &pb.EnterTenantRequest{Token: reqInfo.Token, Password: reqInfo.Password})
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
func UpdateTenantUserRole(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------UpdateTenantUserRole----------")
	tid := p.ByName("tid")
	if tid == "" {
		log.Error("Get tid Failed , ")
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
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
	ctx = context.WithValue(ctx, "tid", tid)
	log.Info("url:", ctx.Value("url"))
	log.Info("tid:", tid)
	//todo 调用rpc
	_, err = rpc.TenantRpcClient().UpdateUserRoleInTenant(ctx, &pb.UpdateUserRoleInTenantRequest{UpdateUserID:updateUserId, UpdateUserRids:rids.Rids})
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
func UpdateTenantUserState(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------UpdateTenantUserState----------")
	tid := p.ByName("tid")
	if tid == "" {
		log.Error("Get tid Failed , ")
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
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
	ctx = context.WithValue(ctx, "tid", tid)
	log.Info("url:", ctx.Value("url"))
	log.Info("tid:", tid)
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
func GetTenantUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------GetTenantUsers----------")
	tid := p.ByName("tid")
	if tid == "" {
		log.Error("Get tid Failed , ")
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	count := r.FormValue("per_page")
	page := r.FormValue("page")
	req := &pb.GetUserInfoInTenantRequest{}
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

	log.Info("GetTenantUsers,","count:", count, ", page:", page)
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
	ctx = context.WithValue(ctx, "tid", tid)
	log.Info("url:", ctx.Value("url"))
	log.Info("tid:", tid)
	//todo 调用rpc
	resp, err := rpc.TenantRpcClient().GetUserInfoInTenant(ctx, req)
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
	JsonReply("Successful", &UserRolesTotalCount{UserRoles:reply, TotalCount:resp.TotalCount}, w)
	return
}

/*DeleteTenantUser
删除租户用户
*/
func DeleteTenantUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------DeleteTenantUser----------")
	tid := p.ByName("tid")
	if tid == "" {
		log.Error("Get tid Failed , ")
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	uidInt, err := strconv.Atoi(p.ByName("uid"))
	if uidInt == 0 || err != nil {
		log.Error("Get tid Failed , ")
		JsonReply("UidIsIncorrectOrEmpty", nil, w)
		return
	}
	userId := int32(uidInt)
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
	ctx = context.WithValue(ctx, "tid", tid)
	log.Info("url:", ctx.Value("url"))
	log.Info("tid:", tid)
	//todo 调用rpc
	_, err = rpc.TenantRpcClient().DeleteUserInTenant(ctx, &pb.DeleteUserInTenantRequest{DeleteUserID:userId})
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



/*AddUserTenantACL
添加用户租户ACL
*/
func AddUserTenantACL(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------AddUserTenantACL----------")
	r.ParseForm()
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var uta AddUserTenantACLReq
	if err := json.Unmarshal(body, &uta); err != nil {
		log.Errorf("JsonError %s ", err)
		JsonReply("JsonError", nil, w)
		return
	}
	log.Info("username:", uta.Username, ",tids:",uta.Tids)
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
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
	didInt,_ := strconv.Atoi(did)
	req := &pb.AddUserTenantACLRequest{
		Username:uta.Username,
		Tids:uta.Tids,
		Did:int32(didInt),
	}
	_, err := rpc.DomainRpcClient().AddUserTenantACL(ctx, req)
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