package main

import (
	log "github.com/cihub/seelog"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func main() {
	NewMongoDB()
	collection := mClient.session.DB("test").C("people")
	err := collection.Insert(&Person{"superWang", "13478808311"}, &Person{"David", "15040268074"})
	if err != nil {
		log.Error(err)
	}
	result := Person{}
	err = collection.Find(bson.M{"name": "superWang"}).One(&result)
	if err != nil {
		log.Error(err)
	}
	fmt.Println("Name:", result.Name)
	fmt.Println("Phone:", result.Phone)
}

type Person struct {
	Name  string
	Phone string
}


// MongoDB 数据源结构体
type MongoDB struct {
	session *mgo.Session
	addr    string
	db      string
}

var (
	mClient *MongoDB
	addr = "192.168.1.6:27017"
	db = "test"
)

// NewMongoDB 创建MongoDB数据源实例，相当于构造方法
func NewMongoDB() {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:     []string{addr},
		Timeout:   10 * time.Second,
		PoolLimit: 4096,
	}
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Error("dial mongodb failed err:%s\n", err)
	}
	session.SetMode(mgo.Monotonic, true)
	mClient = &MongoDB{session: session, addr: addr, db: db}
}