package controller

import (
	"github.com/julienschmidt/httprouter"
)

func RouterMethod(router *httprouter.Router) {
	//邮件发送器
/*	router.POST(AddEmailSend,Auth(AddEmailSender))
	router.DELETE(DeleteEmailSend,Auth(DeleteEmailSender))
	router.PUT(UpdateEmailSend,Auth(UpdateEmailSender))
	router.GET(EmailSends,Auth(GetEmailSenders))
	router.GET(GetEmailSend,Auth(GetEmailSender))*/

	router.POST(AddEmailSend,AddEmailSender)
	router.DELETE(DeleteEmailSend,DeleteEmailSender)
	router.PUT(UpdateEmailSend,UpdateEmailSender)
	router.GET(EmailSends,GetEmailSenders)
	router.GET(GetEmailSend,GetEmailSender)

	//邮件模板
	router.POST(AddEmailTem,AddEmailTemplate)
	router.DELETE(DeleteEmailTem,DeleteEmailTemplate)
	router.PUT(UpdateEmailTem,UpdateEmailTemplate)
	router.GET(EmailTemplates,GetEmailTemplates)
	router.GET(GetEmailTem,GetEmailTemplate)
	//邮件通知
	router.POST(EmailNotice,NoticeByEmail)

}

//HTTP接口地址
const (
	//项目名，版本号
	version_1  = "/v1.0"
	// 邮件发送器
	AddEmailSend                 = version_1 + "/domains/:domainId/senders"
	EmailSends                   = version_1 + "/domains/:domainId/senders"
	GetEmailSend    	         = version_1 + "/domains/:domainId/senders/:senderId"
	DeleteEmailSend	             = version_1 + "/domains/:domainId/senders/:senderId"
	UpdateEmailSend              = version_1 + "/domains/:domainId/senders/:senderId"
	//配置邮件模板
	AddEmailTem                  = version_1 + "/tenants/:tenantId/templates"
	EmailTemplates               = version_1 + "/tenants/:tenantId/templates"
	GetEmailTem    	             = version_1 + "/tenants/:tenantId/templates/:etId"
	DeleteEmailTem	             = version_1 + "/tenants/:tenantId/templates/:etId"
	UpdateEmailTem               = version_1 + "/tenants/:tenantId/templates/:etId"
	//邮件通知
	EmailNotice                  = version_1 + "/notices"
)