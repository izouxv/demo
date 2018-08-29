package result

import (
	"fmt"
	log "github.com/cihub/seelog"
	"net/http"
	"petfone-http/util"
)

//对处理结果进行封装并返回
func RESC(code int32, res http.ResponseWriter) {
	myJson := Jsons(code)
	if myJson.Code == 0 {
		switch code {
		case 33014:
			code = 20026
		case 33012:
			code = 21005
		case 37001:
			code = 33011
		case 37031:
			code = 33013
		default:
			code = 10002
		}
		myJson = Jsons(code)
	}
	myJson.Result = nil
	resw(myJson, res)
}
func REST(result interface{}, res http.ResponseWriter) {
	myjson := Jsons(10000)
	myjson.Result = result
	resw(myjson, res)
}
func resw(myJson Myjson, res http.ResponseWriter) {
	log.Info("resw json:", myJson.Code,myJson.Msg)
	res.Header().Set("Content-Type", "application/json;charset=utf-8")
	res.Header().Set("code", util.Int32ToStr(myJson.Code))
	value, err := util.Json.MarshalToString(myJson)
	if err != nil {
		log.Error("resw Marshal-err:",err)
		RESC(10001, res)
		return
	}
	fmt.Fprint(res, value)
}
func ResCode(code int, res http.ResponseWriter) {
	if code != 0 {
		res.WriteHeader(code)
		res.Write(nil)
	}
}
