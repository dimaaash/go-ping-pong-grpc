package service

import (
	"context"

	"log"

	pb "github.com/dimaaash/go-ping-pong-grpc/protos/gen/ping"
)

// type PingService interface {
// 	Ping() string
// }

type PingServer struct {
	pb.UnimplementedPingServiceServer
}

func (s *PingServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {

	log.Printf("===> [PingServer]: PING request received... %s <===\n", req.In)

	return &pb.PingResponse{
		Status: 0,
		Error:  "",
	}, nil
}
