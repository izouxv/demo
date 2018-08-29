package storage

import (
	"time"
	"github.com/jinzhu/gorm"
	"errors"
	log "github.com/cihub/seelog"
)
var (
	TestUserAlreadyExists = errors.New("testUser already exists")
	TestUserNotExists     = errors.New("testUser not exists")
)

type TestUser struct {
	Id         int32     `gorm:"column:id"`
	Tid        int64     `gorm:"column:tid"`
	UserName   string    `gorm:"column:username"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

/*BeforeCreate添加回调函数*/
func (f *TestUser) BeforeCreate() (err error) {
	f.UpdateTime = time.Now()
	f.CreateTime = time.Now()
	return
}

/*BeforeUpdate修改回调函数*/
func (f *TestUser) BeforeUpdate() {
	f.UpdateTime = time.Now()
}

func (f *TestUser) CreateTestUser(tx *gorm.DB) error {
	feedback := tx.Table("tbl_testuser").Create(f)
	if feedback.Error != nil {
		return feedback.Error
	}
	return nil
}

func (f *TestUser) GetTestUserCount(tx *gorm.DB) (count int32, err error) {
	if err := tx.Table("tbl_testuser").Where("tid = ?",f.Tid).Count(&count).Error; err != nil {
		return 0, err
	}
	return
}

func (f *TestUser) GetTestUsers(tx *gorm.DB, count, page int32) (testusers []*TestUser, totalcount int32, err error) {
	totalcount, err = f.GetTestUserCount(tx)
	if err != nil {
		return nil,0,err
	}
	if err = tx.Table("tbl_testuser").Where("tid  = ?",f.Tid).Order("id desc ").Limit(count).Offset((page - 1) * count).Find(&testusers).Error; err != nil {
		return nil,0,err
	}
	return
}

func (f *TestUser) GetTestUser(tx *gorm.DB) (err error) {
	reply :=tx.Raw("SELECT * FROM tbl_testuser where tid = ? and id = ? ",f.Tid,f.Id).Scan(&f)
	if reply.Error != nil {
		if reply.RowsAffected == 0 {
			return TestUserNotExists
		}
		return err
	}
	return
}

func (f *TestUser) DelTestUser(tx *gorm.DB) (err error) {
	if err = tx.Table("tbl_testuser").Where("id = ? and tid = ?", f.Id,f.Tid).Delete(&f).Error; err != nil {
		return
	}
	return
}

func (f  *TestUser) UpdateTestUser(tx *gorm.DB)(err  error) {
	reply := tx.Table("tbl_testuser").Where("id = ? and tid = ? ", f.Id,f.Tid).Update(&f)
	if reply.Error != nil {
		if reply.RowsAffected == 0 {
			return TestUserNotExists
		}
		return reply.Error
	}
	return
}

func (f *TestUser) GetTestUserByUsername(tx *gorm.DB) (err error) {
	reply :=tx.Raw("SELECT * FROM tbl_testuser where tid = ? and username = ? ORDER BY id DESC LIMIT 1",f.Tid,f.UserName).Scan(&f)
	 if reply.Error != nil {
		 if reply.RowsAffected == 0 {
			 return TestUserNotExists
		 }
		 return reply.Error
	}
	return nil
}


func (f *TestUser) CheckTestUserIsExist(tx *gorm.DB) (*TestUser ,error) {
	reply :=tx.Table("tbl_testuser").Where("tid = ? and username = ?",f.Tid,f.UserName).Find(&f)
	if reply.Error != nil {
		log.Errorf("查询用户名是否已存在有误:",reply.Error)
		return nil,reply.Error
	}
	return f,nil
}



func (f *TestUser) CheckTestUserIsExistPut(tx *gorm.DB) (*TestUser ,error) {
	reply :=tx.Table("tbl_testuser").Where("tid = ? and username = ? and id != ?",f.Tid,f.UserName,f.Id).Find(&f)
	if reply.Error != nil {
		log.Errorf("查询用户名是否已存在有误:",reply.Error)
		return nil,reply.Error
	}
	return f,nil
}




