package common

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"bufio"
	log "github.com/cihub/seelog"
	"strconv"
	"regexp"
)

//邮箱匹配
var (
	MobileReg = regexp.MustCompile("^((1[3,5,8][0-9])|(14[5,7])|(17[0,1,6,7,8]))\\d{8}$")
	EmailReg = regexp.MustCompile("^\\w+([-_.]?\\w+)*@\\w+([\\.-]?\\w+)*(\\.\\w{2,6})+$")
)

//string to int64
func StrToInt64(str string) (int64, error) {
	num, err := strconv.ParseInt(str,10,64)
	return num, err
}

//string to int32
func StrToInt32(str string) (int32, error) {
	num ,err := strconv.Atoi(str)
	if err != nil {
		return  0,err
	}
	return int32(num),nil
}

//处理HTTP请求中的数据
func GetJsonHttpData(req *http.Request, t interface{}) bool {
	//获取body
	body, bodyErr := ioutil.ReadAll(bufio.NewReader(req.Body))
	log.Info("body:",string(body),"bodyErr:",bodyErr)
	if bodyErr != nil {
		log.Info("GetHttp-ReadAll-bodyErr:", bodyErr)
		return true
	}
	if ssoJsonErr := json.Unmarshal(body, &t); ssoJsonErr != nil {
		log.Info("GetHttp-Unmarshal-ssoJsonErr:", ssoJsonErr)
		return true
	}
	return false
}

func ReturnCountPageOrderBy(req *http.Request) (count, page int32,orderBy string) {
	count, _ = StrToInt32(req.FormValue("count"))
	page, _ = StrToInt32(req.FormValue("page"))
	orderBy = req.FormValue("order_by")
	return
}

// 获取Cookie
func GetCookie( r *http.Request,key string) string{
	cookie,err := r.Cookie(key)
	if err != nil {
		return ""
	}
	return cookie.Value
}
