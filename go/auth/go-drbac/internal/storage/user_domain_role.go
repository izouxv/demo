package storage

import (
	"time"

	"github.com/jinzhu/gorm"
	. "auth/go-drbac/common"
)

type UserDomainRole struct {
	Uid        int32     `gorm:"column:uid"`
	Did        int32     `gorm:"column:did"`
	Rid        int32	 `gorm:"column:rid"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}


// BeforeCreate添加回调函数
func (udr *UserDomainRole) BeforeCreate() (err error) {
	udr.UpdateTime = time.Now()
	udr.CreateTime = time.Now()
	return
}
// BeforeUpdate修改回调函数
func (udr *UserDomainRole) BeforeUpdate()  {
	udr.UpdateTime = time.Now()
}
// Create创建一个新的用户-域-角色关系
func (udr *UserDomainRole) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_user_domain_role").Create(udr).Error
	return
}
// DeleteByUid删除uid关联的域
func (udr *UserDomainRole) DeleteByUid(tx *gorm.DB)(err error)  {
	res := tx.Exec(`
		DELETE FROM
			tbl_user_domain_role
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
func (udr *UserDomainRole) DeleteByDid(tx *gorm.DB)(err error) {
	res := tx.Exec(`
		DELETE FROM
			tbl_user_domain_role
		WHERE
			tid = ?
		AND
			isDefault = ?`,
		udr.Did,
		false,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}
// DeleteByRid删除rid关联的域关系
func (udr *UserDomainRole) DeleteByRid(tx *gorm.DB)(err error) {
	res := tx.Exec(`
		DELETE FROM
			tbl_user_domain_role
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
func (udr *UserDomainRole) DeleteByUidAndDid(tx *gorm.DB)(err error)  {
	res := tx.Exec(`
		DELETE FROM
			tbl_user_domain_role
		WHERE
			uid = ?
		AND
			did = ?`,
		udr.Uid,
		udr.Did,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}
// GetUserDomainRoleByUid基于Uid查询其中的域角色列表
func (udr *UserDomainRole) GetUserDomainRoleByUid(tx *gorm.DB)(udrs []*UserDomainRole,err error)  {
	err = tx.Table("tbl_user_domain_role").Where("uid = ?",udr.Uid).Find(&udrs).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}
// GetUserRoleByUidAndDid基于用户和域获取角色
func (udr *UserDomainRole) GetUserRoleByUidAndDid(tx *gorm.DB)(udrs []*UserDomainRole, err error) {
	err = tx.Table("tbl_user_domain_role").Where("uid = ? and did = ?",udr.Uid,udr.Did).Find(&udrs).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}
// GetUserDomainRoleByDid基于Did查询其中的用户角色信息
func (udr *UserDomainRole) GetUserDomainRoleByDid(tx *gorm.DB)(udrs []*UserDomainRole,err error)  {
	err = tx.Table("tbl_user_domain_role").Where("did = ?",udr.Did).Find(&udrs).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}

// GetUidsByDid基于Did查询Uids
func (udr *UserDomainRole) GetUidsByDid(tx *gorm.DB, page, count int32)(uids []int64,totalCount int32, err error)  {

	if err = tx.Table("tbl_user_domain_role").Where("did = ?",udr.Did).Count(&totalCount).Error; err != nil {
		return
	}
	if count == 0 || count == -1{
		count = totalCount
	}
	err = tx.Table("tbl_user_domain_role").Where("did = ? and rid = ?",udr.Did,0).Order("uid desc").Limit(count).Offset((page - 1) * count).Pluck("uid", &uids).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}


// GetUserCountByDid基于Did查询域中成员数量
func (udr *UserDomainRole) GetUserCountByDid(tx *gorm.DB) (count int32,err error)  {
	err = tx.Table("tbl_user_domain_role").Where("Did = ?",udr.Did).Count(&count).Error
	return
}
/*UpdateUserRoleByDidAndUid基于Uid修改用户的角色信息*/
func (udr *UserDomainRole) UpdateUserRoleByDidAndUid(tx *gorm.DB) (err error) {
	res := tx.Exec(`
		UPDATE
			tbl_user_domain_role
		SET
			rid = ?
		WHERE
			did = ?
		AND
			uid = ?`,
		udr.Rid,
		udr.Did,
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
func (udr *UserDomainRole) GetUserDomainMaxRoleByUid(tx *gorm.DB)(maxRole int32,err error)  {
	var rids []int32
	err = tx.Table("tbl_user_domain_role").Where("uid = ?",udr.Uid).Order("rid").Limit(1).Pluck("rid",&rids).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	maxRole = rids[0]
	return
}
/*GetUserRoleByDid
获取用户对应的角色信息
*/
func (udr *UserDomainRole) GetUserRoleByDid(tx *gorm.DB)(userRoles []*UserRole,err error)  {
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
			tbl_user_domain_role,
			tbl_domain,
			tbl_role,
			tbl_user
		WHERE
			tbl_user_domain_role.did = ?
		AND
			tbl_user_domain_role.did = tbl_domain.did
		AND
			tbl_user_domain_role.rid = tbl_role.rid
		AND
			tbl_user_domain_role.uid = tbl_user.uid
	`,udr.Did).Scan(&userRoles).Error
	return
}
/*IsExistence
判断用户是否在域中
*/
func (udr *UserDomainRole) IsExistedInDomain(tx *gorm.DB) (isExist bool) {
	count := 0
	if err := tx.Table("tbl_user_domain_role").Where("did = ? and uid = ?",udr.Did,udr.Uid).Count(&count).Error;err != nil {
		return
	}
	if count >= 1 {
		isExist = true
	}
	return
}
/*SetDefaultDomain设置为默认域*/
func (udr *UserDomainRole) SetDefaultDomain(tx *gorm.DB) error {
	isDefault := true
	res := tx.Exec(`
		UPDATE
			tbl_user_domain_role
		SET
			isDefault = ?
		WHERE
			did = ?
		AND
			uid = ?`,
		isDefault,
		udr.Did,
		udr.Uid,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return nil
}

/*通过uid删除用户角色关系*/
func (udr *UserDomainRole) DeleteUserRoleByUid(tx *gorm.DB)(err error)  {
	err = tx.Exec(`
		DELETE FROM
			tbl_user_domain_role
		WHERE
			uid = ?`,
		udr.Uid,
	).Error
	return
}

/*通过uid删除用户角色关系*/
func (udr *UserDomainRole) DeleteUserRoleByUidAndDid(tx *gorm.DB)(err error)  {
	err = tx.Exec(`
		DELETE FROM
			tbl_user_domain_role
		WHERE
			uid = ?
		AND
			did = ?`,
		udr.Uid, udr.Did,
	).Error
	return
}





/*删除角色时删除用户角色关系*/
func (udr *UserDomainRole) DeleteUserRole(tx *gorm.DB)(err error)  {
	err = tx.Exec(`
		DELETE FROM
			tbl_user_domain_role
		WHERE
			rid = ?`,
		udr.Rid,
	).Error
	return
}

/* 通过uid查询rids*/
func (udr *UserDomainRole) GetRolesByUidAndDid(tx *gorm.DB)(roles []*Role,err error)  {
	err = tx.Raw(`
		SELECT
			tbl_role.rid,
			tbl_role.roleName
		FROM
			tbl_user_domain_role,tbl_role
		WHERE
			tbl_user_domain_role.uid = ?
		AND
			tbl_role.rid = tbl_user_domain_role.rid
		AND
			tbl_user_domain_role.did = ?
	`,udr.Uid,udr.Did).Scan(&roles).Error
	return
}

/* 通过uid查询rids*/
func (udr *UserDomainRole) GetRolesByUid(tx *gorm.DB)(roles []*Role,err error)  {
	err = tx.Raw(`
		SELECT
			tbl_role.rid,
			tbl_role.roleName
		FROM
			tbl_user_domain_role,tbl_role
		WHERE
			tbl_user_domain_role.uid = ?
		AND
			tbl_role.rid = tbl_user_domain_role.rid
	`,udr.Uid).Scan(&roles).Error
	return
}

/* 通过uid查询dids*/
func (udr *UserDomainRole) GetDidsByUid(tx *gorm.DB) (dids []int64, err error)  {
	err = tx.Table("tbl_user_domain_role").Where("uid = ?",udr.Uid).Pluck("did", &dids).Error
	return
}