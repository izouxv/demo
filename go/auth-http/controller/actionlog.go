package controller

import (
	pb "auth-http/api"
	"auth-http/rpc"
	."auth-http/util"
	"context"
	log "github.com/cihub/seelog"

	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
	"google.golang.org/grpc/status"
	"strings"
)

type ActionLogs struct {
	Logs []*Log				`json:"logs"`
	TotalCount int32 		`json:"total_count"`
}

type Log struct {
	Id        		int32 `json:"id"`
	ActionUsername	string `json:"action_username"`
	ActionTime		int64 `json:"action_time"`
	ActionType		int32 `json:"action_type"`
	ActionName		string `json:"action_name"`
	ActionObject	string `json:"action_object"`
}

type Username struct {
	ActionUsername	string `json:"username"`
}

//条件查找操作日志
func GetActionLogs(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("-----Start GetActionLogByType-----")
	alreq := pb.GetActionLogsRequest{}
	count := r.FormValue("per_page")
	page := r.FormValue("page")
	typeid := r.FormValue("type")
	username := r.FormValue("username")
	tid := r.FormValue("tid")
	did := r.FormValue("did")
	log.Infof("GetActionLogs,","count:", count, " ", ",page:", page)
	log.Infof("GetActionLogs,","tid:", tid, " ", ",did:", did)
	log.Infof("GetActionLogs,","username:", username, " ", ",typeid:", typeid)
	if typeid != "" {
		t, _ := strconv.Atoi(typeid)
		alreq.Type = int32(t)
	}
	if username != "" {
		alreq.Username = username
	}
	if count != "" {
		c, err := strconv.Atoi(count)
		if err != nil || c == 0 {
			log.Error("strconv.Atoi(count) Failed,", err)
			JsonReply("PerpageIsIncorrectOrEmpty", nil, w)
			return
		}
		alreq.Count = int32(c)
	}
	if page != "" {
		p, err := strconv.Atoi(page)
		if err != nil || p == 0 {
			log.Error("strconv.Atoi(page) Failed,", err)
			JsonReply("PageIsIncorrectOrEmpty", nil, w)
			return
		}
		alreq.Page = int32(p)
	}
	if tid != "" {
		tidInt, err := strconv.Atoi(tid)
		alTid := int32(tidInt)
		if err != nil || alTid == 0 {
			log.Error("strconv.ParseInt(tid) Failed,", err)
		}else {
			alreq.Tid = alTid
		}
	}
	if did != "" {
		didInt, err := strconv.Atoi(did)
		alDid := int32(didInt)
		if err != nil || alDid == 0 {
			log.Error("strconv.ParseInt(did) Failed,", err)
		}else {
			alreq.Did = alDid
		}
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
	log.Info("url:", ctx.Value("url"))

	log.Info("--alreq:",alreq)
	reply,err := rpc.ActionLogRpcClient().GetActionLogs(ctx,&alreq)
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
	var actionLogs ActionLogs
	var logs []*Log
	for _,v := range reply.ActionLogs {
		logs = append(logs, &Log{
			Id:v.Id,
			ActionUsername:v.ActionUsername,
			ActionTime:v.ActionTime,
			ActionName:v.ActionName,
			ActionType:v.ActionType,
			ActionObject:v.ActionObject,
		})
	}
	actionLogs.Logs = logs
	actionLogs.TotalCount = reply.TotalCount
	JsonReply("Successful", actionLogs, w)
	return
}

func GetAllActionLogsByTid(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}


func GetAllActionLogsByDid(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}