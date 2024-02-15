package services

type Group1Dto struct {
	BATTERYID   string `json:"BATTERY ID"`
	DATALOGID   string `json:"DATALOG ID"`
	Field       string `json:"_field"`
	Measurement string `json:"_measurement"`
	Result      string `json:"result"`
	Start       string `json:"_start"`
	Stop        string `json:"_stop"`
	Table       int    `json:"table"`
	Time        string `json:"_time"`
	Value       string `json:"_value"`
}

type IGroup1 interface {
	InsertDataToAmita() error
}
