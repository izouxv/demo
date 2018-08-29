package db

import (
	"time"
	log "github.com/cihub/seelog"
	"github.com/jinzhu/gorm"
	"petfone-rpc/core"
)

type PetTrainPo struct {
	Id				int32		`gorm:"column:id;primary_key;unique"`
	Pid				int32		`gorm:"column:pid"`			//宠物id
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
func (this *PetTrainPo) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", 0)
}

//跟随宠物信息插入训练信息
func (this *PetTrainPo) SetPetTrainDB(db *gorm.DB) int64 {
	return db.Table("pet_train").Create(this).RowsAffected
}

//修改宠物训练信息
func (this *PetTrainPo) UpdatePetTrainDB() error {
	log.Info("GetPetTrainDB:", this)
	return core.MysqlClient.Table("pet_train").
		Where("id = ? AND pid = ? AND data_state = 1", this.Id, this.Pid).Update(&this).Error
}

//计数宠物训练
func (this *PetTrainPo) CounterPetTrainDB() int64 {
	log.Info("CounterPetTrainDB:", this)
	return core.MysqlClient.Table("pet_train").Where("id = ? AND pid = ? AND data_state = 1",this.Id,this.Pid).
		Updates(map[string]interface{}{"counter":gorm.Expr("counter+?",this.Counter),
		"update_time":this.UpdateTime}).RowsAffected
}

//删除宠物训练记录
func (this *PetTrainPo) DeletePetTrainsDB(dbConn *gorm.DB,startTime, endTime int64) error {
	log.Info("DeletePetTrainsDB:", this)
	return dbConn.Exec("UPDATE pet_train SET data_state = 2 " +
		"WHERE pid = ? AND data_state = 1 AND creation_time >= FROM_UNIXTIME(?) AND creation_time < FROM_UNIXTIME(?);",
		this.Pid, startTime,endTime).Error
}

//查询宠物训练信息
func (this *PetTrainPo) GetPetTrainDB() error {
	log.Info("GetPetTrainDB:", this)
	return core.MysqlClient.Table("pet_train").
		Where("id = ? AND pid = ? AND data_state = 1", this.Id, this.Pid).Last(&this).Error
}

//查询批量宠物时间范围内的训练信息
func (this *PetTrainPo) GetPetsTrainsDB(pids []int32, startTime, endTime int64) (petTrains []*PetTrainPo,err error) {
	log.Info("GetPetTrainsDB:", pids, startTime, endTime)
	return petTrains,core.MysqlClient.Table("pet_train").
		Where("creation_time >= FROM_UNIXTIME(?) AND creation_time < FROM_UNIXTIME(?) AND pid IN (?)",
		startTime, endTime,pids).Find(&petTrains).Error
}

//查询宠物时间范围内的训练信息
func (this *PetTrainPo) GetPetTrainsDB(startTime, endTime int64) (petTrain []*PetTrainPo,err error) {
	log.Info("GetPetTrainsDB:", this,startTime, endTime)
	return petTrain,core.MysqlClient.Table("pet_train").
		Where("pid = ? AND creation_time >= FROM_UNIXTIME(?) AND creation_time < FROM_UNIXTIME(?)",
		this.Pid, startTime, endTime).Find(&petTrain).Error
}

//查询宠物时间范围内的训练信息
func (this *PetTrainPo) IsExistPetTrains() (petTrain []*PetTrainPo,err error) {
	row := core.MysqlClient.Table("pet_train").Where("pid = ? and data_state = 1",this.Pid).Find(&petTrain)
	if gorm.ErrRecordNotFound == row.Error {
		log.Infof("IsExistPetTrains...pet_train 表里没有找到数据")
		err = row.Error
	}
	if row.Error != nil {
		err = row.Error
	}
	return
}

