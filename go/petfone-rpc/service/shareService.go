package service

import (
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"petfone-rpc/db"
	"petfone-rpc/pb"
	"petfone-rpc/util"
	"net/url"
	"petfone-rpc/core"
)

/**
宠物信息
 */
type ShareRpc struct {
}

//添加共享信息
func (this *ShareRpc) SetShare(ctx context.Context, req *pb.ShareRequest) (*pb.ShareReply, error) {
	log.Info("SetShare-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.ShareReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetOwnerUid(),req.MemberUid) {
		return &pb.ShareReply{Code: util.Params_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.Pids...) || len(req.Pids) == 0 {
		return &pb.ShareReply{Code: util.Params_err_empty}, nil
	}
	dbc := core.MysqlClient.Begin()
	nowTime := util.GetNowTime()
	//todo 分享记录
	shareUserPo := &db.ShareUserPo{Uid:req.OwnerUid,Fuid:req.MemberUid, CreationTime:nowTime,UpdateTime:nowTime,DataState:1}
	err := shareUserPo.GetShareDB()
	if err != nil && err.Error() != core.ConstStr.NotFound {
		dbc.Rollback()
		log.Info("SetShare-err:",err)
		return &pb.ShareReply{Code: 10001}, nil
	}
	userDevicePo := &db.UserDevicePo{}
	if shareUserPo.Id == 0 {
		err = shareUserPo.SetShareDB(dbc)
		if err != nil {
			dbc.Rollback()
			return &pb.ShareReply{Code: 10001}, nil
		}
		//todo 检查人端
		userDevicePo.Uid = req.OwnerUid
		userDevicePos, err := userDevicePo.GetUserDevicesDB()
		if err != nil {
			dbc.Rollback()
			return &pb.ShareReply{Code: 10001}, nil
		}
		for _, v := range userDevicePos {
			if 0 == v.Permit && 1 == v.Types {
				userDevicePo.Did = v.Did
			}
		}
		//todo 没有绑定人端不允许共享
		if 0 == userDevicePo.Did {
			dbc.Rollback()
			return &pb.ShareReply{Code: 33012}, nil
		}
		//todo 分享人端
		userDevicePo.Uid=req.MemberUid; userDevicePo.Types=1; userDevicePo.Permit=1
		userDevicePo.CreationTime=nowTime; userDevicePo.UpdateTime=nowTime; userDevicePo.DataState=1
		num := userDevicePo.SetUserDeviceDB(dbc)
		if num != 1 {
			dbc.Rollback()
			return &pb.ShareReply{Code: 10001}, nil
		}
	} else {
		shareUserPo.UpdateTime = util.GetNowTime()
		err = shareUserPo.UpdateShareDB(dbc)
		if err != nil {
			dbc.Rollback()
			return &pb.ShareReply{Code: 10001}, nil
		}
	}
	log.Info("SetShare-shareUserPos:",shareUserPo)
	//todo 查询对方是否已被共享
	userPetPo := &db.UserPetPo{Uid:req.MemberUid}
	userPetPos, err := userPetPo.GetUserPetsDB()
	if err != nil && err.Error() != core.ConstStr.NotFound {
		dbc.Rollback()
		return &pb.ShareReply{Code: 10001}, nil
	}
	log.Info("F-userPetPos:",len(userPetPos))
	for _, v := range req.Pids {
		flag := func () bool {
			for _, vu := range userPetPos {
				if v == vu.Pid && vu.Permit != 0 {
					return false
				}
			}
			return true
		}
		if flag() {
			continue
		}
		dbc.Rollback()
		return &pb.ShareReply{Code: util.User_permission}, nil
	}
	//todo 查询owner的宠物与其对应的设备
	userPetPo = &db.UserPetPo{Uid:req.OwnerUid}
	userPetPos, err = userPetPo.GetUserPetsDB()
	if err != nil {
		dbc.Rollback()
		if err.Error() == core.ConstStr.NotFound {
			return &pb.ShareReply{Code: util.User_manipulate}, nil
		}
		return &pb.ShareReply{Code: 10001}, nil
	}
	log.Info("userPetPos:",len(userPetPos))
	for _, v := range req.Pids {
		flag := func () bool {
			for _, vu := range userPetPos {
				if v == vu.Pid && vu.Permit == 0 {
				return true
				}
			}
			return false
		}
		if flag() {
			continue
		}
		dbc.Rollback()
		return &pb.ShareReply{Code: util.User_permission}, nil
	}
	devicePetPo := &db.DevicePetPo{}
	devicePetPos, err := devicePetPo.GetDevicesPetsPidsDB(req.Pids)
	if err != nil {
		dbc.Rollback()
		if err.Error() == core.ConstStr.NotFound {
			return &pb.ShareReply{Code: util.User_manipulate}, nil
		}
		return &pb.ShareReply{Code: 10001}, nil
	}
	if len(devicePetPos) != len(req.Pids) {
		dbc.Rollback()
		return &pb.ShareReply{Code: util.User_manipulate}, nil
	}
	log.Info("devicePetPos:",len(devicePetPos))
	//todo 添加设备与宠物分享
	var userDevicePos1 []*db.UserDevicePo
	for _, v := range devicePetPos {
		userDevicePos1 = append(userDevicePos1, &db.UserDevicePo{Uid:req.MemberUid, Did:v.Did,Permit:1,Types:0,CreationTime:nowTime,UpdateTime:nowTime,DataState:1})
	}
	err = userDevicePo.SetUserDevicesDB(dbc, userDevicePos1)
	if err != nil {
		dbc.Rollback()
		return &pb.ShareReply{Code: 10001}, nil
	}
	var userPetPos1 []*db.UserPetPo
	for _, v := range req.Pids {
		userPetPos1 = append(userPetPos1, &db.UserPetPo{Uid:req.MemberUid, Pid:v,Permit:1,CreationTime:nowTime,UpdateTime:nowTime,DataState:1})
	}
	err = userPetPo.SetUserPetsDB(dbc, userPetPos1)
	if err != nil {
		dbc.Rollback()
		return &pb.ShareReply{Code: 10001}, nil
	}
	dbc.Commit()
	return &pb.ShareReply{Code: 10000}, nil
}

//删除共享信息
func (this *ShareRpc) DeleteShare(ctx context.Context, req *pb.ShareRequest) (*pb.ShareReply, error) {
	log.Info("DeleteShare-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.ShareReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetOwnerUid(), req.GetMemberUid()) {
		return &pb.ShareReply{Code: util.Params_err_empty}, nil
	}
	//todo 查询分享记录
	shareUserPo := &db.ShareUserPo{Uid:req.OwnerUid,Fuid:req.MemberUid}
	err := shareUserPo.GetShareDB()
	if err != nil {
		log.Info("DeleteShare-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.ShareReply{Code: 33014}, nil
		}
		return &pb.ShareReply{Code: 10001}, nil
	}
	pidLen := len(req.Pids)
	//todo 查询用户的宠物关系
	nowTime := util.GetNowTime()
	userPetPo := &db.UserPetPo{Uid:req.OwnerUid,UpdateTime : nowTime}
	OwnerPetPos, err := userPetPo.GetUserPetsDB()
	if err != nil {
		log.Info("GetPetInfoByPid-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.ShareReply{Code: 33014}, nil
		}
		return &pb.ShareReply{Code: 10001}, nil
	}

	userPetPo.Uid = req.MemberUid
	MemberPetPos, err := userPetPo.GetUserPetsDB()
	if err != nil {
		log.Info("GetPetInfoByPid-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.ShareReply{Code: 33014}, nil
		}
		return &pb.ShareReply{Code: 10001}, nil
	}

	if 0 != len(OwnerPetPos) || 0 == len(MemberPetPos) {
		return &pb.ShareReply{Code: 33014}, nil
	}
	if pidLen != len(MemberPetPos) {
		return &pb.ShareReply{Code: 33014}, nil
	}
	////todo 筛选删除关系
	//for _, pid := range req.Pids {
	//	for _, ownerPetPo := range OwnerPetPos {
	//		if pid == ownerPetPo.Pid && 0 == ownerPetPo.Permit {
	//			break
	//		}
	//		return &pb.ShareReply{Code: 33014}, nil
	//	}
	//	for _, memberPetPo := range MemberPetPos {
	//		if pid == memberPetPo.Pid && 1 == memberPetPo.Permit {
	//			break
	//		}
	//		return &pb.ShareReply{Code: 33014}, nil
	//	}
	//}
	//todo 查询宠物与设备关系
	devicePetPo := &db.DevicePetPo{}
	devicePetPos, err := devicePetPo.GetDevicesPetsPidsDB(req.Pids)
	if err != nil {
		return &pb.ShareReply{Code: 10001}, nil
	}
	if pidLen != len(devicePetPos) {
		return &pb.ShareReply{Code: 33014}, nil
	}
	var dids []int32
	for _, v := range devicePetPos {
		dids = append(dids,v.Did)
	}
	//todo 删除用户对宠物与设备的关系
	dbc := core.MysqlClient.Begin()
	if userPetPo.DeleteUserPetsDB(dbc,req.Pids) != int64(pidLen) {
		dbc.Rollback()
		return &pb.ShareReply{Code: 33014}, nil
	}
	userDevicePo := &db.UserDevicePo{Uid:req.MemberUid, UpdateTime:nowTime}
	num := userDevicePo.DeleteUserDevicesDB(dbc,dids)
	if num != int64(len(dids)) {
		dbc.Rollback()
		return &pb.ShareReply{Code: 33014}, nil
	}
	//todo 查询Member的宠物
	userPetPo = &db.UserPetPo{Uid:req.MemberUid}
	userPetPos1, err := userPetPo.GetUserPetsDB()
	if err != nil {
		dbc.Rollback()
		log.Info("ShareUserResourceByUid-err:", err)
		return &pb.ShareReply{Code: 10001}, nil
	}
	log.Info("Member-userPetPos1:",len(userPetPos1))
	//todo 查询Owner的宠物
	userPetPo.Uid = req.OwnerUid
	userPetPos, err := userPetPo.GetUserPetsDB()
	if err != nil {
		dbc.Rollback()
		return &pb.ShareReply{Code: 10001}, nil
	}
	log.Info("Owner-userPetPos:",len(userPetPos))
	//todo 筛选
	userPetPos2 := userPetPo.SliceDiff2(userPetPos, userPetPos1)
	log.Info("userPetPos2:",len(userPetPos2))
	if len(userPetPos) == len(userPetPos2) {
		//todo 没有分享则删除共享关系与共享人端
		num = shareUserPo.DeleteShareDB(dbc)
		if num != 1 {
			dbc.Rollback()
			return &pb.ShareReply{Code: 33012}, nil
		}
		userDevicePo.Uid = req.OwnerUid
		userDevicePos, err := userDevicePo.GetUserDevicesDB()
		if err != nil {
			dbc.Rollback()
			return &pb.ShareReply{Code: 10001}, nil
		}
		var udids []int32
		for _, v := range userDevicePos {
			if 0 == v.Permit {
				udids = append(udids,v.Did)
			}
		}
		devicePo := &db.DevicePo{}
		devicePos, err := devicePo.GetDevicesDB(udids)
		for _, v := range devicePos {
			if v.Types == 1 {
				userDevicePo.Uid = req.MemberUid
				userDevicePo.Did = v.Did
				num = userDevicePo.DeleteUserDeviceDB(dbc)
				if num != 1 {
					dbc.Rollback()
					return &pb.ShareReply{Code: 10001}, nil
				}
				break
			}
		}
	}
	dbc.Commit()
	return &pb.ShareReply{Code: 10000}, nil
}

//批量查询共享信息
func (this *ShareRpc) GetShare(ctx context.Context, req *pb.ShareRequest) (*pb.ShareMapReply, error) {
	log.Info("GetShare-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.ShareMapReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.OwnerUid) {
		return &pb.ShareMapReply{Code: util.User_params_err}, nil
	}
	//todo 查询共享记录
	shareUserPo := &db.ShareUserPo{Uid:req.OwnerUid,Fuid:req.OwnerUid}
	shareUserPos, err := shareUserPo.GetSharesDB()
	if err != nil {
		log.Info("GetShare-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.ShareMapReply{Code: 33012}, nil
		}
		return &pb.ShareMapReply{Code: 10001}, nil
	}
	if len(shareUserPos) == 0 {
		return &pb.ShareMapReply{Code: 33013}, nil
	}
	//todo 查询用户的宠物
	userPetPo := &db.UserPetPo{Uid:req.OwnerUid}
	userPetPos, err := userPetPo.GetUserPetsDB()
	if err != nil {
		log.Info("GetShare-err:", err)
		if err.Error() ==  core.ConstStr.NotFound {
			return &pb.ShareMapReply{Code: 33012}, nil
		}
		return &pb.ShareMapReply{Code: 10001}, nil
	}
	//todo 查询宠物的用户
	var pids []int32
	for _,v := range userPetPos {
		pids = append(pids,v.Pid)
	}
	userPetPo.Uid = req.OwnerUid
	userPetPos, err = userPetPo.GetUserPetsMemberDB(pids)
	if err != nil {
		return &pb.ShareMapReply{Code: 10001}, nil
	}
	var uids []int32
	for _,v := range userPetPos {
		uids = append(uids,v.Uid)
	}
	uidSet := util.NewSet()
	uidSet.Adds(uids)
	var accounts []*db.Account
	//todo 查询共享宠物的用户信息
	account := &db.Account{}
	accounts, err = account.GetBatchAccount(uidSet.List())
	if err != nil && err.Error() != core.ConstStr.NotFound {
		return &pb.ShareMapReply{Code: 10001}, nil
	}
	log.Info("accounts:",accounts)
	//todo 聚合信息
	var nickname string
	shareRes := make(map[int32]*pb.ShareReply)
	for k,v := range pids {
		shareRe := &pb.ShareReply{}
		ownerInfo := &pb.AccountReply{}
		var members []*pb.ShareMember
		//todo 遍历用户与宠物关联关系
		for _, up := range userPetPos {
			//todo 遍历共享宠物用户信息
			for _, as := range accounts {
				if v == up.Pid {
					var shareTime int64
					for _, su := range shareUserPos {
						if su.Fuid == as.Id {
							shareTime = su.UpdateTime.Unix()
							break
						}
					}
					nickname, _ = url.QueryUnescape(as.Nickname)
					if as.Id == up.Uid && up.Permit == 0 {
						//todo 遍历共享记录
						ownerInfo = &pb.AccountReply{Uid:as.Id, Nickname:nickname, Avatar:core.ConstStr.FileServer+as.Avatar}
						break
					}
					if as.Id == up.Uid && up.Permit != 0 {
						member := &pb.ShareMember{MemberInfo:&pb.AccountReply{Uid:as.Id, Nickname:nickname, Avatar:core.ConstStr.FileServer+as.Avatar},
							ShareTime:shareTime}
						members = append(members, member)
						break
					}
				}
			}
		}
		shareRe.Pid = v
		shareRe.OwnerInfo = ownerInfo
		shareRe.Members = members
		shareRes[int32(k)] = shareRe
	}
	return &pb.ShareMapReply{Code: 10000, Shares:shareRes}, nil
}

//分享设备与宠物 -------------------
func (this *ShareRpc) ShareUserResourceByUid(ctx context.Context, req *pb.ShareRequest) (*pb.ShareReply, error) {
	log.Info("ShareUserResourceByUid-req:", req)
	//if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
	//	return &pb.DevicesReply{Code: util.Source_err_empty}, nil
	//}
	//if util.VerifyParamsUInt32(req.GetUid(),req.GetTouid()) {
	//	return &pb.DevicesReply{Code: util.Params_err_empty}, nil
	//}
	//userDevicePo := &db.UserDevicePo{Uid:req.Touid}
	//userPetPo := &db.UserPetPo{Uid:req.Touid}
	////todo 查询对方的设备与宠物
	//userDevicePos, err := userDevicePo.GetUserDevicesDB()
	//userPetPos, err := userPetPo.GetUserPetsDB()
	//if err != nil {
	//	log.Info("ShareUserResourceByUid-err:", err)
	//	return &pb.DevicesReply{Code: 10001}, nil
	//}
	//log.Info("userDevicePos:",len(userDevicePos),",userPetPos:",len(userPetPos))
	////todo 查询自己的设备与宠物
	//userDevicePo.Uid = req.Uid
	//userDevicePos1, err := userDevicePo.GetUserDevicesDB()
	//if err != nil {
	//	return &pb.DevicesReply{Code: 10001}, nil
	//}
	//userPetPo.Uid = req.Uid
	//userPetPos1, err := userPetPo.GetUserPetsDB()
	//if err != nil {
	//	return &pb.DevicesReply{Code: 10001}, nil
	//}
	//log.Info("userDevicePos1:",len(userDevicePos1),",userPetPos1:",len(userPetPos1))
	////todo 筛选要分享的资源
	//userDevicePos1 = userDevicePo.SliceDiff1(userDevicePos, userDevicePos1)
	//userPetPos1 = userPetPo.SliceDiff2(userPetPos, userPetPos1)
	////todo 添加设备与宠物分享
	//err = userDevicePo.SetUserDevicesDB(userDevicePos1)
	//if err != nil {
	//	return &pb.DevicesReply{Code: 10001}, nil
	//}
	//err = userPetPo.SetUserPetsDB(userPetPos1)
	//if err != nil {
	//	return &pb.DevicesReply{Code: 10001}, nil
	//}
	return &pb.ShareReply{Code: 10000}, nil
}

//修改共享信息----------
func (this *ShareRpc) UpdateShare(ctx context.Context, req *pb.ShareRequest) (*pb.ShareReply, error) {
	log.Info("UpdatePetInfoByPid:", req)
	//if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
	//	return &pb.PetInfoReply{Code: util.Source_err_empty}, nil
	//}
	//if util.VerifyParamsUInt32(req.GetPid()) {
	//	return &pb.PetInfoReply{Code: util.Source_err_empty}, nil
	//}
	////todo 检查权限
	//userPetPo := &db.UserPetPo{Uid:req.Uid,Pid:req.Pid}
	//err := userPetPo.GetUserPetDB()
	//if err != nil {
	//	log.Info("UpdatePetInfoByPid-err:", err)
	//	if err.Error() ==  util.NotFound {
	//		return &pb.PetInfoReply{Code: 33012}, nil
	//	}
	//	return &pb.PetInfoReply{Code: 10001}, nil
	//}
	//petinfoPo := db.PetInfoPo{Pid: req.Pid, Avatar: req.Avatar, Nickname: req.Nickname, Breed: req.Breed, Gender: req.Gender,
	//	Weight: req.Weight, Somatotype: req.Somatotype, UpdateTime: util.GetNowTime()}
	//if req.Birthday > 0 {
	//	birthday, err := util.Int64ToTime(req.Birthday)
	//	if err != nil {
	//		return &pb.PetInfoReply{Code: 10001}, nil
	//	}
	//	petinfoPo.Birthday = birthday
	//}
	//err = petinfoPo.UpdatePetInfoDB()
	//if err != nil {
	//	return &pb.PetInfoReply{Code: 10001}, nil
	//}
	//if req.Avatar != "" {
	//	return &pb.PetInfoReply{Code: 10000,Avatar: util.FileServer +petinfoPo.Avatar}, nil
	//}
	return &pb.ShareReply{Code: 10000}, nil
}
