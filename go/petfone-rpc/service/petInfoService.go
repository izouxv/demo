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

/**
宠物信息
 */
type PetinfoRpc struct {
}

//添加宠物信息
func (this *PetinfoRpc) SetPetInfo(ctx context.Context, req *pb.PetInfoRequest) (*pb.PetInfoReply, error) {
	log.Info("SetPetInfo-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.PetInfoReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetUid()) {
		return &pb.PetInfoReply{Code: util.Params_err_empty}, nil
	}
	//todo 添加宠物与训练信息
	times, err :=util.Int64ToTime(req.Birthday)
	if err != nil {
		log.Info("SetPetInfo-err:", err)
		return &pb.PetInfoReply{Code: 10001}, nil
	}
	nowTime := util.GetNowTime()
	petinfoPo := &db.PetInfoPo{Avatar: req.Avatar, Nickname:req.Nickname, Gender:req.Gender, Breed:req.Breed,
		Birthday:times, Weight:req.Weight, Somatotype:req.Somatotype, Duration:core.ConstStr.PetDuration, Brightness:0, CreationTime:nowTime, UpdateTime:nowTime, DataState:1}
	pettrainPo := &db.PetTrainPo{Num:0, CreationTime:nowTime, UpdateTime:nowTime, DataState:1}
	log.Info("SetPetInfo:", petinfoPo,pettrainPo)
	if req.Avatar == "" {
		petinfoPo.Avatar = core.ConstStr.PetAvatar
	}
	dbc := core.MysqlClient.Begin()
	petTrains,err := petinfoPo.SetPetInfoDB(dbc,pettrainPo)
	if err != nil {
		dbc.Rollback()
		log.Info("SetPetInfo-err:", err)
		return &pb.PetInfoReply{Code: 10001}, nil
	}
	//todo 添加用户关联
	userPetPo := &db.UserPetPo{Uid:req.Uid, Pid:petinfoPo.Pid, CreationTime:nowTime, UpdateTime:nowTime, DataState:1}
	num := userPetPo.SetUserPetDB(dbc)
	if num != 1 {
		dbc.Rollback()
		return &pb.PetInfoReply{Code: 10001}, nil
	}
	dbc.Commit()
	var petTrainReply []*pb.PetTrainReply
	for k,v := range petTrains {
		petTrainReply = append(petTrainReply, &pb.PetTrainReply{Id:v.Id,Name:v.Name,Voice:v.Voice,DevFID:v.DevFid,Counter:v.Counter,Num:v.Num,SmallId:uint32(k)})
	}
	return &pb.PetInfoReply{Code:10000,Pid:petinfoPo.Pid, Avatar:core.ConstStr.FileServer+petinfoPo.Avatar,
		Nickname: req.Nickname, Breed: req.Breed, Gender: req.Gender, Birthday: req.Birthday,
		Weight: req.Weight, Somatotype: req.Somatotype, Duration:petinfoPo.Duration, Brightness:petinfoPo.Brightness,
		CreateTime:petinfoPo.CreationTime.Unix(), Permit:userPetPo.Permit, Trains:petTrainReply}, nil
}

//删除宠物信息
func (this *PetinfoRpc) DeletePetInfoByPid(ctx context.Context, req *pb.PetInfoRequest) (*pb.PetInfoReply, error) {
	log.Info("DeletePetInfoByPid-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.PetInfoReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetPid(), req.GetUid()) {
		return &pb.PetInfoReply{Code: util.Params_err_empty}, nil
	}
	//todo 检查权限
	userPetPo := &db.UserPetPo{Uid:req.Uid,Pid:req.Pid}
	userPetPos, err := userPetPo.GetUsersPetDB()
	if err != nil {
		log.Info("DeletePetInfoByPid-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.PetInfoReply{Code: 33012}, nil
		}
		return &pb.PetInfoReply{Code: 10001}, nil
	}
	if len(userPetPos) != 1  {
		return &pb.PetInfoReply{Code: util.User_permission}, nil
	}
	if userPetPos[0].Permit != 0 || userPetPos[0].Uid != req.Uid {
		return &pb.PetInfoReply{Code: util.User_permission}, nil
	}
	//todo 检查关系
	devicePetPo := &db.DevicePetPo{Pid:req.Pid}
	err = devicePetPo.GetDevicesPetDB()
	log.Info("devicePetPo:",devicePetPo)
	if devicePetPo.Id != 0 {
		return &pb.PetInfoReply{Code: util.User_manipulate}, nil
	}
	nowTime := util.GetNowTime()
	//todo 删除
	dbc := core.MysqlClient.Begin()
	petinfoPo := db.PetInfoPo{Pid: req.Pid, UpdateTime: nowTime}
	num := petinfoPo.DeletePetInfoDB(dbc)
	if num != 1 {
		dbc.Rollback()
		return &pb.PetInfoReply{Code: util.User_permission}, nil
	}
	userPetPo.UpdateTime = nowTime
	err = userPetPo.DeleteUserPetDB(dbc)
	if err != nil {
		dbc.Rollback()
		log.Info("DeletePetInfoByPid-err:",err)
		return &pb.PetInfoReply{Code: 10001}, nil
	}
	petTrainPo := &db.PetTrainPo{Pid:req.Pid}
	err = petTrainPo.DeletePetTrainsDB(dbc,util.GetZeroTime(nowTime.Unix()),nowTime.Unix())
	if err != nil {
		dbc.Rollback()
		log.Info("DeletePetInfoByPid-err:",err)
		return &pb.PetInfoReply{Code: 10001}, nil
	}
	dbc.Commit()
	return &pb.PetInfoReply{Code: 10000}, nil
}

