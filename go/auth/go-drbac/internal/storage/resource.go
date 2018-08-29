package storage

import (
	"github.com/jinzhu/gorm"
	. "auth/go-drbac/common"
	log "github.com/cihub/seelog"
)

type Resource struct {
	ResId       int32     	`gorm:"column:res_id;primary_key;unique"`
	ResName   	string   	`gorm:"column:res_name"`
	ResUrl		string		`gorm:"column:res_url"`
	ResOpt		string		`gorm:"column:res_opt"`
}

//分页条件查询resources
func (r *Resource) GetResources(tx *gorm.DB, page, count int32) ([]*Resource, int32, error) {
	var totalCount int32
	resources := make([]*Resource, 0, totalCount)
	if r.ResName != "" {
		resName := "%" + r.ResName + "%"
		if err := tx.Table("tbl_resource").Where(&Resource{ResOpt:r.ResOpt},).Where("res_name like ?",resName).Count(&totalCount).Error; err != nil {
			log.Error("GetResources err,",err)
			return nil, 0, err
		}
		if count == 0 {
			count = totalCount
		}
		if err := tx.Table("tbl_resource").Where(&Resource{ResOpt:r.ResOpt},).Where("res_name like ?",resName).Order("res_id asc").Limit(count).Offset((page - 1) * count).Find(&resources).Error; err != nil {
			log.Error("GetResources from db error", err)
			return nil, 0, err
		}
	}else {
		if err := tx.Table("tbl_resource").Where(&Resource{ResOpt:r.ResOpt, ResName:r.ResName}).Count(&totalCount).Error; err != nil {
			log.Error("GetResources err,",err)
			return nil, 0, err
		}
		if count == 0 {
			count = totalCount
		}
		if err := tx.Table("tbl_resource").Where(&Resource{ResOpt:r.ResOpt, ResName:r.ResName}).Order("res_id asc").Limit(count).Offset((page - 1) * count).Find(&resources).Error; err != nil {
			log.Error("GetResources from db error", err)
			return nil, 0, err
		}
	}
	return resources, totalCount, nil
}

//通过id查询单个resource
func (r *Resource) GetResourceByResId(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_resource").Where("res_id = ? ",r.ResId).Find(r).Error
	return
}

//通过ids批量查询resource
func (r *Resource) GetResourceByResIds(ids []int32,tx *gorm.DB) (res []*Resource, err error) {
	err = tx.Table("tbl_resource").Where("res_id in (?) ",ids).Find(&res).Error
	return
}

//创建一个新权限模块
func (r *Resource) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_resource").Create(r).Error
	return
}

//修改权限模块
func (r *Resource) UpdateByResId(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_resource").Where("res_id = ? ",r.ResId).Update(r).Error
	return
}

//删除权限模块
func (r *Resource) DeleteByResId(tx *gorm.DB) (err error) {
	res := tx.Exec(`
		DELETE FROM
			tbl_resource
		WHERE res_id = ?`,
		r.ResId,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}

//基于角色ID查询权限信息
func GetRolePermissionInfoByRid(tx *gorm.DB, perIds []int32)(resource []*Resource,err error) {
	err = tx.Raw(`
		SELECT
			tbl_resource.res_id,
			tbl_resource.res_name,
			tbl_resource.res_opt,
			tbl_resource.res_url
		FROM
			tbl_resource,tbl_permission_resource
		WHERE
			tbl_permission_resource.per_id in (?)
		AND
			tbl_permission_resource.res_id = tbl_resource.res_id
	`,perIds).Scan(&resource).Error
	return
}


