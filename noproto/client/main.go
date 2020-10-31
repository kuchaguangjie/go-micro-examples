package main

import (
	"context"
	"fmt"

	"github.com/asim/nitro/v3/client"
	"github.com/asim/nitro/v3/client/mucp"
)

func main() {
	c := mucp.NewClient()

	request := c.NewRequest("greeter", "Greeter.Hello", "john", client.WithContentType("application/json"))
	var response string

	if err := c.Call(context.TODO(), request, &response); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
}
