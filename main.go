package main

import (
	handlers "datalog-go/handler"
	adapter "datalog-go/ports"
	"datalog-go/services"
	"datalog-go/utils/logs"
	"datalog-go/utils/validators"
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
	measurement2 := viper.GetString("db.measurement2")

	client := influxdb2.NewClient(url, token)
	amita := adapter.NewAmita(client, org, bucket)
	g1 := services.NewGroup1(amita, measurement1)
	g2 := services.NewGroup2(amita, measurement2)
	h := handlers.NewHandler(g1, g2)

	app.Post("/g1", validators.ValidateG1, h.InsertG1)
	app.Post("/g2", validators.ValidateG2, h.InsertG2)

	l := fmt.Sprintf("App start on port %v", port)
	logs.Info(l)
	app.Listen(port)
}
