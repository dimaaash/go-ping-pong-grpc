package ping_client

import (
	"context"
	"fmt"
	"log"

	"github.com/dimaaash/go-ping-pong-grpc/game_controller/pkg/config"
	pb "github.com/dimaaash/go-ping-pong-grpc/protos/gen/ping"
	"google.golang.org/grpc"
)

type PingClient struct {
	pc pb.PingServiceClient
}

func InitPingClient(c *config.Config) *PingClient {

	log.Println("===> [PingClient]: calling PingService... <===")

	cc, err := grpc.Dial(c.PingSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	log.Println("===> [PingClient]: PingService OK... <===")

	return &PingClient{
		pc: pb.NewPingServiceClient(cc),
	}
}

func (pc *PingClient) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {

	log.Println("===> [PingClient]: calling Ping method... <===")

	return pc.pc.Ping(ctx, req)
}
