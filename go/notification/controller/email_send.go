package controller

import (
	log "github.com/cihub/seelog"
	"net/http"
	"github.com/julienschmidt/httprouter"
	."notification/storage"
	"notification/common"
	."notification/common"
	"time"
	"notification/config"
)

//基于did配置邮件发送器
func AddEmailSender(res http.ResponseWriter, req *http.Request, params httprouter.Params)  {
	log.Info("AddEmailSender...")
	es:= &EmailSender{}
	if flag := common.GetJsonHttpData(req, es); flag {
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	did,err := common.StrToInt64(params.ByName("domainId"))
	if err != nil {
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	es.Did = did
	log.Debugf("参数 %#v",es)
	if  es.Did == 0 || es.EmailSender == "" || es.Password == "" || es.EmailSender == "" || es.Username == "" {
		log.Infof("did (%d) email (%s) username (%s) password (%s) 异常",es.Did,es.EmailSender,es.Username,es.Password)
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	count,_:=  es.GetEmailSenderCount(config.C.MySQL.DB)
	if count == 1{
		log.Infof("邮件发送器已经存在! count (%d)",count)
		common.JsonReply("AlreadyExists",nil,res)
		return
	}
	if err := es.CreateEmailSender(config.C.MySQL.DB);err != nil {
		log.Errorf("db CreateEmailSender err (%s)",err)
		common.JsonReply("Unimplemented",nil,res)
		return
	}
	go SendEmailBySender(es.SmtpServer,
		es.Username,
		es.Password,
		es.EmailSender,
		"测试发送器邮件",
		"This is test email ...",[]byte(`{"test":"test"}`))
	reply := ES{Id:es.Id,Did:es.Did,CreateTime:es.CreateTime.Unix(),UpdateTime:es.UpdateTime.Unix(),
	SmtpServer:es.SmtpServer,EmailSender:es.EmailSender,Username:es.Username,Password:es.Password}
	common.JsonReply("OK",reply,res)
	return
}

//基于did删除邮件发送器
func DeleteEmailSender(res http.ResponseWriter, req *http.Request, params httprouter.Params)  {
	log.Info("DeleteEmailSender...")
	es:= &EmailSender{}
	did,err := common.StrToInt64(params.ByName("domainId"))
	if err != nil {
		log.Infof("did (%d) 异常",es.Did)
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	es.Did = did
	esId,err := common.StrToInt64(params.ByName("senderId"))
	if es.Did == 0 || err != nil {
		log.Infof("esId (%d) 异常",esId)
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	es.Id = esId
	log.Debugf("参数 %#V",es)
	if err := es.DeleteEmailSender(config.C.MySQL.DB);err != nil {
		if err == ErrDoesNotExist{
			common.JsonReply("NotFound",nil,res)
			return
		}
		log.Errorf("db DeleteEmailSender err (%s)",err)
		common.JsonReply("Unimplemented",nil,res)
		return
	}
	common.JsonReply("OK",nil,res)
	return
}

//基于did修改邮件发送器
func UpdateEmailSender(res http.ResponseWriter, req *http.Request, params httprouter.Params)  {
	log.Info("UpdateEmailSender...")
	es:= &EmailSender{}
	if flag := common.GetJsonHttpData(req, es); flag {
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	did,err := common.StrToInt64(params.ByName("domainId"))
	if err != nil {
		log.Infof("did (%d) 异常",did)
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	esId,err := common.StrToInt64(params.ByName("senderId"))
	if err != nil {
		log.Infof("esId 异常(%s)",esId)
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	es.Id = esId
	es.Did = did
	es.UpdateTime = time.Now()
	log.Infof("http 接收参数:",es)
	if err := es.UpdateEmailSender(config.C.MySQL.DB);err != nil {
		log.Errorf("db UpdateEmailSender err (%s)",err)
		common.JsonReply("Unimplemented",nil,res)
		return
	}
	go SendEmailBySender(es.SmtpServer,
		es.Username,
		es.Password,
		es.EmailSender,
		"测试发送器邮件",
		"This is test email ...",
		[]byte(`{"test":"test"}`))
	reply := ES{Id:es.Id,Did:es.Did,CreateTime:es.CreateTime.Unix(),UpdateTime:es.UpdateTime.Unix(),
		SmtpServer:es.SmtpServer,EmailSender:es.EmailSender,Username:es.Username,Password:es.Password}
	common.JsonReply("OK",reply,res)
	return
}

//基于did获取邮件发送器列表
func GetEmailSenders(res http.ResponseWriter, req *http.Request, params httprouter.Params)  {
	log.Info("GetEmailSenders...")
	did,err := common.StrToInt64(params.ByName("domainId"))
	if err != nil {
		log.Infof("did (%d) 异常",did)
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	es:= &EmailSender{Did:did}
	count,page,orderBy := common.ReturnCountPageOrderBy(req)
	emailSenders,totalCount,err := es.GetEmailSenders(config.C.MySQL.DB,page,count,orderBy)
	if err != nil {
		log.Errorf("db GetEmailSenders err (%s)",err)
		common.JsonReply("Unimplemented",nil,res)
		return
	}
	el := EmailSenderList {}
	el.TotalCount = totalCount
	for _,v := range emailSenders{
		reply := ES{Id:v.Id,Did:v.Did,CreateTime:v.CreateTime.Unix(),UpdateTime:v.UpdateTime.Unix(),
			SmtpServer:v.SmtpServer,EmailSender:v.EmailSender,Username:v.Username,Password:v.Password}
		el.EmailSender = append(el.EmailSender,reply)
	}
	common.JsonReply("OK",el,res)
	return
}

//基于domainId获取邮件发送器
func GetEmailSender(res http.ResponseWriter, req *http.Request, params httprouter.Params)  {
	log.Info("GetEmailSender...")
	es:= &EmailSender{}
	did,err := common.StrToInt64(params.ByName("domainId"))
	if err != nil {
		log.Infof("did (%d) 异常",did)
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	esId,err := common.StrToInt64(params.ByName("senderId"))
	if err != nil {
		log.Infof("esId (%d) 异常",esId)
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	es.Did = did
	es.Id = esId
	if err := es.GetEmailSender(config.C.MySQL.DB);err != nil {
		log.Errorf("db GetEmailSender err (%s)",err)
		common.JsonReply("NotFound",nil,res)
		return
	}
	reply := ES{Id:es.Id,Did:es.Did,CreateTime:es.CreateTime.Unix(),UpdateTime:es.UpdateTime.Unix(),
		SmtpServer:es.SmtpServer,EmailSender:es.EmailSender,Username:es.Username,Password:es.Password}
	common.JsonReply("OK",reply,res)
	return
}

type  EmailSenderList struct {
	  EmailSender []ES        `json:"email_sender"`
	  TotalCount  int32      `json:"total_count"`
}

type ES struct{
	Id           int64     `json:"id"`
	Did          int64     `json:"did"`
	SmtpServer   string    `json:"smtp_server"`
	EmailSender  string    `json:"sender_email"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	CreateTime   int64     `json:"create_time"`
	UpdateTime  int64      `json:"update_time"`
}