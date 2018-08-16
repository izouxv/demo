package storage

import (
	"github.com/jinzhu/gorm"
	"time"
	."notification/common"
)

type EmailSender  struct {
	Id           int64     `json:"id" gorm:"column:id;primary_key:true;unique;AUTO_INCREMENT"`
	Did          int64     `json:"did" gorm:"column:did"`
	SmtpServer   string    `json:"smtp_server" gorm:"column:smtp_server"`
	EmailSender  string    `json:"sender_email" gorm:"column:email_sender"`
	Username     string    `json:"username" gorm:"column:username"`
	Password     string    `json:"password" gorm:"column:password"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime   time.Time `json:"update_time" gorm:"column:update_time"`
}

var noticeOrderColumn = []string{
	"id", "id desc", "id asc",
	"create_time", "create_time desc", "create_time asc",
	"update_time", "update_time desc", "update_time asc",
}

func validateEmailNoticeOrderColumn(orderColumn *string) {
	var defaultOrderColumn = "id"
	validate := false
	for _, v := range noticeOrderColumn {
		if v == *orderColumn {
			validate = true
		}
	}
	if !validate {
		orderColumn = &defaultOrderColumn
	}
}

/*BeforeCreate添加回调函数*/
func (n *EmailSender) BeforeCreate() (err error) {
	if err != nil {
		return
	}
	n.CreateTime = time.Now()
	n.UpdateTime = time.Now()
	return
}

/*BeforeUpdate修改回调函数*/
func (n *EmailSender) BeforeUpdate() {
	n.UpdateTime = time.Now()
}

/*增加邮件发送器*/
func (n *EmailSender) CreateEmailSender(tx *gorm.DB) error {
	notice  := tx.Table("tbl_email_notice").Create(n)
	if notice.Error != nil {
		return notice.Error
	}
	return nil
}

/*基于did 的sp中DeleteAlarm删除邮件发送器*/
func (n *EmailSender) DeleteEmailSender(tx *gorm.DB) (err error) {
	rs := tx.Table("tbl_email_notice").Where("id = ? and did = ?", n.Id,n.Did).Delete(&n)
	if rs.Error != nil {
	if rs.RowsAffected == 0 {
		return ErrDoesNotExist
		}
		return rs.Error
	}
	return nil
}


/*基于did 获取邮件发送器信息*/
func (n *EmailSender) GetEmailSender(tx *gorm.DB) (err error) {
	if err = tx.Table("tbl_email_notice").Where("id = ? and did = ?", n.Id,n.Did).Find(&n).Error; err != nil {
		return err
	}
	return
}
/*基于did 获取邮件发送器信息*/
func (n *EmailSender) GetEmailSenderByDid (tx *gorm.DB) (err error) {
	if err = tx.Table("tbl_email_notice").Where("did = ?", n.Did).Find(&n).Error; err != nil {
		return err
	}
	return
}



/*基于did 获取邮件发送器数量*/
func (n *EmailSender) GetEmailSenderCount(tx *gorm.DB) (count int32, err error) {
	if err := tx.Table("tbl_email_notice").Where("did = ?",n.Did).Count(&count).Error; err != nil {
		return 0, err
	}
	return count,err
}

/*基于did 获取邮件发送器列表*/
func (n *EmailSender) GetEmailSenders(tx *gorm.DB, page ,count int32,order string) (notices []*EmailSender , totalcount int32,err error) {
	totalcount, err = n.GetEmailSenderCount(tx)
	if err != nil {
		return
	}
	switch count{
	case -1:
		if err = tx.Table("tbl_email_notice").Where("did = ?", n.Did).Order("id desc").Find(&notices).Error; err != nil {
			return
		}
		return
	default :
		validatePageCountOrder(&page,&count)
		validateEmailNoticeOrderColumn(&order)
		offset := (page - 1) * count
		if err = tx.Table("tbl_email_notice").Where("did = ?", n.Did).Order("id desc").Limit(count).Offset(offset).Find(&notices).Error; err != nil {
			return
		}
		return
	}
}

/*基于did 修改邮件发送器信息*/
func (n *EmailSender) UpdateEmailSender(tx *gorm.DB) (err error) {
	res := tx.Table("tbl_email_notice").Where("id = ? and did = ?", n.Id,n.Did).
		Updates(map[string]interface{}{"smtp_server":n.SmtpServer,"email_sender":n.EmailSender,"username":n.Username,"password":n.Password,"update_time":n.UpdateTime})
	if err = res.Error; err != nil {
		return err
	}
	return
}
