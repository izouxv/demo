package storage

import (
	. "auth/go-drbac/common"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"time"
)

type Role struct {
	Rid        int32     `gorm:"column:rid;primary_key;unique" json:"rid"`
	RoleName   string    `gorm:"column:roleName" json:"roleName"`
	Description   string `gorm:"column:description" json:"description"`
	Tid   		int32    `gorm:"column:tid" json:"tid"`
	Did   		int32    `gorm:"column:did" json:"did"`
	CreateTime time.Time `gorm:"column:create_time" json:"-"`
	UpdateTime time.Time `gorm:"column:update_time" json:"-"`
}

func (r *Role) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data, r)
}

func (r *Role) MarshalJson() ([]byte, error) {
	return json.Marshal(r)
}

//添加完成后的回调函数
func (r *Role) AfterCreate(tx *gorm.DB) {
	r.GetMaxRID(tx)
}

//创建一个新角色
func (r *Role) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_role").Create(r).Error
	return
}

//删除角色
func (r *Role) DeleteByRID(tx *gorm.DB) (err error) {
	res := tx.Exec(`
		DELETE FROM
			tbl_role
		WHERE rid = ?`,
		r.Rid,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}

//删除角色
func (r *Role) DeleteByTid(tx *gorm.DB) (err error) {
	err = tx.Exec(`
		DELETE FROM
			tbl_role
		WHERE tid = ?`,
		r.Tid,
	).Error
	return
}


//查询角色信息
func (r *Role) GetRoleByRid(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_role").Where("rid = ?", r.Rid).First(&r).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

//查询角色信息
func (r *Role) GetRoleByRidAndDid(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_role").Where("rid = ? and tid = ?", r.Rid, r.Tid).First(&r).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}


//修改角色信息
func (r *Role) UpdateRoleByRid(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_role").Where("rid = ?", r.Rid).Update(&r).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

//查询所有的角色信息
func (r *Role) GetRoles(tx *gorm.DB, page, count int32) (roles []*Role, totalCount int32, err error) {
	if err = tx.Table("tbl_role").Where("tid = ?",r.Tid).Count(&totalCount).Error; err != nil {
		return
	}
	if count == 0 || count == -1{
		count = totalCount
	}
	err = tx.Table("tbl_role").Where("tid = ?",r.Tid).Order("rid asc").Limit(count).Offset(((page - 1) * count)-1).Find(&roles).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

//查询所有的角色信息
func (r *Role) GetDomainRoles(tx *gorm.DB, page, count int32) (roles []*Role, totalCount int32, err error) {
	if err = tx.Table("tbl_role").Where("did = ?",r.Did).Count(&totalCount).Error; err != nil {
		return
	}
	if count == 0 || count == -1{
		count = totalCount
	}
	err = tx.Table("tbl_role").Where("did = ?",r.Did).Order("rid asc").Limit(count).Offset(((page - 1) * count)-1).Find(&roles).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

//查询最大id GetMaxID
func (r *Role) GetMaxRID(tx *gorm.DB) (err error) {
	err = tx.Raw(`SELECT max(rid) as rid FROM tbl_role`).Scan(&r).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

