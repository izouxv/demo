package router

import (
	"github.com/julienschmidt/httprouter"
	. "file-server/controller"
	. "file-server/common"
)


func NewRouter(router *httprouter.Router)  {
	router.POST(FilesUrl,   Upload)
	router.GET(FileUrl,  Download)
	router.DELETE(FilesUrl,  Delete)
}
