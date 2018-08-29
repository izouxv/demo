package controller

import (
	log "github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"petfone-http/pb"
	"petfone-http/po"
	"petfone-http/result"
	"petfone-http/rpc"
	"petfone-http/util"
)

/**
通知信息
*/

//添加通知
func SetNotice(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("SetNotice...")
	sso := util.GetContext(req)
	log.Info("SetNotice-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	notice := new(pb.NoticeRequest)
	nofund := util.GetHttpData(req, util.ReqMethodJson, &notice)
	if nofund {
		result.RESC(21002, res)
		return
	}
	//todo 校验消息类型
	switch notice.Types {
	//todo 类型为调动资源时校验权限
	case 3:
		if notice.Nstate != 3 {
			result.RESC(21001, res)
			return
		}
		//todo 对方用户存在 并且不能是自己
		ssoTo := &pb.SsoRequest{Source: sso.Source, Username: notice.Tou}
		ssoR := rpc.SsoRpc(ssoTo, "GetUserByName")
		if ssoR.Code != 10000 {
			result.RESC(ssoR.Code, res)
			return
		}
		if sso.Uid == ssoR.Uid {
			result.RESC(21002, res)
			return
		}
		//to do 校验用户对资源的权限
		//num, err := db.Redis_Zcount(6379,sso.Source[:2]+db.UDevices+util.Int32ToStr(sso.Uid), 1, 1)
		//if err != nil {
		//	result.RESC(10001, res)
		//	return
		//}
		//if num < 1 {
		//	result.RESC(21008, res)
		//	return
		//}
		//todo 组成描述语句
		notice.Info = "用户 " + sso.Username + " 向你共享宠物与设备"
		notice.Tos = ssoR.Uid
		break
		//todo 其它类型为不允许
	default:
		result.RESC(21002, res)
		return
	}
	notice.Source = sso.Source
	notice.Froms = sso.Uid
	//调用rpc
	noticeRe := rpc.NoticeRpc(notice, "SetNotice")
	if noticeRe.Code != 10000 {
		result.RESC(noticeRe.Code, res)
		return
	}
	devicePo := &po.NoticePo{Id: noticeRe.Id, Info: "" + sso.Username + "", StartTime: noticeRe.Times}
	result.REST(devicePo, res)
}

//删除通知
func DeleteNotice(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("DeleteDeviceP...")
	idStr := params.ByName("id")
	sso := util.GetContext(req)
	log.Info("DeleteDeviceP-sso:", sso, ",idStr:", idStr)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	notice := new(pb.NoticeRequest)
	nofund := util.GetHttpData(req, util.ReqMethodJson, &notice)
	if nofund {
		result.RESC(21002, res)
		return
	}
	log.Info("DeleteDeviceP-deviceJson:", notice)
	var id int32
	var err error
	if id, err = util.StrToInt32(idStr); err != nil {
		result.RESC(10001, res)
		return
	}
	notice.Source = sso.Source
	notice.Froms = sso.Uid
	notice.Id = id
	//调用rpc
	deviceRe := rpc.NoticeRpc(notice, "DeleteNotice")
	result.RESC(deviceRe.Code, res)
	return
}

//对通知应答
func UpdateNotice(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("UpdateDeviceP...")
	idStr := params.ByName("id")
	sso := util.GetContext(req)
	log.Info("UpdateDeviceP-sso:", sso, ",idStr:", idStr)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	notice := new(pb.NoticeRequest)
	nofund := util.GetHttpData(req, util.ReqMethodJson, &notice)
	if nofund {
		result.RESC(21002, res)
		return
	}
	var id int32
	var err error
	if id, err = util.StrToInt32(idStr); err != nil {
		result.RESC(10001, res)
		return
	}
	log.Info("UpdateDeviceP-noticeJson:", notice)
	if util.VerifyParamsUInt32(notice.Id, notice.Froms, notice.Nstate, notice.Types) {
		result.RESC(20015, res)
		return
	}
	notice.Source = sso.Source
	notice.Tos = sso.Uid
	notice.Id = id
	log.Info("UpdateDeviceP-noticeJson:", notice)
	//todo 判断消息类型
	switch notice.Types {
	//todo 调动资源时校验权限
	case 3:
		switch notice.Nstate {
		case 1:
			notice.Types = 1
			//todo 修改通知消息状态
			noticeRe := rpc.NoticeRpc(notice, "UpdateNotice")
			//todo 若不成功，则返回
			if noticeRe.Code != 10000 {
				result.RESC(noticeRe.Code, res)
				return
			}
			//todo 添加该用户对资源的权限
			device := &pb.DeviceRequest{Source: sso.Source, Uid: sso.Uid, Touid: notice.Froms}
			deviceRe := rpc.DevicesRpc(device, "ShareUserResourceByUid")
			if deviceRe.Code != 10000 {
				result.RESC(10001, res)
				return
			}
			break
		case 2:
			//todo 修改通知消息状态
			deviceRe := rpc.NoticeRpc(notice, "UpdateNotice")
			//todo 若不成功，则返回
			if deviceRe.Code != 10000 {
				result.RESC(deviceRe.Code, res)
				return
			}
			break
		}
		break
	default:
		result.RESC(21002, res)
		return
	}
	result.RESC(10000, res)
}

//获取通知
func GetNotice(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("GetNotice...")
	sso := util.GetContext(req)
	log.Info("GetNotice-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	notice := new(pb.NoticeRequest)
	notice.Source = sso.Source
	notice.Froms = sso.Uid
	log.Info("GetNotice-deviceJson:", notice)
	//调用rpc
	noticeMapRe := rpc.NoticesRpc(notice, "GetNotice")
	log.Info("GetNotice-noticeMapRe:", noticeMapRe)
	var arr []po.NoticePo
	for _, v := range noticeMapRe.Notices {
		arr = append(arr, po.NoticePo{Id: v.Id, Froms: v.Froms, Tos: v.Tos, Nstate: v.Nstate, Types: v.Types, Info: v.Info, StartTime: v.Times})
	}
	result.REST(arr, res)
}
