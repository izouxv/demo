package controller

import (
	pb "account-domain-http/api/feedback"
	"account-domain-http/rpc"
	"account-domain-http/util"
	"context"
	log "github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
	"io/ioutil"
	"encoding/json"
)

type Feedback struct {
	Id          int32  `json:"id"`
	Tid         int64  `json:"tid"`
	Type        int32  `json:"type"`
	Description string `json:"description"`
	DeviceInfo  string `json:"deviceInfo,omitempty"`
	AppInfo     string `json:"appInfo,omitempty"`
	UserInfo    string `json:"userInfo,omitempty"`
	MobileInfo  string `json:"mobileInfo,omitempty"`
	ExtendInfo  string `json:"extendInfo,omitempty"`
	Files       string `json:"files"`
	Contact     string `json:"contact,omitempty"`
	CreateTime  int64  `json:"createTime"`
	UpdateTime  int64  `json:"updateTime"`
}

type FeedbackList struct {
	Feedback   []Feedback `json:"feedBack"`
	TotalCount int32      `json:"totalCount"`
}

const  Del_file_url  = "http://file.radacat.com:88/v1.0/file"

//增加反馈  TODO http由chatting提供
func AddFeedback(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start AddFeedback")
	r.ParseForm()
	feedback := &pb.Feedback{}
	errCode := util.GetHttpData(r, "application/json;charset=UTF-8", &feedback)
	if errCode != util.Successfull {
		if errCode == 404 {
			util.ResCode(http.StatusNotFound, w)
			return
		}
		util.JsonReply("Body_is_incorrect_or_empty", nil, w)
		return
	}
	log.Info("http参数,feedback :",feedback)
	if feedback.Description == "" {
		util.JsonReply("Description_is_incorrect_or_empty", nil, w)
		return
	}
	reply, err := rpc.FeedbackRpcClient().AddFeedback(context.Background(), &pb.AddFeedbackRequest{Description:feedback.Description})
	if err != nil {
		log.Error("调用account-domain-rpc AddFeedback failed :", err)
		util.JsonReply("System_error", nil, w)
		return
	}
	log.Info("rpc返回状态码:", reply.ErrorCode)
	util.JsonReply("Successful",nil, w)
	return
}

//基于租户id增加反馈
func AddFeedbackBaseTenant(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start AddFeedbackBaseTenant")
	r.ParseForm()
	tid, err := util.StrToInt64(p.ByName("tid"))
	if err != nil {
		log.Infof("tid (%d),error :(%s)",tid,err)
		util.JsonReply("TenantId_is_incorrect_or_empty", nil, w)
		return
	}
	feedback := &pb.Feedback{}
	errCode := util.GetHttpData(r, "application/json;charset=UTF-8", &feedback)
	if errCode != util.Successfull {
		if errCode == 404 {
			util.ResCode(http.StatusNotFound, w)
			return
		}
		util.JsonReply("Body_is_incorrect_or_empty", nil, w)
		return
	}
	if feedback.Description == "" {
		util.JsonReply("Description_is_incorrect_or_empty", nil, w)
		return
	}
	if feedback.Type == 0 {
		util.JsonReply("Type_is_incorrect_or_empty", nil, w)
		return
	}
	reply, err := rpc.FeedbackRpcClient().AddFeedbackBaseTenant(context.Background(), &pb.AddFeedbackBaseTenantRequest{Tid:tid,Type:feedback.Type ,Description:feedback.Description,Files:feedback.Files})
	if err != nil {
		log.Error("调用account-domain-rpc AddFeedbackBaseTenant failed: ", err)
		util.JsonReply("System_error", nil, w)
		return
	}
	if reply.ErrorCode != util.Successfull{
		util.JsonReply("System_error", nil, w)
		return
	}
	log.Info("rpc返回状态码:", reply.ErrorCode)
	util.JsonReply("Successful",nil, w)
	return
}

