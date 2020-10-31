package main

import (
	"fmt"
	"log"
	"time"

	"github.com/asim/nitro/v3/broker"
	"github.com/asim/nitro/v3/broker/memory"
)

var (
	topic = "go.micro.topic.foo"
	b     = memory.NewBroker()
)

func pub() {
	tick := time.NewTicker(time.Second)
	i := 0
	for _ = range tick.C {
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
		}
		if err := b.Publish(topic, msg); err != nil {
			log.Printf("[pub] failed: %v", err)
		} else {
			fmt.Println("[pub] pubbed message:", string(msg.Body))
		}
		i++
	}
}

func main() {
	if err := b.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}

	if err := b.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	pub()
}
