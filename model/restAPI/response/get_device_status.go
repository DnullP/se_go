package response

type DeviceStatus struct {
	// a float
	EnvTemperature float64 `json:"env_temperature"`
	// true is warm mode
	Mode bool `json:"mode"`
	// low, mid, high
	Speed *string `json:"speed"`
	// a float
	TargetTemperature float64 `json:"target_temperature"`
	// a float
	TotalCost float64 `json:"total_cost"`
	// true is working
	Working bool `json:"working"`
}
