package module

import (
	"time"

	log "github.com/cihub/seelog"
	"github.com/garyburd/redigo/redis"
)

const (
	LoginRedisKeyPrefix    string = "Ag:domain-rpc:login:"
	FindpwReidsKeyPrefix   string = "Ag:domain-rpc:findpw:"
	ActiviteRedisKeyPrefix string = "Ag:domain-rpc:activite:"
	UidRedisKey            string = "Ag:domain-rpc:uid"
	Nopersistence          string = "nopersistence"
	Persistence            string = "persistence"
)
const (
	redisMaxIdle        = 3
	redisIdleTimeoutSec = 24)

func NewRedisPool(redisURL string) *redis.Pool {
	log.Infof("url %s", redisURL)
	return  &redis.Pool{
		MaxIdle:     redisMaxIdle,
		IdleTimeout: redisIdleTimeoutSec * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(redisURL)
			if err != nil {
				log.Errorf("redis connection error: %s",err)
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				log.Errorf("ping redis error: %s",err)
				return err
			}
			return nil
		},
	}
}
