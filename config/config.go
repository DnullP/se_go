package config

// mock config
var (
	MockNum             = 5
	MockTemperatureList = []float32{32.0, 28.0, 30.0, 29.0, 35.0}
)

const (
	// work mode
	COLD = false
	WARM = true
)

// central control config
var (
	ValidUpperTemperature float32 = 28.0
	ValidLowerTemperature float32 = 18.0

	PriceRateLow float32 = 1.0
	PriceRateMid float32 = 1.0
	PriceRateHig float32 = 1.0

	WorkMode = COLD
)