//修改宠物信息
func (this *PetinfoRpc) UpdatePetInfoByPid(ctx context.Context, req *pb.PetInfoRequest) (*pb.PetInfoReply, error) {
	log.Info("UpdatePetInfoByPid:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.PetInfoReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetPid()) {
		return &pb.PetInfoReply{Code: util.Source_err_empty}, nil
	}
	//todo 检查权限
	userPetPo := &db.UserPetPo{Uid:req.Uid,Pid:req.Pid}
	err := userPetPo.GetUserPetDB()
	if err != nil {
		log.Info("UpdatePetInfoByPid-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.PetInfoReply{Code: 33012}, nil
		}
		return &pb.PetInfoReply{Code: 10001}, nil
	}
	petinfoPo := db.PetInfoPo{Pid: req.Pid, Avatar: req.Avatar, Nickname: req.Nickname, Breed: req.Breed, Gender: req.Gender,
		Weight: req.Weight, Somatotype: req.Somatotype, Duration:req.Duration, UpdateTime: util.GetNowTime()}
	if req.Birthday > 0 {
		birthday, err := util.Int64ToTime(req.Birthday)
		if err != nil {
			return &pb.PetInfoReply{Code: 10001}, nil
		}
		petinfoPo.Birthday = birthday
	}
	err = petinfoPo.UpdatePetInfoDB()
	if err != nil {
		log.Error("UpdatePetInfoByPid-UpdatePetInfoDB:",err)
		return &pb.PetInfoReply{Code: 10001}, nil
	}
	if len(req.Avatar) != 0 {
		return &pb.PetInfoReply{Code: 10000,Avatar: core.ConstStr.FileServer+petinfoPo.Avatar}, nil
	}
	return &pb.PetInfoReply{Code: 10000}, nil
}

//获取宠物信息
func (this *PetinfoRpc) GetPetInfoByPid(ctx context.Context, req *pb.PetInfoRequest) (*pb.PetInfoReply, error) {
	log.Info("GetPetInfoByPid-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.PetInfoReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetPid(),req.GetUid()) {
		return &pb.PetInfoReply{Code: util.Source_err_empty}, nil
	}
	//todo 检查权限
	userPetPo := &db.UserPetPo{Uid:req.Uid,Pid:req.Pid}
	err := userPetPo.GetUserPetDB()
	if err != nil {
		log.Info("GetPetInfoByPid-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.PetInfoReply{Code: 33012}, nil
		}
		return &pb.PetInfoReply{Code: 10001}, nil
	}
	petInfoPo := &db.PetInfoPo{Pid: req.Pid, DataState: 1}
	err = petInfoPo.GetPetInfoDB()
	if err != nil {
		log.Info("err", err)
		return &pb.PetInfoReply{Code: 10001}, nil
	}
	//todo 查询宠物训练信息
	nowTime := util.GetNowTimeSecond()
	petTrainPo := &db.PetTrainPo{Pid:req.Pid}
	petTrainPos, err := petTrainPo.GetPetTrainsDB(util.GetZeroTime(nowTime),nowTime)
	if err != nil {
		log.Info("err", err)
		return &pb.PetInfoReply{Code: 10001}, nil
	}
	var petTrainReply []*pb.PetTrainReply
	for k,v := range petTrainPos {
		petTrainReply = append(petTrainReply, &pb.PetTrainReply{Id:v.Id,Name:v.Name,Voice:v.Voice,DevFID:v.DevFid,Counter:v.Counter,Num:v.Num,SmallId:uint32(k)})
	}
	return &pb.PetInfoReply{Code: 10000, Pid: petInfoPo.Pid, Avatar: core.ConstStr.FileServer +petInfoPo.Avatar,
		Nickname: petInfoPo.Nickname, Breed: petInfoPo.Breed, Gender: petInfoPo.Gender, Birthday: petInfoPo.Birthday.Unix(),
		Weight: petInfoPo.Weight, Somatotype: petInfoPo.Somatotype, Duration:petInfoPo.Duration, Brightness:petInfoPo.Brightness,
		CreateTime:petInfoPo.CreationTime.Unix(), Permit:userPetPo.Permit, Trains:petTrainReply}, nil
}

