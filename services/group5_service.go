package services

import (
	port "datalog-go/ports"
	"datalog-go/utils/logs"
	"encoding/json"
	"fmt"
)

type group5 struct {
	amita       port.IAmita
	measurement string
}

func NewGroup5(amita port.IAmita, measurement string) group5 {
	return group5{
		amita:       amita,
		measurement: measurement,
	}
}

func (g group5) InsertDataToAmita(body interface{}) error {
	var m map[string]any
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
			t[k] = ms.(string)
		case "BATTERY ID":
			t[k] = ms.(string)
		case "EvBmsStatus":
			t[k] = ms.(string)
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
