package core

import (
	log "github.com/cihub/seelog"
	"github.com/garyburd/redigo/redis"
	"net"
	"time"
)

var (
	redisClient = make(map[string]*redis.Pool, 0)
)

func RedisClient(port int32) redis.Conn {
	if port == 6379 {
		return redisClient["redis6379"].Get()
	}
	return redisClient["redis6380"].Get()
}

//todo redis初始化
func RedisInit(name, host, port string, maxIdle, maxActive int, password string) {
	redisClient[name] = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			C, err := redis.Dial("tcp", net.JoinHostPort(host, port))
			if err != nil {
				log.Error("redis连接错误:", err)
				panic(err)
			}
			if _, err := C.Do("AUTH", password); err != nil {
				log.Error("redis密码错误:", err)
				panic(err)
			}
			return C, nil
		},
	}
}

//todo 测试redis连接
func RedisPing(port int32) {
	conn := RedisClient(port)
	defer conn.Close()
	v, err := conn.Do("Ping")
	log.Info("ping:",port,v)
	if err != nil {
		log.Error("redis连接错误:", err)
		panic(err)
	}
}

func RedisClose() {
	for k, conn := range redisClient {
		log.Info(k,conn.Close())
	}
}

