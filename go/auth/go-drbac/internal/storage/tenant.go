package storage

import (
	"time"
	"encoding/json"
	"github.com/jinzhu/gorm"
	. "auth/go-drbac/common"
	log "github.com/cihub/seelog"
)

type Tenant struct {
	Tid        		int32     `gorm:"column:tid;primary_key;unique" json:"tid"`
	TenantName 		string    `gorm:"column:tenantName" json:"tenantName"`
	Pid        		int32     `gorm:"column:pid" json:"pid"`
	Did        		int32     `gorm:"column:did" json:"did"`
	TenantURL		string		`gorm:"column:tenantURL" json:"tenantURL"`
	TenantState		int32		`gorm:"column:tenantState" json:"tenantState"`
	Description		string		`gorm:"column:description" json:"description"`
	Contacts		string		`gorm:"column:contacts" json:"contacts"`
	Email			string		`gorm:"column:email" json:"email"`
	Phone			string		`gorm:"column:phone" json:"phone"`
	CreateTime 		time.Time 	`gorm:"column:create_time" json:"createTime"`
	UpdateTime 		time.Time 	`gorm:"column:update_time" json:"updateTime"`
	Icon			string		`gorm:"column:icon" json:"icon"`
	Logo			string		`gorm:"column:logo" json:"logo"`
}

func (d *Tenant) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data,d)
}

func (d *Tenant) MarshalJson() ([]byte,error) {
	return json.Marshal(d)
}
//添加回调函数
func (d *Tenant) BeforeCreate() (err error) {
	d.UpdateTime = time.Now()
	d.CreateTime = time.Now()
	return
}
//修改回调函数
func (d *Tenant) BeforeUpdate()  {
	d.UpdateTime = time.Now()
}
//创建一个新域
func (d *Tenant) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_tenant").Create(d).Error
	return
}
//基于域id删除域
func (d *Tenant) DeleteByTID(tx *gorm.DB) (err error) {
	res := tx.Exec(`
		DELETE FROM
			tbl_tenant
		WHERE tid = ?`,
			d.Tid,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}
//基于父id批量删除域
func (d *Tenant) DeleteByPid(tx *gorm.DB) (err error) {
	res := tx.Exec(`
		DELETE FROM
			tbl_tenant
		WHERE pid = ?`,
		d.Pid,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}
//基于ids批量删除域
func (d *Tenant) DeleteByTIDs(ids []int64,tx *gorm.DB) (err error) {
	res := tx.Exec(`
		DELETE FROM
			tbl_tenant
		WHERE tid in (?)`,
		ids,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}
//修改租户信息
func (d *Tenant) Update(tx *gorm.DB) (err error) {
	res := tx.Table("tbl_tenant").Where("tid = ?",d.Tid).Update(d)
	if err != nil {
		return err
	}else if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}
//基于tid查询域
func (d *Tenant) GetByID(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_tenant").Where("tid = ?",d.Tid).Scan(&d).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}

//基于IDs查询域
func (d *Tenant) GetTenantsByTids(tids []int32,tx *gorm.DB)(tenants []*Tenant,err error)  {
	err = tx.Table("tbl_tenant").Where("tid in (?)",tids).Find(&tenants).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}

//基于did查询域
func (d *Tenant) GetByDid(tx *gorm.DB) (tenants []*Tenant, err error) {
	err = tx.Table("tbl_tenant").Where("did = ?",d.Did).Scan(&tenants).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}

//基于pid查询域
func (d *Tenant) GetTenantsByPid(tx *gorm.DB) (domains []*Tenant,err error) {
	err = tx.Table("tbl_tenant").Where("pid = ?",d.Pid).Find(&domains).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}
//基于父id查询子域数量
func (d *Tenant) GetTenantCountByPid(tx *gorm.DB) (count int32,err error) {
	err = tx.Table("tbl_tenant").Where("pid = ? ",d.Pid).Count(&count).Error
	return
}
//基于pid判断是否有其他
func (d *Tenant) IsExistTenantByPid(tx *gorm.DB) (exist bool) {
	count,err := d.GetTenantCountByPid(tx)
	if err != nil || count > 0 {
		exist = true
	}
	return
}

//获取域的pid
func (d *Tenant) GetPidByTID(tx *gorm.DB) (pid int32,err error) {
	var ids []int32
	err = tx.Table("tbl_tenant").Where("tid = ?", d.Tid).Pluck("pid",&ids).Error
	log.Errorf("err is (%s)",err)
	if err != nil {
		log.Errorf("err is not nil is (%s)",err)
		if err == gorm.ErrRecordNotFound {
			log.Errorf("not find",err)
			err = ErrDoesNotExist
		}
		return
	}
	if len(ids) == 0{
		log.Errorf("err is not nil is (%s),ids is (%s)",err,len(ids),ids)
		err = ErrDoesNotExist
		return
	}
	pid = ids[0]
	return
}
//基于Pid获取Dids
func (d *Tenant) GetIDsByPid(tx *gorm.DB) (ids []int64,err error)  {
	err = tx.Table("tbl_tenant").Where("pid = ?", d.Pid).Pluck("tid",&ids).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}

//基于Url获取tid
func (d *Tenant) GetTidByUrl(tx *gorm.DB) (err error)  {
	err = tx.Table("tbl_tenant").Where("tenantURL = ?", d.TenantURL).First(&d).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}

//基于Tid获取Did
func (d *Tenant) GetDidByTid(tx *gorm.DB) (err error)  {
	err = tx.Table("tbl_tenant").Where("tid = ?", d.Tid).Find(&d).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}

//基于Did获取域的深度
func (d *Tenant) GetTenantDepth(tx *gorm.DB) (depth int32,err error) {
	pid,err := d.GetPidByTID(tx)
	if err != nil {
		return
	}
	for pid != 0{
		pid,err = d.GetPidByTID(tx)
		if err != nil {
			return
		}
		d.Tid = pid
	}
	return
}
/*IsExceedMaxDepth
判断该域是否超过最大深度
*/
func (d *Tenant) IsExceedMaxDepth(tx *gorm.DB) (isExceed bool,err error)  {
	depth,err := d.GetTenantDepth(tx)
	if err != nil {
		return
	}
	if depth > MaxDepth {
		isExceed = true
	}
	return
}
