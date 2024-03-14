package validators

import (
	"datalog-go/utils/logs"
	"encoding/json"
	"fmt"
	"strconv"
)

type Group3Dto struct {
	DATALOGID string  `json:"DATALOG ID" `
	BATTERYID string  `json:"BATTERY ID" `
	T1        float64 `json:"T1"`
	T2        float64 `json:"T2"`
	T3        float64 `json:"T3"`
	T4        float64 `json:"T4"`
	T5        float64 `json:"T5"`
	T6        float64 `json:"T6"`
	T7        float64 `json:"T7"`
	T8        float64 `json:"T8"`
	T9        float64 `json:"T9"`
	T10       float64 `json:"T10"`
	T11       float64 `json:"T11"`
	T12       float64 `json:"T12"`
	T13       float64 `json:"T13"`
	T14       float64 `json:"T14"`
	T15       float64 `json:"T15"`
	T16       float64 `json:"T16"`
	T17       float64 `json:"T17"`
	T18       float64 `json:"T18"`
	T19       float64 `json:"T19"`
	T20       float64 `json:"T20"`
	T21       float64 `json:"T21"`
	T22       float64 `json:"T22"`
	T23       float64 `json:"T23"`
	T24       float64 `json:"T24"`
	T25       float64 `json:"T25"`
	T26       float64 `json:"T26"`
	T27       float64 `json:"T27"`
	T28       float64 `json:"T28"`
	T29       float64 `json:"T29"`
	T30       float64 `json:"T30"`
	T31       float64 `json:"T31"`
	T32       float64 `json:"T32"`
	T33       float64 `json:"T33"`
	T34       float64 `json:"T34"`
	T35       float64 `json:"T35"`
	T36       float64 `json:"T36"`
	T37       float64 `json:"T37"`
	T38       float64 `json:"T38"`
	T39       float64 `json:"T39"`
	T40       float64 `json:"T40"`
	T41       float64 `json:"T41"`
	T42       float64 `json:"T42"`
	T43       float64 `json:"T43"`
	T44       float64 `json:"T44"`
	T45       float64 `json:"T45"`
	T46       float64 `json:"T46"`
	T47       float64 `json:"T47"`
	T48       float64 `json:"T48"`
}

func (g *Group3Dto) MapType(b []byte) []interface{} {
	// tmp := []Group1Dto{}
	var con []interface{}
	var m []map[string]string
	metric := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	if err != nil {
		logs.Error(err)
	}
	for _, v := range m {
		t := Group3Dto{
			DATALOGID: v["DATALOG ID"],
			BATTERYID: v["BATTERY ID"],
		}
		fmt.Println(v)
		fmt.Println("-----")
		for k, vv := range v {
			if k != "DATALOG ID" && k != "BATTERY ID" {

				f, err := strconv.ParseFloat(vv, 64)
				if err != nil {
					metric[k] = 0
				}
				metric[k] = f
			}
		}
		fmt.Println(metric)

		b, err := json.Marshal(metric)
		if err != nil {
			logs.Error(err)
		}
		json.Unmarshal(b, &t)
		con = append(con, t)
		fmt.Println(con...)
	}

	return con
}
