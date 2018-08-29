package database

import (
	"log"
	"time"
	"gopkg.in/mgo.v2"
	"mynotes/blog/config"
	"mynotes/blog/logger"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

// MongoDB 数据源结构体
type MgoClient struct {
	session 	*mgo.Session
	addr		string
	db			string
	poolLimit	int
}

var (
	Mongo *MgoClient
	mongoDBDialInfo  *mgo.DialInfo
)

//初始化MongoDB数据源
func NewMongoDB(mongo config.Database) error {
	if len(mongo.Addr) <= 0 || len(mongo.Addr) <= 0 {
		return errors.New("mongodb addr or db failed:"+mongo.Addr+mongo.Name)
	}
	mongoDBDialInfo = &mgo.DialInfo{
		Addrs:     []string{mongo.Addr},
		Timeout:   10 * time.Second,
		PoolLimit: mongo.ConnMax,
	}
	MgoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		return errors.New("mongodb DialWithInfo failed:"+err.Error())
	}
	MgoSession.SetMode(mgo.Monotonic, true)
	Mongo = &MgoClient{session:MgoSession,addr:mongo.Addr,db:mongo.Name,poolLimit:mongo.ConnMax}
	return nil
}

func CloseMongo()  {
	if Mongo != nil {
		Mongo.session.Close()
	}
}

func (self *MgoClient) sessionClone() *mgo.Session {
	if self.session == nil {
		MgoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
		if err != nil {
			log.Fatalf("dial mongodb failed err:%s\n", err)
		}
		MgoSession.SetMode(mgo.Monotonic, true)
	}
	return self.session.Clone()
}

func (self *MgoClient) SetMongo(i ...interface{}) error {
	s := self.sessionClone()
	defer s.Close()
	c := s.DB(self.db).C("test")
	err := c.Insert(i...)
	if err != nil {
		logger.Error("TestMongo err:", err)
		return err
	}
	return nil
}

func (self *MgoClient) GetMongoOne(b bson.M) (bson.M,error) {
	s := self.sessionClone()
	defer s.Close()
	c := s.DB(self.db).C("test")
	err := c.Find(b).One(&b)
	if err != nil {
		logger.Error("TestMongo err:", err)
		return b, err
	}
	return b, nil
}

func (self *MgoClient) SetGridFs(prefix,name string,bs []byte) error {
	s := self.sessionClone()
	defer s.Close()
	g := s.DB(self.db).GridFS(prefix)
	f,err := g.Create(name)
	if err != nil {
		logger.Error("TestMongo err:", err)
		return err
	}
	defer f.Close()
	f.SetId(1)
	f.SetName(name)
	if _,err = f.Write(bs); err != nil {
		logger.Error("TestMongo err:", err)
		return err
	}
	return nil
}
