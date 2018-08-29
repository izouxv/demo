package storage

import (
	. "auth/go-drbac/common"
	"github.com/jinzhu/gorm"
)

type Service struct {
	Sid         int32  `gorm:"column:service_id;primary_key;unique"json:"sid"`
	ServiceName string `gorm:"column:service_name" json:"serviceName"`
	ServiceKey  string `gorm:"column:service_key" json:"serviceKey"`
	ServiceUrl  string `gorm:"column:service_url" json:"serviceUrl"`
	ServiceType int32  `gorm:"column:service_type" json:"serviceType"`
	ServiceDescription string `gorm:"column:service_description" json:"serviceDescription"`
	ServiceState int32  `gorm:"column:state" json:"serviceState"`
}

//创建一个新服务
func (s *Service) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_service").Create(s).Error
	return
}

//基于sid删除服务
func (s *Service) DeleteBySid(tx *gorm.DB) (err error) {
	err = tx.Exec(`
		DELETE FROM
			tbl_service
		WHERE
			service_id = ?
		AND
			service_type != 3`,
		s.Sid,
	).Error
	return
}

//修改服务信息
func (s *Service) Update(tx *gorm.DB) (err error) {
	res := tx.Table("tbl_service").Where("service_id = ?", s.Sid).Update(s)
	if err != nil {
		return err
	} else if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}

//基于tid查询域
func (s *Service) GetBySid(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_service").Where("service_id = ?", s.Sid).Scan(&s).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

//基于IDs查询域
func (s *Service) GetByTid(tx *gorm.DB,tid int32) (services []*Service, err error) {
	err = tx.Raw(`
		SELECT
			tbl_service.service_id,
			tbl_service.service_name,
			tbl_service.service_key,
			tbl_service.service_url,
			tbl_service.service_type,
			tbl_service.service_description,
			tbl_tenant_service_policy.state
		FROM
			tbl_tenant_service_policy,tbl_service
		WHERE
			tbl_tenant_service_policy.tid = ?
		AND
			tbl_service.service_id = tbl_tenant_service_policy.sid
	`,tid).Scan(&services).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}
