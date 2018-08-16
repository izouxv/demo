package test02

import (
	"github.com/gorilla/websocket"
)

var Manager = ClientManager{
	clients		:	make(map[*Client]bool),
	broadcast	:	make(chan []byte),
	register	:	make(chan *Client),
	unregister	:	make(chan *Client)}

type ClientManager struct {
	clients map[*Client]bool
	broadcast chan []byte
	register chan *Client
	unregister chan *Client
}

type Client struct {
	id string
	socket *websocket.Conn
	send chan []byte
}

type Message struct {
	Sender		string `json:"sender,omitempty"`
	Recipient	string `json:"recipient,omitempty"`
	Content		string `json:"content,omitempty"`
}
