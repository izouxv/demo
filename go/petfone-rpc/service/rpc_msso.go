package service

import (
	"golang.org/x/net/context"
	log "github.com/cihub/seelog"
	"petfone-rpc/db"
	"petfone-rpc/pb"
	"petfone-rpc/util"
	"net/url"
	"petfone-rpc/core"
)

type Rpc_msso struct {
}

//分页查询sso与操作信息
func (this *Rpc_msso) GetPageSsoInfos(ctx context.Context, req *pb.PageRequest) (*pb.PageSsoReply, error) {
	log.Info("GetPageSsoInfos-req:",req)
	if req.Source == "" || req.Page <= 0 || req.Count <= 0 || req.Count > 100 {
		log.Info("req.Ssos is empty")
		return &pb.PageSsoReply{Code: util.Params_err_empty }, nil
	}
	sso := &db.Sso{}
	ssos,totalCount,err := sso.GetPageSsoInfo((req.Page-1)* req.Count, req.Count, req.Sort.String())
	if err != nil {
		return &pb.PageSsoReply{Code: util.System_error }, nil
	}
	if len(ssos) == 0 {
		log.Info("GetPageSsoInfos-err-GetBatchSsoInfo size is 0")
		return &pb.PageSsoReply{Code: util.Success}, nil
	}
	ssoRe, err := getMsso(ssos...)
	return &pb.PageSsoReply{Code: util.Success, MSsos:ssoRe, TotalCount:totalCount}, nil
}

//查询sso信息（uid/username）
func (this *Rpc_msso) SearchSsoInfo(ctx context.Context, req *pb.MSsoInfo) (*pb.MSsoInfo, error) {
	log.Info("SearchSsoInfo-req:",req)
	if req.Source == "" {
		log.Info("req.Ssos is empty")
		return &pb.MSsoInfo{Code: util.Params_err_empty }, nil
	}
	sso := &db.Sso{Id:req.Uid,Username:req.Username}
	var err error
	if sso.Id != 0 {
		err = sso.GetUserInfoById()
	} else {
		err = sso.GetByName()
	}
	if err != nil {
		if err.Error() == core.ConstStr.NotFound {
			return &pb.MSsoInfo{Code: util.User_unexist }, nil
		}
		log.Error("SearchSsoInfo err:",err)
		return &pb.MSsoInfo{Code: util.System_error }, nil
	}
	ssoRe, err := getMsso(sso)
	ssoRe[0].Code = 10000
	return ssoRe[0], nil
}

//删除账号信息username
func (this *Rpc_msso) DeleteAccount(ctx context.Context, req *pb.MSsoInfo) (*pb.MSsoReply, error) {
	log.Info("DeleteAccount-req:",req)
	if req.Source == "" {
		log.Info("DeleteAccount source is empty")
		return &pb.MSsoReply{Code: util.Params_err_empty }, nil
	}
	sso := &db.Sso{Username:req.Username}
	err := sso.GetByName()
	if err != nil {
		if err.Error() == core.ConstStr.NotFound {
			return &pb.MSsoReply{Code: util.User_exists}, nil
		}
		log.Error("DeleteAccount GetByName-err:",err)
		return &pb.MSsoReply{Code: util.System_error }, nil
	}
	dbc := core.MysqlClient.Begin()
	err = sso.UpdateUserName(req.Username+"_test_"+util.TimeToStr())
	if err != nil {
		dbc.Rollback()
		log.Error("DeleteAccount UpdateUserName-err:",err)
		return &pb.MSsoReply{Code: util.System_error }, nil
	}
	dbc.Commit()
	return &pb.MSsoReply{Code: util.Success}, nil
}

// 分页查询设备信息
func (this *Rpc_msso) GetPageDevices(ctx context.Context, req *pb.PageRequest) (*pb.BatchDeviceRe, error) {
	log.Info("GetPageDevices-req:",req)
	if req.Source == "" || req.Page <= 0 || req.Count <= 0 || req.Count > 100 {
		log.Info("GetPageDevices Params err or isEmpty")
		return &pb.BatchDeviceRe{Code: util.Params_err_empty }, nil
	}
	devicePo := &db.DevicePo{}
	devicePos, totalCount, err := devicePo.GetPageDeviceDB(
		(req.Page-1)*req.Count, req.Count,req.Sort.String())
	if err != nil {
		log.Info("GetPageDevices-GetPageDeviceDB err:",err)
		return &pb.BatchDeviceRe{Code: util.System_error }, nil
	}
	if len(devicePos) == 0 {
		log.Info("GetPageDevices-err-GetPageDeviceDB size is 0")
		return &pb.BatchDeviceRe{Code: util.Success}, nil
	}
	userDevicePo := &db.UserDevicePo{}
	var dids []int32
	for _, v := range devicePos {
		dids = append(dids,v.Did)
	}
	userDevicePos, err := userDevicePo.GetBatchDeviceMasterDB(dids)
	if err != nil && err.Error() != core.ConstStr.NotFound {
		log.Info("GetPageDevices-GetBatchDeviceMasterDB err:",err)
		return &pb.BatchDeviceRe{Code: util.System_error }, nil
	}
	var deviceRes []*pb.DeviceReply
	for _, v := range devicePos {
		deviceRe := &pb.DeviceReply{
			Did:v.Did,
			Uid:v.Uid,
			Sn:v.Sn,
			DeviceMac:v.DeviceMac,
			DeviceName:v.DeviceName,
			DeviceVersion:v.DeviceVersion,
			SoftwareVersion:v.SoftwareVersion,
			Permit:0,
			Types:v.Types,
		}
		for _, vu := range userDevicePos {
			if v.Did == vu.Did {
				deviceRe.Uid = vu.Uid
				break
			}
		}
		deviceRes = append(deviceRes, deviceRe)
	}
	batchDeviceRe := &pb.BatchDeviceRe{Code:util.Success}
	batchDeviceRe.TotalCount = totalCount
	batchDeviceRe.Devices = deviceRes
	return batchDeviceRe, nil
}

