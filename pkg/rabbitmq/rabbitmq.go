package rabbitmq

import "github.com/rabbitmq/amqp091-go"

func OpenChannel() (*amqp091.Channel, error) {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func Consume(ch *amqp091.Channel, out chan amqp091.Delivery) error {
	msgs, err := ch.Consume(
		"orders",
		"order-consume",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}

	return nil
}
