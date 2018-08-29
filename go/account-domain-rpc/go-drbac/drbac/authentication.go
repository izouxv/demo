package drbac

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	. "account-domain-rpc/go-drbac/internal/storage"

	log "github.com/cihub/seelog"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

type UserToken struct {
	Token                 string
	User                  *User
	DomainRolePermissions []*DomainRolePermission
}

/*
 *	认证
 */
func authentication(username, password string, tx *gorm.DB, pool *redis.Pool) (userToken *UserToken, isTrue bool, err error) {
	u := &User{Username: username, Password: password}
	isTrue, err = u.IsTruePassword(password, tx)
	if err != nil {
		log.Errorf("password is error,username %s , password %s, err is %s", err, username, password)
		return
	}
	if !isTrue {
		log.Errorf("password is error,username %s , password %s", username, password)
		return
	}
	drps, err := getDomainRolePermissionByUid(u.Uid, tx)
	if err != nil {
		log.Errorf("GetDomainRolePermissionByUid is error %s,u.Uid is %d ", err, u.Uid)
		return
	}
	token := generateToken(u.Uid)
	userToken = &UserToken{User: u, DomainRolePermissions: drps, Token: token}
	err = saveToken(userToken, token, pool)
	if err != nil {
		log.Errorf("saveToken error %s", err)
		return
	}
	return
}

/*
 *	重认证，修改redis中的认证信息（刷新缓存）
 */
func reAuthentication(token string, tx *gorm.DB, pool *redis.Pool) (userToken *UserToken, err error) {
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
	drps, err := getDomainRolePermissionByUid(u.Uid, tx)
	if err != nil {
		log.Errorf("GetDomainRolePermissionByUid error %s uid is %d", err, u.Uid)
		return
	}
	userToken = &UserToken{User: u, DomainRolePermissions: drps, Token: token}
	err = saveToken(userToken, token, pool)
	if err != nil {
		log.Errorf("saveToken error %s", err)
		return
	}
	return
}

/*
	基于uid重新认证用户信息（刷新缓存）
*/
func reAuthenticationByUid(uid int64, tx *gorm.DB, pool *redis.Pool,
	reAuthentication func(token string, tx *gorm.DB, pool *redis.Pool) (userToken *UserToken, err error)) (userToken *UserToken, err error) {
	token := getTokenByUid(uid, tx, pool)
	return reAuthentication(token, tx, pool)
}

/*
	生成token值
*/
func generateToken(base int64) string {
	rand.Seed(time.Now().Unix())
	s := strconv.FormatInt(base, 10) + strconv.Itoa(rand.Int())
	session := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", session)
}

/*
	保存token
*/
func saveToken(userToken *UserToken, token string, pool *redis.Pool) (err error) {
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
	client.Send("SET", "drbac:"+userToken.User.Username, token, "EX", 24*3600)
	_, err = client.Do("EXEC")
	return
}

/*
	删除token
*/
func deleteToken(token string, pool *redis.Pool) (err error) {
	client := pool.Get()
	defer client.Close()
	_, err = client.Do("DEL", "drbac:"+token)
	if err != nil {
		log.Errorf("delete token error is %s, redis key is %s", err, "drbac:"+token)
	}
	return
}

/*
	基于username获取用户的token
*/
func getTokenByUsername(username string, pool *redis.Pool) (token string) {
	client := pool.Get()
	defer client.Close()
	token, err := redis.String(client.Do("GET", "drbac:"+username))
	if err != nil {
		log.Errorf("get Token By Username in redis error:%s,redis key :%s", err, "drbac:"+username)
	}
	return
}

/*
	基于用户id获取token值
*/
func getTokenByUid(uid int64, tx *gorm.DB, pool *redis.Pool) (token string) {
	user := User{Uid: uid}
	err := user.GetUserByUID(tx)
	if err != nil {
		log.Errorf("GetUserByUID is error %s,uid is %d", err, uid)
		return
	}
	token = getTokenByUsername(user.Username, pool)
	return
}

/*
	基于uid提出登录用户
*/
func deleteTokenByUid(uid int64, tx *gorm.DB, pool *redis.Pool) {
	user := User{Uid: uid}
	err := user.GetUserByUID(tx)
	if err != nil {
		log.Errorf("GetUserByUID is error %s,uid is %d", err, uid)
		return
	}
	token := getTokenByUsername(user.Username, pool)
	deleteToken(token, pool)
	return
}
