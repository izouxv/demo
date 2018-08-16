package main

import (
	"net"
	"fmt"
	"io"
	"crypto/tls"
	"io/ioutil"
	"encoding/pem"
	"crypto/x509"
	"bytes"
	"encoding/gob"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/rand"
)

const clientBufLength = 12
var data = ([]byte)("hahahahahah")
var clientRsaLabel = []byte("server")
func main() {
	//tcpClient()
	//tlsTcpClient()
	//udpClient()
	tlsUdpClient()
	//tlsUdp1()
}
//go run $GOROOT/src/crypto/tls/generate_cert.go -ca=true -host="localhost"
func tlsTcpClient()  {
	var config tls.Config
	if certificate, e := tls.LoadX509KeyPair("./tcpIp/test01/cert.pem", "./tcpIp/test01/key.pem"); e ==nil {
		config = tls.Config{
			Certificates: []tls.Certificate{certificate},
			InsecureSkipVerify:true}
	}
	if conn, err := tls.Dial("tcp", "localhost:1024",&config); err == nil {
		defer conn.Close()
		n,err := conn.Write([]byte("hahahahaha"))
		fmt.Println("write:",n,err)
		data := make([]byte, 0)
		buf := make([]byte, clientBufLength)
		for {
			n, err := conn.Read(buf)
			if err != nil && err != io.EOF {
				fmt.Println("err:",err)
			}
			data = append(data, buf[:n]...)
			if n != clientBufLength {
				break
			}
		}
		fmt.Println("data:",string(data))
	} else {
		fmt.Println("error:", err)
	}
}
func tcpClient()  {
	if conn, e := net.Dial("tcp", "localhost:1024"); e == nil {
		defer conn.Close()
		n,err := conn.Write(data)
		fmt.Println("write:",n,err)
		data := make([]byte, 0)//输入缓冲以免数据过长读取到不完整的数据
		buf := make([]byte, clientBufLength)
		for {
			fmt.Println("aaaa")
			n, err := conn.Read(buf)
			if err != nil && err != io.EOF {
				fmt.Println("err:",err)
			}
			data = append(data, buf[:n]...)
			if n != clientBufLength {
				break
			}
			fmt.Println("aaaa")
		}
		fmt.Println("data:",string(data))
	} else {
		fmt.Println("error", e)
	}
}

func udpClient()  {
	if addr, err := net.ResolveUDPAddr("udp", "localhost:1026"); err == nil {
		if conn, err := net.DialUDP("udp", nil, addr); err == nil {
			defer conn.Close()
			n, err := conn.Write(data)
			fmt.Println("write:",n,err)
			data := make([]byte, 0)
			buf := make([]byte, clientBufLength)
			for {
				fmt.Println("Read1:")
				n, remoteAddr,err := conn.ReadFromUDP(buf)
				fmt.Println("Read:",n, err,",buf:",buf)
				if err != nil && err != io.EOF {
					fmt.Println("err:",err)
				}
				data = append(data, buf[:n]...)
				if n != clientBufLength {
					break
				}
				fmt.Println("Read2:",remoteAddr)
			}
			fmt.Println("data:",string(data))
		} else {
			fmt.Println("error:", err)
		}
	} else {
		fmt.Println("error:", err)
	}
}

func tlsUdp1() {
	Connect("localhost:1025", func(server *net.UDPConn, priKey *rsa.PrivateKey) {
		cipherText := make([]byte, 1024)
		if n, e := server.Read(cipherText); e == nil {
			fmt.Println("cipher text:", string(cipherText[:n]))
			// cipherText是加密的数据，需要解密
			if plainText, e := rsa.DecryptOAEP(sha1.New(), rand.Reader, priKey, cipherText[:n], clientRsaLabel); e == nil {
				fmt.Println("receive decrypt string:", string(plainText))
			}
		}
	})
}
func Connect(address string, f func(*net.UDPConn, *rsa.PrivateKey)) {
	LoadPrivateKey("./tcpIp/test01/key.pem", func(pk *rsa.PrivateKey) { // pk is private key
		if addr, e := net.ResolveUDPAddr("udp", address); e ==nil {
			if server, e := net.DialUDP("udp", nil, addr); e == nil {
				defer server.Close()
				SendKey(server, pk.PublicKey, func() {
					f(server, pk)
				})
			} else {
				fmt.Println("err1", e)
			}
		} else {
			fmt.Println("err2", e)
		}
	})
}
func LoadPrivateKey(file string, f func(*rsa.PrivateKey)) {
	if file, e := ioutil.ReadFile(file); e == nil {
		if block, _ := pem.Decode(file); block != nil {
			if block.Type == "RSA PRIVATE KEY" {
				if key, _ := x509.ParsePKCS1PrivateKey(block.Bytes); key !=nil {
					fmt.Println(key)
					f(key)
				}
			}
		}
	} else {
		fmt.Println(e)
	}
	return
}
func SendKey(server *net.UDPConn, publicKey rsa.PublicKey, f func()) {
	var encodedKey bytes.Buffer
	if e := gob.NewEncoder(&encodedKey).Encode(publicKey); e == nil {
		fmt.Println("aaa:", len(encodedKey.Bytes()))
		if _, e = server.Write(encodedKey.Bytes()); e == nil {
			f()
		}
	}
}

func tlsUdpClient()  {
	key := genKey()
	addr, err := net.ResolveUDPAddr("udp", "localhost:1025")
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	server, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	defer server.Close()
	fmt.Println("qqq:", len(getKeyBs(key.PublicKey).Bytes()))
	n, err := server.Write(getKeyBs(key.PublicKey).Bytes())
	fmt.Println("a:",n)
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	//server.Write([]byte("aaaaaaaaaa"))
	cipherText := make([]byte, 1000)
	n, err = server.Read(cipherText)
	fmt.Println("aa:",n)
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	// cipherText是加密的数据，需要解密
	plainText, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, key, cipherText[:n], clientRsaLabel)
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	fmt.Println("receive decrypt string:", string(plainText))
}
func getKeyBs(key rsa.PublicKey) *bytes.Buffer {
	var encodedKey bytes.Buffer
	err := gob.NewEncoder(&encodedKey).Encode(key)
	if err != nil {
		fmt.Println("err:",err)
		return nil
	}
	return &encodedKey
}
func genKey() *rsa.PrivateKey {
	file, err := ioutil.ReadFile("./tcpIp/test01/key.pem")
	if err != nil {
		fmt.Println("err:",err)
		return nil
	}
	block, _ := pem.Decode(file)
	if block == nil {
		fmt.Println("err block nil")
		return nil
	}
	if block.Type != "RSA PRIVATE KEY" {
		fmt.Println("err block.Type:",block.Type)
		return nil
	}
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	fmt.Println("key:",err,key)
	if key == nil || err != nil {
		fmt.Println("err ParsePKCS1PrivateKey:",key,err)
		return nil
	}
	return key
}

