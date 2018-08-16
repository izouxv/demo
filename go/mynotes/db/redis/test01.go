package main

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)

func main() {
	RedisPollInit()
	defer RedisPoll.Close()
	conn := RedisPoll.Get()
	defer conn.Close()
	if err := conn.Send("geoadd",""); err!= nil {
		fmt.Println("err:",err)
		return
	}
	if err := conn.Flush(); err!= nil {
		fmt.Println("err:",err)
		return
	}
	fmt.Print("Receive:")
	fmt.Println(conn.Receive())
}

var RedisPoll *redis.Pool

func RedisPollInit() {
	fmt.Println("aaa:",RedisPoll)
	if RedisPoll != nil {
		return
	}
	RedisPoll = &redis.Pool{
		MaxIdle:3,
		MaxActive:8,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp","192.168.1.6:6379")
			if err != nil {
				return nil,err
			}
			result,err := redis.String(conn.Do("AUTH","radacat1234"))
			if err != nil {
				conn.Close()
				return nil,err
			}
			fmt.Println("result:",result)
			return conn,nil
		},
	}
}
