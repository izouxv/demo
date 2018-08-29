package controller

import (
	"file-server/internal/storage"
	"file-server/util"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func Download(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fid := p.ByName("fid")
	fileInfo := storage.File{Fid: fid}
	fileInfo.GetFileForFid()
	file, err := os.Open(fileInfo.Path)
	defer file.Close()
	if err != nil {
		util.JsonReply(util.SystemError, nil, w)
		return
	}
	fileName := url.QueryEscape(fileInfo.Name)
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("content-disposition", "attachment; filename=\""+fileName+"\"")
	body, err := ioutil.ReadAll(file)
	if err != nil {
		util.JsonReply(util.SystemError, nil, w)
		return
	}
	w.Write(body)
	fmt.Fprint(w)
}
