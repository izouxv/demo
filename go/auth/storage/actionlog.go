package storage

import (
	"time"
	log "github.com/cihub/seelog"
	"auth/config"
)

type ActionLog struct {
	Id        		int32      		`gorm:"column:id;primary_key;unique"`
	ActionUsername  string			`gorm:"column:action_username"`
	ActionTime      int64			`gorm:"column:action_time"`
	ActionType		int32			`gorm:"column:action_type"`
	ActionName 		string			`gorm:"column:action_name"`
	ActionObject	string			`gorm:"column:action_object"`
	CreateTime		time.Time		`gorm:"column:create_time"`
	Tid				int32			`gorm:"column:tid"`
	Did				int32			`gorm:"column:did"`
}

//添加操作日志
func (a *ActionLog) AddActionLog() error {
	if err := config.C.MySQL.DB.Table("tbl_action_log").Create(a).Error; err != nil {
		log.Error("New ActionLog to db error", err)
		return err
	}
	return nil
}

//批量获取操作日志
func (a *ActionLog) GetAllActionLogs(page, count int32) ([]ActionLog, int32, error) {
	var totalCount int32
	if err := config.C.MySQL.DB.Table("tbl_action_log").Count(&totalCount).Error; err != nil {
		log.Error("GetAllAddActions err,",err)
		return nil, 0, err
	}
	if count == 0 {
		count = totalCount
	}
	als := make([]ActionLog, 0, totalCount)
	if err := config.C.MySQL.DB.Table("tbl_action_log").Order("id desc").Limit(count).Offset((page - 1) * count).Find(&als).Error; err != nil {
		log.Error("Get ActionLogs from db error", err)
		return nil, 0, err
	}
	return als, totalCount, nil
}

//批量获取指定租户操作日志
func (a *ActionLog) GetActionLogsByTid(page, count int32) ([]ActionLog, int32, error) {
	var totalCount int32
	if err := config.C.MySQL.DB.Table("tbl_action_log").Where("tid = ?",a.Tid).Count(&totalCount).Error; err != nil {
		log.Error("GetAllAddActions err,",err)
		return nil, 0, err
	}
	if count == 0 {
		count = totalCount
	}
	als := make([]ActionLog, 0, totalCount)
	if err := config.C.MySQL.DB.Table("tbl_action_log").Where("tid = ?",a.Tid).Order("id desc").Limit(count).Offset((page - 1) * count).Find(&als).Error; err != nil {
		log.Error("Get ActionLogs from db error", err)
		return nil, 0, err
	}
	return als, totalCount, nil
}

//批量获取指定域操作日志
func (a *ActionLog) GetActionLogsByDid(page, count int32) ([]ActionLog, int32, error) {
	var totalCount int32
	if err := config.C.MySQL.DB.Table("tbl_action_log").Where("did = ?",a.Did).Count(&totalCount).Error; err != nil {
		log.Error("GetAllAddActions err,",err)
		return nil, 0, err
	}
	if count == 0 {
		count = totalCount
	}
	als := make([]ActionLog, 0, totalCount)
	if err := config.C.MySQL.DB.Table("tbl_action_log").Where("did = ?",a.Did).Order("id desc").Limit(count).Offset((page - 1) * count).Find(&als).Error; err != nil {
		log.Error("Get ActionLogs from db error", err)
		return nil, 0, err
	}
	return als, totalCount, nil
}


//根据用户名批量获取操作日志
func (a *ActionLog) GetActionLogsByUsername(page, count int32) ([]ActionLog, int32, error) {
	var totalCount int32
	if err := config.C.MySQL.DB.Table("tbl_action_log").Where("action_username = ?",a.ActionUsername).Count(&totalCount).Error; err != nil {
		log.Error("GetActionLogsByUsername err,",err)
		return nil, 0, err
	}
	if count == 0 {
		count = totalCount
	}
	als := make([]ActionLog, 0, totalCount)
	if err := config.C.MySQL.DB.Table("tbl_action_log").Where("action_username = ?",a.ActionUsername).Order("id desc").Limit(count).Offset((page - 1) * count).Find(&als).Error; err != nil {
		log.Error("GetActionLogsByUsername from db error", err)
		return nil, 0, err
	}
	return als, totalCount, nil
}

