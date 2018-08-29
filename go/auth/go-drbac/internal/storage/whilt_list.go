package storage

import (
	"time"
	"github.com/jinzhu/gorm"
)


type WhiteList struct {
	Id               int32         `gorm:"column:id;primary_key;unique" json:"id"`
	PermissionName    string       `gorm:"column:permissionName" json:"prmissionName"`
	Url               string       `gorm:"column:url" json:"url"`
	Opt			      string       `gorm:"column:opt" json:"opt"`
	CreateTime 		  time.Time    `gorm:"column:create_time" json:"-"`
	UpdateTime 		  time.Time    `gorm:"column:update_time" json:"-"`
}

func (w WhiteList) GetWhiteList(tx *gorm.DB) (count int64,err error)  {
	count = 0
	err = tx.Table("tbl_white_list").Where("url = ? AND opt = ?",w.Url, w.Opt).Count(&count).Error
	return
}

func (w WhiteList) GetAllWhiteList(tx *gorm.DB) (allWhiteList []*WhiteList,err error)  {
	err = tx.Table("tbl_white_list").Find(&allWhiteList).Error
	return
}