package db

import (
	log "github.com/cihub/seelog"
	"petfone-rpc/core"
	"time"
)

type FaqCommonPo struct {
	Id				int32		`gorm:"column:id"`
	NameCn			string		`gorm:"column:name_cn"`
	InfoCn			string    	`gorm:"column:info_cn"`
	NameEn			string		`gorm:"column:name_en"`
	InfoEn			string    	`gorm:"column:info_en"`
	Parent			int32    	`gorm:"column:parent"`
	CreationTime 	time.Time	`gorm:"column:creation_time"` 	//创建时间
	DataState    	int32		`gorm:"column:data_state"`    	//数据状态
}

func (this *FaqCommonPo) SetFAQById() error {
	log.Info("SetFAQById:",this)
	return core.MysqlClient.Table("faqcommon").Create(this).Error
}

func (this *FaqCommonPo) GetFAQById() error {
	log.Info("GetFAQById:",this)
	return core.MysqlClient.Table("faqcommon").
		Where("id = ? and data_state = ?", this.Id,this.DataState).Find(&this).Error
}

func (this *FaqCommonPo) GetFAQByKeyword() (faqcommon []FaqCommonPo, err error) {
	log.Info("GetFAQByKeyword:",this)
	return faqcommon,core.MysqlClient.Table("faqcommon").
		Where("name_cn LIKE ?", "%"+this.NameCn+"%").Find(&faqcommon).Error
}

//批量查询
func (this *FaqCommonPo) GetFAQs() (faqcommon []FaqCommonPo, err error) {
	log.Info("GetFAQs:",this)
	return faqcommon,core.MysqlClient.Table("faqcommon").
		Where("data_state = ?", this.DataState).Find(&faqcommon).Error
}
