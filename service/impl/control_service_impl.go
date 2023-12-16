package impl

import (
	"github.com/DnullP/se_work/config"
	"github.com/DnullP/se_work/mock"
	dto "github.com/DnullP/se_work/model/DTO"
)

type ControlServiceImpl struct{}

var workingQueue []uint
var waitingQueue []uint

// HandleAdminCommand handles the admin command and performs the corresponding actions based on the command type.
// It takes a command of type dto.Command as input.
// HandleAdminCommand handles the admin command and performs the corresponding actions based on the command type.
// It takes a command of type dto.Command as input.
func (controlService ControlServiceImpl) HandleAdminCommand(command dto.Command) {
	switch command.Command {

	case dto.SetMode:
		// SetMode case handles the command to set the work mode.
		// It expects the mode as the first argument in the command's Args field.
		// Valid modes are "cool" and "warm".
		// If the mode is "cool", the config.WorkMode is set to config.COLD.
		// If the mode is "warm", the config.WorkMode is set to config.WARM.
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
		// SetPrice case handles the command to set the price rate.
		// It expects the price rate as the first argument in the command's Args field.
		// The price rate should be of type float32.
		priceRate, ok1 := command.Args[0].(float32)
		if !ok1 {
			panic("invalid type")
		}
		config.PriceRate = priceRate

	case dto.SetValidRange:
		// SetValidRange case handles the command to set the valid temperature range.
		// It expects the valid lower temperature as the first argument and the valid upper temperature as the second argument in the command's Args field.
		// The valid temperatures should be of type float32.
		validLowerTemperature, ok1 := command.Args[0].(float32)
		validUpperTemperature, ok2 := command.Args[1].(float32)
		if !ok1 || !ok2 {
			panic("invalid type")
		}
		config.ValidLowerTemperature = validLowerTemperature
		config.ValidUpperTemperature = validUpperTemperature

	case dto.TurnOffAll:
		// TurnOffAll case handles the command to turn off all air conditioning devices.
		// It sets the Working field of each device in the workingQueue to false.
		// The devices are then moved from the workingQueue to the waitingQueue.
		for i := 0; i < len(workingQueue); i++ {
			device := mock.GetAirConditionByDeviceID(workingQueue[i])
			device.Working = false
			mock.SetAirCondition(device)
		}
		waitingQueue = append(waitingQueue, workingQueue...)
		workingQueue = []uint{}

	case dto.TurnOnAll:
		// TurnOnAll case handles the command to turn on all air conditioning devices.
		// It sets the Working field of each device in the waitingQueue to true.
		// The devices are then moved from the waitingQueue to the workingQueue.
		for i := 0; i < len(waitingQueue); i++ {
			device := mock.GetAirConditionByDeviceID(waitingQueue[i])
			device.Working = true
			mock.SetAirCondition(device)
		}
		workingQueue = append(workingQueue, waitingQueue...)
		waitingQueue = []uint{}
	}
}

// HandleUserCommand handles the user command for controlling the air conditioner device.
// It takes a dto.Command as input and performs the corresponding action based on the command type.
// The command can be TurnOn, TurnOff, SetTemperature, or SetSpeed.
// For TurnOn command, it turns on the air conditioner device and adds the device ID to the working queue.
// For TurnOff command, it turns off the air conditioner device and removes the device ID from both the working queue and waiting queue.
// For SetTemperature command, it sets the target temperature of the air conditioner device.
// For SetSpeed command, it sets the speed of the air conditioner device.
func (controlService ControlServiceImpl) HandleUserCommand(command dto.Command) {
	device := mock.GetAirConditionByDeviceID(command.DeviceID)

	switch command.Command {
	case dto.TurnOn:
		// Turn on the air conditioner device.
		device.Working = true
		mock.SetAirCondition(device)
		workingQueue = append(workingQueue, uint(command.DeviceID))
	case dto.TurnOff:
		// Turn off the air conditioner device.
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
		// Set the target temperature of the air conditioner device.
		targetTemperature, ok := command.Args[0].(float32)
		if !ok {
			panic("invalid type")
		}
		device.TargetTemperature = float32(targetTemperature)
		mock.SetAirCondition(device)
	case dto.SetSpeed:
		// Set the speed of the air conditioner device.
		speed, ok := command.Args[0].(string)
		if !ok {
			panic("invalid type")
		}
		device.Speed = speed
		mock.SetAirCondition(device)
	}
}
