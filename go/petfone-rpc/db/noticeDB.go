package db

import (
	"time"
	"petfone-rpc/core"
	log "github.com/cihub/seelog"
)

type NoticePo struct {
	Id			int32		`gorm:"column:id;primary_key;unique"`
	From    	int32		`gorm:"column:froms"`
	To    		int32		`gorm:"column:tos"`
	State    	int32		`gorm:"column:n_state"`	//3（待确认）-1（确认）
	Types 		int32		`gorm:"column:types"`	//1.分享设备通知
	Info     	string		`gorm:"column:info"`
	CreationTime time.Time	`gorm:"column:creation_time"` 	//创建时间
	UpdateTime   time.Time	`gorm:"column:update_time"`		//修改时间
	DataState	int32		`gorm:"column:data_state"`    	//数据状态
}

func (this *NoticePo) SetNoticePo() error {
	log.Info("SetNoticePo-this:", this)
	return core.MysqlClient.Table("notice").Create(this).Error
}

func (this *NoticePo) DeleteNoticePo() error {
	log.Info("SetNoticePo-this:", this)
	return core.MysqlClient.Table("notice").Where("id = ? and (froms = ? OR tos = ?)",this.Id, this.To).
		Update(this).Error
}

func (this *NoticePo) UpdateNoticePo() int64 {
	log.Info("SetNoticePo-this:", this)
	return core.MysqlClient.Table("notice").Where("id = ? and froms = ? and tos = ? and data_state = ?",
		this.Id, this.From, this.To, 1).Update(this).RowsAffected
}

func (this *NoticePo) GetNoticePo() ([]NoticePo,error) {
	log.Info("GetNoticePo-this:", this)
	var notices []NoticePo
	err := core.MysqlClient.Table("notice").
		Where("tos = ? AND data_state = ?",this.To, 1).Find(&notices).Error
	if err != nil {
		log.Info("GetNoticePo-err:", err)
		return notices,err
	}
	return notices,nil
}

