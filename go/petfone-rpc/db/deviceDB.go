package db

import (
	log "github.com/cihub/seelog"
	"petfone-rpc/core"
	"time"
	"petfone-rpc/pb"
	"github.com/jinzhu/gorm"
)

//设备表对应的结构体
type DevicePo struct {
	Did          	int32		`gorm:"column:did;primary_key;unique"`		//设备id
	Uid          	int32		`gorm:"column:uid"`							//用户id
	Sn				string    	`gorm:"column:sn"`							//设备sn
	Types			pb.DeviceTypes	`gorm:"column:types"`					//设备类型
	DeviceMac	 	string		`gorm:"device_mac"`							//设备mac
	DeviceName		string		`gorm:"device_name"`						//设备名称
	DeviceVersion	string		`gorm:"device_version"`						//设备版本
	SoftwareVersion	string		`gorm:"software_version"`					//软件名称
	CreationTime 	time.Time	`gorm:"column:creation_time"` 				//创建时间
	UpdateTime   	time.Time 	`gorm:"column:update_time"`   				//修改时间
	DataState    	int32		`gorm:"column:data_state"`    				//数据状态
	LedModel        int8        `gorm:"column:led_model;default:1"`    		//设备灯闪烁模式
	LedColor       	int32       `gorm:"column:led_color;default:153"`       //设备灯颜色
	LedLight        int8        `gorm:"column:led_light;default:100"`    	//设备灯亮度
	LedState        int8        `gorm:"column:led_state;default:0"`    		//宠物设备灯开关状态
	AudioId         int8		`gorm:"column:audio_id;default:1"`    		//宠物播放录音的id
}

//绑定设备
func (this *DevicePo) SetDeviceDB(dbc *gorm.DB) error {
	log.Info("SetPetDeviceDB:", this)
	return dbc.Table("device").Create(&this).Error
}

//修改设备
func (this *DevicePo) UpdateDeviceDB(dbc *gorm.DB) error {
	log.Info("UpdateDeviceDB:", this)
	return dbc.Table("device").Where("did = ? AND data_state = ?", this.Did,1).Update(&this).Error
}

//批量did查询
func (this *DevicePo) GetDevicesDB(dids []int32) (devicePos []*DevicePo, err error) {
	log.Info("GetDevicesDB")
	return devicePos,core.MysqlClient.Table("device").
		Where("did in (?) AND data_state = 1", dids).Find(&devicePos).Error
}

//分页查询
func (this *DevicePo) GetPageDeviceDB(startId,count int32, sort string) (devicePos []*DevicePo, totalCount int32, err error) {
	log.Info("GetPageDeviceDB:",startId,count,sort)
	return devicePos, totalCount, core.MysqlClient.Table("device").Where("data_state = 1").
		Order("did "+sort).Count(&totalCount).Offset(startId).Limit(count).Find(&devicePos).Error
}

//查询did
func (this *DevicePo) GetDeviceByDid() error {
	log.Info("GetDeviceByDid:", this)
	return core.MysqlClient.Table("device").Where("did = ? AND data_state = 1", this.Did).First(&this).Error
}

//查询SN
func (this *DevicePo) GetDeviceBySn() error {
	log.Info("GetDeviceBySn:", this)
	return core.MysqlClient.Table("device").Where("sn = ? AND data_state = 1", this.Sn).Find(&this).Error
}

//查询mac
func (this *DevicePo) GetDeviceByMac() error {
	log.Info("GetDeviceByMac:", this)
	return core.MysqlClient.Table("device").
		Where("device_mac = ? AND data_state = 1", this.DeviceMac).Find(&this).Error
}

//修改设备
func (this *DevicePo) UpdateDeviceByDid(tx *gorm.DB) (err error) {
	res := tx.Table("device").Where("did = ?", this.Did).Update(&this)
	//res := tx.Table("device").Where("did = ?", this.Did).
	//	Update(map[string]interface{}{"led_model": this.LedModel, "led_color": this.LedColor, "led_light": this.LedLight,"led_state":this.LedState,"audio_id":this.AudioId,"update_time":this.UpdateTime})
	if err := res.Error; err != nil {
		log.Errorf("修改设备信息失败 %s", err)
		return err
	}
	return
}
