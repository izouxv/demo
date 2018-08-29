package rpc

import (
	pb "account-domain-rpc/api/feedback"
	"account-domain-rpc/module"
	"account-domain-rpc/storage"
	"account-domain-rpc/util"
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin/json"
	"golang.org/x/net/context"
	"github.com/jinzhu/gorm"
)

type FeedBackServer struct{}

func (fd *FeedBackServer) AddFeedback(ctx context.Context, in *pb.AddFeedbackRequest) (*pb.AddFeedbackReply, error) {
	log.Infof("Start AddFeedback %#v",in)
	if in.MobileInfo == "" || in.Description == "" || in.AppInfo == "" || in.Source == "" {
		log.Infof("输入参数异常 source:(%s),description :(%s),MobileInfo:(%s),AppInfo:(%s)", in.Source, in.Description, in.MobileInfo, in.AppInfo)
		return &pb.AddFeedbackReply{ErrorCode: util.Input_parameter_error}, nil
	}
	if util.SourceToTid[in.Source] == 0{
		log.Infof("输入source有误:(%s),", in.Source)
		return &pb.AddFeedbackReply{ErrorCode: util.Input_parameter_error}, nil
	}
	feedback := storage.FeedBack{
		Description: in.Description,
		DeviceInfo:  in.DeviceInfo,
		UserInfo:    in.UserInfo,
		AppInfo:     in.AppInfo,
		ExtendInfo:  in.ExtendInfo,
		MobileInfo:  in.MobileInfo,
		Tid:         util.SourceToTid[in.Source],
		Type:        1,
		Contact:     in.Contact}
	if in.Files != nil {
		f, err := json.Marshal(in.Files)
		if err != nil {
			log.Infof("文件json.Marshal有误, err (%s) ", err)
			return &pb.AddFeedbackReply{ErrorCode: util.System_error}, nil
		}
		feedback.Files = string(f)
	}
	if err := feedback.CreateFeedBack(module.MysqlClient()); err != nil {
		log.Infof("创建反馈异常,error:(%s),feedback :(%s)", err, feedback)
		return &pb.AddFeedbackReply{ErrorCode: util.System_error}, nil
	}
	log.Info("返回的工单ID:", feedback.Id)
	return &pb.AddFeedbackReply{ErrorCode: util.Successfull, Id: feedback.Id}, nil
}

func (fd *FeedBackServer) AddFeedbackBaseTenant(ctx context.Context, in *pb.AddFeedbackBaseTenantRequest) (*pb.AddFeedbackBaseTenantReply, error) {
	log.Infof("Start AddFeedbackBaseTenant %#v", in)
	if in.Tid  == 0 || in.Type == 0 || in.Description == ""{
		log.Infof("输入参数异常 tid:(%s),description :(%s),BugType:(%s)", in.Tid, in.Description, in.Type)
		return &pb.AddFeedbackBaseTenantReply{ErrorCode: util.Input_parameter_error}, nil
	}
	feedback := storage.FeedBack{
		Description: in.Description,
		Type:        in.Type,
		Tid:         in.Tid,
		Files:       in.Files}
	log.Debug("AddFeedbackBaseTenant:", feedback)
	if err := feedback.CreateFeedBackBaseTenant(module.MysqlClient()); err != nil {
		log.Infof("创建反馈异常,error:(%s),feedback :(%s)", err, feedback)
		return &pb.AddFeedbackBaseTenantReply{ErrorCode: util.System_error}, nil
	}
	return &pb.AddFeedbackBaseTenantReply{ErrorCode: util.Successfull}, nil
}

func (fd *FeedBackServer) GetFeedbacks(ctx context.Context, in *pb.GetFeedbacksRequest) (*pb.GetFeedbacksReply, error) {
	log.Infof("Start CetFeedbacks %#v",in)
	if in.Page == 0 || in.Count == 0 || in.Tid == 0  {
		log.Infof("输入参数异常 page :(%s),count :(%s),tenantId :(%s)", in.Page, in.Count,in.Tid)
		return &pb.GetFeedbacksReply{ErrorCode: util.Input_parameter_error}, nil
	}
	feedback := storage.FeedBack{Tid:in.Tid}
	feedbacks, totalCount, err := feedback.GetFeedbacks(module.MysqlClient(), in.Count, in.Page)
	if err != nil {
		log.Infof("分页获取反馈异常.error is (%s),page is (%d),count is (%d)  ", err, in.Page, in.Count)
		return &pb.GetFeedbacksReply{}, err
	}
	feedbacksResp := make([]*pb.Feedback, 0)
	for _, v := range feedbacks {
		feedbacksResp = append(feedbacksResp, &pb.Feedback{
			Id:          v.Id,
			Tid:         v.Tid,
			Description: v.Description,
			AppInfo:     v.AppInfo,
			UserInfo:    v.UserInfo,
			MobileInfo:  v.MobileInfo,
			ExtendInfo:  v.ExtendInfo,
			Contact:     v.Contact,
			DeviceInfo:  v.DeviceInfo,
			Files:       v.Files,
			Type:        v.Type,
			CreateTime:  v.CreateTime.Unix(),
			UpdateTime:  v.UpdateTime.Unix(),
		})
	}
	return &pb.GetFeedbacksReply{ErrorCode: util.Successfull, FeedBack: feedbacksResp, TotalCount: totalCount}, nil
}

