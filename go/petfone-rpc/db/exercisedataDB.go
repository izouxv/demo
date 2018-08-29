package db

import (
	"time"
	"petfone-rpc/core"
	log "github.com/cihub/seelog"
	"github.com/jinzhu/gorm"
)

type ExerciseDataPetPo struct {
	Id				int64		`gorm:"column:id;primary_key;unique"`
	Uid    			int32		`gorm:"column:uid"`
	Pid    			int32		`gorm:"column:pid"`
	Pdid    		int32		`gorm:"column:pdid"`
	Cals       		float32		`gorm:"column:cals"`			//消耗卡
	ReportTime    	time.Time	`gorm:"column:report_time"`		//上报时间
	Coordinates  	string		`gorm:"column:coordinates"`		//坐标与运动信息
	CreationTime 	time.Time	`gorm:"column:creation_time"` 	//创建时间
	DataState    	int32		`gorm:"column:data_state"`    	//数据状态
}

//存数据
func (this *ExerciseDataPetPo) SetExerciseDataPet(dbc *gorm.DB) error {
	log.Info("SetExerciseDataPet-this:", this)
	return dbc.Table("exercise_pet").Create(this).Error
}

//查询时间段数据
func (this *ExerciseDataPetPo) GetExerciseDataPet(start int64, end int64) (exers []*ExerciseDataPetPo, err error) {
	log.Info("GetExerciseDataPet-this:", this)
	err = core.MysqlClient.Table("exercise_pet").
		Where("pid = ? AND report_time >= FROM_UNIXTIME(?) AND report_time < FROM_UNIXTIME(?)", this.Pid, start, end).
			Find(&exers).Error
	return exers, err
}

func (this *ExerciseDataPetPo) GetExerciseDataPetTimes() (exers []*ExerciseDataPetPo, err error) {
	log.Info("GetExerciseDataPet-this:", this)
	rows, err := core.MysqlClient.Exec(
		`(SELECT id, pid, pdid, uid, cals, coordinates,report_time FROM exercise_pet 
				WHERE pid = 1 AND data_state= 1 order by id LIMIT 1)
		UNION ALL
			(SELECT id, pid, pdid, uid, cals, coordinates,report_time FROM exercise_pet 
				WHERE pid = 1 AND data_state= 1 order by id DESC LIMIT 1)`).Rows()
	if err != nil {
		log.Info("GetExerciseDataPet-err:", err)
		return nil,err
	}
	defer rows.Close()
	for rows.Next() {
		exer := &ExerciseDataPetPo{}
		core.MysqlClient.ScanRows(rows,exer)
		exers = append(exers,exer)
	}
	return exers, err
}

type ExerciseDataUserPo struct {
	Id				int64		`gorm:"column:id;primary_key;unique"`
	Uid    			int32		`gorm:"column:uid"`
	Udid    		int32		`gorm:"column:udid"`
	Steps       	int32		`gorm:"column:steps"`			//运动步数
	Cals       		float32		`gorm:"column:cals"`			//消耗卡
	CreationTime 	time.Time	`gorm:"column:creation_time"` 	//创建时间
	DataState    	int32		`gorm:"column:data_state"`    	//数据状态
}

func (this *ExerciseDataUserPo) SetExerciseDataUser(dbc *gorm.DB) error {
	log.Info("SetExerciseDataUser-this:", this)
	return dbc.Table("exercise_user").Create(this).Error
}