package validators

import (
	"datalog-go/utils/logs"
	"encoding/json"
)

type Group3Dto struct {
	DATALOGID string  `json:"DATALOG ID" `
	BATTERYID string  `json:"BATTERY ID" `
	T1        float64 `json:"T1,string" `
	T2        float64 `json:"T2,string" `
	T3        float64 `json:"T3,string" `
	T4        float64 `json:"T4,string" `
	T5        float64 `json:"T5,string" `
	T6        float64 `json:"T6,string" `
	T7        float64 `json:"T7,string" `
	T8        float64 `json:"T8,string" `
	T9        float64 `json:"T9,string" `
	T10       float64 `json:"T10,string" `
	T11       float64 `json:"T11,string" `
	T12       float64 `json:"T12,string" `
	T13       float64 `json:"T13,string" `
	T14       float64 `json:"T14,string" `
	T15       float64 `json:"T15,string" `
	T16       float64 `json:"T16,string" `
	T17       float64 `json:"T17,string" `
	T18       float64 `json:"T18,string" `
	T19       float64 `json:"T19,string" `
	T20       float64 `json:"T20,string" `
	T21       float64 `json:"T21,string" `
	T22       float64 `json:"T22,string" `
	T23       float64 `json:"T23,string" `
	T24       float64 `json:"T24,string" `
	T25       float64 `json:"T25,string" `
	T26       float64 `json:"T26,string" `
	T27       float64 `json:"T27,string" `
	T28       float64 `json:"T28,string" `
	T29       float64 `json:"T29,string" `
	T30       float64 `json:"T30,string" `
	T31       float64 `json:"T31,string" `
	T32       float64 `json:"T32,string" `
	T33       float64 `json:"T33,string" `
	T34       float64 `json:"T34,string" `
	T35       float64 `json:"T35,string" `
	T36       float64 `json:"T36,string" `
	T37       float64 `json:"T37,string" `
	T38       float64 `json:"T38,string" `
	T39       float64 `json:"T39,string" `
	T40       float64 `json:"T40,string" `
	T41       float64 `json:"T41,string" `
	T42       float64 `json:"T42,string" `
	T43       float64 `json:"T43,string" `
	T44       float64 `json:"T44,string" `
	T45       float64 `json:"T45,string" `
	T46       float64 `json:"T46,string" `
	T47       float64 `json:"T47,string" `
	T48       float64 `json:"T48,string" `
}

func (g *Group3Dto) MapType(b []byte) []interface{} {
	var tmp []Group3Dto
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
