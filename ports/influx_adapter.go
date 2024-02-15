package port

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/influxdata/influxdb-client-go/api/write"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type amita struct {
	client      influxdb2.Client
	org         string
	bucket      string
	measurement string
}

func NewAmita(client influxdb2.Client, org string, bucket string, measurement string) amita {
	return amita{
		client:      client,
		org:         org,
		bucket:      bucket,
		measurement: measurement,
	}
}

func (a amita) Write(tags map[string]string, fields map[string]interface{}) error {
	fmt.Println(tags)
	fmt.Println(fields)
	writeAPI := a.client.WriteAPIBlocking(a.org, a.bucket)

	point := write.NewPoint(a.measurement, tags, fields, time.Now())
	_ = point
	err := writeAPI.WritePoint(context.Background(), point)
	if err != nil {
		return err
	}

	return nil
}

func (a amita) Select(measurement string) (interface{}, error) {
	var results []interface{}
	queryAPI := a.client.QueryAPI(a.org)
	query := fmt.Sprintf(`from(bucket: "%v")
	|> range(start: -40m)
	|> filter(fn: (r) => r._measurement == "%v")`, a.bucket, a.measurement)

	r, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	for r.Next() {
		a := AmitaDto{}
		_ = a
		o := r.Record().Values()
		j, _ := json.Marshal(o)
		// json.Unmarshal(j, &a)
		results = append(results, string(j))
	}
	if err := r.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
