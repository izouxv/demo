package storage

import (
	. "account-domain-rpc/go-drbac/common"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"time"
)

type Permission struct {
	Pid            int32     `gorm:"column:pid;primary_key;unique" json:"pid"`
	PermissionName string    `gorm:"column:permissionName" json:"prmissionName"`
	Url            string    `gorm:"column:url" json:"url"`
	Opt            string    `gorm:"column:opt" json:"opt"`
	CreateTime     time.Time `gorm:"column:create_time" json:"-"`
	UpdateTime     time.Time `gorm:"column:update_time" json:"-"`
}

func (p *Permission) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data, p)
}

func (p *Permission) MarshalJson() ([]byte, error) {
	return json.Marshal(p)
}

//添加回调函数
func (p *Permission) BeforeCreate() {
	p.UpdateTime = time.Now()
	p.CreateTime = time.Now()
}

//修改回调函数
func (p *Permission) BeforeUpdate() {
	p.UpdateTime = time.Now()
}
func (p *Permission) AfterCreate(tx *gorm.DB) {
	p.GetMaxPID(tx)
}

//创建一个新域
func (p *Permission) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_permission").Create(p).Error
	return
}

//基于域id权限信息
func (p *Permission) DeleteByPID(tx *gorm.DB) (err error) {
	res := tx.Exec(`
		DELETE FROM
			tbl_permission
		WHERE pid = ?`,
		p.Pid,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}

//修改权限信息
func (p *Permission) Update(tx *gorm.DB) (err error) {
	res := tx.Exec(`
		UPDATE
			tbl_permission
		SET
			permissionName = ?,
			url = ?,
			opt = ?
		WHERE pid = ?`,
		p.PermissionName,
		p.Url,
		p.Opt,
		p.Pid,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}

//基于pid查询权限
func (p *Permission) GetPermissionByID(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_permission").Where("pid = ?", p.Pid).Scan(&p).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

//基于IDs查询权限
func (p *Permission) GetDomainsByIDs(ids []int32, tx *gorm.DB) (permissions []*Permission, err error) {
	err = tx.Table("tbl_permission").Where("pid in (?)", ids).Find(&permissions).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

//查询最大id GetMaxID
func (p *Permission) GetMaxPID(tx *gorm.DB) (err error) {
	err = tx.Raw(`SELECT max(pid) as pid FROM tbl_permission`).Scan(&p).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}
