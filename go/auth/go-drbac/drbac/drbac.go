package drbac

import (
	"github.com/jinzhu/gorm"

	"github.com/garyburd/redigo/redis"

	. "auth/go-drbac/internal/storage"

	log "github.com/cihub/seelog"

	. "auth/go-drbac/common"
	"auth/util"
	"errors"
	"github.com/gin-gonic/gin/json"
	"time"
)

type DrbacServer struct {
	db   *gorm.DB
	pool *redis.Pool
}

type Modules struct {
	Mid        int32
	ModuleName string
	Operations []string
}

type ModuleInfo struct {
	Mid        int32
	ModuleName string
}

type TenantInfo struct {
	Tid         int32
	TenantName  string
	TenantURL   string
	TenantState int32
	Description string
	Contacts    string
	Email       string
	Phone       string
	Icon        string
	Logo        string
}

func NewDrbacServer(db *gorm.DB, pool *redis.Pool) (*DrbacServer, error) {
	ds := &DrbacServer{db: db, pool: pool}
	return ds, nil
}

/*AuthenticationWithTid
 *认证 指定租户
 */
func (ds *DrbacServer) Authentication(uid int32, password string) (bool, int32, error) {
	return authentication(uid, password, ds.db, ds.pool)
}

/*AuthenticationWithTid
 *认证 指定租户
 */
func (ds *DrbacServer) AuthenticationWithTid(username, password string, tid int32) (*UserTokenTenants, bool, int32, error) {
	return authenticationWithTid(username, password, tid, ds.db, ds.pool)
}

/*AuthenticationWithoutTid
 *认证 不指定租户
 */
func (ds *DrbacServer) AuthenticationWithDid(username, password string, did int32) (*UserToken, bool, int32, error) {
	return authenticationWithDid(username, password, did, ds.db, ds.pool)
}

/*Authorization
 *授权
 */
func (ds *DrbacServer) AuthorizationTenant(tid int32, url, opt, token string) (*UserToken, bool) {
	return authorizationTenant(tid, url, opt, token, ds.pool)
}

/*Authorization
 *授权
 */
