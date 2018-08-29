package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/cihub/seelog"
	."auth-http/util"
	"auth-http/rpc"
	"golang.org/x/net/context"
	pb "auth-http/api"
	"strconv"
	"google.golang.org/grpc/status"
	"strings"
	"io/ioutil"
	"encoding/json"
)

type RoleInfo struct {
	Rid int32				`json:"rid"`
	RoleName string			`json:"role_name"`
	Description string 		`json:"description"`
	Mids []*Mids 			`json:"mids"`
}

type RoleInfos struct {
	RoleInfo []*RoleInfo 	`json:"role_info"`
	TotalCount	int32		`json:"total_count"`
}

type Mids struct {
	Mid int32 				`json:"mid"`
	ModuleName string 		`json:"module_name"`
	Operation []string		`json:"operation"`
}


type Role struct {
	Rid int32		`json:"rid"`
	RoleName string	`json:"role_name"`
	Description string `json:"description"`
}

type Module struct {
	Mid 		int32		`json:"mid"`
	ModuleName 	string		`json:"module_name"`
}

type RoleModules struct {
	Role 		Role      	`json:"role"`
	Mids 	[]*Mids 	`json:"modules"`
}




func AddTenantRole(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start AddRole")
	tid := p.ByName("tid")
	if tid == "" {
		log.Error("Get tid Failed , ")
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	roleInfo := RoleInfo{}
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err := json.Unmarshal(body, &roleInfo); err != nil {
		log.Errorf("JsonError %s ", err)
		JsonReply("JsonError", nil, w)
		return
	}
	log.Infof("AddRoleInfo is %s", roleInfo)

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
	ctx = context.WithValue(ctx, "tid", tid)
	log.Info("url:", ctx.Value("url"))
	log.Info("tid:",tid)

	var mids []*pb.Module
	for _,v := range roleInfo.Mids {
		if v.Mid != 0 {
			mids = append(mids, &pb.Module{Mid:v.Mid, Operation:v.Operation})
		}
	}
	log.Info("mids:",mids)
	tidInt,_ := strconv.Atoi(tid)
	req := pb.AddRoleRequest{
		RoleName:roleInfo.RoleName,
		Description:roleInfo.Description,
		Mids:mids,
		Tid:int32(tidInt),
	}
	if req.RoleName == "" {
		log.Info("RoleName不能为空")
		JsonReply("RoleNameIsIncorrectOrEmpty", nil, w)
		return
	}
	if len(req.Mids) == 0 {
		log.Info("Mids不能为空")
		JsonReply("MidsIsIncorrectOrEmpty", nil, w)
		return
	}
	//todo 调用rpc
	_, err := rpc.RoleRpcClient().AddRole(ctx,&req)
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
	JsonReply("Successful",nil, w)
	return
}



func UpdateTenantRole(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start UpdateRoleModule")
	tid := p.ByName("tid")
	if tid == "" {
		log.Error("Get tid Failed , ")
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	roleInfo := RoleInfo{}
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err := json.Unmarshal(body, &roleInfo); err != nil {
		log.Errorf("JsonError %s ", err)
		JsonReply("JsonError", nil, w)
		return
	}
	log.Infof("AddRoleInfo is %s", roleInfo)

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
	ctx = context.WithValue(ctx, "tid", tid)
	log.Info("url:", ctx.Value("url"))
	log.Info("tid:",tid)

	//todo 数据处理
	log.Info("UpdateRoleModule:",roleInfo)
	var mids []*pb.Module
	for _,v := range roleInfo.Mids {
		if v.Mid != 0 {
			mids = append(mids, &pb.Module{Mid:v.Mid, Operation:v.Operation})
		}
	}
	log.Info("mids:",mids)
	req := pb.UpdateRoleRequest{Rid:roleInfo.Rid,Mids:mids,RoleName:roleInfo.RoleName,Description:roleInfo.Description}
	if req.Rid == 0 {
		log.Info("Rid不能为空")
		JsonReply("RidIsIncorrectOrEmpty", nil, w)
		return
	}
	if req.RoleName == "" {
		log.Info("RoleName不能为空")
		JsonReply("RoleNameIsIncorrectOrEmpty", nil, w)
		return
	}
	if len(req.Mids) == 0 {
		log.Info("Mids不能为空")
		JsonReply("MidsIsIncorrectOrEmpty", nil, w)
		return
	}
	//todo 调用rpc
	_, err := rpc.RoleRpcClient().UpdateRole(ctx,&req)
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
	JsonReply("Successful",nil, w)
	return
}
func DeleteTenantRole(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start DeleteRole")
	tid := p.ByName("tid")
	if tid == "" {
		log.Error("Get tid Failed , ")
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	Rid := p.ByName("rid")
	req := pb.DeleteRoleRequest{}
	if Rid != "" {
		advId, err := strconv.Atoi(Rid)
		if err != nil || advId == 0 {
			log.Error("strconv.Atoi(Rid) Failed,", err)
			JsonReply("RidIsIncorrectOrEmpty", nil, w)
			return
		}
		req.Rid = int32(advId)
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
	ctx = context.WithValue(ctx, "tid", tid)
	log.Info("url:", ctx.Value("url"))
	log.Info("tid:",tid)


	//todo 调用rpc
	_, err := rpc.RoleRpcClient().DeleteRole(ctx,&req)
	if err != nil {
		log.Errorf("调用rpc错误码,err is (%s)", err)
		s,ok := status.FromError(err)
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
		case CodeMap["RoleNameIsIncorrectOrEmpty"].Code:
			log.Info("RoleName为空或格式有误，", err)
			JsonReply("RoleNameIsIncorrectOrEmpty", nil, w)
			return
		case CodeMap["MidsIsIncorrectOrEmpty"].Code:
			log.Info("Mids为空或输入有误，", err)
			JsonReply("MidsIsIncorrectOrEmpty", nil, w)
			return
		case CodeMap["RidIsIncorrectOrEmpty"].Code:
			log.Info("Rid为空或输入有误，", err)
			JsonReply("RidIsIncorrectOrEmpty", nil, w)
			return
		default:
			log.Errorf("未知异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
	}
	JsonReply("Successful",nil, w)
	return
}

//获取所有角色
func GetTenantRoles(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start GetRoles")
	tid := p.ByName("tid")
	if tid == "" {
		log.Error("Get tid Failed , ")
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	count := r.FormValue("per_page")
	page := r.FormValue("page")
	req := &pb.GetRolesRequest{}
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

	log.Info("GetTenantRoles,","count:", count, ", page:", page)
	tidInt,_ := strconv.Atoi(tid)
	req.Tid = int32(tidInt)
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
	ctx = context.WithValue(ctx, "tid", tid)
	log.Info("url:", ctx.Value("url"))
	log.Info("tid:",tid)

	//todo 调用RPC
	reply,err := rpc.RoleRpcClient().GetRoles(ctx,req)
	if err != nil {
		log.Errorf("调用rpc错误码,err is (%s)", err)
		s,ok := status.FromError(err)
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
		case CodeMap["RoleNameIsIncorrectOrEmpty"].Code:
			log.Info("RoleName为空或格式有误，", err)
			JsonReply("RoleNameIsIncorrectOrEmpty", nil, w)
			return
		case CodeMap["MidsIsIncorrectOrEmpty"].Code:
			log.Info("Mids为空或输入有误，", err)
			JsonReply("MidsIsIncorrectOrEmpty", nil, w)
			return
		case CodeMap["RidIsIncorrectOrEmpty"].Code:
			log.Info("Rid为空或输入有误，", err)
			JsonReply("RidIsIncorrectOrEmpty", nil, w)
			return
		default:
			log.Errorf("未知异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
	}
	var roleInfos []*RoleInfo
	for _,v := range reply.RoleModules {
		if v.RoleInfo.Rid != 0 {
			roleInfo := RoleInfo{}
			roleInfo.Rid = v.RoleInfo.Rid
			roleInfo.RoleName = v.RoleInfo.RoleName
			roleInfo.Description = v.RoleInfo.Description
			for _,m := range v.Mids {
				if m.Mid != 0 {
					roleInfo.Mids = append(roleInfo.Mids, &Mids{Mid:m.Mid,ModuleName:m.ModuleName, Operation:m.Operation})
				}
			}
			roleInfos = append(roleInfos, &roleInfo)
		}
	}
	JsonReply("Successful", &RoleInfos{RoleInfo:roleInfos, TotalCount:reply.TotalCount}, w)
	return
}

func GetTenantRoleByRid(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start GetRoles")
	tid := p.ByName("tid")
	if tid == "" {
		log.Error("Get tid Failed , ")
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	ridInt64,err := strconv.ParseInt(p.ByName("rid"),10,64)
	if ridInt64 == 0 || err != nil{
		log.Error("Get did Failed , ")
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	rid := int32(ridInt64)

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
	ctx = context.WithValue(ctx, "tid", tid)
	log.Info("url:", ctx.Value("url"))
	log.Info("tid:",tid)

	//todo 调用RPC
	reply,err := rpc.RoleRpcClient().GetRoleByRid(ctx,&pb.GetRoleByRidRequest{Rid:rid})
	if err != nil {
		log.Errorf("调用rpc错误码,err is (%s)", err)
		s,ok := status.FromError(err)
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
		case CodeMap["RoleNameIsIncorrectOrEmpty"].Code:
			log.Info("RoleName为空或格式有误，", err)
			JsonReply("RoleNameIsIncorrectOrEmpty", nil, w)
			return
		case CodeMap["MidsIsIncorrectOrEmpty"].Code:
			log.Info("Mids为空或输入有误，", err)
			JsonReply("MidsIsIncorrectOrEmpty", nil, w)
			return
		case CodeMap["RidIsIncorrectOrEmpty"].Code:
			log.Info("Rid为空或输入有误，", err)
			JsonReply("RidIsIncorrectOrEmpty", nil, w)
			return
		default:
			log.Errorf("未知异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
	}

	roleInfo := &RoleInfo{}
		if reply.RoleModules.RoleInfo.Rid != 0 {
			roleInfo.Rid = reply.RoleModules.RoleInfo.Rid
			roleInfo.RoleName = reply.RoleModules.RoleInfo.RoleName
			roleInfo.Description = reply.RoleModules.RoleInfo.Description
			for _,m := range reply.RoleModules.Mids {
				if m.Mid != 0 {
					roleInfo.Mids = append(roleInfo.Mids, &Mids{Mid:m.Mid,ModuleName:m.ModuleName, Operation:m.Operation})
				}
			}
		}

	JsonReply("Successful", roleInfo, w)
	return
}

func GetDomainRoleByRid(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start GetRoles")
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	ridInt64,err := strconv.ParseInt(p.ByName("rid"),10,64)
	if ridInt64 == 0 || err != nil{
		log.Error("Get did Failed , ")
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	rid := int32(ridInt64)

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
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:",did)

	//todo 调用RPC
	reply,err := rpc.RoleRpcClient().GetRoleByRid(ctx,&pb.GetRoleByRidRequest{Rid:rid})
	if err != nil {
		log.Errorf("调用rpc错误码,err is (%s)", err)
		s,ok := status.FromError(err)
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
		case CodeMap["RoleNameIsIncorrectOrEmpty"].Code:
			log.Info("RoleName为空或格式有误，", err)
			JsonReply("RoleNameIsIncorrectOrEmpty", nil, w)
			return
		case CodeMap["MidsIsIncorrectOrEmpty"].Code:
			log.Info("Mids为空或输入有误，", err)
			JsonReply("MidsIsIncorrectOrEmpty", nil, w)
			return
		case CodeMap["RidIsIncorrectOrEmpty"].Code:
			log.Info("Rid为空或输入有误，", err)
			JsonReply("RidIsIncorrectOrEmpty", nil, w)
			return
		default:
			log.Errorf("未知异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
	}

	roleInfo := &RoleInfo{}
	if reply.RoleModules.RoleInfo.Rid != 0 {
		roleInfo.Rid = reply.RoleModules.RoleInfo.Rid
		roleInfo.RoleName = reply.RoleModules.RoleInfo.RoleName
		roleInfo.Description = reply.RoleModules.RoleInfo.Description
		for _,m := range reply.RoleModules.Mids {
			if m.Mid != 0 {
				roleInfo.Mids = append(roleInfo.Mids, &Mids{Mid:m.Mid,ModuleName:m.ModuleName, Operation:m.Operation})
			}
		}
	}

	JsonReply("Successful", roleInfo, w)
	return
}




func AddDomainRole(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start AddRole")
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	roleInfo := RoleInfo{}
	errCode := GetHttpData(r, "application/json;charset=UTF-8", &roleInfo)
	if errCode != 10000 {
		if errCode == 404 {
			ResCode(http.StatusNotFound, w)
			return
		}
		JsonReply("BodyIsIncorrectOrEmpty",nil, w)
		return
	}
	log.Info("AddRoleInfo:",roleInfo)

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
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:",did)

	var mids []*pb.Module
	for _,v := range roleInfo.Mids {
		if v.Mid != 0 {
			mids = append(mids, &pb.Module{Mid:v.Mid, Operation:v.Operation})
		}
	}
	log.Info("mids:",mids)
	didInt,_ := strconv.Atoi(did)
	req := pb.AddRoleRequest{
		RoleName:roleInfo.RoleName,
		Description:roleInfo.Description,
		Mids:mids,
		Did:int32(didInt),
	}
	if req.RoleName == "" {
		log.Info("RoleName不能为空")
		JsonReply("RoleNameIsIncorrectOrEmpty", nil, w)
		return
	}
	if len(req.Mids) == 0 {
		log.Info("Mids不能为空")
		JsonReply("MidsIsIncorrectOrEmpty", nil, w)
		return
	}
	//todo 调用rpc
	_, err := rpc.RoleRpcClient().AddDomainRole(ctx,&req)
	if err != nil {
		log.Errorf("调用rpc错误码,err is (%s)", err)
		s,ok := status.FromError(err)
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
		case CodeMap["RoleNameIsIncorrectOrEmpty"].Code:
			log.Info("RoleName为空或格式有误，", err)
			JsonReply("RoleNameIsIncorrectOrEmpty", nil, w)
			return
		case CodeMap["MidsIsIncorrectOrEmpty"].Code:
			log.Info("Mids为空或输入有误，", err)
			JsonReply("MidsIsIncorrectOrEmpty", nil, w)
			return
		case CodeMap["RidIsIncorrectOrEmpty"].Code:
			log.Info("Rid为空或输入有误，", err)
			JsonReply("RidIsIncorrectOrEmpty", nil, w)
			return
		default:
			log.Errorf("未知异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
	}
	JsonReply("Successful",nil, w)
	return
}

func UpdateDomainRole(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start UpdateRoleModule")
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	roleInfo := RoleInfo{}
	errCode := GetHttpData(r, "application/json;charset=UTF-8", &roleInfo)
	if errCode != 10000 {
		if errCode == 404 {
			ResCode(http.StatusNotFound, w)
			return
		}
		JsonReply("BodyIsIncorrectOrEmpty",nil, w)
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
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:",did)

	//todo 数据处理
	log.Info("UpdateRoleModule:",roleInfo)
	var mids []*pb.Module
	for _,v := range roleInfo.Mids {
		if v.Mid != 0 {
			mids = append(mids, &pb.Module{Mid:v.Mid, Operation:v.Operation})
		}
	}
	req := pb.UpdateRoleRequest{Rid:roleInfo.Rid,Mids:mids,RoleName:roleInfo.RoleName,Description:roleInfo.Description}
	if req.Rid == 0 {
		log.Info("Rid不能为空")
		JsonReply("RidIsIncorrectOrEmpty", nil, w)
		return
	}
	if req.RoleName == "" {
		log.Info("RoleName不能为空")
		JsonReply("RoleNameIsIncorrectOrEmpty", nil, w)
		return
	}
	if len(req.Mids) == 0 {
		log.Info("Mids不能为空")
		JsonReply("MidsIsIncorrectOrEmpty", nil, w)
		return
	}
	//todo 调用rpc
	_, err := rpc.RoleRpcClient().UpdateDomainRole(ctx,&req)
	if err != nil {
		log.Errorf("调用rpc错误码,err is (%s)", err)
		s,ok := status.FromError(err)
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
		case CodeMap["RoleNameIsIncorrectOrEmpty"].Code:
			log.Info("RoleName为空或格式有误，", err)
			JsonReply("RoleNameIsIncorrectOrEmpty", nil, w)
			return
		case CodeMap["MidsIsIncorrectOrEmpty"].Code:
			log.Info("Mids为空或输入有误，", err)
			JsonReply("MidsIsIncorrectOrEmpty", nil, w)
			return
		case CodeMap["RidIsIncorrectOrEmpty"].Code:
			log.Info("Rid为空或输入有误，", err)
			JsonReply("RidIsIncorrectOrEmpty", nil, w)
			return
		default:
			log.Errorf("未知异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
	}
	JsonReply("Successful",nil, w)
	return
}
func DeleteDomainRole(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start DeleteRole")
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	Rid := p.ByName("rid")
	req := pb.DeleteRoleRequest{}
	if Rid != "" {
		advId, err := strconv.Atoi(Rid)
		if err != nil || advId == 0 {
			log.Error("strconv.Atoi(Rid) Failed,", err)
			JsonReply("RidIsIncorrectOrEmpty", nil, w)
			return
		}
		req.Rid = int32(advId)
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
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:",did)


	//todo 调用rpc
	_, err := rpc.RoleRpcClient().DeleteDomainRole(ctx,&req)
	if err != nil {
		log.Errorf("调用rpc错误码,err is (%s)", err)
		s,ok := status.FromError(err)
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
		case CodeMap["RoleNameIsIncorrectOrEmpty"].Code:
			log.Info("RoleName为空或格式有误，", err)
			JsonReply("RoleNameIsIncorrectOrEmpty", nil, w)
			return
		case CodeMap["MidsIsIncorrectOrEmpty"].Code:
			log.Info("Mids为空或输入有误，", err)
			JsonReply("MidsIsIncorrectOrEmpty", nil, w)
			return
		case CodeMap["RidIsIncorrectOrEmpty"].Code:
			log.Info("Rid为空或输入有误，", err)
			JsonReply("RidIsIncorrectOrEmpty", nil, w)
			return
		default:
			log.Errorf("未知异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
	}
	JsonReply("Successful",nil, w)
	return
}

//获取所有角色
func GetDomainRoles(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start GetRoles")
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	count := r.FormValue("per_page")
	page := r.FormValue("page")
	req := &pb.GetRolesRequest{}
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
	log.Info("GetTenantRoles,","count:", count, ", page:", page)
	didInt,_ := strconv.Atoi(did)
	req.Did = int32(didInt)
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
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:",did)

	//todo 调用RPC
	reply,err := rpc.RoleRpcClient().GetDomainRoles(ctx,req)
	if err != nil {
		log.Errorf("调用rpc错误码,err is (%s)", err)
		s,ok := status.FromError(err)
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
		case CodeMap["RoleNameIsIncorrectOrEmpty"].Code:
			log.Info("RoleName为空或格式有误，", err)
			JsonReply("RoleNameIsIncorrectOrEmpty", nil, w)
			return
		case CodeMap["MidsIsIncorrectOrEmpty"].Code:
			log.Info("Mids为空或输入有误，", err)
			JsonReply("MidsIsIncorrectOrEmpty", nil, w)
			return
		case CodeMap["RidIsIncorrectOrEmpty"].Code:
			log.Info("Rid为空或输入有误，", err)
			JsonReply("RidIsIncorrectOrEmpty", nil, w)
			return
		default:
			log.Errorf("未知异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
	}
	var roleInfos []*RoleInfo

	for _,v := range reply.RoleModules {
		if v.RoleInfo.Rid != 0 {
			roleInfo := RoleInfo{}
			roleInfo.Rid = v.RoleInfo.Rid
			roleInfo.RoleName = v.RoleInfo.RoleName
			roleInfo.Description = v.RoleInfo.Description
			for _,m := range v.Mids {
				if m.Mid != 0 {
					roleInfo.Mids = append(roleInfo.Mids, &Mids{Mid:m.Mid, Operation:m.Operation})
				}
			}
			roleInfos = append(roleInfos, &roleInfo)
		}
	}
	JsonReply("Successful", &RoleInfos{RoleInfo:roleInfos,TotalCount:reply.TotalCount}, w)
	return
}

//获取所有模块
func GetTenantModules(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start GetRoles")
	tid := p.ByName("tid")
	if tid == "" {
		log.Error("Get tid Failed , ")
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
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
	ctx = context.WithValue(ctx, "tid", tid)
	log.Info("url:", ctx.Value("url"))
	log.Info("tid:",tid)

	//todo 调用RPC
	tidInt,_ := strconv.Atoi(tid)
	reply,err := rpc.RoleRpcClient().GetModuleByTid(ctx,&pb.GetModuleByTidRequest{Tid:int32(tidInt)})
	if err != nil {
		log.Errorf("调用rpc错误码,err is (%s)", err)
		s,ok := status.FromError(err)
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
		case CodeMap["RoleNameIsIncorrectOrEmpty"].Code:
			log.Info("RoleName为空或格式有误，", err)
			JsonReply("RoleNameIsIncorrectOrEmpty", nil, w)
			return
		case CodeMap["MidsIsIncorrectOrEmpty"].Code:
			log.Info("Mids为空或输入有误，", err)
			JsonReply("MidsIsIncorrectOrEmpty", nil, w)
			return
		case CodeMap["RidIsIncorrectOrEmpty"].Code:
			log.Info("Rid为空或输入有误，", err)
			JsonReply("RidIsIncorrectOrEmpty", nil, w)
			return
		default:
			log.Errorf("未知异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
	}
	var modules []*Module
	for _,v := range reply.Modules {
		if v.Mid != 0 {
			modules = append(modules, &Module{Mid:v.Mid, ModuleName:v.ModuleName})
		}
	}
	JsonReply("Successful", modules, w)
	return
}

func GetDomainModules(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start GetRoles")
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
	url := strings.Split(r.RequestURI,"?")[0]
	log.Infof("url is (%s),split url is %s",r.RequestURI,url)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "url", url)
	ctx = context.WithValue(ctx, "opt", opt)
	ctx = context.WithValue(ctx, "token", token)
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:",did)

	didInt,_ := strconv.Atoi(did)
	//todo 调用RPC
	reply,err := rpc.RoleRpcClient().GetModuleByDid(ctx,&pb.GetModuleByDidRequest{Did:int32(didInt)})
	if err != nil {
		log.Errorf("调用rpc错误码,err is (%s)", err)
		s,ok := status.FromError(err)
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
		case CodeMap["RoleNameIsIncorrectOrEmpty"].Code:
			log.Info("RoleName为空或格式有误，", err)
			JsonReply("RoleNameIsIncorrectOrEmpty", nil, w)
			return
		case CodeMap["MidsIsIncorrectOrEmpty"].Code:
			log.Info("Mids为空或输入有误，", err)
			JsonReply("MidsIsIncorrectOrEmpty", nil, w)
			return
		case CodeMap["RidIsIncorrectOrEmpty"].Code:
			log.Info("Rid为空或输入有误，", err)
			JsonReply("RidIsIncorrectOrEmpty", nil, w)
			return
		default:
			log.Errorf("未知异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
	}
	var modules []*Module
	for _,v := range reply.Modules {
		if v.Mid != 0 {
			modules = append(modules, &Module{Mid:v.Mid, ModuleName:v.ModuleName})
		}
	}
	JsonReply("Successful", modules, w)
	return
}

