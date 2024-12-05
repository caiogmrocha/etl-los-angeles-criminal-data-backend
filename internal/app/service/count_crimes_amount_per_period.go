package service

import (
	"strconv"
	"sync"
	"time"

	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/domain/entity"
	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/pkg/utils"
)

type CountCrimesAmountPerPeriodService struct{}

type CountCrimesAmountPerPeriodDataYear map[string]int

type CountCrimesAmountPerPeriodData map[string](CountCrimesAmountPerPeriodDataYear)

const (
	CRIMES_AMOUNT_PER_PERIOD_OUTPUT_KEY         = "crimes_amount_per_period"
	PROCESS_CRIMES_AMOUNT_PER_PERIOD_QUEUE_NAME = "process.crimes-amount-per-period"
)

func (s *CountCrimesAmountPerPeriodService) Execute(output *sync.Map, record *entity.Record, mu *sync.Mutex) {
	var crimesAmountPerPeriodData *CountCrimesAmountPerPeriodData

	value, ok := output.Load(CRIMES_AMOUNT_PER_PERIOD_OUTPUT_KEY)

	mu.Lock()

	if !ok {
		crimesAmountPerPeriodData = &CountCrimesAmountPerPeriodData{}

		for i := 2020; i <= 2024; i++ {
			(*crimesAmountPerPeriodData)[strconv.Itoa(i)] = CountCrimesAmountPerPeriodDataYear{}

			for j := 1; j <= 12; j++ {
				(*crimesAmountPerPeriodData)[strconv.Itoa(i)][strconv.Itoa(j)] = 0
			}
		}
	} else {
		crimesAmountPerPeriodData = value.(*CountCrimesAmountPerPeriodData)
	}

	period, err := time.Parse("01/02/2006 03:04:05 PM", record.DATEOCC)

	if err != nil {
		period, err = time.Parse("01/02/2006 03:04:05", record.DATEOCC)

		if err != nil {
			utils.FailOnError(err, "Error parsing date")
		}
	}

	year := period.Year()
	month := int(period.Month())

	switch {
	case year == 2020:
		(*crimesAmountPerPeriodData)["2020"][strconv.Itoa(month)]++
	case year == 2021:
		(*crimesAmountPerPeriodData)["2021"][strconv.Itoa(month)]++
	case year == 2022:
		(*crimesAmountPerPeriodData)["2022"][strconv.Itoa(month)]++
	case year == 2023:
		(*crimesAmountPerPeriodData)["2023"][strconv.Itoa(month)]++
	case year == 2024:
		(*crimesAmountPerPeriodData)["2024"][strconv.Itoa(month)]++
	}

	output.Store(CRIMES_AMOUNT_PER_PERIOD_OUTPUT_KEY, crimesAmountPerPeriodData)

	mu.Unlock()
}

func NewCountCrimesAmountPerPeriodService() Service {
	return &CountCrimesAmountPerPeriodService{}
}
