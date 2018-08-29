package controller

import  (
	pb "auth-http/api"
	"auth-http/rpc"
	. "auth-http/util"
	"encoding/json"
	log "github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type AliPay struct {
	Did        				int64     	`json:"did,omitempty"`
	AppId 					string    	`json:"appId"`
	MerchantPrivateKey      string     	`json:"merchantPrivateKey"`
	Key						string     	`json:"key"`
}

type WechatPay struct {
	Did        				int64     	`json:"did,omitempty"`
	AppId 					string    	`json:"appId"`
	MchId 					string    	`json:"mchId"`
	AppSecret      			string     	`json:"appSecret"`
	Key						string     	`json:"key"`
}


func AddAliPay(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------AddAliPay----------")

	didInt, err := strconv.Atoi(p.ByName("did"))
	if didInt == 0 || err != nil {
		log.Error("Get did Failed , ", err)
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	did := int32(didInt)
	r.ParseForm()
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var alipay AliPay
	if err := json.Unmarshal(body, &alipay); err != nil {
		log.Errorf("JsonError %s ", err)
		JsonReply("JsonError", nil, w)
		return
	}
	log.Infof("alipay.AppId is %s", alipay.AppId)

	//todo ctx中存入信息
	token := GetCookie(r, TokenCookieKey)
	log.Infof(" get token is (%s)", token)
	if token == "" {
		JsonReply("TokenIsInvalid", nil, w)
		return
	}
	opt := r.Method
	url := strings.Split(r.RequestURI, "?")[0]
	log.Infof("url is (%s),split url is %s", r.RequestURI, url)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "url", url)
	ctx = context.WithValue(ctx, "opt", opt)
	ctx = context.WithValue(ctx, "token", token)
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:", did)
	//todo 调用rpc
	req := &pb.AddAliPayRequest{AliPay:&pb.AliPay{
		AppId:alipay.AppId,
		MerchantPrivateKey:alipay.MerchantPrivateKey,
		Key:alipay.Key,
		Did:did,
	}}
	_,err = rpc.PaymentRpcClient().AddAliPay(ctx,req)
	if err != nil {
		log.Errorf("调用rpc错误码,err is (%s)", err)
		s, ok := status.FromError(err)
		if !ok {
			log.Errorf("系统异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
		code := s.Proto().Code
		for k, v := range CodeMap {
			if code == v.Code {
				log.Info("ErrorInfo:", k)
				JsonReply(k, nil, w)
				return
			}
		}
		JsonReply("SystemError", nil, w)
		return
	}
	JsonReply("Successful", nil, w)
	return
}

func DeleteAliPay(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------DeleteAliPay----------")

	didInt, err := strconv.Atoi(p.ByName("did"))
	if didInt == 0 || err != nil {
		log.Error("Get did Failed , ", err)
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	did := int32(didInt)
	//todo ctx中存入信息
	token := GetCookie(r, TokenCookieKey)
	log.Infof(" get token is (%s)", token)
	if token == "" {
		JsonReply("TokenIsInvalid", nil, w)
		return
	}
	opt := r.Method
	url := strings.Split(r.RequestURI, "?")[0]
	log.Infof("url is (%s),split url is %s", r.RequestURI, url)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "url", url)
	ctx = context.WithValue(ctx, "opt", opt)
	ctx = context.WithValue(ctx, "token", token)
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:", did)
	//todo 调用rpc
	_,err = rpc.PaymentRpcClient().DeleteAliPay(ctx,&pb.DeleteAliPayRequest{Did:did})
	if err != nil {
		log.Errorf("调用rpc错误码,err is (%s)", err)
		s, ok := status.FromError(err)
		if !ok {
			log.Errorf("系统异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
		code := s.Proto().Code
		for k, v := range CodeMap {
			if code == v.Code {
				log.Info("ErrorInfo:", k)
				JsonReply(k, nil, w)
				return
			}
		}
		JsonReply("SystemError", nil, w)
		return
	}
	JsonReply("Successful", nil, w)
	return
}

func UpdateAliPay(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------UpdateAliPay----------")

	didInt, err := strconv.Atoi(p.ByName("did"))
	if didInt == 0 || err != nil {
		log.Error("Get did Failed , ", err)
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	did := int32(didInt)
	r.ParseForm()
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var alipay AliPay
	if err := json.Unmarshal(body, &alipay); err != nil {
		log.Errorf("JsonError %s ", err)
		JsonReply("JsonError", nil, w)
		return
	}
	log.Infof("alipay.AppId is %s", alipay.AppId)

	//todo ctx中存入信息
	token := GetCookie(r, TokenCookieKey)
	log.Infof(" get token is (%s)", token)
	if token == "" {
		JsonReply("TokenIsInvalid", nil, w)
		return
	}
	opt := r.Method
	url := strings.Split(r.RequestURI, "?")[0]
	log.Infof("url is (%s),split url is %s", r.RequestURI, url)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "url", url)
	ctx = context.WithValue(ctx, "opt", opt)
	ctx = context.WithValue(ctx, "token", token)
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:", did)
	//todo 调用rpc
	req := &pb.UpdateAliPayRequest{AliPay:&pb.AliPay{
		AppId:alipay.AppId,
		MerchantPrivateKey:alipay.MerchantPrivateKey,
		Key:alipay.Key,
		Did:did,
	}}
	_,err = rpc.PaymentRpcClient().UpdateAliPay(ctx,req)
	if err != nil {
		log.Errorf("调用rpc错误码,err is (%s)", err)
		s, ok := status.FromError(err)
		if !ok {
			log.Errorf("系统异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
		code := s.Proto().Code
		for k, v := range CodeMap {
			if code == v.Code {
				log.Info("ErrorInfo:", k)
				JsonReply(k, nil, w)
				return
			}
		}
		JsonReply("SystemError", nil, w)
		return
	}
	JsonReply("Successful", nil, w)
	return
}

func GetAliPay(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------GetAliPay----------")
	didInt, err := strconv.Atoi(p.ByName("did"))
	if didInt == 0 || err != nil {
		log.Error("Get did Failed , ", err)
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	did := int32(didInt)
	//todo ctx中存入信息
	token := GetCookie(r, TokenCookieKey)
	log.Infof(" get token is (%s)", token)
	if token == "" {
		JsonReply("TokenIsInvalid", nil, w)
		return
	}
	opt := r.Method
	url := strings.Split(r.RequestURI, "?")[0]
	log.Infof("url is (%s),split url is %s", r.RequestURI, url)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "url", url)
	ctx = context.WithValue(ctx, "opt", opt)
	ctx = context.WithValue(ctx, "token", token)
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:", did)
	//todo 调用rpc
	res,err := rpc.PaymentRpcClient().GetAliPay(ctx,&pb.GetAliPayRequest{Did:did})
	if err != nil {
		log.Errorf("调用rpc错误码,err is (%s)", err)
		s, ok := status.FromError(err)
		if !ok {
			log.Errorf("系统异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
		code := s.Proto().Code
		for k, v := range CodeMap {
			if code == v.Code {
				log.Info("ErrorInfo:", k)
				JsonReply(k, nil, w)
				return
			}
		}
		JsonReply("SystemError", nil, w)
		return
	}
	reply := &AliPay{
		AppId:res.AliPay.AppId,
		MerchantPrivateKey:res.AliPay.MerchantPrivateKey,
		Key:res.AliPay.Key,
	}
	log.Info("reply:",reply)
	JsonReply("Successful", reply, w)
	return
}

func AddWechatPay(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------AddWechatPay----------")

	didInt, err := strconv.Atoi(p.ByName("did"))
	if didInt == 0 || err != nil {
		log.Error("Get did Failed , ", err)
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	did := int32(didInt)
	r.ParseForm()
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var wechatPay WechatPay
	if err := json.Unmarshal(body, &wechatPay); err != nil {
		log.Errorf("JsonError %s ", err)
		JsonReply("JsonError", nil, w)
		return
	}
	log.Infof("wechatPay.AppId is %s", wechatPay.AppId)

	//todo ctx中存入信息
	token := GetCookie(r, TokenCookieKey)
	log.Infof(" get token is (%s)", token)
	if token == "" {
		JsonReply("TokenIsInvalid", nil, w)
		return
	}
	opt := r.Method
	url := strings.Split(r.RequestURI, "?")[0]
	log.Infof("url is (%s),split url is %s", r.RequestURI, url)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "url", url)
	ctx = context.WithValue(ctx, "opt", opt)
	ctx = context.WithValue(ctx, "token", token)
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:", did)
	//todo 调用rpc
	req := &pb.AddWechatPayRequest{WechatPay:&pb.WechatPay{
		Did:did,
		AppId:wechatPay.AppId,
		MchId:wechatPay.MchId,
		Key:wechatPay.Key,
		AppSecret:wechatPay.AppSecret,
	}}
	_,err = rpc.PaymentRpcClient().AddWechatPay(ctx,req)
	if err != nil {
		log.Errorf("调用rpc错误码,err is (%s)", err)
		s, ok := status.FromError(err)
		if !ok {
			log.Errorf("系统异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
		code := s.Proto().Code
		for k, v := range CodeMap {
			if code == v.Code {
				log.Info("ErrorInfo:", k)
				JsonReply(k, nil, w)
				return
			}
		}
		JsonReply("SystemError", nil, w)
		return
	}
	JsonReply("Successful", nil, w)
	return
}

func DeleteWechatPay(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------AddWechatPay----------")
	didInt, err := strconv.Atoi(p.ByName("did"))
	if didInt == 0 || err != nil {
		log.Error("Get did Failed , ", err)
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	did := int32(didInt)
	//todo ctx中存入信息
	token := GetCookie(r, TokenCookieKey)
	log.Infof(" get token is (%s)", token)
	if token == "" {
		JsonReply("TokenIsInvalid", nil, w)
		return
	}
	opt := r.Method
	url := strings.Split(r.RequestURI, "?")[0]
	log.Infof("url is (%s),split url is %s", r.RequestURI, url)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "url", url)
	ctx = context.WithValue(ctx, "opt", opt)
	ctx = context.WithValue(ctx, "token", token)
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:", did)
	//todo 调用rpc
	_,err = rpc.PaymentRpcClient().DeleteWechatPay(ctx,&pb.DeleteWechatPayRequest{Did:did})
	if err != nil {
		log.Errorf("调用rpc错误码,err is (%s)", err)
		s, ok := status.FromError(err)
		if !ok {
			log.Errorf("系统异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
		code := s.Proto().Code
		for k, v := range CodeMap {
			if code == v.Code {
				log.Info("ErrorInfo:", k)
				JsonReply(k, nil, w)
				return
			}
		}
		JsonReply("SystemError", nil, w)
		return
	}
	JsonReply("Successful", nil, w)
	return
}

func UpdateWechatPay(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------UpdateWechatPay----------")
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	r.ParseForm()
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var wechatPay WechatPay
	if err := json.Unmarshal(body, &wechatPay); err != nil {
		log.Errorf("JsonError %s ", err)
		JsonReply("JsonError", nil, w)
		return
	}
	log.Infof("wechatPay.AppId is %s", wechatPay.AppId)
	//todo ctx中存入信息
	token := GetCookie(r, TokenCookieKey)
	log.Infof(" get token is (%s)", token)
	if token == "" {
		JsonReply("TokenIsInvalid", nil, w)
		return
	}
	opt := r.Method
	url := strings.Split(r.RequestURI, "?")[0]
	log.Infof("url is (%s),split url is %s", r.RequestURI, url)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "url", url)
	ctx = context.WithValue(ctx, "opt", opt)
	ctx = context.WithValue(ctx, "token", token)
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:", did)
	//todo 调用rpc
	didInt,_ := strconv.Atoi(did)
	req := &pb.UpdateWechatPayRequest{WechatPay:&pb.WechatPay{
		AppId:wechatPay.AppId,
		MchId:wechatPay.MchId,
		Key:wechatPay.Key,
		AppSecret:wechatPay.AppSecret,
		Did:int32(didInt),
	}}
	_,err := rpc.PaymentRpcClient().UpdateWechatPay(ctx,req)
	if err != nil {
		log.Errorf("调用rpc错误码,err is (%s)", err)
		s, ok := status.FromError(err)
		if !ok {
			log.Errorf("系统异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
		code := s.Proto().Code
		for k, v := range CodeMap {
			if code == v.Code {
				log.Info("ErrorInfo:", k)
				JsonReply(k, nil, w)
				return
			}
		}
		JsonReply("SystemError", nil, w)
		return
	}
	JsonReply("Successful", nil, w)
	return
}

func GetWechatPay(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------GetWechatPay----------")
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	//todo ctx中存入信息
	token := GetCookie(r, TokenCookieKey)
	log.Infof(" get token is (%s)", token)
	if token == "" {
		JsonReply("TokenIsInvalid", nil, w)
		return
	}
	opt := r.Method
	url := strings.Split(r.RequestURI, "?")[0]
	log.Infof("url is (%s),split url is %s", r.RequestURI, url)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "url", url)
	ctx = context.WithValue(ctx, "opt", opt)
	ctx = context.WithValue(ctx, "token", token)
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:", did)
	//todo 调用rpc
	didInt,_ := strconv.Atoi(did)
	res,err := rpc.PaymentRpcClient().GetWechatPay(ctx,&pb.GetWechatPayRequest{Did:int32(didInt)})
	if err != nil {
		log.Errorf("调用rpc错误码,err is (%s)", err)
		s, ok := status.FromError(err)
		if !ok {
			log.Errorf("系统异常,", err)
			JsonReply("SystemError", nil, w)
			return
		}
		code := s.Proto().Code
		for k, v := range CodeMap {
			if code == v.Code {
				log.Info("ErrorInfo:", k)
				JsonReply(k, nil, w)
				return
			}
		}
		JsonReply("SystemError", nil, w)
		return
	}
	reply := WechatPay{
		AppId:res.WechatPay.AppId,
		MchId:res.WechatPay.MchId,
		Key:res.WechatPay.Key,
		AppSecret:res.WechatPay.AppSecret,
	}

	JsonReply("Successful", reply, w)
	return
}