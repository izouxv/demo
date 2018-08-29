package drbac

import (
	"github.com/jinzhu/gorm"

	"github.com/garyburd/redigo/redis"

	. "account-domain-rpc/go-drbac/internal/storage"

	log "github.com/cihub/seelog"

	. "account-domain-rpc/go-drbac/common"
)

type DrbacServer struct {
	db   *gorm.DB
	pool *redis.Pool
}

func NewDrbacServer(db *gorm.DB, pool *redis.Pool) (*DrbacServer, error) {
	ds := &DrbacServer{db: db, pool: pool}
	return ds, nil
}

type Auth interface {
	Authentication(username, password string) (*UserToken, bool, error)
	Authorization(did int64, url, opt, token string) (*UserToken, bool)
	AuthenticationUserDomain(uid, did int64) bool
	LogoutToken(token string) (err error)
	GetAuthorizationInfo(token string) (userToken *UserToken)
	GetAuthorizationInfoByUid(uid int64) (userToken *UserToken)
}

type RoleBase interface {
	GetRoles() (roles []*Role)
	GetRoleByRID(rid int32) (role *Role)
	GetRolePermissions() (RolePermissionInfos []*RolePermissionInfo, err error)
}

type DomainBase interface {
	CreateDomainBaseDomain(pid, uid int64, domainName string) (*Domain, error)
	CreateUserBaseDomain(username, password, nickname string, did int64, rid int32) (*User, error)
	GetUserBaseDomain(did int64) (domainUsers *DomainUsers, err error)
	GetUserCountBaseDomain(did int64) (count int32, err error)
	GetDomain(did int64) (domain Domain, err error)
	GetChildrenDomains(pid int64) (domains []*Domain, err error)
	UpdateDomainName(did int64, domainName, token string) (domain *Domain, err error)
	DeleteDomain(did int64, token string) (err error)
	GetDomains(token string) (domainRolePermissions []*DomainRolePermission)
	DeleteUserDomain(uid, did int64) (err error)
	IsExistUserInDomain(uid, did int64) (exist bool)
	UpdateUserRole(did, uid int64, rid int32) (err error)
	GetDefaultDomain(uid int64) (domain *Domain, err error)
	GetDomainDepth(did int64) (int32, error)
}

type UserBase interface {
	UpdatePassword(uid int64, password string) (err error)
	UpdateNickname(uid int64, nickname string) (err error)
	GetUserByUid(uid int64) (user *User, err error)
	GetUserByUsername(username string) (user *User, err error)
	GetTokenByUid(uid int64) (token string)
	GetTokenByUsername(username string) (token string)
}

/*Authentication
 *认证
 */
func (ds *DrbacServer) Authentication(username, password string) (*UserToken, bool, error) {
	return authentication(username, password, ds.db, ds.pool)
}

/*Authorization
 *授权
 */
func (ds *DrbacServer) Authorization(did int64, url, opt, token string) (*UserToken, bool) {
	return authorization(did, url, opt, token, ds.pool)
}

/*AuthenticationUserDomain
认证uid操作的did是否有权限
*/
func (ds *DrbacServer) AuthenticationUserDomain(uid, did int64) bool {
	userToken := ds.GetAuthorizationInfoByUid(uid)
	if userToken == nil {
		return false
	}
	return judgmentDomain(userToken, did)
}

/*LogoutToken
 *注销token
 */
func (ds *DrbacServer) LogoutToken(token string) (err error) {
	deleteToken(token, ds.pool)
	return
}

/*GetAuthorizationInfo
 *基于token获取认证完的用户信息
 */
func (ds *DrbacServer) GetAuthorizationInfo(token string) (userToken *UserToken) {
	userToken = getToken(token, ds.pool)
	return
}

/*GetAuthorizationInfoByUid
基于uid获取认证信息
*/
func (ds *DrbacServer) GetAuthorizationInfoByUid(uid int64) (userToken *UserToken) {
	user, err := ds.GetUserByUid(uid)
	if err != nil {
		return
	}
	userToken = getToken(ds.GetTokenByUsername(user.Username), ds.pool)
	return
}

/*GetRoles
获取所有的角色信息
*/
func (ds *DrbacServer) GetRoles() (roles []*Role) {
	if len(RoleMap) == 0 {
		SetRoleMap(ds.db)
	}
	for i, v := range RoleMap {
		role := &Role{Rid: i, RoleName: v}
		roles = append(roles, role)
	}
	return
}

/*GetRoleByRID
基于rid获取角色信息
*/
func (ds *DrbacServer) GetRoleByRID(rid int32) (role *Role) {
	role = &Role{Rid: rid, RoleName: RoleMap[rid]}
	return
}

