package handlers

type Group1Dto struct {
	DATALOGID string  `json:"DATALOG ID" validate:"required"`
	BATTERYID string  `json:"BATTERY ID" validate:"required"`
	TotalVV   float64 `json:"Total V (V),string" validate:"required"`
	TotalAA   float64 `json:"Total A (A),string" validate:"required"`
	MinVMV    float64 `json:"MinV (mV),string" validate:"required"`
	MaxVMV    float64 `json:"MaxV (mV),string" validate:"required"`
}
