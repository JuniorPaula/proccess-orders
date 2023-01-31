package main

import (
	"gorabbitmq/internal/config"
	"gorabbitmq/pkg/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	config.InitVariables()

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	out := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, out)

	for msg := range out {
		println(string(msg.Body))
	}
}
