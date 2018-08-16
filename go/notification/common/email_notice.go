package common

import (
	"bytes"
	"net/smtp"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"github.com/alecthomas/template"
	log "github.com/cihub/seelog"
	"crypto/tls"
	"net"
)

//通过发送器发送邮件
func SendEmailBySender(serverAddr, username ,password string,receiveAddr ,tem ,subject string,a []byte)(err error )  {
	bods := new(bytes.Buffer)
	var emails = make([]string ,0)
	con := strings.Split(serverAddr,":")
	auth := smtp.PlainAuth("",username,password,con[0])
	email := strings.Split(receiveAddr,",")
	for _,v := range email{
		emails = append(emails,v)
	}
	alarm := make(map[string]interface{}, 0)
	err = json.Unmarshal(a,&alarm)
	if err != nil {
		fmt.Println(err)
	}
	tmpl, err := template.New("template").Parse(tem)
	if err != nil {
		fmt.Println("tmpl err : ",err)
		return
	}
	log.Info("发邮件的信息 :",alarm)
	for k,v := range alarm{
		str ,err := json.Marshal(v)
		if err != nil {
			log.Infof("err (%s)",err)
		}
		a,err  := strconv.Unquote( string(str))
		if err != nil {
			tmpl.New(k).Parse(string(str))
		}
		tmpl.New(k).Parse(a)
	}
	err = tmpl.Execute(bods, nil)
	if err != nil {
		log.Errorf("邮件模板错误！err (%s)",err)
		return err
	}
	content_type := "Content-Type:text/html;charset=utf-8"
	msg := []byte("To:" + receiveAddr  + "\r\nFrom:"+ username + "\r\nSubject:"+ subject + "\r\n" + content_type +  "\r\n\r\n"+ bods.String())
	err  = sendMailUsingTLS(serverAddr, auth, username, emails, msg)
	if err != nil{
		log.Errorf("发送邮件错误!邮箱号(%s) err (%s)",email,err)
		return err
	}else{
		log.Infof("邮件发送成功")
		return
	}
	return
}


//return a smtp client
func dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, &tls.Config{InsecureSkipVerify:true})
	if err != nil {
		log.Info("Dialing Error:", err)
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}


func sendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {

	//create smtp client
	c, err := dial(addr)
	if err != nil {
		log.Info("Create smpt client error:", err)
		return err
	}
	defer c.Close()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Info("Error during AUTH", err)
				return err
			}
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}