//批量查询宠物信息
func (this *PetinfoRpc) GetPetInfoByUid(ctx context.Context, req *pb.PetInfoRequest) (*pb.PetInfoMapReply, error) {
	log.Info("req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.PetInfoMapReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetUid()) {
		return &pb.PetInfoMapReply{Code: util.User_params_err}, nil
	}
	//todo 查询用户宠物关系
	userPetPo := &db.UserPetPo{Uid:req.Uid}
	userPetPos, err := userPetPo.GetUserPetsDB()
	if err != nil {
		log.Info("GetPetInfoByUid-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.PetInfoMapReply{Code: 33012}, nil
		}
		return &pb.PetInfoMapReply{Code: 10001}, nil
	}
	var pids []int32
	for _,v := range userPetPos {
		pids = append(pids,v.Pid)
	}
	//todo 查询批量宠物信息
	petInfoPo := &db.PetInfoPo{}
	petInfoPos, err := petInfoPo.GetPetInfosDB(pids)
	if err != nil {
		log.Info("err", err)
		return &pb.PetInfoMapReply{Code: 10001}, nil
	}
	//todo 查询设备与宠物关系
	devicePetPo := &db.DevicePetPo{}
	devicePetPos, err := devicePetPo.GetDevicesPetsPidsDB(pids)
	if err != nil {
		log.Info("err", err)
		return &pb.PetInfoMapReply{Code: 10001}, nil
	}
	//todo 查询宠物训练信息
	nowTime := util.GetNowTimeSecond()
	petTrainPo := &db.PetTrainPo{Pid:req.Pid}
	petTrainPos, err := petTrainPo.GetPetsTrainsDB(pids,util.GetZeroTime(nowTime),nowTime)
	if err != nil {
		log.Info("err", err)
		return &pb.PetInfoMapReply{Code: 10001}, nil
	}
	//todo 聚合信息
	petRes := make(map[int32]*pb.PetInfoReply)
	for k,v := range petInfoPos {
		petinfoRe := &pb.PetInfoReply{Pid: v.Pid, Avatar: core.ConstStr.FileServer+v.Avatar, Nickname: v.Nickname, Gender: v.Gender,
		Breed: v.Breed, Birthday: v.Birthday.Unix(), Weight: v.Weight, Somatotype: v.Somatotype,
		Duration:v.Duration, Brightness:v.Brightness, CreateTime:v.CreationTime.Unix()}
		//todo 遍历设备与宠物关联关系
		for _, vp := range devicePetPos {
			if vp.Pid == v.Pid {
				petinfoRe.Did = vp.Did
				break
			}
		}
		//todo 遍历用户与宠物关联关系
		for _, up := range userPetPos {
			if up.Pid == v.Pid {
				petinfoRe.Permit = up.Permit
				break
			}
		}
		//todo 遍历训练信息
		var smallId uint32
		var petTrainReply []*pb.PetTrainReply
		for _,vt := range petTrainPos {
			if v.Pid == vt.Pid {
				petTrainReply = append(petTrainReply,
					&pb.PetTrainReply{Id: vt.Id, Name: vt.Name, Voice: vt.Voice, DevFID:vt.DevFid,Counter: vt.Counter, Num: vt.Num, SmallId: smallId})
				smallId++
			}
		}
		petinfoRe.Trains = petTrainReply
		petRes[int32(k)] = petinfoRe
	}
	return &pb.PetInfoMapReply{Code: 10000, Petinfos:petRes}, nil
}

/**
设备与宠物关系
 */
