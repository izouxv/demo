package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"cotx-http/result"
	log "github.com/cihub/seelog"
	"cotx-http/pb"
	"cotx-http/utils"
	"cotx-http/rpcClient"
	"context"
)

func AddFeedback(res http.ResponseWriter,req *http.Request,param httprouter.Params)   {
	userinfo ,ok := req.Context().Value("userInfo").(*pb.SsoReply)
	if !ok {
		log.Info("没有找到帐号")
		result.JsonReply("Account_abnormality",nil,res)
		return
	}
   log.Debug(userinfo)
	var feedback = new(pb.Feedback)
	var addFeedbackRequest = new(pb.AddFeedbackRequest)
	errCode := utils.GetHttpData(req, "application/json;charset=UTF-8", &feedback)
	log.Debug(feedback)
	if errCode != 10000 {
		if errCode == 404 {
			result.ResCode(http.StatusNotFound,res)
			return
		}
		result.JsonReply("Body_is_incorrect_or_empty",nil, res)
		return
	}
	addFeedbackRequest.Feedback = feedback

	addFeedbackResponse,err := rpcClient.GetFeedbackClient().AddFeedback(context.Background(),addFeedbackRequest)
	log.Debug(addFeedbackResponse)
	if addFeedbackResponse.ErrCode != 10000 || err != nil {
		log.Error("error:",err)
		switch addFeedbackResponse.ErrCode {
		case 30021:
			result.JsonReply("NOFind_Gateway",nil, res)
			return
		case 30033:
			result.JsonReply("SysTem_ERROR",nil, res)
			return
		case 30027:
			result.JsonReply("Gateway_UnResponse",nil, res)
			return
		default:
			result.JsonReply("Unknown_error",nil, res)
			return
		}
	}
	result.JsonReply("Successful",addFeedbackResponse.Feedback,res)
}
