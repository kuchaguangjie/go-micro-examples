package main

import (
	"github.com/asim/nitro-examples/helloworld/handler"
	"github.com/asim/nitro/v3/logger"
	"github.com/asim/nitro/v3/service"
	"github.com/asim/nitro/v3/service/mucp"

	pb "github.com/asim/nitro-examples/helloworld/proto"
)

func main() {
	// New Service
	helloworld := mucp.NewService(
		service.Name("helloworld"),
		service.Version("latest"),
	)

	// Initialise service
	helloworld.Init()

	// Register Handler
	pb.RegisterHelloworldHandler(helloworld.Server(), new(handler.Helloworld))

	// Run service
	if err := helloworld.Run(); err != nil {
		logger.Fatal(err)
	}
}
