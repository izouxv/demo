package controller

import (
	pb "account-domain-http/api/adv"
	"account-domain-http/rpc"
	"account-domain-http/util"
	log "github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
	"net/http"
)

type AdvId struct {
	Id int32 `json:"id"`
}
type AdvInfo struct {
	Id         int32  `json:"id"`
	Tid        int64  `json:"tid"`
	Name       string `json:"name"`
	State      int32  `json:"state"`
	StartTime  int64  `json:"starttime"`
	Md5        string `json:"md5"`
	EndTime    int64  `json:"endtime"`
	Advertiser string `json:"advertiser"`
	FileUrl    string `json:"fileurl"`
	AdvUrl     string `json:"advurl"`
	FileName   string `json:"fileName"`
}
type AdvList struct {
	Advs       []AdvInfo `json:"advs"`
	TotalCount int32     `json:"totalCount"`
}
type NodeAd struct {
	DevEUI string `json:"devEUI"`
	Ad     []struct {
		ID  int32  `json:"id"`
		Md5 string `json:"md5"`
	} `json:"ad"`
}

/*设备广告*/
type NodeAdvs struct {
	DevEUI string `json:"devEUI"`
	Ad     []*Ad  `json:"ad"`
}
type Ad struct {
	Id  int32  `json:"id"`
	Url string `json:"url"`
	Md5 string `json:"md5"`
}

//添加新广告
func NewAdvertisement(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start NewAdvertisement")
	tid, err := util.StrToInt64(p.ByName("tid"))
	if err != nil {
		log.Infof("tid (%d),error :(%s)",tid,err)
		util.JsonReply("TenantId_is_incorrect_or_empty", nil, w)
		return
	}
	advreq := pb.AdvertisementRequest{}
	errCode := util.GetHttpData(r, "application/json;charset=UTF-8", &advreq)
	if errCode != util.Successfull {
		if errCode == 404 {
			util.ResCode(http.StatusNotFound, w)
			return
		}
		util.JsonReply("Body_is_incorrect_or_empty", nil, w)
		return
	}
	log.Info("NewAdvertisementInfo :", advreq)
	if advreq.Name == "" {
		util.JsonReply("AdvName_is_incorrect_or_empty", nil, w)
		return
	}
	if advreq.FileName == ""{
		util.JsonReply("FileName_is_incorrect_or_empty", nil, w)
		return
	}
	if advreq.StartTime == 0 {
		util.JsonReply("StartTime_is_incorrect_or_empty", nil, w)
		return
	}
	if advreq.EndTime == 0 {
		util.JsonReply("EndTime_is_incorrect_or_empty", nil, w)
		return
	}
	if advreq.Advertiser == "" {
		util.JsonReply("Advertiser_is_incorrect_or_empty", nil, w)
		return
	}
	if advreq.FileUrl == "" {
		util.JsonReply("FileUrl_is_incorrect_or_empty", nil, w)
		return
	}
	if advreq.Md5 == "" {
		util.JsonReply("md5_is_incorrect_or_empty", nil, w)
		return
	}
	//todo 调用rpc
	reply, err := rpc.AdvRpcClient().NewAdvertisement(context.Background(), &pb.AdvertisementRequest{
		Tid:        tid,
		Name:       advreq.Name,
		Md5:        advreq.Md5,
		StartTime:  advreq.StartTime,
		EndTime:    advreq.EndTime,
		FileName:   advreq.FileName,
		FileUrl:    advreq.FileUrl,
		AdvUrl:     advreq.AdvUrl,
		Advertiser: advreq.Advertiser,
	})
	if err != nil {
		log.Error("调用rpc NewAdvertisement failed :", err)
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
		}
	}
	util.JsonReply("Successful", nil, w)
	return
}

//获取广告
func GetAdvertisements(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start GetAdvertisements")
	tid, err := util.StrToInt64(p.ByName("tid"))
	if err != nil {
		log.Infof("tid (%d),error :(%s)",tid,err)
		util.JsonReply("TenantId_is_incorrect_or_empty", nil, w)
		return
	}
	countStr := r.FormValue("per_page")
	pageStr := r.FormValue("page")
	count ,err := util.StringToInt32(countStr)
	if err != nil || count == 0 {
			log.Infof("strconv.Atoi(count) Failed:(%s) count :(%d)", err,countStr)
			util.JsonReply("Per_page_is_incorrect_or_empty", nil, w)
			return
		}
	//todo page
	page ,err := util.StringToInt32(pageStr)
	if err != nil || page == 0 {
		log.Infof("strconv.Atoi(page) Failed:(%s)  page: (%d)", err,pageStr)
		util.JsonReply("Page_is_incorrect_or_empty", nil, w)
		return
	}
	reply, err := rpc.AdvRpcClient().GetAllAdvertisement(context.Background(), &pb.AdvertisementRequest{Page:page,Count:count,Tid:tid})
	if err != nil {
		log.Error("调用rpc GetAllAdvertisement failed :", err)
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
		case util.NO_Advertisement_Can_Be_Find:
			util.JsonReply("No_Advertisement_Can_Be_Find", nil, w)
			return
		}
	}
	advInfos := AdvList{TotalCount: reply.TotalCount}
	for _, v := range reply.Advs {
		if v.Id != 0 {
			advInfo := AdvInfo{
				Id:         v.Id,
				Tid:        v.Tid,
				Name:       v.Name,
				State:      v.State,
				StartTime:  v.StartTime,
				EndTime:    v.EndTime,
				Advertiser: v.Advertiser,
				FileUrl:    v.FileUrl,
				FileName:   v.FileName,
				AdvUrl:     v.AdvUrl}
			advInfos.Advs = append(advInfos.Advs, advInfo)
		}
	}
	util.JsonReply("Successful", advInfos, w)
	return
}

