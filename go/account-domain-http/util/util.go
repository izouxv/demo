package util

import  (
	"strconv"
	"net/http"
	"bufio"
	"io/ioutil"
	"encoding/json"
	log "github.com/cihub/seelog"
	"regexp"
)

var TidToSource  = map[int64]string{
	100001:"AQIDAA==",
	100002:"AgIDAA==",
	100003:"BAIDAA==",
	100004:"AwIDAA==",
}

var DecodeSourceToTid  = map[byte]int64{
	1:100001,
	2:100002,
	3:100004,
	4:100003,
}

var VersionCodeRegexp = regexp.MustCompile("^[0-9]+[.][0-9]+[.][0-9]+$")


func StringToInt32 (input string ) (output int32,err error ){
	inputstr, err := strconv.Atoi(input)
	return  int32(inputstr), err
}

func StrToInt64(str string) (int64, error) {
	num, err := strconv.ParseInt(str,10,64)
	return num, err
}


func ReturnCountPageOrderBy(req *http.Request) (count, page int32,orderBy string) {
	count, _ = StringToInt32(req.FormValue("count"))
	page,  _ = StringToInt32(req.FormValue("page"))
	orderBy  = req.FormValue("order_by")
	return
}

func ReturnCountPage(req *http.Request) (count, page int32) {
	count, _ = StringToInt32(req.FormValue("per_page"))
	page,  _ = StringToInt32(req.FormValue("page"))
	return
}

//处理HTTP请求中的数据
func GetJsonHttpData(req *http.Request, t interface{}) bool {
	//获取body
	body, bodyErr := ioutil.ReadAll(bufio.NewReader(req.Body))
	log.Info("request body:",string(body))
	if bodyErr != nil {
		log.Error("GetHttp-ReadAll-bodyErr:", bodyErr)
		return true
	}
	if ssoJsonErr := json.Unmarshal(body, &t); ssoJsonErr != nil {
		log.Error("GetHttp-Unmarshal-ssoJsonErr:", ssoJsonErr)
		return true
	}
	return false
}
