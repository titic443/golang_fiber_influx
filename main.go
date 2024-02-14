package main

import (
	adapter "datalog-go/ports"
	"datalog-go/utils/logs"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var log *zap.Logger

func main() {
	initConfig()
	initTimeZone()
	logs.Info("start app")
	token := viper.GetString("db.token")
	url := viper.GetString("db.url")
	org := viper.GetString("db.org")
	bucket := viper.GetString("db.bucket")

	client := influxdb2.NewClient(url, token)
	amita := adapter.NewAmita(client, org, bucket)

	err := amita.Write("measurement4")
	if err != nil {
		logs.Error(err)
	}

}
