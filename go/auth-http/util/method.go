package util

import (
	"bufio"
	"encoding/json"
	log "github.com/cihub/seelog"
	"net/http"
	"io/ioutil"
	"strconv"
)

//处理HTTP请求中的数据
func GetHttpData(req *http.Request, types string,t interface{}) int32 {
	bodyType := req.Header.Get("Content-Type")
	log.Info("GetHttp-Content-Type:",bodyType)
	//index := strings.Contains(types, bodyType)
	if types != bodyType {
		return 404
	}
	//获取body
	body, bodyErr := ioutil.ReadAll(bufio.NewReader(req.Body))
	if bodyErr != nil {
		log.Error("GetHttp-ReadAll-bodyErr:",bodyErr)
		return 10006
	}
	bodyStr := string(body)
	//bodyStr = strings.Replace(strings.Replace(bodyStr," ","",-1),"\n","",-1)
	log.Info("GetHttp-bodyStr:",bodyStr)
	if len(body) < 8 {
		return 10006
	}
	if ssoJsonerr := json.Unmarshal(body, &t); ssoJsonerr !=nil {
		log.Error("GetHttp-Unmarshal-ssoJsonerr:",ssoJsonerr)
		return 10006
	}
	return 10000
}


func ResCode(code int, res http.ResponseWriter)  {
	if code != 0 {
		res.WriteHeader(code)
		res.Write(nil)
	}
}


// 对参数进行int32转换
func ChangeInt64(did string) int64 {
	d, err := strconv.Atoi(did)
	log.Info("did :",did)
	if err != nil || d == 0 {
		log.Error("strconv.Atoi(did) Failed,", err)
		return 0
	}
	domainId := int64(d)
	return domainId
}





