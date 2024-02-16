package services

import (
	port "datalog-go/ports"
	"datalog-go/utils/logs"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

type group2 struct {
	amita       port.IAmita
	measurement string
}

func NewGroup2(amita port.IAmita, measurement string) group2 {
	return group2{
		amita:       amita,
		measurement: measurement,
	}
}

func (g group2) InsertDataToAmita(body interface{}) error {
	var m map[string]string
	t := make(map[string]string)
	f := make(map[string]interface{})
	i := Group2Dto{}
	mapstructure.Decode(body, &i)

	j, err := json.Marshal(body)
	if err != nil {
		return err
	}
	json.Unmarshal(j, &m)
	fmt.Println("========")
	fmt.Println(m)

	for k, ms := range m {
		switch k {
		case "DATALOG ID":
			t[k] = ms
		case "BATTERY ID":
			t[k] = ms
		default:
			f[k] = ms
		}
	}

	fmt.Println(reflect.TypeOf(f["Soc"]))
	err = g.amita.Write(t, f, g.measurement)
	if err != nil {
		return err
	}
	l := fmt.Sprintf("Insert batteryId:%v datalogId:%v to measurement:%v", t["DATALOG ID"], t["BATTERY ID"], g.measurement)
	logs.Info(l)

	return nil
}
