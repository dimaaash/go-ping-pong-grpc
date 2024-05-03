package service

import (
	"context"
	"log"

	pb "github.com/dimaaash/go-ping-pong-grpc/protos/gen/pong"
)

// type PongService interface {
// 	Pong() string
// }

type PongServer struct {
	pb.UnimplementedPongServiceServer
}

func (s *PongServer) Pong(ctx context.Context, req *pb.PongRequest) (*pb.PongResponse, error) {

	log.Println("===> PongServer: Pong request received... <===")

	return &pb.PongResponse{
		Status: 0,
		Error:  "",
	}, nil
}
