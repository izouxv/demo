package drbac
import (
	. "auth/go-drbac/internal/storage"

	"github.com/jinzhu/gorm"

	log "github.com/cihub/seelog"

	"github.com/garyburd/redigo/redis"
)

type UserRoleInfo struct {
	User User
	Role Role
	Permissions []*Permission
}

type TenantUsers struct {
	Tenant
	UserRoleInfos []*UserRoleInfo
}

type UserRoles struct {
	User User
	Roles []*Role
}


func getUserTenantRoleInfo(tid int32,tx *gorm.DB) (userTenantRoleInfos []*UserRoleInfo,err error) {
	udr := UserTenantRole{Tid:tid}
	userRoles,err := udr.GetUserRoleByTid(tx)
	if err != nil {
		log.Errorf("GetUserRoleByDid error is %s",err)
		return
	}
	rp := RolePermission{}
	for _,v := range userRoles {
		userTenantRoleInfo := &UserRoleInfo{}
		rp.Rid = v.Rid
		permissions,err := rp.GetRolePermissionInfoByRid(tx)
		if err != nil {
			log.Errorf("GetRolePermissionInfoByRid error is %s",err)
		}
		userTenantRoleInfo.User = v.User
		userTenantRoleInfo.Role = v.Role
		userTenantRoleInfo.Permissions = permissions
		userTenantRoleInfos = append(userTenantRoleInfos,userTenantRoleInfo)
	}
	return
}
/*删除域中用户*/
func deleteUserTenantRoleInTenant(uid,tid int32,tx *gorm.DB,pool *redis.Pool) (err error) {
	udr := UserTenantRole{Uid:uid,Tid:tid}
	err = Transaction(tx, func(tx *gorm.DB) error {
		udrs,err := udr.GetUserTenantRoleByUid(tx)
		if err != nil {
			return err
		}
		if len(udrs) == 1 {
			/*用户最后一个域，直接删除用户信息*/
			user := User{Uid:uid}
			err = user.DeleteByUID(tx)
			if err != nil {
				return err
			}
			err = udr.DeleteByUidAndTid(tx)
			if err != nil {
				return err
			}
		} else {
			err = udr.DeleteByUidAndTid(tx)
			if err != nil {
				return err
			}
			/*非最后一个域，直接删除域用户关系*/
			if udr.IsDefault {
				/*如果是默认域，则赋值其他域为默认域*/
				maxTime := udr.CreateTime
				for _,v := range udrs {
					if v.CreateTime.After(maxTime) {
						maxTime = v.CreateTime
						udr.Tid = v.Tid
					}
				}
				/*设置最新的域为默认域*/
				err = udr.SetDefaultTenant(tx)
				if err != nil {
					return err
				}
			}
		}
		deleteTokenByUidAndTid(uid,tid,tx,pool)
		return nil
	})
	return
}

//获取全部用户信息及角色信息
func getAllUserRoles(tid int32, page, count int32, tx *gorm.DB) (allUserRoles []*UserRoles,totalCount int32, err error) {
	udr1 := UserTenantRole{Tid:tid}
	uids, totalCount, err := udr1.GetUidsByTid(tx, page, count)
	if err != nil || len(uids) == 0 || totalCount == 0{
		log.Error("GetUidsByDid Error, ", err)
		return
	}
	user := User{}
	users,err := user.GetUserInfoByUids(uids,tx)
	if err != nil {
		log.Errorf("GetUserDomainRoleByDid error is %s",err)
		return
	}
	for _,v := range users {
			udr := UserTenantRole{Uid:v.Uid,Tid:tid}
			roles, err  := udr.GetRolesByUidAndTid(tx)
			if err != nil {
				log.Errorf("GetRolesByUid error is %s",err)
			}
			allUserRoles = append(allUserRoles, &UserRoles{User:*v,Roles:roles})
	}
	return
}

//获取全部用户信息及角色信息
func getAllUserRolesInDomain(did int32,page, count int32, tx *gorm.DB) (allUserRoles []*UserRoles,totalCount int32, err error) {
	udr1 := UserDomainRole{Did:did}
	uids, totalCount, err := udr1.GetUidsByDid(tx, page,count)
	log.Info("totalCount:",totalCount)
	log.Info("len(uids):",len(uids))
	if err != nil || len(uids) == 0 {
		log.Error("GetUidsByDid Error, ", err)
		return
	}
	user := User{}
	users,err := user.GetUserInfoByUids(uids,tx)
	if err != nil {
		log.Errorf("GetUserDomainRoleByDid error is %s",err)
		return
	}
	log.Info("len(users):",len(users))
	for _,v := range users {
		udr := UserDomainRole{Uid:v.Uid,Did:did}
		roles, err  := udr.GetRolesByUidAndDid(tx)
		if err != nil {
			log.Errorf("GetRolesByUid error is %s",err)
		}
		allUserRoles = append(allUserRoles, &UserRoles{User:*v,Roles:roles})
	}
	log.Info("len(allUserRoles):",len(allUserRoles))
	return
}
