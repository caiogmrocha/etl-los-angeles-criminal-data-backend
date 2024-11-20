package service

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/app/interfaces"
	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/domain/entity"
	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/pkg/utils"
)

type ProduceProcessingTasksService struct {
	queue interfaces.Queue
}

const (
	PROCESS_CRIMES_DATA_ROUTING_KEY   = "process.*"
	PROCESS_CRIMES_DATA_EXCHANGE_NAME = "process-crimes-data"
)

func (s *ProduceProcessingTasksService) Execute(ctx context.Context, databasePath string, recordsTotal *int) {
	file, err := os.Open(databasePath)

	utils.FailOnError(err, "Failed to open database file")

	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Read()

	queuesToBind := []interfaces.AssertExchangeOptionsQueuesToBind{
		{
			QueueName:  PROCESS_CRIMES_AMOUNT_PER_SEX_QUEUE_NAME,
			RoutingKey: PROCESS_CRIMES_DATA_ROUTING_KEY,
		},
		{
			QueueName:  PROCESS_CRIMES_AMOUNT_PER_AGE_QUEUE_NAME,
			RoutingKey: PROCESS_CRIMES_DATA_ROUTING_KEY,
		},
	}

	for _, queue := range queuesToBind {
		err := s.queue.AssertQueue(queue.QueueName)

		utils.FailOnError(err, fmt.Sprintf("error while asserting queue: %s", queue))
	}

	err = s.queue.AssertExchange(interfaces.AssertExchangeOptions{
		ExchangeName: PROCESS_CRIMES_DATA_EXCHANGE_NAME,
		ExchangeType: "topic",
		QueuesToBind: queuesToBind,
	})

	utils.FailOnError(err, fmt.Sprintf("error while asserting exchange: %s", PROCESS_CRIMES_DATA_EXCHANGE_NAME))

	mu := sync.Mutex{}

	for {
		select {
		case <-ctx.Done():
			log.Print("Database .csv reading aborted")
			return
		default:
			row, err := csvReader.Read()

			if err != nil {
				if errors.Is(err, io.EOF) {
					return
				}

				utils.FailOnError(err, "Error while reading rows from database .csv")
			}

			mu.Lock()
			*recordsTotal++
			mu.Unlock()

			record := &entity.Record{
				DR_NO:        row[0],
				DateRptd:     row[1],
				DATEOCC:      row[2],
				TIMEOCC:      row[3],
				AREA:         row[4],
				AREANAME:     row[5],
				RptDistNo:    row[6],
				Part12:       row[7],
				CrmCd:        row[8],
				CrmCdDesc:    row[9],
				Mocodes:      row[10],
				VictAge:      row[11],
				VictSex:      row[12],
				VictDescent:  row[13],
				PremisCd:     row[14],
				PremisDesc:   row[15],
				WeaponUsedCd: row[16],
				WeaponDesc:   row[17],
				Status:       row[18],
				StatusDesc:   row[19],
				CrmCd1:       row[20],
				CrmCd2:       row[21],
				CrmCd3:       row[22],
				CrmCd4:       row[23],
				LOCATION:     row[24],
				CrossStreet:  row[25],
				LAT:          row[26],
				LON:          row[27],
			}

			marshalledRecord, _ := json.Marshal(record)

			err = s.queue.Produce(interfaces.ProduceOptions{
				Message:      marshalledRecord,
				ExchangeName: PROCESS_CRIMES_DATA_EXCHANGE_NAME,
				ExchangeType: "topic",
				RoutingKey:   "process.*",
				ContentType:  "application/json",
			})

			if err != nil {
				utils.FailOnError(err, fmt.Sprintf("error while producing record: %+v", record))
			}
		}
	}
}

func NewProduceProcessingTasksService(
	queue interfaces.Queue,
) *ProduceProcessingTasksService {
	return &ProduceProcessingTasksService{
		queue: queue,
	}
}
