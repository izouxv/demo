package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/cihub/seelog"
)

const (
	NotificationIp string = "notification.radacat.com"
	//NotificationIp string = "127.0.0.1"
	MailModeRegister     string = "51"
	MailModeInvitation   string = "52"
	MailModeRestPwd      string = "53"
	MailModeNotification string = "54"
)

type Model string

const (
	//RegisterMail     Model = "51"
	InvitationMail   Model = "64"
	FindPasswordMail Model = "65"
	//NotificationMail Model = "54"
)

type MailResp struct {
	Code   string
	Msg    string
	Result interface{}
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
