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

type Policy struct {
	Pid        		int32     `json:"pid"`
	PolicyName 		string    `json:"policyName"`
	PolicyType      int32     `json:"policyType"`
	PolicyCycle		int32     `json:"policyCycle"`
	PolicyFeeType   int32     `json:"policyFeeType"`
	PolicyUnitPrice	float32   `json:"policyUnitPrice"`
	PolicyUnitType  int32     `json:"policyUnitType"`
	PolicyUnitCount	int32     `json:"policyUnitCount"`
	PolicySid	    int32     `json:"policySid"`
}

type GetPolicyBySidReply struct {
	Policies	[]*Policy		`json:"policies,omitempty"`
	TotalCount	int32			`json:"total_count,omitempty"`
}

func AddPolicy(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------AddPolicy----------")

	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	sid64, err := strconv.ParseInt(p.ByName("sid"), 10, 64)
	if sid64 == 0 || err != nil {
		log.Error("Get tid Failed , ", err)
		JsonReply("InvalidArgument", nil, w)
		return
	}
	sid := int32(sid64)
	r.ParseForm()
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var policy Policy
	if err := json.Unmarshal(body, &policy); err != nil {
		log.Errorf("JsonError %s ", err)
		JsonReply("JsonError", nil, w)
		return
	}
	log.Infof("PolicyName is %s", policy.PolicyName)

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
	req := &pb.AddPolicyRequest{Policy:&pb.Policy{
		PolicyName: policy.PolicyName,
		PolicySid:sid,
		PolicyType:policy.PolicyType,
		PolicyCycle:policy.PolicyCycle,
		PolicyFeeType:policy.PolicyFeeType,
		PolicyUnitPrice:policy.PolicyUnitPrice,
		PolicyUnitType:policy.PolicyUnitType,
		PolicyUnitCount:policy.PolicyUnitCount,
	}}
	_,err = rpc.PolicyRpcClient().AddPolicy(ctx,req)
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

func DeletePolicyByPid(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------DeletePolicyByPid----------")
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	pid64, err := strconv.ParseInt(p.ByName("pid"), 10, 64)
	if pid64 == 0 || err != nil {
		log.Error("Get pid Failed , ", err)
		JsonReply("InvalidArgument", nil, w)
		return
	}
	pid := int32(pid64)
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

	_,err = rpc.PolicyRpcClient().DeletePolicyByPid(ctx,&pb.DeletePolicyByPidRequest{Pid:pid})
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
func DeletePolicyBySid(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------DeletePolicyBySid----------")
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
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
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:", did)
	//todo 调用rpc
	_,err = rpc.PolicyRpcClient().DeletePolicyBySid(ctx,&pb.DeletePolicyBySidRequest{PolicySid:sid})
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
func UpdatePolicy(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------UpdatePolicy----------")
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	pid64, err := strconv.ParseInt(p.ByName("pid"), 10, 64)
	if pid64 == 0 || err != nil {
		log.Error("Get pid Failed , ", err)
		JsonReply("InvalidArgument", nil, w)
		return
	}
	pid := int32(pid64)
	r.ParseForm()
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var policy Policy
	if err := json.Unmarshal(body, &policy); err != nil {
		log.Errorf("JsonError %s ", err)
		JsonReply("JsonError", nil, w)
		return
	}
	log.Infof("PolicyName is %s", policy.PolicyName)
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
	req := &pb.UpdatePolicyRequest{Policy:&pb.Policy{
		Pid:pid,
		PolicyName: policy.PolicyName,
		PolicyType:policy.PolicyType,
		PolicyCycle:policy.PolicyCycle,
		PolicyFeeType:policy.PolicyFeeType,
		PolicyUnitPrice:policy.PolicyUnitPrice,
		PolicyUnitType:policy.PolicyUnitType,
		PolicyUnitCount:policy.PolicyUnitCount,
	}}
	_,err = rpc.PolicyRpcClient().UpdatePolicy(ctx,req)
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
func GetPolicyByPid(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------GetPolicyByPid----------")
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	pid64, err := strconv.ParseInt(p.ByName("pid"), 10, 64)
	if pid64 == 0 || err != nil {
		log.Error("Get pid Failed , ", err)
		JsonReply("InvalidArgument", nil, w)
		return
	}
	pid := int32(pid64)
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
	req := &pb.GetPolicyByPidRequest{Pid:pid}
	reply,err := rpc.PolicyRpcClient().GetPolicyByPid(ctx,req)
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
	policy := &Policy{
		Pid:pid,
		PolicyName: reply.Policy.PolicyName,
		PolicyType:reply.Policy.PolicyType,
		PolicyCycle:reply.Policy.PolicyCycle,
		PolicyFeeType:reply.Policy.PolicyFeeType,
		PolicyUnitPrice:reply.Policy.PolicyUnitPrice,
		PolicyUnitType:reply.Policy.PolicyUnitType,
		PolicyUnitCount:reply.Policy.PolicyUnitCount,
		PolicySid:reply.Policy.PolicySid,
	}
	JsonReply("Successful", policy, w)
	return
}
func GetPolicyBySid(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------GetPolicyBySid----------")
	did := p.ByName("did")
	if did == "" {
		log.Error("Get did Failed , ")
		JsonReply("DidIsIncorrectOrEmpty", nil, w)
		return
	}
	sid64, err := strconv.ParseInt(p.ByName("sid"), 10, 64)
	if sid64 == 0 || err != nil {
		log.Error("Get pid Failed , ", err)
		JsonReply("InvalidArgument", nil, w)
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
	ctx = context.WithValue(ctx, "did", did)
	log.Info("url:", ctx.Value("url"))
	log.Info("did:", did)
	//todo 调用rpc
	req := &pb.GetPolicyBySidRequest{Sid:sid}
	reply,err := rpc.PolicyRpcClient().GetPolicyBySid(ctx,req)
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
	var policies []*Policy
	for _,v := range reply.Policy {
		policy := &Policy{
			Pid:v.Pid,
			PolicyName: v.PolicyName,
			PolicyType:v.PolicyType,
			PolicyCycle:v.PolicyCycle,
			PolicyFeeType:v.PolicyFeeType,
			PolicyUnitPrice:v.PolicyUnitPrice,
			PolicyUnitType:v.PolicyUnitType,
			PolicyUnitCount:v.PolicyUnitCount,
			PolicySid:v.PolicySid,
		}
		policies = append(policies, policy)
	}
	JsonReply("Successful", &GetPolicyBySidReply{Policies:policies, TotalCount:reply.TotalCount}, w)
	return
}
