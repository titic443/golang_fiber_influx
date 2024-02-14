package adapter

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type amita struct {
	client *influxdb2.Client
}

func NewAmita(client *influxdb2.Client) *amita {
	return &amita{
		client: client,
	}
}
