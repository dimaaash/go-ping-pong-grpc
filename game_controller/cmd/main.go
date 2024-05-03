package main

import (
	"fmt"
	"log"
	"net"


	pc "github.com/dimaaash/go-ping-pong-grpc/game_controller/pkg/core/ping_client"
)
	"github.com/dimaaash/go-ping-pong-grpc/game_controller/pkg/config"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	pingSvc := pc.InitProductServiceClient(c.ProductSvcUrl)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Order Svc on", c.Port)

}
