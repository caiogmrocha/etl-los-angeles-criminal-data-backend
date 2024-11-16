package configs

import (
	"fmt"
	"os"

	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/pkg/utils"

	amqp "github.com/rabbitmq/amqp091-go"
)

var AMQP *amqp.Connection

func ConfigRabbitMQ() {
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASS"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
	)

	conn, err := amqp.Dial(connStr)

	utils.FailOnError(err, "Failed to connect to RabbitMQ")

	AMQP = conn
}
