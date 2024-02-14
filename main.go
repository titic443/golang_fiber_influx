package main

import (
	"datalog-go/adapter"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	initTimeZone()
	token := viper.GetString("db.token")
	url := viper.GetString("db.url")
	client := influxdb2.NewClient(url, token)

	adapter.NewAmita(&client)

	// org := "ea"
	// bucket := "amita"
	// writeAPI := client.WriteAPIBlocking(org, bucket)
	// for value := 0; value < 10000; value++ {
	// 	tags := map[string]string{
	// 		"tagname1": "tagvalue1",
	// 	}
	// 	fields := map[string]interface{}{
	// 		"field1": value,
	// 	}
	// 	point := write.NewPoint("measurement1", tags, fields, time.Now())
	// 	time.Sleep(1 * time.Second) // separate points by 1 second

	// 	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// query(client, org)
}

// func query(client influxdb2.Client, org string) {
// 	queryAPI := client.QueryAPI(org)
// 	query := `from(bucket: "amita")
//             |> range(start: -10m)
//             |> filter(fn: (r) => r._measurement == "measurement1")`
// 	results, err := queryAPI.Query(context.Background(), query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for results.Next() {
// 		fmt.Println(results.Record())
// 	}
// 	if err := results.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	query = `from(bucket: "amita")
//               |> range(start: -10m)
//               |> filter(fn: (r) => r._measurement == "measurement1")
//               |> mean()`
// 	results, err = queryAPI.Query(context.Background(), query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for results.Next() {
// 		fmt.Println(results.Record())
// 	}
// 	if err := results.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }
