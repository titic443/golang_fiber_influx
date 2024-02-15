package main

import (
	adapter "datalog-go/ports"
	"datalog-go/services"
	"datalog-go/utils/logs"
	"fmt"

	"github.com/gofiber/fiber/v2"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/spf13/viper"
)

func main() {

	app := fiber.New()

	port := ":" + viper.GetString("app.port")
	token := viper.GetString("db.token")
	url := viper.GetString("db.url")
	org := viper.GetString("db.org")
	bucket := viper.GetString("db.bucket")
	measurement1 := viper.GetString("db.measurement1")

	client := influxdb2.NewClient(url, token)
	amita := adapter.NewAmita(client, org, bucket)
	g1 := services.NewGroup1(amita, measurement1)

	err := g1.InsertDataToAmita()
	if err != nil {
		logs.Error(err)
	}
	l := fmt.Sprintf("App start on port: %v", port)
	logs.Info(l)
	app.Listen(port)
}
