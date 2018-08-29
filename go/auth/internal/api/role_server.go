package api

import (
	pb "auth/api"
	"auth/go-drbac/drbac"
	. "auth/util"
	log "github.com/cihub/seelog"
	"golang.org/x/net/context"
)

type RoleServer struct {
	DrbacServer *drbac.DrbacServer
}

func (this *RoleServer) GetModuleByTid(ctx context.Context, in *pb.GetModuleByTidRequest) (*pb.GetModuleByTidResponse, error) {
	log.Info("Start AddRole")
	if in.Tid == 0 {
		log.Errorf("GetModuleByTid Input_parameter_error")
		return nil, InvalidArgument
	}
	modules, err := this.DrbacServer.GetModulesByTid(in.Tid)
	if err != nil || len(modules) == 0 {
		log.Errorf("GetModulesByTid Error, ", err)
		return nil, SystemError
	}
	var reply []*pb.ModuleInfo
	for _, v := range modules {
		reply = append(reply, &pb.ModuleInfo{Mid: v.Mid, ModuleName: v.ModuleName})
	}
	return &pb.GetModuleByTidResponse{Modules: reply}, nil
}

func (this *RoleServer) GetModuleByDid(ctx context.Context, in *pb.GetModuleByDidRequest) (*pb.GetModuleByDidResponse, error) {
	log.Info("Start AddRole")
	if in.Did == 0 {
		log.Errorf("GetModuleByTid Input_parameter_error")
		return nil, InvalidArgument
	}
	modules, err := this.DrbacServer.GetModulesByDid(in.Did)
	if err != nil || len(modules) == 0 {
		log.Errorf("GetModulesByTid Error, ", err)
		return nil, SystemError
	}
	var reply []*pb.ModuleInfo
	for _, v := range modules {
		reply = append(reply, &pb.ModuleInfo{Mid: v.Mid, ModuleName: v.ModuleName})
	}
	return &pb.GetModuleByDidResponse{Modules: reply}, nil
}

func (this *RoleServer) AddRole(ctx context.Context, in *pb.AddRoleRequest) (*pb.AddRoleResponse, error) {
	log.Info("Start AddRole")
	if in.RoleName == "" || len(in.Mids) == 0 || in.Tid == 0 {
		log.Errorf("AddRole Input_parameter_error")
		return nil, InvalidArgument
	}
	var modules []*drbac.Modules
	for _, v := range in.Mids {
		modules = append(modules, &drbac.Modules{Mid: v.Mid, Operations: v.Operation})
	}
	if _, err := this.DrbacServer.AddRole(in.RoleName, in.Description, modules, in.Tid); err != nil {
		log.Errorf("AddRole Error, ", err)
		return nil, SystemError
	}
	return nil, Successful
}

func (this *RoleServer) DeleteRole(ctx context.Context, in *pb.DeleteRoleRequest) (*pb.DeleteRoleResponse, error) {
	log.Info("Start DeleteRole")
	if in.Rid == 0 {
		log.Errorf("DeleteRole Input_parameter_error")
		return nil, InvalidArgument
	}
	if err := this.DrbacServer.DeleteTenantRole(in.Rid); err != nil {
		log.Errorf("DeleteRole Error,", err)
		return nil, SystemError
	}
	return nil, Successful
}

func (this *RoleServer) UpdateRole(ctx context.Context, in *pb.UpdateRoleRequest) (*pb.UpdateRoleResponse, error) {
	log.Info("Start UpdateRoleName")
	if in.Rid == 0 || len(in.Mids) == 0 || in.RoleName == "" {
		log.Errorf("UpdateRole Input_parameter_error, input:", in)
		return nil, InvalidArgument
	}
	log.Info("in.Mids:", in.Mids)
	var modules []*drbac.Modules
	for _, v := range in.Mids {
		modules = append(modules, &drbac.Modules{Mid: v.Mid, Operations: v.Operation})
	}
	if err := this.DrbacServer.UpdateRoleInfo(in.Rid, in.RoleName, in.Description, modules); err != nil {
		log.Errorf("UpdateRoleModule Error,", err)
		return nil, SystemError
	}
	return nil, Successful
}

