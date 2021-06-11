package server

import (
	"context"
	"time"

	gpsservice "iot.api/internal/services/gps"
	gps "iot.api/pkg/protos/gps"
)

type S struct {
	gps.UnimplementedGPSServiceServer
}

/**
 * used for Postgres
 */
func (s S) CreateGPS(ctx context.Context, in *gps.CreateRequest) (*gps.CreateResponse, error) {
	// log.Println("Received: ", in)
	resp := &gps.CreateResponse{Success: true}

	err := gpsservice.PG_create(
		gpsservice.GPS{
			Time: time.Now(),
			DeviceId: in.DeviceId,
			Lat: in.Lat,
			Lng: in.Lng,
			Direction: in.Direction,
			Speed: in.Speed,
		})
	if err != nil {
		resp.Success = false
		return resp, err
	}

	return resp, nil
}

/**
 * used for influx
 */
// func (s S) CreateGPS(ctx context.Context, in *gps.CreateRequest) (*gps.CreateResponse, error) {
// 	// log.Println("Received: ", in)
// 	resp := &gps.CreateResponse{Success: true}

// 	err := gpsservice.AsyncWrite(in.DeviceId, in.Lat, in.Lng, in.Direction, in.Speed)
// 	if err != nil {
// 		resp.Success = false
// 		return resp, err
// 	}

// 	return resp, nil
// }

