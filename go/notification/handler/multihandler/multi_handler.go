package multihandler

import (
	"fmt"
	"notification/handler"
	"github.com/pkg/errors"
	"notification/handler/mqtthandler"
)

var Multihandler map[string]handler.Handler = make(map[string]handler.Handler, 0)

func NewMultiHandler(handler ...string) (err error) {
	for _, v := range handler {
		switch v {
		case "mqtt":
			Multihandler["mqtt"] = mqtthandler.MqttHandler
		default:
			errStr := fmt.Sprintf("not have handler is %s", v)
			err = errors.New(errStr)
			return
		}
	}
	return
}

func GetHandler(handler string) handler.Handler {
	h, ok := Multihandler[handler]
	if ok {
		return h
	}
	return nil
}
