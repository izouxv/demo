package mqtthandler

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"sync"
	"time"
	"notification/handler"
	log "github.com/cihub/seelog"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/garyburd/redigo/redis"
	"fmt"
)

const EmailTopic = "notice/email"

var MqttHandler handler.Handler

type MQTTHandler struct {
	conn                      mqtt.Client
	receivedPayloadEmail      chan handler.PayloadEmail
	wg                        sync.WaitGroup
	redisPool                 *redis.Pool
}

/*NewHandler 创建Mq客户端*/
func NewHandler(server, username, password, cafile string) (handler.Handler,error) {
	log.Info("启动MQTT服务")
	h := MQTTHandler{
		receivedPayloadEmail:      make(chan handler.PayloadEmail),
	}
	opts := mqtt.NewClientOptions()
	opts.AddBroker(server)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetOnConnectHandler(h.onConnected)
	opts.SetConnectionLostHandler(h.onConnectionLost)

	if cafile != "" {
		log.Info("没有ca文件指定")
		tlsConfig, err := newTLSConfig(cafile)
		if err != nil {
			log.Errorf("Error with the mqtt CA certificate: %s", err)
		} else {
			opts.SetTLSConfig(tlsConfig)
		}
	}
	log.Infof("server %s handler/mqtt: connecting to mqtt broker", server)
	h.conn = mqtt.NewClient(opts)
	for {
		if token := h.conn.Connect(); token.Wait() && token.Error() != nil {
			log.Errorf("handler/mqtt: connecting to broker error, will retry in 2s: %s", token.Error())
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}
	return &h, nil
}

func newTLSConfig(cafile string) (*tls.Config, error) {
	cert, err := ioutil.ReadFile(cafile)
	if err != nil {
		log.Errorf("backend: couldn't load cafile: %s", err)
		return nil, err
	}
	certpool := x509.NewCertPool()
	certpool.AppendCertsFromPEM(cert)

	return &tls.Config{

		RootCAs: certpool,
	}, nil
}

//关闭mqtt连接
func (h *MQTTHandler) Close() error {
	log.Info("handler/mqtt: closing handler...")
    	log.Infof("notice_email topic", EmailTopic, "handler/mqtt: unsubscribing from tx topic")
	if token := h.conn.Unsubscribe(EmailTopic); token.Wait() && token.Error() != nil {
		return fmt.Errorf("handler/mqtt: unsubscribe from %s error: %s", EmailTopic, token.Error())
	}
	log.Info("handler/mqtt: handling last items in queue")
	h.wg.Wait()
	close(h.receivedPayloadEmail)
	return nil
}

func (h *MQTTHandler) ReceivedChanEmail() chan handler.PayloadEmail {
	return h.receivedPayloadEmail
}

///mqtt主题接收到的数据处理到通道
func (h *MQTTHandler) EmailPayloadHandler(c mqtt.Client, msg mqtt.Message) {
	h.wg.Add(1)
	defer h.wg.Done()
	log.Infof("topic", msg.Topic(), "handler/mqtt: data-down payload received")
	var pl handler.PayloadEmail
	if  err := json.Unmarshal([]byte(msg.Payload()),&pl); err != nil {
		log.Infof("mqtt接收到的数据umarshal错误 err (%s)",err)
		return
	}
	h.receivedPayloadEmail <- pl
}

func (h *MQTTHandler) onConnected(c mqtt.Client) {
	log.Info("handler/mqtt: connected to mqtt")
	for {
		log.Infof("EmailTopic %s handler/mqtt: subscribling to tx topic",EmailTopic)
		if token := h.conn.Subscribe(EmailTopic, 2, h.EmailPayloadHandler); token.Wait() && token.Error() != nil {
			log.Infof("topic", EmailTopic, "handler/mqtt: subscribe error: %s", token.Error())
			time.Sleep(time.Second)
			continue
		}
		return
	}
}

func (h *MQTTHandler) onConnectionLost(c mqtt.Client, reason error) {
	log.Errorf("handler/mqtt: mqtt connection error: %s", reason)
}