func (ds *DrbacServer) AuthorizationDomain(did int32, url, opt, token string) (*UserToken, bool) {
	return authorizationDomain(did, url, opt, token, ds.pool)
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
func (ds *DrbacServer) GetAuthorizationInfoByUid(uid, tid int32) (userToken *UserToken) {
	user, err := ds.GetUserByUid(uid)
	if err != nil {
		log.Error("GetUserByUid Error,", err)
		return
	}
	tenant, err := ds.GetTenant(tid)
	if err != nil || tenant.TenantName == "" {
		log.Error("GetDomainByDid Error,", err, " tenant:", tenant)
		return
	}
	userToken = getToken(ds.GetTokenByUsernameAndTenantName(user.Username, tenant.TenantName), ds.pool)
	return
}

/*GetRoles
获取所有的角色信息
*/
func (ds *DrbacServer) GetRoles(tid int32, page, count int32) (roles []*Role, totalCount int32, err error) {
	role := Role{Tid: tid}
	return role.GetRoles(ds.db, page, count)
}

func (ds *DrbacServer) GetDomainRoles(did int32, page, count int32) (roles []*Role, totalCount int32, err error) {
	role := Role{Did: did}
	return role.GetDomainRoles(ds.db, page, count)
}

//todo GetResource
func (ds *DrbacServer) GetResources(page, count int32, resName, resOpt string) (resources []*Resource, totalCount int32, err error) {
	resource := Resource{ResName:resName,ResOpt:resOpt}
	return resource.GetResources(ds.db, page, count)
}

func (ds *DrbacServer) CreateResource(resName, resUrl, resOpt string) (err error) {
	resource := Resource{ResName:resName,ResOpt:resOpt,ResUrl:resUrl}
	return resource.Create(ds.db)
}
func (ds *DrbacServer) UpdateResource(resId int32, resName, resUrl, resOpt string) (err error) {
	resource := Resource{ResId:resId, ResName:resName,ResOpt:resOpt,ResUrl:resUrl}
	return resource.UpdateByResId(ds.db)
}
func (ds *DrbacServer) DeleteResource(resId int32) (err error) {
	resource := Resource{ResId:resId}
	return resource.DeleteByResId(ds.db)
}
/*GetRoleByRID
基于rid获取角色信息
*/
func (ds *DrbacServer) GetRoleByRid(rid int32) (role *Role, err error) {
	role = &Role{Rid: rid}
	err = role.GetRoleByRid(ds.db)
	return
}

/*GetRoleByRID
基于rid获取角色信息
*/
func (ds *DrbacServer) GetRoleByRidAndTid(rid int32, tid int32) (err error) {
	role := &Role{Rid: rid, Tid: tid}
	err = role.GetRoleByRid(ds.db)
	return
}

/*GetRolePermissions
获取数据库中的角色权限信息
*/
//func (ds *DrbacServer) GetRolePermissions() (RolePermissionInfos []*RolePermissionInfo,err error) {
//	for i,v :=range RoleMap {
//		role := &Role{Rid:i,RoleName:v}
//		rp := RolePermission{Rid:i}
//		ids,err := rp.GetRolePermissionByRid(ds.db)
//		if err != nil {
//			break
//		}
//		p := Permission{}
//		permissions,err := p.GetDomainsByIDs(ids,ds.db)
//		if err != nil {
//			break
//		}
//		RolePermissionInfos = append(RolePermissionInfos,&RolePermissionInfo{Role:role,Permissions:permissions})
//	}
//	return
//}

/*CreateUserBaseDomain
基于租户创建用户
*/
func (ds *DrbacServer) CreateUserBaseTenant(username, password, nickname, token string, tid int32, rid int32) (*User, error) {
	user := &User{Username: username, Nickname: nickname, Password: password}
	err := Transaction(ds.db, func(tx *gorm.DB) error {
		err := user.Create(tx)
		udr := UserTenantRole{Uid: user.Uid, Tid: tid, Rid: rid, IsDefault: true}
		if err != nil {
			log.Error("CreateUserBaseDomain", err)
			if err == ErrAlreadyExists {
				/*判断子域中存在*/
				drps, err := getTenantRoleResourceByUidAndTid(user.Uid, tid, tx)
				if err != nil {
					log.Errorf("GetDomainRolePermissionByUid error %s uid is %d", err, user.Uid)
					return err
				}
				if drps[0].Tenant.Tid == tid {
					/*did是第一级的父域中存在*/
					return ErrAlreadyExists
				}
			}
		} else {
			log.Errorf("系统异常")
			return err
		}
		udr.IsDefault = false
		err = udr.Create(tx)
		if err != nil {
			log.Error("CreateUserBaseDomain", err)
			return err
		}
		userToken := ds.GetAuthorizationInfo(ds.GetTokenByUidAndTid(user.Uid, tid))
		if userToken != nil {
			_, err = reAuthenticationByUid(token, tid, tx, ds.pool, reAuthentication)
			if err != nil {
				log.Errorf("reAuthenticationByUid error %s", err)
				return err
			}
		}
		return nil
	})
	return user, err
}

/*
创建租户
*/
func (ds *DrbacServer) CreateTenantBaseTenant(did, pid int32, tenantInfo TenantInfo, token, inviterNickname string) (*Tenant, error) {
	tenant := &Tenant{
		TenantName:  tenantInfo.TenantName,
		TenantURL:   tenantInfo.TenantURL,
		TenantState: 1,
		Description: tenantInfo.Description,
		Contacts:    tenantInfo.Contacts,
		Email:       tenantInfo.Email,
		Phone:       tenantInfo.Phone,
		Pid:         pid,
		Did:         did,
		Icon:        tenantInfo.Icon,
		Logo:        tenantInfo.Logo,
	}
	err := tenant.Create(ds.db)
	if err != nil || tenant.Tid == 0 {
		log.Error("CreateTenantError", err)
		return nil, errors.New("CreateTenantError")
	}

	//todo 创建默认权限模块
	//服务模块
	module := Module{ModuleName: "服务", ModuleTid: tenant.Tid}
	err = module.Create(ds.db)
	if err != nil || module.Mid == 0 {
		log.Error("module.Create Error,", err)
		return nil, err
	}
	mp := ModulePermission{Mid: module.Mid, Pid: 1}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 2}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 3}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 4}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	//用户模块
	module = Module{ModuleName: "用户", ModuleTid: tenant.Tid}
	err = module.Create(ds.db)
	if err != nil {
		log.Error("module.Create Error")
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 5}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 6}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 7}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 8}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	//角色模块
	module = Module{ModuleName: "角色", ModuleTid: tenant.Tid}
	err = module.Create(ds.db)
	if err != nil {
		log.Error("module.Create Error")
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 9}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 10}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 11}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 12}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	//日志模块
	module = Module{ModuleName: "日志", ModuleTid: tenant.Tid}
	err = module.Create(ds.db)
	if err != nil {
		log.Error("module.Create Error")
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 13}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 14}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 15}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 16}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	//设置模块
	module = Module{ModuleName: "设置", ModuleTid: tenant.Tid}
	err = module.Create(ds.db)
	if err != nil {
		log.Error("module.Create Error")
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 17}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 18}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 19}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 20}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	//密码模块
	module = Module{ModuleName: "密码", ModuleTid: tenant.Tid}
	err = module.Create(ds.db)
	if err != nil {
		log.Error("module.Create Error")
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 21}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 22}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 23}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 24}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	//文件库模块
	module = Module{ModuleName: "文件库", ModuleTid: tenant.Tid}
	err = module.Create(ds.db)
	if err != nil {
		log.Error("module.Create Error")
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 25}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 26}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 27}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 28}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	//保留规则模块
	module = Module{ModuleName: "保留规则", ModuleTid: tenant.Tid}
	err = module.Create(ds.db)
	if err != nil {
		log.Error("module.Create Error")
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 29}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 30}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 31}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 32}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	//财务总览模块
	module = Module{ModuleName: "财务总览", ModuleTid: tenant.Tid}
	err = module.Create(ds.db)
	if err != nil {
		log.Error("module.Create Error")
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 33}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 34}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 35}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 36}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	//消费总览模块
	module = Module{ModuleName: "消费总览", ModuleTid: tenant.Tid}
	err = module.Create(ds.db)
	if err != nil {
		log.Error("module.Create Error")
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 37}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 38}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 39}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 40}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}

	//设备管理服务模块
	module = Module{ModuleName: "设备管理服务", ModuleTid: tenant.Tid}
	err = module.Create(ds.db)
	if err != nil {
		log.Error("module.Create Error")
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 81}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 82}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 83}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 84}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}

	//KMS服务模块
	module = Module{ModuleName: "KMS服务", ModuleTid: tenant.Tid}
	err = module.Create(ds.db)
	if err != nil {
		log.Error("module.Create Error")
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 85}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 86}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 87}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}
	mp = ModulePermission{Mid: module.Mid, Pid: 88}
	err = mp.Create(ds.db)
	if err != nil {
		log.Error("mp.Create Error,", err)
		return nil, err
	}

	//todo 创建默认角色
	module = Module{ModuleTid: tenant.Tid}
	ReplyModules, err := module.GetModulesByTid(ds.db)
	var modules []*Modules
	admin := []string{"ADMIN", "READ", "UPDATE", "CREATE"}
	for _, v := range ReplyModules {
		log.Info("ReplyModules.Mid:", v.Mid)
		modules = append(modules, &Modules{Mid: v.Mid, Operations: admin})
	}
	log.Info("len(modules):", len(modules))
	rid, err := ds.AddRole("管理员", "新建租户的默认角色", modules, tenant.Tid)
	if err != nil {
		log.Error("AddRole Error")
		return nil, err
	}
	//todo 创建管理员用户
	emailToken := string(Krand(32, KC_RAND_KIND_ALL))
	user := &User{Username: tenantInfo.Email, Nickname: tenantInfo.Contacts, Token: emailToken, Tid: tenant.Tid}
	err = user.Create(ds.db)
	if err != nil || user.Uid == 0 {
		log.Info("user.Create Error,", err)
		return nil, errors.New("user.Create Error")
	}
	log.Info("AddUser UserInfo:", user)
	//todo 给创建者赋予默认角色
	urt0 := &UserTenantRole{Uid: user.Uid, Tid: tenant.Tid, Rid: 0, IsDefault: true}
	err = urt0.Create(ds.db)
	if err != nil {
		log.Error("CreateUserTenantRoleError", err)
		return nil, err
	}
	urt := &UserTenantRole{Uid: user.Uid, Tid: tenant.Tid, Rid: rid, IsDefault: true}
	err = urt.Create(ds.db)
	if err != nil {
		log.Error("CreateUserTenantRoleError", err)
		return nil, err
	}
	//todo 添加默认开通服务
	//设备管理服务
	tsp := TenantServicePolicy{Sid:1,Tid:tenant.Tid,Pid:12,StartTime:time.Now()}
	err = tsp.Create(ds.db)
	if err != nil {
		log.Error("TenantServicePolicy Create Error", err)
		return nil, err
	}
	//模拟器服务
	tsp2 := TenantServicePolicy{Sid:2,Tid:tenant.Tid,Pid:12,StartTime:time.Now()}
	err = tsp2.Create(ds.db)
	if err != nil {
		log.Error("TenantServicePolicy Create Error", err)
		return nil, err
	}
	//todo 添加租户账户
	ta := TenantAccount{Tid:tenant.Tid,Balance:0,CreateTime:time.Now(),UpdateTime:time.Now()}
	ta.CreateTenantAccount(ds.db)
	if err != nil {
		log.Error("CreateTenantAccount Error", err)
		return nil, err
	}
	//todo 发送邀请邮件
	sendData := util.SendData{
		To:       user.Username,
		Nickname: "admin",
		Token:    user.Token,
		Company:  "元安物联",
		Url:      tenantInfo.TenantURL,
	}
	mailInfo := util.MailInfo{
		TemId:     3,
		EmailAddr: user.Username,
		SendData:  sendData,
	}
	req, err := json.Marshal(mailInfo)
	if err != nil {
		log.Info("json.Marshal(mailInfo) Error,", err)
		return nil, err
	}
	ok := util.HttpSendMailPost(req)
	log.Info("HttpSendMailPost:", ok)
	//_, err = reAuthenticationByUid(token, tenant.Tid, tx, ds.pool, reAuthentication)
	//if err != nil {
	//	return err
	//}
	log.Info("Create Tenant Successful")
	return tenant, err
}

