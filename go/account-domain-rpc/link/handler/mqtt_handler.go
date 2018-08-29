package handler

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"

	"account-domain-rpc/module"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/garyburd/redigo/redis"
)

const txTopic = "application/+/node/+/tx"
const downlinkLockTTL = time.Millisecond * 100

var txTopicRegex = regexp.MustCompile(`application/(\w+)/node/(\w+)/tx`)

// MQTTHandler implements a MQTT handler for sending and receiving data by
// an application.
type MQTTHandler struct {
	conn         mqtt.Client
	dataDownChan chan DataDownPayload
	wg           sync.WaitGroup
	redisPool    *redis.Pool
}

// NewMQTTHandler creates a new MQTTHandler.
func NewMQTTHandler(server, username, password, cafile string) (Handler, error) {
	p := module.RedisClient(module.Nopersistence)
	h := MQTTHandler{
		dataDownChan: make(chan DataDownPayload),
		redisPool:    p,
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(server)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetOnConnectHandler(h.onConnected)
	opts.SetConnectionLostHandler(h.onConnectionLost)

	if cafile != "" {
		tlsconfig, err := newTLSConfig(cafile)
		if err != nil {
			log.Fatalf("Error with the mqtt CA certificate: %s", err)
		} else {
			opts.SetTLSConfig(tlsconfig)
		}
	}

	log.WithField("server", server).Info("handler/mqtt: connecting to mqtt broker")
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
	// Import trusted certificates from CAfile.pem.

	cert, err := ioutil.ReadFile(cafile)
	if err != nil {
		log.Errorf("backend: couldn't load cafile: %s", err)
		return nil, err
	}

	certpool := x509.NewCertPool()
	certpool.AppendCertsFromPEM(cert)

	// Create tls.Config with desired tls properties
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certpool,
	}, nil
}

// Close stops the handler.
func (h *MQTTHandler) Close() error {
	log.Info("handler/mqtt: closing handler")
	log.WithField("generate", txTopic).Info("handler/mqtt: unsubscribing from tx generate")
	if token := h.conn.Unsubscribe(txTopic); token.Wait() && token.Error() != nil {
		return fmt.Errorf("handler/mqtt: unsubscribe from %s error: %s", txTopic, token.Error())
	}
	log.Info("handler/mqtt: handling last items in queue")
	h.wg.Wait()
	close(h.dataDownChan)
	return nil
}

// SendDataUp sends a DataUpPayload.
func (h *MQTTHandler) SendDataUp(payload DataUpPayload) error {
	b, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("handler/mqtt: data-up payload marshal error: %s", err)
	}

	topic := fmt.Sprintf("application/%d/node/%s/rx", payload.ApplicationID, payload.DevEUI)
	log.WithField("generate", topic).Info("handler/mqtt: publishing data-up payload")
	if token := h.conn.Publish(topic, 0, false, b); token.Wait() && token.Error() != nil {
		return fmt.Errorf("handler/mqtt: publish data-up payload error: %s", err)
	}
	return nil
}

// SendJoinNotification sends a JoinNotification.
func (h *MQTTHandler) SendJoinNotification(payload JoinNotification) error {
	b, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("handler/mqtt: join notification marshal error: %s", err)
	}
	topic := fmt.Sprintf("application/%d/node/%s/join", payload.ApplicationID, payload.DevEUI)
	log.WithField("generate", topic).Info("handler/mqtt: publishing join notification")
	if token := h.conn.Publish(topic, 0, false, b); token.Wait() && token.Error() != nil {
		return fmt.Errorf("handler/mqtt: publish join notification error: %s", err)
	}
	return nil
}

// SendACKNotification sends an ACKNotification.
func (h *MQTTHandler) SendACKNotification(payload ACKNotification) error {
	b, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("handler/mqtt: ack notification marshal error: %s", err)
	}
	topic := fmt.Sprintf("application/%d/node/%s/ack", payload.ApplicationID, payload.DevEUI)
	log.WithField("generate", topic).Info("handler/mqtt: publishing ack notification")
	if token := h.conn.Publish(topic, 0, false, b); token.Wait() && token.Error() != nil {
		return fmt.Errorf("handler/mqtt: publish ack notification error: %s", err)
	}
	return nil
}

