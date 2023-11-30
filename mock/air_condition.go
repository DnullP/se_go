package mock

import (
	"sync"
	"time"

	"github.com/DnullP/se_work/config"
)

type AirConditionMock struct {
	Data unsafeAirConditionMock
	lock sync.Mutex
}

type unsafeAirConditionMock struct {
	DeviceID           uint
	Working            bool
	NatureTemperature  float32
	CurrentTemperature float32
	TargetTemperature  float32
	Speed              string
	TotalCost          float32
}

var AirConditionMockList = []AirConditionMock{}

func GetAirConditionByDeviceID(deviceID uint) *AirConditionMock {
	return &AirConditionMockList[deviceID]
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
			Data: unsafeAirConditionMock{
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
