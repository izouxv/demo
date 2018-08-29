package service

import (
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"petfone-rpc/db"
	"petfone-rpc/pb"
	"petfone-rpc/util"
	"time"
	"net/url"
)

type NoticeRpc struct {
}

//发送通知
func (this *NoticeRpc) SetNotice(ctx context.Context, req *pb.NoticeRequest) (*pb.NoticeReply, error) {
	log.Info("SetNotice-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.NoticeReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetId(),req.GetFroms(),req.GetTos(),req.GetNstate(),req.GetTypes()) {
		return &pb.NoticeReply{Code: util.Params_err_empty}, nil
	}
	if util.VerifyParamsStr(req.GetInfo()) {
		return &pb.NoticeReply{Code: util.Params_err_empty}, nil
	}
	////todo 校验权限
	//num1, num2, err := db.Redis_ZscoreZcard(6379, req.Source[:2]+db.DUsers+methods.Int32ToStr(req.GetDid()), util.Int32ToStr(req.Uid))
	//if err != nil {
	//	log.Info("SetNotice-err:", err)
	//	return &pb.NoticeReply{Code: 10001}, nil
	//}
	//if num1 < 1 || num2 == 0 {
	//	log.Info("没有权限")
	//	return &pb.NoticeReply{Code: util.User_params_err}, nil
	//}
	//todo 向对方发送消息
	notice := &db.NoticePo{From:req.Froms, To:req.Tos, State:req.Nstate, Types:req.Types, Info:url.QueryEscape(req.Info),
		CreationTime:util.GetNowTime(), UpdateTime:util.GetNowTime(), DataState:1}
	err := notice.SetNoticePo()
	if err != nil {
		log.Info("SetNotice-err:", err)
		return &pb.NoticeReply{Code: 10001}, nil
	}
	return &pb.NoticeReply{Code: 10000, Id:notice.Id, Times:notice.CreationTime.Unix()}, nil
}

//删除通知消息
func (this *NoticeRpc) DeleteNotice(ctx context.Context, req *pb.NoticeRequest) (*pb.NoticeReply, error) {
	log.Info("req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.NoticeReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetFroms(), req.GetId()) {
		return &pb.NoticeReply{Code: util.Params_err_empty}, nil
	}
	noticePo := &db.NoticePo{Id: req.Id, From:req.Froms, UpdateTime: util.GetNowTime(), DataState:2}
	t := time.Now()
	err := noticePo.DeleteNoticePo()
	log.Info("time:", time.Now().Sub(t))
	if err != nil {
		return &pb.NoticeReply{Code: 10001}, nil
	}
	return &pb.NoticeReply{Code: 10000}, nil
}

//修改通知
func (this *NoticeRpc) UpdateNotice(ctx context.Context, req *pb.NoticeRequest) (*pb.NoticeReply, error) {
	log.Info("SetNotice-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.NoticeReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetId(),req.GetFroms(),req.GetNstate(),req.GetTypes()) {
		return &pb.NoticeReply{Code: util.Params_err_empty}, nil
	}
	//todo 修改消息状态
	notice := &db.NoticePo{Id:req.Id, From:req.Froms, To:req.Tos, State:req.Nstate, UpdateTime:util.GetNowTime()}
	num := notice.UpdateNoticePo()
	if num != 1 {
		log.Info("SetNotice-num:", num)
		return &pb.NoticeReply{Code: util.User_params_err}, nil
	}
	return &pb.NoticeReply{Code: 10000}, nil
}

//获取通知
func (this *NoticeRpc) GetNotice(ctx context.Context, req *pb.NoticeRequest) (*pb.NoticeMapReply, error) {
	log.Info("GetDevices-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.NoticeMapReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetFroms()) {
		return &pb.NoticeMapReply{Code: util.Params_err_empty}, nil
	}
	notice := &db.NoticePo{To:req.Froms}
	notices, err := notice.GetNoticePo()
	if err != nil {
		log.Info("err", err)
		return &pb.NoticeMapReply{Code: 10001}, nil
	}
	log.Info("notices:", notices)
	noticeMap := make(map[int32]*pb.NoticeReply)
	for k, v := range notices {
		info, err := url.QueryUnescape(v.Info)
		if err != nil {
			log.Info("err", err)
			return &pb.NoticeMapReply{Code: 10001}, nil
		}
		noticeMap[int32(k)] = &pb.NoticeReply{
			Id:v.Id, Froms:v.From, Tos:v.To, Nstate:v.State, Types:v.Types, Info:info, Times:v.CreationTime.Unix(),
		}
	}
	return &pb.NoticeMapReply{Code: 10000, Notices:noticeMap}, nil
}