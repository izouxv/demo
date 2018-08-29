package storage

import (
	. "auth/go-drbac/common"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"time"
)

type Domain struct {
	Did        		int32     `gorm:"column:did;primary_key;unique"`
	DomainName 		string    `gorm:"column:domainName"`
	DomainUrl  		string    `gorm:"column:domainUrl"`
	DomainState  	int32     `gorm:"column:domainState"`
	CreateTime 		time.Time `gorm:"column:create_time"`
	UpdateTime 		time.Time `gorm:"column:update_time"`
}

func (d *Domain) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data, d)
}

func (d *Domain) MarshalJson() ([]byte, error) {
	return json.Marshal(d)
}

//添加回调函数
func (d *Domain) BeforeCreate() (err error) {
	d.UpdateTime = time.Now()
	d.CreateTime = time.Now()
	return
}

//修改回调函数
func (d *Domain) BeforeUpdate() {
	d.UpdateTime = time.Now()
}

//创建一个新域
func (d *Domain) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_domain").Create(d).Error
	return
}

//基于域id删除域
func (d *Domain) DeleteByDID(tx *gorm.DB) (err error) {
	res := tx.Exec(`
		DELETE FROM
			tbl_domain
		WHERE did = ?`,
		d.Did,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}

//基于ids批量删除域
func (d *Domain) DeleteByTIDs(ids []int64, tx *gorm.DB) (err error) {
	res := tx.Exec(`
		DELETE FROM
			tbl_domain
		WHERE did in (?)`,
		ids,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}

//修改域名称
func (d *Domain) Update(tx *gorm.DB) (err error) {
	res := tx.Exec(`
		UPDATE
			tbl_domain
		SET
			domainName = ?
		WHERE did = ?`,
		d.DomainName,
		d.Did,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}

//基于did查询域
func (d *Domain) GetByID(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_domain").Where("did = ?", d.Did).Scan(&d).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

//基于IDs查询域
func (d *Domain) GetDomainsByIDs(ids []int64, tx *gorm.DB) (domains []*Domain, err error) {
	err = tx.Table("tbl_domain").Where("did in (?)", ids).Find(&domains).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

func (d *Domain) GetDomains(tx *gorm.DB) (domains []*Domain, err error) {
	err = tx.Table("tbl_domain").Find(&domains).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}