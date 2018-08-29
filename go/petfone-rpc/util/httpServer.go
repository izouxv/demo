package util

import (
	"net/http"
	"bytes"
	"time"
	"io/ioutil"
	"petfone-rpc/core"
	log "github.com/cihub/seelog"
)

var (
	httpClient = &http.Client{Timeout:time.Second*3}
	mailPath = "/v1.0/notices"
	smsPath = "/v1.0/sms"
)

//发送邮件
func SendMail(to, nikename,token string) {
	bodyStr := StrAdd(`{"tem_id":9,"email_addr":"`,to,`",
			"send_data":{"to":"`,to,`","nickname":"`,nikename,`","token":"`,token,`",
			"company":"北京鹏联优思科技","url":"http://www.penslink.com"}}`)
	buffer := bytes.NewBuffer([]byte(bodyStr))
	log.Info("SendMail res:",string(httpReq(core.ConstStr.NoticeServer+mailPath, "POST", buffer)))
}

//发送短信
func SendSms(phone, code string) {
	buffer := bytes.NewBuffer([]byte(StrAdd(`{"phone":"`,phone,`","code":"`,code,`","tid":100002}`)))
	log.Info("SendSms res:",string(httpReq(core.ConstStr.NoticeServer+smsPath, "POST", buffer)))
}

func httpReq(address,method string,body *bytes.Buffer) []byte {
	log.Info("httpPost body:", body)
	req, err := http.NewRequest(method,address, body)
	if err != nil {
		log.Error("httpPost err:", err)
		return nil
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	res, err := httpClient.Do(req)
	if err != nil {
		log.Error("httpPost err:", err)
		return nil
	}
	result, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Error("httpPost err:", err)
		return nil
	}
	if res.StatusCode != 200 {
		log.Error("httpPost StatusCode:",res.StatusCode)
		return nil
	}
	return result
}

