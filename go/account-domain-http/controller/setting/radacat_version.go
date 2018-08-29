package controller

import (
	pb "account-domain-http/api/setting"
	"account-domain-http/rpc"
	"account-domain-http/util"
	log "github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
)

type RVList struct {
	RadacatVersions []*RadacatVersion `json:"radacatVersions"`
	TotalCount      int32             `json:"totalCount"`
}

type RadacatVersion struct {
	Device        string `json:"device"`
	VersionName   string `json:"versionName"`
	VersionCode   string `json:"versionCode"`
	Md5           string `json:"md5"`
	Filename      string `json:"filename"`
	Length        int64  `json:"length"`
	Path          string `json:"path"`
	DescriptionCn string `json:"descriptionCn"`
	DescriptionEn string `json:"descriptionEn"`
	UploaderUid   int64  `json:"uploaderUid,omitempty"`
	Status        int32  `json:"status"`
	Id            int32  `json:"id"`
	CreateTime    int64  `json:"createTime"`

}

type UpdateRadacatVersion struct {
	Md5           string `json:"md5"`
	Filename      string `json:"filename"`
	Length        int64  `json:"length"`
	Device        string `json:"device"`
	Path          string `json:"path"`
	VersionName   string `json:"versionName"`
	VersionCode   string `json:"versionCode"`
	DescriptionCn string `json:"descriptionCn"`
	DescriptionEn string `json:"descriptionEn"`
	Status        int32  `json:"status"`
	Id            int32  `json:"id"`
	Tid           int64  `json:"tid"`
}

func AddNewVersion(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Infof("Start AddNewVersion")
	rv := RadacatVersion{}
	errCode := util.GetHttpData(r, "application/json;charset=UTF-8", &rv)
	if errCode != util.Successfull {
		if errCode == 404 {
			util.ResCode(http.StatusNotFound, w)
			return
		}
		util.JsonReply("Body_is_incorrect_or_empty", nil, w)
		return
	}
	tid, err := util.StrToInt64(p.ByName("tid"))
	if err != nil {
		log.Infof("tid (%d),error :(%s)",tid,err)
		util.JsonReply("TenantId_is_incorrect_or_empty", nil, w)
		return
	}
	if rv.Path == "" {
		log.Info("Path 不能为空")
		util.JsonReply("FileUrl_is_incorrect_or_empty", nil, w)
		return
	}
	log.Info("版本号：",rv.VersionCode)
	if !util.VersionCodeRegexp.MatchString(rv.VersionCode){
		log.Info("版本号为空或格式有误")
		util.JsonReply("VersionCode_is_incorrect_or_empty", nil, w)
		return
	}
	if rv.Device == "" {
		log.Info("Device不能为空")
		util.JsonReply("Device_is_incorrect_or_empty", nil, w)
		return
	}
	if rv.Md5 == "" {
		log.Info("md5不能为空")
		util.JsonReply("md5_is_incorrect_or_empty", nil, w)
		return
	}
	if rv.Filename == "" {
		log.Info("fileName 不能为空")
		util.JsonReply("filename_is_nil", nil, w)
		return
	}
	if rv.Length == 0 {
		log.Info("filelength 不能为空")
		util.JsonReply("filelength_is_incorrect_or_empty", nil, w)
		return
	}
	if rv.Status == 0 {
		log.Info("Status 不能为0")
		util.JsonReply("Status_is_incorrect_or_empty", nil, w)
		return
	}
	req := pb.AddNewVersionRequest{Version: &pb.Version{
		Tid:           tid,
		Device:        rv.Device,
		VersionName:   rv.VersionName,
		VersionCode:   rv.VersionCode,
		DescriptionCn: rv.DescriptionCn,
		DescriptionEn: rv.DescriptionEn,
		Status:        rv.Status,
		FileName:      rv.Filename,
		Md5:           rv.Md5,
		Path:          "http://file.radacat.com:88/v1.0/file/"+ rv.Path,
		Length:        rv.Length,
	}}

	//todo 调用rpc
	reply, err  := rpc.RadacatVersionRpcClient().AddNewVersion(context.Background(), &req)
	if err != nil {
		log.Error("调用rpc AddNewVersion failed : ", err)
		util.JsonReply("System_error", nil, w)
		return
	}
	if reply.ErrorCode != util.Successfull {
		switch reply.ErrorCode {
		case util.Input_parameter_error:
			util.JsonReply("Body_is_incorrect_or_empty", nil, w)
			return
		case util.Version_is_exist :
			util.JsonReply("RadacatVersion_Can_Be_Find", nil, w)
			return
		case util.System_error:
			util.JsonReply("System_error", nil, w)
			return
		}
	}
	util.JsonReply("Successful", nil, w)
	return
}