func (this *RoleServer) GetRoles(ctx context.Context, in *pb.GetRolesRequest) (*pb.GetRolesResponse, error) {
	log.Info("Start GetRoles")
	log.Info("GetRoles")
	if in.Page == 0 {
		log.Info("GetUserInfoInDomain Input is Empty")
		return nil,InvalidArgument
	}
	modules, err := this.DrbacServer.GetModulesByTid(in.Tid)
	if err != nil {
		log.Errorf("GetModulesByTid Error,", err)
		return nil, InvalidArgument
	}
	roles,totalCount, err := this.DrbacServer.GetRoles(in.Tid,in.Page,in.Count)
	log.Info("totalCount:",totalCount)
	log.Info("len(roles):",len(roles))
	if err != nil {
		log.Errorf("GetRoles Error,", err)
		return nil, InvalidArgument
	}
	var reply []*pb.RoleModules
	for i := 0; i < len(roles); i++ {
		if roles[i].Rid != 0 {
			rr, err := this.DrbacServer.GetModulesByRid(roles[i].Rid, modules)
			if err != nil {
				log.Errorf("GetMidsByRid Error,", err)
				return nil, SystemError
			}
			var mids []*pb.Module
			for _, v := range rr {
				mids = append(mids, &pb.Module{Mid: v.Mid, ModuleName: v.ModuleName, Operation: v.Operations})
			}
			reply = append(reply, &pb.RoleModules{RoleInfo: &pb.RoleInfo{Rid: roles[i].Rid, RoleName: roles[i].RoleName, Description: roles[i].Description}, Mids: mids})
		}
	}
	log.Info("reply:", reply)
	return &pb.GetRolesResponse{RoleModules: reply,TotalCount:totalCount}, nil
}

func (this *RoleServer) GetRoleByRid(ctx context.Context, in *pb.GetRoleByRidRequest) (*pb.GetRoleByRidResponse, error) {
	log.Info("Start GetRoleByRid")
	if in.Rid == 0 {
		log.Errorf("DeleteRole Input_parameter_error")
		return nil, InvalidArgument
	}
	role, err := this.DrbacServer.GetRoleByRid(in.Rid)
	if err != nil {
		log.Errorf("GetRoles Error,", err)
		return nil, InvalidArgument
	}
	var modules []*drbac.ModuleInfo
	if role.Tid != 0 {
		modules, err = this.DrbacServer.GetModulesByTid(role.Tid)
		if err != nil {
			log.Errorf("GetModulesByTid Error,", err)
			return nil, InvalidArgument
		}
	} else {
		modules, err = this.DrbacServer.GetModulesByDid(role.Did)
		if err != nil {
			log.Errorf("GetModulesByTid Error,", err)
			return nil, InvalidArgument
		}
	}
	rr, err := this.DrbacServer.GetModulesByRid(role.Rid, modules)
	if err != nil {
		log.Errorf("GetMidsByRid Error,", err)
		return nil, SystemError
	}
	var mids []*pb.Module
	for _, v := range rr {
		mids = append(mids, &pb.Module{Mid: v.Mid, ModuleName: v.ModuleName, Operation: v.Operations})
	}
	reply := &pb.RoleModules{RoleInfo: &pb.RoleInfo{Rid: role.Rid, RoleName: role.RoleName, Description: role.Description}, Mids: mids}
	log.Info("reply:", reply)
	return &pb.GetRoleByRidResponse{RoleModules: reply}, nil
}

func (this *RoleServer) AddDomainRole(ctx context.Context, in *pb.AddRoleRequest) (*pb.AddRoleResponse, error) {
	log.Info("Start AddDomainRole")
	if in.RoleName == "" || len(in.Mids) == 0 || in.Did == 0 {
		log.Errorf("AddRole Input_parameter_error")
		return nil, InvalidArgument
	}
	var modules []*drbac.Modules
	for _, v := range in.Mids {
		modules = append(modules, &drbac.Modules{Mid: v.Mid, Operations: v.Operation})
	}
	if _, err := this.DrbacServer.AddDomainRole(in.RoleName, in.Description, modules, in.Did); err != nil {
		log.Errorf("AddRole Error, ", err)
		return nil, SystemError
	}
	return nil, Successful
}

func (this *RoleServer) DeleteDomainRole(ctx context.Context, in *pb.DeleteRoleRequest) (*pb.DeleteRoleResponse, error) {
	log.Info("Start DeleteRole")
	if in.Rid == 0 {
		log.Errorf("DeleteRole Input_parameter_error")
		return nil, InvalidArgument
	}
	if err := this.DrbacServer.DeleteDomainRole(in.Rid); err != nil {
		log.Errorf("DeleteRole Error,", err)
		return nil, SystemError
	}
	return nil, Successful
}

func (this *RoleServer) UpdateDomainRole(ctx context.Context, in *pb.UpdateRoleRequest) (*pb.UpdateRoleResponse, error) {
	log.Info("Start UpdateDomainRole")
	if in.Rid == 0 || len(in.Mids) == 0 || in.RoleName == "" {
		log.Errorf("UpdateRole Input_parameter_error, input:", in)
		return nil, InvalidArgument
	}
	var modules []*drbac.Modules
	for _, v := range in.Mids {
		modules = append(modules, &drbac.Modules{Mid: v.Mid, Operations: v.Operation})
	}
	if err := this.DrbacServer.UpdateRoleInfo(in.Rid, in.RoleName, in.Description, modules); err != nil {
		log.Errorf("UpdateRoleModule Error,", err)
		return nil, SystemError
	}
	return nil, Successful
}

