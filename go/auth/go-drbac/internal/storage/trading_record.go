package storage

import (
	"github.com/jinzhu/gorm"
	"time"
	. "auth/go-drbac/common"
)

type TradingRecord struct {
	TradingId        	int32     	`gorm:"column:trading_id;primary_key;unique"`
	Tid    				int32   	`gorm:"column:tid"`
	CreateTime 			time.Time 	`gorm:"column:create_time"`
	TradingContent    	string    	`gorm:"column:trading_content"`
	TradingUnitPrice    float32   	`gorm:"column:trading_unit_price"`
	TradingCount    	int32   	`gorm:"column:trading_count"`
	TradingState    	int32 		`gorm:"column:trading_state"`
	TradingTotalPrice 	float32   	`gorm:"column:trading_total_price"`
}


//分页查询交易记录
func (t *TradingRecord) GetTradingRecords(tx *gorm.DB, page, count int32) (tradingRecords []*TradingRecord, totalCount int32, err error) {
	if err = tx.Table("tbl_trading_record").Where("tid = ?",t.Tid).Count(&totalCount).Error; err != nil {
		return
	}
	if count == 0 || count == -1{
		count = totalCount
	}
	err = tx.Table("tbl_trading_record").Where("tid = ?",t.Tid).Order("trading_id asc").Limit(count).Offset((page - 1) * count).Find(&tradingRecords).Error
	if err == gorm.ErrRecordNotFound {
		err = ErrDoesNotExist
	}
	return
}

//插入交易记录
func (t *TradingRecord) AddTradingRecord (tx *gorm.DB)(err error){
	err = tx.Table("tbl_trading_record").Create(t).Error
	return
}
