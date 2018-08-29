package po

import (
	"github.com/gorilla/websocket"
	"cotx-http/pb"
)

type Websocket struct {
	Code     int32          `json:"code"`
	Msg      string         `json:"msg"`
	Result   interface{}   `json:"result"`
}

type  HeartBeat struct {
	Token   string   `json:"token"`
	Source  string   `json:"source"`
}

type WebScoketInfo struct {
	WebsocketId  string
	Conn         *websocket.Conn
	Token        string
	UserId       int32
}

type WebscoketResult struct {
	GatewayId  string   `json:"gatewayid"`
	Code       int32    `json:"code"`
	Describe   string   `json:"describe"`
	State      int32    `json:"state"`
	Value      int32    `json:"value"`
}
type WebSocketMessgage struct {
	Conn    *websocket.Conn
	Message []*pb.PushMessage
}
type WebscoketChan struct {
	Websocketchan chan map[string]WebSocketMessgage
	Websocketmap   map[string]WebSocketMessgage
}
