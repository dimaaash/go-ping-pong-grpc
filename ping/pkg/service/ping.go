package service

import (
	pb "github.com/dimaaash/go-ping-pong-grpc/ping/proto_gen"
)

// type PingService interface {
// 	Ping() string
// }

type PingServer struct {
	// pb2.UnimplementedHelloWorldServiceServer
}

func (s *PingServer) Ping() string {

	var a pb.PingRequest

	return "ping!"
}
