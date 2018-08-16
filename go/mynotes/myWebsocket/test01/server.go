package main

import (
	"net/http"
	"golang.org/x/net/websocket"
	"fmt"
)

func main() {
	http.Handle("/echo",websocket.Handler(Echo))
	err := http.ListenAndServe("localhost:8888",nil)
	fmt.Println(err)
}

func Echo(ws *websocket.Conn) {
	var str string
	for ws.IsClientConn() {
		fmt.Println("11111")
		ws.Read([]byte(str))
		fmt.Println("str:",str)
		ws.Write([]byte("111111"))
		a,_ := ws.NewFrameReader()
		ws.HandleFrame(a)
	}
}
