package drbac

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
	. "auth/go-drbac/internal/storage"
	log "github.com/cihub/seelog"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"errors"
)

type UserToken struct {
	Token                 string                  `json:"Token"`
	User                  *User                   `json:"User"`
	TenantRoleResource   []*TenantRoleResource    `json:"TenantRoleResource,omitempty"`
	DomainRoleResource    *DomainRoleResource     `json:"DomainRoleResource,omitempty"`
}
type UserTokenTenants struct {
	Token                 string                  `json:"Token"`
	User                  *User                   `json:"User"`
	TenantRoleResource []*TenantRoleResource      `json:"TenantRoleResource"`
}

/*
 *	认证
 */
func authenticationWithTid(username, password string, tid int32, tx *gorm.DB, pool *redis.Pool) (userTokenTenants *UserTokenTenants, isTrue bool, state int32, err error) {
	u := &User{Username: username, Tid: tid}
	err = u.GetUserByUsernameAndTid(tx)
	if err != nil || u.Uid == 0 {
		log.Errorf("username: %s does not exist, err is %s", username, err)
		err = errors.New("用户不存在")
		return
	}
	state = u.State
	isTrue, err = u.IsTruePassword(password, tx)
	if err != nil {
		log.Errorf("password is error,username %s , password %s, err is %s", err, username, password)
		return
	}
	if !isTrue {
		log.Errorf("password is error,username %s , password %s", username, password)
		return
	}
	drps, err := getTenantRoleResourceByUidAndTid(u.Uid, tid, tx)
	if err != nil || drps == nil {
		log.Errorf("GetDomainRolePermissionByUid is error %s,u.Uid is %d, Did is %d ", err, u.Uid, tid)
		err = errors.New("用户不属于该域")
		return
	}
	token := generateTokenWithTid(u.Uid, tid)
	userTokenTenants = &UserTokenTenants{User: u, TenantRoleResource: drps, Token: token}
	err = saveUserTokenTenants(userTokenTenants, token, pool)
	if err != nil {
		log.Errorf("saveToken error %s", err)
		return
	}
	return
}

/*
 *	认证
 */
func authenticationWithDid(username, password string, did int32, tx *gorm.DB, pool *redis.Pool) (userToken *UserToken, isTrue bool, state int32, err error) {
	log.Info("Did:",did)
	u := &User{Username: username, Did: did}
	err = u.GetUserByUsernameAndDid(tx)
	if err != nil {
		log.Errorf("username: %s does not exist, err is %s", username, err)
		return
	}
	state = u.State
	isTrue, err = u.IsTruePassword(password, tx)
	if err != nil {
		log.Errorf("password is error,username %s , password %s, err is %s", err, username, password)
		return
	}
	if !isTrue {
		log.Errorf("password is error,username %s , password %s", username, password)
		return
	}
	trps, err := getDomainRolePermissionByUidAndDid(u.Uid, did, tx)
	log.Info("111111111111111111111111111111111, trps:", trps, " err:", err)
	if err != nil || trps == nil {
		log.Errorf("GetDomainRolePermissionByUid is error (%s),u.Uid is (%d)", err, u.Uid)
		return
	}
	//todo 通过uid获取tid
	var trrs []*TenantRoleResource
	if did == 100002 || did == 100003 || did == 100004 {
		//todo 查acl
		log.Info("开始查询域下租户")
		uta := UserTenantACL{Uid:u.Uid}
		tids,err := uta.GetTidsByUid(tx)
		if err != nil {
			log.Error("uta.GetTidsByUid Error")
			return nil, isTrue, state, err
		}
		log.Info("tids:",tids)
		for _,v := range tids {
			trr,err := getTenantRoleResourceByUidAndTid(u.Uid, v, tx)
			if err != nil {
				log.Error("tid:",v,",getTenantRoleResourceByUidAndTid Error,",err)
				return nil, isTrue, state, err
			}
			trrs = append(trrs, trr[0])
		}
	}else {
		//todo 查所有租户
	}
	token := generateToken(u.Uid)
	userToken = &UserToken{User: u, DomainRoleResource: trps, Token: token, TenantRoleResource:trrs}
	log.Info("userTokenDomain:", userToken)
	err = saveUserTokenDomain(userToken, token, pool)
	if err != nil {
		log.Errorf("saveUserTokenDomain error %s", err)
		return
	}
	return
}

/*
 *	认证
 */
