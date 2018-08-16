package received

import (
	"encoding/json"
	"fmt"
	log "github.com/cihub/seelog"
	. "notification/common"
	"notification/handler"
	"notification/handler/multihandler"
	"notification/storage"
	"notification/config"
)

func HandleReceivedPayloadsEmail() {
	for _, v := range multihandler.Multihandler {
		go handleReceivedPayloadEmail(v.ReceivedChanEmail())
	}
}

func handleReceivedPayloadEmail(plChan chan handler.PayloadEmail) {
	log.Info("start handleReceivedPayloadEmail")
	for pl := range plChan {
		go func(pl handler.PayloadEmail) {
			if err := HandlerEmailPayload(pl); err != nil {
				errStr := fmt.Sprint(
					"handle received payload is :", pl,
					"error: %s", err,
				)
				log.Error(errStr)
			}
		}(pl)
	}
}

func HandlerEmailPayload(pl handler.PayloadEmail) error{
	log.Infof("receive  PayloadEmail  %#v ", pl)
	et ,err := getEmailTemplate(pl.Tid,pl.ALarmTemId)
	if err != nil {
		log.Errorf("获取邮件模板异常 error  (%s)", err)
		return err
	}
	es ,err := getEmailSender(et.Did)
	if err != nil {
		log.Errorf("获取邮件发送器异常 error  (%s)", err)
		return err
	}
	byteData ,err := json.Marshal(pl)
	if err != nil {
		log.Errorf("mqtt邮件通知中收到的数据marshal有误 err (%s)", err)
		return err
	}
	go SendEmailBySender(es.SmtpServer, es.Username, es.Password,pl.AlarmEmail,et.Html,et.Subject, byteData)
	return nil
}

//根据tid,alarmId查找邮件模板
func getEmailTemplate (tid,temId int64) (emailTem  storage.EmailTemplate,err error ){
	et := storage.EmailTemplate{Tid:tid,Id:temId}
	err = et.GetEmailTemplateById(config.C.MySQL.DB)
	if err != nil {
		log.Errorf("getEmailTemplate exception !! error (%s)", err)
		return et,err
	}
	return et ,nil
}

//根据did查找邮件发送器
func getEmailSender (did int64) (emailSender *storage.EmailSender,err error ){
	en := &storage.EmailSender{Did:did}
	err = en.GetEmailSenderByDid(config.C.MySQL.DB)
	if err != nil {
		log.Errorf("getEmailSender exception !! error (%s)", err)
		return nil,err
	}
	return en ,nil
}

