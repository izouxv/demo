package storage

import (
	. "account-domain-rpc/go-drbac/common"
	"encoding/json"
	log "github.com/cihub/seelog"
	"github.com/jinzhu/gorm"
	"time"
)

type Domain struct {
	Did        int64     `gorm:"column:did;primary_key;unique" json:"did"`
	DomainName string    `gorm:"column:domainName" json:"domainName"`
	Pid        int64     `gorm:"column:pid" json:"pid"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (d *Domain) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data, d)
}

func (d *Domain) MarshalJson() ([]byte, error) {
	return json.Marshal(d)
}

//添加回调函数
func (d *Domain) BeforeCreate() (err error) {
	node, err := NewNode(DomainID64NewNodeID)
	if err != nil {
		return
	}
	d.Did = node.Generate().Int64()
	d.UpdateTime = time.Now()
	d.CreateTime = time.Now()
	return
}

//修改回调函数
func (d *Domain) BeforeUpdate() {
	d.UpdateTime = time.Now()
}

//创建一个新域
func (d *Domain) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_domain").Create(d).Error
	return
}

//基于域id删除域
func (d *Domain) DeleteByDID(tx *gorm.DB) (err error) {
	res := tx.Exec(`
		DELETE FROM
			tbl_domain
		WHERE did = ?`,
		d.Did,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}

//基于父id批量删除域
func (d *Domain) DeleteByPid(tx *gorm.DB) (err error) {
	res := tx.Exec(`
		DELETE FROM
			tbl_domain
		WHERE pid = ?`,
		d.Pid,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}

//基于ids批量删除域
func (d *Domain) DeleteByDIDs(ids []int64, tx *gorm.DB) (err error) {
	res := tx.Exec(`
		DELETE FROM
			tbl_domain
		WHERE did in (?)`,
		ids,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}

//修改域名称
func (d *Domain) Update(tx *gorm.DB) (err error) {
	res := tx.Exec(`
		UPDATE
			tbl_domain
		SET
			domainName = ?
		WHERE did = ?`,
		d.DomainName,
		d.Did,
	)
	if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}

//基于did查询域
func (d *Domain) GetByID(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_domain").Where("did = ?", d.Did).Scan(&d).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

//基于pid查询域
func (d *Domain) GetDomainsByPid(tx *gorm.DB) (domains []*Domain, err error) {
	err = tx.Table("tbl_domain").Where("pid = ?", d.Pid).Find(&domains).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

//基于父id查询子域数量
func (d *Domain) GetDomainCountByPid(tx *gorm.DB) (count int32, err error) {
	err = tx.Table("tbl_domain").Where("pid = ? ", d.Pid).Count(&count).Error
	return
}

//基于pid判断是否有其他
func (d *Domain) IsExistDomainByPid(tx *gorm.DB) (exist bool) {
	count, err := d.GetDomainCountByPid(tx)
	if err != nil || count > 0 {
		exist = true
	}
	return
}

//基于IDs查询域
func (d *Domain) GetDomainsByIDs(ids []int64, tx *gorm.DB) (domains []*Domain, err error) {
	err = tx.Table("tbl_domain").Where("did in (?)", d.Did).Find(&domains).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

//获取域的pid
func (d *Domain) GetPidByDID(tx *gorm.DB) (pid int64, err error) {
	var ids []int64
	err = tx.Table("tbl_domain").Where("did = ?", d.Did).Pluck("pid", &ids).Error
	log.Errorf("err is (%s)", err)
	if err != nil {
		log.Errorf("err is not nil is (%s)", err)
		if err == gorm.ErrRecordNotFound {
			log.Errorf("not find", err)
			err = ErrDoesNotExist
		}
		return
	}
	if len(ids) == 0 {
		log.Errorf("err is not nil is (%s),ids is (%s)", err, len(ids), ids)
		err = ErrDoesNotExist
		return
	}
	pid = ids[0]
	return
}

//基于Pid获取Dids
func (d *Domain) GetIDsByPid(tx *gorm.DB) (ids []int64, err error) {
	err = tx.Table("tbl_domain").Where("pid = ?", d.Pid).Pluck("did", &ids).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

//基于Did获取域的深度
func (d *Domain) GetDomainDepth(tx *gorm.DB) (depth int32, err error) {
	pid, err := d.GetPidByDID(tx)
	if err != nil {
		return
	}
	for pid != 0 {
		pid, err = d.GetPidByDID(tx)
		if err != nil {
			return
		}
		d.Did = pid
	}
	return
}

/*IsExceedMaxDepth
判断该域是否超过最大深度
*/
func (d *Domain) IsExceedMaxDepth(tx *gorm.DB) (isExceed bool, err error) {
	depth, err := d.GetDomainDepth(tx)
	if err != nil {
		return
	}
	if depth > MaxDepth {
		isExceed = true
	}
	return
}
