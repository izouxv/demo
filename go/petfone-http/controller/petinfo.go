package controller

import (
	log "github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"petfone-http/pb"
	"petfone-http/po"
	"petfone-http/result"
	"petfone-http/rpc"
	"petfone-http/util"
	"sort"
)

/**
宠物信息
 */

//添加宠物信息
func SetPetInfo(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("SetPetInfo...")
	sso := util.GetContext(req)
	log.Info("SetPetInfo-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	petinfo := &pb.PetInfoRequest{}
	if util.GetHttpData(req, util.ReqMethodJson, &petinfo) {
		result.RESC(21002, res)
		return
	}
	if util.VerifyNickname(petinfo.GetNickname()) {
		result.RESC(21002, res)
		return
	}
	if util.VerifyParamsStr(petinfo.GetAvatar()) {
		result.RESC(21002, res)
		return
	}
	petinfo.Source = sso.Source
	petinfo.Uid = sso.Uid
	log.Info("SetPetInfo-petinfo:", petinfo)
	//调用rpc
	petinfoRe := rpc.PetinfoRpc(petinfo, "SetPetInfo")
	if petinfoRe.Code != 10000 {
		result.RESC(petinfoRe.Code, res)
		return
	}
	result.REST(&po.PetinfoPo{Pid:petinfoRe.Pid, Avatar:petinfoRe.Avatar,
		Nickname: petinfoRe.Nickname, Breed: petinfoRe.Breed, Gender: petinfoRe.Gender, Birthday: petinfoRe.Birthday,
		Weight: petinfoRe.Weight, Somatotype: petinfoRe.Somatotype, Duration:petinfoRe.Duration, Brightness:petinfoRe.Brightness,
		CreateTime: petinfoRe.CreateTime, Permit: petinfoRe.Permit,Trains:petinfoRe.Trains}, res)
}

//删除宠物信息
func DeletePetInfo(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("UpdatePetInfo...")
	pidStr := params.ByName("pid")
	sso := util.GetContext(req)
	log.Info("UpdateDeviceP-sso:", sso, ",pid:", pidStr)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	if util.VerifyParamsStr(pidStr) {
		result.RESC(21002, res)
		return
	}
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		result.RESC(10001, res)
		return
	}
	petInfo := &pb.PetInfoRequest{Source:sso.Source,Uid:sso.Uid}
	petInfo.Pid = int32(pid)
	log.Info("UpdatePetInfo-deviceJson:", petInfo)
	//调用rpc
	deviceRe := rpc.PetinfoRpc(petInfo, "DeletePetInfoByPid")
	result.RESC(deviceRe.Code, res)
}

//修改宠物信息
func UpdatePetInfo(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("UpdatePetInfo...")
	//todo 接收参数
	pidstr := params.ByName("pid")
	sso := util.GetContext(req)
	log.Info("UpdateDevice-sso:", sso, ",pidstr:", pidstr)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	petinfo := &pb.PetInfoRequest{}
	nofund := util.GetHttpData(req, util.ReqMethodJson, &petinfo)
	if nofund {
		result.RESC(21002, res)
		return
	}
	log.Info("UpdatePetInfo-petinfo:", petinfo)
	//todo 校验参数
	if util.VerifyParamsStr(pidstr) {
		log.Info("UpdatePetInfo Weight")
		result.RESC(21002, res)
		return
	}
	if len(petinfo.Nickname) > 100 {
		log.Info("UpdatePetInfo Nickname")
		result.RESC(21002, res)
		return
	}
	if len(petinfo.Avatar) > 150 {
		log.Info("UpdatePetInfo Avatar")
		result.RESC(21002, res)
		return
	}
	if util.VerifyParamsUInt32(petinfo.Breed, petinfo.Gender, petinfo.Somatotype, petinfo.Duration) {
		log.Info("UpdatePetInfo Breed Gender Somatotype Duration")
		result.RESC(21002, res)
		return
	}
	if util.VerifyParamsFloat32(petinfo.Weight) {
		log.Info("UpdatePetInfo Weight")
		result.RESC(21002, res)
		return
	}
	if util.VerifyParamsUInt64(petinfo.Birthday) {
		log.Info("UpdatePetInfo Birthday")
		result.RESC(21002, res)
		return
	}
	pid, err := util.StrToInt32(pidstr)
	if err != nil {
		result.RESC(10001, res)
		return
	}
	petinfo.Source = sso.Source
	petinfo.Pid = int32(pid)
	petinfo.Uid = sso.Uid
	//to do 校验用户对资源的权限
	//flag = db.VerifyUserPermiss(6379,petinfo.Source[:2]+db.PUsers+pidstr,strconv.Itoa(int(petinfo.Uid)),res)
	//num1, num2, err := db.Redis_ZscoreZcard(6379,petinfo.Source[:2]+db.PUsers+pidstr,strconv.Itoa(int(petinfo.GetUid())))
	//if err != nil {
	//	log.Info("VerifyUserPermiss-err:",err)
	//	result.RESC(10001, res)
	//	return
	//}
	//if num1 <= 0 || num2 <= 0{
	//	result.RESC(21005, res)
	//	return
	//}
	//todo 调用rpc
	deviceRe := rpc.PetinfoRpc(petinfo, "UpdatePetInfoByPid")
	if deviceRe.Code != 10000 {
		result.RESC(deviceRe.Code, res)
		return
	}
	result.REST(deviceRe, res)
}

