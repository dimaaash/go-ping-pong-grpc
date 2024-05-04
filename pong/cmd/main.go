package main

import (
	"log"
	"net"

	"github.com/dimaaash/go-ping-pong-grpc/pong/pkg/service"
	pongpb "github.com/dimaaash/go-ping-pong-grpc/protos/gen/pong"
	"google.golang.org/grpc"
)

func main() {

	log.Println("Pong Server starting...")

	lis, err := net.Listen("tcp", ":50060")

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	grpcServer := grpc.NewServer()

	s := service.PongServer{}

	pongpb.RegisterPongServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
