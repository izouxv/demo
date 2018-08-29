package util

import (
	"testing"
	"fmt"
	"time"
	"bytes"
)

func TestIPToString(t *testing.T) {
	//fmt.Println(IPToString(3232235787))
	a := "2018-08-21 10:55:01"
	aa,_ := time.Parse("2006-01-02 15:04:05",a)
	b := "2018-08-21 11:07:59"
	bb,_ := time.Parse("2006-01-02 15:04:05",b)
	fmt.Println(aa.Minute())
	fmt.Println(bb.Minute())
	fmt.Println(CountMinute(aa,bb))
}

func TestSendSms(t *testing.T) {
	bodyStr := StrAdd(`{"phone":"`,"17600117962",`","code":"`,"123456",`","tid":100001}`)
	buffer := bytes.NewBuffer([]byte(bodyStr))
	httpReq(address+"/v1.0/sms","POST", buffer)
}

func TestIPToAddr(t *testing.T) {
	//fmt.Println(IPToAddr("203.69.66.102"))
	//fmt.Println(1529053370-1529052340)
	//times := time.Unix(1529022340,0)
	//fmt.Println(strings.Split(times.Format("2006-01-02 15:04:05")," ")[1][:5])
	chanIp := make(chan string)
	var ip string
	ip = "61.149.7.166"
	fmt.Println(time.Now())
	go IPToAddr(chanIp,ip)
	select {
	case times := <-time.After(time.Second*2):
		fmt.Println("times:",times)
	case addr := <-chanIp:
		fmt.Println("addr:",addr)
	}
	time.Sleep(time.Second*3)
}

var (
	address = "http://192.168.1.51:7023"
)

func TestSendMail(t *testing.T) {
	tem_id := int32(9)
	email_addr := "wangdy@radacat.com"
	nikename := "王东阳"
	token := "123456789123456789"
	company := "北京鹏联优思科技"
	domainUrl := "http://www.penslink.com"
	bodyStr := StrAdd(`{"tem_id":`,Int32ToStr(tem_id),`,"email_addr":"`,email_addr,`",
			"send_data":{"to":"`,email_addr,`","nickname":"`,nikename,`","token":"`,token,`",
				"company":"`,company,`","url":"`,domainUrl,`"}}`)
	buffer := bytes.NewBuffer([]byte(bodyStr))
	fmt.Println("result:",string(httpReq(address+"/v1.0/notices","POST", buffer)))
}
