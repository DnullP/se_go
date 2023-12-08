package db

import (
	"github.com/DnullP/se_work/utils"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	return utils.DB
}
