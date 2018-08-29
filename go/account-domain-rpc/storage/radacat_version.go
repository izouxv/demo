package storage

import (
	"account-domain-rpc/module"
	"github.com/pkg/errors"
	"time"
)

var (
	RadacatVersionAlreadyExists = errors.New("radacatVersion already exists")
	RadacatVersionDoesNotExist  = errors.New("radacatVersion does not exist")
	RadacatVersionCanNotDelete  = errors.New("radacatVersion should not delete ")
)

type RadacatVersion struct {
	Id            int32     `gorm:"column:id"`
	Tid           int64     `gorm:"column:tid"`
	Device        string    `gorm:"column:device"`
	VersionName   string    `gorm:"column:version_name"`
	VersionCode   string    `gorm:"column:version_code"`
	MD5           string    `gorm:"column:md5"`
	FileName      string    `gorm:"column:filename"`
	FileLength    int64     `gorm:"column:length"`
	URL           string    `gorm:"column:path"`
	DescriptionCN string    `gorm:"column:description_cn"`
	DescriptionEN string    `gorm:"column:description_en"`
	UploaderUid   int64     `gorm:"column:uploader_uid"`
	Status        int32     `gorm:"column:status"`
	CreateTime    time.Time `gorm:"column:create_time"`
	UpdateTime    time.Time `gorm:"column:update_time"`

}

func (r *RadacatVersion) Create() error {
	r.CreateTime = time.Now()
	r.UpdateTime = time.Now()
	if err := module.MysqlClient().Table("tbl_radacat_version").Create(r).Error; err != nil {
		return errors.Wrap(err, "create radacat_version error")
	}
	return nil
}

func (r *RadacatVersion) Update() error {
	r.UpdateTime = time.Now()
	reply := module.MysqlClient().Table("tbl_radacat_version").Where("id = ? and tid = ? ", r.Id,r.Tid).Update(r)
	if reply.Error != nil {
		return errors.Wrap(reply.Error, "Update radacat_version error")
	}
	if reply.RowsAffected == 0 {
		return RadacatVersionDoesNotExist
	}
	return nil
}

func (r *RadacatVersion) GetLatestVersion() error {
	reply := module.MysqlClient().Raw("SELECT * FROM tbl_radacat_version where device = ? and tid = ? and status !=3 ORDER BY version_code DESC,create_time desc  LIMIT 1", r.Device, r.Tid).Scan(&r)
	if reply.Error != nil {
		if reply.RowsAffected == 0 {
			return RadacatVersionDoesNotExist
		}
		return reply.Error
	}
	return nil
}

func (r *RadacatVersion) GetLatestVersionRelease() error {
	reply := module.MysqlClient().Raw("SELECT * FROM tbl_radacat_version where device = ? and tid = ? and status = 1 ORDER BY version_code DESC  LIMIT 1", r.Device, r.Tid).Scan(&r)
	if reply.Error != nil {
		if reply.RowsAffected == 0 {
			return RadacatVersionDoesNotExist
		}
		return reply.Error
	}
	return nil
}

func (r *RadacatVersion) GetAllVersions(page, count int32) ([]*RadacatVersion, int32, error) {
	var totalCount int32
	if err := module.MysqlClient().Table("tbl_radacat_version").Where("tid = ? and status != 3",r.Tid).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}
	if count == 0 {
		count = totalCount
	}
	rvs := make([]*RadacatVersion, 0, totalCount)
	if err := module.MysqlClient().Table("tbl_radacat_version").Where("tid = ? and status != 3 ",r.Tid).Order("id desc").Limit(count).Offset((page - 1) * count).Find(&rvs).Error; err != nil {
		return nil, 0, err
	}
	return rvs, totalCount, nil
}

/*DeleteVersion删除版本*/
func (r *RadacatVersion) DeleteVersion() (err error) {
	rs := module.MysqlClient().Table("tbl_radacat_version").Where("id = ? and tid = ?",r.Id,r.Tid).Update(map[string]interface{}{"status":3})
	if rs.Error != nil {
		return errors.Wrap(rs.Error, "Delete radacat_version error")
	}
	if rs.RowsAffected == 0 {
		return RadacatVersionDoesNotExist
	}
	return nil
}


func (r *RadacatVersion) GetVersion() (*RadacatVersion,error) {
	rv := module.MysqlClient().Table("tbl_radacat_version").Where("tid = ? and id = ?",r.Tid,r.Id).Find(&r)
	if rv.Error != nil {
		if rv.RowsAffected == 0{
			return nil, RadacatVersionDoesNotExist
		}
		return nil,rv.Error
	}
	return r, nil
}


func (r *RadacatVersion) GetVersionByVersionCode() (*RadacatVersion,error) {
	rv := module.MysqlClient().Table("tbl_radacat_version").Where("tid = ? and version_code = ? and device = ? and status !=3",r.Tid,r.VersionCode,r.Device).Find(&r)
	if rv.Error != nil {
		if rv.RowsAffected == 0{
			return nil, RadacatVersionDoesNotExist
		}
		return nil,rv.Error
	}
	return r, nil
}


func (r *RadacatVersion) GetVersionByVersionCodePut() (*RadacatVersion,error) {
	rv := module.MysqlClient().Table("tbl_radacat_version").Where("tid = ? and version_code = ? and device = ? and id != ? and status != 3",r.Tid,r.VersionCode,r.Device,r.Id).Find(&r)
	if rv.Error != nil {
		if rv.RowsAffected == 0{
			return nil, RadacatVersionDoesNotExist
		}
		return nil,rv.Error
	}
	return r, nil
}


