package interfaces

type ProduceOptions struct {
	// Message to be enqueued
	Message []byte
	// Queue name
	QueueName string
	// Exchange name
	ExchangeName string
	// Routing key
	RoutingKey string
	// Content type
	ContentType string
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

type ConsumeCallback func(string) error

type Queue interface {
	// Push a message to the queue
	Produce(options ProduceOptions) error
	// Pop a message from the queue
	Consume(cb ConsumeCallback, options ConsumeOptions) error
}
