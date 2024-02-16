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

	v := &validators.ValidateBody{}
	vg1 := &validators.Group1Dto{}
	vg2 := &validators.Group2Dto{}
	vg3 := &validators.Group3Dto{}
	vg4 := &validators.Group4Dto{}
	vg5 := &validators.Group4Dto{}

	port := ":" + viper.GetString("app.port")
	token := viper.GetString("db.token")
	url := viper.GetString("db.url")
	org := viper.GetString("db.org")
	bucket := viper.GetString("db.bucket")
	measurement1 := viper.GetString("db.measurement1")
	measurement2 := viper.GetString("db.measurement2")
	measurement3 := viper.GetString("db.measurement3")
	measurement4 := viper.GetString("db.measurement4")
	measurement5 := viper.GetString("db.measurement5")

	client := influxdb2.NewClient(url, token)
	amita := adapter.NewAmita(client, org, bucket)
	g1 := services.NewGroup1(amita, measurement1)
	g2 := services.NewGroup2(amita, measurement2)
	g3 := services.NewGroup3(amita, measurement3)
	g4 := services.NewGroup4(amita, measurement4)
	g5 := services.NewGroup4(amita, measurement5)
	h := handlers.NewHandler(g1, g2, g3, g4, g5)

	app.Post("/g1", v.ValidateBody(vg1), h.InsertG1)
	app.Post("/g2", v.ValidateBody(vg2), h.InsertG2)
	app.Post("/g3", v.ValidateBody(vg3), h.InsertG3)
	app.Post("/g4", v.ValidateBody(vg4), h.InsertG4)
	app.Post("/g5", v.ValidateBody(vg5), h.InsertG5)

	l := fmt.Sprintf("App start on port %v", port)
	logs.Info(l)
	app.Listen(port)
}
