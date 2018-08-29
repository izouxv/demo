package thirdPartyServer

import (
	"encoding/base64"
	"net/url"
	"fmt"
	"strconv"
	"crypto/md5"
	"io"
	"net/http"
	"strings"
	"io/ioutil"
	"petfone-http/util"
	"encoding/json"
	log "github.com/cihub/seelog"
)

const (
	xfVoiceUrl = "http://api.xfyun.cn/v1/service/v1/iat"
	contentType = "application/x-www-form-urlencoded; charset=utf-8"
	params = `{"engine_type":"sms16k","aue":"raw"}`
	xParam = "eyJlbmdpbmVfdHlwZSI6InNtczE2ayIsImF1ZSI6InJhdyJ9"//对params进行base64编码
	appId = "5ac9c38a"
	apiKey = "279bb4d06fc5534522d496f41d5f3aec"
)

type XfResult struct {
	Code string `json:"code"`
	Data string `json:"data"`
	Desc string `json:"desc"`
	Sid string `json:"sid"`
}

/**
语音听写
 */
func VoiceToText(bs []byte) XfResult {
	//对bytes进行base64压缩
	audio := base64.StdEncoding.EncodeToString(bs)
	audioBase64 := url.QueryEscape(audio)
	log.Info("VoiceToText-audioBase64-len:", len(audioBase64))
	//计算头参数
	timeStr := strconv.FormatInt(util.GetNowTimeInt64(),10)
	//xParam := base64.StdEncoding.EncodeToString([]byte(param))
	w := md5.New()
	io.WriteString(w, apiKey+timeStr+xParam)
	checkSum := fmt.Sprintf("%x", w.Sum(nil))
	log.Info("VoiceToText-checkSum:",checkSum)
	//创建请求
	req, _ := http.NewRequest("POST",xfVoiceUrl, strings.NewReader("audio="+audioBase64))
	req.Header.Add("Content-Type",contentType)
	req.Header.Add("X-Appid",appId)
	req.Header.Add("X-CurTime",timeStr)
	req.Header.Add("X-Param",xParam)
	req.Header.Add("X-CheckSum",checkSum)
	res, err := httpClient.Do(req)
	if err != nil || res.StatusCode != 200 {
		log.Error("讯飞语音调用失败-err:", err,",StatusCode:",res.StatusCode)
	}
	result, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	var xfResult XfResult
	json.Unmarshal(result, &xfResult)
	log.Info("HttpSendMailPost-xfResult:",xfResult)
	return xfResult
}
