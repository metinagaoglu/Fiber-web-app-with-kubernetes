package queue

import (
	adapters "websocket-gateway/internal/queue/adapters"
)

type QueueAdapter interface {
	Connect() error
	QueueDeclare() error
	Publish(message string) error
	Subscribe(ProcessMessage func(message string)) error
}

func NewQueueAdapter() QueueAdapter {
	return &adapters.RabbitMQAdapter{}
}
