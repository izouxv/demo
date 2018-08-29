package storage

import (
	"github.com/jinzhu/gorm"
	. "auth/go-drbac/common"
	"time"
)

type RoleModule struct {
	Rid        int32     `gorm:"column:rid;primary_key"`
	Mid        int32     `gorm:"column:mid;primary_key"`
	CreateTime time.Time `gorm:"column:create_time" json:"-"`
	UpdateTime time.Time `gorm:"column:update_time" json:"-"`
}

//添加角色模块关系
func (rm *RoleModule) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_role_module").Create(rm).Error
	return
}

//基于rid删除角色模块关系
func (rm *RoleModule) Delete(tx *gorm.DB) (err error)  {
	res := tx.Exec(`
		DELETE FROM
			tbl_role_module
		WHERE rid = ?`,
		rm.Rid,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}

//基于角色查询权限模块mids
func (rm *RoleModule) GetRoleModuleByRid(tx *gorm.DB) (mids []int32,err error)  {
	err = tx.Table("tbl_role_module").Where("rid = ?", rm.Rid).Pluck("mid",&mids).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}


//基于模块ID查询权限模块信息
func GetPermissionInfoByMids(tx *gorm.DB,mids []int32)(resources []*Resource,err error)  {
	err = tx.Raw(`
		SELECT
			tbl_permission.pid,
			tbl_permission.opt,
			tbl_permission.url
		FROM
			tbl_module_permission,tbl_permission
		WHERE
			tbl_module_permission.module in (?)
		AND
			tbl_permission.pid = tbl_module_permission.pid
	`,mids).Scan(&resources).Error
	return
}