func GetAllVersions(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Infof("Start GetAllVersions")
	tid, err := util.StrToInt64(p.ByName("tid"))
	if err != nil {
		log.Infof("tid (%d),error :(%s)",tid,err)
		util.JsonReply("TenantId_is_incorrect_or_empty", nil, w)
		return
	}
	req := pb.GetAllVersionsRequest{Tid:tid}
	count := r.FormValue("per_page")
	page := r.FormValue("page")
	log.Infof("参数 count: (%s) page: (%s)", count, page)
	//todo count
	if count != "" {
		c, err := strconv.Atoi(count)
		if err != nil || c == 0 {
			log.Info("strconv.Atoi(count) Failed:", err)
			util.JsonReply("Per_page_is_incorrect_or_empty", nil, w)
			return
		}
		req.Count = int32(c)
	}
	//todo page
	if page != "" {
		p, err := strconv.Atoi(page)
		if err != nil || p == 0 {
			log.Info("strconv.Atoi(page) Failed:", err)
			util.JsonReply("Page_is_incorrect_or_empty", nil, w)
			return
		}
		req.Page = int32(p)
	}
	reply, err  := rpc.RadacatVersionRpcClient().GetAllVersions(context.Background(), &req)
	if err != nil {
		log.Error("调用rpc GetAllVersions failed ", err)
		util.JsonReply("System_error", nil, w)
		return
	}
	log.Info("rpc errorCode:", reply.ErrorCode)
	if reply.ErrorCode != util.Successfull {
		switch reply.ErrorCode {
		case util.Input_parameter_error:
			util.JsonReply("Body_is_incorrect_or_empty", nil, w)
			return
		case util.System_error:
			util.JsonReply("System_error", nil, w)
			return
		case util.NO_RadacatVersion_Can_Be_Find:
			util.JsonReply("No_RadacatVersion_Can_Be_Find", nil, w)
			return
		}
	}
	rvs := RVList{TotalCount: reply.TotalCount}
	for k, v := range reply.Versions {
		if v.Id != 0 {
			rvs.RadacatVersions = append(rvs.RadacatVersions, &RadacatVersion{
				Id:            reply.Versions[k].Id,
				Device:        reply.Versions[k].Device,
				VersionName:   reply.Versions[k].VersionName,
				VersionCode:   reply.Versions[k].VersionCode,
				Md5:           reply.Versions[k].Md5,
				Filename:      reply.Versions[k].FileName,
				Length:        reply.Versions[k].Length,
				Path:          reply.Versions[k].Path,
				DescriptionCn: reply.Versions[k].DescriptionCn,
				DescriptionEn: reply.Versions[k].DescriptionEn,
				Status:        reply.Versions[k].Status,
				CreateTime:    reply.Versions[k].CreateTime,
			})
		}
	}
	util.JsonReply("Successful", rvs, w)
	return
}

