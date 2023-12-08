package controller

import (
	"net/http"

	"github.com/DnullP/se_work/model/restAPI/response"
	"github.com/gin-gonic/gin"
)

// GetDailyReportPost - get_daily_report
func GetDailyReportPost(c *gin.Context) {

	response := response.Report{
		RoomList: []response.RoomList{
			{
				BillTimes:       0,
				DailyCost:       0,
				RoomID:          0,
				SchedTimes:      0,
				TurnTimes:       0,
				TweatSpeedTimes: 0,
				TweatTemTimes:   0,
				WorkingTime:     0,
			},
		},
	}

	c.JSON(http.StatusOK, response)
}

// GetWeeklyReportPost - get_weekly_report
func GetWeeklyReportPost(c *gin.Context) {

	response := response.Report{
		RoomList: []response.RoomList{
			{
				BillTimes:       0,
				DailyCost:       0,
				RoomID:          0,
				SchedTimes:      0,
				TurnTimes:       0,
				TweatSpeedTimes: 0,
				TweatTemTimes:   0,
				WorkingTime:     0,
			},
		},
	}

	c.JSON(http.StatusOK, response)
}
