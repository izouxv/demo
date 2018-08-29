package storage

import (
"github.com/jinzhu/gorm"
. "auth/go-drbac/common"
	"time"
)

type Policy struct {
	Pid        		int32     `gorm:"column:policy_id;primary_key;unique"json:"pid"`
	PolicyName 		string    `gorm:"column:policy_name" json:"serviceName"`
	PolicyType      int32     `gorm:"column:policy_type" json:"policyType"`
	PolicyCycle		int32     `gorm:"column:policy_cycle" json:"policyCycle"`
	PolicyFeeType   int32     `gorm:"column:policy_fee_type" json:"policyFeeType"`
	PolicyUnitPrice	float32   `gorm:"column:policy_unit_price" json:"policyUnitPrice"`
	PolicyUnitType  int32     `gorm:"column:policy_unit_type" json:"policyUnitType"`
	PolicyUnitCount	int32     `gorm:"column:policy_unit_count" json:"policyUnitCount"`
	PolicySid	    int32     `gorm:"column:policy_sid" json:"policySid"`
	CreateTime	    time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime	    time.Time `gorm:"column:update_time" json:"updateTime"`
}

//创建一个新计费策略
func (p *Policy) Create(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_policy").Create(p).Error
	return
}


//基于pid删除计费策略
func (p *Policy) DeleteByPid(tx *gorm.DB) (err error) {
	err = tx.Exec(`
		DELETE FROM
			tbl_policy
		WHERE policy_id = ?`,
		p.Pid,
	).Error
	return
}

//基于pid删除计费策略
func (p *Policy) DeleteBySid(tx *gorm.DB) (err error) {
	err = tx.Exec(`
		DELETE FROM
			tbl_policy
		WHERE policy_sid = ?`,
		p.PolicySid,
	).Error
	return
}


//修改策略信息信息
func (p *Policy) Update(tx *gorm.DB) (err error) {
	res := tx.Table("tbl_policy").Where("policy_id = ?",p.Pid).Update(p)
	if err != nil {
		return err
	}else if res.RowsAffected == 0 {
		return ErrDoesNotExist
	}
	return
}

//基于pid查询计费策略
func (p *Policy) GetByPid(tx *gorm.DB) (err error) {
	err = tx.Table("tbl_policy").Where("policy_id = ?",p.Pid).Scan(&p).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}

//基于Sid查询域
func (p *Policy) GetBySid(tx *gorm.DB)(policys []*Policy,err error)  {
	err = tx.Table("tbl_policy").Where("policy_sid = ?",p.PolicySid).Find(&policys).Error
	if err == gorm.ErrRecordNotFound{
		err = ErrDoesNotExist
	}
	return
}

