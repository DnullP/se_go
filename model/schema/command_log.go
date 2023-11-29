package schema

import (
	"gorm.io/gorm"
)

type CommandLog struct {
	CommandLogID uint   `gorm:"primaryKey"`
	DeviceID     uint   `gorm:"not null"`
	Time         int    `grom:"autoCreateTime"`
	Command      string `gorm:"not null"`
	Args         string
}

func (commandLog *CommandLog) TableName() string {
	return "command_log"
}

func (commandLog *CommandLog) CreateCommandLog(db *gorm.DB) *CommandLog {
	db.Create(&commandLog)
	return commandLog
}
