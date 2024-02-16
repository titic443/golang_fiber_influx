package validators

import (
	"encoding/json"
	"fmt"
)

type Group1Dto struct {
	DATALOGID string  `json:"DATALOG ID" `
	BATTERYID string  `json:"BATTERY ID"`
	TotalVV   float64 `json:"Total V (V),string"`
	TotalAA   float64 `json:"Total A (A),string"`
	MinVMV    float64 `json:"MinV (mV),string"`
	MaxVMV    float64 `json:"MaxV (mV),string"`
}

func (g *Group1Dto) MapType(b []byte) []interface{} {
	var tmp []Group1Dto
	var con []interface{}
	fmt.Println(b)
	json.Unmarshal(b, &tmp)
	fmt.Println(len(tmp))
	for _, item := range tmp {
		con = append(con, item)
	}
	return con
}
