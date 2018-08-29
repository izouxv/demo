package api

import (
	pb "auth/api"
	"auth/go-drbac/drbac"
	. "auth/util"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"github.com/pkg/errors"
)

type TradingServer struct {
	DrbacServer *drbac.DrbacServer
}


func (this *TradingServer) GetTenantAccount(ctx context.Context, in *pb.GetTenantAccountRequest) (*pb.GetTenantAccountResponse, error) {
	log.Info("--GetTenantAccount--")
	log.Info("in.Tid:",in.Tid)
	if in.Tid == 0 {
		log.Error("in.Tid == 0")
		return nil,InvalidArgument
	}
	account,err := this.DrbacServer.GetTenantAccountByTid(in.Tid)
	if err != nil || account.Tid == 0 {
		log.Error("GetTenantAccountByTid Error,",err)
		return nil, NotFind
	}
	log.Info("account.Balance:",account.Balance)
	return &pb.GetTenantAccountResponse{Balance:account.Balance},nil
}


func (this *TradingServer) UpdateTenantAccount(ctx context.Context, in *pb.UpdateTenantAccountRequest) (*pb.UpdateTenantAccountResponse, error) {
	log.Info("--UpdateTenantAccount--")
	log.Info("in.Tid:",in.Tid,", in.Balance:",in.Balance,", in.ActionType:",in.ActionType)
	if in.Tid == 0 {
		log.Error("in.Tid == 0")
		return nil,InvalidArgument
	}
	if in.Balance == 0 || in.ActionType == 0 {
		log.Error("in.Tid == 0")
		return nil,InvalidArgument
	}
	//todo 调用修改接口
	err := this.DrbacServer.UpdateTenantAccountByTid(in.Tid,in.Balance,in.ActionType)
	if err != nil {
		log.Info("UpdateTenantAccountByTid Error,",err)
		if err == errors.New("账户余额不足，扣费失败") {
			return nil,InsufficientAccountBalance
		}
		return nil,SystemError
	}
	return nil,Successful
}


func (this *TradingServer) GetTradingRecords(ctx context.Context, in *pb.GetTradingRecordsRequest) (*pb.GetTradingRecordsResponse, error) {
	log.Info("--GetTradingRecords--")
	log.Info("in.Tid:",in.Tid)
	if in.Tid == 0 {
		log.Error("in.Tid == 0")
		return nil,InvalidArgument
	}
	log.Info("GetTradingRecords Input Info, page:", in.Page, ", count:", in.Count)
	if in.Page == 0 {
		log.Info("GetUserInfoInTenant Input is Empty")
		return nil,InvalidArgument
	}
	records,totalCount,err := this.DrbacServer.GetTradingRecordsByTid(in.Tid,in.Page,in.Count)
	if err != nil || len(records) == 0 || totalCount == 0{
		log.Error("GetTradingRecordsByTid Error,",err)
		return nil, NotFind
	}
	var reply []*pb.TradingRecord
	for _,v := range  records {
		if v.TradingId != 0 {
			reply = append(reply, &pb.TradingRecord{
				TradingId:v.TradingId,
				Tid:v.Tid,
				CreateTime:v.CreateTime.Unix(),
				TradingContent:v.TradingContent,
				TradingUnitPrice:v.TradingUnitPrice,
				TradingTotalPrice:v.TradingTotalPrice,
				TradingCount:v.TradingCount,
				TradingState:v.TradingState,
			})
		}
	}
	log.Info("TotalCount:",totalCount)
	return &pb.GetTradingRecordsResponse{TradingRecords:reply, TotalCount:totalCount},nil
}