//修改广告
func UpdateAdvertisement(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start UpdateAdvertisementInfo")
	tid, err := util.StrToInt64(p.ByName("tid"))
	if err != nil {
		log.Infof("tid (%d),error :(%s)",tid,err)
		util.JsonReply("TenantId_is_incorrect_or_empty", nil, w)
		return
	}
	id, err := util.StrToInt64(p.ByName("id"))
	if err != nil {
		log.Infof("get id err (%s)",err)
		util.JsonReply("Id_is_incorrect_or_empty", nil, w)
		return
	}
	advreq := pb.AdvertisementRequest{}
	errCode := util.GetHttpData(r, "application/json;charset=UTF-8", &advreq)
	if errCode != util.Successfull {
		if errCode == 404 {
			util.ResCode(http.StatusNotFound, w)
			return
		}
		util.JsonReply("Body_is_incorrect_or_empty", nil, w)
		return
	}
	log.Infof("UpdateAdvertisementInfo tid (%d) id (%d)", tid,id)
	//todo 调用rpc
	reply, err := rpc.AdvRpcClient().UpdateAdvertisement(context.Background(), &pb.AdvertisementRequest{
		Tid:        tid,
		Id:         int32(id),
		Name:       advreq.Name,
		Md5:        advreq.Md5,
		StartTime:  advreq.StartTime,
		EndTime:    advreq.EndTime,
		FileUrl:    advreq.FileUrl,
		FileName:   advreq.FileName,
		AdvUrl:     advreq.AdvUrl,
		Advertiser: advreq.Advertiser,
		State:      advreq.State,
	})
	if err != nil {
		log.Errorf("调用rpc UpdateAdvertisement failed,err is (%s)", err)
		util.JsonReply("System_error", nil, w)
		return
	}
	if reply.ErrorCode != 10000 {
		switch reply.ErrorCode {
		case util.Input_parameter_error:
			util.JsonReply("Body_is_incorrect_or_empty", nil, w)
			return
		case util.System_error:
			util.JsonReply("System_error", nil, w)
			return
		case util.Advertisement_not_exist:
			util.JsonReply("No_Advertisement_Can_Be_Find", nil, w)
			return
		}
	}
	util.JsonReply("Successful", nil, w)
	return
}

//删除广告
func DeleteAdvertisement(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start DeleteAdvertisement")
	tid, err := util.StrToInt64(p.ByName("tid"))
	if err != nil {
		log.Infof("tid (%d),error :(%s)",tid,err)
		util.JsonReply("TenantId_is_incorrect_or_empty", nil, w)
		return
	}
	id, err := util.StrToInt64(p.ByName("id"))
	if err != nil {
		log.Infof("get id err (%s)",err)
		util.JsonReply("Id_is_incorrect_or_empty", nil, w)
		return
	}
	log.Infof("参数 tid (%d) id (%d)",tid,id)
	reply, err := rpc.AdvRpcClient().DelAdvertisement(context.Background(), &pb.AdvertisementRequest{Id:int32(id),Tid:tid})
	if err != nil {
		log.Errorf("调用 rpc DelAdvertisement failed,err is (%s)", err)
		util.JsonReply("System_error", nil, w)
		return
	}
	if reply.ErrorCode != 10000 {
		switch reply.ErrorCode {
		case util.Input_parameter_error:
			util.JsonReply("Body_is_incorrect_or_empty", nil, w)
			return
			break
		case util.System_error:
			util.JsonReply("System_error", nil, w)
			return
			break
		case util.Advertisement_not_exist:
			util.JsonReply("No_Advertisement_Can_Be_Find", nil, w)
			return
			break
		}
	}
	util.JsonReply("Successful", nil, w)
	return
}

//获取单个广告
func GetAdvertisement(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("Start GetAdvertisement")
	tid, err := util.StrToInt64(p.ByName("tid"))
	if err != nil {
		log.Infof("tid (%d),error :(%s)",tid,err)
		util.JsonReply("TenantId_is_incorrect_or_empty", nil, w)
		return
	}
	id, err := util.StrToInt64(p.ByName("id"))
	if err != nil {
		log.Infof("get id err (%s)",err)
		util.JsonReply("Id_is_incorrect_or_empty", nil, w)
		return
	}
	log.Infof("参数 tid (%d) id (%d)",tid,id)
	reply, err := rpc.AdvRpcClient().GetOneAdvertisement(context.Background(), &pb.AdvertisementRequest{Id:int32(id),Tid:tid})
	if err != nil {
		log.Errorf("调用 rpc GetAdvertisement failed,err is (%s)", err)
		util.JsonReply("System_error", nil, w)
		return
	}
	if reply.ErrorCode != 10000 {
		switch reply.ErrorCode {
		case util.Input_parameter_error:
			util.JsonReply("Body_is_incorrect_or_empty", nil, w)
			return
			break
		case util.System_error:
			util.JsonReply("System_error", nil, w)
			return
			break
		case util.Advertisement_not_exist:
			util.JsonReply("No_Advertisement_Can_Be_Find", nil, w)
			return
			break
		}
	}
	advInfo := AdvInfo{
		Id:         reply.Id,
		Tid:        reply.Tid,
		Name:       reply.Name,
		State:      reply.State,
		StartTime:  reply.StartTime,
		EndTime:    reply.EndTime,
		Advertiser: reply.Advertiser,
		FileUrl:    reply.FileUrl,
		FileName:   reply.FileName,
		AdvUrl:     reply.AdvUrl}
	util.JsonReply("Successful",advInfo , w)
	return
}
