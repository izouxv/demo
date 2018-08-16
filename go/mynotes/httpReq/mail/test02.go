package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

/*解析html文件并填充数据*/

type User struct {
	Nickname string
	Token    string
}

func main() {
	host := "smtp.qq.com:465"
	user  := "1161115315@qq.com"
	pwd := "qycbvzlgxljzjhid"          // 这里填你的授权码
	to := []string{"wangdy@radacat.com"} //  目标地址

	fmt.Println("SEND EMAIL ...")
	str := `<pre>联系方式：q \r\n  描述信息：q \r\n 手机信息：q \r\n 应用信息：q \r\n 设备信息：q \r\n 用户信息：q\r\n 文件信息：q \r\n 扩展信息：q</pre>`
	err := SendToMail(user,pwd,host,to[0],"123144",str,"html")
	fmt.Println(err)
}

func SendToMail(user, pwd, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	fmt.Println("hp[0] :", hp[0])
	//鉴权无需端口
	auth := smtp.PlainAuth("", user, pwd, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	fmt.Println("host :", host)
	err := smtp.SendMail(host, auth, user, send_to, msg)
	fmt.Println("host :", msg)
	return err

}