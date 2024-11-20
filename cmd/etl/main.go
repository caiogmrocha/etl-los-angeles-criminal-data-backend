package main

import (
	"context"
	"encoding/json"
	"log"
	"path/filepath"
	"sync"

	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/configs"
	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/app/interfaces"
	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/app/service"
	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/domain/entity"
	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/infra"
)

func main() {
	defer configs.Close()

	rabbitmqQueue := infra.NewRabbitMQQueue()
	produceProcessingTasksService := service.NewProduceProcessingTasksService(rabbitmqQueue)

	recordsTotal := 0

	mutex := &sync.Mutex{}

	go func() {
		databasePath := filepath.Join("..", "..", "assets", "crime_data_from_2020_to_2024_los_angeles_minified.csv")

		mutex.Lock()
		produceProcessingTasksService.Execute(context.Background(), databasePath, &recordsTotal)
		mutex.Unlock()
	}()

	outputDataMap := &sync.Map{}

	type ConsumersMap map[string](*struct {
		ServiceConstructor func() service.Service
		Counter            int
		DoneChannel        chan bool
	})

	consumersMap := ConsumersMap{
		service.PROCESS_CRIMES_AMOUNT_PER_AGE_QUEUE_NAME: {
			ServiceConstructor: service.NewCountCrimesAmountPerAgeService,
			Counter:            0,
			DoneChannel:        make(chan bool),
		},
		service.PROCESS_CRIMES_AMOUNT_PER_SEX_QUEUE_NAME: {
			ServiceConstructor: service.NewCountCrimesAmountPerSexService,
			Counter:            0,
			DoneChannel:        make(chan bool),
		},
	}

	for queueName := range consumersMap {
		go rabbitmqQueue.Consume(func(message []byte) error {
			svc := consumersMap[queueName].ServiceConstructor()

			var record entity.Record

			err := json.Unmarshal(message, &record)

			if err != nil {
				return err
			}

			svc.Execute(outputDataMap, &record)

			mutex.Lock()
			consumersMap[queueName].Counter++
			mutex.Unlock()

			if recordsTotal == consumersMap[queueName].Counter {
				consumersMap[queueName].DoneChannel <- true

				return nil
			}

			return nil
		}, interfaces.ConsumeOptions{
			QueueName:    queueName,
			ExchangeName: "",
			RoutingKey:   queueName,
			ContentType:  "application/json",
		})

		log.Printf("Record consumer from %s started", queueName)
	}

	for queueName := range consumersMap {
		<-consumersMap[queueName].DoneChannel

		log.Printf("All records from %s processed", queueName)

		close(consumersMap[queueName].DoneChannel)
	}

	log.Println("Press CTRL+C to stop the service")

	storeOutputReportService := service.NewStoreOutputReportService()
	outputPath := filepath.Join("..", "..", "assets", "output.json")

	storeOutputReportService.Execute(outputDataMap, outputPath)
}
