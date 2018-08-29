package db

import (
	log "github.com/cihub/seelog"
	"petfone-rpc/core"
	"time"
	"github.com/pkg/errors"
	"github.com/jinzhu/gorm"
)

//宠物信息
type PetInfoPo struct {
	Pid            	int32     	`gorm:"column:pid;primary_key;unique"`
	Uid          	int32		`gorm:"column:uid"`							//用户id
	Avatar         	string    	`gorm:"column:avatar"`
	Nickname       	string    	`gorm:"column:nickname"`
	Breed          	int32     	`gorm:"column:breed"`
	Gender         	int32     	`gorm:"column:gender"`
	Birthday       	time.Time 	`gorm:"column:birthday"`
	Weight         	float32		`gorm:"column:weight"`
	Somatotype     	int32     	`gorm:"column:somatotype"`
	Duration		int32		`gorm:"column:duration"`
	Brightness		int32		`gorm:"column:brightness"`
	CreationTime	time.Time 	`gorm:"column:creation_time"`
	UpdateTime		time.Time 	`gorm:"column:update_time"`
	DataState      	int32     	`gorm:"column:data_state"`
}

//插入宠物信息
func (this *PetInfoPo) SetPetInfoDB(dbc *gorm.DB, pettrain *PetTrainPo) ([]*PetTrainPo,error) {
	log.Info("SetPetInfoDB:", this)
	var petTrains []*PetTrainPo
	err := dbc.Table("pet").Create(this).Error
	if err != nil {
		return nil,err
	}
	//todo 添加宠物训练信息
	pettrain.Pid = this.Pid
	var key = 1
	for key <= len(core.Names) {
		pettrain.Name = core.Names[key]
		pettrain.Voice = core.Voices[key]
		num := pettrain.SetPetTrainDB(dbc)
		petTrains = append(petTrains,pettrain)
		log.Info("pettrain id:",pettrain)
		if  num != 1 {
			return nil,errors.New("插入训练信息失败")
		}
		key += int(num)
	}
	return petTrains,nil
}

//修改宠物信息
func (this *PetInfoPo) UpdatePetInfoDB() error {
	log.Info("UpdatePetInfoDB:",this)
	return core.MysqlClient.Table("pet").Where("pid = ? and data_state = 1", this.Pid).Update(this).Error
}

//删除宠物信息
func (this *PetInfoPo) DeletePetInfoDB(dbConn *gorm.DB) int64 {
	log.Info("DeletePetInfoDB:",this)
	return dbConn.Table("pet").Where("pid = ? and data_state = 1", this.Pid).
		Updates(map[string]interface{}{"update_time": this.UpdateTime, "data_state": 2}).RowsAffected
}

//查询宠物信息
func (this *PetInfoPo) GetPetInfoDB() error {
	log.Info("GetPetInfoDB:",this)
	return core.MysqlClient.Table("pet").Where("pid = ?", this.Pid).First(this).Error
}

//批量查询宠物信息
func (this *PetInfoPo) GetPetInfosDB(pids []int32) (petinfos []PetInfoPo, err error) {
	log.Info("GetPetInfosDB:",this)
	return petinfos, core.MysqlClient.Table("pet").Where("pid in (?) AND data_state = 1", pids).Find(&petinfos).Error
}

//批量查询宠物信息
func (this *PetInfoPo) GetPagePetInfoDB() (petinfos []PetInfoPo, err error) {
	log.Info("GetPagePetInfoDB-this:",this)
	return petinfos, core.MysqlClient.Table("pet").Where("data_state = ?", this.DataState).Find(&petinfos).Error
}

