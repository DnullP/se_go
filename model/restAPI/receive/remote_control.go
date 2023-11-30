package receive

type UserCommandReceive struct {
	Args    []interface{} `json:"args"`
	Command string        `json:"command"`
}

const (
	SetSpeed       = "set_speed"
	SetTemperature = "set_temperature"
	TurnOff        = "turn_off"
	TurnOn         = "turn_on"
)