//关联设备与宠物
func (this *PetinfoRpc) SetDevicePet(ctx context.Context, req *pb.PetInfoRequest) (*pb.PetInfoReply, error) {
	log.Info("SetDevicePet-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.PetInfoReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetUid(), req.GetDid(), req.GetPid()) {
		return &pb.PetInfoReply{Code: util.Params_err_empty}, nil
	}
	//todo 检查权限
	userPetPo := &db.UserPetPo{Uid:req.Uid,Pid:req.Pid}
	err := userPetPo.GetUserPetDB()
	log.Info("SetDevicePet-userPetPo:",userPetPo)
	if err != nil {
		log.Info("SetDevicePet-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.PetInfoReply{Code: 33012}, nil
		}
		return &pb.PetInfoReply{Code: 10001}, nil
	}
	if userPetPo == nil {
		return &pb.PetInfoReply{Code: util.User_permission}, nil
	}
	userDevicePo := &db.UserDevicePo{Uid:req.Uid, Did:req.Did}
	err = userDevicePo.GetUserDeviceDB()
	log.Info("SetDevicePet-userDevicePo:",userDevicePo)
	if err != nil {
		log.Info("SetDevicePet-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.PetInfoReply{Code: 33012}, nil
		}
		return &pb.PetInfoReply{Code: 10001}, nil
	}
	//todo 查询关联关系
	nowTime := util.GetNowTime()
	devicePetPo := &db.DevicePetPo{Pid:req.Pid,Did:req.Did,CreationTime:nowTime, UpdateTime: nowTime, DataState:1}
	err = devicePetPo.GetDevicePetDB()
	log.Info("SetDevicePet-devicePetPo:",devicePetPo)
	if err != nil {
		if err.Error() == core.ConstStr.NotFound {
			//todo 可关联
			num := devicePetPo.SetDevicePetDB()
			if num != 1 {
				return &pb.PetInfoReply{Code: 10001}, nil
			}
			return &pb.PetInfoReply{Code: 10000}, nil
		}
		return &pb.PetInfoReply{Code: 10001}, nil
	}
	return &pb.PetInfoReply{Code: 33014}, nil
}

//取消关联设备与宠物
func (this *PetinfoRpc) DeleteDevicePet(ctx context.Context, req *pb.PetInfoRequest) (*pb.PetInfoReply, error) {
	log.Info("DeleteDevicePet-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.PetInfoReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetUid(), req.GetDid(), req.GetPid()) {
		return &pb.PetInfoReply{Code: util.Params_err_empty}, nil
	}
	//todo 检查关系
	userPetPo := &db.UserPetPo{Uid:req.Uid,Pid:req.Pid}
	userPetPos, err := userPetPo.GetUsersPetDB()
	if err != nil {
		log.Info("DeleteDevicePet-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.PetInfoReply{Code: 33012}, nil
		}
		return &pb.PetInfoReply{Code: 10001}, nil
	}
	userDevicePo := &db.UserDevicePo{Uid:req.Uid, Did:req.Did}
	userDevicePos, err := userDevicePo.GetUsersDeviceDB()
	if err != nil {
		log.Info("DeleteDevicePet-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.PetInfoReply{Code: 33012}, nil
		}
		return &pb.PetInfoReply{Code: 10001}, nil
	}
	//todo 检查共享或权限
	if len(userPetPos) != 1 || len(userDevicePos) != 1 {
		return &pb.PetInfoReply{Code: util.User_permission}, nil
	}
	if userPetPos[0].Permit != 0 || userPetPos[0].Uid != req.Uid || userDevicePos[0].Uid != req.Uid || userDevicePos[0].Permit != 0 {
		return &pb.PetInfoReply{Code: util.User_permission}, nil
	}
	//todo 查询关联关系
	devicePetPo := &db.DevicePetPo{Did:req.Did, Pid:req.Pid}
	err = devicePetPo.GetDevicePetDB()
	log.Info("devicePetPo:",devicePetPo)
	if err != nil {
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.PetInfoReply{Code: 33012}, nil
		}
		log.Info("DeleteDevicePet-err:", err)
		return &pb.PetInfoReply{Code: 10001}, nil
	}
	devicePetPo.UpdateTime = util.GetNowTime()
	num := devicePetPo.DeleteDevicePetDB()
	if num != 1 {
		return &pb.PetInfoReply{Code: 10001}, nil
	}
	return &pb.PetInfoReply{Code: 10000}, nil
}

//查询关联设备与宠物
func (this *PetinfoRpc) GetDevicePet(ctx context.Context, req *pb.PetInfoRequest) (*pb.PetInfoReply, error) {
	log.Info("GetDevicePet-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.PetInfoReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetUid(), req.GetDid(), req.GetPid()) {
		return &pb.PetInfoReply{Code: util.Params_err_empty}, nil
	}
	//todo 查询关联关系
	devicePetPo := &db.DevicePetPo{Did:req.Did, Pid:req.Pid}
	err := devicePetPo.GetDevicePetDB()
	log.Info("devicePetPo:",devicePetPo)
	if err != nil {
		return &pb.PetInfoReply{Code: 33012}, nil
	}
	return &pb.PetInfoReply{Code: 10000}, nil
}

/**
宠物训练信息
 */
