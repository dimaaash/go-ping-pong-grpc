package service

import (
	pb "github.com/dimaaash/go-ping-pong-grpc/protos/ping/pb"
)

// type PingService interface {
// 	Ping() string
// }

type PingServer struct {
	pb.UnimplementedHelloWorldServiceServer
}

func (s *PingServer) Ping() string {

	return "ping!"
}
