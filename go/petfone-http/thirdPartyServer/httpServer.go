package thirdPartyServer

import (
	log "github.com/cihub/seelog"
	"net/http"
	"io/ioutil"
	"bytes"
	"petfone-http/core"
	"time"
	"petfone-http/result"
	. "petfone-http/util"
	"errors"
	"fmt"
)

var (
	httpClient = &http.Client{Timeout:time.Second*3}
	mailPath = "/v1.0/notices"
	httpType = "application/json;charset=UTF-8"
)

//发送反馈邮件
func FeedbackMail(to string, subject string, content string) {
	bodyStr := fmt.Sprint(
		`{"tem_id":6,"email_addr":"`,to,`",
		"send_data":{"subject":"`,"项圈建议与反馈-工单ID："+subject,`",
		"content":"`,content,`"}}`)
	log.Info("FeedbackMail result:",string(
		httpReq(core.ConstStr.NoticeServer+mailPath, "POST",httpType,bytes.NewBuffer([]byte(bodyStr)))))

}

//发送post请求
func httpReq(address,method ,contentType string,body *bytes.Buffer) []byte {
	log.Info("httpPost body:", body)
	req, _ := http.NewRequest(method,address, body)
	req.Header.Add("Content-Type", contentType)
	res, err := httpClient.Do(req)
	if err != nil {
		log.Error("httpPost err:", err)
		return nil
	}
	bs, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Error("httpPost err:", err)
		return nil
	}
	if res.StatusCode != 200 {
		log.Error("httpPost StatusCode:",res.StatusCode)
		return nil
	}
	return bs
}


//上传文件
func SendFileHttp(url,contentType string, buff *bytes.Buffer) (string, error) {
	bs := httpReq(url,"POST",contentType,buff)
	fileJson := &result.Myjson{}
	if err := Json.Unmarshal(bs, fileJson); err != nil {
		log.Error("SendFileHttp Unmarshal-err:", err)
		log.Error("SendFileHttp result:",string(bs))
		return "",err
	}
	fid := fileJson.Result.(map[string]interface{})["fid"].(string)
	if "d41d8cd98f00b204e9800998ecf8427e" == fid {
		log.Error("SendFileHttp err: 上传文件为空")
		return "",errors.New("上传文件为空")
	}
	return fid, nil
}

//下载请求
func GetFileHttp(url string) (*http.Response,int32) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		log.Error("GetFileHttp err:", err)
		log.Error("GetFileHttp Status:",res.Status+",url:"+url)
		return nil,10001
	}
	if res.Header.Get("content-type") != "application/octet-stream" {
		resBody, _ := ioutil.ReadAll(res.Body)
		log.Error("GetFileHttp 下载异常：",string(resBody))
		return res,21007
	}
	return res,10000
}

