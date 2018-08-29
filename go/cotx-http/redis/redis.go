package redis

import (
	"github.com/garyburd/redigo/redis"
	"net"
	"time"
	log "github.com/cihub/seelog"
)

const (
	SourceRedisKeyPrefix    string = "pub:source:"
)

var (
	redisClient map[string]*redis.Pool = make(map[string]*redis.Pool, 0)
)

var RedisConn redis.Conn

func RedisClient(name string) *redis.Pool {
	return redisClient[name]
}

func NewRedisClient(name, host, port string, maxIdle, maxActive int, password string) {
	redisClient[name] = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			C, err := redis.Dial("tcp", net.JoinHostPort(host, port))
			if err != nil {

				return nil, err
			}
			if _, err := C.Do("AUTH", password); err != nil {
				log.Error("redis password error", err)
				return nil, err
			}
			return C, nil
		},
	}
}

//检验用户source是否正确
func CheckSource(source string)  (bool ,error){
	log.Info("Start checkSource")
	client := RedisClient("persistence").Get()
	defer client.Close()
	res, err := client.Do("get",SourceRedisKeyPrefix + source)
	if err != nil || res == nil{
		log.Error("验证source失败 ",source)
		return false,err
	}
	return true,nil
}

func InitRedisConn() {
	RedisConn = RedisClient("persistence").Get()
}


func GetRedisConn() redis.Conn  {
	return RedisConn
}

//redis的读操作
func ReadFromRedis(key string) (interface{},error) {
	log.Info("Redis:read  from redis")
	value ,err:= GetRedisConn().Do("get",key)
	return value,err
}
func WriteToRedis(key string,data []byte) error  {
	log.Info("Redis:write  To redis")
	_,err := GetRedisConn().Do("set",key,data,"EX",60*5)
	log.Info("Redis:write  over")
	return  err
}

func DeleteFromRedis(key string) error {
	_,err := GetRedisConn().Do("DEL",key)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}


