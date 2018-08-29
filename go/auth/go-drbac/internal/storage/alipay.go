package storage

import (
	. "auth/go-drbac/common"
	"github.com/jinzhu/gorm"
	"time"
)

type AliPay struct {
	Did         		int32  		`gorm:"column:did;primary_key;unique"json:"did"`
	AppId       		string  	`gorm:"column:app_id" json:"app_id"`
	MerchantPrivateKey  string  	`gorm:"column:merchant_private_key" json:"merchant_private_key"`
	Key  				string 		`gorm:"column:key" json:"key"`
	CreateTime			time.Time 	`gorm:"column:create_time" json:"create_time"`
	UpdateTime			time.Time 	`gorm:"column:update_time" json:"update_time"`
}

//创建一个新服务
func (a *AliPay) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_domain_alipay").Create(a).Error
	return
}

//基于sid删除服务
func (a *AliPay) DeleteByDid(tx *gorm.DB) (err error) {
	err = tx.Exec(`
		DELETE FROM
			tbl_domain_alipay
		WHERE
			did = ?`,
		a.Did,
	).Error
	return
}

//修改服务信息
func (a *AliPay) Update(tx *gorm.DB) (err error) {
	res := tx.Table("tbl_domain_alipay").Where("did = ?", a.Did).Update(a)
	if err != nil {
		return err
	} else if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}

//基于tid查询域
func (a *AliPay) GetByDid(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_domain_alipay").Where("did = ?", a.Did).Scan(&a).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}
