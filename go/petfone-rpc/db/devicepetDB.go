package db

import (
	"petfone-rpc/core"
	"time"
	"github.com/jinzhu/gorm"
	log "github.com/cihub/seelog"
	"petfone-rpc/pb"
)
/**
用户设备关系
 */
type UserDevicePo struct {
	Id           int32        `gorm:"column:id;primary_key;unique"`
	Uid          int32        `gorm:"column:uid"`           //用户id
	Did          int32        `gorm:"column:did"`           //设备id
	Types        pb.DeviceTypes		`gorm:"column:types"`	//设备类型
	Permit       pb.DevPermit `gorm:"column:permit"`        //是否被分享次数
	CreationTime time.Time    `gorm:"column:creation_time"` //创建时间
	UpdateTime   time.Time    `gorm:"column:update_time"`   //修改时间
	DataState    int32        `gorm:"column:data_state"`    //数据状态
}

//以第一个切片为主去除第二个切片中的重复元素
func (this *UserDevicePo) SliceDiff1(slice1, slice2 []*UserDevicePo) (diffslice []*UserDevicePo) {
	for _, v := range slice1 {if v.Permit == 0 && inSliceIf1(v.Did, slice2) {v.Uid = this.Uid
		v.Permit = 1;diffslice = append(diffslice, v)}};return
}
func inSliceIf1(id int32, sl []*UserDevicePo) bool {
	for _, vv := range sl {if vv.Did == id {return false}};return true
}

//初始化
func (this *UserDevicePo) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("id", 0)
	return nil
}

//绑定用户与设备
func (this *UserDevicePo) SetUserDeviceDB(dbc *gorm.DB) int64 {
	log.Info("SetUserDeviceDB:", this)
	return dbc.Table("user_device").Create(&this).RowsAffected
}

//批量插入用户与设备关系
func (this *UserDevicePo) SetUserDevicesDB(dbc *gorm.DB,userDevicePos []*UserDevicePo) error {
	log.Info("SetUserDevicesDB")
	var err error
	for _, v := range userDevicePos {
		err = dbc.Table("user_device").Create(&v).Error
		if  err != nil {
			log.Info("SetUserDevicesDB-err:", err)
			return err
		}
	}
	return nil
}

//解绑用户与设备
func (this *UserDevicePo) DeleteUserDeviceDB(dbc *gorm.DB) int64 {
	log.Info("DeleteUserDeviceDB:", this)
	return dbc.Table("user_device").Where("uid = ? AND did = ? AND data_state = 1", this.Uid, this.Did).
		Updates(map[string]interface{}{"data_state": 2, "update_time": this.UpdateTime}).RowsAffected
}

//批量解绑用户与设备
func (this *UserDevicePo) DeleteUserDevicesDB(dbc *gorm.DB,dids []int32) int64 {
	log.Info("DeleteUserDevicesDB:", this,dids)
	return dbc.Table("user_device").Where("uid = ? AND did in (?) AND data_state = 1",this.Uid,dids).
		Updates(map[string]interface{}{"update_time":this.UpdateTime,"data_state":2}).RowsAffected
}

//查询设备关系
func (this *UserDevicePo) GetUsersDeviceDB() (userDevicePos []*UserDevicePo,err error) {
	log.Info("GetUsesDeviceDB:", this)
	return userDevicePos,core.MysqlClient.Table("user_device").
		Where("did = ? AND data_state = 1",this.Did).Find(&userDevicePos).Error
}

//查询用户关系
func (this *UserDevicePo) GetUserDevicesDB() (userDevicePos []*UserDevicePo,err error) {
	log.Info("GetUserDevicesDB:", this)
	return userDevicePos,core.MysqlClient.Table("user_device").
		Where("uid = ? AND data_state = 1",this.Uid).Find(&userDevicePos).Error
}

//查询用户设备关系
func (this *UserDevicePo) GetUserDeviceDB() (err error) {
	log.Info("GetUserDeviceDB:", this)
	return core.MysqlClient.Table("user_device").Where("uid = ? AND did = ? AND data_state = 1",
		this.Uid, this.Did).Order("id DESC").Limit(1).First(&this).Error
}