//查询宠物训练历史
func (this *PetinfoRpc) GetPetTrainByPid(ctx context.Context, req *pb.PetTrainRequest) (*pb.PetSliceTrainsReply, error) {
	log.Info("GetPetTrainByPid-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.PetSliceTrainsReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetPid(),req.Uid) {
		return &pb.PetSliceTrainsReply{Code: util.User_params_err}, nil
	}
	if util.VerifyParamsUInt64(req.StartTime,req.EndTime) {
		return &pb.PetSliceTrainsReply{Code: util.User_params_err}, nil
	}
	//todo 检查权限
	userPetPo := &db.UserPetPo{Uid:req.Uid,Pid:req.Pid}
	err := userPetPo.GetUserPetDB()
	if err != nil {
		log.Info("UpdatePetInfoByPid-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.PetSliceTrainsReply{Code: 33012}, nil
		}
		return &pb.PetSliceTrainsReply{Code: 10001}, nil
	}
	//查询时间内训练详情
	petTrainPo := &db.PetTrainPo{Pid:req.Pid}
	petTrainPos, err := petTrainPo.GetPetTrainsDB(req.StartTime,req.EndTime)
	if err != nil {
		return &pb.PetSliceTrainsReply{Code: 10001}, nil
	}
	if len(petTrainPos) == 0 {
		return &pb.PetSliceTrainsReply{Code: 33013}, nil
	}
	var sliceTrains []*pb.PetTrains
	endTime := req.EndTime
	var smallId uint32
	for req.StartTime <= endTime {
		petTrains := &pb.PetTrains{}
		var petTrain []*pb.PetTrainReply
		zeroTime := util.GetZeroTime(endTime)
		for _,v := range petTrainPos {
			if zeroTime <= v.CreationTime.Unix() && (zeroTime+86400) > v.CreationTime.Unix() {
				petTrain = append(petTrain,&pb.PetTrainReply{Id:v.Id, Name:v.Name, Voice:core.ConstStr.FileServer+v.Voice,
				DevFID:v.DevFid, Num:v.Num, Counter:v.Counter, Times:v.CreationTime.Unix(), SmallId:smallId})
				if smallId == 2 {
					petTrains.Trains = petTrain
					sliceTrains = append(sliceTrains,petTrains)
					smallId = 0
				} else {
					smallId++
				}
			}
		}
		endTime -= 86400
	}
	return &pb.PetSliceTrainsReply{Code: 10000, SliceTrains:sliceTrains}, nil
}

//修改宠物训练信息
func (this *PetinfoRpc) UpdatePetTrainByPid(ctx context.Context, req *pb.PetTrainRequest) (*pb.PetTrainReply, error) {
	log.Info("UpdatePetTrainByPid-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.PetTrainReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.Id, req.GetPid(),req.Uid) {
		return &pb.PetTrainReply{Code: 33001}, nil
	}
	//todo 检查权限
	userPetPo := &db.UserPetPo{Uid:req.Uid,Pid:req.Pid}
	err := userPetPo.GetUserPetDB()
	if err != nil {
		log.Info("UpdatePetInfoByPid-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.PetTrainReply{Code: 33012}, nil
		}
		return &pb.PetTrainReply{Code: 10001}, nil
	}
	petTrainPo := &db.PetTrainPo{Id:req.Id, Pid:req.Pid, Name:req.Name, Voice:req.Voice, DevFid:req.DevFID,Num:req.Num, UpdateTime:util.GetNowTime()}
	err = petTrainPo.UpdatePetTrainDB()
	if err != nil {
		log.Info("UpdatePetTrainByPid-UpdatePetTrainDB:",err)
		return &pb.PetTrainReply{Code: 33012}, nil
	}
	// 给device_train表里写入数据
	go UpdateDeviceTrain(req)
	return &pb.PetTrainReply{Code: 10000}, nil
}

//计数宠物训练次数
func (this *PetinfoRpc) CounterPetTrainByPid(ctx context.Context, req *pb.PetTrainRequest) (*pb.PetTrainReply, error) {
	log.Info("CounterPetTrainByPid-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.PetTrainReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetPid()) {
		return &pb.PetTrainReply{Code: util.Source_err_empty}, nil
	}
	//todo 检查权限
	userPetPo := &db.UserPetPo{Uid:req.Uid,Pid:req.Pid}
	err := userPetPo.GetUserPetDB()
	if err != nil {
		log.Info("UpdatePetInfoByPid-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.PetTrainReply{Code: 33012}, nil
		}
		return &pb.PetTrainReply{Code: 10001}, nil
	}
	petTrainPo := &db.PetTrainPo{Id:req.Id, Pid:req.Pid,Counter:req.Counter,UpdateTime:util.GetNowTime()}
	num := petTrainPo.CounterPetTrainDB()
	if num != 1 {
		log.Info("UpdatePetInfoByPid-num:", num)
		return &pb.PetTrainReply{Code: 33012}, nil
	}
	return &pb.PetTrainReply{Code: 10000}, nil
}

