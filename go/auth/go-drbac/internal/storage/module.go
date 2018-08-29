package storage

import (
	"github.com/jinzhu/gorm"
	. "auth/go-drbac/common"
)

type Module struct {
	Mid        int32     	`gorm:"column:module_id;primary_key;unique" json:"mid"`
	ModuleName   string   	`gorm:"column:module_name" json:"moduleName"`
	ModuleDid	int32		`gorm:"column:module_did" json:"moduleDid"`
	ModuleTid	int32		`gorm:"column:module_tid" json:"moduleTid"`
}

//创建一个新权限模块
func (m *Module) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_module").Create(m).Error
	return
}

//修改权限模块
func (m *Module) UpdateByMid(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_module").Where("mid = ? ",m.Mid).Update(m).Error
	return
}

//查询权限模块
func (m *Module) GetModuleByMid(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_module").Where("mid = ? ",m.Mid).Find(m).Error
	return
}

//查询所有的权限模块
func (m *Module) GetModulesByTid(tx *gorm.DB,) (modules []*Module, err error) {
	err = tx.Table("tbl_module").Where("module_tid = ?",m.ModuleTid).Find(&modules).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

//查询所有的权限模块
func (m *Module) GetModulesByDid(tx *gorm.DB,) (modules []*Module, err error) {
	err = tx.Table("tbl_module").Where("module_did = ?",m.ModuleDid).Find(&modules).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

//删除权限模块
func (m *Module) DeleteByMID(tx *gorm.DB) (err error) {
	res := tx.Exec(`
		DELETE FROM
			tbl_module
		WHERE mid = ?`,
		m.Mid,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}
