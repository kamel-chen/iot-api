package main

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	gps "iot.api/pkg/protos/gps"
)


func main() {
	// 192.168.0.161, 192.168.0.167, 192.168.0.151
	// addr := "192.168.0.167:" + os.Getenv("APP_PORT")
	// addr := "192.168.1.4:8080"
	addr := "localhost:8080"
	fmt.Println(addr)

	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Can not connect to gRPC server: %v", err)
	}
	defer conn.Close()

	c := gps.NewGPSServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.CreateGPS(ctx, &gps.CreateRequest{DeviceId: 1, Lat: 24.1724173, Lng: 120.6775856, Direction: 10.000001, Speed: 40.21})
	if err != nil {
		log.Fatalf("Can not create gps: %v", err)
	}

	fmt.Println("Response:", r.Success)
}
