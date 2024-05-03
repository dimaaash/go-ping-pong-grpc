package main

import (
	"log"
	"net"

	"github.com/dimaaash/go-ping-pong-grpc/ping/pkg/service"
	pingpb "github.com/dimaaash/go-ping-pong-grpc/protos/gen/ping"
	"google.golang.org/grpc"
)

func main() {

	log.Println("Ping Server starting...")

	lis, err := net.Listen("tcp", "50059")

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	grpcServer := grpc.NewServer()

	s := service.PingServer{}

	pingpb.RegisterPingServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

}
