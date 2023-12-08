package controller

import (
	"net/http"

	restAPI "github.com/DnullP/se_work/model/restAPI/response"
	getService "github.com/DnullP/se_work/utils/service"
	"github.com/gin-gonic/gin"
)

// LoginGet - login
func LoginGet(c *gin.Context) {
	permission, _ := getService.GetUserService().Login("", "")
	response := restAPI.Status{}

	if permission {
		response.Status = "success"
	} else {
		response.Status = "fail"
	}

	c.JSON(http.StatusOK, response)
}

// LogoutGet - logout
func LogoutGet(c *gin.Context) {
	permission, _ := getService.GetUserService().Logout("")
	response := restAPI.Status{}

	if permission {
		response.Status = "success"
	} else {
		response.Status = "fail"
	}

	c.JSON(http.StatusOK, response)
}