//获取宠物信息
func GetPetInfo(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("GetPetInfo...")
	pidstr := params.ByName("pid")
	sso := util.GetContext(req)
	log.Info("GetPetInfo-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	if util.VerifyParamsStr(pidstr) {
		result.RESC(21002, res)
		return
	}
	pid, err := strconv.Atoi(pidstr)
	if err != nil {
		result.RESC(10001, res)
		return
	}
	petinfo := &pb.PetInfoRequest{Source:sso.Source,Pid:int32(pid),Uid:sso.Uid}
	log.Info("GetPetInfo-petinfo:", petinfo)
	//调用rpc
	petinfoRe := rpc.PetinfoRpc(petinfo, "GetPetInfoByPid")
	if petinfoRe.Code != 10000 {
		result.RESC(petinfoRe.Code, res)
		return
	}
	result.REST(&po.PetinfoPo{Pid: petinfoRe.Pid, Did: petinfoRe.Did, Avatar: petinfoRe.Avatar,
		Nickname: petinfoRe.Nickname, Breed: petinfoRe.Breed, Gender: petinfoRe.Gender, Birthday: petinfoRe.Birthday,
		Weight: petinfoRe.Weight, Somatotype: petinfoRe.Somatotype, Duration:petinfoRe.Duration, Brightness:petinfoRe.Brightness,
		CreateTime: petinfoRe.CreateTime, Permit: petinfoRe.Permit,Trains:petinfoRe.Trains}, res)
}

