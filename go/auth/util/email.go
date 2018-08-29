package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/cihub/seelog"
	"bytes"
)

const (
	NotificationIp string = "notification.radacat.com"
	//NotificationIp string = "192.168.1.178"
)

type Model string

const (
	FindPasswordMail Model = "65"
	
	CotxInvitationMail   Model = "64"
	CotxFindPasswordMail Model = "65"

	RadacatInvitationMail   Model = "67"
	RadacatFindPasswordMail Model = "68"
	PensLinkInvitationMail   Model = "69"
	PensLinkFindPasswordMail Model = "70"
)

type MailResp struct {
	Code   string
	Msg    string
	Result interface{}
}


type MailInfo struct {
	TemId 		int32 		`json:"tem_id"`
	EmailAddr 	string 		`json:"email_addr"`
	SendData 	SendData 	`json:"send_data"`
}


type SendData struct {
	To 			string 	`json:"to"`
	Nickname 	string 	`json:"nickname"`
	Token 		string 	`json:"token"`
	Company 	string 	`json:"company"`
	Url 		string 	`json:"url"`
}




/*SendMail customModel填写
SendMail(user.Username,[]string{"token:"+token,"nickname:"+user.Nickname},util.FindPasswordMail)
*/
func SendMail(to string, customModel []string, model Model) {
	emailRequest := make(url.Values)
	emailRequest["tos"] = []string{to}
	emailRequest["model"] = []string{string(model)}
	emailRequest["customModel"] = customModel
	HttpSendMailPostForm(emailRequest)
}

func HttpSendMailPostForm(req url.Values) bool {
	resp, err := http.PostForm("http://"+NotificationIp+":8008/notification/v1.0/mailmodel", req)

	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	var rest MailResp
	json.Unmarshal(result, &rest)
	fmt.Println(rest)
	if rest.Code != "10000" {
		log.Errorf("发送邮件失败%#v", rest)
		return false
	}
	log.Infof("发送邮件成功,%#v", rest)
	return true
}

func HttpSendMailPost(requestInfo []byte) bool {
	resp, err := http.Post("http://118.190.152.145:88/v1.0/notices","application/json",bytes.NewBuffer(requestInfo))
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	var rest MailResp
	json.Unmarshal(result, &rest)
	fmt.Println(rest)
	if rest.Code != "10000" {
		log.Errorf("发送邮件失败%#v", rest)
		return false
	}
	log.Infof("发送邮件成功,%#v", rest)
	return true
}