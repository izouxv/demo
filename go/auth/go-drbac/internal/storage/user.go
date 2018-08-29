package storage

import (
	"time"
	"github.com/jinzhu/gorm"
	. "auth/go-drbac/common"
	log "github.com/cihub/seelog"
)

type User struct {
	Uid        int32     `gorm:"column:uid;primary_key;unique" json:"uid"`
	Username   string    `gorm:"column:username" json:"username"`
	State      int32     `gorm:"column:state" json:"state"`
	Nickname   string    `gorm:"column:nickname" json:"nickname"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	Password   string    `gorm:"column:password" json:"-"`
	Salt       string    `gorm:"column:salt" json:"-"`
	Token	   string	 `gorm:"column:token" json:"-"`
	Tid      	int32    `gorm:"column:tid" json:"tid"`
	Did      	int32    `gorm:"column:did" json:"did"`
}
//添加回调函数
func (u *User) BeforeCreate() (err error) {
	if err != nil {
		log.Info("BeforeCreate 生成uid异常")
		return
	}
	if u.Password != "" {
		u.Salt = string(Krand(6, KC_RAND_KIND_ALL))
		log.Info("u.Salt:",u.Salt)
		u.Password = EncryptWithSalt(u.Password,u.Salt)
		u.State = 3
	}else {
		u.State = 1
	}
	u.UpdateTime = time.Now()
	u.CreateTime = time.Now()
	return
}
//修改回调函数
func (u *User) BeforeUpdate()  {
	u.UpdateTime = time.Now()
}
//创建一个新用户
func (u *User) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_user").Create(u).Error
	return
}
//基于Uid删除用户
func (u *User) DeleteByUID(tx *gorm.DB)(err error) {
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

//基于Tid删除用户
func (u *User) DeleteByTid(tx *gorm.DB)(err error) {
	err = tx.Exec(`
		DELETE FROM
			tbl_user
		WHERE tid = ?`,
		u.Tid,
	).Error
	return
}

//基于username查询用户信息
func (u *User) GetUserByUsernameAndTid(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_user").Where("username = ? and tid = ?",u.Username, u.Tid).Scan(&u).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}

//基于username查询用户信息
func (u *User) GetUserByUsernameAndDid(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_user").Where("username = ? and did = ?",u.Username, u.Did).Scan(&u).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}

//基于Token查询用户信息
func (u *User) GetUserByToken(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_user").Where("token = ?",u.Token).Scan(&u).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}


//基于Uid查询用户信息
func (u *User) GetUserByUID(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_user").Where("uid = ?",u.Uid).Scan(&u).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}
//基于UIDs批量查询用户信息
func (u *User) GetUserByUIDs(ids []int64,tx *gorm.DB) (users []*User,err error) {
	err = tx.Table("tbl_user").Where("uid in (?)",ids).Scan(&users).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}

//修改用户密码
func (u *User) UpdatePassword(tx *gorm.DB) (err error) {
	u.Salt = string(Krand(6,KC_RAND_KIND_UPPER))
	u.Password = EncryptWithSalt(u.Password,u.Salt)
	res := tx.Exec(`
		UPDATE
			tbl_user
		SET
			password = ?,
			salt = ?
		WHERE
			uid = ?`,
		u.Password,
		u.Salt,
		u.Uid,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}

//修改用户密码
func (u *User) UpdateNicknameAndPassword(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_user").Where("uid = ?",u.Uid).Update(map[string]interface{}{"nickname":u.Nickname,"password":u.Password,"salt":u.Salt}).Error
	return
}

//判断密码是否正确
func (u *User) IsTruePassword(password string,tx *gorm.DB) (isTrue bool,err error)  {
	if u.Password == EncryptWithSalt(password,u.Salt) {
		isTrue = true
	}
	return
}
//修改用户昵称
func (u *User) UpdateNickname(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_user").Where("uid = ?",u.Uid).Update("nickname",u.Nickname).Error
	return
}

//完善用户信息
func (u *User) UpdateUser(tx *gorm.DB) (ra int64,err error) {
	reply := tx.Table("tbl_user").Where("uid = ?",u.Uid).Update(&u)
	err = reply.Error
	ra = reply.RowsAffected
	return
}

//获取所有用户信息
func (u *User) GetAllUserInfo(tx *gorm.DB) (users []*User, err error) {
	err = tx.Table("tbl_user").Find(&users).Error
	return
}

//获取所有用户信息
func (u *User) GetUserInfoByUids(uids []int64, tx *gorm.DB) (users []*User, err error) {
	err = tx.Table("tbl_user").Where("uid in (?)",uids).Find(&users).Error
	return
}

//获取所有用户信息
func (u *User) GetUserInfoByUidsAndTid(uids []int64, tx *gorm.DB) (users []*User, err error) {
	err = tx.Table("tbl_user").Where("uid in (?) and tid = ?",uids, u.Tid).Find(&users).Error
	return
}

//获取所有用户信息
func (u *User) GetUserInfoByUidsAndDid(uids []int64, tx *gorm.DB) (users []*User, err error) {
	err = tx.Table("tbl_user").Where("uid in (?) and did = ?",uids, u.Did).Find(&users).Error
	return
}