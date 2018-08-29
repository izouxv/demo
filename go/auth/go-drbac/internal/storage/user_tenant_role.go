package storage

import (
	"time"

	"github.com/jinzhu/gorm"
	. "auth/go-drbac/common"
)

type UserTenantRole struct {
	Uid        int32     `gorm:"column:uid"`
	Tid        int32     `gorm:"column:tid"`
	Rid        int32	 `gorm:"column:rid"`
	IsDefault  bool		 `gorm:"column:isDefault"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}


type UserRole struct {
	User
	Role
}

// BeforeCreate添加回调函数
func (udr *UserTenantRole) BeforeCreate() (err error) {
	udr.UpdateTime = time.Now()
	udr.CreateTime = time.Now()
	return
}
// BeforeUpdate修改回调函数
func (udr *UserTenantRole) BeforeUpdate()  {
	udr.UpdateTime = time.Now()
}
// Create创建一个新的用户-域-角色关系
func (udr *UserTenantRole) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_user_tenant_role").Create(udr).Error
	return
}
// DeleteByUid删除uid关联的域
func (udr *UserTenantRole) DeleteByUid(tx *gorm.DB)(err error)  {
	res := tx.Exec(`
		DELETE FROM
			tbl_user_tenant_role
		WHERE
			uid = ?
		AND
			isDefault = ?`,
		udr.Uid,
		false,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}
// DeleteByDid删除did关联的域关系
func (udr *UserTenantRole) DeleteByTid(tx *gorm.DB)(err error) {
	err = tx.Exec(`
		DELETE FROM
			tbl_user_tenant_role
		WHERE
			tid = ?`,
		udr.Tid,
	).Error
	return
}

// DeleteByRid删除rid关联的域关系
func (udr *UserTenantRole) DeleteByRid(tx *gorm.DB)(err error) {
	res := tx.Exec(`
		DELETE FROM
			tbl_user_tenant_role
		WHERE
			rid = ?
		AND
			isDefault = ?`,
		udr.Rid,
		false,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}
/*DeleteByUidAndDid基于uid，did删除域用户关系*/
func (udr *UserTenantRole) DeleteByUidAndTid(tx *gorm.DB)(err error)  {
	res := tx.Exec(`
		DELETE FROM
			tbl_user_tenant_role
		WHERE
			uid = ?
		AND
			tid = ?`,
		udr.Uid,
		udr.Tid,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}
// GetUserDomainRoleByUid基于Uid查询其中的域角色列表
func (udr *UserTenantRole) GetUserTenantRoleByUid(tx *gorm.DB)(udrs []*UserTenantRole,err error)  {
	err = tx.Table("tbl_user_tenant_role").Where("uid = ?",udr.Uid).Find(&udrs).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}
// GetUserRoleByUidAndDid基于用户和域获取角色
func (udr *UserTenantRole) GetUserRoleByUidAndTid(tx *gorm.DB)(udrs []*UserTenantRole, err error) {
	err = tx.Table("tbl_user_tenant_role").Where("uid = ? and tid = ?",udr.Uid,udr.Tid).Find(&udrs).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}
// GetUserDomainRoleByDid基于Did查询其中的用户角色信息
func (udr *UserTenantRole) GetUserTenantRoleByTid(tx *gorm.DB)(udrs []*UserTenantRole,err error)  {
	err = tx.Table("tbl_user_tenant_role").Where("tid = ?",udr.Tid).Find(&udrs).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}

// GetUidsByDid基于Did查询Uids
func (udr *UserTenantRole) GetUidsByTid(tx *gorm.DB, page, count int32)(uids []int64, totalCount int32,err error)  {
	if err := tx.Table("tbl_user_tenant_role").Where("tid = ?",udr.Tid).Count(&totalCount).Error; err != nil {
		return nil, 0 ,err
	}
	if count == 0 || count == -1{
		count = totalCount
	}
	err = tx.Table("tbl_user_tenant_role").Where("tid = ? and rid = ?",udr.Tid, 0).Order("uid desc").Limit(count).Offset((page - 1) * count).Pluck("uid", &uids).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}


// GetUserCountByDid基于Did查询域中成员数量
func (udr *UserTenantRole) GetUserCountByTid(tx *gorm.DB) (count int32,err error)  {
	err = tx.Table("tbl_user_tenant_role").Where("tid = ?",udr.Tid).Count(&count).Error
	return
}
/*UpdateUserRoleByDidAndUid基于Uid修改用户的角色信息*/
func (udr *UserTenantRole) UpdateUserRoleByTidAndUid(tx *gorm.DB) (err error) {
	res := tx.Exec(`
		UPDATE
			tbl_user_tenant_role
		SET
			rid = ?
		WHERE
			tid = ?
		AND
			uid = ?`,
		udr.Rid,
		udr.Tid,
		udr.Uid,
	)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}
/*GetUserDomainMaxRoleByUid
查询uid最大的角色对应的id
*/
func (udr *UserTenantRole) GetUserTenantMaxRoleByUid(tx *gorm.DB)(maxRole int32,err error)  {
	var rids []int32
	err = tx.Table("tbl_user_tenant_role").Where("uid = ?",udr.Uid).Order("rid").Limit(1).Pluck("rid",&rids).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	maxRole = rids[0]
	return
}
/*GetUserRoleByDid
获取用户对应的角色信息
*/
func (udr *UserTenantRole) GetUserRoleByTid(tx *gorm.DB)(userRoles []*UserRole,err error)  {
	err = tx.Raw(`
		SELECT
			tbl_role.rid,
			tbl_role.roleName,
			tbl_user.uid,
			tbl_user.username,
			tbl_user.nickname,
			tbl_user.phone,
			tbl_user.email,
			tbl_user.state,
			tbl_user.create_time
		FROM
			tbl_user_tenant_role,
			tbl_tenant,
			tbl_role,
			tbl_user
		WHERE
			tbl_user_tenant_role.did = ?
		AND
			tbl_user_tenant_role.did = tbl_tenant.did
		AND
			tbl_user_tenant_role.rid = tbl_role.rid
		AND
			tbl_user_tenant_role.uid = tbl_user.uid
	`,udr.Tid).Scan(&userRoles).Error
	return
}
/*IsExistence
判断用户是否在域中
*/
func (udr *UserTenantRole) IsExistedInTenant(tx *gorm.DB) (isExist bool) {
	count := 0
	if err := tx.Table("tbl_user_tenant_role").Where("tid = ? and uid = ?",udr.Tid,udr.Uid).Count(&count).Error;err != nil {
		return
	}
	if count >= 1 {
		isExist = true
	}
	return
}
/*SetDefaultDomain设置为默认域*/
func (udr *UserTenantRole) SetDefaultTenant(tx *gorm.DB) error {
	isDefault := true
	res := tx.Exec(`
		UPDATE
			tbl_user_tenant_role
		SET
			isDefault = ?
		WHERE
			tid = ?
		AND
			uid = ?`,
		isDefault,
		udr.Tid,
		udr.Uid,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return nil
}

/*通过uid删除用户角色关系*/
func (udr *UserTenantRole) DeleteUserRoleByUid(tx *gorm.DB)(err error)  {
	err = tx.Exec(`
		DELETE FROM
			tbl_user_tenant_role
		WHERE
			uid = ?`,
		udr.Uid,
	).Error
	return
}

/*通过uid删除用户角色关系*/
func (udr *UserTenantRole) DeleteUserRoleByUidAndTid(tx *gorm.DB)(err error)  {
	err = tx.Exec(`
		DELETE FROM
			tbl_user_tenant_role
		WHERE
			uid = ?
		AND
			tid = ?`,
		udr.Uid, udr.Tid,
	).Error
	return
}





/*删除角色时删除用户角色关系*/
func (udr *UserTenantRole) DeleteUserRole(tx *gorm.DB)(err error)  {
	err = tx.Exec(`
		DELETE FROM
			tbl_user_tenant_role
		WHERE
			rid = ?`,
		udr.Rid,
	).Error
	return
}

/* 通过uid查询rids*/
func (udr *UserTenantRole) GetRolesByUidAndTid(tx *gorm.DB)(roles []*Role,err error)  {
	err = tx.Raw(`
		SELECT
			tbl_role.rid,
			tbl_role.roleName
		FROM
			tbl_user_tenant_role,tbl_role
		WHERE
			tbl_user_tenant_role.uid = ?
		AND
			tbl_role.rid = tbl_user_tenant_role.rid
		AND
			tbl_user_tenant_role.tid = ?
	`,udr.Uid,udr.Tid).Scan(&roles).Error
	return
}

/* 通过uid查询rids*/
func (udr *UserTenantRole) GetRolesByUid(tx *gorm.DB)(roles []*Role,err error)  {
	err = tx.Raw(`
		SELECT
			tbl_role.rid,
			tbl_role.roleName
		FROM
			tbl_user_tenant_role,tbl_role
		WHERE
			tbl_user_tenant_role.uid = ?
		AND
			tbl_role.rid = tbl_user_tenant_role.rid
	`,udr.Uid).Scan(&roles).Error
	return
}

/* 通过uid查询tids*/
func (udr *UserTenantRole) GetTidsByUid(tx *gorm.DB) (tids []int64, err error)  {
	err = tx.Table("tbl_user_tenant_role").Where("uid = ?",udr.Uid).Pluck("tid", &tids).Error
	return
}