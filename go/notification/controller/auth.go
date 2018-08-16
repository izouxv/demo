package controller

import (
	"context"
	"notification/api"
	"notification/common"
	"notification/rpc"
	log "github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
	"strings"
)

var TokenCookieKey = "token"

type  MailSender struct {
	Id           int64     `json:"id"`
	Did          int64     `json:"did"`
	SmtpServer   string    `json:"smtp_server"`
	EmailSender  string    `json:"sender_email"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
}

func Auth(handle httprouter.Handle) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		log.Info("auth...")
		token := common.GetCookie(req, TokenCookieKey)
		log.Infof("get token is (%s)", token)
		if token == "" {
			common.JsonReply("token_is_invalid", nil, res)
			return
		}
		did, err := common.StrToInt64(params.ByName("domainId"))
		if err != nil {
			common.JsonReply("InvalidArgument", nil, res)
			return
		}
		str := strings.Split(req.URL.String(),"?")
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		authReq := &api.AuthorizationWithDidRequest{Did: did, Url: str[0], Opt: req.Method, Token: token}
		log.Info("AuthorizationRequest: ", authReq)
		ctx = context.WithValue(ctx, "source", "notification")
		//鉴权部分
		//handle(res, req, params)
		//return
		_, err = rpc.AuthClient().AuthorizationWithDid(ctx, authReq)
		log.Info("调用auth-rpc鉴权结果:" ,err)
		authCode, ok := status.FromError(err)
		log.Info("取到的状态码code:", authCode.Code())
		if !ok {
			common.JsonReply("ERR", nil, res)
			return
		}
		switch authCode.Code() {
		case 10000:
			handle(res, req, params)
		case 10001:
			common.JsonReply("ERR", nil, res)
		case 10002:
			common.JsonReply("token_is_invalid", nil, res)
		case 10003:
			common.JsonReply("permission denied", nil, res)
		case 10007:
			common.JsonReply("PARAMS_ERR", nil, res)
		}
		//req.WithContext(context.WithValue(req.Context(), "UserContext", authReply))
		return
	}
}
