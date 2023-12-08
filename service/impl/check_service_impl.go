package impl

import (
	"time"

	dto "github.com/DnullP/se_work/model/DTO"
	"github.com/DnullP/se_work/model/schema"
	"github.com/DnullP/se_work/utils/db"
)

type CheckServiceImpl struct{}

func (checkService CheckServiceImpl) CheckIn(info dto.CheckInfo) {
	db := db.GetDB()

	var checkInTime string = time.Now().Format("2006-01-02 15:04:05")

	guest := schema.Guest{
		GuestName:    info.GuestName,
		DeviceID:     info.RoomID,
		CheckInTime:  checkInTime,
		CheckOutTime: "still_in",
	}

	guest.CreateGuest(db)
}

func (checkService CheckServiceImpl) CheckOut(info dto.CheckInfo) {
	db := db.GetDB()

	var checkOutTime string = time.Now().Format("2006-01-02 15:04:05")

	guest := schema.Guest{
		GuestName:    info.GuestName,
		CheckOutTime: checkOutTime,
	}

	guest.SetCheckOutTime(db)

}
