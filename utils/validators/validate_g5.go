package validators

import (
	"datalog-go/utils/logs"
	"encoding/json"
	"strconv"
)

type Group5Dto struct {
	DATALOGID   string  `json:"DATALOG ID"`
	BATTERYID   string  `json:"BATTERY ID"`
	Soh         float64 `json:"SoH"`
	EvBmsStatus string  `json:"EvBmsStatus"`
}

func (g *Group5Dto) MapType(b []byte) []interface{} {
	var con []interface{}
	var m []map[string]string
	metric := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	if err != nil {
		logs.Error(err)
	}

	for _, v := range m {
		t := Group5Dto{
			DATALOGID:   v["DATALOG ID"],
			BATTERYID:   v["BATTERY ID"],
			EvBmsStatus: v["EvBmsStatus"],
		}

		for k, vv := range v {
			if k != "DATALOG ID" && k != "BATTERY ID" && k != "EvBmsStatus" {
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