//获取批量宠物信息
func GetPetInfos(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("GetPetInfos...")
	sso := util.GetContext(req)
	log.Info("GetPetInfos-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	petInfo := &pb.PetInfoRequest{	Source:sso.Source,Uid:sso.Uid}
	log.Info("GetPetInfos-petInfo:", petInfo)
	//调用rpc
	petInfoMapRe := rpc.PetinfosRpc(petInfo, "GetPetInfoByUid")
	if petInfoMapRe.Code != 10000 {
		result.RESC(petInfoMapRe.Code, res)
		return
	}
	var arr []po.PetinfoPo
	for _, v := range petInfoMapRe.Petinfos {
		arr = append(arr, po.PetinfoPo{Pid: v.Pid, Did: v.Did, Avatar: v.Avatar, Nickname: v.Nickname,
		Breed: v.Breed, Gender: v.Gender, Birthday: v.Birthday, Weight: v.Weight, Somatotype: v.Somatotype,
		Duration:v.Duration, Brightness:v.Brightness, CreateTime: v.CreateTime, Permit: v.Permit,Trains:v.Trains})
	}
	result.REST(arr, res)
}

// todo modify xiaorx 2018/7/4
//获取宠物信息和宠端设备录音
func GetPetInfo1_1(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("GetPetInfo1_1...")
	sso := util.GetContext(req)
	log.Info("GetPetInfo-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	pid,err := util.StrToInt32(params.ByName("pid"))
	if err != nil {
		result.RESC(10001, res)
		return
	}
	if util.VerifyParamsStr(req.FormValue("did")) {
		result.RESC(21002, res)
		return
	}
	did,err := util.StrToInt32(req.FormValue("did"))
	log.Infof("str is: (%#v)" ,req.FormValue("did"))
	log.Infof("did(%#v),err(%#v)",did,err)
	if err != nil {
		result.RESC(10001, res)
		return
	}
	petinfo := &pb.PetInfoRequest{Source:sso.Source,Pid:pid,Uid:sso.Uid,Did:did}
	log.Info("GetPetInfo1_1-petinfo:", petinfo)
	//调用rpc
	petinfoRe := rpc.GetPetInfoBydid(petinfo)
	if petinfoRe.Code != 10000 {
		result.RESC(petinfoRe.Code, res)
		return
	}
	result.REST(&po.PetinfoPo{Pid: petinfoRe.Pid, Did: petinfoRe.Did, Avatar: petinfoRe.Avatar,
		Nickname: petinfoRe.Nickname, Breed: petinfoRe.Breed, Gender: petinfoRe.Gender, Birthday: petinfoRe.Birthday,
		Weight: petinfoRe.Weight, Somatotype: petinfoRe.Somatotype, Duration:petinfoRe.Duration, Brightness:petinfoRe.Brightness,
		CreateTime: petinfoRe.CreateTime, Permit: petinfoRe.Permit,Trains:petinfoRe.Trains}, res)
}
//获取批量宠物信息和设备录音
func GetPetInfos1_1(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("GetPetInfos1_1...")
	sso := util.GetContext(req)
	log.Info("GetPetInfos1_1-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	did,err := util.StrToInt32(req.FormValue("did"))
	if err != nil {
		result.RESC(10001, res)
		return
	}
	petInfo := &pb.PetInfoRequest{Source:sso.Source,Uid:sso.Uid,Did:did}
	log.Info("GetPetInfos1_1-petInfo:", petInfo)
	//调用rpc
	petInfoMapRe := rpc.GetPetInfosBydid(petInfo)
	if petInfoMapRe.Code != 10000 {
		result.RESC(petInfoMapRe.Code, res)
		return
	}
	var arr []po.PetinfoPo
	for _, v := range petInfoMapRe.Petinfos {
		arr = append(arr, po.PetinfoPo{Pid: v.Pid, Did: v.Did, Avatar: v.Avatar, Nickname: v.Nickname,
			Breed: v.Breed, Gender: v.Gender, Birthday: v.Birthday, Weight: v.Weight, Somatotype: v.Somatotype,
			Duration:v.Duration, Brightness:v.Brightness, CreateTime: v.CreateTime, Permit: v.Permit,Trains:v.Trains})
	}
	result.REST(arr, res)
}

/**
宠物训练信息
 */

//查询宠物训练信息
func GetPetTrainByPid(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("GetPetTrainByPid...")
	pidStr := params.ByName("pid")
	startStr := req.FormValue("start")
	endStr := req.FormValue("end")
	sso := util.GetContext(req)
	log.Info("GetPetTrainByPid-sso:", sso, ",pidStr:", pidStr,",startStr:",startStr,",endStr:",endStr)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	if util.VerifyParamsStr(pidStr) {
		result.RESC(21002, res)
		return
	}
	pid, err := util.StrToInt32(pidStr)
	var start, end int64
	if util.VerifyParamsStr(startStr,endStr) {
		end = util.GetNowTimeInt64()
		start = util.GetZeroTime(end)
	} else {
		start, err = util.StrToInt64(startStr)
		end, err = util.StrToInt64(endStr)
	}
	if err != nil {
		result.RESC(21002, res)
		return
	}
	if util.VerifyParamsUInt32(pid) {
		result.RESC(21002, res)
		return
	}
	if util.VerifyParamsUInt64(start,end) {
		result.RESC(21002, res)
		return
	}
	train := &pb.PetTrainRequest{Source:sso.Source,Uid:sso.Uid,Pid:pid,StartTime:start,EndTime:end}
	log.Info("GetPetTrainByPid-train:", train)
	//调用rpc
	trainRe := rpc.TrainsRpc(train, "GetPetTrainByPid")
	if trainRe.Code != 10000 {
		result.RESC(trainRe.Code, res)
		return
	}
	var dayTrains [][]*pb.PetTrainReply
	for _,v := range trainRe.SliceTrains {
		dayTrains = append(dayTrains,v.Trains)
	}
	result.REST(dayTrains, res)
}

//修改宠物训练信息
func UpdatePetTrainByPid(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("CounterPetTrainByPid...")
	pidStr := params.ByName("pid")
	idStr := params.ByName("id")
	sso := util.GetContext(req)
	log.Info("CounterPetTrainByPid-sso:", sso, ",pidStr:", pidStr, ",idStr:", idStr)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	if util.VerifyParamsStr(pidStr,idStr) {
		result.RESC(21002, res)
		return
	}
	pid, err := util.StrToInt32(pidStr)
	id, err := util.StrToInt32(idStr)
	if err != nil {
		result.RESC(21002, res)
		return
	}
	if util.VerifyParamsUInt32(pid,id) {
		result.RESC(21002, res)
		return
	}
	train := &pb.PetTrainRequest{}
	if util.GetHttpData(req, util.ReqMethodJson, &train) {
		result.RESC(21002, res)
		return
	}
	if train.Name != "" {
		if len(train.Name) > 100 {
			result.RESC(21002, res)
			return
		}
	}
	if train.Voice != "" {
		if len(train.Voice) > 100 {
			result.RESC(21002, res)
			return
		}
	}
	if train.Num < 0 || train.Num > 1000 {
		result.RESC(21002, res)
		return
	}
	train.Source=sso.Source; train.Pid=pid; train.Uid=sso.Uid; train.Id=id
	log.Info("CounterPetTrainByPid-train:", train)
	//调用rpc
	trainRe := rpc.TrainRpc(train, "UpdatePetTrainByPid")
	result.RESC(trainRe.Code, res)
}
//修改宠端设备训练信息
func UpdateDeviceTrainByPid(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("CounterPetTrainByPid...")
	pidStr := params.ByName("did")
	idStr := params.ByName("id")
	sso := util.GetContext(req)
	log.Info("CounterPetTrainByPid-sso:", sso, ",pidStr:", pidStr, ",idStr:", idStr)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	if util.VerifyParamsStr(pidStr,idStr) {
		result.RESC(21002, res)
		return
	}
	did, err := util.StrToInt32(pidStr)
	id, err := util.StrToInt32(idStr)
	if err != nil {
		result.RESC(21002, res)
		return
	}
	if util.VerifyParamsUInt32(did,id) {
		result.RESC(21002, res)
		return
	}
	train := &pb.DeviceTrainRequest{}
	if util.GetHttpData(req, util.ReqMethodJson, &train) {
		result.RESC(21002, res)
		return
	}
	if train.Name != "" {
		if len(train.Name) > 100 {
			result.RESC(21002, res)
			return
		}
	}
	if train.Voice != "" {
		if len(train.Voice) > 100 {
			result.RESC(21002, res)
			return
		}
	}
	if train.Num < 0 || train.Num > 1000 {
		result.RESC(21002, res)
		return
	}
	train.Source=sso.Source; train.Did=did; train.Uid=sso.Uid; train.Id=id
	log.Info("CounterPetTrainByPid-train:", train)
	//调用rpc
	trainRe := rpc.UpdateDeviceTrainByDid(train)
	result.RESC(trainRe.Code, res)
}

//计数宠物训练次数
func CounterPetTrainByPid(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("CounterPetTrainByPid...")
	pidStr := params.ByName("pid")
	idStr := params.ByName("id")
	log.Info("CounterPetTrainByPid-pidStr:", pidStr, ",idStr:", idStr)
	if util.VerifyParamsStr(pidStr,idStr) {
		result.RESC(21002, res)
		return
	}
	pid, err := util.StrToInt32(pidStr)
	id, err := util.StrToInt32(idStr)
	if err != nil {
		result.RESC(21002, res)
		return
	}
	if util.VerifyParamsUInt32(pid,id) {
		result.RESC(21002, res)
		return
	}
	sso := util.GetContext(req)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	log.Info("CounterPetTrainByPid-sso:", sso)
	train := &pb.PetTrainRequest{}
	nofund := util.GetHttpData(req, util.ReqMethodJson, &train)
	if nofund {
		result.RESC(21002, res)
		return
	}
	log.Info("CounterPetTrainByPid-train:", train)
	if train.Counter > 10 {
		result.RESC(21002, res)
		return
	}
	if train.Counter == 0 {
		train.Counter = 1
	}
	train.Source=sso.Source; train.Pid=pid; train.Uid=sso.Uid; train.Id=id
	log.Info("CounterPetTrainByPid-train:", train)
	//调用rpc
	trainRe := rpc.TrainRpc(train, "CounterPetTrainByPid")
	result.RESC(trainRe.Code, res)
}

/**
关联信息
 */
func RelevanceDevicePet(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("RelevanceDevicePet...")
	//todo 接收参数
	pidstr := params.ByName("pid")
	didstr := params.ByName("did")
	sso := util.GetContext(req)
	log.Info("RelevanceDevicePet-sso:", sso, ",pidstr:", pidstr, ",didstr:", didstr)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	//todo 校验参数
	if util.VerifyParamsStr(pidstr, didstr) {
		result.RESC(21002, res)
		return
	}
	pid, err := util.StrToInt32(pidstr)
	did, err := util.StrToInt32(didstr)
	if err != nil {
		result.RESC(10001, res)
		return
	}
	petInfo := &pb.PetInfoRequest{Source:sso.Source,Pid:pid,Uid:sso.Uid,Did:did}
	log.Info("RelevanceDevicePet-petinfo:", petInfo)
	//todo 调用rpc
	deviceRe := rpc.PetinfoRpc(petInfo, "SetDevicePet")
	result.RESC(deviceRe.Code, res)
}

func UnRelevanceDevicePet(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("UnRelevanceDevicePet...")
	pidstr := params.ByName("pid")
	didstr := params.ByName("did")
	sso := util.GetContext(req)
	log.Info("RelevanceDevicePet-sso:", sso, ",pidstr:", pidstr, ",didstr:", didstr)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	//todo 校验参数
	if util.VerifyParamsStr(pidstr, didstr) {
		result.RESC(21002, res)
		return
	}
	pid, err := util.StrToInt32(pidstr)
	did, err := util.StrToInt32(didstr)
	if err != nil {
		result.RESC(10001, res)
		return
	}
	petinfo := &pb.PetInfoRequest{Source:sso.Source, Pid:pid, Uid:sso.Uid, Did:did}
	log.Info("RelevanceDevicePet-petinfo:", petinfo)
	//todo 调用rpc
	deviceRe := rpc.PetinfoRpc(petinfo, "DeleteDevicePet")
	result.RESC(deviceRe.Code, res)
}

/*
获取品种信息
 */
func GetBreedInfos(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("GetBreedInfos...")
	typesStr := params.ByName("types")
	idStr := req.FormValue("id")
	numberStr := req.FormValue("number")
	sso := util.GetContext(req)
	log.Info("GetBreedInfos-sso:", sso, ",typesStr:", typesStr, ",idStr:", idStr, ",numberStr:", numberStr)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	if util.VerifyParamsStr(typesStr, numberStr) {
		result.RESC(21002, res)
		return
	}
	var types, id, number int32 = 0, 1, 0
	var err error
	if idStr != "" {
		id, err = util.StrToInt32(idStr)
	}
	types, err = util.StrToInt32(typesStr)
	number, err = util.StrToInt32(numberStr)
	if err != nil {
		result.RESC(21002, res)
		return
	}
	if util.VerifyParamsUInt32(types, id, number) {
		result.RESC(21002, res)
		return
	}
	if number > 60 {
		result.RESC(21002, res)
		return
	}
	files := &pb.FilesRequest{}
	files.Source = sso.Source
	files.Id = id
	files.Types = types
	files.Number = number
	log.Info("CounterPetTrainByPid-files:", files)
	//调用rpc
	filesRe := rpc.FilesRpc(files, "GetBreeds")
	if filesRe.Code != 10000 {
		result.RESC(filesRe.Code, res)
		return
	}
	var breeds po.BreedInfoPoSlice
	for _, v := range filesRe.GetFiles() {
		breeds = append(breeds, &po.BreedInfoPo{Id: v.Id, Name: v.Name, Types: v.Types, Address: v.Address})
	}
	sort.Stable(breeds)
	result.REST(breeds, res)
}

func GetBreedInfos1_1(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("GetBreedInfos...")
	typesStr := params.ByName("types")
	idStr := req.FormValue("id")
	numberStr := req.FormValue("number")
	sso := util.GetContext(req)
	log.Info("GetBreedInfos-sso:", sso, ",typesStr:", typesStr, ",idStr:", idStr, ",numberStr:", numberStr)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	if util.VerifyParamsStr(typesStr, numberStr) {
		result.RESC(21002, res)
		return
	}
	var types, id, number int32 = 0, 1, 0
	var err error
	if idStr != "" {
		id, err = util.StrToInt32(idStr)
	}
	types, err = util.StrToInt32(typesStr)
	number, err = util.StrToInt32(numberStr)
	if err != nil {
		result.RESC(21002, res)
		return
	}
	if util.VerifyParamsUInt32(types, id, number) {
		result.RESC(21002, res)
		return
	}
	if number > 60 {
		result.RESC(21002, res)
		return
	}
	files := &pb.FilesRequest{}
	files.Source = sso.Source
	files.Id = id
	files.Types = types
	files.Number = number
	log.Info("CounterPetTrainByPid-files:", files)
	//调用rpc
	filesRe := rpc.FilesRpc(files, "GetBreeds")
	if filesRe.Code != 10000 {
		result.RESC(filesRe.Code, res)
		return
	}
	var breeds po.BreedInfoPoSlice
	for _, v := range filesRe.GetFiles() {
		breeds = append(breeds, &po.BreedInfoPo{Id: v.Id, NameCh: v.NameCh, NameEn: v.NameEn, Types: v.Types, Address: v.Address})
	}
	sort.Stable(breeds)
	result.REST(breeds, res)
}
