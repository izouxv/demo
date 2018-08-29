package db

import (
	log "github.com/cihub/seelog"
	"petfone-rpc/core"
	"time"
	"github.com/jinzhu/gorm"
)

//宠物品种对应的结构体
type BreedInfoPo struct {
	Id			int32		`gorm:"column:id"`
	NameCh		string		`gorm:"column:name_ch"`			//名称
	NameEn		string		`gorm:"column:name_en"`			//英文名称
	Address		string		`gorm:"column:address"`			//内容
	Types		int32		`gorm:"column:types"`
	Info		string		`gorm:"column:info"`
	UpdateTime	time.Time	`gorm:"column:update_time"`   //修改时间
}

//添加信息
func (this *BreedInfoPo) SetBreedinfoDB() error {
	log.Info("SetFilesinfoDB:", this)
	db := core.MysqlClient.Begin()
	if err := db.Exec(
		`insert ignore into
		breedinfo (name_ch, name_en, address, info, update_time) VALUES (?, ?, ?, ?, ?)`,
		this.NameCh, this.NameEn, this.Address, this.Info, this.UpdateTime).Error; err != nil {
		log.Info("SetFilesinfoDB-err:", err)
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}

//获取信息
func (this *BreedInfoPo) GetBreedinfoDB() error {
	log.Info("GetBreedinfoDB:", this)
	err := core.MysqlClient.Table("breedinfo").
		Where("id = ? AND types = ?", this.Id, this.Types).Select("id, name_ch, name_en, address").First(this).Error
	if err != nil {
		return err
	}
	return nil
}

//批量获取信息
func (this *BreedInfoPo) GetBreedinfosDB(number int32) (breedInfoPo []*BreedInfoPo, err error) {
	log.Info("GetBreedinfosDB:", this.Id,number)
	return breedInfoPo,core.MysqlClient.Table("breedinfo").
		Where("types = ? AND id >= ? AND id < ?", this.Types, this.Id,number).Order("id").Find(&breedInfoPo).Error
}

//修改信息
func (this *BreedInfoPo) UpdateInfoDB(tableName string,dbc *gorm.DB) error {
	log.Info("UpdateInfoDB:", this)
	return dbc.Table(tableName).Where("id = ?", this.Id).Update(&this).Error
}
