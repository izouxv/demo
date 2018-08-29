package drbac

import (
	"encoding/json"
	"regexp"

	. "auth/go-drbac/internal/storage"

	log "github.com/cihub/seelog"

	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

/*url白名单*/
func checkWhitelist(url, opt string, tx *gorm.DB) bool {
	log.Info("开始校验白名单")
	wl := WhiteList{}
	allWhiteList,err := wl.GetAllWhiteList(tx)
	log.Info("len(allWhiteList)",len(allWhiteList))
	if len(allWhiteList) == 0 || err != nil {
		return false
	}
	log.Info("请求url:",url," opt:",opt)
	for _,v := range allWhiteList {
		log.Info("v:",v)
		match, _ := regexp.MatchString(v.Url, url)
		if match && v.Opt == opt {
			log.Info("校验通过")
			return true
		}
	}
	return false
}


/*授权*/
func authorizationTenant(tid int32, url, opt, token string, pool *redis.Pool) (userToken *UserToken, isTrue bool) {
	userToken = getToken(token, pool)
	if userToken == nil {
		log.Errorf("getToken is nil")
		return
	}
	if userToken.TenantRoleResource == nil {
		log.Errorf("getToken is nil")
		return
	}
	isTrue = authorityJudgmentTenant(tid,userToken, url, opt)
	return
}

/*授权*/
func authorizationDomain(did int32, url, opt, token string, pool *redis.Pool) (userToken *UserToken, isTrue bool) {
	userToken = getToken(token, pool)
	if userToken == nil {
		log.Errorf("getToken is nil")
		return
	}
	if userToken.DomainRoleResource == nil {
		log.Errorf("getToken is nil")
		return
	}
	isTrue = judgmentPermissions(userToken.DomainRoleResource.Resource, url, opt)
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
		log.Info("Json Unmarshal userToken Error", err)
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
func authorityJudgmentTenant(tid int32, userToken *UserToken, url, opt string) bool {
	if tid != 0 {
		if tid == 1 || tid == 3 {
			return judgmentPermissions(userToken.DomainRoleResource.Resource, url, opt)
		}
		for k,v := range userToken.TenantRoleResource {
			if v.Tenant.Tid == tid {
				return judgmentPermissions(userToken.TenantRoleResource[k].Resource, url, opt)
			}
		}
	}
	return judgmentPermissions(userToken.TenantRoleResource[0].Resource, url, opt)
}


/*judgmentPermissions
判断权限
*/
func judgmentPermissions(resources []*Resource, url, opt string) bool {
	log.Info("resources:",resources)
	for _, v := range resources {
		log.Info("permissions URL: ",v.ResUrl, " permissions OPT: ",v.ResOpt)
		match, _ := regexp.MatchString(v.ResUrl, url)
		if match && v.ResOpt == opt {
			log.Info("校验成功")
			return true
		}
	}
	return false
}

/*judgmentChildren
判断子域中是否存在该did
*/
func judgmentChildren(Children []*TenantTree, tid int32) bool {
	for _, v := range Children {
		if v.Tenant.Tid == tid {
			return true
		}
		if judgmentChildren(v.Children, tid) {
			return true
		}
	}
	return false
}