//根据角色批量获取操作日志
//func (a *ActionLog) GetActionLogsByRole(page, count int32) ([]ActionLog, int32, error) {
//	var totalCount int32
//	if err := module.MysqlClient().Table("tbl_action_log").Where("action_role = ?",a.ActionRole).Count(&totalCount).Error; err != nil {
//		log.Error("GetActionLogsByRole err,",err)
//		return nil, 0, err
//	}
//	if count == 0 {
//		count = totalCount
//	}
//	als := make([]ActionLog, 0, totalCount)
//	if err := module.MysqlClient().Table("tbl_action_log").Where("action_role = ?",a.ActionRole).Order("id desc").Limit(count).Offset((page - 1) * count).Find(&als).Error; err != nil {
//		log.Error("GetActionLogsByRole from db error", err)
//		return nil, 0, err
//	}
//	return als, totalCount, nil
//}

//根据类型批量获取操作日志
func (a *ActionLog) GetActionLogsByType(page, count int32) ([]ActionLog, int32, error) {
	var totalCount int32
	if err := config.C.MySQL.DB.Table("tbl_action_log").Where("action_type = ?",a.ActionType).Count(&totalCount).Error; err != nil {
		log.Error("GetActionLogsByRole err,",err)
		return nil, 0, err
	}
	if count == 0 {
		count = totalCount
	}
	als := make([]ActionLog, 0, totalCount)
	if err := config.C.MySQL.DB.Table("tbl_action_log").Where("action_type = ?",a.ActionType).Order("id desc").Limit(count).Offset((page - 1) * count).Find(&als).Error; err != nil {
		log.Error("GetActionLogsByRole from db error", err)
		return nil, 0, err
	}
	return als, totalCount, nil
}

//根据用户名批量获取操作日志
func (a *ActionLog) GetActionLogsByUsernameAndType(page, count int32) ([]ActionLog, int32, error) {
	var totalCount int32
	if err := config.C.MySQL.DB.Table("tbl_action_log").Where("action_username = ? And action_type = ?",a.ActionUsername,a.ActionType).Count(&totalCount).Error; err != nil {
		log.Error("GetActionLogsByUsernameAndType err,",err)
		return nil, 0, err
	}
	if count == 0 {
		count = totalCount
	}
	als := make([]ActionLog, 0, totalCount)
	if err := config.C.MySQL.DB.Table("tbl_action_log").Where("action_username = ? And action_type = ?",a.ActionUsername,a.ActionType).Order("id desc").Limit(count).Offset((page - 1) * count).Find(&als).Error; err != nil {
		log.Error("GetActionLogsByUsernameAndType from db error", err)
		return nil, 0, err
	}
	return als, totalCount, nil
}

//根据条件批量获取操作日志
func (a *ActionLog) GetActionLogs(page, count int32) ([]ActionLog, int32, error) {
	var totalCount int32
	als := make([]ActionLog, 0, totalCount)
	if a.ActionUsername != "" {
		actionUsername := "%" + a.ActionUsername + "%"
		if err := config.C.MySQL.DB.Table("tbl_action_log").Where(&ActionLog{ActionType:a.ActionType, Tid:a.Tid, Did:a.Did},).Where("action_username like ?",actionUsername).Count(&totalCount).Error; err != nil {
			log.Error("GetActionLogsByUsernameAndType err,",err)
			return nil, 0, err
		}
		if count == 0 {
			count = totalCount
		}
		if err := config.C.MySQL.DB.Table("tbl_action_log").Where(&ActionLog{ActionType:a.ActionType, Tid:a.Tid, Did:a.Did},).Where("action_username like ?",actionUsername).Order("id desc").Limit(count).Offset((page - 1) * count).Find(&als).Error; err != nil {
			log.Error("GetActionLogsByUsernameAndType from db error", err)
			return nil, 0, err
		}
	}else {
		if err := config.C.MySQL.DB.Table("tbl_action_log").Where(&ActionLog{ActionType:a.ActionType, Tid:a.Tid, Did:a.Did}).Count(&totalCount).Error; err != nil {
			log.Error("GetActionLogsByUsernameAndType err,",err)
			return nil, 0, err
		}
		if count == 0 {
			count = totalCount
		}
		if err := config.C.MySQL.DB.Table("tbl_action_log").Where(&ActionLog{ActionType:a.ActionType, Tid:a.Tid, Did:a.Did}).Order("id desc").Limit(count).Offset((page - 1) * count).Find(&als).Error; err != nil {
			log.Error("GetActionLogsByUsernameAndType from db error", err)
			return nil, 0, err
		}
	}
	return als, totalCount, nil
}