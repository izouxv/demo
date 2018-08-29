package storage

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TenantAccount struct {
	Tid        int32     `gorm:"column:tid;primary_key;unique"`
	Balance    float32   `gorm:"column:tenant_balance"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}


//查询账户金额
func (t *TenantAccount) CreateTenantAccount(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_tenant_account").Create(t).Error
	return
}

//查询账户金额
func (t *TenantAccount) GetTenantAccount(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_tenant_account").Where("tid = ?",t.Tid).Find(&t).Error
	return
}

//修改账户金额
func (t *TenantAccount) UpdateTenantAccount(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_tenant_account").Where("tid = ?",t.Tid).Update(&t).Error
	return
}