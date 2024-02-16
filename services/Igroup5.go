package services

type Group5Dto struct {
	DATALOGID   string
	BATTERYID   string
	Soh         float64
	EvBmsStatus float64
}

type IGroup5 interface {
	InsertDataToAmita(interface{}) error
}
