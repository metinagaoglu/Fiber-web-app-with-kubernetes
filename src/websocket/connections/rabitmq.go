package amqpclient

import (
	"github.com/streadway/amqp"
)

func Connection() *amqp.Connection {

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial("amqp://rabbitmq:5672/")
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	return connectRabbitMQ
}
