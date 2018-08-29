package storage

import (
	. "auth/go-drbac/common"
	"github.com/jinzhu/gorm"
	"time"
)

type WeChatPay struct {
	Did         int32  		`gorm:"column:did;primary_key;unique"json:"did"`
	AppId       string  	`gorm:"column:app_id" json:"app_id"`
	MchId  		string  	`gorm:"column:mch_id" json:"mch_id"`
	Key  		string 		`gorm:"column:key" json:"key"`
	AppSecret   string  	`gorm:"column:app_secret" json:"app_secret"`
	CreateTime	time.Time 	`gorm:"column:create_time" json:"create_time"`
	UpdateTime	time.Time 	`gorm:"column:update_time" json:"update_time"`
}

//创建一个新服务
func (w *WeChatPay) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_domain_wechatpay").Create(w).Error
	return
}

//基于sid删除服务
func (w *WeChatPay) DeleteByDid(tx *gorm.DB) (err error) {
	err = tx.Exec(`
		DELETE FROM
			tbl_domain_wechatpay
		WHERE
			did = ?`,
		w.Did,
	).Error
	return
}

//修改服务信息
func (w *WeChatPay) Update(tx *gorm.DB) (err error) {
	res := tx.Table("tbl_domain_wechatpay").Where("did = ?", w.Did).Update(w)
	if err != nil {
		return err
	} else if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}

//基于tid查询域
func (w *WeChatPay) GetByDid(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_domain_wechatpay").Where("did = ?", w.Did).Scan(&w).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

