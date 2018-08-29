package service

import (
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"petfone-rpc/db"
	"petfone-rpc/pb"
	"petfone-rpc/util"
	"petfone-rpc/core"
	"github.com/jinzhu/gorm"
	"time"
)

type DeviceRpc struct {
}

//验证设备
func (this *DeviceRpc) VerificationDeviceBySn(ctx context.Context, req *pb.DeviceRequest) (*pb.DeviceReply, error) {
	log.Info("VerificationDeviceBySn-req:", req)
	if util.VerifyParamsStr(req.GetSource())||"AgIDAA==" != req.GetSource() {
		return &pb.DeviceReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetUid()) {
		return &pb.DeviceReply{Code: util.Params_err_empty}, nil
	}
	device := &db.DevicePo{Sn:req.Sn}
	err := device.GetDeviceBySn()
	log.Info("VerificationDeviceBySn-device:", device)
	if err != nil {
		if core.ConstStr.NotFound == err.Error() {
			return &pb.DeviceReply{Code: 10000}, nil
		}
		log.Info("VerificationDeviceBySn-err:", err)
		return &pb.DeviceReply{Code: util.System_error}, nil
	}
	//todo 查询设备关系
	userDevicePo := &db.UserDevicePo{Did:device.Did}
	userDevicePos, err := userDevicePo.GetUsersDeviceDB()
	log.Info("userDevicePos-len:",len(userDevicePos))
	if err != nil {
		log.Info("VerificationDeviceBySn-err:", err)
		return &pb.DeviceReply{Code: 10001}, nil
	}
	if len(userDevicePos) == 0 {
		return &pb.DeviceReply{Code: 10000}, nil
	}
	for _, v := range userDevicePos {
		if v.Uid == req.Uid {
			return &pb.DeviceReply{Code: 10000}, nil
		}
	}
	return &pb.DeviceReply{Code: 33012}, nil
}

