package handler

import (
	"encoding/hex"
	"github.com/brocaar/lorawan"
	"time"
)

const (
	JoinRequest = 0x00
	JoinAccept  = 0x01
)

type Payload interface {
	MarshalBinary() (data []byte, err error)
	UnmarshalBinary(data []byte) error
}

type Header byte

// MIC represents the message integrity code.
type MIC [4]byte

// String implements fmt.Stringer.
func (m MIC) String() string {
	return hex.EncodeToString(m[:])
}

// MarshalText implements encoding.TextMarshaler.
func (m MIC) MarshalText() ([]byte, error) {
	return []byte(m.String()), nil
}

//pyyPayload请求
type PHYPayload struct {
	Header  Header  `json:"mhdr"`
	Payload Payload `json:"data"`
	MIC     MIC     `json:"mic"`
}

// DataRate contains the data-rate related fields.
type DataRate struct {
	Modulation   string `json:"modulation"`             //调制
	Bandwidth    int    `json:"bandwidth"`              //带宽
	SpreadFactor int    `json:"spreadFactor,omitempty"` //传播因子
	Bitrate      int    `json:"bitrate,omitempty"`      //比特率
}

// RXInfo contains the RX information.
type RXInfo struct {
	MAC       lorawan.EUI64 `json:"mac"`
	Time      *time.Time    `json:"time,omitempty"`
	RSSI      int           `json:"rssi"`    //接受信号强度
	LoRaSNR   float64       `json:"loRaSNR"` //性躁比
	Name      string        `json:"name"`
	Latitude  float64       `json:"latitude"`  //纬度
	Longitude float64       `json:"longitude"` //精度
	Altitude  float64       `json:"altitude"`
}

// TXInfo contains the TX information.
type TXInfo struct {
	Frequency int      `json:"frequency"` //频率
	DataRate  DataRate `json:"dataRate"`  //数率
	ADR       bool     `json:"adr"`       //自动协调速率
	CodeRate  string   `json:"codeRate"`  //码率
}

// DataUpPayload represents a data-up payload.
type DataUpPayload struct {
	ApplicationID   int64         `json:"applicationID,string"`
	ApplicationName string        `json:"applicationName"`
	NodeName        string        `json:"nodeName"`
	DevEUI          lorawan.EUI64 `json:"devEUI"`
	RXInfo          []RXInfo      `json:"rxInfo"`
	TXInfo          TXInfo        `json:"txInfo"`
	FCnt            uint32        `json:"fCnt"`
	FPort           uint8         `json:"fPort"`
	Data            []byte        `json:"data"`
}

// DataDownPayload represents a data-down payload.
type DataDownPayload struct {
	ApplicationID int64         `json:"applicationID,string"`
	DevEUI        lorawan.EUI64 `json:"devEUI"`
	Reference     string        `json:"reference"`
	Confirmed     bool          `json:"confirmed"`
	FPort         uint8         `json:"fPort"`
	Data          []byte        `json:"data"`
}

// JoinNotification defines the payload sent to the application on
// a JoinNotificationType event.
type JoinNotification struct {
	ApplicationID   int64           `json:"applicationID,string"`
	ApplicationName string          `json:"applicationName"`
	NodeName        string          `json:"nodeName"`
	DevEUI          lorawan.EUI64   `json:"devEUI"`
	DevAddr         lorawan.DevAddr `json:"devAddr"`
}

// ACKNotification defines the payload sent to the application
// on an ACK event.
type ACKNotification struct {
	ApplicationID   int64         `json:"applicationID,string"`
	ApplicationName string        `json:"applicationName"`
	NodeName        string        `json:"nodeName"`
	DevEUI          lorawan.EUI64 `json:"devEUI"`
	Reference       string        `json:"reference"`
}

// ErrorNotification defines the payload sent to the application
// on an error event.
type ErrorNotification struct {
	ApplicationID   int64         `json:"applicationID,string"`
	ApplicationName string        `json:"applicationName"`
	NodeName        string        `json:"nodeName"`
	DevEUI          lorawan.EUI64 `json:"devEUI"`
	Type            string        `json:"type"`
	Error           string        `json:"error"`
}
