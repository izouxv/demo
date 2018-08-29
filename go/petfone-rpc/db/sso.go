package db

import (
	"petfone-rpc/core"
	"time"
	"github.com/jinzhu/gorm"
	log "github.com/cihub/seelog"
)

type Sso struct {
	Id         int32     `gorm:"column:id"`
	Username   string    `gorm:"column:username"`
	Password   string    `gorm:"column:password"`
	Salt       string    `gorm:"column:salt"`
	State      int32     `gorm:"column:state"`
	Nickname   string    `gorm:"column:nickname"`
	UpdateTime time.Time `gorm:"column:update_time"`
	CreateTime time.Time `gorm:"column:create_time"`
	RegIp		int64	`gorm:"column:reg_ip"`
	RegAddr		string	`gorm:"column:reg_addr"`
	LastLoginTime   time.Time 	`gorm:"last_login_time"`
	LastLoginIp     int64     	`gorm:"last_login_ip"`
	LastLoginAddr	string		`gorm:"last_login_addr"`
	LastLoginDevInfo	string	`gorm:"last_login_dev_info"`
}

//账号添加
func (s *Sso) Insert(dbc *gorm.DB) error {
	return dbc.Table("account").Create(s).Error
}

//通过用户名查找用户信息
func (s *Sso) GetByName() error {
	return core.MysqlClient.Table("account").Where("username = ?", s.Username).First(&s).Error
}

//修改密码
func (s *Sso) UpdatePasswordAndSalt() error {
	return core.MysqlClient.Table("account").Where("id = ?", s.Id).Update(s).Error
}

//通过账号修改密码
func (s *Sso) UpdatePasswordAndSaltByName() error {
	sso := Sso{Username: s.Username, Password: s.Password, Salt: s.Salt}
	return core.MysqlClient.Table("account").Where("username = ?", sso.Username).Update(sso).Error
}

//修改用户
func (this *Sso) UpdateUserName(username string) error {
	return core.MysqlClient.Table("account").Where("username = ?", this.Username).Updates(
		map[string]interface{}{"username": username,"state":0}).Error
}

//修改sso,有则修改
func (s *Sso) UpdateSso(dbc *gorm.DB) error {
	return dbc.Table("account").Where("id = ?", s.Id).Update(&s).Error
}

//通过ID查找用户信息
func (s *Sso) GetUserInfoById() error {
	return core.MysqlClient.Table("account").Where("id = ?", s.Id).First(s).Error
}

//批量查询用户信息
func (s *Sso) GetBatchSsoInfo(uids []int32) (ssos []Sso, err error) {
	return ssos, core.MysqlClient.Table("account").Where("id in (?)", uids).Find(&ssos).Error
}

//分页查询用户信息
func (s *Sso) GetPageSsoInfo(startId,count int32, sort string) (ssos []*Sso,totalCount int32, err error) {
	log.Info("GetPageSsoInfo:",startId,count,sort)
	return ssos, totalCount,core.MysqlClient.Table("account").Where("state = 3").
		Order("id "+sort).Count(&totalCount).Offset(startId).Limit(count).Find(&ssos).Error
}