//绑定设备
func (this *DeviceRpc) SetDeviceBySn(ctx context.Context, req *pb.DeviceRequest) (*pb.DeviceReply, error) {
	log.Info("SetDeviceBySn-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.DeviceReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetUid()) {
		return &pb.DeviceReply{Code: util.Params_err_empty}, nil
	}
	dbc := core.MysqlClient.Begin()
	nowTime := util.GetNowTime()
	devicePo := &db.DevicePo{Sn:req.Sn,CreationTime:nowTime, UpdateTime:nowTime, DataState:1}
	userDevicePo := &db.UserDevicePo{Uid:req.Uid, CreationTime:nowTime, UpdateTime:nowTime, DataState:1}
	err := devicePo.GetDeviceBySn()
	if err != nil && core.ConstStr.NotFound != err.Error() {
		dbc.Rollback()
		log.Info("VerificationDeviceBySn-err:", err)
		return &pb.DeviceReply{Code: util.System_error}, nil
	}
	//todo 查询用户的人端设备
	userDevicePos, err := userDevicePo.GetUserDevicesDB()
	if err != nil && core.ConstStr.NotFound != err.Error() {
		dbc.Rollback()
		log.Info("VerificationDeviceBySn-err:", err)
		return &pb.DeviceReply{Code: util.System_error}, nil
	}
	for _,v := range userDevicePos {
		if v.Types == 1 && req.Types == 1 {
			dbc.Rollback()
			return &pb.DeviceReply{Code: util.User_permission}, nil
		}
	}
	//todo 如果不存在，则创建
	if devicePo.Did == 0 {
		devicePo.Uid	=	req.Uid
		devicePo.Types = req.Types
		devicePo.DeviceName = req.DeviceName
		devicePo.DeviceVersion = req.DeviceVersion
		devicePo.DeviceMac = req.DeviceMac
		devicePo.SoftwareVersion = req.SoftwareVersion
		devicePo.LedModel = util.LedModel
		devicePo.LedState = util.LedState
		devicePo.LedLight = util.LedLight
		devicePo.LedColor = util.LedColor
		devicePo.AudioId = util.AudioId
		err = devicePo.SetDeviceDB(dbc)
		if err != nil {
			dbc.Rollback()
			log.Info("err:", err)
			return &pb.DeviceReply{Code: 10001}, nil
		}
		userDevicePo.Did = devicePo.Did
		// todo modifyz xiaorx 2018\7\4 	创建宠端设备录音关系表
		//err := db.Transaction(core.MysqlClient, func(tx *gorm.DB) error {
		//	deviceTrainPo := &db.DeviceTrainPo{Did:devicePo.Did,CreationTime:nowTime,UpdateTime:nowTime,DataState:1,Num:10}
		//	for index,value := range core.Names {
		//		log.Infof("创建宠端设备录音关系表：index(%#v)value(%#v)",index,value)
		//		deviceTrainPo.Name = value
		//		deviceTrainPo.Voice = core.Voices[index]
		//		num := deviceTrainPo.Create(dbc)
		//		if  num != 1 {
		//			dbc.Rollback()
		//			return	errors.New("插入宠端设备训练录音失败")
		//		}
		//	}
		//	return nil
		//})
		//if err != nil {
		//	dbc.Rollback()
		//	log.Errorf("SetDeviceBySn...%#v",err.Error())
		//	return &pb.DeviceReply{Code: util.Mysql_err}, nil
		//}
	} else {
		//todo 如果已存在，并且该did已解绑
		userDevicePo.Did = devicePo.Did
		log.Info("SetDeviceBySn-userDevicePo:", userDevicePo)
		userDevicePos, err := userDevicePo.GetUsersDeviceDB()
		if err != nil {
			dbc.Rollback()
			log.Info("SetDeviceBySn-err:", err)
			return &pb.DeviceReply{Code: 10001}, nil
		}
		if len(userDevicePos) > 0 {
			dbc.Rollback()
			return &pb.DeviceReply{Code: 33012}, nil
		}
	}
	userDevicePo.Types = req.Types
	//todo 绑定用户
	if userDevicePo.SetUserDeviceDB(dbc) != 1 {
		dbc.Rollback()
		log.Info("SetDeviceBySn-SetUserDeviceDB-err:数据库插入错误")
		return &pb.DeviceReply{Code: 10001}, nil
	}
	dbc.Commit()
	return &pb.DeviceReply{Code: 10000, Did:devicePo.Did,Types:req.Types,LedModel:int32(devicePo.LedModel),
	LedState:int32(devicePo.LedState),LedLight:int32(devicePo.LedLight),LedColor:int32(devicePo.LedColor),AudioId:int32(devicePo.AudioId)}, nil
}

//解绑设备
func (this *DeviceRpc) DeleteDeviceByDid(ctx context.Context, req *pb.DeviceRequest) (*pb.DeviceReply, error) {
	log.Info("DeleteDeviceByDid-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.DeviceReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetUid(), req.GetDid()) {
		return &pb.DeviceReply{Code: util.Params_err_empty}, nil
	}
	//todo 查询设备权限
	userDevicePo := &db.UserDevicePo{Uid:req.Uid, Did:req.Did, DataState:1}
	userDevicePos, err := userDevicePo.GetUsersDeviceDB()
	log.Info("DeleteDeviceByDid-userDevicePos:", len(userDevicePos))
	if err != nil {
		log.Info("DeleteDeviceByDid-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.DeviceReply{Code: 33012}, nil
		}
		return &pb.DeviceReply{Code: 10001}, nil
	}
	// todo
	if len(userDevicePos) != 1 {
		return &pb.DeviceReply{Code: 33012}, nil
	}
	if userDevicePos[0].Permit != 0 || userDevicePos[0].Uid != req.Uid {
		return &pb.DeviceReply{Code: 33012}, nil
	}
	//TODO 查询与宠物的关系
	devicePetPo := &db.DevicePetPo{Did:req.Did}
	err = devicePetPo.GetDevicePetsDB()
	log.Info("DeleteDeviceByDid-devicePetPo:",devicePetPo)
	if err != nil && err.Error() !=  core.ConstStr.NotFound {
		log.Info("DeleteDeviceByDid-err:", err)
		return &pb.DeviceReply{Code: 10001}, nil
	}
	if devicePetPo.Pid != 0 {
		return &pb.DeviceReply{Code: util.User_manipulate}, nil
	}
	//todo 修改关系表
	nowTime := util.GetNowTime()
	dbc := core.MysqlClient.Begin()
	userDevicePo.UpdateTime = nowTime
	num := userDevicePo.DeleteUserDeviceDB(dbc)
	if num == 0 {
		dbc.Rollback()
		return &pb.DeviceReply{Code: 33012}, nil
	}
	devicePo := &db.DevicePo{Did:req.Did,UpdateTime:nowTime,}
	err = devicePo.UpdateDeviceDB(dbc)
	if err != nil {
		dbc.Rollback()
		log.Info("DeleteDeviceByDid-err:", err)
		return &pb.DeviceReply{Code: 10001}, nil
	}
	dbc.Commit()
	// 修改设备的默认颜色
	device := db.DevicePo{
		Did:req.Did,
		LedModel:util.LedModel,
		LedColor:util.LedColor,
		LedLight:util.LedLight,
		LedState:util.LedState,
		AudioId:util.AudioId,
		UpdateTime:time.Now(),
	}
	err = db.Transaction(core.MysqlClient, func(tx *gorm.DB) error {
		device.UpdateDeviceByDid(core.MysqlClient)
		return nil
	})
	if err != nil {
		log.Errorf("DeleteDeviceByDid...修改设备信息为默认值失败！")
	}
	return &pb.DeviceReply{Code: 10000}, nil
}

