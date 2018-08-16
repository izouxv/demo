package main


import (
	"net"
	"fmt"
	"io"
	"crypto/tls"
	"crypto/rand"
)

/**
 *  tcp 启动 链接 三次握手  关闭四次
 *  慢启动 + 拥塞窗口
 *  门限控制 < 拥塞窗口 进入拥塞避免
 *  在发送数据后 检测 ack 确认超时 或者 收到重复ack 确认 操作门限值 和 拥塞窗口值 来限流
 *  接收方 使用通告窗口 告知发送可接受多少字节
 *
 *  发送方：滑动窗口协议
 */

const tcpServerBufLength = 8
func main() {
	tcp()
	//tlsTcp()
}

func tlsTcp()  {
	var config tls.Config
	if certificate, e := tls.LoadX509KeyPair("./tcpIp/test01/cert.pem", "./tcpIp/test01/key.pem"); e ==nil {
		config = tls.Config{
			Certificates: []tls.Certificate{certificate},
			Rand: rand.Reader,
		}
	}
	if listener, err := tls.Listen("tcp", ":1024",&config); err == nil {
		for {
			if conn, e := listener.Accept(); e==nil {
				go handle(conn)
			}
		}
	} else {
		fmt.Println("error:",err)
	}
}

func tcp()  {
	if listener, e := net.Listen("tcp", ":1024"); e == nil {
		for {
			if conn, e := listener.Accept(); e==nil {
				go handle(conn)
			}
		}
	} else {
		fmt.Println("error")
	}
}

func handle(c net.Conn)  {
	defer c.Close()
	data := make([]byte, 0)//此处做一个输入缓冲以免数据过长读取到不完整的数据
	buf := make([]byte, tcpServerBufLength)
	for {
		n, err := c.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("err:",err)
		}
		data = append(data, buf[:n]...)
		if n != tcpServerBufLength {
			break
		}
	}
	fmt.Println("data:",string(data))
	c.Write(data)
}