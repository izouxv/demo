package user

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/cihub/seelog"
	"account-domain-http/util"
	"account-domain-http/rpc"
	"account-domain-http/api/user/pb"
	api"account-domain-http/api/user"
	"golang.org/x/net/context"
	"time"
	"encoding/base64"
)

const
(
	LLRedisTestUser = "AQ:testuser:"
	PLRedisTestUser = "Ag:testuser:"
	XGRedisTestUser = "BA:testuser:"
)


type MUser struct {
	Uid         int32  `json:"uid"`
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	LoginState  int32  `json:"loginState"`
	State       int32  `json:"state"`
	RegTime     int64  `json:"regTime"`
	RegIp       string  `json:"regIp"`
	RegAddr     string  `json:"regAddr"`
	LoginTime   int64   `json:"loginTime"`
	QuitTime    int64   `json:"quitTime"`
	NewIP      string   `json:"newIP"`
	NewAddr    string   `json:"newAddr"`
	Token      string   `json:"token"`
	DevInfo    string   `json:"devInfo"`
}

type MUserList  struct {
	Accounts   []MUser     `json:"accounts"`
	TotalCount  int32      `json:"totalCount"`
}

//分页查询用户
func GetPageAccounts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("GetPageAccounts")
	pageStr  := r.FormValue("page")
	countStr := r.FormValue("per_page")
	tid, err := util.StrToInt64(p.ByName("tid"))
	_,ok := util.TidToSource[tid]
	if err != nil  ||  !ok {
		log.Infof("strconv.Atoi(tenantid) Failed,err is (%s) ,tid (%d) to source (%s)  not success  ", err,tid)
		util.JsonReply("TenantId_is_incorrect_or_empty", nil, w)
		return
	}
	page, err := util.StringToInt32(pageStr)
	count, err := util.StringToInt32(countStr)
	if err != nil || page == 0 || count == 0 {
		log.Infof("参数有误 page (%d) count (%d)  err (%s) ", pageStr,countStr,err)
		util.JsonReply("Params_error", nil, w)
		return
	}
	var  accountsRe *pb.PageSsoReply
	switch tid {
	case 100001:
		log.Infof("调用account-rpc的用户管理 tid (%d) source (%s)", tid, util.TidToSource[tid])
		sTime := time.Now()
		accountsRe, err = rpc.MUserRpcClient().GetPageSsoInfos(context.Background(), &pb.PageRequest{Source: util.TidToSource[tid], Page: page, Count: count, Order: "id"})
		if err != nil {
			log.Error("调用account-rpc 错误 :", err)
			util.JsonReply("System_error", nil, w)
			return
		}
		log.Infof("调用account-rpc用时",time.Now().Sub(sTime))
		log.Infof("account-rpc code (%d)", accountsRe.Code)
	case 100002:
		log.Infof("调用petfone-rpc的用户管理 tid (%d) source (%s)", tid, util.TidToSource[tid])
		sTime := time.Now()
		accountsRe, err = rpc.MUserPetRpcClient().GetPageSsoInfos(context.Background(), &pb.PageRequest{Source: util.TidToSource[tid], Page: page, Count: count, Order: "id"})
		if err != nil {
			log.Error("调用account-rpc 错误 :", err)
			util.JsonReply("System_error", nil, w)
			return
		}
		log.Infof("调用petfone-rpc用时",time.Now().Sub(sTime))
		log.Infof("petfone-rpc code (%d)", accountsRe.Code)
	case 100003:
		log.Infof("调用account-rpc的用户管理 tid (%d) source (%s)", tid, util.TidToSource[tid])
		sTime := time.Now()
		accountsRe, err = rpc.MUserRpcClient().GetPageSsoInfos(context.Background(), &pb.PageRequest{Source: util.TidToSource[tid], Page: page, Count: count, Order: "id"})
		if err != nil {
			log.Error("调用account-rpc 错误 :", err)
			util.JsonReply("System_error", nil, w)
			return
		}
		log.Infof("调用account-rpc用时",time.Now().Sub(sTime))
		log.Infof("account-rpc code (%d)", accountsRe.Code)
	default:
		log.Infof("参数 tid (%d)有误",tid)
	}
	if len(accountsRe.MSsos) == 0 {
		util.JsonReply("User_does_not_exist_in_did", nil, w)
		return
	}
	mUserList := MUserList{TotalCount: accountsRe.TotalCount}
		for _, v := range accountsRe.MSsos{
			users := MUser{
				Uid:           v.Uid,           Username:      v.Username,
				Nickname:      v.Nickname,      LoginState:    v.LoginState,
				State:         v.State,         RegTime:       v.RegTime,
				RegIp:         v.RegIP,         RegAddr:       v.RegAddr,
				LoginTime:     v.LoginTime,
				NewAddr:       v.NewAddr,       QuitTime:      v.QuitTime,
				NewIP:         v.NewIP,         Token:         v.Token,
				DevInfo:       v.DevInfo,
			}
			mUserList.Accounts = append(mUserList.Accounts, users)
		}
	util.JsonReply("Successful", mUserList, w)
	return
}