/*GetRolePermissions
获取数据库中的角色权限信息
*/
func (ds *DrbacServer) GetRolePermissions() (RolePermissionInfos []*RolePermissionInfo, err error) {
	for i, v := range RoleMap {
		role := &Role{Rid: i, RoleName: v}
		rp := RolePermission{Rid: i}
		ids, err := rp.GetRolePermissionByRid(ds.db)
		if err != nil {
			break
		}
		p := Permission{}
		permissions, err := p.GetDomainsByIDs(ids, ds.db)
		if err != nil {
			break
		}
		RolePermissionInfos = append(RolePermissionInfos, &RolePermissionInfo{Role: role, Permissions: permissions})
	}
	return
}

/*CreateUserBaseDomain
基于域创建用户
*/
func (ds *DrbacServer) CreateUserBaseDomain(username, password, nickname string, did int64, rid int32) (*User, error) {
	user := &User{Username: username, Nickname: nickname, Password: password}
	err := Transaction(ds.db, func(tx *gorm.DB) error {
		err := user.Create(tx)
		udr := UserDomainRole{Uid: user.Uid, Did: did, Rid: rid, IsDefault: true}
		if err != nil {
			log.Error("CreateUserBaseDomain", err)
			if err == ErrAlreadyExists {
				/*判断子域中存在*/
				drps, err := getDomainRolePermissionByUid(user.Uid, tx)
				if err != nil {
					log.Errorf("GetDomainRolePermissionByUid error %s uid is %d", err, user.Uid)
					return err
				}
				for _, v := range drps {
					if v.Domain.Did == did {
						/*did是第一级的父域中存在*/
						return ErrAlreadyExists
					}
					/*did是子域id*/
					if judgmentChildren(v.Children, did) {
						return ErrAlreadyExists
					}
				}
			} else {
				log.Errorf("系统异常")
				return err
			}
			udr.IsDefault = false
		}
		err = udr.Create(tx)
		if err != nil {
			log.Error("CreateUserBaseDomain", err)
			return err
		}
		userToken := ds.GetAuthorizationInfo(ds.GetTokenByUid(user.Uid))
		if userToken != nil {
			_, err = reAuthenticationByUid(user.Uid, tx, ds.pool, reAuthentication)
			if err != nil {
				log.Errorf("reAuthenticationByUid error %s", err)
				return err
			}
		}
		return nil
	})
	return user, err
}

/*CreateDomainBaseDomain
基于父域创建子域
*/
func (ds *DrbacServer) CreateDomainBaseDomain(pid, uid int64, domainName string) (*Domain, error) {
	domain := &Domain{DomainName: domainName, Pid: pid}
	err := Transaction(ds.db, func(tx *gorm.DB) error {
		err := domain.Create(tx)
		if err != nil {
			log.Error("CreateDomainBaseDomain", err)
			return err
		}
		_, err = reAuthenticationByUid(uid, tx, ds.pool, reAuthentication)
		if err != nil {
			return err
		}
		return nil
	})
	return domain, err
}

/*GetDomainDepth
获取域的深度
*/
func (ds *DrbacServer) GetDomainDepth(did int64) (int32, error) {
	domain := &Domain{Did: did}
	return domain.GetDomainDepth(ds.db)
}

/*GetUserBaseDomain
获取域中成员信息
*/
func (ds *DrbacServer) GetUserBaseDomain(did int64) (domainUsers *DomainUsers, err error) {
	domainUsers = &DomainUsers{}
	domain := Domain{Did: did}
	err = domain.GetByID(ds.db)
	if err != nil {
		return
	}
	userDomainRoles, err := getUserDomainRoleInfo(did, ds.db)
	if err != nil {
		return
	}
	domainUsers.UserRoleInfos = userDomainRoles
	domainUsers.Domain = domain
	return
}

/*GetUserCountBaseDomain
获取域中成员数量
*/
func (ds *DrbacServer) GetUserCountBaseDomain(did int64) (count int32, err error) {
	udr := UserDomainRole{Did: did}
	count, err = udr.GetUserCountByDid(ds.db)
	return
}

/*GetDomain
获取单个域信息
*/
func (ds *DrbacServer) GetDomain(did int64) (domain Domain, err error) {
	domain.Did = did
	err = domain.GetByID(ds.db)
	if err != nil {
		return
	}
	return
}

/*GetChildrenDomains
基于父id查询子域信息
*/
func (ds *DrbacServer) GetChildrenDomains(pid int64) (domains []*Domain, err error) {
	domain := Domain{Pid: pid}
	return domain.GetDomainsByPid(ds.db)
}

/*UpdateDomainName
基于did修改域信息
*/
func (ds *DrbacServer) UpdateDomainName(did int64, domainName, token string) (domain *Domain, err error) {
	domain = &Domain{Did: did, DomainName: domainName}
	err = Transaction(ds.db, func(tx *gorm.DB) error {
		err = domain.Update(ds.db)
		if err != nil {
			return err
		}
		_, err = reAuthentication(token, tx, ds.pool)
		if err != nil {
			return err
		}
		return nil
	})
	return
}

