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

	client := influxdb2.NewClient(url, token)
	amita := adapter.NewAmita(client, org, bucket)

	err := amita.Write("measurement1")
	if err != nil {
		logs.Error(err)
	}
	_, err = amita.Select("measurement1")
	if err != nil {
		logs.Error(err)
	}
	app.Listen(port)
}
