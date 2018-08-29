package drbac

import (
	. "auth/go-drbac/internal/storage"

	. "auth/go-drbac/common"

	log "github.com/cihub/seelog"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type DomainRoleResource struct {
	Domain             *Domain `json:"Domain"`
	IsDefaultDomain    bool    `json:"IsDefaultDomain"`
	RoleResourceInfo  			`json:"RoleResourceInfo"`
}

type TenantRoleResource struct {
	Tenant             *Tenant `json:"Tenant"`
	IsDefaultTenant    bool    `json:"IsDefaultTenant"`
	RoleResourceInfo 			`json:"RoleResourceInfo"`
	Children           []*TenantTree
}

type TenantTree struct {
	Tenant   *Tenant
	Children []*TenantTree
}

type RoleResourceInfo struct {
	Role        []*Role       `json:"Role"`
	Resource    []*Resource `json:"Resource"`
}

/*通过uid和did获取域的角色权限*/
func getDomainRolePermissionByUidAndDid(uid, did int32, tx *gorm.DB) (*DomainRoleResource, error) {
	//todo 获取租户信息
	domain := Domain{Did: did}
	err := domain.GetByID(tx)
	if err != nil {
		log.Errorf("domain.GetByID Error, %s,  uid is %d", err, uid)
		return nil, err
	}
	drp := &DomainRoleResource{}
	drp.Domain = &domain
	drp.IsDefaultDomain = true
	udr := UserDomainRole{Uid: uid, Did: did}
	//todo 获取用户所有的角色
	udrs, err := udr.GetUserRoleByUidAndDid(tx)
	if err != nil{
		log.Errorf("GetRolesByUidAndDid Error, err is %s, uid is %d", err, uid)
		return nil, errors.New("获取用户角色失败")
	}
	if len(udrs) == 0 {
		log.Info("用户无角色")
		return drp,nil
	}
	var rids []int32
	var roles []*Role
	for _, udr := range udrs {
		log.Info("udr.Rid:",udr.Rid)
		rids = append(rids, udr.Rid)
		//todo 获取角色信息
		role := &Role{Rid: udr.Rid}
		err = role.GetRoleByRid(tx)
		if err != nil {
			continue
		}
		roles = append(roles, role)
	}
	log.Info("rids:",rids)
	drp.Role = roles
	//todo 获取所有权限pids
	perIds,err := GetPermissionByRids(rids,tx)
	if err != nil {
		log.Error("GetPermissionByRids Error, ",err)
		return drp,nil
	}
	//todo 获取所有的resources
	resource,err := GetRolePermissionInfoByRid(tx,perIds)
	if err != nil {
		log.Error("GetRolePermissionInfoByRid Error,", err)
	}
	log.Info("roleMids:", rids)
	log.Info("resource:", resource)
	drp.Resource = resource
	log.Info("drp:", drp)
	return drp, nil
}

/*获取用户角色权限*/
func getTenantRoleResourceByUidAndTid(uid, tid int32, tx *gorm.DB) ([]*TenantRoleResource, error) {
	//todo 获取租户信息
	tenant := Tenant{Tid: tid}
	err := tenant.GetByID(tx)
	if err != nil {
		log.Errorf("domain.GetByID Error, %s,  uid is %d", err, uid)
		return nil, err
	}
	var drps []*TenantRoleResource
	drp := TenantRoleResource{}
	drp.Tenant = &tenant
	drp.IsDefaultTenant = true
	udr := UserTenantRole{Uid: uid, Tid: tid}
	//todo 获取用户所有的角色
	udrs, err := udr.GetUserRoleByUidAndTid(tx)
	if err != nil{
		log.Errorf("GetRolesByUidAndTid Error, err is %s, uid is %d", err, uid)
		return nil, errors.New("获取用户角色失败")
	}
	if len(udrs) == 0 {
		log.Info("用户无角色")
		drps = append(drps, &drp)
		return drps,nil
	}
	var rids []int32
	var roles []*Role
	for _, udr := range udrs {
		log.Info("udr.Rid:",udr.Rid)
		rids = append(rids, udr.Rid)
		//todo 获取角色信息
		role := &Role{Rid: udr.Rid}
		err = role.GetRoleByRid(tx)
		if err != nil {
			continue
		}
		roles = append(roles, role)
	}
	log.Info("rids:",rids)
	drp.Role = roles
	//todo 获取所有权限pids
	perIds,err := GetPermissionByRids(rids,tx)
	if err != nil {
		log.Error("GetPermissionByRids Error, ",err)
		return drps,nil
	}
	//todo 获取所有的resources
	resource,err := GetRolePermissionInfoByRid(tx,perIds)
	if err != nil {
		log.Error("GetRolePermissionInfoByRid Error,", err)
	}
	log.Info("roleMids:", rids)
	log.Info("resource:", resource)
	drp.Resource = resource
	log.Info("drp:", drp)
	drps = append(drps, &drp)
	return drps, nil
}

/*getChildrenByDid
获取子域的信息（其中的角色-权限信息与父域中的信息相同）*/
func getChildrenByTid(tid int32, tx *gorm.DB) (children []*TenantTree, err error) {
	tenant := Tenant{Pid: tid}
	tenants, err := tenant.GetTenantsByPid(tx)
	log.Info("==========", tenants)
	if err != nil {
		if err != ErrDoesNotExist {
			return
		} else {
			return nil, nil
		}
	}
	if len(tenants) == 0 {
		return
	}
	for i, v := range tenants {
		log.Info("获取子域的次数", i)
		tree := &TenantTree{Tenant: v}
		//递归获取子域
		trees, err := getChildrenByTid(v.Tid, tx)
		if err != nil {
			children = append(children, tree)
			continue
		}
		if tree == nil {
			children = append(children, tree)
			continue
		}
		tree.Children = trees
		children = append(children, tree)
	}
	return
}
