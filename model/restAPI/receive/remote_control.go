package receive

type UserCommandReceive struct {
	Args    Args   `json:"args"`
	Command string `json:"command"`
}

type Args struct {
	CurrentRoomTemp *float64 `json:"current_room_temp,omitempty"`
	Speed           *string  `json:"speed,omitempty"`
	TargetTemp      *float64 `json:"target_temp,omitempty"`
}

type Command string

const (
	SetSpeed       = "set_speed"
	SetTemperature = "set_temperature"
	TurnOff        = "turn_off"
	TurnOn         = "turn_on"
)
