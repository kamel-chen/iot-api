package gpsservice

import (
	"log"
	"time"

	database "iot.api/internal/database"
)

type GPS struct {
	Time time.Time
	DeviceId int64
	Lat float64
	Lng float64
	Direction float64
	Speed float64
}

var ctx, db = database.GetPGDb()

func PG_findAll() {
	// db := database.GetPGDb()
	// defer db.Close()

	if db == nil {
		log.Println("db is nil")
		return
	}

	q := `SELECT * FROM gps_matrics`

	_, err := db.Exec(ctx, q)
	if err != nil {
		log.Fatal(err)
	}
	// defer rows.Close()

	// data := []GPS{}
	// for rows.Next() {
	// 	row := GPS{}

	// 	err := rows.Scan(&row.Time, &row.DeviceId, &row.Lat, &row.Lng)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	// data = append(data, row)
	// }

	// if err = rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }
}

func PG_create(params GPS) error {
	// db := database.GetPGDb()
	// defer db.Close()
	
	q := "INSERT INTO gps_matrics(time, device_id, lat, lng, direction, speed) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := db.Exec(ctx, q, params.Time, params.DeviceId, params.Lat, params.Lng, params.Direction, params.Speed)
	
	if err != nil {
		log.Fatal("create errors: ", err)
		return err
	}
	
	// log.Println(row)
	// err = row.Close()
	// if err != nil {
	// 	log.Fatal("close row error: ", err)
	// }

	// log.Println(row)
	return nil
}
