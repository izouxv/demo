package storage

import (
	. "file-server/common"
	"file-server/module"
	log "github.com/cihub/seelog"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

type File struct {
	Fid        string    `gorm:"column:fid"`
	Name       string    `gorm:"column:name"`
	Ext        string    `gorm:"column:ext"`
	Path       string    `gorm:"column:path"`
	Size       int       `gorm:"column:size"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
	Url        string    `gorm:column:url`
}

/*BeforeCreate
添加回调函数
*/
func (this *File) BeforeCreate() (err error) {
	this.Url = ServerAddress + strings.Replace(FileUrl, ":fid", this.Fid, -1)
	this.CreateTime = time.Now()
	this.UpdateTime = time.Now()
	return
}

/*BeforeUpdate
修改回调函数
*/
func (this *File) BeforeUpdate() {
	this.UpdateTime = time.Now()
}

/*Create
添加数据点
*/
func (this *File) Create(tx *gorm.DB) (err error) {
	if err = tx.Table("tbl_file").Create(this).Error; err != nil {
		log.Errorf("添加文件失败 %s", err)
		return
	}
	return
}

/*GetFileForFid
获取文件信息
*/
func (this *File) GetFileForFid() (err error) {
	if err = module.MysqlClient().Table("tbl_file").Where("fid = ?", this.Fid).First(&this).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrDoesNotExist
		}
		return
	}
	return
}

func (this *File) GetList(fids ...string) (files []*File,err error) {
	log.Info("GetList fids:",fids)
	return files,module.MysqlClient().Table("tbl_file").Where("fid in (?)", fids).Select("fid,path").Find(&files).Error
}

func (this *File) DeleteList(fids ...string) error {
	log.Info("DeleteList fids:", fids)
	return module.MysqlClient().Table("tbl_file").Where("fid in (?)", fids).Delete(File{}).Error
}
