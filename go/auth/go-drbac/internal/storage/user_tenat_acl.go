package storage

import (
	"github.com/jinzhu/gorm"
)

type UserTenantACL struct {
	Uid        int32     `gorm:"column:uid"`
	Tid        int32     `gorm:"column:tid"`
}

// Create创建一个新的用户-租户acl关系
func (uta *UserTenantACL) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_user_tenant_acl").Create(uta).Error
	return
}

//通过uid查询所有绑定的acl关系
func (uta *UserTenantACL) GetTidsByUid(tx *gorm.DB) (tids []int32, err error) {
	err = tx.Table("tbl_user_tenant_acl").Where("uid = ?",uta.Uid).Pluck("tid",&tids).Error
	return
}