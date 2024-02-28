package validators

import (
	"datalog-go/utils/logs"
	"encoding/json"
)

type Group4Dto struct {
	DATALOGID  string  `json:"DATALOG ID" `
	BATTERYID  string  `json:"BATTERY ID" `
	Fault1     float64 `json:"Fault1,string" `
	Fault2     float64 `json:"Fault2,string" `
	Fault3     float64 `json:"Fault3,string" `
	Fault4     float64 `json:"Fault4,string" `
	Fault5     float64 `json:"Fault5,string" `
	Fault6     float64 `json:"Fault6,string" `
	Fault7     float64 `json:"Fault7,string" `
	Fault8     float64 `json:"Fault8,string" `
	RelayState float64 `json:"RelayState,string" `
}

func (g *Group4Dto) MapType(b []byte) []interface{} {
	var tmp []Group4Dto
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
