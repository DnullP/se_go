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

func updateAirConditionMockList() {
	for {
		for i := 0; i < len(AirConditionMockList); i++ {
			AirConditionMockList[i].lock.Lock()
			if AirConditionMockList[i].Data.CurrentTemperature < AirConditionMockList[i].Data.TargetTemperature {
				AirConditionMockList[i].Data.CurrentTemperature += 0.5
			} else if AirConditionMockList[i].Data.CurrentTemperature > AirConditionMockList[i].Data.TargetTemperature {
				AirConditionMockList[i].Data.CurrentTemperature -= 0.5
			}
			AirConditionMockList[i].lock.Unlock()
		}
		time.Sleep(10 * time.Second)
	}
}

func UpdateWorkingAirCondition(workingQueue []int) {
	for i := 0; i < len(workingQueue); i++ {
		aircondition := &AirConditionMockList[workingQueue[i]]

		aircondition.lock.Lock()
		if aircondition.Data.Working {
			switch AirConditionMockList[i].Data.Speed {
			case "low":
				AirConditionMockList[i].Data.TotalCost += 0.5 + 1.0/3
			case "medium":
				AirConditionMockList[i].Data.TotalCost += 0.5 + 1.0/2
			case "high":
				AirConditionMockList[i].Data.TotalCost += 0.5 + 1.0
			default:

			}
		}
		aircondition.lock.Unlock()
	}
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
