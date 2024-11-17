package main

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/configs"
	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/app/interfaces"
	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/app/service"
	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/domain/entity"
	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/infra"
)

func main() {
	defer configs.Close()

	wg := sync.WaitGroup{}

	rabbitmqQueue := infra.NewRabbitMQQueue()
	produceProcessingTasksService := service.NewProduceProcessingTasksService(rabbitmqQueue)

	recordsTotal := 0

	wg.Add(1)
	go func() {
		produceProcessingTasksService.Execute(context.Background(), "../../assets/crime_data_from_2020_to_2024_los_angeles_minified.csv", &recordsTotal)

		wg.Done()
	}()

	outputDataMap := &sync.Map{}

	countCrimesAmountPerSexServiceExecutions := 0

	wg.Add(1)
	go rabbitmqQueue.Consume(func(message []byte) error {
		countCrimesAmountPerSexService := service.NewCountCrimesAmountPerSexService()

		var record entity.Record

		err := json.Unmarshal(message, &record)

		if err != nil {
			return err
		}

		countCrimesAmountPerSexService.Execute(&record, outputDataMap)

		countCrimesAmountPerSexServiceExecutions++

		if recordsTotal == countCrimesAmountPerSexServiceExecutions {
			wg.Done()
		}

		return nil
	}, interfaces.ConsumeOptions{
		QueueName:    service.PROCESS_CRIMES_AMOUNT_PER_SEX_QUEUE_NAME,
		ExchangeName: "",
		RoutingKey:   service.PROCESS_CRIMES_AMOUNT_PER_SEX_QUEUE_NAME,
		ContentType:  "application/json",
	})

	log.Println("Press CTRL+C to stop the service")

	wg.Wait()

	outputDataMap.Range(func(key, value any) bool {
		log.Printf("%s %+v\n", key, value)

		return true
	})
}
