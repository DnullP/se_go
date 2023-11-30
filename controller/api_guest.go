package controller

import (
	"strconv"

	dto "github.com/DnullP/se_work/model/DTO"
	"github.com/DnullP/se_work/model/restAPI/receive"
	"github.com/DnullP/se_work/model/restAPI/response"
	"github.com/DnullP/se_work/utils"
	"github.com/gin-gonic/gin"
)

// RemoteControlPost - remote_control
func RemoteControlPost(c *gin.Context) {
	command := receive.UserCommandReceive{}
	err := c.BindJSON(&command)
	if err != nil {
		panic(err)
	}

	deviceID, _ := strconv.ParseUint(c.Query("device_id"), 10, 64)

	requireHandling := dto.Command{
		Command:  command.Command,
		DeviceID: uint(deviceID),
		Args:     command.Args,
	}

	utils.GetControlService().HandleUserCommand(requireHandling)

	response_client := response.Status{}
	response_client.Status = "success"
	c.JSON(200, response_client)
}