func authentication(uid int32, password string, tx *gorm.DB, pool *redis.Pool) (isTrue bool, state int32, err error) {
	u := &User{Uid: uid}
	err = u.GetUserByUID(tx)
	if err != nil {
		log.Errorf("uid: %s does not exist, err is %s", uid, err)
		return
	}
	state = u.State
	isTrue, err = u.IsTruePassword(password, tx)
	if err != nil {
		log.Errorf("password is error,uid %s , password %s, err is %s", err, uid, password)
		return
	}
	if !isTrue {
		log.Errorf("password is error,uid %s , password %s", uid, password)
		return
	}
	return
}

/*
 *	重认证，修改redis中的认证信息（刷新缓存）
 */
func reAuthentication(token string, tid int32, tx *gorm.DB, pool *redis.Pool) (userTokenTenants *UserTokenTenants, err error) {
	oldUserToken := getToken(token, pool)
	if oldUserToken == nil {
		log.Errorf("getToken error %s token is %s", err, token)
		return
	}
	u := oldUserToken.User
	err = u.GetUserByUID(tx)
	if err != nil {
		log.Errorf("GetUserByUID error %s uid is %d", err, u.Uid)
		return
	}
	drps, err := getTenantRoleResourceByUidAndTid(u.Uid, tid, tx)
	if err != nil {
		log.Errorf("GetDomainRolePermissionByUid error %s uid is %d", err, u.Uid)
		return
	}
	userTokenTenants = &UserTokenTenants{User: u, TenantRoleResource: drps, Token: token}
	err = saveUserTokenTenants(userTokenTenants, token, pool)
	if err != nil {
		log.Errorf("saveToken error %s", err)
		return
	}
	return
}


func reAuthenticationByDid(token string, did int32, tx *gorm.DB, pool *redis.Pool) (userTokenTenants *UserToken, err error) {
	oldUserToken := getToken(token, pool)
	if oldUserToken == nil {
		log.Errorf("getToken error %s token is %s", err, token)
		return
	}
	u := oldUserToken.User
	err = u.GetUserByUID(tx)
	if err != nil {
		log.Errorf("GetUserByUID error %s uid is %d", err, u.Uid)
		return
	}
	drps, err := getDomainRolePermissionByUidAndDid(u.Uid, did, tx)
	if err != nil {
		log.Errorf("GetDomainRolePermissionByUid error %s uid is %d", err, u.Uid)
		return
	}
	userTokenTenants = &UserToken{User: u, DomainRoleResource: drps, Token: token}
	err = saveUserTokenDomain(userTokenTenants, token, pool)
	if err != nil {
		log.Errorf("saveToken error %s", err)
		return
	}
	return
}

/*
	基于uid重新认证用户信息（刷新缓存）
*/
func reAuthenticationByUid(token string, tid int32, tx *gorm.DB, pool *redis.Pool,
	reAuthentication func(token string, did int32, tx *gorm.DB, pool *redis.Pool) (userTokenTenants *UserTokenTenants, err error)) (userTokenTenants *UserTokenTenants, err error) {
	return reAuthentication(token, tid, tx, pool)
}

/*
	生成token值
*/
func generateToken(base int32) string {
	rand.Seed(time.Now().Unix())
	s := strconv.Itoa(int(base)) + strconv.Itoa(rand.Int())
	session := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", session)
}

/*
	生成token值 with tid
*/
func generateTokenWithTid(base1, base2 int32) string {
	rand.Seed(time.Now().Unix())
	s := strconv.Itoa(int(base1)) + strconv.Itoa(int(base2)) + strconv.Itoa(rand.Int())
	session := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", session)
}

/*
	保存token
*/
//func saveToken(userTokenTenants *UserTokenTenants, token string, pool *redis.Pool) (err error) {
//	if userTokenTenants == nil {
//		err = errors.New("userToken is nil")
//		return
//	}
//	if userTokenTenants.User.Username == ""{
//		log.Info("usertoken:",userTokenTenants)
//		err = errors.New("userToken is nil")
//		return
//	}
//	client := pool.Get()
//	defer client.Close()
//	resp, err := json.Marshal(userTokenTenants)
//	if err != nil {
//		log.Errorf("json Marshal is error %+v", userTokenTenants)
//		return
//	}
//	log.Infof("UserToken json is \n %s", string(resp))
//	log.Infof("redis key is %s,%s","drbac:"+token,"drbac:"+userTokenTenants.User.Username,"time is 24*3600 second")
//	client.Send("MULTI")
//	client.Send("SET", "drbac:"+token, resp, "EX", 24*3600)
//	//client.Send("SET", "drbac:"+userToken.User.Username+":"+userToken.TenantRolePermissions.Tenant.TenantName, token, "EX", 24*3600)
//	_, err = client.Do("EXEC")
//	return
//}

