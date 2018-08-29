package result

import (
	"net/http"
	"fmt"
	"encoding/json"
	log "github.com/cihub/seelog"
)

func JsonReply(msg string, result interface{}, w http.ResponseWriter) {
	r := GetCodeAndMsg(msg)
	r.Result = result
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	value, _ := json.Marshal(r)
	log.Info("RES-json:",value)
	fmt.Fprintf(w, "%s", value)
	return
}

func ResCode(code int, res http.ResponseWriter)  {
	if code != 0 {
		res.WriteHeader(code)
		res.Write(nil)
	}
}