func GetAccountByUid (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("GetAccount  by  uid")
	tid, err := util.StrToInt64(p.ByName("tid"))
	_,ok := util.TidToSource[tid]
	if err != nil  ||  !ok {
		log.Infof("strconv.Atoi(tid) Failed,err is (%s) ,tid (%d) to source (%s)  not success  ", err,tid)
		util.JsonReply("TenantId_is_incorrect_or_empty", nil, w)
		return
	}
	uid, err := util.StringToInt32(p.ByName("uid"))
	if err != nil {
		log.Infof("get uid (%d) err (%s) ",uid,err)
		util.JsonReply("Params_error", nil, w)
		return
	}
	var  account *pb.MSsoInfo
	switch tid {
	case 100001:
		log.Infof("参数 tid (%d) uid (%d)", tid, uid)
		account, err = rpc.MUserRpcClient().SearchSsoInfo(context.Background(), &pb.MSsoInfo{Source: util.TidToSource[tid], Uid: uid})
		if err != nil || account.Code == util.System_error {
			log.Errorf("调用account-rpc错误 (%s) 状态码 (%d) source (%s)", err, account.Code,util.TidToSource[tid])
			util.JsonReply("System_error", nil, w)
			return
		}
		log.Infof("account-rpc code (%d)", account.Code)
	case 100002:
		log.Infof("参数 tid (%d) uid (%d) source (%s)", tid, uid,util.TidToSource[tid])
		account, err = rpc.MUserPetRpcClient().SearchSsoInfo(context.Background(), &pb.MSsoInfo{Source: util.TidToSource[tid], Uid: uid})
		if err != nil || account.Code == util.System_error {
			log.Errorf("调用petfone-rpc错误 (%s) 状态码 (%d) ", err, account.Code)
			util.JsonReply("System_error", nil, w)
			return
		}
		log.Infof("petfone-rpc code (%d)", account.Code)
	case 100003:
		log.Infof("参数 tid (%d) uid (%d)", tid, uid)
		account, err = rpc.MUserRpcClient().SearchSsoInfo(context.Background(), &pb.MSsoInfo{Source: util.TidToSource[tid], Uid: uid})
		if err != nil || account.Code == util.System_error {
			log.Errorf("调用account-rpc错误 (%s) 状态码 (%d) source (%s)", err, account.Code,util.TidToSource[tid])
			util.JsonReply("System_error", nil, w)
			return
		}
		log.Infof("account-rpc code (%d)", account.Code)
	default:
		log.Infof("参数 tid (%d)有误",tid)
	}
	if account.Uid == 0 {
		util.JsonReply("User_does_not_exist_in_did", nil, w)
		return
	}
	user := MUser{
		Uid:           account.Uid,           Username:      account.Username,
		Nickname:      account.Nickname,      LoginState:    account.LoginState,
		State:         account.State,         RegIp:         account.RegIP,
		RegTime:       account.RegTime,       RegAddr:       account.RegAddr,
		QuitTime:      account.QuitTime,      LoginTime:     account.LoginTime,
		NewIP:        account.NewIP,          NewAddr:       account.NewAddr,
		Token:         account.Token,         DevInfo:        account.DevInfo,
	}
	util.JsonReply("Successful", user, w)
	return
}

