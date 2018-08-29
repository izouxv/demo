package db

import (
	"time"
	"petfone-rpc/core"
	log "github.com/cihub/seelog"
	"github.com/jinzhu/gorm"
)

type MotionDataPetPo struct {
	Id				int64		`gorm:"column:id;primary_key;unique"`
	Uid    			int32		`gorm:"column:uid"`
	Pid    			int32		`gorm:"column:pid"`
	Pdid    		int32		`gorm:"column:pdid"`
	ReportTime    	time.Time	`gorm:"column:report_time"`		//上报时间
	StartTime		time.Time 	`gorm:"column:start_time"`		//起始时间
	EndTime			time.Time 	`gorm:"column:end_time"`		//结束时间
	CardioTimes		int64		`gorm:"column:cardio_times"`	//普通时长
	StrenuousTimes	int64		`gorm:"column:strenuous_times"`	//剧烈运动时长
	Calorie       	float32		`gorm:"column:calorie"`			//消耗卡
	Steps			int32		`gorm:"column:steps"`			//运动步数
	Url				string		`gorm:"column:url"`				//图片地址
	Width 			int32		`gorm:"column:width"`			//原宽度
	Height 			int32		`gorm:"column:height"`			//原高度
	Size			int32		`gorm:"column:size"`
	Name			string		`gorm:"column:name"`
	CreationTime 	time.Time	`gorm:"column:creation_time"` 	//创建时间
	DataState    	int32		`gorm:"column:data_state"`    	//数据状态
}

//查询md5
func (m *MotionDataPetPo) GetMotionDataPetExist(tableSuffix string) error {
	log.Info("GetMotionDataPetExist this:", m)
	return core.MysqlClient.Table("motion_data_"+tableSuffix).Where("(pid = ? and start_time = ? and end_time = ? ) or url = ?",m.Pid,m.StartTime,m.EndTime,m.Url).Select("id").First(&m).Error
}

//存数据
func (m *MotionDataPetPo) SetMotionDataPet(tableSuffix string) error {
	log.Info("SetMotionDataPet this:", m)
	return core.MysqlClient.Table("motion_data_"+tableSuffix).Create(m).Error
}

//查询时间内数据
func (m *MotionDataPetPo) GetMotionDataPetByTime(tableSuffix string, start,end int64) (motions []*MotionDataPetPo, err error) {
	log.Info("GetMotionDataPetByTime pid:", m.Pid,start,end)
	err = core.MysqlClient.Table("motion_data_"+tableSuffix).
		Where("pid = ? AND report_time >= FROM_UNIXTIME(?) AND report_time < FROM_UNIXTIME(?)", m.Pid, start, end).
			Find(&motions).Error
	return motions, err
}

type TestPicPo struct {
	Id 		int64 	`gorm:"column:id;primary_key;unique"`
	Pic 	[]byte	`gorm:"column:pic"`
	State 	int32 	`gorm:"column:state"`
}


//初始化
func (t *TestPicPo) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("id", 0)
	return nil
}

func (t *TestPicPo) SetPic() error {
	return core.MysqlClient.Table("test").Create(t).Error
}

func (t *TestPicPo) GetPic() error {
	log.Info("GetPic:",t.Id)
	return core.MysqlClient.Table("test").Where("id = ?", t.Id).First(&t).Error
}