package main

import (
	"github.com/go-gomail/gomail"
	"fmt"
)
func main() {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "2930026384@qq.com" /*"发件人地址"*/, "发件人") // 发件人
	m.SetHeader("To", m.FormatAddress("wangdy@radacat.com", "收件人")) // 收件人
	//m.SetHeader("Cc", m.FormatAddress("wangdy@radacat.com", "收件人")) //抄送
	//m.SetHeader("Bcc", m.FormatAddress("wangdy@radacat.com", "收件人")) //暗送
	m.SetHeader("Subject", "测试")     // 主题
	//m.SetBody("text/html",xxxxx ") // 可以放html..还有其他的
	m.SetBody("text/html;chartset=UTF-8",
		"<pre>联系方式：q \r\n 描述信息：q \r\n 手机信息：q \r\n 应用信息：q \r\n 设备信息：q \r\n 用户信息：q\r\n 文件信息：q \r\n 扩展信息：q</pre>") // 正文

	m.Attach("我是附件")  //添加附件

	d := gomail.NewDialer("smtp.qq.com", 465, "2930026384@qq.com", "gvjvblygnadudeac") // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("发送失败", err)
		return
	}
	fmt.Println("done.发送成功")
}