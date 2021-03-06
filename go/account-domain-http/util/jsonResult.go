package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetJsonResult(code string, msg string, result interface{}) map[string]interface{} {
	j1 := make(map[string]interface{})
	j1["code"] = code
	j1["msg"] = msg
	j1["result"] = result
	return j1
}

func JsonReply(msg string, result interface{}, w http.ResponseWriter) {
	r := GetCodeAndMsg(msg)
	r.Result = result
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	value, _ := json.Marshal(r)
	fmt.Fprintf(w, "%s", value)
	return
}

