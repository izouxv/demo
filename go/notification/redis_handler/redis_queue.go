package redis_handler

import (
	log "github.com/cihub/seelog"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"encoding/json"
	"notification/common"
	"notification/config"
	"time"
)

func ReceiveRedisMess() {
	log.Infof("start redis message ....")
	client := config.C.Redis.Pool.Get()
	defer client.Close()
	psc := redis.PubSubConn{client}
	psc.Subscribe("filebeat")
	var rn common.RedisNotice
	for {
		switch v := psc.Receive().(type) {
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case redis.Message:
			log.Infof("redis接收到的数据是:(%s)",v.Data)
			if err := json.Unmarshal(v.Data,&rn);err != nil {
				log.Infof("数据umarshal出错, err (%s)",err)
				return
			}
			rn.Timestamp = 	time.Now().Format("2006-01-02 15:04:05")
			log.Infof("struct处理之后的数据是:(%s)",rn)
			if err := common.WechatNotice(rn);err != nil {
				log.Infof("微信通知失败...err (%s)",err)
				return
			}
		case redis.PMessage:
			log.Infof("模式订阅psubscribe...")
			fmt.Printf("PMessage: %s %s %s\n", v.Pattern, v.Channel, v.Data)
		case error:
			return
		}
	}
}





