package module

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"

	log "github.com/cihub/seelog"
)

var mysqlClient *gorm.DB

func MysqlClient() *gorm.DB {
	mysqlClient.LogMode(true)
	return mysqlClient
}
func NewMysqlClient(host, port, name, user, password string) {
	var err error
	dbUrl := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8&loc=Asia%2FShanghai&parseTime=true"
	log.Info("New MysqlClient", dbUrl)
	mysqlClient, err = gorm.Open("mysql", dbUrl)
	if err != nil {
		log.Error("Connected mysql failed:", err)
		panic(err)
	}
	mysqlClient.DB().SetMaxIdleConns(10)
	mysqlClient.DB().SetMaxOpenConns(20)
}

func SetLog(logger gorm.Logger) {
	mysqlClient.SetLogger(logger)
}
