package storage

import (
	"github.com/jinzhu/gorm"
)

type ModulePermission struct {
	Mid			int32			`gorm:"column:module_id"`
	Pid			int32			`gorm:"column:per_id"`
}

//创建一个新模块权限关系
func (m *ModulePermission) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_module_permission").Create(m).Error
	return
}

