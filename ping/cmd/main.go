package main

import (
	"log"
	"net"

	"github.com/dimaaash/go-ping-pong-grpc/ping/pkg/service"
	pb "github.com/dimaaash/go-ping-pong-grpc/protos/gen/ping"
	"google.golang.org/grpc"
)

func main() {

	log.Println("Ping Server starting...")

	lis, err := net.Listen("tcp", "50003")

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	grpcServer := grpc.NewServer()

	s := service.PingServer{}

	pb.RegisterPingServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

}
