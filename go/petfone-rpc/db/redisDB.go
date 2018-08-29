package db

import (
	"petfone-rpc/core"
	log "github.com/cihub/seelog"
	"strconv"
	"github.com/garyburd/redigo/redis"
	"petfone-rpc/util"
)

//对key的操作
func Redis_key()  {
	
}

//对string操作
func Redis_Str(port int32, cmd, key string) (interface{}, error) {
	client := core.RedisClient(port)
	defer client.Close()
	value, err := client.Do(cmd, key)
	if err != nil {
		log.Errorf("Redis_Str err:%s,port:%d,cmd:%s,key:%s", err,port,cmd,key)
		return nil,err
	}
	return value,nil
}

//对string操作
func RedisCmd(port int32, cmd string,args ...interface{}) (interface{}, error) {
	client := core.RedisClient(port)
	defer client.Close()
	value, err := client.Do(cmd, args)
	if err != nil {
		log.Errorf("Redis_Str err:%s,port:%d,cmd:%s,args:%s", err,port,cmd,args)
		return nil,err
	}
	return value,nil
}







//添加元素并赋予分数
func Redis_Zadd(db int32, key string, num int32, member string) (bool, error) {
	client := core.RedisClient(db)
	reply , err := client.Do(util.Zadd, key, num, member)
	defer client.Close()
	if err != nil || reply == 0 {
		log.Info("Redis_Zadd-err:",err)
		return true,err
	}
	return false,nil
}

//移除元素
func Redis_Zrem(db int32, key string, member string) (bool, error) {
	client := core.RedisClient(db)
	reply , err := redis.Int(client.Do(util.Zrem, key, member))
	defer client.Close()
	if err != nil{
		log.Info("Redis_Zrange-err:",err)
		return true,err
	}
	log.Info("Redis_Zrange-reply:",reply)
	return false,nil
}

//返回有序集中，成员的分数值与成员数
func Redis_ZscoreZcard(db int32, key string, member string) (int32, int32, error) {
	client := core.RedisClient(db)
	reply, err := client.Do(util.Zscore, key, member)
	reply2, err := client.Do(util.Zcard, key)
	defer client.Close()
	log.Info("Redis_ZscoreZcard-reply:",reply,",reply2:",reply2)
	if err != nil || reply == 0 || reply2 == 0 {
		log.Info("Redis_Zscore-err:",err)
		return 0,0,err
	}
	re, err := strconv.Atoi(string(reply.([]uint8)))
	re2, err := strconv.Atoi(string(reply2.([]uint8)))
	if err != nil {
		log.Info("Redis_Zscore-err:",err)
		return 0,0,err
	}
	return int32(re),int32(re2),nil
}

//返回有序集中，分数区间内的成员
func Redis_Zrevrangebyscore(db int32, key string , min int32, max int32) ([]string, error) {
	client := core.RedisClient(db)
	reply, err := redis.Values(client.Do(util.Zrevrangebyscore, key, min, max))
	defer client.Close()
	log.Info("Redis_Zrevrangebyscore-reply:",reply)
	if err != nil {
		log.Info("Redis_Zrevrangebyscore-err:",err)
		return nil,err
	}
	var strs []string
	for _,v := range reply {
		strs = append(strs, string(v.([]uint8)))
	}
	return strs,nil
}

//set string expire
func Redis_SetEx(db int32, key string, value string, num int32) (bool, error) {
	client := core.RedisClient(db)
	reply , err := client.Do(util.SetStr, key, value, util.Expire, num)
	defer client.Close()
	log.Info("Redis_SetEx-reply:",reply)
	if err != nil || "OK" != reply {
		log.Info("Redis_SetEx-err:",err)
		return true,err
	}
	return false,nil
}
