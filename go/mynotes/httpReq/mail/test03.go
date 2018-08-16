package main

import (
	"log"
	"net/smtp"
	"crypto/tls"
	"net"
	"fmt"
)


func main() {
	host :="mail.radacat.com"
	port := 465
	from  := "admin@radacat.com"
	pwd := "RadacatMail2017"  // 这里填你的授权码
	toEmail := "wangdy@radacat.com"  // 目标地址

	header := make(map[string]string)
	header["From"] = "admin"
	header["To"] = toEmail
	header["Subject"] = "邮件标题11111"
	header["Content-Type"] = "text/html;chartset=UTF-8"
	body  := "<pre>联系方式：q \r\n 描述信息：q \r\n 手机信息：q \r\n 应用信息：q \r\n 设备信息：q \r\n 用户信息：q\r\n 文件信息：q \r\n 扩展信息：q</pre>"

	message := "123"
	for k,v :=range header{
		message  += fmt.Sprintf("%s:%s\r\n",k,v)
	}
	message +="\r\n"+body
	auth :=smtp.PlainAuth("", from, pwd, host)
	client, err := dial(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Println("Create smpt client error:", err)
		return
	}
	defer client.Close()
	err = sendMailUsingTLS(client,auth, from, []string{"wangdy@radacat.com"}, []byte(message))
	if err !=nil{
		panic(err)
	}

}
//return a smtp client
func dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, &tls.Config{InsecureSkipVerify:true})
	if err != nil {
		log.Panicln("Dialing Error:", err)
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}


func sendMailUsingTLS(c *smtp.Client,auth smtp.Auth, from string, to []string, msg []byte) (err error) {
	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Println("Error during AUTH", err)
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		fmt.Println("email :",addr)
		if err = c.Rcpt(addr); err != nil {
			fmt.Println("rcpt err :",err)
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