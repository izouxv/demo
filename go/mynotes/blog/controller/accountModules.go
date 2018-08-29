package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"mynotes/blog/logger"
)

func Index(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	logger.Info("aaa")
	res.Write([]byte("aaa"))
}