//基于租户获取反馈列表
func GetFeedbacksBaseTenant(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start GetFeedbacksBaseTenant")
	r.ParseForm()
	countStr := r.FormValue("per_page")
	pageStr := r.FormValue("page")
	tid, err := util.StrToInt64(p.ByName("tid"))
	if err != nil {
		log.Infof("tid (%d),error :(%s)",tid,err)
		util.JsonReply("TenantId_is_incorrect_or_empty", nil, w)
		return
	}
	bugType, err := strconv.ParseInt(r.FormValue("type"), 10, 64)
	if err != nil {
		log.Infof("strconv.Atoi(type) Failed,err is %s ", err)
		util.JsonReply("Type_is_incorrect_or_empty", nil, w)
		return
	}
	count, err := util.StrToInt64(countStr)
	if err != nil || count == 0  {
		log.Info("Per_page不能为空", count)
		util.JsonReply("Per_page_is_incorrect_or_empty", nil, w)
		return
	}
	page, err := util.StrToInt64(pageStr)
	if err != nil || page == 0  {
		log.Infof("strconv.Atoi(page) Failed (%s) ", err)
		util.JsonReply("Page_is_incorrect_or_empty", nil, w)
		return
	}
	log.Infof("tid (%d) bugtype (%d) page(%d) per_page (%d)",tid,bugType, page, count)
	if bugType == 0 {
		reply, err := rpc.FeedbackRpcClient().GetFeedbacks(context.Background(), &pb.GetFeedbacksRequest{Tid: tid, Page: int32(page), Count: int32(count)})
		if err != nil {
			log.Error("调用account-domain-rpc GetFeedbacks failed:", err)
			util.JsonReply("System_error", nil, w)
			return
		}
		log.Infof("rpc返回状态码 (%d) ", reply.ErrorCode)
		feedbackList := FeedbackList{TotalCount: reply.TotalCount}
		for _, v := range reply.FeedBack {
			feedBack := Feedback{
				Id:          v.Id,
				Tid:         v.Tid,
				Description: v.Description,
				DeviceInfo:  v.DeviceInfo,
				Files:       v.Files,
				UserInfo:    v.UserInfo,
				MobileInfo:  v.MobileInfo,
				ExtendInfo:  v.ExtendInfo,
				AppInfo:     v.AppInfo,
				Contact:     v.Contact,
				Type:        v.Type,
				CreateTime:  v.CreateTime,
				UpdateTime:  v.UpdateTime,
			}
			feedbackList.Feedback = append(feedbackList.Feedback, feedBack)
		}
		if feedbackList.Feedback == nil {
			util.JsonReply("No_Feedback_Can_Be_Find", nil, w)
			return
		}
		log.Infof("feedback totalCount :(%d)",feedbackList.TotalCount)
		util.JsonReply("Successful", feedbackList, w)
		return
	}else{
		reply, err := rpc.FeedbackRpcClient().GetFeedbacksByType(context.Background(), &pb.GetFeedbacksByTypeRequest{Tid: tid,Type:int32(bugType), Page: int32(page), Count: int32(count)})
		if err != nil {
			log.Error("调用account-domain-rpc GetFeedbacksByType failed ", err)
			util.JsonReply("System_error", nil, w)
			return
		}
		log.Infof("rpc返回状态码 (%d)", reply.ErrorCode)
		feedbackList := FeedbackList{TotalCount: reply.TotalCount}
		for _, v := range reply.FeedBack {
			feedBack := Feedback{
				Id:          v.Id,
				Tid:         v.Tid,
				Description: v.Description,
				DeviceInfo:  v.DeviceInfo,
				Files:       v.Files,
				UserInfo:    v.UserInfo,
				MobileInfo:  v.MobileInfo,
				ExtendInfo:  v.ExtendInfo,
				AppInfo:     v.AppInfo,
				Contact:     v.Contact,
				Type:        v.Type,
				CreateTime:  v.CreateTime,
				UpdateTime:  v.UpdateTime,
			}
			feedbackList.Feedback = append(feedbackList.Feedback, feedBack)
		}
		if feedbackList.Feedback == nil {
			util.JsonReply("No_Feedback_Can_Be_Find", nil, w)
			return
		}
		log.Infof("feedback totalCount: (%d)",feedbackList.TotalCount)
		util.JsonReply("Successful", feedbackList, w)
		return
	}
	return
}

