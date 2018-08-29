package controller

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	. "file-server/common"
	log "github.com/cihub/seelog"
	"file-server/internal/storage"
	"file-server/module"
	"file-server/util"
	"github.com/julienschmidt/httprouter"
)

type File struct {
	Fid string `json:"fid"`
}

func Upload(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	file, handler, err := r.FormFile("file")
	if err != nil || file == nil {
		log.Error("Upload FormFile err:",err)
		util.JsonReply(util.UploadFail, nil, w)
		return
	}
	defer file.Close()
	w.Header().Set("Access-Control-Allow-Origin", "*")             // 允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") // header的类型
	w.Header().Set("content-type", "application/json;charset=utf-8")// 返回数据格式是json

	fileExt := filepath.Ext(handler.Filename)
	body, err := ioutil.ReadAll(file)
	log.Info("Upload len:",handler.Size)
	if err != nil {
		log.Error("Upload ReadAll err:",err)
		util.JsonReply(util.UploadFail, nil, w)
		return
	}
	fid := fmt.Sprintf("%x", md5.Sum(body))
	filename := fid + fileExt
	f, err := os.Create(UploadDir + filename)
	if err != nil {
		log.Error("Upload Create err:",err)
		util.JsonReply(util.UploadFail, nil, w)
		return
	}
	defer f.Close()
	size, err := f.Write(body)
	if err != nil || size == 0 {
		defer os.Remove(UploadDir + filename)
		log.Error("Upload Write err:",err)
		util.JsonReply(util.UploadFail, nil, w)
		return
	}
	path, _ := filepath.Abs(UploadDir + filename)
	saveFile(fid, handler.Filename, fileExt, path, size)
	util.JsonReply(util.Successful, File{Fid: fid}, w)
}

func saveFile(fid, name, ext, path string, size int) error {
	file := storage.File{Fid: fid, Name: name, Ext: ext, Size: size, Path: path}
	return file.Create(module.MysqlClient())
}
