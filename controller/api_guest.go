package controller

import (
	"strconv"

	dto "github.com/DnullP/se_work/model/DTO"
	"github.com/DnullP/se_work/model/restAPI/receive"
	"github.com/DnullP/se_work/model/restAPI/response"
	getService "github.com/DnullP/se_work/utils/service"
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

	args := make([]interface{}, 2)

	switch command.Command {
	case receive.SetSpeed:
		args[0] = *command.Args.Speed
	case receive.SetTemperature:
		args[0] = float32(*command.Args.TargetTemp)
		args[1] = float32(*command.Args.CurrentRoomTemp)
	}

	requireHandling := dto.Command{
		Command:  command.Command,
		DeviceID: uint(deviceID),
		Args:     args,
	}

	getService.GetControlService().HandleUserCommand(requireHandling)

	response_client := response.Status{}
	response_client.Status = "success!"
	c.JSON(200, response_client)
}

// GetPanelStatus - get_panel_status
func GetPanelStatusGet(c *gin.Context) {
	device_id := c.Query("device_id")
	deviceID, _ := strconv.ParseUint(device_id, 10, 64)

	device := getService.GetQueryService().GetDeviceStatus(uint(deviceID))

	response_client := response.PanelStatus{
		EnvTemperature:    float64(device.CurrentTemperature),
		Working:           device.Working,
		TargetTemperature: float64(device.TargetTemperature),
		Speed:             &device.Speed,
	}

	c.JSON(200, response_client)
}