//搜索设备信息（sn/msc/did）
func (this *Rpc_msso) SearchDevice(ctx context.Context, req *pb.DeviceRequest) (*pb.DeviceReply, error) {
	log.Info("SearchDevice-req:",req)
	if req.Source == "" {
		log.Info("SearchDevice-Source is empty")
		return &pb.DeviceReply{Code: util.Params_err_empty }, nil
	}
	flag1,flag2,flag3 := req.DeviceMac == "",req.GetSn() == "",req.Did == 0
	if !(flag1 && flag2 && flag3) {
		devicePo := &db.DevicePo{}
		var err error
		if req.DeviceMac != "" {
			devicePo.DeviceMac = req.DeviceMac
			err = devicePo.GetDeviceByMac()
			goto Returns
		}
		if req.Sn != "" {
			devicePo.Sn = req.Sn
			err = devicePo.GetDeviceBySn()
			goto Returns
		}
		if req.Did != 0 {
			devicePo.Did = req.Did
			err = devicePo.GetDeviceByDid()
			goto Returns
		}
	Returns:
		if err != nil {
			if err.Error() == core.ConstStr.NotFound {
				return &pb.DeviceReply{Code: 33013 }, nil
			}
			log.Error("GetPageDevices-GetDevice err:",err)
			return &pb.DeviceReply{Code: util.System_error }, nil
		}
		userDevicePo := &db.UserDevicePo{Did:devicePo.Did}
		err = userDevicePo.GetDeviceMasterDB()
		if err != nil && err.Error() != core.ConstStr.NotFound {
			log.Info("GetPageDevices-GetBatchDeviceMasterDB err:",err)
			return &pb.DeviceReply{Code: util.System_error }, nil
		}
		devicePetPo := &db.DevicePetPo{Did:devicePo.Did}
		err = devicePetPo.GetDevicePetDB()
		if err != nil && err.Error() != core.ConstStr.NotFound {
			log.Info("GetPageDevices-GetDevicePetDB err:",err)
			return &pb.DeviceReply{Code: util.System_error }, nil
		}
		return &pb.DeviceReply{Did:devicePo.Did, Uid:userDevicePo.Uid, Pid:devicePetPo.Pid, Sn:devicePo.Sn,
			DeviceMac:devicePo.DeviceMac, DeviceName:devicePo.DeviceName, DeviceVersion:devicePo.DeviceVersion,
			SoftwareVersion:devicePo.SoftwareVersion, Permit:0, Types:devicePo.Types,}, nil
	}
	log.Info("SearchDevice Params_err_empty")
	return &pb.DeviceReply{Code: util.Params_err_empty }, nil
}

//------------------   分割线   --------------------------------

//批量uid查询用户操作记录
func getMsso(ssos ...*db.Sso) (ssoRe []*pb.MSsoInfo, err error) {
	for _, ssov := range ssos {
		nickname, _ := url.QueryUnescape(ssov.Nickname)
		ip, _ := util.IPToString(ssov.RegIp)
		mSsoRe := &pb.MSsoInfo{Uid: ssov.Id, Username: ssov.Username, Nickname: nickname, State: ssov.State,
			RegTime: ssov.CreateTime.Unix(),RegIP:ip,RegAddr:ssov.RegAddr}
		if ssov.State != 3 || (util.GetNowTime().Unix() - ssov.LastLoginTime.Unix()) > 3600*24*7 {
			mSsoRe.LoginState = 2
		} else {
			mSsoRe.LoginState = 1
		}
		mSsoRe.LoginTime = ssov.CreateTime.Unix()
		loginIp, _ := util.IPToString(ssov.LastLoginIp)
		mSsoRe.NewIP = loginIp
		mSsoRe.NewAddr = ssov.LastLoginAddr
		mSsoRe.DevInfo = ssov.LastLoginDevInfo
		ssoRe = append(ssoRe, mSsoRe)
	}
	return ssoRe,nil
}
