package storage

import (
	"time"

	. "account-domain-rpc/common"

	"account-domain-rpc/module"

	log "github.com/cihub/seelog"
)

type Advertisement struct {
	Id         int32     `gorm:"column:id;primary_key;unique"`
	Tid        int64     `gorm:"column:tid"`
	Name       string    `gorm:"column:name"`
	Md5        string    `gorm:"column:md5"`
	StartTime  time.Time `gorm:"column:start_time"`
	EndTime    time.Time `gorm:"column:end_time"`
	Advertiser string    `gorm:"column:advertiser"`
	State      int32     `gorm:"column:state"`
	AdvUrl     string    `gorm:"column:adv_url"`
	FileUrl    string    `gorm:"column:file_url"`
	FileName   string    `gorm:"column:file_name"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

//添加新广告
func (a *Advertisement) NewAdvertisement() error {
	if err := module.MysqlClient().Table("tbl_advertisement").Create(a).Error; err != nil {
		log.Error("New Advertisement to db error", err)
		return err
	}
	return nil
}

//修改广告
func (a *Advertisement) UpdateAdvertisement() error {
	r := module.MysqlClient().Table("tbl_advertisement").Where("id = ?  and tid = ? and state = 1", a.Id,a.Tid).Update(a)
	if r.Error != nil {
		log.Error("Update advertisement info to db error", r.Error)
	}
	if r.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return nil
}

//app获取最新广告
func (a *Advertisement) GetLatestdAdvertisement(tid int64) error {
	r := module.MysqlClient().Table("tbl_advertisement").
		Where("tid = ? and state = 1 ", tid).Order("id desc ").First(a)
	if r.Error != nil {
		log.Error("Get advertisement from db error", r.Error)
		return  r.Error
	}
	return nil
}


func (a *Advertisement) GetAdvertisement() error {
	r := module.MysqlClient().Table("tbl_advertisement").
		Where("tid = ? and id = ? ", a.Tid,a.Id).Find(&a)
	if r.Error != nil {
		if r.RowsAffected == 0 {
			return ErrDoesNotExist
		}
		log.Error("Get one advertisement from db error", r.Error)
		return  r.Error
	}
	return nil
}


//批量获取广告信息
func (a *Advertisement) GetAllAdvertisement(page, count int32) ([]Advertisement, int32, error) {
	var totalCount int32
	if err := module.MysqlClient().Table("tbl_advertisement").Where("tid = ? and state = 1",a.Tid).Count(&totalCount).Error; err != nil {
		log.Error("GetAllAdvertisement err,", err)
		return nil, 0, err
	}
	if count == 0 {
		count = totalCount
	}
	advs := make([]Advertisement, 0, totalCount)
	if err := module.MysqlClient().Table("tbl_advertisement").Where("tid = ? and state = 1",a.Tid).Order("id desc").Limit(count).Offset((page - 1) * count).Find(&advs).Error; err != nil {
		log.Error("Get advertisement from db error", err)
		return nil, 0, err
	}
	return advs, totalCount, nil
}

//删除广告信息
func (a *Advertisement) DelAdvertisement() error {
	r := module.MysqlClient().Table("tbl_advertisement").Where("tid = ? and id = ?",a.Tid,a.Id).Update(map[string]interface{}{"state": 2})
	if r.Error != nil {
		log.Error("Delete advertisement info to db error", r.Error)
	}
	if r.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return nil
}
