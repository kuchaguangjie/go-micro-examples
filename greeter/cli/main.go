package main

import (
    "context"
    "fmt"
    "github.com/micro/go-micro/v2/registry"
    "github.com/micro/go-micro/v2/registry/etcd"

    hello "github.com/micro/examples/greeter/srv/proto/hello"
    "github.com/micro/go-micro/v2"
)

func main() {
    // create a new service
    service := micro.NewService(
        micro.Registry(etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))),
    )

    // parse command line flags
    service.Init()

    // Use the generated client stub
    cl := hello.NewSayService("go.micro.srv.greeter", service.Client())

    // Make request
    rsp, err := cl.Hello(context.Background(), &hello.Request{
        Name: "John",
    })
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(rsp.Msg)
}
