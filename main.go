package main

import (
	adapter "datalog-go/ports"
	"datalog-go/utils/logs"

	"github.com/gofiber/fiber/v2"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/spf13/viper"
)

func main() {

	app := fiber.New()

	logs.Info("start app")
	port := ":" + viper.GetString("app.port")
	token := viper.GetString("db.token")
	url := viper.GetString("db.url")
	org := viper.GetString("db.org")
	bucket := viper.GetString("db.bucket")
	measurement1 := viper.GetString("db.measurement1")

	client := influxdb2.NewClient(url, token)
	amita := adapter.NewAmita(client, org, bucket, measurement1)
	t := map[string]string{
		"DATALOG ID": "L0001",
		"BATTERY ID": "BAT0002",
	}
	f := map[string]interface{}{
		"Total V (V)": 690,
		"Total A (A)": -6.7,
		"MinV (mV)":   4072,
		"MaxV (mV)":   4126,
	}
	err := amita.Write(t, f)
	if err != nil {
		logs.Error(err)
	}
	_, err = amita.Select("measurement1")
	if err != nil {
		logs.Error(err)
	}
	app.Listen(port)
}
