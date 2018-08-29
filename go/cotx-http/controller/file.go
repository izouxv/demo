package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"cotx-http/result"
	"cotx-http/pb"
	log "github.com/cihub/seelog"
	"strconv"
	"fmt"
	"cotx-http/utils"
	"cotx-http/rpcClient"
	"golang.org/x/net/context"
)

func AddFile(res http.ResponseWriter ,req *http.Request ,param httprouter.Params  )  {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
	log.Debug(userinfo)
	var files = make([]*pb.File,0)
	var addFilesRequest = new(pb.AddFilesRequest)
	numString := req.FormValue("num")
	num,_ := strconv.Atoi(numString)
	typeString := req.FormValue("type")
	t,_ := strconv.Atoi(typeString)
	log.Debugf("num:%d,type:%d",num,t)
	for i := 0 ; i < num;i++  {
		var f = new(pb.File)
		name := fmt.Sprintf("file%d",i+1)
		file,header,err := req.FormFile(name)
		if err != nil {
			log.Error("error1")
			result.JsonReply("Update_File_Error",nil,res)
			return
		}
		 path := fmt.Sprintf("/data1/www/htdocs/gateway.cotxnetworks.com/upload/photo/%s",header.Filename)
		 //path := fmt.Sprintf("D:/ca/upload/photo/%s",header.Filename)
		 code := utils.CopyFileTo(path,file)
		if code != 200 {
			result.JsonReply("Update_File_Error",nil,res)
			return
		}
		url := fmt.Sprintf("http://gateway.cotxnetworks.com/upload/photo/%s",header.Filename)
		f.Url = url
		f.Type = int32(t)
		f.Name = header.Filename
		files = append(files,f)
	}
	addFilesRequest.Files = files
	log.Debug(addFilesRequest)
	addfilesResponse ,err:= rpcClient.GetFileClient().AddFiles(context.Background(),addFilesRequest)
	if addfilesResponse.ErrCode != 10000 || err != nil {
		log.Error("error:",err)
		switch addfilesResponse.ErrCode {
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		case 30027:
			result.JsonReply("Gateway_UnResponse",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",addfilesResponse.Files,res)
}
