package service

import (
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"petfone-rpc/db"
	"petfone-rpc/pb"
	"petfone-rpc/util"
	"petfone-rpc/core"
)

/**
常见问题信息
 */
type FaqCommonRpc struct {
}

//根据id获取常见问题信息
func (this *FaqCommonRpc) GetFaqCommonById(ctx context.Context, req *pb.FaqCommonRequest) (*pb.FaqCommonReply, error) {
	log.Info("req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.FaqCommonReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsUInt32(req.GetId()) {
		return &pb.FaqCommonReply{Code: util.Source_err_empty}, nil
	}
	faqc := &db.FaqCommonPo{Id: req.GetId(), DataState: 1}
	err := faqc.GetFAQById()
	log.Info("GetFaqCommonById-faqc:", faqc)
	if err != nil {
		if err.Error() == core.ConstStr.NotFound {
			return &pb.FaqCommonReply{Code: 33011}, nil
		}
		log.Info("GetFaqCommonById-err", err)
		return &pb.FaqCommonReply{Code: 10001}, nil
	}
	return &pb.FaqCommonReply{Code:10000, Id:faqc.Id,Name:faqc.NameCn, Info:faqc.InfoCn,
		NameCn:faqc.NameCn, InfoCn:faqc.InfoCn, NameEn:faqc.NameEn, InfoEn:faqc.InfoEn, Parent:faqc.Parent}, nil
}

//根据name获取常见问题信息
func (this *FaqCommonRpc) GetFaqCommonByKeyword(ctx context.Context, req *pb.FaqCommonRequest) (*pb.FaqCommonsReply, error) {
	log.Info("GetFaqCommonByKeyword-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.FaqCommonsReply{Code: util.Source_err_empty}, nil
	}
	if util.VerifyParamsStr(req.GetNameCn()) {
		return &pb.FaqCommonsReply{Code: util.Source_err_empty}, nil
	}
	faqc := &db.FaqCommonPo{NameCn: req.GetNameCn(), DataState: 1}
	faqcs, err := faqc.GetFAQByKeyword()
	log.Info("GetFaqCommonByKeyword-faqc:", faqc)
	if err != nil {
		log.Info("GetFaqCommonByKeyword-err", err)
		return &pb.FaqCommonsReply{Code: 10001}, nil
	}
	if len(faqcs) == 0 {
		log.Info("GetFaqCommonByKeyword-nil")
		return &pb.FaqCommonsReply{Code: 33013}, nil
	}
	var faqCmRes []*pb.FaqCommonReply
	for _, v := range faqcs {
		faqCmRes = append(faqCmRes,&pb.FaqCommonReply{Id: v.Id, Name:v.NameCn, Info:v.InfoCn, Parent:v.Parent})
	}
	return &pb.FaqCommonsReply{Code:10000, Faqcs:faqCmRes}, nil
}

//批量常见问题信息
func (this *FaqCommonRpc) GetFaqCommons(ctx context.Context, req *pb.FaqCommonRequest) (*pb.FaqCommonsReply, error) {
	log.Info("GetFaqCommons-req:", req)
	if util.VerifyParamsStr(req.GetSource()) || "AgIDAA==" != req.GetSource() {
		return &pb.FaqCommonsReply{Code: util.Source_err_empty}, nil
	}
	faqc := &db.FaqCommonPo{DataState: 1}
	faqcs, err := faqc.GetFAQs()
	log.Info("GetFaqCommons-faqc:", len(faqcs))
	if err != nil || len(faqcs) == 0{
		log.Info("GetFaqCommons-err", err)
		return &pb.FaqCommonsReply{Code: 10001}, nil
	}
	var faqCmRes []*pb.FaqCommonReply
	for _, v := range faqcs {
		faqCmRes = append(faqCmRes,&pb.FaqCommonReply{Id: v.Id,
			NameCn:v.NameCn, InfoCn:v.InfoCn, NameEn:v.NameEn, InfoEn:v.InfoEn})
	}
	return &pb.FaqCommonsReply{Code:10000, Faqcs:faqCmRes}, nil
}
