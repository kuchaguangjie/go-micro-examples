package main

import (
	"log"

	"github.com/asim/nitro-examples/server/handler"
	"github.com/asim/nitro/v3/server"
	"github.com/asim/nitro/v3/server/mucp"
)

func main() {
	// Initialise Server
	server := mucp.NewServer(
		server.Name("go.micro.srv.example"),
	)

	// Register Handlers
	server.Handle(
		server.NewHandler(
			new(handler.Example),
		),
	)

	// Run server
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	select {}
}
