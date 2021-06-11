package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	gps "iot.api/pkg/protos/gps"
	server "iot.api/server"
)

func main() {
	port := ":" + os.Getenv("APP_PORT")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Server listening on Port", port)

	grpcServer := grpc.NewServer()
	gps.RegisterGPSServiceServer(grpcServer, &server.S{})

	// run server
	go func ()  {
		if err := grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	} ()

	// graceful shutdown (避免 tcp 佔用資源)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				log.Println("Graceful shutdown start")
				grpcServer.GracefulStop()
				return
		}
	}
}
