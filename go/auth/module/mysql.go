package module

import (
	log "github.com/cihub/seelog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
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
		return nil, fmt.Errorf("database connection error: %s", err)
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