/*
	保存token
*/
func saveUserTokenDomain(userToken *UserToken, token string, pool *redis.Pool) (err error) {
	if userToken == nil {
		err = errors.New("userToken is nil")
		return
	}
	if userToken.User.Username == "" {
		log.Info("usertoken:", userToken)
		err = errors.New("userToken is nil")
		return
	}
	client := pool.Get()
	defer client.Close()
	resp, err := json.Marshal(userToken)
	if err != nil {
		log.Errorf("json Marshal is error %+v", userToken)
		return
	}
	log.Infof("UserToken json is \n %s", string(resp))
	log.Infof("redis key is %s,%s", "drbac:"+token, "drbac:"+userToken.User.Username, "time is 24*3600 second")
	client.Send("MULTI")
	client.Send("SET", "drbac:"+token, resp, "EX", 24*3600)
	//client.Send("SET", "drbac:"+userToken.User.Username+":"+userToken.TenantRolePermissions.Tenant.TenantName, token, "EX", 24*3600)
	_, err = client.Do("EXEC")
	return
}

/*
	保存token
*/
func saveUserTokenTenants(userTokenTenants *UserTokenTenants, token string, pool *redis.Pool) (err error) {
	if userTokenTenants == nil {
		err = errors.New("userToken is nil")
		return
	}
	if userTokenTenants.User.Username == "" {
		log.Info("usertoken:", userTokenTenants)
		err = errors.New("userToken is nil")
		return
	}
	client := pool.Get()
	defer client.Close()
	resp, err := json.Marshal(userTokenTenants)
	if err != nil {
		log.Errorf("json Marshal is error %+v", userTokenTenants)
		return
	}
	log.Infof("UserToken json is \n %s", string(resp))
	log.Infof("redis key is %s,%s", "drbac:"+token, "drbac:"+userTokenTenants.User.Username, "time is 24*3600 second")
	client.Send("MULTI")
	client.Send("SET", "drbac:"+token, resp, "EX", 24*3600)
	//client.Send("SET", "drbac:"+userToken.User.Username+":"+userToken.TenantRolePermissions.Tenant.TenantName, token, "EX", 24*3600)
	_, err = client.Do("EXEC")
	return
}

/*
	删除token
*/
func deleteToken(token string, pool *redis.Pool) (err error) {
	client := pool.Get()
	defer client.Close()
	userToken := UserToken{}
	resp, err := client.Do("GET", "drbac:"+token)
	if err != nil || resp == nil {
		log.Errorf("redis get token %s error %s, resp is %s", token, err, resp)
		return
	}
	err = json.Unmarshal([]byte(resp.([]uint8)), &userToken)
	_, err = client.Do("DEL", "drbac:"+token)
	if err != nil {
		log.Errorf("delete token error is %s, redis key is %s", err, "drbac:"+token)
	}
	return
}

/*
	基于username获取用户的token
*/
func getTokenByUsernameAndTenantName(username, tenantName string, pool *redis.Pool) (token string) {
	client := pool.Get()
	defer client.Close()
	token, err := redis.String(client.Do("GET", "drbac:"+username+":"+tenantName))
	if err != nil {
		log.Errorf("get Token By Username in redis error:%s,redis key :%s", err, "drbac:"+username+":"+tenantName)
	}
	return
}

/*
	基于用户id获取token值
*/
func getTokenByUidAndTid(uid, tid int32, tx *gorm.DB, pool *redis.Pool) (token string) {
	user := User{Uid: uid}
	err := user.GetUserByUID(tx)
	if err != nil {
		log.Errorf("GetUserByUID is error %s,uid is %d", err, uid)
		return
	}
	tenant := Tenant{Tid: tid}
	err = tenant.GetByID(tx)
	if err != nil || tenant.TenantName == "" {
		log.Errorf("GetByID is error %s,did is %d", err, tid)
		return
	}
	token = getTokenByUsernameAndTenantName(user.Username, tenant.TenantName, pool)
	return
}

/*
	基于uid踢出登录用户
*/
func deleteTokenByUidAndTid(uid, tid int32, tx *gorm.DB, pool *redis.Pool) {
	user := User{Uid: uid}
	err := user.GetUserByUID(tx)
	if err != nil {
		log.Errorf("GetUserByUID is error %s,uid is %d", err, uid)
		return
	}
	tenant := Tenant{Tid: tid}
	err = tenant.GetByID(tx)
	if err != nil {
		log.Errorf("GetByID is error %s,did is %d", err, tid)
		return
	}
	token := getTokenByUsernameAndTenantName(user.Username, tenant.TenantName, pool)
	deleteToken(token, pool)
	return
}
