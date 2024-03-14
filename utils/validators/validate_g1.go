package validators

import (
	"datalog-go/utils/logs"
	"encoding/json"
	"strconv"
)

type Group1Dto struct {
	DATALOGID string  `json:"DATALOG ID" `
	BATTERYID string  `json:"BATTERY ID"`
	TotalVV   float64 `json:"Total V (V)"`
	TotalAA   float64 `json:"Total A (A)"`
	MinVMV    float64 `json:"MinV (mV)"`
	MaxVMV    float64 `json:"MaxV (mV)"`
}

func (g *Group1Dto) MapType(b []byte) []interface{} {
	var con []interface{}
	var m []map[string]string
	metric := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	if err != nil {
		logs.Error(err)
	}
	for _, v := range m {
		t := Group1Dto{
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
