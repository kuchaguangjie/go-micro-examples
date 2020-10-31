package main

import (
	"context"
	"log"
	"time"

	"github.com/asim/nitro/v3/service"
	"github.com/asim/nitro/v3/service/mucp"
)

func main() {
	// cancellation context
	ctx, cancel := context.WithCancel(context.Background())

	// shutdown after 5 seconds
	go func() {
		<-time.After(time.Second * 5)
		log.Println("Shutdown example: shutting down service")
		cancel()
	}()

	// create service
	service := mucp.NewService(
		// with our cancellation context
		service.Context(ctx),
	)

	// init service
	service.Init()

	// run service
	service.Run()
}
