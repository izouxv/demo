package thirdPartyServer

import (
	"testing"
	"fmt"
	"github.com/disintegration/imaging"
	"strings"
	"os"
	"bytes"
	"mime/multipart"
	"io"
	"petfone-http/util"
)


var (
	fileServer = "http://file.penslink.com:88/v1.0/file"
	noticeServer = "http://192.168.1.51:7023"
)

func TestFeedbackMail(t *testing.T) {
	temId := int32(6)
	emailAddr := "wangdy@radacat.com"
	bodyStr := fmt.Sprint(
		`{"tem_id":`,util.Int32ToStr(temId),
		`,"email_addr":"`,emailAddr,`",
		"send_data":{"subject":"`,"项圈建议与反馈-工单ID："+"123456",`",
		"content":"`,
		`<pre>联系方式：`,util.PrintStr("4324"),`</pre>`,
		`<pre>描述信息：`,util.PrintStr("4324"),`</pre>`,
		`<pre>手机信息：`,util.PrintStr("4324"),`</pre>`,
		`<pre>应用信息：`,util.PrintStr("4324"),`</pre>`,
		`<pre>设备信息：`,util.PrintStr("4324"),`</pre>`,
		`<pre>用户信息：`,util.PrintStr("4324"),`</pre>`,
		`<pre>文件信息：`,util.PrintStr("4324"),`</pre>`,
		`<pre>扩展信息：`,util.PrintStr("4324"),`</pre>`,
		`"}}`)
	buffer := bytes.NewBuffer([]byte(bodyStr))
	fmt.Println("result:",string(httpReq(noticeServer+"/v1.0/notices","POST",httpType, buffer)))
}

func TestPostFileHttp(t *testing.T) {
	file,err := os.Open("../temp/d787132ea1c62bbb311e05f166f85416.png")
	fmt.Println("Name:",file.Name())
	buff := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(buff)
	fileWriter, err := bodyWriter.CreateFormFile("file", "05f166f85416.png")
	if err != nil {
		fmt.Println("TestPostFileHttp CreateFormFile err:",err)
		return
	}
	copySize,err :=io.Copy(fileWriter, file)
	fmt.Println(copySize)
	if err != nil || copySize == 0 {
		fmt.Println("TestPostFileHttp Copy err:",err,copySize)
		return
	}
	bodyWriter.Close()
	fid,err := SendFileHttp(fileServer,"1.png",buff)
	fmt.Println("TestPostFileHttp:", fid,err)
}

func TestGetFileHttp(t *testing.T) {
	res,code := GetFileHttp(fileServer+"da4e8a3db87d848f3ce7e739beac058e")
	defer res.Body.Close()
	if code != 10000 {
		fmt.Println("err1:",code)
		return
	}
	image,err := imaging.Decode(res.Body)
	//调整成为适合的尺寸
	dst := imaging.Fit(image,250,200,imaging.Lanczos)
	//将结果图像保存为JPEG
	err = imaging.Save(dst,strings.Split(res.Header.Get("content-disposition"),"\"")[1])
	if err != nil {
		fmt.Println("failed to save image:", err)
	}
	//将结果图像保存为JPEG
	err = imaging.Save(dst,strings.Split(res.Header.Get("content-disposition"),"\"")[1])
	if err != nil {
		fmt.Println("failed to save image:", err)
	}
}