//修改设备
func (this *DeviceRpc) UpdateDeviceByDid(ctx context.Context, req *pb.DeviceRequest) (*pb.DeviceReply, error) {
	log.Infof("UpdateDeviceByDid...req:(%#v)",*req)
	if req.Did == 0 {
		log.Errorf("UpdateDeviceByDid...参数错误！")
		return &pb.DeviceReply{Code: util.User_params_err}, nil
	}
	device := db.DevicePo{
		Did:req.Did,
		LedModel:int8(req.LedModel),
		LedColor:req.LedColor,
		LedLight:int8(req.LedLight),
		LedState:int8(req.LedState),
		AudioId:int8(req.AudioId),
		UpdateTime:time.Now(),
		SoftwareVersion:req.SoftwareVersion,
	}
	err := db.Transaction(core.MysqlClient, func(tx *gorm.DB) error {
		device.UpdateDeviceByDid(core.MysqlClient)
		return nil
	})
	if err != nil {
		log.Errorf("UpdateDeviceByDid...修改设备信息失败！")
		return &pb.DeviceReply{Code: util.Mysql_err}, nil
	}
	return &pb.DeviceReply{Code: util.Success}, nil
}

//获取设备
func (this *DeviceRpc) GetDevicesByDid(ctx context.Context, req *pb.DeviceRequest) (*pb.DeviceReply, error) {
	log.Info("GetDevicesByUid-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.DeviceReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetUid()) {
		return &pb.DeviceReply{Code: util.Params_err_empty}, nil
	}
	userDevicePo := &db.UserDevicePo{Uid:req.Uid, Did:req.Did}
	err := userDevicePo.GetUserDeviceDB()
	if err != nil {
		log.Info("GetDevicesByDid-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.DeviceReply{Code: 33012}, nil
		}
		return &pb.DeviceReply{Code: 10001}, nil
	}
	devicePo := &db.DevicePo{Did:req.Did}
	err = devicePo.GetDeviceByDid()
	if err != nil {
		log.Info("err", err)
		return &pb.DeviceReply{Code: 10001}, nil
	}
	//todo 查询设备关系信息
	devicePetPo := &db.DevicePetPo{Did: req.Did}
	err = devicePetPo.GetDevicePetDB()
	if err != nil {
		return &pb.DeviceReply{Code: 33012}, nil
	}
	log.Infof("...............devicePo(%#v)", *devicePo)
	// todo 为了适配APP端
	if devicePo.LedState == -1 {
		devicePo.LedState = 0
	}
	if devicePo.LedLight == -1 {
		devicePo.LedLight = 0
	}
	return &pb.DeviceReply{Code: 10000, Did: devicePo.Did, Types: devicePo.Types, Pid: devicePetPo.Pid,
		DeviceName:devicePo.DeviceName,DeviceVersion:devicePo.DeviceVersion,
		DeviceMac:devicePo.DeviceMac,SoftwareVersion:devicePo.SoftwareVersion,
		LedModel:int32(devicePo.LedModel),
		LedColor:devicePo.LedColor,
		LedLight:int32(devicePo.LedLight),
		LedState:int32(devicePo.LedState),
		AudioId:int32(devicePo.AudioId)}, nil
}

