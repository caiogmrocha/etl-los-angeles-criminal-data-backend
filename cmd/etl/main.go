package main

import (
	"log"
	"sync"

	configs "github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/configs"
	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/app/interfaces"
	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/infra"
	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/pkg/utils"
)

func main() {
	defer configs.Close()

	queue := infra.NewRabbitMQQueue()

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		log.Println("Producing messages")

		for i := 0; i < 1000; i++ {
			err := queue.Produce(interfaces.ProduceOptions{
				Message:      "Hello, World!",
				QueueName:    "hello",
				ExchangeName: "",
				RoutingKey:   "hello",
				ContentType:  "text/plain",
			})

			utils.FailOnError(err, "Failed to produce message")
		}
	}()

	wg.Add(1)

	go func() {
		err := queue.Consume(func(msg string) error {
			log.Printf("Received message: %s", msg)
			return nil
		}, interfaces.ConsumeOptions{
			QueueName:    "hello",
			ExchangeName: "",
			RoutingKey:   "hello",
			ContentType:  "text/plain",
		})

		utils.FailOnError(err, "Failed to consume message")
	}()

	wg.Wait()

	log.Println("Press CTRL+C to stop consuming messages")
}
