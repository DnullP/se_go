package response

type PanelStatus struct {
	// a float
	EnvTemperature float64 `json:"env_temperature"`
	// low, mid, high
	Speed *string `json:"speed"`
	// a float
	TargetTemperature float64 `json:"target_temperature"`
	// true is working
	Working bool `json:"working"`
}
