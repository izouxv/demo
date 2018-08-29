package storage

import (
	//. "auth/go-drbac/common"
	//"github.com/jinzhu/gorm"
	"time"
)

type RetentionRule struct {
	Tid						int32			`gorm:"column:tid;primary_key;unique"json:"tid"`
	RuleId					int32			`gorm:"column:rule_id" json:"rule_id"`
	RuleName         		string  		`gorm:"column:rule_name" json:"rule_name"`
	RuleType       			string  		`gorm:"column:rule_type" json:"rule_type"`
	RuleSource  			string  		`gorm:"column:rule_source" json:"rule_source"`
	MaxRetentionTime  		int32 			`gorm:"column:key" json:"key"`
	CreateTime				time.Time 		`gorm:"column:create_time" json:"create_time"`
}

////创建一个新服务
//func (r *RetentionRule) Create(tx *gorm.DB) (err error) {
//	err = tx.Table("tbl_retention_rule").Create(r).Error
//	return
//}
//
////基于sid删除服务
//func (r *RetentionRule) DeleteByDid(tx *gorm.DB) (err error) {
//	err = tx.Exec(`
//		DELETE FROM
//			tbl_domain_wechatpay
//		WHERE
//			did = ?`,
//		r.Tid,
//	).Error
//	return
//}
//
////修改服务信息
//func (w *RetentionRule) Update(tx *gorm.DB) (err error) {
//	res := tx.Table("tbl_domain_wechatpay").Where("did = ?", w.Did).Update(w)
//	if err != nil {
//		return err
//	} else if res.RowsAffected == 0 {
//		return ErrDoesNotExist
//	}
//	return
//}
//
////基于tid查询域
//func (w *RetentionRule) GetByDid(tx *gorm.DB) (err error) {
//	err = tx.Table("tbl_domain_wechatpay").Where("did = ?", w.Did).Scan(&w).Error
//	if err == gorm.ErrRecordNotFound {
//		err = ErrDoesNotExist
//	}
//	return
//}

