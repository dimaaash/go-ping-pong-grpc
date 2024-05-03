package main

import (
	"fmt"
	"log"
	"net"

	"github.com/dimaaash/go-ping-pong-grpc/game_controller/pkg/config"
	pingc "github.com/dimaaash/go-ping-pong-grpc/game_controller/pkg/core/ping_client"
	cs "github.com/dimaaash/go-ping-pong-grpc/game_controller/pkg/core/service"
	pb "github.com/dimaaash/go-ping-pong-grpc/protos/gen/ping"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	pingSvc := pingc.InitPingClient(&c)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Order Svc on", c.Port)

	s := cs.ControllerServer{
		PingSvc: pingSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterPingServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

}
