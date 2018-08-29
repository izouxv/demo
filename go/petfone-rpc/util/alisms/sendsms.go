package alisms

import (
	log "github.com/cihub/seelog"
)

const (
	gatewayUrl = "http://dysmsapi.aliyuncs.com/"
	accessKeyId = "LTAI3pqHeu2ht5p5"//LTAI3pqHeu2ht5p5   LTAI0cnPz9TxFqni
	accessKeySecret = "N87mmOHNA3q2AoubRFRaYW32CqUFVp"//N87mmOHNA3q2AoubRFRaYW32CqUFVp UoBTysYVr5MiXHhNwFvNsy5Eo1IQzy
	signName = "小鲑"//小鲑  鹏联优思
	templateCode = "SMS_123673875"//SMS_123673875  SMS_115930310
	// 模板变量赋值，json格式
	templateParam1 = "{\"code\":\""
	templateParam2 = "\"}"
	phoneNumbers = "17600117962"
)

func SendSms(phone string, code string) map[string]interface{} {
	log.Info("SendSms-phone:",phone,",code:",code)
	smsClient := NewSmsClient(gatewayUrl)
	template := templateParam1+code+templateParam2
	result, err := smsClient.Execute(accessKeyId, accessKeySecret, phone, signName, templateCode, template)
	log.Info("SendSms-result:",result)
	if err != nil && result["Message"].(string) != "OKOK" {
		log.Error("发送验证码异常:",phone,result["Message"],result["Code"],", SendSms-error:",err.Error())
	}
	return result
}
