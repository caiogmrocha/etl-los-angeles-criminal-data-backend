package interfaces

import amqp "github.com/rabbitmq/amqp091-go"

type ProduceOptions struct {
	// Message to be enqueued
	Message []byte
	// Queue name
	QueueName string
	// Exchange name
	ExchangeName string
	// Exchange type
	ExchangeType string
	// Routing key
	RoutingKey string
	// Content type
	ContentType string
	// Channel
	Channel *amqp.Channel
}

type ConsumeOptions struct {
	// Queue name
	QueueName string
	// Exchange name
	ExchangeName string
	// Routing key
	RoutingKey string
	// Content type
	ContentType string
}

type AssertExchangeOptionsQueuesToBind struct {
	// Queue name
	QueueName string
	// Routing key
	RoutingKey string
}

type AssertExchangeOptions struct {
	// Exchange name
	ExchangeName string
	// Exchange type
	ExchangeType string
	// Queues to bind
	QueuesToBind []AssertExchangeOptionsQueuesToBind
}

type ConsumeCallback func([]byte) error

type Queue interface {
	// Push a message to the queue
	Produce(options *ProduceOptions) error
	// Pop a message from the queue
	Consume(cb ConsumeCallback, options ConsumeOptions) error
	// Assert the queue
	AssertQueue(queueName string) error
	// Assert the exchange
	AssertExchange(options AssertExchangeOptions) error
}
