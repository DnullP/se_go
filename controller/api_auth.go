package controller

import (
	"net/http"

	restAPI "github.com/DnullP/se_work/model/restAPI/response"
	"github.com/DnullP/se_work/utils"
	"github.com/gin-gonic/gin"
)

// LoginPost - login
func LoginPost(c *gin.Context) {
	permission, _ := utils.GetUserService().Login("", "")
	response := restAPI.Status{}

	if permission {
		response.Status = "success"
	} else {
		response.Status = "fail"
	}

	c.JSON(http.StatusOK, response)
}

// LogoutPost - logout
func LogoutPost(c *gin.Context) {
	permission, _ := utils.GetUserService().Logout("")
	response := restAPI.Status{}

	if permission {
		response.Status = "success"
	} else {
		response.Status = "fail"
	}

	c.JSON(http.StatusOK, response)
}
