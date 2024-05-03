package pong_client

import (
	"context"
	"fmt"
	"log"

	"github.com/dimaaash/go-ping-pong-grpc/game_controller/pkg/config"
	pb "github.com/dimaaash/go-ping-pong-grpc/protos/gen/pong"
	"google.golang.org/grpc"
)

type PongClient struct {
	pc pb.PongServiceClient
}

func InitPingClient(c *config.Config) *PongClient {

	log.Println("===> [PongClient]: calling PongService... <===")

	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())

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
