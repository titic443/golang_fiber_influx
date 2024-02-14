package port

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type amita struct {
	client influxdb2.Client
	org    string
	bucket string
}

func NewAmita(client influxdb2.Client, org string, bucket string) amita {
	return amita{
		client: client,
		org:    org,
		bucket: bucket,
	}
}

func (a amita) Write(measurement string) error {
	writeAPI := a.client.WriteAPIBlocking(a.org, a.bucket)
	for value := 0; value < 10; value++ {
		tags := map[string]string{
			"DATALOG ID": "L0001",
			"BATTERY ID": "BAT0002",
		}
		fields := map[string]interface{}{
			"Total V (V)": 690,
			"Total A (A)": -6.7,
			"MinV (mV)":   4072,
			"MaxV (mV)":   4126,
		}
		point := write.NewPoint(measurement, tags, fields, time.Now())
		time.Sleep(1 * time.Second)
		fmt.Println(tags)
		fmt.Println(fields)
		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			return err
		}
	}

	return nil
}

func (a amita) Select(measurement string) (interface{}, error) {
	var result []interface{}
	queryAPI := a.client.QueryAPI(a.org)
	query := fmt.Sprintf(`from(bucket: "%v")
	|> range(start: -40m)
	|> filter(fn: (r) => r._measurement == "%v")`, a.bucket, measurement)

	results, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	for results.Next() {
		result = append(result, results.Record())
	}
	if err := results.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