/*GetDomainDepth
获取域的深度
*/
func (ds *DrbacServer) GetTenantDepth(tid int32) (int32, error) {
	tenant := &Tenant{Tid: tid}
	return tenant.GetTenantDepth(ds.db)
}

/*GetUserBaseDomain
获取租户中成员信息
*/
func (ds *DrbacServer) GetUserRoles(tid int32, page, count int32) (userRoles []*UserRoles, totalCount int32, err error) {
	userRoles, totalCount, err = getAllUserRoles(tid, page, count, ds.db)
	return
}

/*GetUserBaseDomain
获取域中成员信息
*/
func (ds *DrbacServer) GetUserRolesInDomain(did int32, page, count int32) (userRoles []*UserRoles, totalCount int32, err error) {
	userRoles, totalCount, err = getAllUserRolesInDomain(did, page, count, ds.db)
	return
}

/*GetUserCountBaseDomain
获取域中成员数量
*/
func (ds *DrbacServer) GetUserCountBaseTenant(tid int32) (count int32, err error) {
	udr := UserTenantRole{Tid: tid}
	count, err = udr.GetUserCountByTid(ds.db)
	return
}

/*GetDomain
获取单个域信息
*/
func (ds *DrbacServer) GetTenant(tid int32) (tenant Tenant, err error) {
	tenant.Tid = tid
	err = tenant.GetByID(ds.db)
	if err != nil {
		return
	}
	return
}

/*GetDomain
获取单个域信息
*/
func (ds *DrbacServer) GetTenants(did, uid int32) (tenants []*Tenant, err error) {
	if did == 1 {
		//todo 通过uid查询acl
		uta := UserTenantACL{Uid: uid}
		tids, err := uta.GetTidsByUid(ds.db)
		if err != nil {
			log.Error("uta.GetTidsByUid Error, ", err)
			return nil, err
		}
		tenant := Tenant{}
		tenants, err = tenant.GetTenantsByTids(tids, ds.db)
		if err != nil {
			log.Error("tenant.GetTenantsByTids Error,", err)
			return nil, err
		}
		return tenants, err
	}
	tenant := Tenant{Did: did}
	tenants, err = tenant.GetByDid(ds.db)
	if err != nil {
		log.Error("tenant.GetByDid, Error", err)
	}
	return
}

/*GetChildrenDomains
基于父id查询子域信息
*/
func (ds *DrbacServer) GetChildrenTenants(pid int32) (tenants []*Tenant, err error) {
	tenant := Tenant{Pid: pid}
	return tenant.GetTenantsByPid(ds.db)
}

/*UpdateDomainName
基于did修改域信息
*/
func (ds *DrbacServer) UpdateTenantInfo(tenantInfo TenantInfo, token string) (tenant *Tenant, err error) {
	if tenantInfo.Tid == 0 {
		log.Error("tid为空")
		return nil, errors.New("tid为空")
	}
	tenant = &Tenant{
		Tid:         tenantInfo.Tid,
		TenantName:  tenantInfo.TenantName,
		TenantURL:   tenantInfo.TenantURL,
		Description: tenantInfo.Description,
		Contacts:    tenantInfo.Contacts,
		Email:       tenantInfo.Email,
		Phone:       tenantInfo.Phone,
		Icon:        tenantInfo.Icon,
		Logo:        tenantInfo.Logo,
	}
	err = Transaction(ds.db, func(tx *gorm.DB) error {
		err = tenant.Update(ds.db)
		if err != nil {
			if err == ErrDoesNotExist {
				log.Info("租户信息无修改")
				return nil
			}
			return err
		}
		//重认证（刷新缓存），待修改
		//_, err = reAuthentication(token, tenantInfo.Tid, tx, ds.pool)
		//if err != nil {
		//	return err
		//}
		return nil
	})
	return
}

/*UpdateDomainName
基于tid修改租户状态
*/
func (ds *DrbacServer) UpdateTenantState(tid int32, state int32, token string) (tenant *Tenant, err error) {
	if tid == 0 || state == 0 {
		log.Error("tid为空或state为空")
		return nil, errors.New("tid为空")
	}
	tenant = &Tenant{
		Tid:         tid,
		TenantState: state,
	}
	err = Transaction(ds.db, func(tx *gorm.DB) error {
		err = tenant.Update(ds.db)
		if err != nil {
			return err
		}
		//重认证（刷新缓存），待修改
		//_, err = reAuthentication(token, tenantInfo.Tid, tx, ds.pool)
		//if err != nil {
		//	return err
		//}
		return nil
	})
	return
}

