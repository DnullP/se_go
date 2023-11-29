package schema

import (
	"gorm.io/gorm"
)

type Guest struct {
	GuestLogID   uint   `gorm:"primaryKey"`
	GuestName    string `gorm:"not null"`
	DeviceID     uint   `gorm:"not null"`
	CheckInTime  string
	CheckOutTime string
}

func (guest *Guest) TableName() string {
	return "guest"
}

func (guest *Guest) CreateGuest(db *gorm.DB) *Guest {
	db.Create(&guest)
	return guest
}

func SetCheckInTime(db *gorm.DB, guest *Guest) *Guest {
	db.Model(&guest).Update("check_in_time", guest.CheckInTime)
	return guest
}

func SetCheckOutTime(db *gorm.DB, guest *Guest) *Guest {
	db.Model(&guest).Update("check_out_time", guest.CheckOutTime)
	return guest
}