/*DeleteDomain
基于did删除域
*/
func (ds *DrbacServer) DeleteDomain(did int64, token string) (err error) {
	err = Transaction(ds.db, func(tx *gorm.DB) error {
		domain := Domain{Did: did, Pid: did}
		if domain.IsExistDomainByPid(tx) {
			log.Errorf("有子域不能删除")
			return ErrCanNotDelete
		}
		udr := UserDomainRole{Did: did}
		err = domain.DeleteByDID(ds.db)
		if err != nil {
			return err
		}
		err = udr.DeleteByDid(tx)
		if err != nil {
			if err != ErrDoesNotExist {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return
	}
	_, err = reAuthentication(token, ds.db, ds.pool)
	if err != nil {
		return err
	}
	return
}

/*GetDomains
基于token获取用户的域信息子父级列表
*/
func (ds *DrbacServer) GetDomains(token string) (domainRolePermissions []*DomainRolePermission) {
	userToken := getToken(token, ds.pool)
	domainRolePermissions = userToken.DomainRolePermissions
	return
}

/*GetDefaultDomain
获取用户默认域
*/
func (ds *DrbacServer) GetDefaultDomain(uid int64) (domain *Domain, err error) {
	udr := UserDomainRole{Uid: uid}
	udrs, err := udr.GetUserDomainRoleByUid(ds.db)
	if err != nil || len(udrs) == 0 {
		err = ErrDoesNotExist
		return
	}
	for _, v := range udrs {
		if v.IsDefault {
			domain = &Domain{Did: v.Did}
			err = domain.GetByID(ds.db)
			return
		}
	}
	return
}

/*DeleteUserDomain
删除域中用户
*/
func (ds *DrbacServer) DeleteUserDomain(uid, did int64) (err error) {
	err = deleteUserDomainRoleInDomain(uid, did, ds.db, ds.pool)
	return
}

/*IsExistUserInDomain
用户是否存在域中
*/
func (ds *DrbacServer) IsExistUserInDomain(uid, did int64) (exist bool) {
	udr := UserDomainRole{Uid: uid, Did: did}
	err := udr.GetUserRoleByUidAndDid(ds.db)
	if err == ErrDoesNotExist {
		return
	}
	exist = true
	return
}

/*UpdateUserRole
修改用户在域中的角色
*/
func (ds *DrbacServer) UpdateUserRole(did, uid int64, rid int32) (err error) {
	err = Transaction(ds.db, func(tx *gorm.DB) error {
		udr := UserDomainRole{Uid: uid, Did: did, Rid: rid}
		err = udr.UpdateUserRoleByDidAndUid(tx)
		if err != nil {
			return err
		}
		_, err = reAuthenticationByUid(uid, tx, ds.pool, reAuthentication)
		if err != nil {
			log.Errorf("重新认证失败,可能用户未登录")
		}
		return nil
	})
	return
}

/*UpdatePassword
修改用户认证密码
*/
func (ds *DrbacServer) UpdatePassword(uid int64, password string) (err error) {
	user := User{Uid: uid, Password: password}
	err = user.UpdatePassword(ds.db)
	return
}

/*UpdateNickname
修改用户昵称
*/
func (ds *DrbacServer) UpdateNickname(uid int64, nickname string) (err error) {
	err = Transaction(ds.db, func(tx *gorm.DB) error {
		user := User{Uid: uid, Nickname: nickname}
		err = user.UpdateNickname(ds.db)
		if err != nil {
			return err
		}
		_, err = reAuthenticationByUid(uid, ds.db, ds.pool, reAuthentication)
		if err != nil {
			return err
		}
		return nil
	})
	return
}

/*GetUserByUid
基于uid获取username
*/
func (ds *DrbacServer) GetUserByUid(uid int64) (user *User, err error) {
	user = &User{Uid: uid}
	err = user.GetUserByUID(ds.db)
	return
}

/*GetUserByUsername
根据用户名查询用户
*/
func (ds *DrbacServer) GetUserByUsername(username string) (user *User, err error) {
	user = &User{Username: username}
	err = user.GetUserByUsername(ds.db)
	return
}

/*GetTokenByUid
基于uid湖片区用户对应的token
*/
func (ds *DrbacServer) GetTokenByUid(uid int64) (token string) {
	token = getTokenByUid(uid, ds.db, ds.pool)
	return
}

/*GetTokenByUsername
基于username获取用户对应的token
*/
func (ds *DrbacServer) GetTokenByUsername(username string) (token string) {
	token = getTokenByUsername(username, ds.pool)
	return
}
