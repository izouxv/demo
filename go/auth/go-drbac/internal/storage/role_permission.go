package storage

import (
	"github.com/jinzhu/gorm"
	. "auth/go-drbac/common"
	"time"
)

type RolePermission struct {
	Rid        int32     `gorm:"column:rid;primary_key"`
	Pid        int32     `gorm:"column:per_id;primary_key"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

//添加角色权限关系
func (rp *RolePermission) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_role_permission").Create(rp).Error
	return
}

//基于rid删除角色权限关系
func (rp *RolePermission) DeleteByRid(tx *gorm.DB) (err error)  {
	err = tx.Exec(`
		DELETE FROM
			tbl_role_permission
		WHERE rid = ?`,
		rp.Rid,
	).Error
	return
}
//基于pid删除角色权限关系
func (rp *RolePermission) DeleteByPid(tx *gorm.DB) (err error)  {
	res := tx.Exec(`
		DELETE FROM
			tbl_role_permission
		WHERE pid = ?`,
		rp.Pid,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}
//基于角色查询权限ID
func (rp *RolePermission) GetRolePermissionByRid(tx *gorm.DB) (ids []int32,err error)  {
	err = tx.Table("tbl_role_permission").Where("rid = ?", rp.Rid).Pluck("pid",&ids).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}

//基于角色rids查询权限ID
func GetPermissionByRids(rids []int32, tx *gorm.DB) (ids []int32,err error)  {
	err = tx.Table("tbl_role_permission").Where("rid in (?)", rids).Pluck("per_id",&ids).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}


//基于角色ID查询权限信息
func (rp *RolePermission) GetModulesByRid(tx *gorm.DB)(midPermission []*MidPermission,err error) {
	err = tx.Raw(`
		SELECT
			tbl_module_permission.module_id,
			tbl_permission.per_id,
			tbl_permission.per_operation
		FROM
			tbl_role_permission,tbl_permission,tbl_module_permission
		WHERE
			tbl_role_permission.rid = ?
		AND
			tbl_role_permission.per_id = tbl_permission.per_id
		AND
			tbl_module_permission.per_id = tbl_permission.per_id
	`,rp.Rid).Scan(&midPermission).Error
	return
}



//基于角色ID查询权限信息
func (rp *RolePermission) GetRolePermissionInfoByRid(tx *gorm.DB)(permission []*Permission,err error) {
	err = tx.Raw(`
		SELECT
			tbl_permission.per_id,
			tbl_permission.per_opt,
			tbl_permission.url
		FROM
			tbl_role,tbl_role_permission,tbl_permission
		WHERE
			tbl_role.rid = ?
		AND
			tbl_role.rid = tbl_role_permission.rid
		AND
			tbl_role_permission.pid = tbl_permission.pid
	`,rp.Rid).Scan(&permission).Error
	return
}
