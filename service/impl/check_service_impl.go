package impl

import (
	"time"

	dto "github.com/DnullP/se_work/model/DTO"
	"github.com/DnullP/se_work/model/schema"
	"github.com/DnullP/se_work/utils/db"
)

type CheckServiceImpl struct{}

// CheckIn checks in a guest by creating a new guest record in the database.
// It takes a CheckInfo object as input, which contains the guest's name and room ID.
// The current date and time are recorded as the check-in time.
// The check-out time is set to "still_in" indicating that the guest is still checked in.
// The guest record is then created in the database.
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

// CheckOut checks out a guest by updating the check-out time in the database.
// It takes a CheckInfo object as input, which contains the guest name.
// The current date and time are obtained using time.Now() and formatted as "2006-01-02 15:04:05".
// The guest's check-out time is set in the Guest struct and saved to the database using the SetCheckOutTime method.
func (checkService CheckServiceImpl) CheckOut(info dto.CheckInfo) {
	db := db.GetDB()

	var checkOutTime string = time.Now().Format("2006-01-02 15:04:05")

	guest := schema.Guest{
		GuestName:    info.GuestName,
		CheckOutTime: checkOutTime,
	}

	guest.SetCheckOutTime(db)

}
