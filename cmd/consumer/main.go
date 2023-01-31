package main

import (
	"encoding/json"
	"fmt"
	"gorabbitmq/internal/config"
	"gorabbitmq/internal/order/infra/database"
	"gorabbitmq/internal/order/usecases"
	"gorabbitmq/pkg/msql"
	"gorabbitmq/pkg/rabbitmq"
	"time"

	_ "github.com/go-sql-driver/mysql"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	config.InitVariables()

	db, err := msql.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := database.NewOrderRepository(db)
	uc := usecases.CalculateFinalPriceUsecase{OrderRepository: repository}

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	out := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, out)

	for msg := range out {
		var inputDTO usecases.OrderInputDTO
		err := json.Unmarshal(msg.Body, &inputDTO)
		if err != nil {
			panic(err)
		}

		outputDTO, err := uc.Execute(inputDTO)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)

		fmt.Println(outputDTO)
		time.Sleep(500 * time.Millisecond)
	}
}
