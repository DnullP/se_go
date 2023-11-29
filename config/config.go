package config

import (
	"github.com/DnullP/se_work/mock"
)

// mock config
var (
	MockNum              = 5
	AirConditionMockList = []mock.AirConditionMock{
		{
			DeviceID:           0,
			NatureTemperature:  32,
			CurrentTemperature: 32,
			TargetTemperature:  32,
			Speed:              "low",
		},
		{
			DeviceID:           1,
			NatureTemperature:  28,
			CurrentTemperature: 28,
			TargetTemperature:  28,
			Speed:              "low",
		}, {
			DeviceID:           2,
			NatureTemperature:  30,
			CurrentTemperature: 30,
			TargetTemperature:  30,
			Speed:              "low",
		}, {
			DeviceID:           3,
			NatureTemperature:  29,
			CurrentTemperature: 29,
			TargetTemperature:  29,
			Speed:              "low",
		}, {
			DeviceID:           4,
			NatureTemperature:  35,
			CurrentTemperature: 35,
			TargetTemperature:  35,
			Speed:              "low",
		},
	}
)
