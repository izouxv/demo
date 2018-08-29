package module

import (
	"net"
	"time"

	log "github.com/cihub/seelog"
	"github.com/garyburd/redigo/redis"
)

const (
	Nopersistence          string = "nopersistence"
	Persistence            string = "persistence"
)

var (
	redisClient map[string]*redis.Pool = make(map[string]*redis.Pool, 0)
)

func RedisClient(name string) *redis.Pool {
	return redisClient[name]
}

func NewRedisClient(name, host, port string, maxIdle, maxActive int, password string) {
	log.Info("New Redis Client")
	redisClient[name] = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: 30 * time.Second,
		Dial: func() (redis.Conn, error) {
			C, err := redis.Dial("tcp", net.JoinHostPort(host, port))
			if err != nil {
				panic(err)
				return nil, err
			}
			if _, err = C.Do("AUTH", password); err != nil {
				log.Error("redis password error", err)
				panic(err)
				return nil, err
			}
			return C, nil
		},
	}
}
