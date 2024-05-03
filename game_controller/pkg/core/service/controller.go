package controller

import (
	"context"

	"log"

	pingc "github.com/dimaaash/go-ping-pong-grpc/game_controller/pkg/core/ping_client"
	pongc "github.com/dimaaash/go-ping-pong-grpc/game_controller/pkg/core/pong_client"
	pb "github.com/dimaaash/go-ping-pong-grpc/protos/gen/ping"
)

type ControllerServer struct {
	PingSvc *pingc.PingClient
	PongSvc *pongc.PongClient
}

func (cs *ControllerServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {

	res, err := cs.PingSvc.Ping(ctx, req)

	if err != nil {
		log.Fatalf("Failed to ping: %v", err)
	}
	return res, nil
}
