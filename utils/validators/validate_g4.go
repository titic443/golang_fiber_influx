package validators

import (
	"datalog-go/utils/logs"
	"encoding/json"
	"strconv"
)

type Group4Dto struct {
	DATALOGID  string  `json:"DATALOG ID" `
	BATTERYID  string  `json:"BATTERY ID" `
	Fault1     float64 `json:"Fault1" `
	Fault2     float64 `json:"Fault2" `
	Fault3     float64 `json:"Fault3" `
	Fault4     float64 `json:"Fault4" `
	Fault5     float64 `json:"Fault5" `
	Fault6     float64 `json:"Fault6" `
	Fault7     float64 `json:"Fault7" `
	Fault8     float64 `json:"Fault8" `
	RelayState float64 `json:"RelayState" `
}

func (g *Group4Dto) MapType(b []byte) []interface{} {
	var con []interface{}
	var m []map[string]string
	metric := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	if err != nil {
		logs.Error(err)
	}
	for _, v := range m {
		t := Group4Dto{
			DATALOGID: v["DATALOG ID"],
			BATTERYID: v["BATTERY ID"],
		}

		for k, vv := range v {
			if k != "DATALOG ID" && k != "BATTERY ID" {

				f, err := strconv.ParseFloat(vv, 64)
				if err != nil {
					metric[k] = 0
				}
				metric[k] = f
			}
		}

		b, err := json.Marshal(metric)
		if err != nil {
			logs.Error(err)
		}
		json.Unmarshal(b, &t)
		con = append(con, t)

	}

	return con
}
