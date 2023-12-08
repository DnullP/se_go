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

func GetAirConditionByDeviceID(deviceID uint) UnsafeAirConditionMock {
	return AirConditionMockList[deviceID].Data
}

func SetAirCondition(set UnsafeAirConditionMock) {
	aircondition := &AirConditionMockList[set.DeviceID]

	aircondition.lock.Lock()
	aircondition.Data.Working = set.Working
	aircondition.Data.TargetTemperature = set.TargetTemperature
	aircondition.Data.Speed = set.Speed
	aircondition.lock.Unlock()
}

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
