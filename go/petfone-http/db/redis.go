package db

import (
	log "github.com/cihub/seelog"
	"github.com/garyburd/redigo/redis"
	"reflect"
	"petfone-http/core"
	"petfone-http/util"
	"petfone-http/result"
	"net/http"
)

const (
	FeedBackMail		string = "Ag:petfone:v1.0:feedback:mail"
	RedisKeySource      string = "pub:source:"
	Version				string = "Ag:petfone:v1.0:test:version:"
	UDevices            string = "Ag:petfone:uid:dids:" //用户-设备权限
	DUsers              string = "Ag:petfone:did:uids:" //设备-用户权限
	UPets               string = "Ag:petfone:uid:pids:" //用户-宠物权限
	PUsers              string = "Ag:petfone:pid:uids:" //宠物-用户权限
	PetDeviceShareToken string = "Ag:petfone:uid:suid"  //设备分享token
)

const (
	expire           string = "EX"
	set              string = "set"
	get              string = "get"
	del              string = "del"
	zadd             string = "zadd"
	zscore           string = "zscore"
	zcard            string = "zcard"
	zscan            string = "zscan"
	zcount           string = "zcount"
	zrevrangebyscore string = "zrevrangebyscore"
)

//校验source
func VerifySource(source string) bool {
	if util.VerifyParamsStr(source) {
		return true
	}
	//从redis验证source
	if _, flag := GetSource(source); flag {
		switch source {
		case "AQIDAA==":
		case "AgIBAA==":
		case "AgICAA==":
		case "AgIDAA==":
		case "AgECAA==":
		case "AgEBAA==":
		case "AgEDAA==":
		default:
			return true
		}
	}
	return false
}

//redis 校验权限
func VerifyUserPermiss(port int32, t1 string, t2 string, res http.ResponseWriter) bool {
	log.Info("VerifyUserPermiss-dbs:", port, ",t1:", t1, ",t2:", t2)
	num1, num2, err := Redis_ZscoreZcard(port, t1, t2)
	if err != nil {
		log.Info("VerifyUserPermiss-err:", err)
		result.RESC(10001, res)
		return true
	}
	if num1 <= 0 || num2 <= 0 {
		result.RESC(21005, res)
		return true
	}
	return false
}

/**
redis命令
 */

//获取redis source
func GetSource(source string) (string, bool) {
	log.Info("GetSource-source:", source)
	client := core.RedisClient(6379)
	defer client.Close()
	key := RedisKeySource + source
	value, err := redis.String(client.Do("get", key))
	if err != nil || value == "" {
		log.Info("GetSource-err:", err)
		return "", true
	}
	log.Info("GetSource-value:", value)
	return value, false
}

//set string expire
func Redis_GetEx(db int32, key string) (bool, error) {
	client := core.RedisClient(db)
	reply, err := client.Do(get, key)
	defer func() { client.Close() }()
	log.Info("Redis_SetEx-reply:", reply)
	if err != nil || "OK" != reply {
		log.Info("Redis_SetEx-err:", err)
		return true, err
	}
	return false, nil
}

func Redis_DelEx(db int32, key string) (bool, error) {
	client := core.RedisClient(db)
	reply, err := client.Do(del, key)
	defer func() { client.Close() }()
	log.Info("Redis_DelEx-reply:", reply)
	if err != nil || "OK" != reply {
		log.Error("Redis_DelEx-err:", err)
		return true, err
	}
	return false, nil
}

//返回有序集中，区间内的成员数量
func Redis_Zcount(db int32, key string, min int32, max int32) (int32, error) {
	client := core.RedisClient(db)
	reply2, err := redis.Int64(client.Do(zcount, key, min, max))
	defer func() { client.Close() }()
	log.Info("ParamsType-type:", reflect.TypeOf(reply2))
	log.Info("Redis_ZscoreZcard-reply2:", reply2)
	if err != nil || reply2 == 0 {
		log.Error("Redis_Zscore-err:", err)
		return 0, err
	}
	re2 := util.Int64ToInt32(reply2)
	return int32(re2), nil
}

//返回有序集中，分数区间内的成员
func Redis_Zrevrangebyscore(db int32, key string, min int32, max int32) ([]string, error) {
	client := core.RedisClient(db)
	reply, err := redis.Values(client.Do(zrevrangebyscore, key, min, max))
	defer func() { client.Close() }()
	log.Info("ParamsType-type:", reflect.TypeOf(reply))
	log.Info("Redis_ZscoreZcard-reply2:", reply)
	if err != nil {
		log.Error("Redis_Zscore-err:", err)
		return nil, err
	}
	var strs []string
	for _, v := range reply {
		strs = append(strs, string(v.([]uint8)))
	}
	return strs, nil
}

//返回有序集中，成员的分数值与成员数
func Redis_ZscoreZcard(db int32, key string, member string) (int32, int32, error) {
	client := core.RedisClient(db)
	defer func() { client.Close() }()
	reply1, err := redis.Int(client.Do(zscore, key, member))
	reply2, err := redis.Int64(client.Do(zcard, key))
	log.Info("Redis_ZscoreZcard-reply1:", reply1, ",reply2:", reply2)
	if err != nil {
		log.Error("Redis_Zscore-err:", err)
		return 0, 0, err
	}
	return int32(reply1), util.Int64ToInt32(reply2), nil
}

//返回有序集中，迭代成员的分数值与成员
func Redis_Zscan(db int32, key string) (map[string]int, error) {
	client := core.RedisClient(db)
	reply, err := redis.Values(client.Do(zscan, key, 0))
	defer func() { client.Close() }()
	if err != nil {
		log.Error("Redis_Zscore-err:", err)
		return nil, err
	}
	re, err := redis.IntMap(reply[1], err)
	if err != nil {
		log.Error("Redis_Zscore-err:", err)
		return nil, err
	}
	return re, nil
}
