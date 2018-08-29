package controller

import (
	log "github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"petfone-http/pb"
	"petfone-http/pb/api"
	"petfone-http/po"
	"petfone-http/result"
	"petfone-http/rpc"
	"petfone-http/util"
	"petfone-http/core"
	. "petfone-http/thirdPartyServer"
	"sort"
	"encoding/base64"
	"github.com/disintegration/imaging"
	"strings"
	"net/url"
	"bytes"
	"image"
	"github.com/nfnt/resize"
	"image/jpeg"
	"mime/multipart"
	"os"
)

/**
设备
*/
type deviceRequest struct {
	Source               string      `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Did                  int32       `protobuf:"varint,2,opt,name=did,proto3" json:"did,omitempty"`
	Uid                  int32       `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Pid                  int32       `protobuf:"varint,4,opt,name=pid,proto3" json:"pid,omitempty"`
	Touid                int32       `protobuf:"varint,5,opt,name=touid,proto3" json:"touid,omitempty"`
	Sn                   string      `protobuf:"bytes,6,opt,name=sn,proto3" json:"sn,omitempty"`
	DeviceMac            string      `protobuf:"bytes,7,opt,name=deviceMac,proto3" json:"deviceMac,omitempty"`
	DeviceName           string      `protobuf:"bytes,8,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	DeviceVersion        string      `protobuf:"bytes,9,opt,name=deviceVersion,proto3" json:"deviceVersion,omitempty"`
	SoftwareVersion      string      `protobuf:"bytes,10,opt,name=softwareVersion,proto3" json:"softwareVersion,omitempty"`
	//Permit               DevPermit   `protobuf:"varint,11,opt,name=permit,proto3,enum=pb.DevPermit" json:"permit,omitempty"`
	//Types                DeviceTypes `protobuf:"varint,12,opt,name=types,proto3,enum=pb.DeviceTypes" json:"types,omitempty"`
	Isdel                uint32      `protobuf:"varint,13,opt,name=isdel,proto3" json:"isdel,omitempty"`
	ShareUrl             string      `protobuf:"bytes,14,opt,name=shareUrl,proto3" json:"shareUrl,omitempty"`
	Code                 int32       `protobuf:"varint,15,opt,name=code,proto3" json:"code,omitempty"`
	LedModel             int32       `protobuf:"varint,16,opt,name=ledModel,proto3" json:"ledModel,omitempty"`
	LedColor             int32       `protobuf:"varint,17,opt,name=ledColor,proto3" json:"ledColor,omitempty"`
	LedLight             *int32       `protobuf:"varint,18,opt,name=ledLight,proto3" json:"ledLight,omitempty"`
	LedState             *int32       `protobuf:"varint,19,opt,name=ledState,proto3" json:"ledState,omitempty"`
	AudioId              int32       `protobuf:"varint,20,opt,name=audioId,proto3" json:"audioId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

var backPath = "./temp/backUpData.log"

//校验设备
func VerificationDevices(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("VerificationDevices...")
	snStr := params.ByName("did")
	sso := util.GetContext(req)
	log.Info("SetDeviceP-sso:", sso, ",snStr:", snStr)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	if util.VerifyParamsStr(snStr) {
		result.RESC(21001, res)
		return
	}
	if util.VerifySN(snStr) {
		result.RESC(21002, res)
		return
	}
	devices := new(pb.DeviceRequest)
	devices.Source = sso.Source
	devices.Uid = sso.Uid
	devices.Sn = snStr
	//调用rpc
	deviceRe := rpc.DevicesRpc(devices, "VerificationDeviceBySn")
	result.RESC(deviceRe.Code, res)
}

//添加设备
func SetDevices(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("SetDevices...")
	snStr := params.ByName("did")
	if util.VerifyParamsStr(snStr) || util.VerifySN(snStr) {
		result.RESC(21002, res)
		return
	}
	sso := util.GetContext(req)
	log.Info("SetDeviceP-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	devices := &pb.DeviceRequest{}
	if util.GetHttpData(req, util.ReqMethodJson, &devices) {
		result.RESC(21002, res)
		return
	}
	if util.VerifyParamsStr(devices.DeviceName, devices.DeviceVersion, devices.DeviceMac, devices.SoftwareVersion) {
		result.RESC(21001, res)
		return
	}
	devices.Source = sso.Source
	devices.Uid = sso.Uid
	devices.Sn = snStr
	//调用rpc
	deviceRe := rpc.DevicesRpc(devices, "SetDeviceBySn")
	if deviceRe.Code != 10000 {
		result.RESC(deviceRe.Code, res)
		return
	}
	result.REST(&po.DevicePo{Did: deviceRe.Did, Types: deviceRe.Types, DeviceMac:devices.DeviceMac,
		DeviceName:devices.DeviceName,DeviceVersion:devices.DeviceVersion,SoftwareVersion:devices.SoftwareVersion,
		LedColor:deviceRe.LedColor,LedLight:deviceRe.LedLight,LedState:deviceRe.LedState,LedModel:deviceRe.LedModel,AudioId:deviceRe.AudioId}, res)
}

//解绑设备
func DeleteDevice(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("DeleteDevice...")
	didStr := params.ByName("did")
	sso := util.GetContext(req)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	log.Info("DeleteDevice-sso:", sso, ",didStr:", didStr)
	did, err := util.StrToInt32(didStr)
	if err != nil {
		result.RESC(21002, res)
		return
	}
	devices := &pb.DeviceRequest{Source:sso.Source,Uid:sso.Uid,Did:did}
	log.Info("DeleteDevice-devices:", devices)
	//调用rpc
	deviceRe := rpc.DevicesRpc(devices, "DeleteDeviceByDid")
	if deviceRe.Code == 33012 {
		result.RESC(21005, res)
		return
	}
	result.RESC(deviceRe.Code, res)
}

//修改设备
func UpdateDeviceInfo(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("UpdateDevice...")
	didStr := params.ByName("did")
	sso := util.GetContext(req)
	log.Info("UpdateDeviceP-sso:", sso, ",didStr:", didStr)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	devicesRequest := &deviceRequest{}
	devices := &pb.DeviceRequest{}
	log.Infof("&&& device(%#v)",devicesRequest)
	if util.GetHttpData(req, util.ReqMethodJson, &devicesRequest) {
		result.RESC(21002, res)
		return
	}
	log.Infof("devices(%#v)", devicesRequest)
	if devicesRequest.AudioId>3 || devicesRequest.AudioId<0 || devicesRequest.LedModel<0 || devicesRequest.LedModel>3 {
		log.Infof("参数有误！（%#v）",devicesRequest)
		result.RESC(21002, res)
		return
	}
	if devicesRequest.LedLight != nil {
		if *devicesRequest.LedLight == 0 {
			devices.LedLight = -1
		}else {
			devices.LedLight = *devicesRequest.LedLight
		}
	}
	if devicesRequest.LedState != nil {
		if *devicesRequest.LedState == 0 {
			devices.LedState = -1
		}else {
			devices.LedState = *devicesRequest.LedState
		}
	}
	var did int32
	var err error
	if did, err = util.StrToInt32(didStr); err != nil {
		result.RESC(10001, res)
		return
	}
	devices.Source = sso.Source
	devices.Uid = sso.Uid
	devices.Did = did
	devices.AudioId = devicesRequest.AudioId
	devices.LedModel = devicesRequest.LedModel
	devices.LedColor = devicesRequest.LedColor
	devices.SoftwareVersion = devicesRequest.SoftwareVersion
	log.Info("UpdateDeviceP-deviceJson:", devices)
	//调用rpc
	deviceRe := rpc.DevicesRpc(devices, "UpdateDeviceByDid")
	result.RESC(deviceRe.Code, res)
}

//获取设备
func GetDevices(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("GetDevices...")
	sso := util.GetContext(req)
	log.Info("GetDevicesP-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	device := &pb.DeviceRequest{}
	device.Source = sso.Source
	device.Uid = sso.Uid
	log.Info("GetDevicesP-deviceJson:", device)
	//调用rpc
	deviceMapRe := rpc.DeviceMapRpc(device, "GetDevicesByUid")
	log.Info("deviceMapRe:", len(deviceMapRe.Devices))
	if deviceMapRe.Code != 10000 {
		result.RESC(deviceMapRe.Code, res)
		return
	}
	var arr []po.DevicePo
	for _, v := range deviceMapRe.Devices {
		log.Infof("v(%#v)",v)
		arr = append(arr, po.DevicePo{Did: v.Did, Pid: v.Pid, Types:v.Types,
			DeviceName: v.DeviceName, DeviceVersion: v.DeviceVersion, DeviceMac: v.DeviceMac,
			SoftwareVersion: v.SoftwareVersion, Permit: v.Permit,
			LedModel:v.LedModel,
			LedColor:v.LedColor,
			LedState:v.LedState,
			LedLight:v.LedLight,
			AudioId:v.AudioId,
		})
	}
	result.REST(arr, res)
}

/**
设备数据
 */

//上报运动信息
func SetExerciseData(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("SetExerciseData...")
	pidStr := params.ByName("pid")
	sso := util.GetContext(req)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	log.Info("SetDeviceP-sso:", sso, ",pidStr:", pidStr)
	exercise := &pb.ExerciseDataRequest{}
	if util.GetHttpData(req, util.ReqMethodJson, &exercise) {
		result.RESC(21002, res)
		return
	}
	//todo 校验参数
	if util.VerifyParamsStr(pidStr) {
		result.RESC(21001, res)
		return
	}
	pid, err := util.StrToInt32(pidStr)
	if err != nil {
		result.RESC(21002, res)
		return
	}
	if util.VerifyParamsFloat32(exercise.Pcals) {
		result.RESC(21002, res)
		return
	}
	if exercise.Pcoordinates == nil || len(exercise.Pcoordinates) == 0 || exercise.Pcoordinates[0].Nowtime == 0 {
		result.RESC(21002, res)
		return
	}
	if exercise.StartTime > 0 {
		nowTime := util.GetNowTime().Unix()
		if util.Int64DifferenceAbs(nowTime,exercise.StartTime) > 86400 {
			result.RESC(21002, res)
			return
		}
	}
	exercise.Source = sso.Source
	exercise.Uid = sso.Uid
	exercise.Pid = pid
	log.Info("SetDeviceP-exercise:", len(exercise.Pcoordinates))
	//todo 调用rpc
	deviceRe := rpc.ExerciseRpc(exercise, "SetExerciseData")
	result.RESC(deviceRe.Code, res)
}
func SetExerciseData1(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("SetExerciseData1...")
	sso := util.GetContext(req)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	exercise := &pb.ExerciseDataRequest{}
	if util.GetHttpData(req, util.ReqMethodJson, &exercise) {
		result.RESC(21002, res)
		return
	}
	log.Info("SetExerciseData1-sso:", sso,",exercise:",exercise.Pid,exercise.Pdid)
	if util.VerifyParamsUInt32(exercise.Pid,exercise.Pdid) {
		result.RESC(21001, res)
		return
	}
	if util.VerifyParamsFloat32(exercise.Pcals) {
		result.RESC(21002, res)
		return
	}
	if exercise.Pcoordinates == nil || len(exercise.Pcoordinates) == 0 || exercise.Pcoordinates[0].Nowtime == 0 {
		result.RESC(21002, res)
		return
	}
	nowTime := util.GetNowTime().Unix()
	if exercise.StartTime <= 0 || util.Int64DifferenceAbs(nowTime,exercise.StartTime) > 86400 {
		result.RESC(21002, res)
		return
	}
	if exercise.ImageInfo == nil || exercise.ImageInfo.ImageEnCode == "" || exercise.ImageInfo.Md5 == "" || exercise.ImageInfo.Name == "" {
		log.Info("SetExerciseData1 exercise.ImageInfo empty")
		result.RESC(21002, res)
		return
	}
	suffix,err := imaging.FormatFromFilename(exercise.ImageInfo.Name)
	if err == imaging.ErrUnsupportedFormat {
		log.Info("SetExerciseData1 name err:", err)
		result.RESC(21002, res)
		return
	}
	base64Str, _ := url.QueryUnescape(exercise.ImageInfo.ImageEnCode)
	bs,_ := base64.StdEncoding.DecodeString(base64Str)
	if util.CalMd5(bs) != exercise.ImageInfo.Md5 {
		log.Info("SetExerciseData1 Md5:",exercise.ImageInfo.Md5)
		result.RESC(21002, res)
		return
	}
	imageBuff := &bytes.Buffer{}
	imageBuff.Write(bs)
	img, _, err := image.Decode(imageBuff)
	if err != nil {
		log.Info("SetExerciseData1 imaging.Decode err:", err)
		result.RESC(21002, res)
		return
	}
	x := img.Bounds().Size().X
	y := img.Bounds().Size().Y
	compressImage := resize.Thumbnail(uint(x), uint(y), img, resize.Bilinear)
	imageBuff.Reset()
	jpeg.Encode(imageBuff,compressImage,nil)
	compressMd5 := util.CalMd5(imageBuff.Bytes())
	bodyBuff := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuff)
	fileName := compressMd5+"."+strings.ToLower(suffix.String())
	fileWriter, err := bodyWriter.CreateFormFile("file", fileName)
	if err != nil {
		log.Error("SetExerciseData1 CreateFormFile err:",err)
		result.RESC(21002, res)
		return
	}
	fileWriter.Write(imageBuff.Bytes())
	bodyWriter.Close()
	fid, err := SendFileHttp(core.ConstStr.FileServer,bodyWriter.FormDataContentType(),bodyBuff)
	if err != nil || fid == "" {
		log.Error("SetExerciseData1 SendFile err:", err)
		log.Error("SetExerciseData1 fileName:", fileName)
		backUpData, _ := os.OpenFile(backPath,os.O_CREATE|os.O_APPEND,0666)
		defer backUpData.Close()
		backUpData.Write([]byte(exercise.String()+"\n"))
		result.RESC(10001, res)
		return
	}
	exercise.Source = sso.Source
	exercise.Uid = sso.Uid
	exercise.ImageInfo.Url = fid
	exercise.ImageInfo.Name = fileName
	exercise.ImageInfo.Size = int32(imageBuff.Len())
	exercise.ImageInfo.Width = int32(x)
	exercise.ImageInfo.Height = int32(y)
	log.Info("SetExerciseData1-exercise:", len(exercise.Pcoordinates))
	log.Info("aaaaaaaaaaaa:",exercise.ImageInfo.Md5)
	log.Info("aaaaaaaaaaaa:",compressMd5,fid)
	deviceRe := rpc.ExerciseRpc(exercise, "SetExerciseData")
	result.RESC(deviceRe.Code, res)
}

//获取运动历史信息
func GetExerciseData(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("GetExerciseData...")
	pidStr := params.ByName("pid")
	startStr := req.FormValue("start")
	endStr := req.FormValue("end")
	sso := util.GetContext(req)
	log.Info("GetExerciseData-sso:", sso, ",pidStr:", pidStr, ",startStr:", startStr, ",endStr:", endStr)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	//todo 校验参数
	if util.VerifyParamsStr(pidStr, startStr, endStr) {
		result.RESC(21001, res)
		return
	}
	pid, err := util.StrToInt32(pidStr)
	end, err := util.StrToInt64(endStr)
	start, err := util.StrToInt64(startStr)
	if err != nil {
		result.RESC(21001, res)
		return
	}
	exercise := &pb.ExerciseDataRequest{Source: sso.Source, Uid: sso.Uid, Pid: pid, StartTime: start, EndTime: end}
	log.Info("GetExerciseData-exercise:", exercise)
	//todo 调用rpc
	exersRe := rpc.MotionDataMapRpc(exercise, "GetExerciseDataPet")
	if exersRe.Code != 10000 {
		result.RESC(exersRe.Code, res)
		return
	}
	log.Info("GetExerciseDatasPet-exersRe:", len(exersRe.Data))
	var dayExerciseDataPos []*po.DayExerciseDataPo
	for _,v := range exersRe.Data {
		var coordinates po.CoordinateSlice
		for _,v := range v.Coordinates {
			coordinates = append(coordinates,&po.Coordinate{NowTime:v.NowTime, Longitude:v.Longitude, Latitude:v.Latitude, State:v.State,})
		}
		sort.Stable(coordinates)
		dayExerciseDataPos = append(dayExerciseDataPos,&po.DayExerciseDataPo{DayTime:v.DayTime,
			CardioTimes:v.CardioTimes, StrenuousTimes:v.StrenuousTimes, Steps:v.Steps, Cals:v.Calorie, Coordinates:coordinates,})
	}
	result.REST(&po.DaysExerciseDataPo{Code:exersRe.Code,Pid:exersRe.Pid,Pdid:exersRe.Pdid,Data:dayExerciseDataPos}, res)
}

func GetExerciseData1(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("GetExerciseData1.1...")
	pidStr := req.FormValue("pid")
	startStr := req.FormValue("start")
	endStr := req.FormValue("end")
	sso := util.GetContext(req)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	log.Infof("GetExerciseData1.1 pid:%v,start:%v,end:%v", pidStr,startStr,endStr)
	//todo 校验参数
	if util.VerifyParamsStr(pidStr,startStr,endStr) {
		result.RESC(21001, res)
		return
	}
	pid, err := util.StrToInt32(pidStr)
	start, err := util.StrToInt64(startStr)
	end, err := util.StrToInt64(endStr)
	if err != nil {
		result.RESC(21001, res)
		return
	}
	exercise := &pb.ExerciseDataRequest{Source:sso.Source,Uid:sso.Uid,Pid:pid, StartTime:start,EndTime:end}
	log.Info("GetExerciseData1.1-exercise:", exercise)
	//todo 调用rpc
	exersRe := rpc.MotionDataMapRpc(exercise, "GetMotionDataPetByTime")
	if exersRe.Code != 10000 {
		result.RESC(exersRe.Code, res)
		return
	}
	log.Info("GetExerciseData1.1-exersRe:", len(exersRe.Data))
	var days []*po.DayMotionDataPo
	for _,v := range exersRe.Data {
		var records []*po.Record
		for _,v := range v.Records {
			records = append(records,&po.Record{Pdid:v.Pdid, TimeRecord:v.TimeRecord, Steps:v.Steps, Calorie:v.Calorie,
				CardioTimes:v.CardioTimes, StrenuousTimes:v.StrenuousTimes,
				CardioDurationMinute:v.CardioDurationMinute,StrenuousDurationMinute:v.StrenuousDurationMinute,
				ImageInfo:&po.ImageInfo{Url:core.ConstStr.ImageServer+v.ImageInfo.Url,ViewUrl:core.ConstStr.ImageServer+v.ImageInfo.Url,Name:v.ImageInfo.Name,
					Size:v.ImageInfo.Size,Width:v.ImageInfo.Width,Height:v.ImageInfo.Height,FileUrl:core.ConstStr.FileServer+v.ImageInfo.Url}})
		}
		days = append(days, &po.DayMotionDataPo{DayTime:v.DayTime, Records:records, StepsTotal:v.Steps, CalorieTotal:v.Calorie,
			CardioTimesTotal:v.CardioTimes,StrenuousTimesTotal:v.StrenuousTimes,
			CardioDurationMinuteTotal:v.CardioDurationMinute,StrenuousDurationMinuteTotal:v.StrenuousDurationMinute})
	}
	result.REST(days, res)
}

//上报设备数据
func UploadTwinsAgent(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("UploadDeviceAgent...")
	dev := params.ByName("dev")
	if dev != "dev" {
		result.RESC(21002, res)
		return
	}
	sso := util.GetContext(req)
	log.Info("UploadDeviceAgent-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	deviceReq := &pb.DeviceRequest{}
	if util.GetHttpData(req, util.ReqMethodJson, &deviceReq) {
		result.RESC(21002, res)
		return
	}
	if util.VerifyParamsStr(deviceReq.Input) {
		result.RESC(21002, res)
		return
	}
	if deviceReq.Did <= 0  {
		result.RESC(21002, res)
		return
	}
	deviceReq.Source = sso.Source
	deviceReq.Uid = sso.Uid
	//todo 调用rpc
	deviceRe := rpc.DevicesRpc(deviceReq, "GetDeviceSn")
	if deviceRe.Code != 10000 {
		result.RESC(deviceRe.Code, res)
		return
	}
	go func() {
		//todo 解析input调用rpc
		var aJson map[string]interface{}
		util.Json.Unmarshal([]byte(deviceReq.Input),&aJson)
		aJson["sn"] = deviceRe.Sn
		aJson["source"] = sso.Source
		aJson["uid"] = sso.Uid
		input, err := util.Json.MarshalToString(&aJson)
		if err != nil {
			log.Error("UploadDeviceAgent MarshalToString:", err)
		}
		agentReq := &api.AddTwinsAgentRequest{Reported:input}
		rpc.TwinsRpc(agentReq,"AddTwinsAgent")
	}()
	result.RESC(10000, res)
}