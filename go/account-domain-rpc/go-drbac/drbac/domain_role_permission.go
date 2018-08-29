package drbac

import (
	. "account-domain-rpc/go-drbac/internal/storage"

	. "account-domain-rpc/go-drbac/common"

	log "github.com/cihub/seelog"

	"github.com/jinzhu/gorm"
)

type DomainRolePermission struct {
	Domain          *Domain
	IsDefaultDomain bool
	RolePermissionInfo
	Children []*DomainTree
}

type DomainTree struct {
	Domain   *Domain
	Children []*DomainTree
}

type RolePermissionInfo struct {
	Role        *Role
	Permissions []*Permission
}

/*获取用户角色权限*/
func getDomainRolePermissionByUid(uid int64, tx *gorm.DB) (drps []*DomainRolePermission, err error) {
	udr := UserDomainRole{Uid: uid}
	udrs, err := udr.GetUserDomainRoleByUid(tx)
	if err != nil {
		log.Errorf("GetUserDomainRoleByUid %s uid is %d", err, uid)
		return
	}
	for _, udr := range udrs {
		drp := &DomainRolePermission{}
		/*获取对应的域*/
		domain := &Domain{Did: udr.Did}
		err = domain.GetByID(tx)
		if err != nil {
			continue
		}
		/*获取对应的角色*/
		role := &Role{Rid: udr.Rid}
		err = role.SetRoleName(tx)
		if err != nil {
			continue
		}
		/*获取对应的权限*/
		rp := RolePermission{Rid: udr.Rid}
		ids, err := rp.GetRolePermissionByRid(tx)
		if err != nil {
			continue
		}
		per := Permission{}
		permissions, err := per.GetDomainsByIDs(ids, tx)
		if err != nil {
			continue
		}
		/*聚合对象*/
		drp.Domain = domain
		drp.IsDefaultDomain = udr.IsDefault
		drp.Role = role
		drp.Permissions = permissions
		/*获取子域*/
		children, err := getChildrenByDid(udr.Did, tx)
		if err != nil {
			continue
		}
		drp.Children = children
		drps = append(drps, drp)
	}
	err = nil
	return
}

/*getChildrenByDid
获取子域的信息（其中的角色-权限信息与父域中的信息相同）*/
func getChildrenByDid(did int64, tx *gorm.DB) (children []*DomainTree, err error) {
	domain := Domain{Pid: did}
	domains, err := domain.GetDomainsByPid(tx)
	log.Info("==========", domains)
	if err != nil {
		if err != ErrDoesNotExist {
			return
		} else {
			return nil, nil
		}
	}
	if len(domains) == 0 {
		return
	}
	for i, v := range domains {
		log.Info("获取子域的次数", i)
		tree := &DomainTree{Domain: v}
		//递归获取子域
		trees, err := getChildrenByDid(v.Did, tx)
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
