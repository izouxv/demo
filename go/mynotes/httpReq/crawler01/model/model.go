package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB

	username string = "root"
	password string = "Radacat2017"
	dbName   string = "reptile"
)

//root:Radacat2017@tcp(mm3306.mysql.radacat.com:3306)/petfone?charset=utf8mb4
func init() {
	var err error
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(192.168.1.6)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, dbName))
	if err != nil {
		log.Fatalf(" gorm.Open.err: %v", err)
	}
	log.Println("init----------")
	DB.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		log.Println("defaultTableName:",defaultTableName)
		return "job"
	}
}