//查询设备主人
func (this *UserDevicePo) GetDeviceMasterDB() error {
	log.Info("GetDeviceMasterDB:", this)
	return core.MysqlClient.Table("user_device").
		Where("did = ? AND permit = 0 AND data_state = 1", this.Did).First(&this).Error
}

//查询用户绑定的宠端设备
func (this *UserDevicePo) GetUserDevicesForPet() (err error) {
	log.Info("GetUserDevicesDB:", this)
	return core.MysqlClient.Table("user_device").
		Where("uid = ? AND types = ? AND data_state = 1",this.Uid,this.Types).Find(&this).Error
}

//批量查询设备主人
func (this *UserDevicePo) GetBatchDeviceMasterDB(dids []int32) (userDevicePos []*UserDevicePo,err error) {
	log.Info("GetBatchDeviceMasterDB")
	return userDevicePos, core.MysqlClient.Table("user_device").
		Where("did in (?) AND permit = 0 AND data_state = 1", dids).Find(&userDevicePos).Error
}

/**
用户宠物关系
 */
type UserPetPo struct {
	Id           int32        `gorm:"column:id;primary_key;unique"`
	Uid          int32        `gorm:"column:uid"`           //用户id
	Pid          int32        `gorm:"column:pid"`           //宠物id
	Permit       pb.PetPermit `gorm:"column:permit"`		//权限级别
	CreationTime time.Time    `gorm:"column:creation_time"` //创建时间
	UpdateTime   time.Time    `gorm:"column:update_time"`   //修改时间
	DataState    int32        `gorm:"column:data_state"`    //数据状态
}

func (this *UserPetPo) SliceDiff2(slice1, slice2 []*UserPetPo) (diffslice []*UserPetPo) {
	for _, v := range slice1 {if InSliceIf2(v, slice2) {v.Uid = this.Uid
		v.Permit = 1;diffslice = append(diffslice, v)}};return
}
func InSliceIf2(v *UserPetPo, sl []*UserPetPo) bool {
	for _, vv := range sl {if vv.Pid == v.Pid && v.Permit == 0 {return false}};return true
}

//初始化
func (this *UserPetPo) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("id", 0)
	return nil
}

//绑定用户与宠物
func (this *UserPetPo) SetUserPetDB(dbc *gorm.DB) int64 {
	log.Info("RelevanceUserPetDB:", this)
	return dbc.Table("user_pet").Create(&this).RowsAffected
}

//批量插入用户与宠物关系
func (this *UserPetPo) SetUserPetsDB(dbc *gorm.DB, userPetPos []*UserPetPo) error {
	log.Info("SetUserPetsDB:", this)
	var err error
	for _, v := range userPetPos {
		err = dbc.Table("user_pet").Create(&v).Error
		if  err != nil {
			log.Info("SetUserPetsDB-err:", err)
			return err
		}
	}
	return nil
}

//删除用户关系
func (this *UserPetPo) DeleteUserPetDB(dbConn *gorm.DB) error {
	log.Info("DeleteUserPetDB:", this)
	return dbConn.Table("user_pet").Where("uid = ? AND pid = ? AND data_state = 1",this.Uid,this.Pid).
		Updates(map[string]interface{}{"update_time":this.UpdateTime,"data_state":2}).Error
}

//批量删除用户关系
func (this *UserPetPo) DeleteUserPetsDB(dbc *gorm.DB,pids []int32) int64 {
	log.Info("DeleteUserPetsDB:", pids,this)
	return dbc.Table("user_pet").Where("uid = ? AND pid in (?) AND data_state = 1",this.Uid,pids).
		Updates(map[string]interface{}{"update_time":this.UpdateTime,"data_state":2}).RowsAffected
}

//查询用户的宠物
func (this *UserPetPo) GetUserPetsDB() (userPetPos []*UserPetPo,err error) {
	log.Info("GetUserPetsDB:", this)
	return userPetPos,core.MysqlClient.Table("user_pet").Where("uid = ? AND data_state = 1",this.Uid).Find(&userPetPos).Error
}

//批量查询用户宠物关系
func (this *UserPetPo) GetUsersPetsDB(pids []int32) (userPetPos []*UserPetPo,err error) {
	log.Info("GetUserPetsDB:", this)
	return userPetPos,core.MysqlClient.Table("user_pet").
		Where("uid = ? AND permit = ? AND pid in (?) AND data_state = 1",this.Uid,this.Permit,pids).
		Find(&userPetPos).Error
}