func (this *RoleServer) GetDomainRoles(ctx context.Context, in *pb.GetRolesRequest) (*pb.GetRolesResponse, error) {
	log.Info("Start GetRoles")
	if in.Did == 0 {
		log.Errorf("GetDomainRoles Input_parameter_error, input:", in)
		return nil, InvalidArgument
	}
	if in.Page == 0 {
		log.Info("GetUserInfoInDomain Input is Empty")
		return nil,InvalidArgument
	}
	modules, err := this.DrbacServer.GetModulesByDid(in.Did)
	if err != nil {
		log.Errorf("GetModulesByTid Error,", err)
		return nil, InvalidArgument
	}
	roles, totalCount, err := this.DrbacServer.GetDomainRoles(in.Did,in.Page,in.Count)
	if err != nil {
		log.Errorf("GetRoles Error,", err)
		return nil, InvalidArgument
	}
	var reply []*pb.RoleModules
	for i := 0; i < len(roles); i++ {
		if roles[i].Rid != 0 {
			rr, err := this.DrbacServer.GetModulesByRid(roles[i].Rid, modules)
			if err != nil {
				log.Errorf("GetMidsByRid Error,", err)
				return nil, SystemError
			}
			var mids []*pb.Module
			for _, v := range rr {
				mids = append(mids, &pb.Module{Mid: v.Mid, ModuleName: v.ModuleName, Operation: v.Operations})
			}
			reply = append(reply, &pb.RoleModules{RoleInfo: &pb.RoleInfo{Rid: roles[i].Rid, RoleName: roles[i].RoleName, Description: roles[i].Description}, Mids: mids})
		}
	}
	log.Info("reply:", reply)
	return &pb.GetRolesResponse{RoleModules: reply,TotalCount:totalCount}, nil
}

//todo Resources

func (this *RoleServer) GetResources(ctx context.Context, in *pb.GetResourcesRequest) (*pb.GetResourcesResponse, error) {
	log.Info("Start GetResources")
	if in.Page == 0 {
		log.Info("GetResources Input is Empty")
		return nil,InvalidArgument
	}
	resources,totalCount,err := this.DrbacServer.GetResources(in.Page, in.Count,in.ResName,in.ResOpt)
	log.Info("totalCount:",totalCount)
	log.Info("len(resources):",len(resources))
	if err != nil {
		log.Errorf("GetResources Error,", err)
		return nil, InvalidArgument
	}
	var reply []*pb.Resource
	for _,v := range resources {
		if v.ResId != 0 {
			reply = append(reply, &pb.Resource{ResId:v.ResId,ResName:v.ResName,ResOpt:v.ResOpt,ResUrl:v.ResUrl})
		}
	}
	log.Info("reply:", reply)
	return &pb.GetResourcesResponse{Resources:reply,TotalCount:totalCount}, nil
}

func (this *RoleServer) CreateResource(ctx context.Context, in *pb.CreateResourceRequest) (*pb.CreateResourceResponse, error) {
	log.Info("Start CreateResource")
	if in.Resources == nil {
		log.Info("in.Resources is nil")
		return nil,InvalidArgument
	}
	log.Info("in.Resources:",in.Resources)
	if in.Resources.ResName == "" || in.Resources.ResOpt == "" || in.Resources.ResUrl == ""{
		log.Info("in.Resources input is empty")
		return nil,InvalidArgument
	}
	err := this.DrbacServer.CreateResource(in.Resources.ResName,in.Resources.ResUrl,in.Resources.ResOpt)
	if err != nil {
		log.Error("CreateResource Error,",err)
		return nil,SystemError
	}
	return nil,Successful
}

func (this *RoleServer) UpdateResource(ctx context.Context, in *pb.UpdateResourceRequest) (*pb.UpdateResourceResponse, error) {
	log.Info("Start UpdateResource")
	if in.Resources == nil {
		log.Info("in.Resources is nil")
		return nil,InvalidArgument
	}
	log.Info("in.Resources:",in.Resources)
	err := this.DrbacServer.UpdateResource(in.Resources.ResId,in.Resources.ResName,in.Resources.ResUrl,in.Resources.ResOpt)
	if err != nil {
		log.Error("CreateResource Error,",err)
		return nil,SystemError
	}
	return nil,Successful
}

func (this *RoleServer) DeleteResource(ctx context.Context, in *pb.DeleteResourceRequest) (*pb.DeleteResourceResponse, error) {
	log.Info("Start UpdateResource")
	log.Info("in.ResId:",in.ResId)
	if in.ResId == 0 {
		log.Info("in.Resources is nil")
		return nil,InvalidArgument
	}
	err := this.DrbacServer.DeleteResource(in.ResId)
	if err != nil {
		log.Error("CreateResource Error,",err)
		return nil,SystemError
	}
	return nil,Successful
}
