package impl

import "github.com/DnullP/se_work/mock"

type QueryServiceImpl struct{}

// GetDeviceStatus retrieves the device status for the given device ID.
// It returns a mock.UnsafeAirConditionMock object representing the device status.
func (queryService QueryServiceImpl) GetDeviceStatus(deviceID uint) mock.UnsafeAirConditionMock {
	return mock.GetAirConditionByDeviceID(deviceID)
}
