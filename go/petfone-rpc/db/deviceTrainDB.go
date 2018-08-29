package db

import (
	"time"
	log "github.com/cihub/seelog"
	"github.com/jinzhu/gorm"
	"petfone-rpc/core"
)

type DeviceTrainPo struct {
	Id				int32		`gorm:"column:id;primary_key;unique"`
	Did				int32		`gorm:"column:did"`			//宠端设备id
	Name			string		`gorm:"column:name"`		//训练1，名称
	Num				int32    	`gorm:"column:num"`			//训练1，目标次数
	Counter			int32    	`gorm:"column:counter"`		//训练1，次数
	Voice			string    	`gorm:"column:voice"`		//训练1
	DevFid			uint32		`gorm:"dev_fid"`			//设备文件唯一id
	CreationTime   	time.Time	`gorm:"column:creation_time"`
	UpdateTime     	time.Time	`gorm:"column:update_time"`
	DataState		int32		`gorm:"column:data_state"`
}

//初始化
func (this *DeviceTrainPo) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", 0)
}

//跟随宠物信息插入训练信息
func (this *DeviceTrainPo) Create(db *gorm.DB) int64 {
	return db.Table("device_train").Create(this).RowsAffected
}

//修改宠物训练信息
func (this *DeviceTrainPo) UpdateDeviceTrainDB() error {
	log.Info("UpdateDeviceTrainDB:", this)
	return core.MysqlClient.Table("device_train").
		Where("id = ? AND did = ? AND data_state = 1", this.Id, this.Did).Update(&this).Error
}

//基于宠端设备名修改宠物训练信息
func (this *DeviceTrainPo) UpdateDeviceTrainForName() error {
	log.Info("UpdateDeviceTrainDB:", this)
	return core.MysqlClient.Table("device_train").
		Where("name = ? AND did = ? AND data_state = 1", this.Name, this.Did).Update(&this).Error
}

//计数宠物训练
func (this *DeviceTrainPo) CounterPetTrainDB() int64 {
	log.Info("CounterPetTrainDB:", this)
	return core.MysqlClient.Table("device_train").Where("id = ? AND did = ? AND data_state = 1",this.Id,this.Did).
		Updates(map[string]interface{}{"counter":gorm.Expr("counter+?",this.Counter),
		"update_time":this.UpdateTime}).RowsAffected
}

//删除宠物训练记录
func (this *DeviceTrainPo) DeletePetTrainsDB(dbConn *gorm.DB,startTime, endTime int64) error {
	log.Info("DeletePetTrainsDB:", this)
	return dbConn.Exec("UPDATE device_train SET data_state = 2 " +
		"WHERE did = ? AND data_state = 1 AND creation_time >= FROM_UNIXTIME(?) AND creation_time < FROM_UNIXTIME(?);",
		this.Did, startTime,endTime).Error
}

//查询宠物训练信息
func (this *DeviceTrainPo) GetPetTrainDB() error {
	log.Info("GetPetTrainDB:", this)
	return core.MysqlClient.Table("device_train").
		Where("id = ? AND did = ? AND data_state = 1", this.Id, this.Did).Last(&this).Error
}

//查询批量宠物时间范围内的训练信息
func (this *DeviceTrainPo) GetDevicesTrainsDB(did int32, startTime, endTime int64) (petTrains []*PetTrainPo,err error) {
	log.Info("GetPetTrainsDB:", did, startTime, endTime)
	return petTrains,core.MysqlClient.Table("device_train").
		Where("creation_time >= FROM_UNIXTIME(?) AND creation_time < FROM_UNIXTIME(?) AND did = ?",
		startTime, endTime,did).Find(&petTrains).Error
}

//查询宠物时间范围内的训练信息
func (this *DeviceTrainPo) GetDeviceTrainsDB(startTime, endTime int64) (petTrain []*PetTrainPo,err error) {
	log.Info("GetPetTrainsDB:", this,startTime, endTime)
	return petTrain,core.MysqlClient.Table("device_train").
		Where("did = ? AND creation_time >= FROM_UNIXTIME(?) AND creation_time < FROM_UNIXTIME(?)",
		this.Did, startTime, endTime).Find(&petTrain).Error
}

//查询宠物的训练信息
func (this *DeviceTrainPo) GetDeviceTrains() (petTrain []*PetTrainPo,err error) {
	row := core.MysqlClient.Table("device_train").Where("did = ? and data_state = 1",this.Did).Find(&petTrain)
	if gorm.ErrRecordNotFound == row.Error {
		log.Infof("GetDeviceTrains...device_train 表里没有找到数据")
		err = row.Error
	}
	if row.Error != nil {
		err = row.Error
	}
	return
}

