package controller

import (
	pb "account-domain-http/api"
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/cihub/seelog"
	"account-domain-http/util"
	"account-domain-http/rpc"
	"context"
)
type DevType struct {
	Id  int32         `json:"id"`
	Tid int64         `json:"tid"`
	DeviceType string `json:"deviceType"`
	Status    int32   `json:"status"`
}

type DevTypeList struct {
	TotalCount int32   `json:"totalCount"`
	DevType []DevType `json:"device_type"`
}

func GetDeviceTypes(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start GetDeviceTypes")
	countStr := r.FormValue("count")
	pageStr := r.FormValue("page")
	orderStr := r.FormValue("order")
	tid, err := util.StrToInt64(p.ByName("tid"))
	if err != nil {
		log.Infof("tid (%d),error :(%s)", tid, err)
		util.JsonReply("TenantId_is_incorrect_or_empty", nil, w)
		return
	}
	count, err := util.StringToInt32(countStr)
	page, err := util.StringToInt32(pageStr)
	if err != nil {
		util.JsonReply("Params_error", nil, w)
		return
	}
	reply, err := rpc.DevTypeRpcClient().GetDeviceTypes(context.Background(), &pb.GetDeviceTypesRequest{Tid: tid, Page: page, Count: count, OrderBy: orderStr})
	if err != nil {
		log.Error("调用account-domain-rpc GetDeviceTypes failed:", err)
		util.JsonReply("System_error", nil, w)
		return
	}
	if reply.ErrorCode == util.DeviceType_Does_Not_Exist{
		util.JsonReply("No_DeviceType_Can_Be_Find", nil, w)
		return
	}
	log.Infof("rpc返回状态码 (%d) ", reply.ErrorCode)
	log.Debug(reply.DeviceType)
	devTypeList := DevTypeList{TotalCount: reply.TotalCount}
	for _, v := range reply.DeviceType {
		dt := DevType{
			Id:  v.Id,
			Tid: v.Tid,
			DeviceType:v.DeviceType,
			Status:v.Status,
		}
		devTypeList.DevType = append(devTypeList.DevType, dt)
	}
	util.JsonReply("Successful", devTypeList, w)
	return
}

//基于租户id增加设备类型
func AddDeviceType(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start AddDeviceType")
	r.ParseForm()
	tid, err := util.StrToInt64(p.ByName("tid"))
	if err != nil || tid == 0 {
		log.Infof("tid (%d),error :(%s)",tid,err)
		util.JsonReply("TenantId_is_incorrect_or_empty", nil, w)
		return
	}
	dt := &pb.Type{}
	errCode := util.GetHttpData(r, "application/json;charset=UTF-8", &dt)
	if errCode != util.Successfull {
		if errCode == 404 {
			util.ResCode(http.StatusNotFound, w)
			return
		}
		util.JsonReply("Body_is_incorrect_or_empty", nil, w)
		return
	}
	if dt.DeviceType == "" {
		util.JsonReply("Params_error", nil, w)
		return
	}
	dt.Tid = tid
	reply, err := rpc.DevTypeRpcClient().AddDeviceType(context.Background(), &pb.AddDeviceTypeRequest{DeviceType:dt})
	if err != nil {
		log.Error("调用account-domain-rpc AddDeviceType failed: ", err)
		util.JsonReply("System_error", nil, w)
		return
	}
	if reply.ErrorCode != util.Successfull{
		util.JsonReply("System_error", nil, w)
		return
	}
	log.Info("rpc返回状态码:", reply.ErrorCode)
	util.JsonReply("Successful",nil, w)
	return
}

