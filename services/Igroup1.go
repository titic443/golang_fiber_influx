package services

type Group1Dto struct {
	DATALOGID string
	BATTERYID string
	TotalVV   float64
	TotalAA   float64
	MinVMV    float64
	MaxVMV    float64
}

type IGroup1 interface {
	InsertDataToAmita(interface{}) error
}
