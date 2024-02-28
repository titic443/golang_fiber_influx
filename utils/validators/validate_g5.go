package validators

import (
	"datalog-go/utils/logs"
	"encoding/json"
)

type Group5Dto struct {
	DATALOGID   string  `json:"DATALOG ID"`
	BATTERYID   string  `json:"BATTERY ID"`
	Soh         float64 `json:"SoH,string"`
	EvBmsStatus float64 `json:"EvBmsStatus,string"`
}

func (g *Group5Dto) MapType(b []byte) []interface{} {
	var tmp []Group5Dto
	var con []interface{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		logs.Error(err)
		logs.Info("convert to default value")
		// return nil
	}
	for _, t := range tmp {
		con = append(con, t)
	}
	return con
}
