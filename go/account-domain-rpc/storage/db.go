package storage

import (
	log "github.com/cihub/seelog"
	"github.com/jinzhu/gorm"
)

// Transaction wraps the given function in a transaction. In case the given
// functions returns an error, the transaction will be rolled back.
func Transaction(db *gorm.DB, f func(tx *gorm.DB) error) error {
	tx := db.Begin()
	log.Info("tx begin")
	err := f(tx)
	if err != nil {
		log.Errorf("func error is %s,tx rollback", err)
		if tx.Rollback().Error != nil {
			log.Errorf("tx rollback Error %s", err)
		}
		return err
	}
	if err = tx.Commit().Error; err != nil {
		log.Errorf("tx commit error %s", err)
		return err
	}
	log.Info("tx commit successful ")
	return nil
}
