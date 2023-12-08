/*
 * air_condition
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package controller

import (
	"net/http"
	"strconv"

	"github.com/DnullP/se_work/config"
	dto "github.com/DnullP/se_work/model/DTO"
	"github.com/DnullP/se_work/model/restAPI/receive"
	"github.com/DnullP/se_work/model/restAPI/response"
	"github.com/DnullP/se_work/utils"
	"github.com/gin-gonic/gin"
)

// AdminControlPost - admin_control
func AdminControlPost(c *gin.Context) {
	device_id := c.Query("device_id")
	deviceID, _ := strconv.ParseUint(device_id, 10, 64)

	command := receive.AdminCommandReceive{}
	c.BindJSON(&command)

	var args []interface{} = make([]interface{}, 2)

	switch command.Command {
	case receive.SetMode:
		args[0] = *command.Args.Mode
	case receive.SetPrice:
		args[0] = float32(*command.Args.FeeRate)
	case receive.SetValidRange:
		args[0] = float32(*command.Args.ValidRangeLow)
		args[1] = float32(*command.Args.ValidRangeHigh)
	}

	requiredHandling := dto.Command{
		Command:  command.Command,
		DeviceID: uint(deviceID),
		Args:     args,
	}
	utils.GetControlService().HandleAdminCommand(requiredHandling)

	response_client := response.Status{}
	response_client.Status = "success!"
	c.JSON(http.StatusOK, response_client)
}

// GetAllDeviceStatusGet - get_all_device_status
func GetAllDeviceStatusGet(c *gin.Context) {
	response_client := response.DeviceStatusList{}
	for i := 0; i < config.MockNum; i++ {
		device := utils.GetQueryService().GetDeviceStatus(uint(i))
		response_client.DeviceList = append(response_client.DeviceList, response.DeviceStatus_room_id{
			DeviceStatus: response.DeviceStatus{
				EnvTemperature:    float64(device.CurrentTemperature),
				Mode:              config.WorkMode,
				TargetTemperature: float64(device.TargetTemperature),
				Speed:             &device.Speed,
				TotalCost:         float64(device.TotalCost),
				Working:           device.Working,
			},
			RoomID: uint(i),
		})
	}
	c.JSON(http.StatusOK, response_client)
}

// GetDeviceStatusGet - get_device_status
func GetDeviceStatusGet(c *gin.Context) {
	device_id := c.Query("device_id")
	deviceID, _ := strconv.ParseUint(device_id, 10, 64)

	device := utils.GetQueryService().GetDeviceStatus(uint(deviceID))

	response_client := response.DeviceStatus{
		EnvTemperature:    float64(device.CurrentTemperature),
		Mode:              config.WorkMode,
		TargetTemperature: float64(device.TargetTemperature),
		Speed:             &device.Speed,
		TotalCost:         float64(device.TotalCost),
		Working:           device.Working,
	}
	c.JSON(http.StatusOK, response_client)
}
