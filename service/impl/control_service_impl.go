package impl

import (
	"github.com/DnullP/se_work/config"
	"github.com/DnullP/se_work/mock"
	dto "github.com/DnullP/se_work/model/DTO"
)

type ControlServiceImpl struct{}

var workingQueue []uint
var waitingQueue []uint

func (controlService ControlServiceImpl) HandleAdminCommand(command dto.Command) {
	switch command.Command {

	case dto.SetMode:
		mode, ok := command.Args[0].(string)
		if !ok {
			panic("invalid type")
		}
		switch mode {
		case "cool":
			config.WorkMode = config.COLD
		case "warm":
			config.WorkMode = config.WARM
		}

	case dto.SetPrice:
		priceRate, ok1 := command.Args[0].(float32)
		if !ok1 {
			panic("invalid type")
		}
		config.PriceRate = priceRate

	case dto.SetValidRange:
		validLowerTemperature, ok1 := command.Args[0].(float32)
		validUpperTemperature, ok2 := command.Args[1].(float32)
		if !ok1 || !ok2 {
			panic("invalid type")
		}
		config.ValidLowerTemperature = validLowerTemperature
		config.ValidUpperTemperature = validUpperTemperature

	case dto.TurnOffAll:
		for i := 0; i < len(workingQueue); i++ {
			device := mock.GetAirConditionByDeviceID(workingQueue[i])
			device.Working = false
			mock.SetAirCondition(device)
		}
		waitingQueue = append(waitingQueue, workingQueue...)
		workingQueue = []uint{}

	case dto.TurnOnAll:
		for i := 0; i < len(waitingQueue); i++ {
			device := mock.GetAirConditionByDeviceID(waitingQueue[i])
			device.Working = true
			mock.SetAirCondition(device)
		}
		workingQueue = append(workingQueue, waitingQueue...)
		waitingQueue = []uint{}
	}
}

func (controlService ControlServiceImpl) HandleUserCommand(command dto.Command) {
	device := mock.GetAirConditionByDeviceID(command.DeviceID)

	switch command.Command {
	case dto.TurnOn:
		device.Working = true
		mock.SetAirCondition(device)
		workingQueue = append(workingQueue, uint(command.DeviceID))
	case dto.TurnOff:
		device.Working = false
		mock.SetAirCondition(device)
		for i := 0; i < len(workingQueue); i++ {
			if workingQueue[i] == uint(command.DeviceID) {
				workingQueue = append(workingQueue[:i], workingQueue[i+1:]...)
			}
		}
		for i := 0; i < len(waitingQueue); i++ {
			if waitingQueue[i] == uint(command.DeviceID) {
				waitingQueue = append(waitingQueue[:i], waitingQueue[i+1:]...)
			}
		}
	case dto.SetTemperature:
		targetTemperature, ok := command.Args[0].(float32)
		if !ok {
			panic("invalid type")
		}
		device.TargetTemperature = float32(targetTemperature)
		mock.SetAirCondition(device)
	case dto.SetSpeed:
		speed, ok := command.Args[0].(string)
		if !ok {
			panic("invalid type")
		}
		device.Speed = speed
		mock.SetAirCondition(device)
	}
}
