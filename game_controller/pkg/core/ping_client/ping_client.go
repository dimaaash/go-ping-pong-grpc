package ping_client

import (
	"fmt"

	"github.com/dimaaash/go-ping-pong-grpc/game_controller/pkg/config"
	pb "github.com/dimaaash/go-ping-pong-grpc/protos/gen/ping"
	"google.golang.org/grpc"
)

type PingClient struct {
	pc pb.PingServiceClient
}

func InitPingClient(c *config.Config) *PingClient {

	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &pb.NewPingServiceClient(cc)
}
