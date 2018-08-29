package storage

import (
	"time"
	"github.com/jinzhu/gorm"
)

type TenantServicePolicy struct {
	Sid        int32     `gorm:"column:sid"`
	Tid        int32     `gorm:"column:tid"`
	Pid        int32	 `gorm:"column:pid"`
	StartTime time.Time `gorm:"column:start_time"`
	State        int32	 `gorm:"column:state;default:1"`
}

// Create
func (tsp *TenantServicePolicy) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_tenant_service_policy").Create(tsp).Error
	return
}

func (tsp *TenantServicePolicy) UpdateBySid(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_tenant_service_policy").Where("sid = ?",tsp.Sid).Update(tsp).Error
	return
}

func (tsp *TenantServicePolicy) UpdateBySidAndTid(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_tenant_service_policy").Where("sid = ? and tid = ?",tsp.Sid, tsp.Tid).Update(tsp).Error
	return
}

// delelte
func (tsp *TenantServicePolicy) DeleteBySid(tx *gorm.DB) (err error) {
	err = tx.Exec(`
		DELETE FROM
			tbl_tenant_service_policy
		WHERE
			sid = ?`,
		tsp.Sid, ).Error
	return
}

func (tsp *TenantServicePolicy) DeleteByPid(tx *gorm.DB) (err error) {
	err = tx.Exec(`
		DELETE FROM
			tbl_tenant_service_policy
		WHERE
			pid = ?`,
		tsp.Pid, ).Error
	return
}