package db

import (
	"petfone-rpc/core"
	"time"
	"github.com/jinzhu/gorm"
)

type Account struct {
	Id         int32     `gorm:"column:id"`
	Nickname   string    `gorm:"column:nickname"`
	Gender     int32     `gorm:"column:gender"`
	Birthday   time.Time `gorm:"column:birthday"`
	Avatar     string    `gorm:"column:avatar"`
	Signature  string    `gorm:"column:signature"`
	Address    string    `gorm:"column:address"`
	Email      string    `gorm:"column:email"`
	Phone      string    `gorm:"column:phone"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}
type AgentUser struct {
	Uid      int32  `gorm:"column:uid"`
	Username string `gorm:"column:username"`
	Nickname string `gorm:"column:nickname"`
}

//插入用户基本信息
func (a *Account) InsertAccount() error {
	if err := core.MysqlClient.Table("account").Create(a).Error; err != nil {
		return err
	}
	return nil
}

//通过ID查询用户基本信息
func (a *Account) GetAccountById() error {
	if err := core.MysqlClient.Table("account").Where("id = ?", a.Id).First(a).Error; err != nil {
		return err
	}
	return nil
}

//修改用户基本信息(忽略空值)
func (a *Account) UpdateAccount(dbc *gorm.DB) error {
	return dbc.Table("account").Where("id = ?", a.Id).Update(a).Error
}

//修改用户基本信息（不忽略空值）
func (a *Account) UpdateAccountInfoAll() error {
	if err := core.MysqlClient.Table("account").Where("id = ?", a.Id).Update(a).
		//Updates(map[string]interface{}{"nickname": a.Nickname, "signature": a.Signature, "gender": a.Gender,
		//	"avatar": a.Avatar, "birthday": a.Birthday, "address": a.Address, "email": a.Email, "phone": a.Phone,
		//	"update_time": time.Now()}).
				Error; err != nil {
		return err
	}
	return nil
}

//批量查询用户信息
func (a *Account) GetBatchAccount(uids []int32) ([]*Account, error) {
	var accounts []*Account
	if err := core.MysqlClient.Table("account").Where("id in (?) AND state = 3", uids).Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}
