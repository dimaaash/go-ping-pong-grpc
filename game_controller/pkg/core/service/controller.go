package controller

import (
	"context"

	"log"

	pingc "github.com/dimaaash/go-ping-pong-grpc/game_controller/pkg/core/ping_client"
	pongc "github.com/dimaaash/go-ping-pong-grpc/game_controller/pkg/core/pong_client"
	pingpb "github.com/dimaaash/go-ping-pong-grpc/protos/gen/ping"
	pongpb "github.com/dimaaash/go-ping-pong-grpc/protos/gen/pong"
)

type ControllerServer struct {
	PingSvc *pingc.PingClient
	PongSvc *pongc.PongClient
}

func (cs *ControllerServer) Ping(ctx context.Context, req *pingpb.PingRequest) (*pingpb.PingResponse, error) {

	res, err := cs.PingSvc.Ping(ctx, req)

	if err != nil {
		log.Fatalf("Failed to ping: %v", err)
	}
	return res, nil
}

func (cs *ControllerServer) Pong(ctx context.Context, req *pongpb.PongRequest) (*pongpb.PongResponse, error) {

	res, err := cs.PongSvc.Pong(ctx, req)

	if err != nil {
		log.Fatalf("Failed to pong: %v", err)
	}
	return res, nil
}
