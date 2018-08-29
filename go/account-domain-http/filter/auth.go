package filter

import (
	"context"
	log "github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
	"account-domain-http/util"
	"account-domain-http/rpc"
	pb "account-domain-http/api/auth/api"
	"strings"
)

var tokenCookieKey = "token"

type AccountContext struct {
	Tid      int64
	Username string
	Did      int64
	Token    string
	NickName string
}

func Auth(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		log.Info("auth...")
		token := getCookie(req, tokenCookieKey)
		log.Infof("get token is (%s)", token)
		if token == "" {
			util.JsonReply("Session_is_invalid", nil, w)
			return
		}
		tid, err := util.StrToInt64(params.ByName("tid"))
		if err != nil {
			log.Infof("get tid (%d) 异常  err (%s) ",err)
			util.JsonReply("TenantId_is_incorrect_or_empty", nil, w)
			return
		}
		log.Infof("get tid is (%d)", tid)
		str := strings.Split(req.URL.String(),"?")
		handle(w, req, params)
		return
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		authRequest := &pb.AuthorizationRequest{Tid: tid,Url:str[0], Opt: req.Method, Token: token}
		ctx = context.WithValue(ctx, "source", "account-domain-http")
       //鉴权部分
		log.Infof("AuthRequest %#v ", authRequest)
		_, err = rpc.AuthRpcClient().Authorization(ctx, authRequest)
		log.Error("调用鉴权错误:",err)
		if err != nil {
		authCode, ok := status.FromError(err)
		log.Info("鉴权返回的状态码:",ok)
		log.Info("鉴权返回的状态码:",int32(authCode.Code()))
		if !ok {
			util.JsonReply("Session_is_invalid", nil, w)
			return
	    }
	    log.Infof("code:%d",int32(authCode.Code()))  //code 为14表示连接rpc失败
		switch int32(authCode.Code()){
		case 10000:
			handle(w, req, params)
		case 10001:
			util.JsonReply("System_error", nil, w)
			break
		case 10002:
			util.JsonReply("Session_is_invalid", nil, w)
			break
		case 10003:
			util.JsonReply("Permission_denied", nil, w)
			break
		case 10007:
			util.JsonReply("Params_error", nil, w)
			break
		default :
			log.Error("auth返回状态码:",int32(authCode.Code()))
			util.JsonReply("System_error", nil, w)
			break
		}
		return
		}
	}
}

/*----------GetCookie--------------*/

func getCookie(r *http.Request, key string) string {
	cookie, err := r.Cookie(key)
	if err != nil {
		return ""
	}
	return cookie.Value
}
