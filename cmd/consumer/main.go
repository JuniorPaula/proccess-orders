package main

import (
	"encoding/json"
	"fmt"
	"gorabbitmq/internal/config"
	"gorabbitmq/internal/order/infra/database"
	"gorabbitmq/internal/order/usecases"
	"gorabbitmq/pkg/msql"
	"gorabbitmq/pkg/rabbitmq"
	"net/http"
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

	qtdWorkers := 5

	for i := 1; i <= qtdWorkers; i++ {
		go worker(out, &uc, i)
	}

	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		getTotalUC := usecases.GetTotalUseCase{OrderRepository: repository}
		total, err := getTotalUC.Execute()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		json.NewEncoder(w).Encode(total)
	})

	http.ListenAndServe(":3032", nil)

}

func worker(deliveryMessage <-chan amqp.Delivery, uc *usecases.CalculateFinalPriceUsecase, workerID int) {
	for msg := range deliveryMessage {
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

		fmt.Printf("Worker %d has processed order %s\n", workerID, outputDTO.ID)
		time.Sleep(1 * time.Second)
	}
}