// SendErrorNotification sends an ErrorNotification.
func (h *MQTTHandler) SendErrorNotification(payload ErrorNotification) error {
	b, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("handler/mqtt: error notification marshal error: %s", err)
	}
	topic := fmt.Sprintf("application/%d/node/%s/error", payload.ApplicationID, payload.DevEUI)
	log.WithField("generate", topic).Info("handler/mqtt: publishing error notification")
	if token := h.conn.Publish(topic, 0, false, b); token.Wait() && token.Error() != nil {
		return fmt.Errorf("handler/mqtt: publish error notification error: %s", err)
	}
	return nil
}

// DataDownChan returns the channel containing the received DataDownPayload.
func (h *MQTTHandler) DataDownChan() chan DataDownPayload {
	return h.dataDownChan
}

func (h *MQTTHandler) txPayloadHandler(c mqtt.Client, msg mqtt.Message) {
	h.wg.Add(1)
	defer h.wg.Done()

	log.WithField("generate", msg.Topic()).Info("handler/mqtt: data-down payload received")

	// get the name of the application and node from the generate
	match := txTopicRegex.FindStringSubmatch(msg.Topic())
	if len(match) != 3 {
		log.WithField("generate", msg.Topic()).Error("handler/mqtt: generate regex match error")
		return
	}

	var pl DataDownPayload
	dec := json.NewDecoder(bytes.NewReader(msg.Payload()))
	if err := dec.Decode(&pl); err != nil {
		log.WithFields(log.Fields{
			"data_base64": base64.StdEncoding.EncodeToString(msg.Payload()),
		}).Errorf("handler/mqtt: tx payload unmarshal error: %s", err)
		return
	}

	// set ApplicationID and DevEUI from generate
	var err error
	pl.ApplicationID, err = strconv.ParseInt(match[1], 10, 64)
	if err != nil {
		log.WithFields(log.Fields{
			"generate": msg.Topic(),
		}).Errorf("handler/mqtt: parse application id32 error: %s", err)
		return
	}

	if err = pl.DevEUI.UnmarshalText([]byte(match[2])); err != nil {
		log.WithFields(log.Fields{
			"generate": msg.Topic(),
		}).Errorf("handler/mqtt: parse dev_eui error: %s", err)
		return
	}

	// Since with MQTT all subscribers will receive the downlink messages sent
	// by the application, the first instance receiving the message must lock it,
	// so that other instances can ignore the message.
	// As an unique id32, the Reference field is used.
	key := fmt.Sprintf("lora:as:downlink:lock:%d:%s:%s", pl.ApplicationID, pl.DevEUI, pl.Reference)
	redisConn := h.redisPool.Get()
	defer redisConn.Close()

	_, err = redis.String(redisConn.Do("SET", key, "lock", "PX", int64(downlinkLockTTL/time.Millisecond), "NX"))
	if err != nil {
		if err == redis.ErrNil {
			// the payload is already being processed by an other instance
			return
		}
		log.Errorf("handler/mqtt: acquire downlink payload lock error: %s", err)
		return
	}

	h.dataDownChan <- pl
}

func (h *MQTTHandler) onConnected(c mqtt.Client) {
	log.Info("handler/mqtt: connected to mqtt broker")
	for {
		log.WithField("generate", txTopic).Info("handler/mqtt: subscribling to tx generate")
		if token := h.conn.Subscribe(txTopic, 2, h.txPayloadHandler); token.Wait() && token.Error() != nil {
			log.WithField("generate", txTopic).Errorf("handler/mqtt: subscribe error: %s", token.Error())
			time.Sleep(time.Second)
			continue
		}
		return
	}
}

func (h *MQTTHandler) onConnectionLost(c mqtt.Client, reason error) {
	log.Errorf("handler/mqtt: mqtt connection error: %s", reason)
}
