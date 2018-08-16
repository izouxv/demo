package test02

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/satori/go.uuid"
	"github.com/gorilla/websocket"
	"testing"
)

//
func TestServer_ws(t *testing.T) {
	fmt.Println("Starting application...")
	go Manager.start()
	http.HandleFunc("/ws", wsPage)
	//绑定效果页面
	http.Handle("/", http.StripPrefix("/",http.FileServer(http.Dir("h5"))))
	http.Handle("/static/", http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":12345", nil)
}

func h_index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("11111111")
	http.ServeFile(w, r, "index.html")
}

//将http连接升级为ws连接
func wsPage(res http.ResponseWriter, req *http.Request) {
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
	}}).Upgrade(res, req, nil)
	if err != nil {
		http.NotFound(res, req)
		return
	}
	uid,_ := uuid.NewV4()
	client := &Client{id: uid.String(), socket: conn, send: make(chan []byte)}

	Manager.register <- client

	go client.read()
	go client.write()
}

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.register:
			manager.clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
			manager.send(jsonMessage, conn)
		case conn := <-manager.unregister:
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
				manager.send(jsonMessage, conn)
			}
		case message := <-manager.broadcast:
			for conn := range manager.clients {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(manager.clients, conn)
				}
			}
		}
	}
}

func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.clients {
		if conn != ignore {
			conn.send <- message
		}
	}
}

func (c *Client) read() {
	defer func() {
		Manager.unregister <- c
		c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			Manager.unregister <- c
			c.socket.Close()
			break
		}
		fmt.Println("id:",c.id)
		jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
		Manager.broadcast <- jsonMessage
	}
}

func (c *Client) write() {
	defer func() {
		c.socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

