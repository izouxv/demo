package storage

import (
	"github.com/jinzhu/gorm"
	"time"
	."notification/common"
)

type EmailTemplate  struct {
	Id          int64     `json:"id" gorm:"column:id;primary_key:true;unique;AUTO_INCREMENT"`
	Did         int64     `json:"did" gorm:"column:did"`
	Tid         int64     `json:"tid" gorm:"column:tid"`
	Subject     string    `json:"subject" gorm:"column:subject"`
	Html        string    `json:"html" gorm:"column:html"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time"`
}


/*BeforeCreate添加回调函数*/
func (n *EmailTemplate) BeforeCreate() (err error) {
	if err != nil {
		return
	}
	n.CreateTime = time.Now()
	n.UpdateTime = time.Now()
	return
}

/*BeforeUpdate修改回调函数*/
func (n *EmailTemplate) BeforeUpdate() {
	n.UpdateTime = time.Now()
}

/*增加邮件模板*/
func (n *EmailTemplate) CreateEmailTemplate(tx *gorm.DB) error {
	notice  := tx.Table("tbl_email_template").Create(n)
	if notice.Error != nil {
		return notice.Error
	}
	return nil
}

/*基于tid 的删除邮件模板*/
func (n *EmailTemplate) DeleteEmailTemplate(tx *gorm.DB) (err error) {
	rs := tx.Table("tbl_email_template").Where("id = ? and tid = ?", n.Id,n.Tid).Delete(&n)
	if rs.Error != nil {
	if rs.RowsAffected == 0 {
		return ErrDoesNotExist
		}
		return rs.Error
	}
	return nil
}

/*基于did 获取邮件模板数量*/
func (n *EmailTemplate) GetEmailTemplateCount(tx *gorm.DB) (count int32, err error) {
	if err := tx.Table("tbl_email_template").Where("tid = ?",n.Tid).Count(&count).Error; err != nil {
		return 0, err
	}
	return
}

/*基于tid 获取邮件模板列表*/
func (n *EmailTemplate) GetEmailTemplates(tx *gorm.DB,page ,count int32,order string ) (totalcount int32,templates []*EmailTemplate ,err error) {
	totalcount, err = n.GetEmailTemplateCount(tx)
	if err != nil {
		return
	}
	switch count {
	case -1:
		if err = tx.Table("tbl_email_template").Where("tid = ?", n.Tid).Order("id desc ").Find(&templates).Error; err != nil {
			return
		}
		return
	default:
		validatePageCountOrder(&page,&count)
		validateEmailNoticeOrderColumn(&order)
		offset := (page - 1) * count
		if err = tx.Table("tbl_email_template").Where("tid = ?", n.Tid).Order("id desc ").Limit(count).Offset(offset).Find(&templates).Error; err != nil {
			return
		}
		return
	}
}

/*基于tid 获取邮件模板*/
func (n *EmailTemplate) GetEmailTemplate(tx *gorm.DB) (err error) {
	if err = tx.Table("tbl_email_template").Where(" id = ? and tid = ? ", n.Id,n.Tid).Find(&n).Error; err != nil {
		return
	}
	return
}

/*基于tid 获取邮件模板*/
func (n *EmailTemplate) GetEmailTemplateById(tx *gorm.DB) (err error) {
	if err = tx.Table("tbl_email_template").Where(" id = ? ", n.Id).Find(&n).Error; err != nil {
		return err
	}
	return
}


/*基于did 修改邮件发送器信息*/
func (n *EmailTemplate) UpdateEmailTemplate(tx *gorm.DB) (err error) {
	res := tx.Table("tbl_email_template").Where("id = ? and tid = ?", n.Id,n.Tid).
		Updates(map[string]interface{}{"subject":n.Subject,"html":n.Html,"update_time":n.UpdateTime})
	if err = res.Error; err != nil {
		return err
	}
	return
}