// todo modify xiaorx 2018/7/4
//修改宠端设备训练信息
func (this *PetinfoRpc) UpdateDeviceTrainByDid(ctx context.Context, req *pb.DeviceTrainRequest) (*pb.DeviceTrainReply, error) {
	log.Infof("UpdateDeviceTrainByDid...req(%#v)", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.DeviceTrainReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.Id, req.Did,req.Uid) {
		return &pb.DeviceTrainReply{Code: 33001}, nil
	}
	//todo 检查权限
	userPetPo := &db.UserDevicePo{Uid:req.Uid,Did:req.Did}
	if req.Did != 0 {
		err := userPetPo.GetUserDeviceDB()
		if err != nil {
			log.Info("UpdateDeviceTrainByDid-err:", err)
			if err.Error() ==  core.ConstStr.NotFound {
				return &pb.DeviceTrainReply{Code: 33012}, nil
			}
			return &pb.DeviceTrainReply{Code: 10001}, nil
		}
	}
	deviceTrainPo := &db.DeviceTrainPo{Id:req.Id, Did:req.Did, Name:req.Name, Voice:req.Voice, DevFid:req.DevFID,Num:req.Num, UpdateTime:util.GetNowTime()}
	err := deviceTrainPo.UpdateDeviceTrainDB()
	if err != nil {
		log.Infof("UpdateDeviceTrainByDid-err...err(%#v)",err)
		return &pb.DeviceTrainReply{Code: 33012}, nil
	}
	return &pb.DeviceTrainReply{Code: 10000}, nil
}
//获取单个宠物信息
func (this *PetinfoRpc) GetPetInfoBydid(ctx context.Context, req *pb.PetInfoRequest) (*pb.PetInfoReply, error) {
	log.Info("GetPetInfoByPid-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.PetInfoReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetPid(),req.GetUid()) {
		return &pb.PetInfoReply{Code: util.Source_err_empty}, nil
	}
	//todo 检查权限
	userPetPo := &db.UserPetPo{Uid:req.Uid,Pid:req.Pid}
	err := userPetPo.GetUserPetDB()
	if err != nil {
		log.Info("GetPetInfoByPid-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.PetInfoReply{Code: 33012}, nil
		}
		return &pb.PetInfoReply{Code: 10001}, nil
	}
	petInfoPo := &db.PetInfoPo{Pid: req.Pid, DataState: 1}
	err = petInfoPo.GetPetInfoDB()
	if err != nil {
		log.Info("err", err)
		return &pb.PetInfoReply{Code: 10001}, nil
	}
	//todo 查询宠物训练信息
	nowTime := util.GetNowTimeSecond()
	deviceTrainPo := &db.DeviceTrainPo{Did:req.Did}
	petTrainPos, err := deviceTrainPo.GetDeviceTrainsDB(util.GetZeroTime(nowTime),nowTime)
	if err != nil {
		log.Info("err", err)
		return &pb.PetInfoReply{Code: 10001}, nil
	}
	var petTrainReply []*pb.PetTrainReply
	for k,v := range petTrainPos {
		petTrainReply = append(petTrainReply, &pb.PetTrainReply{Id:v.Id,Name:v.Name,Voice:v.Voice,DevFID:v.DevFid,Counter:v.Counter,Num:v.Num,SmallId:uint32(k)})
	}
	return &pb.PetInfoReply{Code: 10000, Pid: petInfoPo.Pid, Avatar: core.ConstStr.FileServer +petInfoPo.Avatar,
		Nickname: petInfoPo.Nickname, Breed: petInfoPo.Breed, Gender: petInfoPo.Gender, Birthday: petInfoPo.Birthday.Unix(),
		Weight: petInfoPo.Weight, Somatotype: petInfoPo.Somatotype, Duration:petInfoPo.Duration, Brightness:petInfoPo.Brightness,
		CreateTime:petInfoPo.CreationTime.Unix(), Permit:userPetPo.Permit, Trains:petTrainReply}, nil
}
//批量获取宠物信息
func (this *PetinfoRpc) GetPetInfosBydid(ctx context.Context, req *pb.PetInfoRequest) (*pb.PetInfoMapReply, error) {
	log.Info("GetPetInfosBydid...", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.PetInfoMapReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetUid()) {
		return &pb.PetInfoMapReply{Code: util.User_params_err}, nil
	}
	//todo 查询用户宠物
	userPetPo := &db.UserPetPo{Uid:req.Uid}
	userPetPos, err := userPetPo.GetUserPetsDB()
	if err != nil {
		log.Info("GetPetInfosBydid-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.PetInfoMapReply{Code: 33012}, nil
		}
		return &pb.PetInfoMapReply{Code: 10001}, nil
	}
	var pids []int32
	for _,v := range userPetPos {
		pids = append(pids,v.Pid)
	}
	//todo 查询批量宠物信息
	petInfoPo := &db.PetInfoPo{}
	petInfoPos, err := petInfoPo.GetPetInfosDB(pids)
	if err != nil {
		log.Info("err", err)
		return &pb.PetInfoMapReply{Code: 10001}, nil
	}
	//todo 查询设备与宠物关系
	devicePetPo := &db.DevicePetPo{}
	devicePetPos, err := devicePetPo.GetDevicesPetsPidsDB(pids)
	if err != nil {
		log.Info("err", err)
		return &pb.PetInfoMapReply{Code: 10001}, nil
	}


	//todo 查询宠物训练信息
	nowTime := util.GetNowTimeSecond()
	deviceTrainPo := &db.DeviceTrainPo{Did:req.Did}
	petTrainPos, err := deviceTrainPo.GetDevicesTrainsDB(req.Did,util.GetZeroTime(nowTime),nowTime)
	if err != nil {
		log.Info("err", err)
		return &pb.PetInfoMapReply{Code: 10001}, nil
	}
	//todo 聚合信息
	petRes := make(map[int32]*pb.PetInfoReply)
	for k,v := range petInfoPos {
		petinfoRe := &pb.PetInfoReply{Pid: v.Pid, Avatar: core.ConstStr.FileServer+v.Avatar, Nickname: v.Nickname, Gender: v.Gender,
			Breed: v.Breed, Birthday: v.Birthday.Unix(), Weight: v.Weight, Somatotype: v.Somatotype,
			Duration:v.Duration, Brightness:v.Brightness, CreateTime:v.CreationTime.Unix()}
		//todo 遍历设备与宠物关联关系
		for _, vp := range devicePetPos {
			if vp.Pid == v.Pid {
				petinfoRe.Did = vp.Did
				break
			}
		}
		//todo 遍历用户与宠物关联关系
		for _, up := range userPetPos {
			if up.Pid == v.Pid {
				petinfoRe.Permit = up.Permit
				break
			}
		}
		//todo 遍历训练信息
		var smallId uint32
		var petTrainReply []*pb.PetTrainReply
		for _,vt := range petTrainPos {
			if v.Pid == vt.Pid {
				petTrainReply = append(petTrainReply,
					&pb.PetTrainReply{Id: vt.Id, Name: vt.Name, Voice: vt.Voice, DevFID:vt.DevFid,Counter: vt.Counter, Num: vt.Num, SmallId: smallId})
				smallId++
			}
		}
		petinfoRe.Trains = petTrainReply
		petRes[int32(k)] = petinfoRe
	}
	return &pb.PetInfoMapReply{Code: 10000, Petinfos:petRes}, nil
}
/*
//批量获取宠物信息
func (this *PetinfoRpc) GetPetInfosBydid(ctx context.Context, req *pb.PetInfoRequest) (*pb.PetInfoMapReply, error) {
	log.Info("GetPetInfosBydid...", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.PetInfoMapReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetUid()) {
		return &pb.PetInfoMapReply{Code: util.User_params_err}, nil
	}
	//todo 查询用户宠物
	userPetPo := &db.UserPetPo{Uid:req.Uid}
	userPetPos, err := userPetPo.GetUserPetsDB()
	if err != nil {
		log.Info("GetPetInfosBydid-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.PetInfoMapReply{Code: 33012}, nil
		}
		return &pb.PetInfoMapReply{Code: 10001}, nil
	}
	var pids []int32
	for _,v := range userPetPos {
		pids = append(pids,v.Pid)
	}
	//todo 查询批量宠物信息
	petInfoPo := &db.PetInfoPo{}
	petInfoPos, err := petInfoPo.GetPetInfosDB(pids)
	if err != nil {
		log.Info("err", err)
		return &pb.PetInfoMapReply{Code: 10001}, nil
	}
	//todo 查询设备与宠物关系
	devicePetPo := &db.DevicePetPo{}
	devicePetPos, err := devicePetPo.GetDevicesPetsPidsDB(pids)
	if err != nil {
		log.Info("err", err)
		return &pb.PetInfoMapReply{Code: 10001}, nil
	}
	//todo 查询宠物训练信息
	nowTime := util.GetNowTimeSecond()
	deviceTrainPo := &db.DeviceTrainPo{Did:req.Did}
	petTrainPos, err := deviceTrainPo.GetDevicesTrainsDB(req.Did,util.GetZeroTime(nowTime),nowTime)
	if err != nil {
		log.Info("err", err)
		return &pb.PetInfoMapReply{Code: 10001}, nil
	}
	//todo 聚合信息
	petRes := make(map[int32]*pb.PetInfoReply)
	for k,v := range petInfoPos {
		petinfoRe := &pb.PetInfoReply{Pid: v.Pid, Avatar: core.ConstStr.FileServer+v.Avatar, Nickname: v.Nickname, Gender: v.Gender,
			Breed: v.Breed, Birthday: v.Birthday.Unix(), Weight: v.Weight, Somatotype: v.Somatotype,
			Duration:v.Duration, Brightness:v.Brightness, CreateTime:v.CreationTime.Unix()}
		//todo 遍历设备与宠物关联关系
		for _, vp := range devicePetPos {
			if vp.Pid == v.Pid {
				petinfoRe.Did = vp.Did
				break
			}
		}
		//todo 遍历用户与宠物关联关系
		for _, up := range userPetPos {
			if up.Pid == v.Pid {
				petinfoRe.Permit = up.Permit
				break
			}
		}
		//todo 遍历训练信息
		var smallId uint32
		var petTrainReply []*pb.PetTrainReply
		for _,vt := range petTrainPos {
			if v.Pid == vt.Pid {
				petTrainReply = append(petTrainReply,
					&pb.PetTrainReply{Id: vt.Id, Name: vt.Name, Voice: vt.Voice, DevFID:vt.DevFid,Counter: vt.Counter, Num: vt.Num, SmallId: smallId})
				smallId++
			}
		}
		petinfoRe.Trains = petTrainReply
		petRes[int32(k)] = petinfoRe
	}
	return &pb.PetInfoMapReply{Code: 10000, Petinfos:petRes}, nil
}*/

