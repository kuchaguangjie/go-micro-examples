package main

import (
	"fmt"

	"github.com/asim/nitro/v3/config"
	"github.com/asim/nitro/v3/config/memory"
	"github.com/asim/nitro/v3/config/source/file"
)

func main() {
	// load the config from a file source
	c, err := memory.NewConfig(config.WithSource(
		file.NewSource(file.WithPath("./config.json")),
	))
	if err != nil {
		fmt.Println(err)
		return
	}

	// define our own host type
	type Host struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
	}

	var host Host

	// read a database host
	v, err := c.Load("hosts", "database")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := v.Scan(&host); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(host.Address, host.Port)
}
