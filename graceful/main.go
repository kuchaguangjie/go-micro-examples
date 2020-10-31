package main

import (
	"github.com/asim/nitro/v3/server"
	"github.com/asim/nitro/v3/server/mucp"
)

func main() {
	srv := mucp.NewServer(
		server.Wait(nil),
	)

	srv.Start()

	select {}
}
