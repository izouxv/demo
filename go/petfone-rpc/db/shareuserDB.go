package db

import (
	log "github.com/cihub/seelog"
	"petfone-rpc/core"
	"time"
	"github.com/jinzhu/gorm"
)


//用户分享结构体
type ShareUserPo struct {
	Id          	int32		`gorm:"column:id;primary_key;unique"`//id
	Uid          	int32		`gorm:"column:uid"`					//用户id
	Fuid	        int32		`gorm:"fuid"`						//对方用户id
	CreationTime 	time.Time	`gorm:"column:creation_time"` 		//创建时间
	UpdateTime   	time.Time 	`gorm:"column:update_time"`   		//修改时间
	DataState    	int32		`gorm:"column:data_state"`    		//数据状态
}

//添加分享
func (this *ShareUserPo) SetShareDB(dbc *gorm.DB) error {
	log.Info("SetShareDB:", this)
	return dbc.Table("share_user").Create(&this).Error
}

//修改分享
func (this *ShareUserPo) UpdateShareDB(dbc *gorm.DB) error {
	log.Info("UpdateShareDB:", this)
	return dbc.Table("share_user").Where("id = ?", this.Id).
		Updates(map[string]interface{}{"update_time":this.UpdateTime}).Error
}

//删除分享
func (this *ShareUserPo) DeleteShareDB(dbc *gorm.DB) int64 {
	log.Info("DeleteShareDB-this:", this)
	return dbc.Table("share_user").Where("uid = ? AND fuid = ? AND data_state = 1", this.Uid, this.Fuid).
		Updates(map[string]interface{}{"update_time":this.UpdateTime,"data_state":2}).RowsAffected
}

//查询分享
func (this *ShareUserPo) GetShareDB() (err error) {
	log.Info("GetShareDB:", this)
	return core.MysqlClient.Table("share_user").Where("uid = ? AND fuid = ? AND data_state = 1",
		this.Uid, this.Fuid).First(&this).Error
}

//查询分享
func (this *ShareUserPo) GetSharesDB() (shares []*ShareUserPo,err error) {
	log.Info("GetSharesDB:", this)
	return shares,core.MysqlClient.Table("share_user").Where("(uid = ? or fuid = ?) AND data_state = 1",
		this.Uid, this.Fuid).Find(&shares).Error
}