//宠物录音时，在device_train 表里插入数据
func UpdateDeviceTrain(req *pb.PetTrainRequest)  {
	log.Infof("UpdateDeviceTrain...")
	// 获取用户绑定的宠端设备
	device := &db.UserDevicePo{Uid:req.Uid,DataState:1,Types:pb.DeviceTypes_Pet}
	if err := device.GetUserDevicesForPet();err != nil {
		log.Error("错误：用户没用绑定宠端设备")
		return
	}
	// 1.device_train 表里是否存在
	deviceTrain := &db.DeviceTrainPo{Did:device.Did,DataState:1}
	_, err := deviceTrain.GetDeviceTrains()
	if err == gorm.ErrRecordNotFound {
		//不存在
			//1.1查找pet_train表里是否存在
			petTrain := &db.PetTrainPo{Pid:req.Pid}
			petTrains,err := petTrain.IsExistPetTrains()
			log.Debugf("petTrains(%#v)",petTrains)
			if err == gorm.ErrRecordNotFound {
				//获取device_train的默认值
				defaultTrain := &db.DeviceTrainPo{Did:0}
				defaultData,err := defaultTrain.GetDeviceTrains()
				if err != nil {
					log.Errorf("获取宠端默认训练语音报错defaulTrain(%#v)",defaultTrain)
					return
				}
				// 遍历默认数据
				timeNow := time.Now()
				err = db.Transaction(core.MysqlClient, func(tx *gorm.DB) error {
					for _,value := range defaultData {
						deviceTrain := db.DeviceTrainPo{}
						deviceTrain.Name = req.Name
						deviceTrain.Did = device.Did
						deviceTrain.Num = 10
						deviceTrain.Counter = 0
						deviceTrain.DataState = 1
						deviceTrain.DevFid = 0
						deviceTrain.Voice = req.Voice
						deviceTrain.CreationTime = timeNow
						if value.Name == req.Name {
							deviceTrain.Name = value.Name
							deviceTrain.Voice = value.Voice
						}else {
							deviceTrain.Name = req.Name
							deviceTrain.Voice = req.Voice
						}
						if deviceTrain.Create(core.MysqlClient) != 1 {
							log.Error("插入宠端设备录音信息失败！")
							return gorm.ErrInvalidTransaction
						}
					}
					return nil

				})
				if err != nil {
					log.Errorf("插入宠端设备录音失败！err(%#v)",err)
				}
			}
	} else if err == nil {
		//存在,直接更新
		if err = db.Transaction(core.MysqlClient, func(tx *gorm.DB) error {
			deviceTrain := db.DeviceTrainPo{Did:device.Did,Name:req.Name}
			if err = deviceTrain.UpdateDeviceTrainForName(); err != nil {
				log.Error("更新宠端设备录音信息失败！")
				return gorm.ErrInvalidTransaction
			}
			return nil
		});err != nil {
			log.Errorf("更新宠端录音信息失败！")
			return
		}
		return
	}
}
