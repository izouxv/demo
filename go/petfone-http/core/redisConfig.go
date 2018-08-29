package core

import (
	log "github.com/cihub/seelog"
	"github.com/garyburd/redigo/redis"
	"net"
	"time"
)

var (
	//6379端口连接池
	RedisMap = make(map[string]*redis.Pool, 0)
)

func RedisClient(port int32) redis.Conn {
	if port == 6379 {
		return RedisMap["redis6379"].Get()
	}
	return RedisMap["redis6380"].Get()
}

//结束redis连接
func CloseRedis() {
	log.Info("CloseRedis")
	var err error
	for _, conn := range RedisMap {
		err = conn.Close()
		if err != nil {
			log.Info("CloseRedis err:", err)
		}
	}
}

//redis客户端配置
func RedisConfig(name, host, port string, maxIdle, maxActive int, pwd string) {
	RedisMap[name] = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			C, err := redis.Dial("tcp", net.JoinHostPort(host, port))
			if err != nil {
				log.Error("redis连接错误:", err)
				return nil, err
			}
			if _, err := C.Do("AUTH", pwd); err != nil {
				log.Error("redis密码错误:", err)
				return nil, err
			}
			return C, nil
		},
	}
	tmp := RedisMap[name].Get()
	_, err := tmp.Do("Ping")
	defer tmp.Close()
	if err != nil {
		log.Error("redis连接错误:", err)
		panic(err)
	}
	log.Info("redis-name:", name, ",host:", host, ",port:", port, ",maxIdle:", maxIdle, ",maxActive:", maxActive, ",pwd:", pwd)
}
