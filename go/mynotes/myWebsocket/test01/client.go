package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"time"
)

func main() {
	origin := "httpReq://localhost/"
	url := "ws://localhost:8888/echo"
	ws, err := websocket.Dial(url,"",origin)
	fmt.Println("err:",err)
	_,err =ws.Write([]byte("test"))
	ws.SetWriteDeadline(time.Now())
	fmt.Println("err:",err)
	msg := make([]byte,20)
	ws.Read(msg)
	fmt.Println(string(msg))
}
