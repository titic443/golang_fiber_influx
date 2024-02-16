package services

import (
	port "datalog-go/ports"
	"datalog-go/utils/logs"
	"encoding/json"
	"fmt"
	"strconv"
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
	var m map[string]any
	t := make(map[string]string)
	f := make(map[string]interface{})
	j, err := json.Marshal(body)
	if err != nil {
		return err
	}
	json.Unmarshal(j, &m)
	fmt.Println(body)
	for k, ms := range m {
		switch k {
		case "DATALOG ID":
			t[k] = ms.(string)
		case "BATTERY ID":
			t[k] = ms.(string)
		default:
			fl, err := strconv.ParseFloat(ms.(string), 64)
			if err != nil {
				return err
			}
			f[k] = fl
		}
	}
	err = g.amita.Write(t, f, g.measurement)
	if err != nil {
		return err
	}
	l := fmt.Sprintf("Insert batteryId:%v datalogId:%v to measurement:%v", t["DATALOG ID"], t["BATTERY ID"], g.measurement)
	logs.Info(l)

	return nil
}
