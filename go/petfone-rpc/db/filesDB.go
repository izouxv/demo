package db

import (
	"time"
	"petfone-rpc/core"
	log "github.com/cihub/seelog"
)

type FilesPo struct {
	Md5			string		`gorm:"column:md5"`
	Uid	    	int32		`gorm:"column:uid"`
	Address     string		`gorm:"column:address"`
	CreationTime time.Time	`gorm:"column:creation_time"` 	//创建时间
}

func (this *FilesPo) SetFile() error {
	log.Info("SetFile-this:", this)
	return core.MysqlClient.Table("files").Create(this).Error
}

func (this *FilesPo) GetFile() error {
	log.Info("GetFile-this:", this)
	err := core.MysqlClient.Table("files").
		Where("md5 = ?",this.Md5).
			Select("md5, uid, address, creation_time").
				First(this).Error
	log.Info("files:",this)
	if err != nil {
		log.Info("GetFile-err:", err)
		return err
	}
	return nil
}

