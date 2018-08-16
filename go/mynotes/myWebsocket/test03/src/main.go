package main

import (
	"flag"
	"mynotes/myWebsocket/test03/src/gopool"
	"mynotes/myWebsocket/test03/src/chat"
	"log"
	"net"
	"time"
	"github.com/gobwas/ws"
	"github.com/mailru/easygo/netpoll"
	"net/http"
	_ "net/http/pprof"
)

var (
	addr      = flag.String("listen", ":3333", "address to bind to")
	debug     = flag.String("pprof", ":3334", "address for pprof httpReq")
	workers   = flag.Int("workers", 128, "max workers count")
	queue     = flag.Int("queue", 1, "workers task queue size")
	ioTimeout = flag.Duration("io_timeout", time.Millisecond*100, "i/o operations timeout")
)

func main() {
	flag.Parse()
	if x := *debug; x != "" {
		log.Printf("starting pprof server on %s", x)
		go func() {
			log.Printf("pprof server error: %v", http.ListenAndServe(x, nil))
		}()
	}
	//初始化netpoll实例. 使用netpoll侦听进来的连接
	poller, err := netpoll.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	var (
		// 创建X大的池pool，Y长度的工作队列，一个预生成的goroutine池
		pool = gopool.NewPool(*workers, *queue, 1)
		achat = chat.NewChat(pool)
		exit = make(chan struct{})
	)
	//handle是新连接的处理函数
	//这个函数将tcp连接升级到websocket连接，在连接上注册netpoll监听器，并将其存储为聊天实例中的聊天用户，将在accept()下的循环中调用。
	handle := func(conn net.Conn) {
		// NOTE: we wrap conn here to show that ws could work with any kind of
		// io.ReadWriter.
		safeConn := deadliner{conn, *ioTimeout}
		// Zero-copy upgrade to WebSocket connection.
		hs, err := ws.Upgrade(safeConn)
		if err != nil {
			log.Printf("%s: upgrade error: %v", nameConn(conn), err)
			conn.Close()
			return
		}
		log.Printf("%s: established websocket connection: %+v", nameConn(conn), hs)
		// Register incoming user in chat.
		user := achat.Register(safeConn)
		// Create netpoll event descriptor for conn.
		// We want to handle only read events of it.
		desc := netpoll.Must(netpoll.HandleRead(conn))
		// Subscribe to events about conn.
		poller.Start(desc, func(ev netpoll.Event) {
			if ev&(netpoll.EventReadHup|netpoll.EventHup) != 0 {
				// When ReadHup or Hup received, this mean that client has
				// closed at least write end of the connection or connections
				// itself. So we want to stop receive events about such conn
				// and remove it from the chat registry.
				poller.Stop(desc)
				achat.Remove(user)
				return
			}
			// Here we can read some new message from connection.
			// We can not read it right here in callback, because then we will
			// block the poller's inner loop.
			// We do not want to spawn a new goroutine to read single message.
			// But we want to reuse previously spawned goroutine.
			pool.Schedule(func() {
				if err := user.Receive(); err != nil {
					// When receive failed, we can only disconnect broken
					// connection and stop to receive events about it.
					poller.Stop(desc)
					achat.Remove(user)
				}
			})
		})
	}
	//创建监听器连接
	ln, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("websocket is listening on %s", ln.Addr().String())
	//为监听器创建netpoll描述 ，用OneShot手动恢复事件流
	acceptDesc := netpoll.Must(netpoll.HandleListener(
		ln, netpoll.EventRead|netpoll.EventOneShot,
	))
	//accept是接收signal信号的通道，用来接收Accept()连接connection返回的结果。
	accept := make(chan error, 1)
	//订阅监听器Listener的事件
	poller.Start(acceptDesc, func(e netpoll.Event) {
		// We do not want to accept incoming connection when goroutine pool is
		// busy. So if there are no free goroutines during 1ms we want to
		// cooldown the server and do not receive connection for some short
		// time.
		err := pool.ScheduleTimeout(time.Millisecond, func() {
			conn, err := ln.Accept()
			if err != nil {
				accept <- err
				return
			}
			accept <- nil
			handle(conn)
		})
		if err == nil {
			err = <-accept
		}
		if err != nil {
			if err != gopool.ErrScheduleTimeout {
				goto cooldown
			}
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				goto cooldown
			}
			log.Fatalf("accept error: %v", err)
		cooldown:
			delay := 5 * time.Millisecond
			log.Printf("accept error: %v; retrying in %s", err, delay)
			time.Sleep(delay)
		}
		poller.Resume(acceptDesc)
	})
	<-exit
}

func nameConn(conn net.Conn) string {
	return conn.LocalAddr().String() + " > " + conn.RemoteAddr().String()
}

// deadliner is a wrapper around net.Conn that sets read/write deadlines before
// every Read() or Write() call.
type deadliner struct {
	net.Conn
	t time.Duration
}

func (d deadliner) Write(p []byte) (int, error) {
	if err := d.Conn.SetWriteDeadline(time.Now().Add(d.t)); err != nil {
		return 0, err
	}
	return d.Conn.Write(p)
}

func (d deadliner) Read(p []byte) (int, error) {
	if err := d.Conn.SetReadDeadline(time.Now().Add(d.t)); err != nil {
		return 0, err
	}
	return d.Conn.Read(p)
}
