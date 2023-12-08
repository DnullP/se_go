package schema

import (
	"gorm.io/gorm"
)

type Guest struct {
	GuestLogID   uint   `gorm:"primaryKey;autoIncrement"`
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

func (guest *Guest) SetCheckInTime(db *gorm.DB) *Guest {
	db.Model(&guest).Update("check_in_time", guest.CheckInTime)
	return guest
}

func (guest *Guest) SetCheckOutTime(db *gorm.DB) *Guest {
	db.Model(&guest).Where("guest_name = ? and check_out_time = ?", guest.GuestName, "still_in").Update("check_out_time", guest.CheckOutTime)
	return guest
}
