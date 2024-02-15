package handlers

type Group1Dto struct {
	DATALOGID string `json:"DATALOG ID" validate:"required"`
	BATTERYID string `json:"BATTERY ID" validate:"required"`
	TotalVV   string `json:"Total V (V)" validate:"required"`
	TotalAA   string `json:"Total A (A)" validate:"required"`
	MinVMV    string `json:"MinV (mV)" validate:"required"`
	MaxVMV    string `json:"MaxV (mV)" validate:"required"`
}
