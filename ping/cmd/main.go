package main

import (
	"log"

	"github.com/dimaaash/go-ping-pong-grpc/ping/pkg/service"
	"google.golang.org/grpc"
)

func main() {

	log.Println("Ping Server starting...")

	grpcServer := grpc.NewServer()

	s := service.Server{}

	// pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(50003); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

}
