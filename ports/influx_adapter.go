package port

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
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

func (a amita) Write(tags map[string]string, fields map[string]interface{}, measurement string) error {
	writeAPI := a.client.WriteAPIBlocking(a.org, a.bucket)

	p := influxdb2.NewPoint(measurement, tags, fields, time.Now())
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		return err
	}

	return nil
}

func (a amita) Select1H(measurement string) (interface{}, error) {
	var results []interface{}
	queryAPI := a.client.QueryAPI(a.org)
	query := fmt.Sprintf(`from(bucket: "%v")
	|> range(start: -1h)
	|> filter(fn: (r) => r._measurement == "%v")`, a.bucket, measurement)

	r, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	for r.Next() {
		results = append(results, r.Record())
	}
	if err := r.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