//查询用户宠物
func (this *UserPetPo) GetUserPetDB() (err error) {
	log.Info("GetUserPetDB:", this)
	return core.MysqlClient.Table("user_pet").Where("uid = ? AND pid = ? AND data_state = 1",
		this.Uid,this.Pid).First(&this).Error
}

//查询宠物的用户
func (this *UserPetPo) GetUsersPetDB() (userPetPos []*UserPetPo, err error) {
	log.Info("GetUserPetDB:", this)
	return userPetPos,core.MysqlClient.Table("user_pet").Where("pid = ? AND data_state = 1",
		this.Pid).Find(&userPetPos).Error
}

//批量查询宠物用户关系
func (this *UserPetPo) GetUserPetsMemberDB(pids []int32) (userPetPo []*UserPetPo, err error) {
	log.Info("GetUserPetsMemberDB:", this)
	err = core.MysqlClient.Table("user_pet").Where("pid in (?) AND data_state = 1", pids).Find(&userPetPo).Error
	return userPetPo,err
}


/**
设备宠物关系
 */
type DevicePetPo struct {
	Id           int32		`gorm:"column:id;primary_key;unique"`
	Did          int32		`gorm:"column:did"`           	//设备id
	Pid          int32     	`gorm:"column:pid"`				//宠物id
	CreationTime time.Time	`gorm:"column:creation_time"` 	//创建时间
	UpdateTime   time.Time 	`gorm:"column:update_time"`   	//修改时间
	DataState    int32		`gorm:"column:data_state"`    	//数据状态
}

//关联设备与宠物的关系
func (this *DevicePetPo) SetDevicePetDB() int64 {
	log.Info("SetDevicePetDB:", this)
	return core.MysqlClient.Table("device_pet").Create(&this).RowsAffected
}

//取消设备与宠物的关系
func (this *DevicePetPo) DeleteDevicePetDB() int64 {
	log.Info("DeleteDevicePetDB:", this)
	return core.MysqlClient.Table("device_pet").Where("id = ? AND did = ? AND pid = ? AND data_state = 1",
		this.Id, this.Did, this.Pid).Updates(map[string]interface{}{"data_state": 2, "update_time": this.UpdateTime}).RowsAffected
}

//批量设备查询关系
func (this *DevicePetPo) GetDevicesPetsDidsDB(dids []int32) (devicePetPos []*DevicePetPo, err error) {
	log.Info("GetDevicesPetsDidsDB:", this,dids)
	return devicePetPos,core.MysqlClient.Table("device_pet").
		Where("did in (?) AND data_state = 1", dids).Order("id desc").Find(&devicePetPos).Error
}

//批量宠物查询关系
func (this *DevicePetPo) GetDevicesPetsPidsDB(pids []int32) (devicePetPos []*DevicePetPo, err error) {
	log.Info("GetDevicesPetsPidsDB:", this)
	return devicePetPos,core.MysqlClient.Table("device_pet").
		Where("pid in (?) AND data_state = 1", pids).Order("id desc").Find(&devicePetPos).Error
}

//获取宠物、设备的关系
func (this *DevicePetPo) GetDevicePetDB() error {
	log.Info("GetDevicePetDB:", this)
	return core.MysqlClient.Table("device_pet").Where("did = ? AND pid = ? AND data_state = 1",
		this.Did, this.Pid).Order("id DESC").Limit(1).First(&this).Error
}

//获取设备的关系
func (this *DevicePetPo) GetDevicePetsDB() error {
	log.Info("GetDevicePetsDB:", this)
	return core.MysqlClient.Table("device_pet").Where("did = ? AND data_state = 1",
		this.Did).Order("id DESC").First(&this).Error
}

//获取宠物的关系
func (this *DevicePetPo) GetDevicesPetDB() error {
	log.Info("GetDevicePetDB:", this)
	return core.MysqlClient.Table("device_pet").Where("pid = ? AND data_state = 1",
		this.Pid).Order("id DESC").Limit(1).First(&this).Error
}

