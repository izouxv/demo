package db

import (
	"github.com/jinzhu/gorm"
	log "github.com/cihub/seelog"
)

//创建新的宠物训练记录，并清除过期的记录
func SetClearPetTrainDB(dbc *gorm.DB,start, end, expiration int64) error {
	log.Info("Create-train-all-start:",start,",end:",end,",expiration:",expiration)
	err := dbc.Exec("INSERT INTO pet_train (pid, name, voice, dev_fid, num, creation_time, update_time, data_state) " +
		"SELECT p.pid, p.name, p.voice, p.dev_fid, p.num,NOW(), NOW() , 1 FROM pet_train p " +
		"WHERE data_state = 1 AND p.creation_time >= FROM_UNIXTIME(?) AND p.creation_time < FROM_UNIXTIME(?)",start,end).Error
	err = dbc.Exec("UPDATE pet_train SET data_state = 2 " +
		"WHERE data_state = 1 AND creation_time >= FROM_UNIXTIME(?) AND creation_time < FROM_UNIXTIME(?)",start,end).Error
	err = dbc.Table("pet_train").Where("creation_time < FROM_UNIXTIME(?)", expiration).Delete(PetTrainPo{}).Error
	return err
}

//清除过期的关系记录
func ClearRelationDB(dbc *gorm.DB, expiration int64) error {
	log.Info("SetClearRelationDB expiration:",expiration)
	err := dbc.Table("user_device").Where(
		"data_state = 2 AND creation_time < FROM_UNIXTIME(?)", expiration).Delete(UserDevicePo{}).Error
	err = dbc.Table("user_pet").Where(
		"data_state = 2 AND creation_time < FROM_UNIXTIME(?)", expiration).Delete(UserDevicePo{}).Error
	err = dbc.Table("device_pet").Where(
		"data_state = 2 AND creation_time < FROM_UNIXTIME(?)", expiration).Delete(UserDevicePo{}).Error
	return err
}

//清除过期的关系记录
func ClearActionDB(dbc *gorm.DB, expiration int64) error {
	log.Info("SetClearRelationDB expiration:",expiration)
	err := dbc.Table("action_log").Where(
		"creation_time < FROM_UNIXTIME(?)", expiration).Delete(ActionLog{}).Error
	return err
}