func UpdateVersion(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Infof("Start UpdateVersion")
	tid, err := util.StrToInt64(p.ByName("tid"))
	if err != nil {
		log.Infof("tid (%d),error :(%s)",tid,err)
		util.JsonReply("TenantId_is_incorrect_or_empty", nil, w)
		return
	}
	id, err := util.StringToInt32(p.ByName("id"))
	urv := UpdateRadacatVersion{Tid:tid,Id:id}
	errCode := util.GetHttpData(r, "application/json;charset=UTF-8", &urv)
	if errCode != util.Successfull {
		if errCode == 404 {
			util.ResCode(http.StatusNotFound, w)
			return
		}
		util.JsonReply("Body_is_incorrect_or_empty", nil, w)
		return
	}
	log.Info("修改的版本号为:",urv.VersionCode)
	if !util.VersionCodeRegexp.MatchString(urv.VersionCode){
		log.Info("版本号为空或格式有误")
		util.JsonReply("VersionCode_is_incorrect_or_empty", nil, w)
		return
	}
	if urv.Id == 0 {
		log.Info("版本ID不能为空")
		util.JsonReply("Id_is_incorrect_or_empty", nil, w)
		return
	}
	req := &pb.UpdateVersionRequest{Version: &pb.Version{
		Id:            id,
		Tid:           urv.Tid,
		Md5:           urv.Md5,
		FileName:      urv.Filename,
		Length:        urv.Length,
		Path:          urv.Path,
		VersionName:   urv.VersionName,
		VersionCode:   urv.VersionCode,
		DescriptionCn: urv.DescriptionCn,
		DescriptionEn: urv.DescriptionEn,
		Status:        urv.Status,
		Device:        urv.Device,
	}}
	log.Info("device:",urv.Device)
	//todo 调用rpc
	log.Infof("参数  id (%d) tid (%d)",req.Version.Id,tid)
	reply, err := rpc.RadacatVersionRpcClient().UpdateVersion(context.Background(), req)
	if err != nil {
		log.Error("调用rpc UpdateVersion failed: ", err)
		util.JsonReply("System_error", nil, w)
		return
	}
	if reply.ErrorCode != util.Successfull {
		switch reply.ErrorCode {
		case util.Input_parameter_error:
			util.JsonReply("Body_is_incorrect_or_empty", nil, w)
			return
		case util.Version_is_exist :
			util.JsonReply("RadacatVersion_Can_Be_Find", nil, w)
			return
		case util.System_error:
			util.JsonReply("System_error", nil, w)
			return
		}
	}
	util.JsonReply("Successful", nil, w)
	return
}

func DeleteVersion(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Infof("Start DeleteVersion")
	tid, err := util.StrToInt64(p.ByName("tid"))
	if err != nil {
		log.Infof("tid (%d),error :(%s)",tid,err)
		util.JsonReply("TenantId_is_incorrect_or_empty", nil, w)
		return
	}
	id, err := strconv.ParseInt(p.ByName("id"), 10, 64)
	if err != nil {
		log.Infof("strconv.Atoi(id) Failed , err is (%s) id (%d)", err,id)
		util.JsonReply("Id_is_incorrect_or_empty", nil, w)
		return
	}
	//todo 调用rpc
	reply, err := rpc.RadacatVersionRpcClient().DeleteVersion(context.Background(), &pb.DeleteVersionRequest{Id:int32(id),Tid:tid})
	if err != nil {
		log.Error("调用rpc DeleteVersion failed ", err)
		util.JsonReply("System_error", nil, w)
		return
	}
	if reply.ErrorCode != util.Successfull {
		switch reply.ErrorCode {
		case util.System_error:
			util.JsonReply("System_error", nil, w)
			return
		case util.NO_RadacatVersion_Can_Be_Find:
			util.JsonReply("File_is_exist", nil, w)
			return
		}
	}
	util.JsonReply("Successful", nil, w)
	return
}

func GetVersion(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Infof("Start GetVersion")
	tid, err := util.StrToInt64(p.ByName("tid"))
	if err != nil {
		log.Infof("tid (%d),error :(%s)",tid,err)
		util.JsonReply("TenantId_is_incorrect_or_empty", nil, w)
		return
	}
	id, err := strconv.ParseInt(p.ByName("id"), 10, 64)
	if err != nil {
		log.Infof("strconv.Atoi(id) Failed , err is (%s) id (%d)", err,id)
		util.JsonReply("Id_is_incorrect_or_empty", nil, w)
		return
	}
	//todo 调用rpc
	reply, err := rpc.RadacatVersionRpcClient().GetVersion(context.Background(), &pb.GetVersionRequest{Id:int32(id),Tid:tid})
	if err != nil {
		log.Error("调用rpc GetVersion failed ", err)
		util.JsonReply("System_error", nil, w)
		return
	}
	if reply.ErrorCode != util.Successfull {
		switch reply.ErrorCode {
		case util.System_error:
			util.JsonReply("System_error", nil, w)
			return
		case util.NO_RadacatVersion_Can_Be_Find:
			util.JsonReply("File_is_exist", nil, w)
			return
		}
	}
	util.JsonReply("Successful", reply.Version, w)
	return
}


