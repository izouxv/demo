package main

import (
	"net"
	"fmt"
	"io"
	"crypto/rsa"
	"encoding/gob"
	"bytes"
	"crypto/sha1"
	"crypto/rand"
)

/**
DialUDP 是 pre-connected 其实是维持了一个发送地址
ListenUDP 是 unconnect
如果*UDPConn是connected,读写方法是Read和Write。
如果*UDPConn是unconnected,读写方法是ReadFromUDP和WriteToUDP（以及ReadFrom和WriteTo)。
如果使用dail
你将失去SetKeepAlive或TCPConn和UDPConn的SetReadBuffer 这些函数， 除非做类型转换
*/

const udpServerBufLength = 1024
func main() {
	//udp()
	tlsUdp()
}

var MSG = []byte("hello go scure udp")
var RsaLabel = []byte("server")

func tlsUdp() {
	if a, err := net.ResolveUDPAddr("udp", ":1025"); err == nil {
		if conn, err := net.ListenUDP("udp", a); err == nil {
			for {
				buffer := make([]byte, 1600)
				n, client, err := conn.ReadFromUDP(buffer)
				if err != nil {
					fmt.Println("err2:",err)
				}
				go func(conn *net.UDPConn,c *net.UDPAddr, b []byte) {
					publicKey := genPublicKey(b[:n])
					if publicKey == nil {
						return
					}
					resp, err := rsa.EncryptOAEP(sha1.New(),rand.Reader,publicKey,MSG,RsaLabel)
					if err != nil {
						fmt.Println("err4:",err)
					}
					n, _ := conn.WriteToUDP(resp, c)
					fmt.Println("MSG:",string(MSG),n)
					if n != 0 {
						fmt.Println(n, "write to ", c)
					}
				}(conn,client,buffer[:n])
			}
		}
	}
}
func genPublicKey(b []byte) *rsa.PublicKey {
	var key rsa.PublicKey
	err := gob.NewDecoder(bytes.NewBuffer(b)).Decode(&key)
	if err != nil {
		fmt.Println("err3:",err)
		return nil
	}
	return &key
}

func udp()  {
	addr, err := net.ResolveUDPAddr("udp", ":1026")
	if err != nil {
		fmt.Println("error1:", err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("error2:", err)
		return
	}
	defer conn.Close()
	for {
		var data []byte
		buf := make([]byte, udpServerBufLength)
		var remoteAddr *net.UDPAddr
		var n int
		for {
			n, remoteAddr, err = conn.ReadFromUDP(buf)
			if err != nil && err != io.EOF {
				fmt.Println("error3:", err)
				return
			}
			data = append(data,buf[:n]...)
			if n != udpServerBufLength {
				break
			}
		}
		fmt.Println("remoteAddr:",remoteAddr)
		n, err := conn.WriteToUDP(data,remoteAddr)
		if err != nil {
			fmt.Println("error4:", err)
			return
		}
		fmt.Println("data:",string(data),n)
	}
}
