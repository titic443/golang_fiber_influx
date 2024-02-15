package services

type Group2Dto struct {
	DATALOGID string
	BATTERYID string
	TotalVV   float64
	TotalAA   float64
	MinVMV    float64
	MaxVMV    float64
}

type IGroup2 interface {
	InsertDataToAmita(interface{}) error
}
