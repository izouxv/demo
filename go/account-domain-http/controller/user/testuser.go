package user


import (
	pb "account-domain-http/api/user"
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/cihub/seelog"
	"account-domain-http/util"
	"account-domain-http/rpc"
	"context"
)


type TestUser struct {
	Id         int32 `json:"id"`
	Tid        int64 `json:"tid"`
	UserName   string `json:"userName"`
	CreateTime int64 `json:"createTime"`
	UpdateTime int64 `json:"updateTime"`
}

type TestUserList struct {
	TotalCount int32      `json:"totalCount"`
	TestUser   []TestUser `json:"testUsers"`
}

func AddTestUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start AddTestUser")
	r.ParseForm()
	tid, err := util.StrToInt64(p.ByName("tid"))
	if err != nil {
		log.Debugf("tid参数异常tid (%d),error :(%s)",tid,err)
		util.JsonReply("Params_error", nil, w)
		return
	}
	testUser := &pb.TestUser{}
	if flag := util.GetJsonHttpData(r,testUser);flag {
		util.JsonReply("Body_is_incorrect_or_empty", nil, w)
		return
	}
	testUser.Tid = tid
	if testUser.Username == "" {
		log.Debug("用户名不能为空")
		util.JsonReply("Params_error", nil, w)
		return
	}
	reply, err := rpc.TestUserRpcClient().AddTestUser(context.Background(), &pb.AddTestUserReq{TestUser:testUser})
	if err != nil  {
		log.Error("调用account-domain-rpc AddTestUser failed: ", err)
		util.JsonReply("System_error", nil, w)
		return
	}
	switch reply.ErrorCode{
	case util.Successfull:
		util.JsonReply("Successful",reply.TestUser, w)
		return
	case util.Input_parameter_error:
		log.Debug("rpc返回参数异常")
		util.JsonReply("Params_error",nil, w)
		return
	case util.TestUser_is_exist:
		log.Debug("rpc返回测试账号已存在")
		util.JsonReply("TestUser_is_exist",nil, w)
		return
	default:
		log.Debug("rpc返回状态码:",reply.ErrorCode)
		util.JsonReply("System_error", nil, w)
		return
	}
}

func DelTestUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start DelTestUser")
	r.ParseForm()
	tid, err := util.StrToInt64(p.ByName("tid"))
	if err != nil {
		log.Debugf("tid参数异常tid (%d),error :(%s)",tid,err)
		util.JsonReply("Params_error", nil, w)
		return
	}
	id ,err := util.StringToInt32(p.ByName("id"))
	if err != nil || id == 0 {
		log.Debugf("id参数异常 (%d)",id)
		util.JsonReply("Params_error", nil, w)
		return
	}
	reply, err := rpc.TestUserRpcClient().DelTestUser(context.Background(), &pb.DelTestUserReq{Tid:tid,Id:id})
	if err != nil || reply.ErrorCode != util.Successfull {
		log.Error("调用account-domain-rpc DelTestUser failed: ", err)
		util.JsonReply("System_error", nil, w)
		return
	}
	log.Info("rpc返回状态码:", reply.ErrorCode)
	util.JsonReply("Successful",nil, w)
	return
}

func GetTestUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start GetTestUsers")
	r.ParseForm()
	tid, err := util.StrToInt64(p.ByName("tid"))
	if err != nil {
		log.Debugf("tid参数异常tid (%d),error :(%s)",tid,err)
		util.JsonReply("Params_error", nil, w)
		return
	}
	count ,page:= util.ReturnCountPage(r)
	log.Infof("http参数 tid (%d) page (%d) count (%d)",tid,page,count)
	reply, err := rpc.TestUserRpcClient().GetTestUsers(context.Background(), &pb.GetTestUsersReq{Tid:tid,Page:page,Count:count})
	if err != nil || reply.ErrorCode != util.Successfull {
		log.Error("调用account-domain-rpc GetTestUsers failed: ", err)
		util.JsonReply("System_error", nil, w)
		return
	}
	log.Info("reply:",reply.TestUser)
	tul := TestUserList{TotalCount:reply.TotalCount}
	for _,v := range reply.TestUser{
		tu := TestUser{Id:v.Id,Tid:v.Tid,UserName:v.Username,CreateTime:v.CreateTime,UpdateTime:v.UpdateTime}
		tul.TestUser = append(tul.TestUser,tu)
	}
	log.Info("rpc返回状态码:", reply.ErrorCode)
	util.JsonReply("Successful",tul, w)
	return
}

func UpdateTestUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start UpdateTestUser")
	r.ParseForm()
	tid, err := util.StrToInt64(p.ByName("tid"))
	if err != nil {
		log.Debugf("tid参数异常tid (%d),error :(%s)",tid,err)
		util.JsonReply("Params_error", nil, w)
		return
	}
	id ,err := util.StringToInt32(p.ByName("id"))
	if err != nil || id == 0 {
		log.Debugf("id参数异常 (%d)",id)
		util.JsonReply("Params_error", nil, w)
		return
	}
	testUser := &pb.TestUser{}
	if flag := util.GetJsonHttpData(r,testUser);flag {
		util.JsonReply("Body_is_incorrect_or_empty", nil, w)
		return
	}
	testUser.Tid = tid
	testUser.Id = id
	if testUser.Username == "" {
		log.Debug("用户名不能为空")
		util.JsonReply("Params_error", nil, w)
		return
	}
	reply, err := rpc.TestUserRpcClient().PutTestUser(context.Background(), &pb.PutTestUserReq{TestUser:testUser})
	if err != nil  {
		log.Error("调用account-domain-rpc PutTestUser failed: ", err)
		util.JsonReply("System_error", nil, w)
		return
	}
	switch reply.ErrorCode{
	case util.Successfull:
		util.JsonReply("Successful",reply.TestUser, w)
		return
	case util.Input_parameter_error:
		log.Debug("rpc返回参数异常")
		util.JsonReply("Params_error",nil, w)
		return
	case  util.TestUser_is_exist:
		log.Debug("rpc返回测试账号已存在")
		util.JsonReply("TestUser_is_exist",nil, w)
		return
	default:
		log.Debug("rpc返回状态码:",reply.ErrorCode)
		util.JsonReply("System_error", nil, w)
		return
	}
	return
}

func GetTestUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start GetTestUser")
	r.ParseForm()
	tid, err := util.StrToInt64(p.ByName("tid"))
	if err != nil {
		log.Debugf("tid参数异常tid (%d),error :(%s)",tid,err)
		util.JsonReply("Params_error", nil, w)
		return
	}
	id ,err := util.StringToInt32(p.ByName("id"))
	if err != nil || id == 0 {
		log.Debugf("id参数异常 (%d)",id)
		util.JsonReply("Params_error", nil, w)
		return
	}
	reply, err := rpc.TestUserRpcClient().GetTestUser(context.Background(), &pb.GetTestUserReq{Id:id,Tid:tid})
	if err != nil {
		log.Error("调用account-domain-rpc GetTestUser failed: ", err)
		util.JsonReply("System_error", nil, w)
		return
	}
	switch  reply.ErrorCode {
	case util.TestUser_not_exist:
		util.JsonReply("Not_exist", nil, w)
		return
	case util.Successfull:
		util.JsonReply("Successful",reply.TestUser, w)
		return
	default :
		log.Info("rpc 返回状态码:",reply.ErrorCode)
		util.JsonReply("System_error", nil, w)
		return
	}
}