func (fd *FeedBackServer) GetFeedbacksByType(ctx context.Context, in *pb.GetFeedbacksByTypeRequest) (*pb.GetFeedbacksByTypeReply, error) {
	log.Infof("Start GetFeedbacksByType %#v", in)
	if in.Tid == 0 || in.Page == 0 || in.Count == 0 {
		log.Infof("输入参数错误 %#v", in)
		return &pb.GetFeedbacksByTypeReply{ErrorCode: util.Input_parameter_error}, nil
	}
	feedback := storage.FeedBack{Tid:in.Tid,Type:in.Type}
	feedbacks, totalCount, err := feedback.GetFeedbackByType(module.MysqlClient(), in.Count, in.Page)
	if err != nil {
		log.Infof("分类分页获取工单异常,error is (%s),page is (%d),count is (%d)", err, in.Page, in.Count)
		return &pb.GetFeedbacksByTypeReply{ErrorCode: util.System_error}, err
	}
	feedbacksResp := make([]*pb.Feedback, 0)
	for _, v := range feedbacks {
		feedbacksResp = append(feedbacksResp, &pb.Feedback{
			Id:          v.Id,
			Description: v.Description,
			Tid :        v.Tid,
			AppInfo:     v.AppInfo,
			UserInfo:    v.UserInfo,
			MobileInfo:  v.MobileInfo,
			ExtendInfo:  v.ExtendInfo,
			Contact:     v.Contact,
			DeviceInfo:  v.DeviceInfo,
			Files:       v.Files,
			Type:        v.Type,
			CreateTime:  v.CreateTime.Unix(),
			UpdateTime:  v.UpdateTime.Unix(),
		})
	}
	return &pb.GetFeedbacksByTypeReply{TotalCount:totalCount,FeedBack:feedbacksResp,ErrorCode: util.Successfull}, nil
}

func (fd *FeedBackServer) GetFeedback(ctx context.Context, in *pb.GetFeedbackRequest) (*pb.GetFeedbackReply, error) {
	log.Infof("Start CetFeedback %#v", in)
	if in.Id == 0 || in.Tid == 0 {
		log.Infof("输入参数异常 id :(%s) tid (%d)", in.Id,in.Tid)
		return &pb.GetFeedbackReply{ErrorCode: util.Input_parameter_error}, nil
	}
	feedback := storage.FeedBack{Id: in.Id,Tid:in.Tid}
	err := feedback.GetFeedback(module.MysqlClient())
	if err != nil {
		log.Infof("获取反馈异常 error :(%s)", err)
		return &pb.GetFeedbackReply{ErrorCode: util.No_Feedback_Can_Be_Find}, nil
	}
	return &pb.GetFeedbackReply{
		ErrorCode: util.Successfull,
		FeedBack: &pb.Feedback{
			Id:          feedback.Id,
			Tid:         feedback.Tid,
			Description: feedback.Description,
			AppInfo:     feedback.AppInfo,
			UserInfo:    feedback.UserInfo,
			MobileInfo:  feedback.MobileInfo,
			ExtendInfo:  feedback.ExtendInfo,
			Contact:     feedback.Contact,
			DeviceInfo:  feedback.DeviceInfo,
			Files:       feedback.Files,
			Type:        feedback.Type,
			CreateTime:  feedback.CreateTime.Unix(),
			UpdateTime:  feedback.UpdateTime.Unix(),
		}}, nil
}

func (fd *FeedBackServer) DelFeedback(ctx context.Context, in *pb.DelFeedbackRequest) (*pb.DelFeedbackReply, error) {
	log.Infof("Start DelFeedback %#v", in)
	if in.Ids == nil || len(in.Ids) == 0 || in.Tid == 0{
		log.Infof("输入参数异常 ids :(%s) tid (%d)", in.Ids,in.Tid)
		return &pb.DelFeedbackReply{ErrorCode: util.Input_parameter_error}, nil
	}
	feedback := storage.FeedBack{Tid:in.Tid}
	err := storage.Transaction(module.MysqlClient(), func(tx *gorm.DB) error {
		err := feedback.DelFeedback(module.MysqlClient(),in.Ids)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Infof("删除反馈异常 error :(%s)", err)
		return &pb.DelFeedbackReply{ErrorCode: util.System_error}, nil
	}
	return &pb.DelFeedbackReply{ErrorCode: util.Successfull}, nil
}

func (fd *FeedBackServer) BatchFeedback(ctx context.Context, in *pb.BatchFeedbackRequest) (*pb.BatchFeedbackReply, error) {
	log.Infof("Start DelFeedback %#v", in)
	if in.Ids == nil || len(in.Ids) == 0 || in.Tid == 0 {
		log.Infof("输入参数异常 ids :(%s) tid (%d)", in.Ids,in.Tid)
		return &pb.BatchFeedbackReply{ErrorCode: util.Input_parameter_error,FeedBack:nil}, nil
	}
	feedback := storage.FeedBack{Tid:in.Tid}
	feedbacks ,err := feedback.BatchFeedback(module.MysqlClient(),in.Ids)
	if err != nil {
		log.Infof("批量获取要删除的工单信息有误:(%s)", err)
		return &pb.BatchFeedbackReply{ErrorCode: util.System_error,FeedBack:nil}, nil
	}
	feedbacksResp := make([]*pb.Feedback, 0)
	for _, v := range feedbacks {
		feedbacksResp = append(feedbacksResp, &pb.Feedback{
			Id:          v.Id,
			Tid:         v.Tid,
			Description: v.Description,
			AppInfo:     v.AppInfo,
			UserInfo:    v.UserInfo,
			MobileInfo:  v.MobileInfo,
			ExtendInfo:  v.ExtendInfo,
			Contact:     v.Contact,
			DeviceInfo:  v.DeviceInfo,
			Files:       v.Files,
			Type:        v.Type,
			CreateTime:  v.CreateTime.Unix(),
			UpdateTime:  v.UpdateTime.Unix(),
		})
	}
	return &pb.BatchFeedbackReply{ErrorCode: util.Successfull,FeedBack:feedbacksResp}, nil
}