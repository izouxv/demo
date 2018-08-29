package module

import (
	log "github.com/cihub/seelog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var mysqlClient *gorm.DB

func MysqlClient() *gorm.DB {
	mysqlClient.LogMode(true)
	return mysqlClient
}

func NewMysqlClient(host, port, name, user, password string, maxIdle, maxOpen int) {
	var err error
	dbUrl := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8&loc=Asia%2FShanghai&parseTime=true"
	log.Info("New MysqlClient", dbUrl)
	mysqlClient, err = gorm.Open("mysql", dbUrl)
	if err != nil {
		log.Error("Connected mysql failed:", err)
		panic(err)
	}
	mysqlClient.DB().SetMaxIdleConns(maxIdle)
	mysqlClient.DB().SetMaxOpenConns(maxOpen)
}
