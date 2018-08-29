package storage

import (
	. "account-domain-rpc/go-drbac/common"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"time"
)

type Role struct {
	Rid        int32     `gorm:"column:rid;primary_key;unique" json:"rid"`
	RoleName   string    `gorm:"column:roleName" json:"roleName"`
	CreateTime time.Time `gorm:"column:create_time" json:"-"`
	UpdateTime time.Time `gorm:"column:update_time" json:"-"`
}

func (r *Role) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data, r)
}

func (r *Role) MarshalJson() ([]byte, error) {
	return json.Marshal(r)
}

//添加回调函数
func (r *Role) BeforeCreate() {
	r.UpdateTime = time.Now()
	r.CreateTime = time.Now()
}

//修改回调函数
func (r *Role) BeforeUpdate() {
	r.UpdateTime = time.Now()
}

//添加完成后的回调函数
func (r *Role) AfterCreate(tx *gorm.DB) {
	r.GetMaxRID(tx)
	SetRoleMap(tx)
}

//创建一个新角色
func (r *Role) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_role").Create(r).Error
	return
}

//基于组织rid删除角色
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

//基于rid查询角色信息
func (r *Role) GetRoleByRid(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_role").Where("rid = ?", r.Rid).First(&r).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

//查询所有的角色信息
func (r *Role) GetRoles(tx *gorm.DB) (roles []*Role, err error) {
	err = tx.Table("tbl_role").Find(&roles).Error
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

/*SetRoleMap
将角色信息放入map对象中*/
func SetRoleMap(tx *gorm.DB) (err error) {
	role := Role{}
	roles, err := role.GetRoles(tx)
	if err != nil {
		return
	}
	for _, v := range roles {
		RoleMap[v.Rid] = v.RoleName
	}
	return
}

/*SetRoleName
获取角色的名称*/
func (r *Role) SetRoleName(tx *gorm.DB) (err error) {
	roleName, err := GetRoleNameFromRoleMapByRid(r.Rid, tx)
	r.RoleName = roleName
	return
}

/*GetRoleNameFromRoleMapByRid
获取map中的角色信息
*/
func GetRoleNameFromRoleMapByRid(rid int32, tx *gorm.DB) (roleName string, err error) {
	if len(RoleMap) == 0 {
		err = SetRoleMap(tx)
	}
	roleName = RoleMap[rid]
	return
}

/*DeleteRoleMapByRid
删除map中的角色信息
*/
func DeleteRoleMapByRid(rid int32) {
	delete(RoleMap, rid)
}
