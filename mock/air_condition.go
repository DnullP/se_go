package mock

import (
	"sync"
	"time"

	"github.com/DnullP/se_work/config"
)

type AirConditionMock struct {
	Data UnsafeAirConditionMock
	lock sync.Mutex
}

type UnsafeAirConditionMock struct {
	DeviceID           uint
	Working            bool
	NatureTemperature  float32
	CurrentTemperature float32
	TargetTemperature  float32
	Speed              string
	TotalCost          float32
}

var AirConditionMockList = []AirConditionMock{}

// GetAirConditionByDeviceID retrieves the air condition data for the specified device ID.
// It returns an UnsafeAirConditionMock object representing the air condition data.
func GetAirConditionByDeviceID(deviceID uint) UnsafeAirConditionMock {
	return AirConditionMockList[deviceID].Data
}

// SetAirCondition sets the parameters of an air condition device.
// It takes an UnsafeAirConditionMock struct as input, which contains the device ID, working status, target temperature, and speed.
// The function updates the corresponding air condition device with the provided parameters.
// It acquires a lock on the device, updates the parameters, and releases the lock.
func SetAirCondition(set UnsafeAirConditionMock) {
	aircondition := &AirConditionMockList[set.DeviceID]

	aircondition.lock.Lock()
	aircondition.Data.Working = set.Working
	aircondition.Data.TargetTemperature = set.TargetTemperature
	aircondition.Data.Speed = set.Speed
	aircondition.lock.Unlock()
}

// InitAirConditionMock initializes the mock air condition list with the specified number of mock devices.
// It populates the list with mock air condition objects, each having a unique device ID and initial temperature values.
// The mock devices are initially set to be not working, with a low fan speed and zero total cost.
// This function also starts a goroutine to update the mock air condition list periodically.
func InitAirConditionMock() {
	for i := 0; i < config.MockNum; i++ {
		AirConditionMockList = append(AirConditionMockList, AirConditionMock{
			Data: UnsafeAirConditionMock{
				DeviceID:           uint(i),
				Working:            false,
				NatureTemperature:  config.MockTemperatureList[i],
				CurrentTemperature: config.MockTemperatureList[i],
				TargetTemperature:  config.MockTemperatureList[i],
				Speed:              "low",
				TotalCost:          0,
			},
		})
	}
	go updateAirConditionMockList()
}

// updateAirConditionMockList updates the air condition mock list periodically.
// It simulates the behavior of air condition units by changing the temperature based on various conditions.
func updateAirConditionMockList() {
	var minutes uint = 0
	for {
		for i := 0; i < len(AirConditionMockList); i++ {
			natureTemperatureChange(i)

			aircon := AirConditionMockList[i].Data
			if aircon.Working && aircon.Speed == "high" {
				artificialTemperatureChange(i, 1.5)
			}
		}

		if minutes%2 == 0 {
			for i := 0; i < len(AirConditionMockList); i++ {
				aircon := AirConditionMockList[i].Data
				if aircon.Working && aircon.Speed == "mid" {
					artificialTemperatureChange(i, 1.5)
				}
			}
		}

		if minutes%3 == 0 {
			for i := 0; i < len(AirConditionMockList); i++ {
				aircon := AirConditionMockList[i].Data
				if aircon.Working && aircon.Speed == "low" {
					artificialTemperatureChange(i, 1.5)
				}
			}
		}

		minutes += 1
		time.Sleep(10 * time.Second)
	}
}

// smile code, change it or don't read it
// natureTemperatureChange simulates the change in temperature based on the work mode of the air conditioner.
// It takes an index 'i' as input to identify the air conditioner in the AirConditionMockList.
// The function locks the air conditioner, checks the work mode, and adjusts the current temperature accordingly.
// If the work mode is set to config.COLD, the current temperature is increased by 0.5 if it is lower than the nature temperature.
// If the work mode is set to config.WARM, the current temperature is decreased by 0.5 if it is higher than the nature temperature.
// Finally, the function unlocks the air conditioner.
func natureTemperatureChange(i int) {
	AirConditionMockList[i].lock.Lock()

	switch config.WorkMode {

	case config.COLD:
		if AirConditionMockList[i].Data.CurrentTemperature < AirConditionMockList[i].Data.NatureTemperature {
			AirConditionMockList[i].Data.CurrentTemperature += 0.5
		}

	case config.WARM:
		if AirConditionMockList[i].Data.CurrentTemperature > AirConditionMockList[i].Data.NatureTemperature {
			AirConditionMockList[i].Data.CurrentTemperature -= 0.5
		}
	}

	AirConditionMockList[i].lock.Unlock()
}

// artificialTemperatureChange simulates an artificial temperature change for a specific air condition mock.
// It takes an index i and a temperature change value c as parameters.
// The function updates the current temperature of the air condition mock based on the work mode and target temperature.
// If the work mode is set to config.COLD, the current temperature is decreased by c until it reaches the target temperature.
// If the work mode is set to config.WARM, the current temperature is increased by c until it reaches the target temperature.
// If the current temperature reaches the target temperature, the air condition mock's working status is set to false.
// The function also updates the total cost of the air condition mock based on the price rate.
// The function uses a lock to ensure thread safety when accessing and modifying the air condition mock's data.
func artificialTemperatureChange(i int, c float32) {
	AirConditionMockList[i].lock.Lock()

	switch config.WorkMode {

	case config.COLD:
		if AirConditionMockList[i].Data.CurrentTemperature > AirConditionMockList[i].Data.TargetTemperature {
			AirConditionMockList[i].Data.CurrentTemperature -= c
		} else {
			AirConditionMockList[i].Data.Working = false
		}

	case config.WARM:
		if AirConditionMockList[i].Data.CurrentTemperature < AirConditionMockList[i].Data.TargetTemperature {
			AirConditionMockList[i].Data.CurrentTemperature += c
		} else {
			AirConditionMockList[i].Data.Working = false
		}
	}

	AirConditionMockList[i].Data.TotalCost += config.PriceRate

	AirConditionMockList[i].lock.Unlock()
}
