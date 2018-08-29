package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"file-server/util"
	"os"
	"io"
	log "github.com/cihub/seelog"
	"file-server/internal/storage"
	"regexp"
	"encoding/json"
	"io/ioutil"
	"bufio"
	"github.com/pkg/errors"
	"github.com/jinzhu/gorm"
	"sort"
	"fmt"
)

var md5Reg = regexp.MustCompile(`[a-z0-9]+`)

func Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             // 允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") // header的类型
	w.Header().Set("content-type", "application/json;charset=utf-8")// 返回数据格式是json
	var files []string
	err := readBody(r.Body,&files)
	if err != nil || len(files) == 0 {
		log.Error("Delete readBody, err:",err)
		util.JsonReply(util.DeleteFail, nil, w)
		return
	}
	sort.Strings(files)
	files = RemoveDuplicatesAndEmpty(files...)
	var fids []string
	var resultFids []string
	for _,v := range files {
		if md5Reg.MatchString(v) {
			fids = append(fids,v)
		} else {
			resultFids = append(resultFids,v)
		}
	}
	fileInfo := storage.File{}
	fileInfos,err := fileInfo.GetList(fids...)
	if err == gorm.ErrRecordNotFound || len(fileInfos) == 0 {
		util.JsonReply(util.DeleteFail, files, w)
		return
	}
	if err != nil {
		log.Error("Delete GetList, err:",err)
		util.JsonReply(util.SystemError, files, w)
		return
	}
	fids = []string{}
	for _,v := range fileInfos {
		if pathIsNotExist(v.Path) {
			log.Error("Delete pathIsNotExist, Path:",v.Path)
			resultFids = append(resultFids,v.Fid)
			continue
		}
		if os.Remove(v.Path) != nil {
			log.Error("Delete Remove, Path:",fileInfo.Path)
			resultFids = append(resultFids,v.Fid)
			continue
		}
		for k,vv := range files {
			if vv == v.Fid {
				files = append(files[:k],files[k+1:]...)
			}
		}
		fids = append(fids,v.Fid)
	}
	resultFids = append(resultFids,files...)
	if len(fids) == 0 {
		util.JsonReply(util.SystemError, resultFids, w)
		return
	}
	if err = fileInfo.DeleteList(fids...); err != nil {
		log.Error("Delete DeleteList, err:",err)
		util.JsonReply(util.SystemError, files, w)
		return
	}
	fmt.Println("resultFids:",resultFids)
	util.JsonReply(util.Successful, resultFids, w)
}

func compriseFid(fid string,fids ...string) bool {
	for _,v := range fids {
		if fid == v {
			return false
		}
	}
	return true
}

func RemoveDuplicatesAndEmpty(fids ...string) (result []string){
	for k,_ := range fids {
		if (k > 0 && fids[k-1] == fids[k]) || len(fids[k])==0{
			continue
		}
		result = append(result, fids[k])
	}
	return
}

func readBody(read io.Reader, t interface{}) error {
	body, bodyErr := ioutil.ReadAll(bufio.NewReader(read))
	if bodyErr != nil {
		return bodyErr
	}
	size := len(body)
	log.Info("GetHttp-bodyLen:", size)
	if size < 2 {
		return errors.New("数据长度错误")
	}
	if size < 1024 {
		log.Info("GetHttp-bodyStr:", string(body))
	}
	if err := json.Unmarshal(body, &t); err != nil {
		return err
	}
	return nil
}

//true不存在
func pathIsNotExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return false
	}
	if os.IsNotExist(err) {
		return true
	}
	return false
}
