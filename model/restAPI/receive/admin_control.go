package receive

type AdminCommandReceive struct {
	Args    ArgsACR   `json:"args"`
	Command string `json:"command"`
}

type ArgsACR struct {
	FeeRate *float64 `json:"fee_rate,omitempty"`
	// 'cool'/'warm'
	Mode           *string `json:"mode,omitempty"`
	ValidRangeHigh *int64  `json:"valid_range_high,omitempty"`
	ValidRangeLow  *int64  `json:"valid_range_low,omitempty"`
}

const (
	SetMode       = "set_mode"
	SetPrice      = "set_price"
	SetValidRange = "set_valid_range"
)
