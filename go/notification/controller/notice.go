package controller

import (
	log "github.com/cihub/seelog"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	."notification/storage"
	"notification/common"
	"notification/config"
)

type EmailNotices struct{
	Tid        int64      `json:"tid"`
	Did        int64      `json:"did"`
	TemId      int64      `json:"tem_id"`
	EmailAddr string      `json:"email_addr"`
	SendData  interface{} `json:"send_data"`
}

func NoticeByEmail(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Infof("start EmailSend...")
	en := &EmailNotices{}
	if flag := common.GetJsonHttpData(req, en); flag {
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	if en.TemId == 0 {
		log.Infof("temId (%d)",en.Tid,en.TemId)
		common.JsonReply("InvalidArgument",nil,res)
		return
	}
	et := &EmailTemplate{Tid: en.Tid, Id: en.TemId}
	if err := et.GetEmailTemplateById(config.C.MySQL.DB); err != nil {
		log.Errorf("获取邮件模板异常 error (%s) id (%d) ", err, et.Tid)
		return
	}
	es := &EmailSender{Did:et.Did}
	if err := es.GetEmailSenderByDid(config.C.MySQL.DB); err != nil{
		log.Errorf("获取邮件发送器错误 error (%s) did (%d) ", err, es.Did)
		common.JsonReply("Email_failed",nil,res)
		return
	}
	byteData, err := json.Marshal(en.SendData)
	if err != nil || len(byteData) == 0 || byteData == nil {
		log.Infof("数据marshal错误 error (%s) sendData (%s) ", err, en.SendData)
		common.JsonReply("PARAMS_ERR",nil,res)
		return
	}
	common.JsonReply("OK",nil,res)
	go  common.SendEmailBySender(es.SmtpServer, es.Username, es.Password, en.EmailAddr, et.Html, et.Subject, byteData)
	return

}
