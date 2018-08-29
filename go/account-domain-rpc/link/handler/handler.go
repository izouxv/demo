package handler

import (
	log "github.com/cihub/seelog"
	"github.com/pkg/errors"
)

// Handler defines the interface of a handler backend.
type Handler interface {
	Close() error                                          // closes the handler
	SendDataUp(payload DataUpPayload) error                // send data-up payload
	SendJoinNotification(payload JoinNotification) error   // send join notification
	SendACKNotification(payload ACKNotification) error     // send ack notification
	SendErrorNotification(payload ErrorNotification) error // send error notification
	DataDownChan() chan DataDownPayload                    // returns DataDownPayload channel
}

const (
	MQTT = "mqtt"
	DB   = "db"
	RPC  = "rpc"
)

var h Handler

func NewHandler(handler string) {
	var err error

	switch handler {
	case DB:
		h, err = NewDBHandler()
		log.Info("New DB Handler")
		if err != nil {
			panic(errors.New("unknow handler"))
		}
	case MQTT:
		h, err = NewMQTTHandler("", "", "", "")
		log.Info("New MQTT Handler")
		if err != nil {
			panic(errors.New("unknow handler"))
		}
	case RPC:
		log.Info("New RPC Handler")
	}
}

func GetHandler() (Handler, error) {
	if h != nil {
		return h, nil
	}
	return h, errors.New("handler is nil")
}
