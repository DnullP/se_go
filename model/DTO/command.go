package dto

// user command
type Command struct {
	Command  string        `json:"command"`
	DeviceID uint          `json:"device_id"`
	Args     []interface{} `json:"args"`
}

const (
	SetSpeed       = "set_speed"
	SetTemperature = "set_temperature"
	TurnOff        = "turn_off"
	TurnOn         = "turn_on"
)

const (
	SetMode       = "set_mode"
	SetPrice      = "set_price"
	SetValidRange = "set_valid_range"
	TurnOffAll    = "turn_off_all"
	TurnOnAll     = "turn_on_all"
)
