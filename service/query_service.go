package service

import "github.com/DnullP/se_work/mock"

type QueryService interface {
	GetDeviceStatus(uint) mock.UnsafeAirConditionMock
}
