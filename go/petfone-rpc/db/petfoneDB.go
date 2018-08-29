package db

import (
	log "github.com/cihub/seelog"
	"petfone-rpc/core"
	"time"
	"github.com/jinzhu/gorm"
)

//设备表对应的结构体
type PetfonePo struct {
	Uid          int32     `gorm:"column:uid"`           //用户id
	Radius       int32     `gorm:"column:radius"`        //围栏半径
	CreationTime time.Time `gorm:"column:creation_time"` //创建时间
	UpdateTime   time.Time `gorm:"column:update_time"`   //修改时间
	DataState    int32     `gorm:"column:data_state"`    //数据状态
	Map			 int32     `gorm:"column:map"`  		 //用户地图
}

//添加业务信息
func (this *PetfonePo) SetPetfoneDB(dbc *gorm.DB) error {
	log.Info("UpdatePetfoneDB:", this)
	return dbc.Exec(`insert ignore into user_function (uid, radius, creation_time, update_time, data_state) 
		VALUES (?, ?, ?, ?, ?)`, this.Uid, this.Radius, this.CreationTime, this.UpdateTime, this.DataState).Error
}

//获取业务信息
func (this *PetfonePo) GetPetfoneDB() error {
	log.Info("GetDevicesDB:", this)
	return core.MysqlClient.Table("user_function").
		Where("uid = ? and data_state = ?", this.Uid, 1).
		Select("uid, radius,map").First(&this).Error
}

//修改业务信息
func (this *PetfonePo) UpdatePetfoneDB(pe PetfonePo) error {
	log.Info("UpdatePetfoneDB:", this)
	return core.MysqlClient.Table("user_function").Where("uid = ?", this.Uid).Update(pe).Error
}

//宠聊对话模板
type PetChatPo struct {
	Id          int32     	`gorm:"column:id"`          //id
	NameCn		string     	`gorm:"column:name_cn"`		//关键词
	InfoCn		string 		`gorm:"column:info_cn"`		//模版语句
	NameEn		string     	`gorm:"column:name_en"`		//关键词
	InfoEn		string 		`gorm:"column:info_en"`		//模版语句
	UpdateTime  time.Time 	`gorm:"column:update_time"` //修改时间
	DataState   int32     	`gorm:"column:data_state"`  //数据状态
}

//反向匹配
func (this *PetChatPo) GetPetChatKeyDB() (petChatPos []*PetChatPo,err error) {
	return petChatPos,core.MysqlClient.Raw(`SELECT * FROM pet_chat WHERE INSTR(?, name_cn);`, this.NameCn).Scan(&petChatPos).Error
}

//获取信息
func (this *PetChatPo) GetPetChatsDB() (petChatPos []*PetChatPo, err error) {
	log.Info("GetPetChatsDB-this:",this)
	return petChatPos,core.MysqlClient.Table("pet_chat").Where("data_state = 1").Scan(&petChatPos).Error
}

//设备表对应的结构体
type PetChatMsgPo struct {
	Id          int64     	`gorm:"column:id"`			//id
	MsgSource	int32     	`gorm:"column:msg_source"`	//内容来源
	MsgAbout	string     	`gorm:"column:msg_about"`	//内容相关
	Msg			string		`gorm:"column:msg"`			//内容
	CreationTime time.Time 	`gorm:"column:creation_time"` //创建时间
}

func (this *PetChatMsgPo) SetMsg() {
	err := core.MysqlClient.Table("pet_chat_msg").Create(this).Error
	if err != nil {
		log.Error("SetMsg-err:",err)
	}
}

//宠聊对话模板
type PetChatKeyPo struct {
	Id          int32     	`gorm:"column:id"`          //id
	InfoCn		string 		`gorm:"column:info_cn"`		//模版语句
	InfoEn		string 		`gorm:"column:info_en"`		//模版语句
	CreateTime  time.Time 	`gorm:"column:create_time"` //修改时间
	DataState   int32     	`gorm:"column:data_state"`  //数据状态
}

//获取信息
func (this *PetChatKeyPo) GetPetChatKeysDB() (petChatKeys []*PetChatKeyPo, err error) {
	log.Info("GetPetChatKeysDB-this:",this)
	return petChatKeys, core.MysqlClient.Table("pet_chat_key").Where("data_state = 1").Scan(&petChatKeys).Error
}