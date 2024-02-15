package services

import (
	port "datalog-go/ports"
	"datalog-go/utils/logs"
	"fmt"

	"github.com/mitchellh/mapstructure"
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

	b := Group1Dto{}
	mapstructure.Decode(body, &b)

	t := map[string]string{
		"DATALOG ID": b.DATALOGID,
		"BATTERY ID": b.BATTERYID,
	}
	f := map[string]interface{}{
		"Total V (V)": b.TotalVV,
		"Total A (A)": b.TotalAA,
		"MinV (mV)":   b.MinVMV,
		"MaxV (mV)":   b.MaxVMV,
	}

	err := g.amita.Write(t, f, g.measurement)
	if err != nil {
		return err
	}
	l := fmt.Sprintf("Insert batteryId:%v datalogId:%v to measurement:%v", t["DATALOG ID"], t["BATTERY ID"], g.measurement)
	logs.Info(l)

	return nil
}
