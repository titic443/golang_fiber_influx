package services

import (
	port "datalog-go/ports"
	"datalog-go/utils/logs"
	"encoding/json"
	"fmt"
	"strconv"
)

type group3 struct {
	amita       port.IAmita
	measurement string
}

func NewGroup3(amita port.IAmita, measurement string) group3 {
	return group3{
		amita:       amita,
		measurement: measurement,
	}
}

func (g group3) InsertDataToAmita(body interface{}) error {
	var m map[string]string
	t := make(map[string]string)
	f := make(map[string]interface{})

	j, err := json.Marshal(body)
	if err != nil {
		return err
	}
	json.Unmarshal(j, &m)

	for k, ms := range m {
		switch k {
		case "DATALOG ID":
			t[k] = ms
		case "BATTERY ID":
			t[k] = ms
		default:
			if fl, err := strconv.ParseFloat(ms, 64); err != nil {
				return err
			} else {
				f[k] = fl
			}
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
