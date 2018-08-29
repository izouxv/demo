package storage

import (
	"github.com/jinzhu/gorm"
	."account-domain-rpc/common"
)

type DeviceType struct {
	Id          int32     `gorm:"column:id;primary_key;unique"`
	Tid         int64     `gorm:"column:tid"`
	DeviceType  string    `gorm:"column:device_type"`
	Status      int32     `gorm:"column:status"`
}


/*添加设备类型*/
func (f *DeviceType) CreateDeviceType(tx *gorm.DB) error {
	dt := tx.Table("tbl_device_type").Create(f)
	if dt.Error != nil {
		return dt.Error
	}
	return nil
}

/*获取设备类型数量*/
func (f *DeviceType) GetDeviceTypeCount(tx *gorm.DB) (count int32, err error) {
	if err = tx.Table("tbl_device_type").Where("tid = ?",f.Tid).Count(&count).Error; err != nil {
		return 0, err
	}
	return
}

/*分页获取设备类型信息*/
func (f *DeviceType) GetDeviceTypes(tx *gorm.DB, count, page int32,order string) (deciveType []*DeviceType, totalcount int32, err error) {
	totalcount, err = f.GetDeviceTypeCount(tx)
	if err != nil {
		return
	}
	validatePageCount(&page,&count)
	validateOrderColumn(&order)
	switch count {
	case  -1:
		r := tx.Table("tbl_device_type").Where("tid  = ?",f.Tid).Order(order).Find(&deciveType)
		if r.RowsAffected == 0 {
				return nil, 0,ErrDoesNotExist
			}
			return deciveType, totalcount,r.Error
	default :
		offset := (page - 1) * count
		r := tx.Table("tbl_device_type").Where("tid  = ?",f.Tid).Order(order).Limit(count).Offset(offset).Find(&deciveType)
		if r.RowsAffected == 0 {
			return nil, 0,ErrDoesNotExist
		}
		return deciveType, totalcount,r.Error
	}
	return
}

func validatePageCount(page, count *int32) {
	if *page == 0 {
		*page = 1
	}
	if *count <= 0 {
		*count = 10
	}
	if *count > 1000 {
		*count = 1000
	}
	return
}

var OrderColumn = []string{
	"id", "id desc", "id aes",
}



func validateOrderColumn(orderColumn *string) {
	var defaultOrderColumn = "id"
	validate := false
	for _, v := range OrderColumn {
		if v == *orderColumn {
			validate = true
		}
	}
	if !validate {
		orderColumn = &defaultOrderColumn
	}
}


