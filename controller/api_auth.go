package controller

import (
	"net/http"

	restAPI "github.com/DnullP/se_work/model/restAPI/response"
	getService "github.com/DnullP/se_work/utils/service"
	"github.com/gin-gonic/gin"
)

// LoginGet - login
// LoginGet handles the GET request for the login endpoint.
// It calls the Login function of the UserService to authenticate the user.
// If the login is successful, it returns a JSON response with status "login success".
// Otherwise, it returns a JSON response with status "fail".
func LoginGet(c *gin.Context) {
	permission, _ := getService.GetUserService().Login("", "")
	response := restAPI.Status{}

	if permission {
		response.Status = "login success"
	} else {
		response.Status = "fail"
	}

	c.JSON(http.StatusOK, response)
}

// LogoutGet - logout
// LogoutGet handles the GET request for logging out the user.
// It calls the Logout method of the UserService to perform the logout operation.
// The response is sent back as a JSON object with the status of the logout operation.
func LogoutGet(c *gin.Context) {
	permission, _ := getService.GetUserService().Logout("")
	response := restAPI.Status{}

	if permission {
		response.Status = "login success"
	} else {
		response.Status = "fail"
	}

	c.JSON(http.StatusOK, response)
}
