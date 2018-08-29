package storage

import (
	log "github.com/cihub/seelog"
	"github.com/jinzhu/gorm"
	"time"
)

type FeedBack struct {
	Id          int32     `gorm:"column:id;primary_key;unique"`
	Tid         int64     `gorm:"column:tid"`
	DeviceInfo  string    `gorm:"column:device_info"`
	AppInfo     string    `gorm:"column:app_info"`
	UserInfo    string    `gorm:"column:user_info"`
	MobileInfo  string    `gorm:"column:mobile_info"`
	ExtendInfo  string    `gorm:"column:extend_info"`
	Description string    `gorm:"column:description"`
	Files       string    `gorm:"column:files"`
	Contact     string    `gorm:"column:contact"`
	Type        int32     `gorm:"column:bug_type"`
	CreateTime  time.Time `gorm:"column:create_time"`
	UpdateTime  time.Time `gorm:"column:update_time"`
}

/*BeforeCreate添加回调函数*/
func (f *FeedBack) BeforeCreate() (err error) {
	f.UpdateTime = time.Now()
	f.CreateTime = time.Now()
	return
}

/*BeforeUpdate修改回调函数*/
func (f *FeedBack) BeforeUpdate() {
	f.UpdateTime = time.Now()
}

/*添加反馈*/
func (f *FeedBack) CreateFeedBack(tx *gorm.DB) error {
	feedback := tx.Table("tbl_feedback").Create(f)
	if feedback.Error != nil {
		log.Errorf("err :(%s)", feedback.Error)
		return feedback.Error
	}
	return nil
}

/*添加反馈*/
func (f *FeedBack) CreateFeedBackBaseTenant(tx *gorm.DB) error {
	feedback := tx.Table("tbl_feedback").Create(f)
	if feedback.Error != nil {
		log.Errorf("err :(%s)", feedback.Error)
		return feedback.Error
	}
	return nil
}



/*GetFeedbackCount获取反馈数量*/
func (f *FeedBack) GetFeedbackCount(tx *gorm.DB) (count int32, err error) {
	if err := tx.Table("tbl_feedback").Where("tid = ?",f.Tid).Count(&count).Error; err != nil {
		log.Error("Get feedback count to db error:", err)
		return 0, err
	}
	return
}

/*GetFeedbacks分页获取反馈信息*/
func (f *FeedBack) GetFeedbacks(tx *gorm.DB, count, page int32) (feeedbacks []*FeedBack, totalcount int32, err error) {
	totalcount, err = f.GetFeedbackCount(tx)
	if err != nil {
		log.Error("Get FeedbackCount error:", err)
		return
	}
	if err = tx.Table("tbl_feedback").Where("tid  = ?",f.Tid).Order("id desc ").Limit(count).Offset((page - 1) * count).Find(&feeedbacks).Error; err != nil {
		log.Error("Get FeedbackByPage  info to db error:", err)
		return
	}
	return
}

/*GetFeedbackCount分类获取反馈数量*/
func (f *FeedBack) GetFeedbackCountByType(tx *gorm.DB) (count int32, err error) {
	if err := tx.Table("tbl_feedback").Where("bug_type = ? and tid = ?",f.Type,f.Tid).Count(&count).Error; err != nil {
		log.Error("Get feedback count to db error:", err)
		return 0, err
	}
	return
}

/*分类获取反馈信息*/
func (f *FeedBack) GetFeedbackByType(tx *gorm.DB, count, page int32) (feedBacks []*FeedBack, totalcount int32, err error) {
	totalcount, err = f.GetFeedbackCountByType(tx)
	if err != nil {
		log.Error("GetFeedbackCountByType  error", err)
		return
	}
	if err = tx.Table("tbl_feedback").Where("tid = ? and bug_type = ?",f.Tid,f.Type).Order("id desc ").Limit(count).Offset((page - 1) * count).Find(&feedBacks).Error; err != nil {
		log.Error("Get FeedbackByType  info to db error:", err)
		return
	}
	return
}

/*GetFeedbacks获取反馈信息*/
func (f *FeedBack) GetFeedback(tx *gorm.DB) (err error) {
	if err = tx.Table("tbl_feedback").Where("id = ? and tid = ?", f.Id,f.Tid).Find(&f).Error; err != nil {
		log.Error("Get feedback info to db error:", err)
		return
	}
	return
}



/*GetFeedbacks获取反馈信息*/
func (f *FeedBack) DelFeedback(tx *gorm.DB,ids []int32) (err error) {
	if err = tx.Table("tbl_feedback").Delete(f,"id  in (?) and tid = ?", ids,f.Tid).Error; err != nil {
		log.Error("删除工单失败:", err)
		return
	}
	return
}

/*批量获取*/
func (f *FeedBack) BatchFeedback(tx *gorm.DB,ids []int32) (feedBacks []*FeedBack, err error) {
	if err = tx.Table("tbl_feedback").Where("tid = ? and id in (?) ",f.Tid,ids).Find(&feedBacks).Error; err != nil {
		log.Error("批量获取要删除的工单信息:", err)
		return
	}
	return
}



