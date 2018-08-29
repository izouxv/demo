package storage

import (
	"github.com/jinzhu/gorm"
	. "auth/go-drbac/common"
	"time"
)

type RoleResource struct {
	Rid        int32     `gorm:"column:rid;primary_key"`
	ResId      int32     `gorm:"column:res_id;primary_key"`
	CreateTime time.Time `gorm:"column:create_time" json:"-"`
	UpdateTime time.Time `gorm:"column:update_time" json:"-"`
}

//添加角色模块关系
func (rr *RoleResource) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_role_resource").Create(rr).Error
	return
}

//基于rid删除角色模块关系
func (rr *RoleResource) Delete(tx *gorm.DB) (err error)  {
	err = tx.Exec(`
		DELETE FROM
			tbl_role_resource
		WHERE rid = ?`,
		rr.Rid,
	).Error
	return
}

//基于角色查询权限模块mids
func (rr *RoleResource) GetRoleResourceByRid(tx *gorm.DB) (resIds []int32,err error)  {
	err = tx.Table("tbl_role_resource").Where("rid = ?", rr.Rid).Pluck("res_id",&resIds).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}


//基于rid查询所有resource
func GetPermissionInfoByRids(tx *gorm.DB,rids []int32)(resources []*Resource,err error)  {
	err = tx.Raw(`
		SELECT
			tbl_resource.res_id,
			tbl_resource.res_opt,
			tbl_resource.res_url
		FROM
			tbl_role_resource,tbl_resource
		WHERE
			tbl_role_resource.rid in (?)
		AND
			tbl_resource.res_id = tbl_role_resource.res_id
	`,rids).Scan(&resources).Error
	return
}