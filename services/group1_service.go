package services

import (
	port "datalog-go/ports"
	"datalog-go/utils/logs"
	"encoding/json"
	"fmt"
)

type group1 struct {
	amita       port.IAmita
	measurement string
}

func NewGroup1(amita port.IAmita, measurement string) group1 {
	return group1{
		amita:       amita,
		measurement: measurement,
	}
}

func (g group1) InsertDataToAmita(body interface{}) error {

	b, _ := json.Marshal(body)

	fmt.Println(string(b))

	t := map[string]string{
		"DATALOG ID": "L0001",
		"BATTERY ID": "BAT0002",
	}
	f := map[string]interface{}{
		"Total V (V)": 690,
		"Total A (A)": -6.7,
		"MinV (mV)":   4072,
		"MaxV (mV)":   4126,
	}

	err := g.amita.Write(t, f, g.measurement)
	if err != nil {
		return err
	}
	l := fmt.Sprintf("Insert batteryId:%v datalogId:%v to measurement:%v", t["DATALOG ID"], t["BATTERY ID"], g.measurement)
	logs.Info(l)

	return nil
}
