package controller

import (
	pb "auth-http/api"
	"auth-http/rpc"
	. "auth-http/util"
	log "github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"strings"
	"encoding/json"
	"io/ioutil"
)

type Record struct {
	TradingId   		int32  `json:"trading_id"`
	Tid    				int32 `json:"tid"`
	CreateTime    		int64 `json:"create_time"`
	TradingContent    	string `json:"trading_content"`
	TradingUnitPrice 	float32 `json:"trading_unit_price"`
	TradingCount  		int32  `json:"trading_count"`
	TradingState       	int32 `json:"trading_state"`
	TradingTotalPrice   float32  `json:"trading_total_price"`
}

type RecordsResponse struct {
	Record   		[]*Record  `json:"records"`
	TotalCount		int32		`json:"total_count"`
}

type TenantAccount struct {
	Balance float32 `json:"balance"`
}


func GetTenantAccount(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------GetTenantAccount----------")
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
	resp, err := rpc.TradingRpcClient().GetTenantAccount(ctx, &pb.GetTenantAccountRequest{Tid:int32(tidInt)})
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
	log.Info("GetTenantAccount Successful, Balance:",resp.Balance)
	JsonReply("Successful", &TenantAccount{Balance:resp.Balance}, w)
	return
}

func UpdateTenantAccount(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------GetTenantAccount----------")
	tid := p.ByName("tid")
	if tid == "" {
		log.Error("Get tid Failed , ")
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var account TenantAccount
	if err := json.Unmarshal(body, &account); err != nil {
		log.Errorf("JsonError %s ", err)
		JsonReply("JsonError", nil, w)
		return
	}
	log.Infof("Balance is %s", account.Balance)
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
	_, err := rpc.TradingRpcClient().UpdateTenantAccount(ctx, &pb.UpdateTenantAccountRequest{Tid:int32(tidInt),Balance:account.Balance,ActionType:1})
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


func GetTradingRecord(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Info("--------GetTradingRecord----------")
	tid := p.ByName("tid")
	if tid == "" {
		log.Error("Get tid Failed , ")
		JsonReply("TidIsIncorrectOrEmpty", nil, w)
		return
	}
	count := r.FormValue("count")
	page := r.FormValue("page")
	req := &pb.GetTradingRecordsRequest{}
	if page != "" {
		p, err := strconv.Atoi(page)
		if err != nil || p == 0 {
			log.Info("strconv.Atoi(page) Failed,", err)
			JsonReply("PageIsIncorrectOrEmpty", nil, w)
			return
		}
		req.Page = int32(p)
	}
	if count != "" {
		c, err := strconv.Atoi(count)
		if err != nil || c == 0 {
			log.Info("strconv.Atoi(count) Failed,", err)
			JsonReply("PerpageIsIncorrectOrEmpty", nil, w)
			return
		}
		req.Count = int32(c)
	}
	tidInt,_ := strconv.Atoi(tid)
	req.Tid = int32(tidInt)
	log.Info("GetTenantUsers,","count:", count, ", page:", page, ", tid:", tid)
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
	resp, err := rpc.TradingRpcClient().GetTradingRecords(ctx, req)
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
	var records  []*Record
	for _,v := range resp.TradingRecords {
		if v.TradingId != 0 {
			records = append(records, &Record{
				TradingId:v.TradingId,
				Tid:v.Tid,
				CreateTime:v.CreateTime,
				TradingContent:v.TradingContent,
				TradingUnitPrice:v.TradingUnitPrice,
				TradingCount:v.TradingCount,
				TradingState:v.TradingState,
				TradingTotalPrice:v.TradingTotalPrice,
			})
		}
	}
	JsonReply("Successful", &RecordsResponse{Record:records,TotalCount:resp.TotalCount}, w)
	return
}