/*DeleteDomain
基于did删除域
*/
func (ds *DrbacServer) DeleteTenant(tid int32, token string) (err error) {

	err = Transaction(ds.db, func(tx *gorm.DB) error {
		tenant := Tenant{Tid: tid, Pid: tid}
		if tenant.IsExistTenantByPid(tx) {
			log.Errorf("有子域不能删除")
			return ErrCanNotDelete
		}
		//todo 删除所有租户角色
		role := Role{Tid: tid}
		err = role.DeleteByTid(tx)
		if err != nil {
			log.Error("role.DeleteByTid Error,", err)
			return err
		}
		//todo 删除所有租户用户关系
		udr := UserTenantRole{Tid: tid}

		err = udr.DeleteByTid(tx)
		if err != nil {
			log.Error("udr.DeleteByTid Error,", err)
			return err
		}
		//todo 删除所有租户下用户
		user := User{Tid: tid}
		err := user.DeleteByTid(tx)
		if err != nil {
			log.Error("user.DeleteByTid Error,", err)
			return err
		}

		err = tenant.DeleteByTID(ds.db)
		if err != nil {
			log.Error("tenant.DeleteByTID Error,", err)
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
	//_, err = reAuthentication(token, tid, ds.db, ds.pool)
	//if err != nil {
	//	return err
	//}
	return
}

/*GetDomains
基于token获取用户的域信息子父级列表
*/
//func (ds *DrbacServer) GetDomains(token string)(domainRolePermissions []*DomainRolePermission) {
//	userToken := getToken(token,ds.pool)
//	domainRolePermissions = userToken.DomainRolePermissions
//	return
//}
/*GetDefaultDomain
获取用户默认域
*/
func (ds *DrbacServer) GetDefaultTenant(uid int32) (tenant *Tenant, err error) {
	udr := UserTenantRole{Uid: uid}
	udrs, err := udr.GetUserTenantRoleByUid(ds.db)
	if err != nil || len(udrs) == 0 {
		err = ErrDoesNotExist
		return
	}
	for _, v := range udrs {
		if v.IsDefault {
			tenant = &Tenant{Tid: v.Tid}
			err = tenant.GetByID(ds.db)
			return
		}
	}
	return
}

/*DeleteUserDomain
删除域中用户
*/
func (ds *DrbacServer) DeleteUserTenant(uid, tid int32) (err error) {
	err = deleteUserTenantRoleInTenant(uid, tid, ds.db, ds.pool)
	return
}

/*IsExistUserInDomain
用户是否存在域中
*/
func (ds *DrbacServer) IsExistUserInDomain(uid, tid int32) (exist bool) {
	udr := UserTenantRole{Uid: uid, Tid: tid}
	_, err := udr.GetUserRoleByUidAndTid(ds.db)
	if err == ErrDoesNotExist {
		return
	}
	exist = true
	return
}

/*UpdateUserRole
修改用户在租户中的角色
*/
func (ds *DrbacServer) UpdateUserRole(uid, tid int32, rids []int32, token string) (err error) {
	tx := ds.db.Begin()
	deleteUdr := UserTenantRole{Uid: uid, Tid: tid}

	err = deleteUdr.DeleteUserRoleByUidAndTid(tx)
	if err != nil {
		tx.Rollback()
		return
	}
	rids = append(rids, 0)
	for _, v := range rids {
		createUdr := UserTenantRole{Uid: uid, Tid: tid, Rid: v, IsDefault: true}
		err = createUdr.Create(tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	return
}

/*UpdateUserRole
修改用户在域中的角色
*/
func (ds *DrbacServer) UpdateUserRoleInDomain(uid, did int32, rids []int32, token string) (err error) {
	tx := ds.db.Begin()
	deleteUdr := UserDomainRole{Uid: uid, Did: did}

	err = deleteUdr.DeleteUserRoleByUidAndDid(tx)
	if err != nil {
		tx.Rollback()
		return
	}
	rids = append(rids, 0)
	for _, v := range rids {
		createUdr := UserDomainRole{Uid: uid, Did: did, Rid: v}
		err = createUdr.Create(tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	return
}

/*UpdatePassword
修改用户认证密码
*/
func (ds *DrbacServer) UpdatePassword(uid int32, password string) (err error) {
	user := User{Uid: uid, Password: password}
	err = user.UpdatePassword(ds.db)
	return
}

/*UpdatePassword
修改用户认证密码
*/
func (ds *DrbacServer) UpdateNicknameAndPassword(token string, uid int32, nickname, password string) (err error) {
	tx := ds.db.Begin()
	userInfo, err := ds.GetUserByUid(uid)
	if userInfo.Username == "" || err != nil {
		log.Info("Cannot Find User By Uid ", uid)
		return err
	}
	user := User{Uid: uid, Nickname: nickname}
	user.Salt = string(Krand(6, KC_RAND_KIND_UPPER))
	user.Password = EncryptWithSalt(password, user.Salt)
	err = user.UpdateNicknameAndPassword(tx)
	if err != nil {
		log.Info("UpdateNicknameAndPassword Error,", err)
		tx.Rollback()
	}
	//todo 刷新缓存
	if userInfo.Did != 0 {
		_, err = reAuthenticationByDid(token, userInfo.Did, ds.db, ds.pool)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else if userInfo.Tid != 0 {
		_, err = reAuthentication(token, userInfo.Tid, ds.db, ds.pool)
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		tx.Rollback()
		log.Info("Cannot Find Did or Tid By Uid:", uid)
		return errors.New("cannot Find Did or Tid By Uid")
	}
	tx.Commit()
	return
}

/*UpdateNickname
修改用户昵称
*/
func (ds *DrbacServer) UpdateNickname(uid int32, nickname, token string) (err error) {
	err = Transaction(ds.db, func(tx *gorm.DB) error {
		userInfo, err := ds.GetUserByUid(uid)
		if userInfo.Username == "" || err != nil {
			log.Info("Cannot Find User By Uid ", uid)
			return err
		}
		user := User{Uid: uid, Nickname: nickname}
		err = user.UpdateNickname(ds.db)
		if err != nil {
			return err
		}
		//todo 刷新缓存
		if userInfo.Did != 0 {
			_, err = reAuthenticationByDid(token, userInfo.Did, ds.db, ds.pool)
			if err != nil {
				return err
			}
		} else if userInfo.Tid != 0 {
			_, err = reAuthentication(token, userInfo.Tid, ds.db, ds.pool)
			if err != nil {
				return err
			}
		} else {
			log.Info("Cannot Find Did or Tid By Uid:", uid)
			return errors.New("cannot Find Did or Tid By Uid")
		}
		return nil
	})
	return
}

//修改用户状态
func (ds *DrbacServer) UpdateUserState(uid int32, state int32) (ra int64, err error) {
	tx := ds.db.Begin()
	user := User{Uid: uid, State: state}
	ra, err = user.UpdateUser(tx)
	if ra == 0 || err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

/*GetUserByUid
基于uid获取username
*/
func (ds *DrbacServer) GetUserByUid(uid int32) (user *User, err error) {
	user = &User{Uid: uid}
	err = user.GetUserByUID(ds.db)
	return
}

/*GetUserByUsername
根据用户名查询用户
*/
func (ds *DrbacServer) GetUserByUsernameAndDid(username string, did int32) (user *User, err error) {
	user = &User{Username: username, Did: did}
	err = user.GetUserByUsernameAndDid(ds.db)
	return
}

/*GetUserByUsername
根据用户名查询用户
*/
func (ds *DrbacServer) GetUserByUsernameAndTid(username string, tid int32) (user *User, err error) {
	user = &User{Username: username, Tid: tid}
	err = user.GetUserByUsernameAndTid(ds.db)
	return
}

/*GetTokenByUid
基于uid湖片区用户对应的token
*/
func (ds *DrbacServer) GetTokenByUidAndTid(uid, tid int32) (token string) {
	token = getTokenByUidAndTid(uid, tid, ds.db, ds.pool)
	return
}

/*GetTokenByUsername
基于username获取用户对应的token
*/
func (ds *DrbacServer) GetTokenByUsernameAndTenantName(username, tenantName string) (token string) {
	token = getTokenByUsernameAndTenantName(username, tenantName, ds.pool)
	return
}

/*GetTokenByUsername
基于username获取用户对应的token
*/
func (ds *DrbacServer) GetTokenByUsernameAndTid(username string, tid int32) (token string) {
	tenant, err := ds.GetTenant(tid)
	if err != nil {
		return
	}
	token = getTokenByUsernameAndTenantName(username, tenant.TenantName, ds.pool)
	return
}

func (ds *DrbacServer) AddRole(roleName, description string, modules []*Modules, tid int32) (rid int32, err error) {
	role := Role{RoleName: roleName, Description: description, Tid: tid, CreateTime: time.Now(), UpdateTime: time.Now()}
	tx := ds.db.Begin()
	err = role.Create(tx)
	if err != nil || role.Rid == 0 {
		tx.Rollback()
		return
	}
	//todo 通过module获取permission
	var perIds []int32
	for _, v := range modules {
		if v.Mid != 0 {
			ids, err := GetPidsByModule(v.Mid, v.Operations, tx)
			if err != nil {
				log.Error("GetPidsByModule Error,", err)
				tx.Rollback()
				return 0, err
			}
			log.Info("ids:", ids)
			for _, v := range ids {
				perIds = append(perIds, v)
			}
		}
	}
	log.Info("len(perIds):", len(perIds))
	// todo role_permission
	for _, v := range perIds {
		rp := RolePermission{
			Rid:        role.Rid,
			Pid:        v,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		err = rp.Create(tx)
		if err != nil {
			log.Error("role_permission Create Error, ", err)
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	rid = role.Rid
	log.Info("Create Role Successful, Rid:", rid)
	return
}

func (ds *DrbacServer) AddDomainRole(roleName, description string, modules []*Modules, did int32) (rid int32, err error) {
	role := Role{RoleName: roleName, Description: description, Did: did, CreateTime: time.Now(), UpdateTime: time.Now()}
	tx := ds.db.Begin()
	err = role.Create(tx)
	if err != nil || role.Rid == 0 {
		tx.Rollback()
		return
	}
	//todo 通过module获取permission
	var perIds []int32
	for _, v := range modules {
		if v.Mid != 0 {
			ids, err := GetPidsByModule(v.Mid, v.Operations, tx)
			if err != nil {
				log.Error("GetPidsByModule Error,", err)
				tx.Rollback()
				return 0, err
			}
			for _, v := range ids {
				perIds = append(perIds, v)
			}
		}
	}
	// todo role_permission
	for _, v := range perIds {
		rp := RolePermission{
			Rid:        role.Rid,
			Pid:        v,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		err = rp.Create(tx)
		if err != nil {
			log.Error("role_permission Create Error, ", err)
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	rid = role.Rid
	log.Info("Create Role Successful, Rid:", rid)
	return
}

func (ds *DrbacServer) UpdateRoleInfo(rid int32, roleName, description string, modules []*Modules) (err error) {
	tx := ds.db.Begin()
	//todo 修改roleInfo
	role := Role{Rid: rid, RoleName: roleName, Description: description, UpdateTime: time.Now()}
	err = role.UpdateRoleByRid(ds.db)
	if err != nil {
		log.Error("role.UpdateRoleByRid Error,", err)
		tx.Rollback()
		return
	}
	//todo 删除旧rp关系
	rp := RolePermission{Rid: rid}
	err = rp.DeleteByRid(tx)
	if err != nil {
		log.Error("rp.DeleteByRid Error", err)
		tx.Rollback()
		return
	}
	//todo 插入新rp关系
	//todo 通过module获取permission
	var perIds []int32
	for _, v := range modules {
		if v.Mid != 0 {
			ids, err := GetPidsByModule(v.Mid, v.Operations, tx)
			if err != nil {
				log.Error("GetPidsByModule Error,", err)
				tx.Rollback()
				return err
			}
			for _, v := range ids {
				perIds = append(perIds, v)
			}
		}
	}
	// todo role_permission
	for _, v := range perIds {
		rp := RolePermission{
			Rid:        role.Rid,
			Pid:        v,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		err = rp.Create(tx)
		if err != nil {
			log.Error("role_permission Create Error, ", err)
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	return
}

func (ds *DrbacServer) DeleteTenantRole(rid int32) (err error) {
	tx := ds.db.Begin()
	//todo 删除角色
	role := Role{Rid: rid}
	err = role.DeleteByRID(tx)
	if err != nil {
		tx.Rollback()
		return
	}
	//todo 删除角色权限绑定
	rr := RolePermission{Rid: rid}
	err = rr.DeleteByRid(tx)
	if err != nil {
		tx.Rollback()
		return
	}
	//todo 删除用户角色绑定
	udr := UserTenantRole{Rid: rid}
	err = udr.DeleteUserRole(tx)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}
func (ds *DrbacServer) DeleteDomainRole(rid int32) (err error) {
	tx := ds.db.Begin()
	//todo 删除角色
	role := Role{Rid: rid}
	err = role.DeleteByRID(tx)
	if err != nil {
		tx.Rollback()
		return
	}
	//todo 删除角色权限绑定
	rr := RolePermission{Rid: rid}
	err = rr.DeleteByRid(tx)
	if err != nil {
		tx.Rollback()
		return
	}
	//todo 删除用户角色绑定
	udr := UserDomainRole{Rid: rid}
	err = udr.DeleteUserRole(tx)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func (ds *DrbacServer) GetModulesByRid(rid int32, moduleAll []*ModuleInfo) (modules []*Modules, err error) {
	rp := RolePermission{Rid: rid}
	permission, err := rp.GetModulesByRid(ds.db)
	if err != nil {
		log.Error("rp.GetModulesByRid Error, ", err)
		return
	}
	for _, v1 := range moduleAll {
		m := &Modules{Mid: v1.Mid, ModuleName: v1.ModuleName}
		for _, v2 := range permission {
			if m.Mid == v2.Mid {
				m.Operations = append(m.Operations, v2.PerOperation)
			}
		}
		modules = append(modules, m)
	}
	return
}

func RemoveStringDuplicatesAndEmpty(list []string) (set []string) {
	for _, i := range list {
		if i != "" {
			if len(set) == 0 {
				set = append(set, i)
			} else {
				for k, v := range set {
					if i == v {
						break
					}
					if k == len(set)-1 {
						set = append(set, i)
					}
				}
			}
		}
	}
	return
}

/*AddUser
创建被邀请的用户
*/
func (ds *DrbacServer) AddUser(username, nickname, token string, tid int32, rids []int32) (err error) {
	tx := ds.db.Begin()
	user := &User{Username: username, Nickname: nickname, Token: token, Tid: tid}
	err = user.Create(tx)
	if err != nil {
		log.Error("user.Create Error,", err)
		tx.Rollback()
		return
	}
	log.Info("AddUser UserInfo:", user)
	rids = append(rids, 0)
	for _, rid := range rids {
		udr := &UserTenantRole{Uid: user.Uid, Tid: tid, Rid: rid, IsDefault: true}
		err = udr.Create(tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	return
}

/*AddUser
创建被邀请的用户
*/
func (ds *DrbacServer) AddDomainUser(username, nickname, token string, did int32, rids []int32) (err error) {
	tx := ds.db.Begin()
	user := &User{Username: username, Nickname: nickname, Token: token, Did: did, CreateTime:time.Now()}
	err = user.Create(tx)
	if err != nil {
		tx.Rollback()
		return
	}
	log.Info("AddUser UserInfo:", user)
	rids = append(rids, 0)
	for _, rid := range rids {
		udr := &UserDomainRole{Uid: user.Uid, Did: did, Rid: rid, CreateTime:time.Now(), UpdateTime:time.Now()}
		err = udr.Create(tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	return
}

/*AddUser
邀请已存在的用户进入新的租户
*/
func (ds *DrbacServer) AddUdr(uid, tid int32, rids []int32) (err error) {
	tx := ds.db.Begin()
	rids = append(rids, 0)
	for _, rid := range rids {
		udr := &UserTenantRole{Uid: uid, Tid: tid, Rid: rid, IsDefault: true}
		err = udr.Create(tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	return
}

/*AddUser
邀请已存在的用户进入域
*/
func (ds *DrbacServer) AddUserInDomain(uid, did int32, rids []int32) (err error) {
	tx := ds.db.Begin()
	rids = append(rids, 0)
	for _, rid := range rids {
		udr := &UserDomainRole{Uid: uid, Did: did, Rid: rid}
		err = udr.Create(tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	return
}

/*
根绝Token查询被邀请的用户
*/
func (ds *DrbacServer) GetUserByToken(token string) (user User, err error) {
	user = User{Token: token}
	err = user.GetUserByToken(ds.db)
	return
}

/*
完善用户信息
*/
func (ds *DrbacServer) UpdateUser(uid int32, password string) (ra int64, err error) {
	user := User{Uid: uid}
	user.Salt = string(Krand(6, KC_RAND_KIND_ALL))
	log.Info("u.Salt:", user.Salt)
	user.Password = EncryptWithSalt(password, user.Salt)
	user.State = 3
	ra, err = user.UpdateUser(ds.db)
	return
}

/*DeleteUser
删除用户
*/
func (ds *DrbacServer) DeleteUser(uid, tid int32) (err error) {
	tx := ds.db.Begin()
	udr := UserTenantRole{Uid: uid, Tid: tid}
	err = udr.DeleteByUidAndTid(tx)
	if err != nil {
		tx.Rollback()
		return
	}
	utr1 := UserTenantRole{Uid: uid}
	roles, _ := utr1.GetRolesByUid(tx)
	if len(roles) != 0 {
		tx.Commit()
		return
	}
	udr1 := UserDomainRole{Uid: uid}
	rolesDomain, _ := udr1.GetRolesByUid(tx)
	if len(rolesDomain) != 0 {
		tx.Commit()
		return
	}
	user := User{Uid: uid}
	err = user.DeleteByUID(tx)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

/*DeleteUser
删除用户
*/
func (ds *DrbacServer) DeleteUserInDomain(uid, did int32) (err error) {
	tx := ds.db.Begin()
	udr := UserDomainRole{Uid: uid, Did: did}
	err = udr.DeleteByUidAndDid(tx)
	if err != nil {
		tx.Rollback()
		return
	}
	udr1 := UserDomainRole{Uid: uid}
	roles, _ := udr1.GetRolesByUid(tx)
	if len(roles) != 0 {
		tx.Commit()
		return
	}
	utr1 := UserTenantRole{Uid: uid}
	rolesTenant, _ := utr1.GetRolesByUid(tx)
	if len(rolesTenant) != 0 {
		tx.Commit()
		return
	}

	user := User{Uid: uid}
	err = user.DeleteByUID(tx)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func (ds *DrbacServer) IsExistedInTenant(uid, tid int32) (isTrue bool) {
	udr := UserTenantRole{Uid: uid, Tid: tid}
	isTrue = udr.IsExistedInTenant(ds.db)
	return
}

func (ds *DrbacServer) IsExistedInDomain(uid, did int32) (isTrue bool) {
	udr := UserDomainRole{Uid: uid, Did: did}
	isTrue = udr.IsExistedInDomain(ds.db)
	return
}

//校验URL白名单
func (ds *DrbacServer) CheckWhitelist(url, opt string) bool {
	return checkWhitelist(url, opt, ds.db)
}

func RemoveDuplicatesAndEmpty(list []int32) (set []int32) {
	for _, i := range list {
		if i != 0 {
			if len(set) == 0 {
				set = append(set, i)
			} else {
				for k, v := range set {
					if i == v {
						break
					}
					if k == len(set)-1 {
						set = append(set, i)
					}
				}
			}
		}
	}
	return
}

/*
通过URL查询Tid
*/

func (ds *DrbacServer) GetTidByUrl(url string) (*TenantInfo, error) {
	tenant := Tenant{TenantURL: url}
	err := tenant.GetTidByUrl(ds.db)
	if err != nil || tenant.Tid == 0 {
		log.Error("GetTidByUrl Error,", err)
		return nil, errors.New("GetTidByUrl From Mysql Error")
	}
	return &TenantInfo{Tid: tenant.Tid, Icon: tenant.Icon, Logo: tenant.Logo}, nil
}

/*
通过Tid查询Did
*/

func (ds *DrbacServer) GetDidByTid(tid int32) (int32, error) {
	tenant := Tenant{Tid: tid}
	err := tenant.GetDidByTid(ds.db)
	if err != nil || tenant.Did == 0 {
		log.Error("GetTidByUrl Error,", err, " ,did:", tenant.Did)
		return 0, errors.New("GetDidByTid From Mysql Error")
	}
	return tenant.Did, nil
}

/*
查询全部权限模块信息
*/
func (ds *DrbacServer) GetModulesByTid(tid int32) (modules []*ModuleInfo, err error) {
	m := Module{ModuleTid: tid}
	reply, err := m.GetModulesByTid(ds.db)
	if err != nil {
		log.Error("GetModules Error,", err)
		return
	}
	for _, v := range reply {
		modules = append(modules, &ModuleInfo{v.Mid, v.ModuleName})
	}
	return
}

func (ds *DrbacServer) GetModulesByDid(did int32) (modules []*ModuleInfo, err error) {
	m := Module{ModuleDid: did}
	reply, err := m.GetModulesByDid(ds.db)
	if err != nil {
		log.Error("GetModules Error,", err)
		return
	}
	for _, v := range reply {
		modules = append(modules, &ModuleInfo{v.Mid, v.ModuleName})
	}
	return
}

func (ds *DrbacServer) AddUserTenantACL(username string, did int32, tids []int32) (err error) {
	tx := ds.db.Begin()
	user := User{Username: username, Did: did}
	err = user.GetUserByUsernameAndDid(tx)
	if err != nil || user.Uid == 0 {
		log.Error("GetUserByUsernameAndDid Error,", err, " ,Uid:", user.Uid)
		tx.Rollback()
		return
	}
	for _, v := range tids {
		uta := UserTenantACL{Uid: user.Uid, Tid: v}
		err = uta.Create(tx)
		if err != nil {
			log.Error("tid:", v, ",uta.Create Error,", err)
			tx.Rollback()
			return
		}
		log.Info("tid:", v, ",插入acl关系成功")
	}
	tx.Commit()
	return
}

func (ds *DrbacServer) GetUserTenantACL(uid int32) (tids []int32, err error) {
	uta := UserTenantACL{Uid: uid}
	tids, err = uta.GetTidsByUid(ds.db)
	return
}

func (ds *DrbacServer) AddService(sid, serviceType int32, serviceName, serviceKey, serviceUrl string, serviceTid int32) (err error) {
	tx := ds.db.Begin()
	service := Service{Sid: sid, ServiceName: serviceName, ServiceUrl: serviceUrl, ServiceType: serviceType, ServiceKey: serviceKey, ServiceState: 1}
	err = service.Create(tx)
	if err != nil || service.Sid == 0 {
		log.Error("service.Create Error")
		tx.Rollback()
		return
	}
	tsp := &TenantServicePolicy{Tid: serviceTid, Sid: service.Sid, Pid: 3}
	err = tsp.Create(tx)
	if err != nil {
		log.Error("tsp.Create Error")
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func (ds *DrbacServer) UpdateService(sid, serviceType int32, serviceName, serviceKey, serviceUrl string, serviceTid int32, serviceDescription string, serviceState int32) (err error) {
	tx := ds.db.Begin()
	if serviceTid != 0 {
		if serviceState != 0 {
			tsp := TenantServicePolicy{Sid:sid,State:serviceState,Tid:serviceTid}
			err = tsp.UpdateBySidAndTid(tx)
			if err != nil {
				log.Error("TenantServicePolicy UpdateBySid Err,", err)
				tx.Rollback()
				return
			}
		}else {
			service := Service{Sid: sid, ServiceName: serviceName, ServiceUrl: serviceUrl, ServiceType: serviceType, ServiceKey: serviceKey, ServiceDescription: serviceDescription}
			err = service.Update(tx)
			if err != nil {
				log.Error("UpdateService Err,", err)
				tx.Rollback()
				return
			}
		}
	}else {
		service := Service{Sid: sid, ServiceName: serviceName, ServiceUrl: serviceUrl, ServiceType: serviceType, ServiceKey: serviceKey, ServiceDescription: serviceDescription, ServiceState:serviceState}
		err = service.Update(tx)
		if err != nil {
			log.Error("UpdateService Err,", err)
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	return
}

func (ds *DrbacServer) DeleteService(sid int32) (err error) {
	tx := ds.db.Begin()
	tsp := &TenantServicePolicy{Sid: sid}
	err = tsp.DeleteBySid(tx)
	if err != nil {
		log.Error("tsp.DeleteBySid Error")
		tx.Rollback()
		return
	}
	policy := Policy{PolicySid: sid}
	err = policy.DeleteBySid(tx)
	if err != nil {
		log.Error("policy.DeleteBySid Error")
		tx.Rollback()
		return
	}
	service := Service{Sid: sid}
	err = service.DeleteBySid(tx)
	if err != nil {
		log.Error("service.DeleteBySid Error")
		tx.Rollback()
		return
	}
	log.Info("Sid:", sid, " 删除服务成功")
	tx.Commit()
	return
}

func (ds *DrbacServer) GetServiceBySid(sid int32) (service Service, err error) {
	service = Service{Sid: sid}
	err = service.GetBySid(ds.db)
	return
}

func (ds *DrbacServer) GetServiceByTid(tid int32) (services []*Service, err error) {
	service := Service{}
	services, err = service.GetByTid(ds.db, tid)
	return
}

func (ds *DrbacServer) GetTenantAccountByTid(tid int32) (account TenantAccount, err error) {
	account = TenantAccount{Tid: tid}
	err = account.GetTenantAccount(ds.db)
	return
}

func (ds *DrbacServer) UpdateTenantAccountByTid(tid int32, price float32, actionType int32) (err error) {
	account := TenantAccount{Tid: tid}
	err = account.GetTenantAccount(ds.db)
	log.Info("Old Balance:", account.Balance)
	if actionType == 1 {
		account.Balance += price
		log.Info("New Balance:", account.Balance)
		err = account.UpdateTenantAccount(ds.db)
		if err == nil {
			ds.AddTradingRecordsByTid(tid, "账户充值", price, price, 1, 3)
		}
		return
	} else if actionType == 2 {
		if account.Balance < 1 {
			log.Info("账户余额不足，扣费失败")
			return errors.New("账户余额不足，扣费失败")
		}
		account.Balance -= price
		log.Info("New Balance:", account.Balance)
		err = account.UpdateTenantAccount(ds.db)
		if err == nil {
			ds.AddTradingRecordsByTid(tid, "设备接入", price, price, 1, 3)
		}
		return
	}
	return errors.New("系统异常")
}

func (ds *DrbacServer) GetTradingRecordsByTid(tid int32, page, count int32) (records []*TradingRecord, totalCount int32, err error) {
	record := TradingRecord{Tid: tid}
	records, totalCount, err = record.GetTradingRecords(ds.db, page, count)
	return
}

func (ds *DrbacServer) AddTradingRecordsByTid(tid int32, tradingContent string, tradingUnitPrice, tradingTotalPrice float32, tradingCount, TradingState int32) (err error) {
	record := TradingRecord{Tid: tid, CreateTime: time.Now(), TradingContent: tradingContent, TradingUnitPrice: tradingUnitPrice, TradingCount: tradingCount, TradingState: TradingState, TradingTotalPrice: tradingTotalPrice}
	err = record.AddTradingRecord(ds.db)
	return
}

func (ds *DrbacServer) AddPolicy(policyType, policyCycle, policyFeeType, policyUnitType, policyUnitCount, policySid int32, policyName string, policyUnitPrice float32) (err error) {
	policy := Policy{PolicyType: policyType, PolicyCycle: policyCycle, PolicyFeeType: policyFeeType, PolicyUnitType: policyUnitType, PolicyUnitCount: policyUnitCount, PolicySid: policySid, PolicyName: policyName, PolicyUnitPrice: policyUnitPrice, CreateTime: time.Now(), UpdateTime: time.Now()}
	err = policy.Create(ds.db)
	return
}

func (ds *DrbacServer) UpdatePolicy(pid, policyType, policyCycle, policyFeeType, policyUnitType, policyUnitCount, policySid int32, policyName string, policyUnitPrice float32) (err error) {
	policy := Policy{Pid: pid, PolicyType: policyType, PolicyCycle: policyCycle, PolicyFeeType: policyFeeType, PolicyUnitType: policyUnitType, PolicyUnitCount: policyUnitCount, PolicySid: policySid, PolicyName: policyName, PolicyUnitPrice: policyUnitPrice, UpdateTime: time.Now()}
	err = policy.Update(ds.db)
	return
}

func (ds *DrbacServer) DeletePolicyByPid(pid int32) (err error) {
	policy := Policy{Pid: pid}
	err = policy.DeleteByPid(ds.db)
	return
}

func (ds *DrbacServer) DeletePolicyBySid(sid int32) (err error) {
	policy := &Policy{PolicySid: sid}
	err = policy.DeleteBySid(ds.db)
	return
}

func (ds *DrbacServer) GetPolicyByPid(pid int32) (policy Policy, err error) {
	policy = Policy{Pid: pid}
	err = policy.GetByPid(ds.db)
	return
}

func (ds *DrbacServer) GetPolicyBySid(sid int32) (policys []*Policy, err error) {
	policy := Policy{PolicySid: sid}
	policys, err = policy.GetBySid(ds.db)
	return
}

func (ds *DrbacServer) AddAliPay(did int32, merchantPrivateKey, key, appId string) (err error) {
	aliPay := AliPay{Did: did, AppId: appId, MerchantPrivateKey: merchantPrivateKey, Key: key, CreateTime: time.Now(), UpdateTime: time.Now()}
	err = aliPay.Create(ds.db)
	return
}

func (ds *DrbacServer) UpdateAliPay(did int32, merchantPrivateKey, key, appId string) (err error) {
	aliPay := AliPay{Did: did, AppId: appId, MerchantPrivateKey: merchantPrivateKey, Key: key, UpdateTime: time.Now()}
	err = aliPay.Update(ds.db)
	return
}

func (ds *DrbacServer) DeleteAliPay(did int32) (err error) {
	aliPay := AliPay{Did: did}
	err = aliPay.DeleteByDid(ds.db)
	return
}

func (ds *DrbacServer) GetAliPay(did int32) (aliPay *AliPay, err error) {
	aliPay = &AliPay{Did: did}
	err = aliPay.GetByDid(ds.db)
	return
}

func (ds *DrbacServer) AddWechatPay(did int32, key, appId, mchId, appSecret string) (err error) {
	wechatPay := WeChatPay{Did: did, AppId: appId, MchId: mchId, Key: key, AppSecret: appSecret, CreateTime: time.Now(), UpdateTime: time.Now()}
	err = wechatPay.Create(ds.db)
	return
}

func (ds *DrbacServer) UpdateWechatPay(did int32, key, appId, mchId, appSecret string) (err error) {
	wechatPay := WeChatPay{Did: did, AppId: appId, MchId: mchId, Key: key, AppSecret: appSecret, UpdateTime: time.Now()}
	err = wechatPay.Update(ds.db)
	return
}

func (ds *DrbacServer) DeleteWechatPay(did int32) (err error) {
	wechatPay := WeChatPay{Did: did}
	err = wechatPay.DeleteByDid(ds.db)
	return
}

func (ds *DrbacServer) GetWechatPay(did int32) (wechatPay *WeChatPay, err error) {
	wechatPay = &WeChatPay{Did: did}
	err = wechatPay.GetByDid(ds.db)
	return
}

func (ds *DrbacServer) GetDomainByDid(did int32) (domain *Domain, err error) {
	domain = &Domain{Did: did}
	err = domain.GetByID(ds.db)
	return
}


func (ds *DrbacServer) GetDomains() ([]*Domain,error) {
	domain := &Domain{}
	domains, err := domain.GetDomains(ds.db)
	return domains,err
}


func (ds *DrbacServer) AddModuleByDid(moduleName string, did int32, admin, read ,update ,create []int32) error {
	//todo 创建模块信息
	tx := ds.db.Begin()
	module := &Module{}
	err := module.Create(tx)
	if err != nil || module.Mid == 0{
		log.Info("Create Module Error,",err)
		tx.Rollback()
		return err
	}
	//todo 创建permission
	adminPermission := &Permission{PerName:moduleName + "_Admin", PerOperation:"ADMIN"}
	err = adminPermission.Create(tx)
	if err != nil || adminPermission.Pid == 0{
		log.Info("Create Admin_Permission Error,",err)
		tx.Rollback()
		return err
	}

	readPermission := &Permission{PerName:moduleName + "_Read", PerOperation:"READ"}
	err = readPermission.Create(tx)
	if err != nil || readPermission.Pid == 0{
		log.Info("Create Read_Permission Error,",err)
		tx.Rollback()
		return err
	}

	updatePermission := &Permission{PerName:moduleName + "_Update", PerOperation:"UPDATE"}
	err = updatePermission.Create(tx)
	if err != nil || updatePermission.Pid == 0{
		log.Info("Create Update_Permission Error,",err)
		tx.Rollback()
		return err
	}

	createPermission := &Permission{PerName:moduleName + "_Create", PerOperation:"CREATE"}
	err = createPermission.Create(tx)
	if err != nil || createPermission.Pid == 0 {
		log.Info("Create Create_Permission Error,",err)
		tx.Rollback()
		return err
	}
	//todo 创建module_permission
	adminModulePermission := ModulePermission{Mid:module.Mid, Pid:adminPermission.Pid}
	err = adminModulePermission.Create(tx)
	if err != nil {
		log.Info("Create Admin_ModulePermission Error,",err)
		tx.Rollback()
		return err
	}

	readModulePermission := ModulePermission{Mid:module.Mid, Pid:readPermission.Pid}
	err = readModulePermission.Create(tx)
	if err != nil {
		log.Info("Create Read_ModulePermission Error,",err)
		tx.Rollback()
		return err
	}

	updateModulePermission := ModulePermission{Mid:module.Mid, Pid:updatePermission.Pid}
	err = updateModulePermission.Create(tx)
	if err != nil {
		log.Info("Create Update_ModulePermission Error,",err)
		tx.Rollback()
		return err
	}

	createModulePermission := ModulePermission{Mid:module.Mid, Pid:createPermission.Pid}
	err = createModulePermission.Create(tx)
	if err != nil {
		log.Info("Create Create_ModulePermission Error,",err)
		tx.Rollback()
		return err
	}

	//todo 创建permission_resource
	//permissionResource := &permissionress

	return err
}

func (ds *DrbacServer) AddModuleByTid() error {

	return nil
}