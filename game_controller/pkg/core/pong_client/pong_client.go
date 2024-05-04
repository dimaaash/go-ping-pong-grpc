package pong_client

import (
	"context"
	"fmt"
	"log"

	pb "github.com/dimaaash/go-ping-pong-grpc/protos/gen/pong"
	"google.golang.org/grpc"
)

type PongClient struct {
	pc pb.PongServiceClient
}

func InitPongClient() *PongClient {

	log.Println("===> [PongClient]: calling PongService... <===")

	cc, err := grpc.Dial("localhost:50060", grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	log.Println("===> [PongClient]: PongService OK... <===")

	return &PongClient{
		pc: pb.NewPongServiceClient(cc),
	}
}

func (pc *PongClient) Pong(ctx context.Context, req *pb.PongRequest) (*pb.PongResponse, error) {

	log.Println("===> [PongClient]: calling Pong method... <===")

	return pc.pc.Pong(ctx, req)
}
