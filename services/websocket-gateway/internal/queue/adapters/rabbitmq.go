package adapters

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"

	"websocket-gateway/pkg/config"
	"websocket-gateway/pkg/logger"
	"websocket-gateway/pkg/utils"
)

var amqConn *amqp.Connection
var channel *amqp.Channel

type RabbitMQAdapter struct {
}

func (r *RabbitMQAdapter) Connect() error {
	c, _ := config.LoadConfig()

	logger.Info(c.RabbitMqUrl, "Connecting to RabbitMQ")
	conn, err := amqp.Dial(c.RabbitMqUrl)
	if err != nil {
		return err
	}
	defer conn.Close()

	amqConn, err := amqp.Dial(c.RabbitMqUrl)
	if err != nil {
		return err
	}

	// Let's start by opening a channel to our RabbitMQ
	// instance over the connection we have already
	// established.
	channelRabbitMQ, err := amqConn.Channel()
	if err != nil {
		return err
	}
	channel = channelRabbitMQ
	return nil
}

func (r *RabbitMQAdapter) QueueDeclare() error {

	_, err := channel.QueueDeclare(
		fmt.Sprintf("node-%s", utils.GetNodeId()), // queue name with nodeid
		true,  // durable
		false, // auto delete
		true,  // exclusive
		false, // no wait
		nil,   // arguments
	)
	if err != nil {
		panic(err)
	}

	return nil
}

func (r *RabbitMQAdapter) Publish(message string) error {

	// Attempt to publish a message to the queue.
	if err := channel.Publish(
		"",              // exchange
		"QueueService1", // queue name
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	); err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQAdapter) Subscribe(callback func(string)) error {

	// Consume message.
	msgs, err := channel.Consume(
		fmt.Sprintf("node-%s", utils.GetNodeId()), // queue
		"",    // consumer
		true,  // auto ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   // args
	)
	if err != nil {
		panic(err)
	}

	for d := range msgs {
		callback(string(d.Body))
	}
	return nil
}
