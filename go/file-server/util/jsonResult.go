package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JsonReply(msg string, result interface{}, w http.ResponseWriter) {
	r := GetCodeAndMsg(msg)
	r.Result = result
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	value, _ := json.Marshal(r)
	fmt.Fprintf(w, "%s", value)
	return
}
