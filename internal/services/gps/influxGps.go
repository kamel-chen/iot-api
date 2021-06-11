package gpsservice

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"iot.api/internal/database"
	num "iot.api/pkg/util/number"
)

type gpsData struct {
	lat float64;
	lng float64;
}

var client = database.GetClient()
var bucket = os.Getenv("DB_BUCKET")
var org = os.Getenv("DB_ORG")

func SyncWrite() {
	writeAPI := client.WriteAPIBlocking(org, bucket)

	for i := 0; i < 10000; i++ {
		lat := num.GPSRound(num.RandomFloat(26, 25))
		lng := num.GPSRound(num.RandomFloat(122, 121))

		writeAPI.WriteRecord(context.Background(), fmt.Sprintf("gps,deviceId=1 lat=%f,lng=%f", lat, lng))
	}
}

func  AsyncWrite(deviceId int64, lat float64, lng float64, direction float64, speed float64,) error {
	if lat == 0 || lng == 0 {
		return errors.New("params cannot null")
	}

	writeAPI := client.WriteAPI(org, bucket)

	// p := write.NewPoint(
	// 	"gps",
	// 	map[string]string{
	// 		"deviceId": strconv.FormatInt(deviceId, 10),
	// 	},
	// 	map[string]interface{}{
	// 		"lat": lat,
	// 		"lng": lng,
	// 		"direction": direction,
	// 		"speed": speed,
	// 	},
	// 	time.Now())
	// writeAPI.WritePoint(p)

	writeAPI.WriteRecord(fmt.Sprintf("gps,type=gps lat=%f,lng=%f", lat, lng))
	return nil
}

func AsyncRandomWrite(n int) {
	writeAPI := client.WriteAPI(org, bucket)

	errorsCh := writeAPI.Errors()
	// Create go proc for reading and logging errors
	go func() {
		for err := range errorsCh {
			fmt.Printf("write error: %s\n", err.Error())
		}
	}()

	for i := 0; i < n; i++ {
		lat := num.GPSRound(num.RandomFloat(26, 25))
		lng := num.GPSRound(num.RandomFloat(122, 121))

		p := write.NewPoint(
			"gps",
			map[string]string{
				"deviceId": "1",
			},
			map[string]interface{}{
				"lat": lat,
				"lng": lng,
			},
			time.Now())
		writeAPI.WritePoint(p)
		// writeAPI.WriteRecord(fmt.Sprintf("bafang,type=gps lat=%f,lng=%f", lat, lng))
	}
}

func FindAll() int {
	query := fmt.Sprintf(`
		from(bucket:"%v")
			|> range(start: -1h)
			|> filter(fn: (r) => r._measurement == "gps")
			|> pivot(rowKey:["_time"], columnKey:["_field"], valueColumn:"_value")`, bucket)
	log.Println(query)

	l := 0
	queryAPI := client.QueryAPI(org)
	result, err := queryAPI.Query(context.Background(), query)

	if err == nil {
		data := []gpsData{}

		for result.Next() {
			l++

			// Access data
			v := result.Record().Values()

			lng := v["lng"].(float64)
			lat := v["lat"].(float64)

			var m = gpsData{
				lng: lng,
				lat: lat,
			}

			data = append(data, m)
		}
		fmt.Println(len(data))

		if result.Err() != nil {
			fmt.Printf("query parsing error: %v\n", result.Err().Error())
		}
	} else {
		panic(err)
	}

	return l
}