//todo 此接口是给测试专用,没有页面。暂时只有接口
func DeleteTestUser (w http.ResponseWriter, r *http.Request, p httprouter.Params){
	log.Info("Start DeleteTestUser")
	source  :=  r.Header.Get("source")
	userName := p.ByName("username")
	log.Infof("source:%s username:%s",source,userName)
	if source == "" || userName == ""{
		util.JsonReply("Params_error", nil, w)
		return
	}
	decodeBytes, err := base64.StdEncoding.DecodeString(source)
	if err != nil {
		log.Infof("source(%s)解码错误:%s",source,err)
		util.JsonReply("Source_is_incorrect_or_empty", nil, w)
		return
	}
    v,ok := util.DecodeSourceToTid[decodeBytes[0]]
    if !ok {
		util.JsonReply("Source_is_incorrect_or_empty", nil, w)
		return
	}
	var replyUser *pb.MSsoReply
	var errs error
	reply ,err := rpc.TestUserRpcClient().GetUserByUsername(context.Background(), &api.GetUserByUsernameReq{Username:userName,Tid:v})
	if err != nil {
		log.Error("调用domain-rpc GetUserByUsername有误: ",err)
		util.JsonReply("System_error", nil, w)
		return
	}
	switch reply.ErrorCode {
	case  util.TestUser_not_exist:
		log.Debug("不是测试账号,不能删除")
		util.JsonReply("User_does_not_exist_in_db", nil, w)
		return
	case util.Successfull:
		log.Debug("是测试账号,允许删除")
		switch decodeBytes[0] {
		case 1,4:
			log.Info("蓝涟或小鲑的测试账号,调用account-rpc")
			replyUser ,errs = rpc.MUserRpcClient().DeleteAccount(context.Background(), &pb.MSsoInfo{Username:userName,Source:source})
		case 2:
			log.Info("鹏联测试账号,调用petfone-rpc")
			replyUser, errs  = rpc.MUserPetRpcClient().DeleteAccount(context.Background(), &pb.MSsoInfo{Username:userName,Source:source})
		default:
			log.Infof("source有误 (%s)",source)
			util.JsonReply("Source_is_incorrect_or_empty", nil, w)
			return
		}
		if errs != nil {
			log.Errorf("调用rpc错误:%s",err)
			util.JsonReply("System_error", nil, w)
			return
		}
		    switch  replyUser.Code {
		    case 10000:
			   util.JsonReply("Successful", nil, w)
			   return
		    case 10001:
				util.JsonReply("System_error", nil, w)
				return
		   case 33001:
				util.JsonReply("Params_error", nil, w)
				return
		case 33002:
				util.JsonReply("User_does_not_exist", nil, w)
				return
		case 33010:
				util.JsonReply("Source_is_incorrect_or_empty", nil, w)
				return
		}
	case util.System_error:
		util.JsonReply("System_error", nil, w)
		return
	default:
		log.Info("domain-rpc GetUserByUsername 返回状态码:",reply.ErrorCode)
		util.JsonReply("System_error", nil, w)
		return
	}

	/*client := RedisClient(Persistence).Get()
	if client.Err() != nil {
		log.Errorf("RedisClient Error, connection refused ,%s", client.Err())
		util.JsonReply("System_error", nil, w)
		return
	}
	defer client.Close()
	var reply *pb.MSsoReply
	var errs error
	switch  decodeBytes[0] {
	case  1:
		user ,err := client.Do("get",LLRedisTestUser + userName)
		if err != nil || user == nil {
			util.JsonReply("User_does_not_exist_in_redis", nil, w)
			return
		}
		reply ,errs  = rpc.MUserRpcClient().DeleteAccount(context.Background(), &pb.MSsoInfo{Username:userName,Source:source})
	case  2:
		user ,err := client.Do("get",PLRedisTestUser + userName)
		if err != nil || user == nil {
			util.JsonReply("User_does_not_exist_in_redis", nil, w)
			return
		}
		reply, errs = rpc.MUserPetRpcClient().DeleteAccount(context.Background(), &pb.MSsoInfo{Username:userName,Source:source})
	case 4:
		user ,err := client.Do("get",XGRedisTestUser + userName)
		if err != nil || user == nil {
			util.JsonReply("User_does_not_exist_in_redis", nil, w)
			return
		}
		reply ,errs = rpc.MUserRpcClient().DeleteAccount(context.Background(), &pb.MSsoInfo{Username:userName,Source:source})
	default:
		log.Infof("source有误 (%s)",source)
		util.JsonReply("Source_is_incorrect_or_empty", nil, w)
		return
	}
	if errs != nil {
		log.Errorf("调用rpc错误:%s",errs)
		util.JsonReply("System_error", nil, w)
		return
	}
	switch  reply.Code {
	case 10000:
		util.JsonReply("Successful", nil, w)
		return
	case 10001:
		util.JsonReply("System_error", nil, w)
		return
	case 33001:
		util.JsonReply("Params_error", nil, w)
		return
	case 33002:
		util.JsonReply("User_does_not_exist", nil, w)
		return
	case 33010:
		util.JsonReply("Source_is_incorrect_or_empty", nil, w)
		return
	}*/
	return
}

