package module

import (
	log "github.com/cihub/seelog"
	"github.com/garyburd/redigo/redis"
	"net"
	"time"
)

const (
	LoginRedisKeyPrefix    string = "Ag:domain-rpc:login:"
	FindpwReidsKeyPrefix   string = "Ag:domain-rpc:findpw:"
	ActiviteRedisKeyPrefix string = "Ag:domain-rpc:activite:"
	DynamicRedisKey        string = "Ag:domain-rpc:node_dynamic"
	UidRedisKey            string = "Ag:domain-rpc:uid"
	AirInfoRedisKey        string = "Ag:domain-rpc:air"
	Nopersistence          string = "nopersistence"
	Persistence            string = "persistence"
	ActivateNodeRedisKey   string = "Ag:domain-rpc:Activation:"
	UpgradeCount           string = "Ag:domain-rpc:upgrade"
	HeartbeatKeyPrefixSet  string = "Ag:domain-rpc:heartbeat"
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
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			C, err := redis.Dial("tcp", net.JoinHostPort(host, port))
			if err != nil {
				log.Error("redis Dial error:", err)
				panic(err)
				return nil, err
			}
			if _, err := C.Do("AUTH", password); err != nil {
				log.Error("redis password error:", err)
				panic(err)
				return nil, err
			}
			return C, nil
		},
	}
}
