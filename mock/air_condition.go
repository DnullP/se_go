package mock

import "time"

type AirConditionMock struct {
	DeviceID           uint
	NatureTemperature  float32
	CurrentTemperature float32
	TargetTemperature  float32
	Speed              string
}

var AirConditionMockList = []AirConditionMock{}

func updateAirConditionMockList() {
	for ;; {
		for i := 0; i < len(AirConditionMockList); i++ {
			AirConditionMockList[i].CurrentTemperature -= 0.5;
		}
		time.Sleep(10 * time.Second)
	}
}

func InitAirConditionMock(mock_num int, nature_temperature_list []float32) {
	for i := 0; i < mock_num; i++ {
		AirConditionMockList = append(AirConditionMockList, AirConditionMock{
			DeviceID:           uint(i),
			NatureTemperature:  nature_temperature_list[i],
			CurrentTemperature: nature_temperature_list[i],
			TargetTemperature:  25,
			Speed:              "low",
		})
	}
	go updateAirConditionMockList()
}