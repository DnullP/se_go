

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetDailyReportPost - get_daily_report
func GetDailyReportPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetWeeklyReportPost - get_weekly_report
func GetWeeklyReportPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
