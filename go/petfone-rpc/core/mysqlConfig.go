package core

import (
	logger "github.com/cihub/seelog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"time"
	"fmt"
	"strconv"
)

var (
	MysqlClient *gorm.DB
	err         error
)

//Mysql客户端连接池
func NewMysqlClient(DBHost, DBName, DBUser string, DBPassword interface{}, DBPort int) {
	args := DBUser + ":" + fmt.Sprint(DBPassword) + "@tcp(" + DBHost + ":" + strconv.Itoa(DBPort) + ")/" + DBName + "?charset=utf8mb4&loc=Asia%2FShanghai&parseTime=true"
	logger.Info("NewMysqlClient-args:",args)
	MysqlClient, err = gorm.Open("mysql", args)
	MysqlClient.LogMode(true)
	MysqlClient.SetLogger(log.New(os.Stdout,"",log.LstdFlags))
	if err != nil {
		logger.Error("NewMysqlClient-err open mysql:", err)
		panic(err)
	}
	MysqlClient.DB().SetMaxIdleConns(20)
	MysqlClient.DB().SetMaxOpenConns(50)
	MysqlClient.DB().SetConnMaxLifetime(time.Hour)
	logger.Info("NewMysqlClient-init")
}

//结束mysql连接
func MysqlClose() {
	logger.Info("MysqlClose:",MysqlClient.Close())
}
