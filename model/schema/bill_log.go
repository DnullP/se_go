package schema

import (
	"gorm.io/gorm"
)

type BillLog struct {
	CommandLogID       uint    `gorm:"primaryKey"`
	TargetTemperature  float32 `gorm:"not null"`
	CurrentTemperature float32 `gorm:"not null"`
	Speed              string  `gorm:"not null"`
}

func (billLog *BillLog) TableName() string {
	return "bill_log"
}

func (billLog *BillLog) CreateBillLog(db *gorm.DB) *BillLog {
	db.Create(&billLog)
	return billLog
}
