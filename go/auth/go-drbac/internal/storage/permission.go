package storage

import (
	"github.com/jinzhu/gorm"
	"encoding/json"
	. "auth/go-drbac/common"
)

type Permission struct {
	Pid               int32        `gorm:"column:pid;primary_key;unique" json:"per_id"`
	PerName    		  string       `gorm:"column:per_name" json:"per_name"`
	PerOperation	  string	   `gorm:"column:per_operation" json:"per_operation"`
}

type MidPermission struct {
	Pid               int32        `gorm:"column:pid;primary_key;unique" json:"per_id"`
	PerName    		  string       `gorm:"column:per_name" json:"per_name"`
	PerOperation	  string	   `gorm:"column:per_operation" json:"per_operation"`
	Mid	  	  		  int32	       `gorm:"column:module_id" json:"module_id"`
}

func (p *Permission) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data,p)
}

func (p *Permission) MarshalJson() ([]byte,error) {
	return json.Marshal(p)
}


//创建一个新权限
func (p *Permission) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_permission").Create(p).Error
	return
}

//基于id删除权限信息
func (p *Permission) DeleteByPID(tx *gorm.DB) (err error) {
	res := tx.Exec(`
		DELETE FROM
			tbl_permission
		WHERE per_id = ?`,
		p.Pid,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}

//基于id修改权限信息
func (p *Permission) Update(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_permission").Where("per_id = ?",p.Pid).Update(p).Error
	return
}
//基于pid查询权限
func (p *Permission) GetPermission(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_permission").Where("per_id = ?",p.Pid).Scan(&p).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}
//基于ids批量查询权限
func (p *Permission) GetPermissions(ids []int32,tx *gorm.DB)(permissions []*Permission,err error)  {
	err = tx.Table("tbl_permission").Where("per_id in (?)",ids).Find(&permissions).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

//基于module_id和operation批量查询permission
//func GetPidsByModule(mid int32, operations []string, tx *gorm.DB) (perIds []int32,err error) {
//	err = tx.Table("tbl_permission").Where("per_module = ? and per_operation in (?)",mid, operations).Pluck("per_id",&perIds).Error
//	if err == gorm.ErrRecordNotFound {
//		err = ErrDoesNotExist
//	}
//	return
//}
func GetPidsByModule(mid int32, operations []string, tx *gorm.DB) (perIds []int32,err error) {
	err = tx.Raw(`
		SELECT
			tbl_permission.per_id
		FROM
			tbl_module_permission,tbl_permission
		WHERE
			tbl_module_permission.module_id = ?
		AND
			tbl_module_permission.per_id = tbl_permission.per_id
		AND
			tbl_permission.per_operation in (?)
`,mid,operations).Pluck("per_id",&perIds).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}