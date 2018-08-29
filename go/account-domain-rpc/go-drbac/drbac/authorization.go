package drbac

import (
	"encoding/json"
	"regexp"

	. "account-domain-rpc/go-drbac/internal/storage"

	log "github.com/cihub/seelog"

	"github.com/garyburd/redigo/redis"
)

/*授权*/
func authorization(did int64, url, opt, token string, pool *redis.Pool) (userToken *UserToken, isTrue bool) {
	userToken = getToken(token, pool)
	if userToken == nil {
		log.Errorf("getToken is nil")
		return
	}
	isTrue = authorityJudgment(userToken, did, url, opt)
	return
}

/*获取用户信息*/
func getToken(token string, pool *redis.Pool) (userToken *UserToken) {
	client := pool.Get()
	defer client.Close()
	resp, err := client.Do("GET", "drbac:"+token)
	if err != nil || resp == nil {
		log.Errorf("redis get token %s error %s, resp is %s", token, err, resp)
		return
	}
	err = json.Unmarshal([]byte(resp.([]uint8)), &userToken)
	if err != nil {
		log.Error("json Unmarshal userToken error", err)
		return
	}
	go renewToken(token, pool)
	return
}

/*续租token*/
func renewToken(token string, pool *redis.Pool) (err error) {
	client := pool.Get()
	defer client.Close()
	_, err = client.Do("EXPIRE", "drbac:"+token, 24*3600)
	if err != nil {
		log.Errorf("redis EXPIRE token %s error %s", token, err)
	}
	return
}

/*authorityJudgment
权限判断
*/
func authorityJudgment(userToken *UserToken, did int64, url, opt string) bool {
	for _, v := range userToken.DomainRolePermissions {
		if v.Domain.Did == did {
			/*did是第一级的父域*/
			return judgmentPermissions(v.Permissions, url, opt)
		}
		/*did是子域id*/
		if judgmentChildren(v.Children, did) {
			return judgmentPermissions(v.Permissions, url, opt)
		}
	}
	return false
}

/*judgmentPermissions
判断权限
*/
func judgmentPermissions(permissions []*Permission, url, opt string) bool {
	for _, v := range permissions {
		match, _ := regexp.MatchString(v.Url, url)
		if match && v.Opt == opt {
			return true
		}
	}
	return false
}

/*judgmentChildren
判断子域中是否存在该did
*/
func judgmentChildren(Children []*DomainTree, did int64) bool {
	for _, v := range Children {
		if v.Domain.Did == did {
			return true
		}
		if judgmentChildren(v.Children, did) {
			return true
		}
	}
	return false
}

/*
判断域did是否合法
*/
func judgmentDomain(userToken *UserToken, did int64) bool {
	for _, v := range userToken.DomainRolePermissions {
		if v.Domain.Did == did {
			/*did是第一级的父域*/
			return true
		}
		/*did是子域id*/
		if judgmentChildren(v.Children, did) {
			return true
		}
	}
	return false
}
