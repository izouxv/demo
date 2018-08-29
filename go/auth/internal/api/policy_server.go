package api

import (
	pb "auth/api"
	"auth/go-drbac/drbac"
	. "auth/util"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
)

type PolicyServer struct {
	DrbacServer *drbac.DrbacServer
}

//添加自定义服务
func (this *PolicyServer) AddPolicy(ctx context.Context, in *pb.AddPolicyRequest) (*pb.AddPolicyResponse, error) {
	log.Info("Start AddPolicy")
	if in.Policy == nil {
		log.Error("in.Policy is nil")
		return nil,InvalidArgument
	}
	if in.Policy.PolicyName == "" || in.Policy.PolicySid == 0 || in.Policy.PolicyType == 0 || in.Policy.PolicyCycle == 0 || in.Policy.PolicyFeeType == 0 || in.Policy.PolicyUnitPrice == 0 || in.Policy.PolicyUnitType == 0 || in.Policy.PolicyUnitCount == 0 || in.Policy.PolicySid == 0 {
		log.Error("in.Policy is empty,",in.Policy)
		return nil,InvalidArgument
	}
	err := this.DrbacServer.AddPolicy(in.Policy.PolicyType, in.Policy.PolicyCycle, in.Policy.PolicyFeeType, in.Policy.PolicyUnitType, in.Policy.PolicyUnitCount, in.Policy.PolicySid, in.Policy.PolicyName, in.Policy.PolicyUnitPrice)
	if err != nil {
		log.Error("AddPolicy Error,",err)
		return nil, SystemError
	}
	return nil,Successful
}

//删除自定义服务
func (this *PolicyServer) DeletePolicyByPid(ctx context.Context, in *pb.DeletePolicyByPidRequest) (*pb.DeletePolicyByPidResponse, error) {
	log.Info("Start DeletePolicyByPid")
	if in.Pid == 0 {
		log.Error("in.Pid == 0")
		return nil,InvalidArgument
	}
	err := this.DrbacServer.DeletePolicyByPid(in.Pid)
	if err != nil {
		log.Error("DeletePolicyByPid Error,",err)
		return nil,SystemError
	}
	return nil,Successful
}

//修改自定义服务
func (this *PolicyServer) DeletePolicyBySid(ctx context.Context, in *pb.DeletePolicyBySidRequest) (*pb.DeletePolicyBySidResponse, error) {
	log.Info("Start DeletePolicyBySid")
	if in.PolicySid == 0 {
		log.Error("in.PolicySid is nil")
		return nil,InvalidArgument
	}
	log.Info("in.PolicySid:",in.PolicySid)
	err := this.DrbacServer.DeletePolicyBySid(in.PolicySid)
	if err != nil {
		log.Error("DeletePolicyBySid Error,",err)
		return nil, SystemError
	}
	return nil,Successful
}

//通过Sid获取服务信息
func (this *PolicyServer) UpdatePolicy(ctx context.Context, in *pb.UpdatePolicyRequest) (*pb.UpdatePolicyResponse, error) {
	log.Info("Start UpdatePolicy")
	if in.Policy == nil {
		log.Error("in.Policy is nil")
		return nil,InvalidArgument
	}
	if in.Policy.Pid == 0{
		log.Error("in.Policy is empty,",in.Policy)
		return nil,InvalidArgument
	}
	err := this.DrbacServer.UpdatePolicy(in.Policy.Pid, in.Policy.PolicyType, in.Policy.PolicyCycle, in.Policy.PolicyFeeType, in.Policy.PolicyUnitType, in.Policy.PolicyUnitCount, in.Policy.PolicySid, in.Policy.PolicyName, in.Policy.PolicyUnitPrice)
	if err != nil{
		log.Error("GetServiceBySid Error,",err)
		return nil,SystemError
	}
	return nil,Successful
}

func (this *PolicyServer) GetPolicyByPid(ctx context.Context, in *pb.GetPolicyByPidRequest) (*pb.GetPolicyByPidResponse, error) {
	log.Info("Staret GetServiceByTid")
	if in.Pid == 0 {
		log.Error("in.Pid == 0")
		return nil,InvalidArgument
	}
	policy,err := this.DrbacServer.GetPolicyByPid(in.Pid)
	if err != nil || policy.PolicyName == ""{
		log.Error("GetPolicyByPid Error,",err)
		return nil, NotFind
	}
	reply := &pb.GetPolicyByPidResponse{
		Policy: &pb.Policy{
			Pid:policy.Pid,
			PolicyName:policy.PolicyName,
			PolicyType:policy.PolicyType,
			PolicyCycle:policy.PolicyCycle,
			PolicyFeeType:policy.PolicyFeeType,
			PolicyUnitPrice:policy.PolicyUnitPrice,
			PolicyUnitType:policy.PolicyUnitType,
			PolicyUnitCount:policy.PolicyUnitCount,
			PolicySid:policy.PolicySid,
	}}
	return reply,nil
}

func (this *PolicyServer) GetPolicyBySid(ctx context.Context, in *pb.GetPolicyBySidRequest) (*pb.GetPolicyBySidResponse, error) {
	log.Info("Staret GetPolicyBySid")
	if in.Sid == 0 {
		log.Error("in.Sid == 0")
		return nil,InvalidArgument
	}
	policies,err := this.DrbacServer.GetPolicyBySid(in.Sid)
	if err != nil || len(policies) == 0 {
		log.Error("GetPolicyBySid Error,",err)
		return nil, NotFind
	}
	var reply []*pb.Policy
	for _,v := range  policies {
		if v.Pid != 0 {
			reply = append(reply, &pb.Policy{
				Pid:v.Pid,
				PolicyName:v.PolicyName,
				PolicyType:v.PolicyType,
				PolicyCycle:v.PolicyCycle,
				PolicyFeeType:v.PolicyFeeType,
				PolicyUnitPrice:v.PolicyUnitPrice,
				PolicyUnitType:v.PolicyUnitType,
				PolicyUnitCount:v.PolicyUnitCount,
				PolicySid:v.PolicySid,
			})
		}
	}
	return &pb.GetPolicyBySidResponse{Policy:reply, TotalCount:int32(len(policies))},nil
}