package main

import (
	"bytes"
	"net/http"
	"io/ioutil"
	"fmt"
	"os"
	"encoding/base64"
	"net/url"
	"strconv"
	"time"
	"crypto/md5"
	"io"
	"strings"
)

func main() {
	FormPost()
}

//识别
func FormPost1() {
	paramJson := ``
	fmt.Println("FormPost----:",paramJson)
	//计算头参数
	timeStr := strconv.FormatInt(time.Now().Unix(),10)
	param := `{"engine_type":"sms16k","aue":"raw"}`
	xparam := base64.StdEncoding.EncodeToString([]byte(param))
	apikey := "279bb4d06fc5534522d496f41d5f3aec"
	w := md5.New()
	io.WriteString(w, apikey+timeStr+xparam)
	checkSum := fmt.Sprintf("%x", w.Sum(nil))
	//创建请求
	req, _ := http.NewRequest("POST","httpReq://openapi.xfyun.cn/v2/aiui", bytes.NewBuffer([]byte(param)))
	req.Header.Add("Content-Type","application/json; charset=utf-8")
	req.Header.Add("X-Appid","5ac9c38a")
	req.Header.Add("X-CurTime",timeStr)
	req.Header.Add("X-Param",xparam)
	req.Header.Add("X-CheckSum",checkSum)
	client := &http.Client{
		Timeout:time.Millisecond*1500,
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("FormPost-err:", err)
	}
	defer res.Body.Close()
	fmt.Println("FormPost-rest:", res)

	defer res.Body.Close()
	result, _ := ioutil.ReadAll(res.Body)
	fmt.Println("FormPost-rest:",string(result))
}


//表单语音识别
func FormPost() {
	ff, _ := os.Open("E:\\1.wav")
	defer ff.Close()
	sourcebuffer := make([]byte, 50000)
	n, _ := ff.Read(sourcebuffer)
	//base64压缩
	audio := base64.StdEncoding.EncodeToString(sourcebuffer[:n])
	audioBase64 := url.QueryEscape(audio)
	fmt.Println("FormPost----:",audioBase64)
	//计算头参数
	timeStr := strconv.FormatInt(time.Now().Unix(),10)
	param := `{"engine_type":"sms16k","aue":"raw"}`
	xparam := base64.StdEncoding.EncodeToString([]byte(param))
	apikey := "279bb4d06fc5534522d496f41d5f3aec"
	w := md5.New()
	io.WriteString(w, apikey+timeStr+xparam)
	checkSum := fmt.Sprintf("%x", w.Sum(nil))
	fmt.Println("FormPost----:", timeStr,param,xparam,checkSum)
	//创建请求
	req, _ := http.NewRequest("POST","http://api.xfyun.cn/v1/service/v1/iat", strings.NewReader("audio="+audioBase64))
	req.Header.Add("Content-Type","application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Add("X-Appid","5ac9c38a")
	req.Header.Add("X-CurTime",timeStr)
	req.Header.Add("X-Param",xparam)
	req.Header.Add("X-CheckSum",checkSum)
	client := &http.Client{
		Timeout:time.Millisecond*1500,
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("FormPost-err:", err)
	}
	fmt.Println("FormPost-rest:", res)
	result, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	fmt.Println("FormPost-rest:",string(result))
}


/**
发送json请求
 */
func ReqJson(address, method string, buff *bytes.Buffer) []byte {
	req, _ := http.NewRequest(method, address, buff)
	client := &http.Client{}
	req.Header.Set("Content-type", "application/json")
	res, _ := client.Do(req)
	if res.StatusCode == 200 {
		body, _ := ioutil.ReadAll(res.Body)
		return body
	} else {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
		return nil
	}
}
