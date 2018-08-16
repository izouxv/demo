package core

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/cihub/seelog"
	"time"
)

const (
	mysqlMaxIdle = 10
	mysqlMaxOpen = 20
)


func OpenDatabase(dsn string) (*gorm.DB, error) {
	log.Infof("dsn:%s ",dsn)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Errorf("database connection error: %s",err)
		return nil,  err
	}
	for {
		if err = db.DB().Ping(); err != nil {
			log.Errorf("ping database error, will retry in 2s: %s", err)
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}
	db.LogMode(true)
	db.DB().SetMaxIdleConns(mysqlMaxIdle)
	db.DB().SetMaxOpenConns(mysqlMaxOpen)
	return db, nil
}
