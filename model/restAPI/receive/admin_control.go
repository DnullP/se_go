package receive

type AdminCommandReceive struct {
	Args    []interface{} `json:"args"`
	Command string        `json:"command"`
}

const (
	SetMode       = "set_mode"
	SetPrice      = "set_price"
	SetValidRange = "set_valid_range"
)
