package database

import (
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var client influxdb2.Client

func GetClient() influxdb2.Client{	
	if client == nil {
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		token := os.Getenv("DB_TOKEN")
		url := "http://" + host + ":" + port 

		client = influxdb2.NewClient(url, token)
	}

	return client
}
