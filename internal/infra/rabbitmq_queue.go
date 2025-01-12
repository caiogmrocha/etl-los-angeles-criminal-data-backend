package infra

import (
	"log"

	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/configs"
	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/app/interfaces"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQQueue struct {
	// RabbitMQ connection
	conn *amqp.Connection
}

func (q *RabbitMQQueue) AssertQueue(queueName string) error {
	channel, err := q.conn.Channel()

	if err != nil {
		return err
	}

	defer channel.Close()

	_, err = channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	return nil
}

func (q *RabbitMQQueue) AssertExchange(options interfaces.AssertExchangeOptions) error {
	channel, err := q.conn.Channel()

	if err != nil {
		return err
	}

	defer channel.Close()

	err = channel.ExchangeDeclare(
		options.ExchangeName,
		options.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
	)

	for _, queue := range options.QueuesToBind {
		err = channel.QueueBind(
			queue.QueueName,
			queue.RoutingKey,
			options.ExchangeName,
			false,
			nil,
		)
	}

	if err != nil {
		return err
	}

	return nil
}

func (q *RabbitMQQueue) Produce(options *interfaces.ProduceOptions) error {
	var channel *amqp.Channel

	if options.Channel == nil {
		var err error

		channel, err = q.conn.Channel()

		if err != nil {
			return err
		}

		defer channel.Close()
	} else {
		channel = options.Channel
	}

	err := channel.Publish(
		options.ExchangeName,
		options.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: options.ContentType,
			Body:        options.Message,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (q *RabbitMQQueue) Consume(cb interfaces.ConsumeCallback, options interfaces.ConsumeOptions) error {
	channel, err := q.conn.Channel()

	if err != nil {
		return err
	}

	defer channel.Close()

	msgs, err := channel.Consume(
		options.QueueName,
		"",
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
		err = cb(msg.Body)

		if err != nil {
			log.Printf("Failed to consume message: %s", err)
		}

		msg.Ack(false)
	}

	return nil
}

func NewRabbitMQQueue() *RabbitMQQueue {
	return &RabbitMQQueue{
		conn: configs.AMQP,
	}
}
