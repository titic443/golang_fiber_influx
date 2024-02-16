package services

type Group4Dto struct {
	DATALOGID  string
	BATTERYID  string
	Fault1     float64
	Fault2     float64
	Fault3     float64
	Fault4     float64
	Fault5     float64
	Fault6     float64
	Fault7     float64
	Fault8     float64
	RelayState float64
}

type IGroup4 interface {
	InsertDataToAmita(interface{}) error
}
