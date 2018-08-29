package handler

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"account-domain-rpc/module"
	"account-domain-rpc/storage"
	"account-domain-rpc/util/id64"
	log "github.com/cihub/seelog"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// MQTTHandler implements a MQTT handler for sending and receiving data by
// an application.
type DBHandler struct {
	db           *gorm.DB
	dataDownChan chan DataDownPayload
	wg           sync.WaitGroup
	redisPool    *redis.Pool
}

// NewMQTTHandler creates a new MQTTHandler.
func NewDBHandler() (Handler, error) {
	p := module.RedisClient(module.Nopersistence)
	h := DBHandler{
		db:           module.MysqlClient(),
		dataDownChan: make(chan DataDownPayload),
		redisPool:    p,
	}
	return &h, nil
}

func GetRandomId() int64 {
	did, err := id64.NewNode(1)
	if err != nil {
		log.Info("NewNode :get did error")
	}
	id := int64(did.Generate())
	return id
}

func (h *DBHandler) SendDataUp(payload DataUpPayload) error {
	var err error
	node := storage.Node{DevEUI: payload.DevEUI[:]}
	if err = node.GetNodeByDevEUI(); err != nil {
		log.Error("dev_eui is error not found")
		return err
	}
	err = storage.Transaction(module.MysqlClient(), func(tx *gorm.DB) error {
		dataup := storage.DataUp{
			Id:          GetRandomId(),
			DevEUI:      payload.DevEUI[:],
			Frequency:   int32(payload.TXInfo.Frequency),
			Modulation:  payload.TXInfo.DataRate.Modulation,
			BandWidth:   int32(payload.TXInfo.DataRate.Bandwidth),
			Spreafactor: int32(payload.TXInfo.DataRate.SpreadFactor),
			Bitrate:     int32(payload.TXInfo.DataRate.Bitrate),
			CodeRate:    payload.TXInfo.CodeRate,
			ADR:         payload.TXInfo.ADR,
			Fcnt:        payload.FCnt,
			Fport:       payload.FPort,
			Data:        payload.Data,
		}
		err = dataup.CreateDataup(tx)
		if err != nil {
			log.Info("CreateDataup error")
			return errors.Wrap(err, "Create Dataup error")
		}
		for i := 0; i < len(payload.RXInfo); i++ {
			txinfo := storage.TxInfo{
				Id:        GetRandomId(),
				Mac:       payload.RXInfo[i].MAC[:],
				Time:      payload.RXInfo[i].Time,
				RSSI:      payload.RXInfo[i].RSSI,
				LoraSNR:   payload.RXInfo[i].LoRaSNR,
				Longitude: payload.RXInfo[i].Longitude,
				Latitude:  payload.RXInfo[i].Latitude,
				Altitude:  payload.RXInfo[i].Altitude,
			}
			err = txinfo.CreatTxInfo(tx)
			if err != nil {
				log.Info("CreateTxInfo error")
				return errors.Wrap(err, "Create TxInfo error")
			}
			dataup_txinfo := storage.DataupTxinfo{
				DataUpId: dataup.Id,
				RxInfoId: txinfo.Id,
			}
			err = dataup_txinfo.CreateDataupTxinfo(tx)
			if err != nil {
				log.Info("Create dataup_txinfo error")
				return errors.Wrap(err, "Create Dataup_txinfo error")
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	go heartbeat(node.Nid)
	return nil
}

//设备在线状态心跳
func heartbeat(id int64) {
	client := module.RedisClient(module.Persistence).Get()
	defer client.Close()
	node := storage.Node{Nid: id}
	err := node.GetNodeByNid()
	key := fmt.Sprintf("%s:%d:%d", module.HeartbeatKeyPrefixSet, node.Did, node.ApplicationID)
	resp, err := client.Do("ZADD", key, time.Now().Unix(), id)
	if err != nil {
		log.Error(err)
	}
	log.Info(resp)
}

// Close stops the handler.
func (h *DBHandler) Close() error {
	log.Info("handler/db: closing handler")
	close(h.dataDownChan)
	return nil
}

// SendJoinNotification sends a JoinNotification.
func (h *DBHandler) SendJoinNotification(payload JoinNotification) error {
	b, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("handler/db: join notification marshal error: %s", err)
	}
	fmt.Println(b)
	//todo save b >> db
	return nil
}

// SendACKNotification sends an ACKNotification.
func (h *DBHandler) SendACKNotification(payload ACKNotification) error {
	b, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("handler/db: ack notification marshal error: %s", err)
	}
	fmt.Println(b)
	//todo save b >> db
	return nil
}

// SendErrorNotification sends an ErrorNotification.
func (h *DBHandler) SendErrorNotification(payload ErrorNotification) error {
	b, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("handler/db: error notification marshal error: %s", err)
	}
	fmt.Println(b)
	//todo save b >> db
	return nil
}

// DataDownChan returns the channel containing the received DataDownPayload.
func (h *DBHandler) DataDownChan() chan DataDownPayload {
	return h.dataDownChan
}
