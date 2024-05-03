package main

import (
	"context"
	"log"

	"github.com/dimaaash/go-ping-pong-grpc/game_controller/pkg/config"
	pingc "github.com/dimaaash/go-ping-pong-grpc/game_controller/pkg/core/ping_client"
	pongc "github.com/dimaaash/go-ping-pong-grpc/game_controller/pkg/core/pong_client"
	pingpb "github.com/dimaaash/go-ping-pong-grpc/protos/gen/ping"
	pongpb "github.com/dimaaash/go-ping-pong-grpc/protos/gen/pong"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	// lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	pingSvc := pingc.InitPingClient(&c)
	pongSvc := pongc.InitPongClient(&c)

	pingSvc.Ping(context.Background(), &pingpb.PingRequest{})

	pongSvc.Pong(context.Background(), &pongpb.PongRequest{})

}
