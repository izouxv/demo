package drbac

import (
	. "account-domain-rpc/go-drbac/internal/storage"

	"github.com/jinzhu/gorm"

	log "github.com/cihub/seelog"

	"github.com/garyburd/redigo/redis"
)

type UserRoleInfo struct {
	User        User
	Role        Role
	Permissions []*Permission
}

type DomainUsers struct {
	Domain
	UserRoleInfos []*UserRoleInfo
}

func getUserDomainRoleInfo(did int64, tx *gorm.DB) (userDomainRoleInfos []*UserRoleInfo, err error) {
	udr := UserDomainRole{Did: did}
	userRoles, err := udr.GetUserRoleByDid(tx)
	if err != nil {
		log.Errorf("GetUserRoleByDid error is %s", err)
		return
	}
	rp := RolePermission{}
	for _, v := range userRoles {
		userDomainRoleInfo := &UserRoleInfo{}
		rp.Rid = v.Rid
		permissions, err := rp.GetRolePermissionInfoByRid(tx)
		if err != nil {
			log.Errorf("GetRolePermissionInfoByRid error is %s", err)
		}
		userDomainRoleInfo.User = v.User
		userDomainRoleInfo.Role = v.Role
		userDomainRoleInfo.Permissions = permissions
		userDomainRoleInfos = append(userDomainRoleInfos, userDomainRoleInfo)
	}
	return
}

/*删除域中用户*/
func deleteUserDomainRoleInDomain(uid, did int64, tx *gorm.DB, pool *redis.Pool) (err error) {
	udr := UserDomainRole{Uid: uid, Did: did}
	err = Transaction(tx, func(tx *gorm.DB) error {
		udrs, err := udr.GetUserDomainRoleByUid(tx)
		if err != nil {
			return err
		}
		if len(udrs) == 1 {
			/*用户最后一个域，直接删除用户信息*/
			user := User{Uid: uid}
			err = user.DeleteByUID(tx)
			if err != nil {
				return err
			}
			err = udr.DeleteByUidAndDid(tx)
			if err != nil {
				return err
			}
		} else {
			err = udr.DeleteByUidAndDid(tx)
			if err != nil {
				return err
			}
			/*非最后一个域，直接删除域用户关系*/
			if udr.IsDefault {
				/*如果是默认域，则赋值其他域为默认域*/
				maxTime := udr.CreateTime
				for _, v := range udrs {
					if v.CreateTime.After(maxTime) {
						maxTime = v.CreateTime
						udr.Did = v.Did
					}
				}
				/*设置最新的域为默认域*/
				err = udr.SetDefaultDomain(tx)
				if err != nil {
					return err
				}
			}
		}
		deleteTokenByUid(uid, tx, pool)
		return nil
	})
	return
}
