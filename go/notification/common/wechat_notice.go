package common

import (
	log "github.com/cihub/seelog"
	"encoding/json"
	"bytes"
	"net/http"
	"github.com/alecthomas/template"
	"io/ioutil"
)

type RedisNotice struct {
	Source  string `json:"source"`
	Message string `json:"message"`
	Beat    struct {
		Version  string  `json:"version"`
		Name     string  `json:"name"`
		HostName string  `json:"hostname"`
	}  `json:"beat"`
	Timestamp    string `json:"@timestamp"`
	Metadata     struct {
		Beat    string `json:"beat"`
		Type    string `json:"type"`
		Version string `json:"version"`
	}  `json:"@metadata"`
	Offset    int    `json:"offSet"`
}

type Access struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type Sent struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	Invaliduser  string `json:"invaliduser"`
	Invalidparty string `json:"invalidparty"`
	Invalidtag   string `json:"invalidtag"`
}

type Text struct {
	Touser  string `json:"touser"`
	Toparty string `json:"toparty"`
	Totag   string `json:"totag"`
	Msgtype string `json:"msgtype"`
	Agentid int    `json:"agentid"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	Safe int `json:"safe"`
}

func WechatNotice (m RedisNotice) error {
	log.Infof("微信通知....")
	muban1 := `出错服务器:{{template "name"}};\n日志来源:{{template "source"}};\n报错时间:{{template "timestamp"}};\n日志信息:{{template "message"}}`
	tmpl, err := template.New("temp").Parse(muban1)
	if err != nil {
		log.Infof("微信通知模解析失败 :%s",err)
		return err
	}
	log.Debugf("要发送的数据 m %#V",m)
	tmpl.New("name").Parse(m.Beat.Name)
	tmpl.New("source").Parse(m.Source)
	tmpl.New("timestamp").Parse(m.Timestamp)
	tmpl.New("message").Parse(m.Message)
	b := new(bytes.Buffer)
	if err = tmpl.Execute(b, nil);err != nil {
		log.Infof("模板执行失败:%s",err)
		return err
	}
	sent := &Text{
		Touser:  "@all",
		Toparty: "@all",
		Totag:   "@all",
		Msgtype: "text",
		Agentid: 1000002,
		Text: struct {
			Content string `json:"content"`
		}{
			Content: b.String(),
		},
	}
	url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=ww6079b55c0c74769c&corpsecret=Tn66KjP55jhk7wEVvMARP04s3wzLM8veYPloTRYIQag"
	//获取access_token
	request, err := http.NewRequest("Get", url, nil)
	if err != nil {
		log.Infof("获取微信access_token有误:%s",err)
		return err
	}
	client := http.Client{}
	resp, err :=client.Do(request)
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Infof("网络延迟 err (%s)",err)
		return err
	}
	rest := &Access{}
	if err := json.Unmarshal(result, &rest);err != nil {
		return err
	}
	log.Debugf("access token :",rest.AccessToken)
	url = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + rest.AccessToken
	bb, err := json.Marshal(sent)
	if err != nil {
		log.Debugf("sebt 数据marshal err (%s)",err)
		return err
	}
	reader := bytes.NewReader(bb)
	request, err = http.NewRequest("POST", url, reader)
	if err != nil {
		log.Infof("发送微信通知失败!%s",err)
		return err
	}
	client = http.Client{}
	resp, err = client.Do(request)
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Infof("发送失败 err (%s)",err)
		return err
	}
	log.Infof("发送微信通知成功!")
	return nil
}