package controller

import (
	log "github.com/cihub/seelog"
	"net/http"
	"github.com/julienschmidt/httprouter"
	."notification/storage"
	"notification/common"
	"time"
	"notification/rpc"
	"context"
	"notification/api"
	"notification/config"
)

//基于id配置邮件模板
func AddEmailTemplate(res http.ResponseWriter, req *http.Request, params httprouter.Params)  {
	log.Info("AddEmailTemplate...")
	et := &EmailTemplate{}
	if flag := common.GetJsonHttpData(req, et); flag {
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	tid,err := common.StrToInt64(params.ByName("tenantId"))
	if err != nil{
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	et.Tid = tid
	response ,err := GetDidByTid(tid)
	if err != nil || response == nil{
		log.Errorf("调用auth-rpc失败 (%s)",err)
		common.JsonReply("Unimplemented",nil,res)
		return
	}
	et.Did = response.Did
	if et.Tid == 0 || et.Did == 0 ||  et.Subject == "" || et.Html == ""{
		log.Infof("tid (%d) did (%d) subject (%s) ",et.Tid,et.Did,et.Subject)
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	log.Debugf("emailTemplate 参数 %#v",et)
	if  err := et.CreateEmailTemplate(config.C.MySQL.DB);err != nil {
		log.Errorf("db CreateEmailTemplate err (%s)",err)
		common.JsonReply("Unimplemented",nil,res)
		return
	}
	reply := ET{Id:et.Id,Did:et.Did,Tid:et.Tid,CreateTime:et.CreateTime.Unix(),UpdateTime:et.UpdateTime.Unix(),
		Html:et.Html,Subject:et.Subject}
	    common.JsonReply("OK",reply,res)
	    return
}

//基于did删除邮件模板
func DeleteEmailTemplate(res http.ResponseWriter, req *http.Request, params httprouter.Params)  {
	log.Info("DeleteEmailTemplate...")
	tid,err := common.StrToInt64(params.ByName("tenantId"))
	if err != nil {
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	emailTemId,err := common.StrToInt64(params.ByName("etId"))
	if err != nil {
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	et := &EmailTemplate{Tid:tid,Id:emailTemId}
	log.Debugf("参数 tid (%d) id (%d)",et.Tid,et.Id)
	if  err := et.DeleteEmailTemplate(config.C.MySQL.DB);err != nil {
		log.Infof("db DeleteEmailTemplate err (%s)",err)
		common.JsonReply("Unimplemented",nil,res)
		return
	}
	common.JsonReply("OK",nil,res)
	return
}

//基于id修改邮件模板
func UpdateEmailTemplate(res http.ResponseWriter, req *http.Request, params httprouter.Params)  {
	log.Info("UpdateEmailTemplate...")
	et := &EmailTemplate{}
	if flag := common.GetJsonHttpData(req, et); flag {
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	tid,err := common.StrToInt64(params.ByName("tenantId"))
	if err != nil {
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	emailTemId,err := common.StrToInt64(params.ByName("etId"))
	if err != nil {
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	et.Id = emailTemId
	et.Tid = tid
	et.UpdateTime = time.Now()
	response ,err := GetDidByTid(tid)
	if err != nil || response == nil{
		log.Errorf("调用auth-rpc根据tid获取did失败 (%s)",err)
		common.JsonReply("Unimplemented",nil,res)
		return
	}
	et.Did = response.Did
	if et.Tid == 0 || et.Did == 0 || et.Subject == "" || et.Html == ""{
		log.Infof("tid (%d) did (%d) subject(%s) html (...)",et.Tid,et.Did,et.Subject)
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	log.Debugf("参数 %#v",et)
	if  err := et.UpdateEmailTemplate(config.C.MySQL.DB);err != nil {
		log.Errorf("db UpdateEmailTemplate err (%s)",err)
		common.JsonReply("Unimplemented",nil,res)
		return
	}
	reply := ET{Id:et.Id,Did:et.Did,Tid:et.Tid,CreateTime:et.CreateTime.Unix(),UpdateTime:et.UpdateTime.Unix(),
		Html:et.Html,Subject:et.Subject}
	common.JsonReply("OK",reply,res)
	return
}

//基于id获取邮件模板列表
func GetEmailTemplates(res http.ResponseWriter, req *http.Request, params httprouter.Params)  {
	log.Info("GetEmailTemplates...")
	tid,err := common.StrToInt64(params.ByName("tenantId"))
	if err != nil {
		log.Infof("参数异常 tid (%d)",tid)
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	et := &EmailTemplate{Tid:tid}
	count,page,orderBy := common.ReturnCountPageOrderBy(req)
	totalCount ,tems , err := et.GetEmailTemplates(config.C.MySQL.DB,page,count,orderBy)
	if err != nil {
		log.Errorf("db UpdateEmailTemplate err (%s)",err)
		common.JsonReply("Unimplemented",nil,res)
		return
	}
	etl := EmailTemList{}
	etl.TotalCount = totalCount
	for _,v := range tems{
		reply := ET{Id:v.Id,Did:v.Did,Tid:et.Tid,CreateTime:v.CreateTime.Unix(),UpdateTime:v.UpdateTime.Unix(),
			Html:v.Html,Subject:v.Subject}
		etl.EmailTem = append(etl.EmailTem,reply)
	}
	common.JsonReply("OK",etl,res)
	return
}

//基于tid获取邮件模板
func GetEmailTemplate(res http.ResponseWriter, req *http.Request, params httprouter.Params)  {
	log.Info("GetEmailTemplate...")
	tid,err := common.StrToInt64(params.ByName("tenantId"))
	if err != nil {
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	emailTemId,err := common.StrToInt64(params.ByName("etId"))
	if err != nil {
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	et := &EmailTemplate{Tid:tid,Id:emailTemId}
	if  err := et.GetEmailTemplate(config.C.MySQL.DB);err != nil {
		log.Errorf("db GetEmailTemplate err (%s)",err)
		common.JsonReply("Unimplemented",nil,res)
		return
	}
	reply := ET{Id:et.Id,Did:et.Did,Tid:et.Tid,CreateTime:et.CreateTime.Unix(),UpdateTime:et.UpdateTime.Unix(),
		Html:et.Html,Subject:et.Subject}
	common.JsonReply("OK",reply,res)
	return
}

func GetDidByTid (tid int64 ) (reply *api.GetDidByTidResponse,err error){
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	ctx = context.WithValue(ctx, "source", "notification")
	tenantReq := &api.GetDidByTidRequest{Tid:tid}
	reply ,err = rpc.TenantClient().GetDidByTid(ctx,tenantReq)
	return reply ,err
}

type  EmailTemList struct {
	EmailTem   []ET          `json:"email_template"`
	TotalCount  int32        `json:"total_count"`
}

type  ET struct {
	Id          int64     `json:"id"`
	Did         int64     `json:"did"`
	Tid         int64     `json:"tid"`
	Subject     string    `json:"subject"`
	Html        string    `json:"html"`
	CreateTime  int64     `json:"create_time"`
	UpdateTime  int64     `json:"update_time"`
}

