package db

import (
	"time"
	"petfone-rpc/core"
	"github.com/jinzhu/gorm"
)

type ActionLog struct {
	Id			int32    	`gorm:"column:id;primary_key;unique"`
	Uid			int32   	`gorm:"column:uid"`
	Path		string     	`gorm:"column:path"`
	Method		string     	`gorm:"column:method"`
	Duration	string		`gorm:"column:duration"`
	Code		string		`gorm:"column:code"`
	Token		string		`gorm:"column:token"`
	Ip			int64		`gorm:"column:ip"`
	Addr		string		`gorm:"column:addr"`
	DevInfo		string		`gorm:"column:dev_info"`
	CreateTime	time.Time 	`gorm:"column:create_time"`
	DataState	int32 		`gorm:"column:data_state"`
}

//记录操作信息
func (this *ActionLog) SetActionLog(dbc *gorm.DB) error {
	return dbc.Table("action_log").Create(this).Error
}

//批量查询用户操作信息
func (this *ActionLog) GetActionInfos(actions int32, uids []int32) ([]*ActionLog, error) {
	var logs []*ActionLog
	db := core.MysqlClient
	rows, err := db.Raw(
		"SELECT * FROM action_log AS ac " +
			"WHERE " +
			"(SELECT COUNT(*) FROM action_log WHERE actions = ? AND uid = ac.uid AND id > ac.id AND ac.uid IN (?)) < 1 ",
		actions, uids).Rows()
	defer rows.Close()
	if err != nil {
		return logs, err
	}
	for rows.Next() {
		actionLog := &ActionLog{}
		db.ScanRows(rows, actionLog)
		logs = append(logs, actionLog)
	}
	return logs, nil
}