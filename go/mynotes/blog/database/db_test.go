package database

import (
	"testing"
	"mynotes/blog/config"
	"mynotes/blog/api"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	_ "mynotes/blog/initContext"
	"mynotes/blog/logger"
	"time"
	"encoding/json"
)

func Test_TestMongo(t *testing.T) {
	logger.Debug("aaa",time.Now())
	NewMongoDB(config.Config.DataBases["mongo"])
	//mongodb.SetMongo(&api.Person{"superWang", "13478808311"}, &api.Person{"David", "15040268074"})
	person := &api.Person{Name:"superWang"}
	data,err := bson.Marshal(person)
	if err != nil {
		fmt.Println("111",err)
		return
	}
	var b bson.M
	err = bson.Unmarshal(data,&b)
	if err != nil {
		fmt.Println("222",err)
		return
	}
	fmt.Println("bbb:",b)
	b, err = Mongo.GetMongoOne(b)
	fmt.Println("bbb:",b)
	data,err = json.Marshal(b)
	err = json.Unmarshal(data,&person)
	if err != nil {
		fmt.Println("333",err)
		return
	}
	fmt.Println("aaa:",person)
}

func Test_struct(t *testing.T) {
	person1 := &api.Person{Name:"superWang",Age:13}
	rt := reflect.TypeOf(*person1)
	rv := reflect.ValueOf(*person1)
	for i:=0; i<rt.NumField(); i++ {
		f := rt.Field(i)
		key := f.Tag.Get("bson")
		value := rv.FieldByName(f.Name).String()
		fmt.Println(key,value,f.Name)
	}
}


