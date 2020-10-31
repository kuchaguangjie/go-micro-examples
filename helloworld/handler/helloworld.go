package handler

import (
	"context"
	"fmt"

	helloworld "github.com/asim/nitro-examples/helloworld/proto"
)

type Helloworld struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Helloworld) Call(ctx context.Context, req *helloworld.Request, rsp *helloworld.Response) error {
	fmt.Println("Received Helloworld.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
