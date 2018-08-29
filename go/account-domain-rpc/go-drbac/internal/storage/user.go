package storage

import (
	. "account-domain-rpc/go-drbac/common"
	log "github.com/cihub/seelog"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	Uid        int64     `gorm:"column:uid;primary_key;unique" json:"uid"`
	Username   string    `gorm:"column:username" json:"username"`
	State      int32     `gorm:"column:state" json:"state"`
	Nickname   string    `gorm:"column:nickname" json:"nickname"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	Password   string    `gorm:"column:password" json:"-"`
	Salt       string    `gorm:"column:salt" json:"-"`
}

//添加回调函数
func (u *User) BeforeCreate() (err error) {
	node, err := NewNode(UserID64NewNodeID)
	if err != nil {
		return
	}
	u.Uid = node.Generate().Int64()
	if u.Password != "" {
		u.Salt = string(Krand(6, KC_RAND_KIND_ALL))
		log.Info("u.Salt:", u.Salt)
		u.Password = EncryptWithSalt(u.Password, u.Salt)
		u.State = 3
	} else {
		u.State = 0
	}
	u.UpdateTime = time.Now()
	u.CreateTime = time.Now()
	return
}

//修改回调函数
func (u *User) BeforeUpdate() {
	u.UpdateTime = time.Now()
}

//创建一个新用户
func (u *User) Create(tx *gorm.DB) (err error) {
	if u.IsExistence(tx) {
		err = ErrAlreadyExists
		return
	}
	err = tx.Table("tbl_user").Create(u).Error
	return
}

//基于Uid删除用户
func (u *User) DeleteByUID(tx *gorm.DB) (err error) {
	res := tx.Exec(`
		DELETE FROM
			tbl_user
		WHERE uid = ?`,
		u.Uid,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}

//基于username查询用户信息
func (u *User) GetUserByUsername(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_user").Where("username = ?", u.Username).Scan(&u).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

//判断用户名是否已经存在
func (u *User) IsExistence(tx *gorm.DB) (isExistence bool) {
	if u.GetUserByUsername(tx) != ErrDoesNotExist {
		isExistence = true
	}
	return
}

//基于Uid查询用户信息
func (u *User) GetUserByUID(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_user").Where("uid = ?", u.Uid).Scan(&u).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

//基于UIDs批量查询用户信息
func (u *User) GetUserByUIDs(ids []int64, tx *gorm.DB) (users []*User, err error) {
	err = tx.Table("tbl_user").Where("uid in (?)", ids).Scan(&users).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

//修改用户密码
func (u *User) UpdatePassword(tx *gorm.DB) (err error) {
	u.Salt = string(Krand(6, KC_RAND_KIND_UPPER))
	u.Password = EncryptWithSalt(u.Password, u.Salt)
	res := tx.Exec(`
		UPDATE
			tbl_user
		SET
			password = ?,
			salt = ?
		WHERE uid = ?`,
		u.Password,
		u.Salt,
		u.Uid,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}

//判断密码是否正确
func (u *User) IsTruePassword(password string, tx *gorm.DB) (isTrue bool, err error) {
	err = u.GetUserByUsername(tx)
	if err != nil {
		return
	}
	if u.Password == EncryptWithSalt(password, u.Salt) {
		isTrue = true
	}
	return
}

//修改用户昵称
func (u *User) UpdateNickname(tx *gorm.DB) (err error) {
	res := tx.Exec(`
		UPDATE
			tbl_user
		SET
			nickname = ?
		WHERE uid = ?`,
		u.Nickname,
		u.Uid,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}
