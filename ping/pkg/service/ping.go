package service

import (
	pb2 "github.com/dimaaash/go-ping-pong-grpc/protos/gen/ping"
)

// type PingService interface {
// 	Ping() string
// }

type PingServer struct {
	// pb2.UnimplementedHelloWorldServiceServer
}

func (s *PingServer) Ping() string {

	var a pb2.PingRequest

	return "ping!"
}