//基于租户单个反馈详情
func GetFeedbackBaseTenant(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start GetFeedbackBaseTenant")
	//提取form表单信息
	r.ParseForm()
	tid, err := util.StrToInt64(p.ByName("tid"))
	if err != nil {
		log.Infof("tid (%d),error :(%s)",tid,err)
		util.JsonReply("TenantId_is_incorrect_or_empty", nil, w)
		return
	}
	id, err := strconv.ParseInt(p.ByName("id"), 10, 64)
	if err != nil {
		log.Infof("strconv.Atoi(id) Failed,err is %s ", err)
		util.JsonReply("TenantId_is_incorrect_or_empty", nil, w)
		return
	}
	log.Infof("参数 id:(%d) tid (%d) ", id,tid)
	reply, err := rpc.FeedbackRpcClient().GetFeedback(context.Background(), &pb.GetFeedbackRequest{Tid:tid,Id: int32(id)})
	if err != nil {
		log.Errorf("调用account-domain-rpc  GetFeedback failed ", err)
		util.JsonReply("System_error", nil, w)
		return
	}
	if reply.ErrorCode == util.No_Feedback_Can_Be_Find{
		util.JsonReply("No_Feedback_Can_Be_Find", nil, w)
		return
	}
	if reply.ErrorCode != util.Successfull{
		util.JsonReply("System_error", nil, w)
		return
	}
	log.Info("rpc返回状态码:", reply.ErrorCode)
	util.JsonReply("Successful", reply.FeedBack, w)
	return
}

func DelFeedbackBaseTenant(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start DelFeedbackBaseTenant")
	//提取form表单信息
	r.ParseForm()
	tid, err := util.StrToInt64(p.ByName("tid"))
	if err != nil {
		log.Infof("tid (%d),error :(%s)",tid,err)
		util.JsonReply("TenantId_is_incorrect_or_empty", nil, w)
		return
	}
	delFd := &pb.DelFeedbackRequest{}
	if flag := util.GetJsonHttpData(r,delFd);flag{
		util.JsonReply("Body_is_incorrect_or_empty", nil, w)
		return
	}
	delFd.Tid = tid
	log.Info("删除的工单ids:",delFd.Ids,"租户id:",tid)
	res, err := rpc.FeedbackRpcClient().BatchFeedback(context.Background(), &pb.BatchFeedbackRequest{Tid:delFd.Tid,Ids: delFd.Ids})
	if err != nil {
		log.Errorf("批量获取要删除的工单信息有误", err)
		util.JsonReply("System_error", nil, w)
		return
	}
	var md5s = make([]string,0)
	if res.ErrorCode == util.Successfull {
		for _,v :=  range res.FeedBack{
			md5 := GetMd5(v.Files,v.ExtendInfo)
			for _,v := range md5{
				md5s = append(md5s,v)
			}
		}
	}
	log.Info("要删除的工单中文件MD5:",md5s)
	reply, err := rpc.FeedbackRpcClient().DelFeedback(context.Background(), &pb.DelFeedbackRequest{Tid:delFd.Tid,Ids: delFd.Ids})
	if err != nil {
		log.Errorf("调用account-domain-rpc  DelFeedback failed ", err)
		util.JsonReply("System_error", nil, w)
		return
	}
	log.Info("rpc删除工单返回状态码:", reply.ErrorCode)
	switch  reply.ErrorCode  {
	case util.Input_parameter_error:
		util.JsonReply("Params_error", nil, w)
		return
	case util.Successfull:
		if len(md5s) != 0 {
			go DelFile(md5s)
		}
		util.JsonReply("Successful",nil, w)
		return
	default:
		util.JsonReply("System_error", nil, w)
		return
	}
}

func GetMd5(file  ...string)  (strs []string) {
	md5 := make([]string,0)
	for _,v := range file{
		if v != "" && strings.Contains(v, "file/") {
			files := strings.Split(v, "file/")
			for i := 1; i < len(files); i = i + 1 {
				value := files[i]
				md5 = append(md5, value[0:32])
			}
		}
	}
	return md5
}

func DelFile(md5 []string){
	log.Info("批量删除工单中图片及日志文件...")
	client := &http.Client{}
	b ,_:= json.Marshal(md5)
	req, err := http.NewRequest("DELETE",Del_file_url,strings.NewReader(string(b)))
    if err != nil {
    	log.Error(err)
    	return
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return
	}
	var fileres FileRes
	err = json.Unmarshal(body,&fileres)
	log.Infof("批量删除文件返回(%s)",string(body))
}


type FileRes struct {
	Code     int32   `json::"code"`
    Msg      string   `json::"msg"`
    Result   interface{}   `json:"result"`
}



