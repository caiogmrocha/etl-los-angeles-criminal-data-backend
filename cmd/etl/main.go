package main

import (
	"context"
	"log"
	"sync"

	configs "github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/configs"
	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/app/service"
	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/infra"
)

func main() {
	defer configs.Close()

	wg := sync.WaitGroup{}

	rabbitmqQueue := infra.NewRabbitMQQueue()
	produceProcessingTasksService := service.NewProduceProcessingTasksService(rabbitmqQueue)

	go produceProcessingTasksService.Execute(context.Background(), "../../assets/crime_data_from_2020_to_2024_los_angeles_minified.csv")
	wg.Add(1)

	log.Println("Press CTRL+C to stop the service")

	wg.Wait()
}
