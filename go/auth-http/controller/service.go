package controller

import (
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

type Service struct {
	Sid            int32  	`json:"sid,omitempty"`
	ServiceName    string  	`json:"serviceName,omitempty"`
	ServiceUrl     string 	`json:"serviceUrl,omitempty"`
	ServiceKey     string	`json:"serviceKey,omitempty"`
	ServiceType	   int32	`json:"serviceType,omitempty"`
	ServiceDescription string `json:"serviceDescription,omitempty"`
	ServiceState   int32     `json:"serviceState,omitempty"`
}

type GetServiceByTidReply struct {
	Services	[]*Service		`json:"services,omitempty"`
	TotalCount	int32			`json:"total_count,omitempty"`
}

func AddService(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------AddService----------")
	r.ParseForm()
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var service Service
	if err := json.Unmarshal(body, &service); err != nil {
	log.Errorf("JsonError %s ", err)
	JsonReply("JsonError", nil, w)
	return
	}
	log.Infof("ServiceName is %s", service.ServiceName)
	tid := p.ByName("tid")
	if tid == "" {
		log.Error("Get tid Failed , ")
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
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
	ctx = context.WithValue(ctx, "tid", tid)
	log.Info("url:", ctx.Value("url"))
	log.Info("tid:", tid)
	//todo 调用rpc
	tidInt,_ := strconv.Atoi(tid)
	req := &pb.AddServiceRequest{Service:&pb.Service{
		ServiceName:service.ServiceName,
		ServiceUrl:service.ServiceUrl,
		ServiceKey:service.ServiceKey,
		ServiceType:service.ServiceType,
		ServiceTid:int32(tidInt),
	}}
	_,err := rpc.ServiceRpcClient().AddService(ctx,req)
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
func UpdateService(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------UpdateService----------")
	r.ParseForm()
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var service Service
	if err := json.Unmarshal(body, &service); err != nil {
		log.Errorf("JsonError %s ", err)
		JsonReply("JsonError", nil, w)
		return
	}
	log.Infof("ServiceName is %s", service.ServiceName)
	tid := p.ByName("tid")
	if tid == "" {
		log.Error("Get tid Failed , ")
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	sid64, err := strconv.ParseInt(p.ByName("sid"), 10, 64)
	if sid64 == 0 || err != nil {
		log.Error("Get tid Failed , ", err)
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	sid := int32(sid64)
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
	ctx = context.WithValue(ctx, "tid", tid)
	log.Info("url:", ctx.Value("url"))
	log.Info("tid:", tid)
	//todo 调用rpc
	tidInt,_ := strconv.Atoi(tid)
	req := &pb.UpdateServiceRequest{Service:&pb.Service{
		Sid:sid,
		ServiceName:service.ServiceName,
		ServiceUrl:service.ServiceUrl,
		ServiceKey:service.ServiceKey,
		ServiceType:service.ServiceType,
		ServiceTid:int32(tidInt),
		ServiceDescription:service.ServiceDescription,
		ServiceState:service.ServiceState,
	}}
	_,err = rpc.ServiceRpcClient().UpdateService(ctx,req)
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
func DeleteService(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------DeleteService----------")
	tid := p.ByName("tid")
	if tid == "" {
		log.Error("Get tid Failed , ")
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	sid64, err := strconv.ParseInt(p.ByName("sid"), 10, 64)
	if sid64 == 0 || err != nil {
		log.Error("Get tid Failed , ", err)
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	sid := int32(sid64)

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
	ctx = context.WithValue(ctx, "tid", tid)
	log.Info("url:", ctx.Value("url"))
	log.Info("tid:", tid)
	//todo 调用rpc
	req := &pb.DeleteServiceRequest{Sid:sid}
	_,err = rpc.ServiceRpcClient().DeleteService(ctx,req)
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
func GetServiceBySid(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------GetServiceBySid----------")
	tid := p.ByName("tid")
	if tid == "" {
		log.Error("Get tid Failed , ")
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	sid64, err := strconv.ParseInt(p.ByName("sid"), 10, 64)
	if sid64 == 0 || err != nil {
		log.Error("Get sid Failed , ", err)
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	sid := int32(sid64)
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
	ctx = context.WithValue(ctx, "tid", tid)
	log.Info("url:", ctx.Value("url"))
	log.Info("tid:", tid)
	//todo 调用rpc
	req := &pb.GetServiceBySidRequest{Sid:sid}
	reply,err := rpc.ServiceRpcClient().GetServiceBySid(ctx,req)
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
	service := &Service{
		Sid:reply.Service.Sid,
		ServiceName:reply.Service.ServiceName,
		ServiceType:reply.Service.ServiceType,
		ServiceUrl:reply.Service.ServiceUrl,
		ServiceKey:reply.Service.ServiceKey,
		ServiceDescription:reply.Service.ServiceDescription,
		ServiceState:reply.Service.ServiceState,
	}
	JsonReply("Successful", service, w)
	return
}
func GetServiceByTid(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------GetServiceByTid----------")
	tid := p.ByName("tid")
	if tid == "" {
		log.Error("Get tid Failed , ")
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
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
	ctx = context.WithValue(ctx, "tid", tid)
	log.Info("url:", ctx.Value("url"))
	log.Info("tid:", tid)
	//todo 调用rpc
	tidInt,_ := strconv.Atoi(tid)
	req := &pb.GetServiceByTidRequest{Tid:int32(tidInt)}
	reply,err := rpc.ServiceRpcClient().GetServiceByTid(ctx,req)
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
	var services []*Service
	for _,v := range reply.Service {
		service := &Service{
			Sid:v.Sid,
			ServiceName:v.ServiceName,
			ServiceUrl:v.ServiceUrl,
			ServiceKey:v.ServiceKey,
			ServiceType:v.ServiceType,
			ServiceDescription:v.ServiceDescription,
			ServiceState:v.ServiceState,
		}
		services = append(services, service)
	}

	JsonReply("Successful", &GetServiceByTidReply{Services:services,TotalCount:reply.TotalCount}, w)
	return
}