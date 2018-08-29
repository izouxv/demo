package api

import (
	pb "auth/api"
	"auth/go-drbac/drbac"
	. "auth/util"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
)

type ServiceServer struct {
	DrbacServer *drbac.DrbacServer
}

//添加自定义服务
func (this *ServiceServer) AddService(ctx context.Context, in *pb.AddServiceRequest) (*pb.AddServiceResponse, error) {
	log.Info("Start AddService")
	if in.Service == nil {
		log.Error("in.Service is nil")
		return nil,InvalidArgument
	}
	if in.Service.ServiceName == "" || in.Service.ServiceTid == 0 || in.Service.ServiceKey == "" || in.Service.ServiceType == 0 || in.Service.ServiceUrl == "" {
		log.Error("in.Service is empty,",in.Service)
		return nil,InvalidArgument
	}
	err := this.DrbacServer.AddService(in.Service.Sid,in.Service.ServiceType,in.Service.ServiceName,in.Service.ServiceKey,in.Service.ServiceUrl,in.Service.ServiceTid)
	if err != nil {
		log.Error("AddService Error,",err)
		return nil, SystemError
	}
	return nil,Successful
}

//删除自定义服务
func (this *ServiceServer) DeleteService(ctx context.Context, in *pb.DeleteServiceRequest) (*pb.DeleteServiceResponse, error) {
	log.Info("Start DeleteService")
	if in.Sid == 0 {
		log.Error("in.Sid == 0")
		return nil,InvalidArgument
	}
	err := this.DrbacServer.DeleteService(in.Sid)
	if err != nil {
		log.Error("DeleteService Error,",err)
		return nil,SystemError
	}
	return nil,Successful
}

//修改自定义服务
func (this *ServiceServer) UpdateService(ctx context.Context, in *pb.UpdateServiceRequest) (*pb.UpdateServiceResponse, error) {
	log.Info("Start UpdateService")
	log.Info("in.Service:",in.Service)
	if in.Service == nil {
		log.Error("in.Service is nil")
		return nil,InvalidArgument
	}
	if in.Service.Sid == 0 {
		log.Error("in.Service is empty,",in.Service)
		return nil,InvalidArgument
	}
	err := this.DrbacServer.UpdateService(in.Service.Sid,in.Service.ServiceType,in.Service.ServiceName,in.Service.ServiceKey,in.Service.ServiceUrl,in.Service.ServiceTid, in.Service.ServiceDescription, in.Service.ServiceState)
	if err != nil {
		log.Error("UpdateService Error,",err)
		return nil, SystemError
	}
	return nil,Successful
}

//通过Sid获取服务信息
func (this *ServiceServer) GetServiceBySid(ctx context.Context, in *pb.GetServiceBySidRequest) (*pb.GetServiceBySidResponse, error) {
	log.Info("Start GetServiceBySid")
	if in.Sid == 0 {
		log.Error("in.Sid == 0")
		return nil, InvalidArgument
	}
	service,err := this.DrbacServer.GetServiceBySid(in.Sid)
	if err != nil || service.ServiceName == "" {
		log.Error("GetServiceBySid Error,",err)
		return nil,NotFind
	}
	reply := &pb.Service{
		Sid:service.Sid,
		ServiceName:service.ServiceName,
		ServiceType:service.ServiceType,
		ServiceUrl:service.ServiceUrl,
		ServiceKey:service.ServiceKey,
		ServiceDescription:service.ServiceDescription,

	}
	return &pb.GetServiceBySidResponse{Service:reply},nil
}

func (this *ServiceServer) GetServiceByTid(ctx context.Context, in *pb.GetServiceByTidRequest) (*pb.GetServiceByTidResponse, error) {
	log.Info("Staret GetServiceByTid")
	if in.Tid == 0 {
		log.Error("in.Tid == 0")
		return nil,InvalidArgument
	}
	services,err := this.DrbacServer.GetServiceByTid(in.Tid)
	if err != nil || len(services) == 0 {
		log.Error("GetServiceByTid Error,",err)
		return nil, NotFind
	}
	var reply []*pb.Service
	for _,v := range  services {
		if v.Sid != 0 {
			reply = append(reply, &pb.Service{
				Sid:v.Sid,
				ServiceName:v.ServiceName,
				ServiceType:v.ServiceType,
				ServiceUrl:v.ServiceUrl,
				ServiceKey:v.ServiceKey,
				ServiceDescription:v.ServiceDescription,
				ServiceState:v.ServiceState,
			})
		}
	}
	log.Info("TotalCount:",int32(len(services)))
	return &pb.GetServiceByTidResponse{Service:reply,TotalCount:int32(len(services))},nil
}

