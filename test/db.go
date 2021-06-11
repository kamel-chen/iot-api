package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	_ "github.com/joho/godotenv/autoload"
	gps "iot.api/internal/services/gps"
	num "iot.api/pkg/util/number"
)

type res struct {
	lat float64;
	lng float64;
}


var params = gps.GPS{
	Time: time.Now(),
	DeviceId: 1,
	Lat: num.GPSRound(num.RandomFloat(26, 25)),
	Lng: num.GPSRound(num.RandomFloat(122, 121)),
}

func main() {
	// token := os.Getenv("DB_TOKEN")

	// client := influxdb2.NewClient("http://localhost:8086", token)
	// defer client.Close()
	t1 := time.Now()

	// gps.PG_create(params)
	// gps.PG_findAll()

	// gps.SyncWrite()
	gps.FindAll()
	// findAll(client)
	// write(client, 1)

	t2 := time.Now()
	fmt.Println(t2.Sub(t1))

}

func round(num float64) float64 {
	return float64(int(num * 10000000)) / 10000000;
}

func randomFloat() float64 {
	rand.Seed(time.Now().UnixNano())
	return round(rand.Float64())
}

func write(client influxdb2.Client, t int) {
	bucket := os.Getenv("DB_BUCKET")
	org := os.Getenv("DB_ORG")

	// get non-blocking write client
	// writeAPI := client.WriteAPI(org, bucket)

	writeAPI := client.WriteAPIBlocking(org, bucket)

	for i := 0; i < 10000000; i++ {
		lat := 25 + randomFloat()
		lng := 121 + randomFloat()

		writeAPI.WriteRecord(context.Background(), fmt.Sprintf("bafang,type=gps lat=%f,lng=%f", lat, lng))
	}
}

func findAll(client influxdb2.Client) {
	bucket := os.Getenv("DB_BUCKET")
	org := os.Getenv("DB_ORG")

	query := fmt.Sprintf(`
		from(bucket:"%v")
			|> range(start: -12h)
			|> filter(fn: (r) => r._measurement == "bafang")
			|> pivot(rowKey:["_time"], columnKey:["_field"], valueColumn:"_value")`, bucket)

	log.Println(query)
	// Get query client
	queryAPI := client.QueryAPI(org)

	// get QueryTableResult
	result, err := queryAPI.Query(context.Background(), query)
	if err == nil {
		// Iterate over query response
		data := []res{}
		l := 0

		for result.Next() {
			l++
			// Notice when group key has changed
			if result.TableChanged() {
				fmt.Printf("table: %s\n\n", result.TableMetadata().String())
			}

			// Access data
			v := result.Record().Values()

			lng := v["lng"].(float64)
			lat := v["lat"].(float64)

			var m = res{
				lng: lng,
				lat: lat,
			}

			data = append(data, m)
			// fmt.Println(v["_measurement"], v["lng"], v["lat"], data)
		}
		fmt.Println(l)
		// check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %v\n", result.Err().Error())
		}
	} else {
		panic(err)
	}
}
