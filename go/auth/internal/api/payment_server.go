package api

import (
	pb "auth/api"
	"auth/go-drbac/drbac"
	. "auth/util"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
)

type PaymentServer struct {
	DrbacServer *drbac.DrbacServer
}

//添加租户支付宝接口参数
func (this *PaymentServer) AddAliPay(ctx context.Context, in *pb.AddAliPayRequest) (*pb.AddAliPayResponse, error) {
	log.Info("Start AddAilPay")
	if in.AliPay == nil {
		log.Error("AddAilPay input is empty")
		return nil,InvalidArgument
	}
	if in.AliPay.Did == 0 || in.AliPay.AppId == "" || in.AliPay.MerchantPrivateKey == "" || in.AliPay.Key == "" {
		log.Error("AddAilPay input is empty,in.AliPay:",in.AliPay)
		return nil,InvalidArgument
	}
	err := this.DrbacServer.AddAliPay(in.AliPay.Did,in.AliPay.MerchantPrivateKey,in.AliPay.Key,in.AliPay.AppId)
	if err != nil {
		log.Error("AddAliPay Error,",err)
		return nil,SystemError
	}
	return nil,Successful
}

//删除租户支付宝接口参数
func (this *PaymentServer) DeleteAliPay(ctx context.Context, in *pb.DeleteAliPayRequest) (*pb.DeleteAliPayResponse, error) {
	log.Info("Start DeleteAliPay")
	if in.Did == 0 {
		log.Error("DeleteAliPay input is empty")
		return nil,InvalidArgument
	}
	err := this.DrbacServer.DeleteAliPay(in.Did)
	if err != nil {
		log.Error("DeleteAliPay Error,",err)
		return nil,SystemError
	}
	return nil,Successful
}

//修改租户支付宝接口参数
func (this *PaymentServer) UpdateAliPay(ctx context.Context, in *pb.UpdateAliPayRequest) (*pb.UpdateAliPayResponse, error) {
	log.Info("Start UpdateAliPay")
	if in.AliPay == nil {
		log.Error("UpdateAliPay input is empty")
		return nil,InvalidArgument
	}
	if in.AliPay.Did == 0 || in.AliPay.AppId == "" || in.AliPay.MerchantPrivateKey == "" || in.AliPay.Key == "" {
		log.Error("UpdateAliPay input is empty,in.AliPay:",in.AliPay)
		return nil,InvalidArgument
	}
	err := this.DrbacServer.UpdateAliPay(in.AliPay.Did,in.AliPay.MerchantPrivateKey,in.AliPay.Key,in.AliPay.AppId)
	if err != nil {
		log.Error("UpdateAliPay Error,",err)
		return nil,SystemError
	}
	return nil,Successful
}

//通过Sid获取服务信息
func (this *PaymentServer) GetAliPay(ctx context.Context, in *pb.GetAliPayRequest) (*pb.GetAliPayResponse, error) {
	log.Info("Start GetAliPay")
	if in.Did == 0 {
		log.Error("DeleteAliPay input is empty")
		return nil,InvalidArgument
	}
	res,err := this.DrbacServer.GetAliPay(in.Did)
	if err != nil || res.AppId == ""{
		log.Error("GetAliPay Error,",err)
		return nil,NotFind
	}
	aliPay := &pb.AliPay{
		AppId:res.AppId,
		MerchantPrivateKey:res.MerchantPrivateKey,
		Key:res.Key,
	}
	log.Info("aliPay:",aliPay)
	return &pb.GetAliPayResponse{AliPay:aliPay},nil
}

func (this *PaymentServer) AddWechatPay(ctx context.Context, in *pb.AddWechatPayRequest) (*pb.AddWechatPayResponse, error) {
	log.Info("Start AddWechatPay")
	if in.WechatPay == nil {
		log.Error("AddWechatPay input is empty")
		return nil,InvalidArgument
	}
	if in.WechatPay.Did == 0 || in.WechatPay.AppId == "" || in.WechatPay.MchId == "" || in.WechatPay.Key == "" || in.WechatPay.AppSecret == "" {
		log.Error("AddWechatPay input is empty,in.WechatPay:",in.WechatPay)
		return nil,InvalidArgument
	}
	err := this.DrbacServer.AddWechatPay(in.WechatPay.Did,in.WechatPay.Key,in.WechatPay.AppId,in.WechatPay.MchId,in.WechatPay.AppSecret)
	if err != nil {
		log.Error("AddWechatPay Error,",err)
		return nil,SystemError
	}
	return nil,Successful
}

func (this *PaymentServer) DeleteWechatPay(ctx context.Context, in *pb.DeleteWechatPayRequest) (*pb.DeleteWechatPayResponse, error) {
	log.Info("Start DeleteWechatPay")
	if in.Did == 0 {
		log.Error("DeleteWechatPay input is empty")
		return nil,InvalidArgument
	}
	err := this.DrbacServer.DeleteWechatPay(in.Did)
	if err != nil {
		log.Error("DeleteWechatPay Error,",err)
		return nil,SystemError
	}
	return nil,Successful
}

func (this *PaymentServer) UpdateWechatPay(ctx context.Context, in *pb.UpdateWechatPayRequest) (*pb.UpdateWechatPayResponse, error) {
	log.Info("Start UpdateWechatPay")
	if in.WechatPay == nil {
		log.Error("UpdateWechatPay input is empty")
		return nil,InvalidArgument
	}
	if in.WechatPay.Did == 0 || in.WechatPay.AppId == "" || in.WechatPay.MchId == "" || in.WechatPay.Key == "" || in.WechatPay.AppSecret == "" {
		log.Error("UpdateWechatPay input is empty,in.WechatPay:",in.WechatPay)
		return nil,InvalidArgument
	}
	err := this.DrbacServer.UpdateWechatPay(in.WechatPay.Did,in.WechatPay.Key,in.WechatPay.AppId,in.WechatPay.MchId,in.WechatPay.AppSecret)
	if err != nil {
		log.Error("AddWechatPay Error,",err)
		return nil,SystemError
	}
	return nil,Successful
}

func (this *PaymentServer) GetWechatPay(ctx context.Context, in *pb.GetWechatPayRequest) (*pb.GetWechatPayResponse, error) {
	log.Info("Start GetAliPay")
	if in.Did == 0 {
		log.Error("GetWechatPay input is empty")
		return nil,InvalidArgument
	}
	res,err := this.DrbacServer.GetWechatPay(in.Did)
	if err != nil || res.AppId == ""{
		log.Error("GetAliPay Error,",err)
		return nil,NotFind
	}
	wechatPay := &pb.WechatPay{
		AppId:res.AppId,
		MchId:res.MchId,
		AppSecret:res.AppSecret,
		Key:res.Key,
	}
	return &pb.GetWechatPayResponse{WechatPay:wechatPay},nil
}