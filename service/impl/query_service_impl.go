package impl

import "github.com/DnullP/se_work/mock"

type QueryServiceImpl struct{}

func (queryService QueryServiceImpl) GetDeviceStatus(deviceID uint) mock.UnsafeAirConditionMock {
	return mock.GetAirConditionByDeviceID(deviceID)
}