//批量获取设备
func (this *DeviceRpc) GetDevicesByUid(ctx context.Context, req *pb.DeviceRequest) (*pb.BatchDeviceRe, error) {
	log.Info("GetDevicesByUid-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.BatchDeviceRe{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetUid()) {
		return &pb.BatchDeviceRe{Code: util.Params_err_empty}, nil
	}
	//todo 查询用户的设备
	userDevicePo := &db.UserDevicePo{Uid:req.Uid, Did:req.Did}
	userDevicePos, err := userDevicePo.GetUserDevicesDB()
	log.Info("userDevicePos:", userDevicePos)
	if err != nil {
		log.Info("GetDevicesByUid-err:", err)
		return &pb.BatchDeviceRe{Code: 10001}, nil
	}
	if len(userDevicePos) == 0 {
		return &pb.BatchDeviceRe{Code: 10000}, nil
	}
	var dids []int32
	for _, v := range userDevicePos {
		dids = append(dids,v.Did)
	}
	//todo 查询设备信息
	devicePo := &db.DevicePo{}
	devicePos, err := devicePo.GetDevicesDB(dids)
	if err != nil {
		log.Info("err:", err)
		return &pb.BatchDeviceRe{Code: 10001}, nil
	}
	log.Info("devicePos", devicePos)
	//todo 查询设备与宠物关系
	devicePetPo := &db.DevicePetPo{}
	devicePetPos, err := devicePetPo.GetDevicesPetsDidsDB(dids)
	if err != nil && err.Error() != core.ConstStr.NotFound {
		log.Error("SetDeviceBySn-err:", err)
		return &pb.BatchDeviceRe{Code: 10001}, nil
	}
	//todo 信息组合
	if len(devicePos) == 0 {
		return &pb.BatchDeviceRe{Code: 33013}, nil
	}
	var ds []*pb.DeviceReply
	for _, vd := range devicePos {
		// todo 为了适配APP端
		if vd.LedState == -1 {
			vd.LedState = 0
		}
		if vd.LedLight == -1 {
			vd.LedLight = 0
		}
		deviceRe := &pb.DeviceReply{Did: vd.Did, Types:vd.Types, DeviceName:vd.DeviceName,
			DeviceVersion:vd.DeviceVersion, DeviceMac:vd.DeviceMac,SoftwareVersion:vd.SoftwareVersion,
			LedModel:int32(vd.LedModel),
			LedColor:vd.LedColor,
			LedLight:int32(vd.LedLight),
			LedState:int32(vd.LedState),
			AudioId:int32(vd.AudioId),
			}
		for _, dp := range devicePetPos {if deviceRe.Did == dp.Did {deviceRe.Pid = dp.Pid;break}}
		for _, ud := range userDevicePos {if deviceRe.Did == ud.Did {deviceRe.Permit = ud.Permit;break}}
		ds = append(ds,deviceRe)
	}
	return &pb.BatchDeviceRe{Code: 10000, Devices: ds}, nil
}

//获取设备sn
func (this *DeviceRpc) GetDeviceSn(ctx context.Context, req *pb.DeviceRequest) (*pb.DeviceReply, error) {
	log.Info("GetDeviceSn-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.DeviceReply{Code: util.Source_err_empty}, nil
	}
	userDevicePo := &db.UserDevicePo{Uid:req.Uid, Did:req.Did}
	err := userDevicePo.GetUserDeviceDB()
	if err != nil {
		log.Info("GetDeviceSn-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.DeviceReply{Code: 33012}, nil
		}
		return &pb.DeviceReply{Code: 10001}, nil
	}
	devicePo := &db.DevicePo{Did:req.Did}
	err = devicePo.GetDeviceByDid()
	if err != nil {
		log.Info("GetDeviceSn-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.DeviceReply{Code: 33012}, nil
		}
		return &pb.DeviceReply{Code: 10001}, nil
	}
	return &pb.DeviceReply{Code: 10000,Sn:devicePo.Sn}, nil
}
