package models

type OptionChainEntry struct {
	ID uint `gorm:"primaryKey"`

	StrikePrice float64 `json:"strikePrice"`

	// Call Option Data
	CITMP      *float64 `json:"citmp"`
	CallDelta  *float64 `json:"callDelta"`
	CallGamma  *float64 `json:"callGamma"`
	CallIV     *float64 `json:"callIV"`
	CallLTP    *float64 `json:"callLtp"`
	CallOI     *int64   `json:"callOi"`
	CallTheta  *float64 `json:"callTheta"`
	CallVega   *float64 `json:"callVega"`
	CallVolume *int64   `json:"callVolume"`

	// Put Option Data
	PITMP     *float64 `json:"pitmp"`
	PutDelta  *float64 `json:"putDelta"`
	PutGamma  *float64 `json:"putGamma"`
	PutIV     *float64 `json:"putIv"`
	PutLTP    *float64 `json:"putLtp"`
	PutOI     *int64   `json:"putOi"`
	PutTheta  *float64 `json:"putTheta"`
	PutVega   *float64 `json:"putVega"`
	PutVolume *int64   `json:"putVolume"`
}

