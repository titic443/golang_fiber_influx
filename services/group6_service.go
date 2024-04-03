package services

import (
	port "datalog-go/ports"
	"datalog-go/utils/logs"
	"encoding/json"
	"fmt"
	"slices"
)

type group6 struct {
	amita       port.IAmita
	measurement string
}

func NewGroup6(amita port.IAmita, measurement string) group6 {
	return group6{
		amita:       amita,
		measurement: measurement,
	}
}

func (g group6) InsertDataToAmita(body interface{}) error {
	var m map[string]any
	filter := []string{"DATALOG ID", "BATTERY ID", "EvBmsStatus", "date", "ts", "Fault1", "Fault2", "Fault3", "Fault4", "Fault5", "Fault6", "Fault7", "Fault8", "RelayState"}
	t := make(map[string]string)
	f := make(map[string]interface{})

	j, err := json.Marshal(body)
	if err != nil {
		return err
	}
	json.Unmarshal(j, &m)

	for k, ms := range m {
		checkString := slices.Contains(filter, k)
		switch checkString {
		case true:
			t[k] = ms.(string)
		// case "DATALOG ID":
		// 	t[k] = ms.(string)
		// case "BATTERY ID":
		// 	t[k] = ms.(string)
		// case "EvBmsStatus":
		// 	t[k] = ms.(string)
		default:

			f[k] = ms
		}
	}
	// fmt.Println(reflect.TypeOf(f["Soc"]))
	err = g.amita.Write(t, f, g.measurement)
	if err != nil {
		return err
	}
	l := fmt.Sprintf("Insert batteryId:%v datalogId:%v to measurement:%v", t["DATALOG ID"], t["BATTERY ID"], g.measurement)
	logs.Info(l)

	return nil
}
