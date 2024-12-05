package service

import (
	"sync"

	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/domain/entity"
)

type CountCrimesAmountPerAreaService struct{}

type CountCrimesAmountPerAreaData map[string]int

const (
	CRIMES_AMOUNT_PER_AREA_OUTPUT_KEY         = "crimes_amount_per_area"
	PROCESS_CRIMES_AMOUNT_PER_AREA_QUEUE_NAME = "process.crimes-amount-per-area"
)

func (s *CountCrimesAmountPerAreaService) Execute(output *sync.Map, record *entity.Record, mu *sync.Mutex) {
	var crimesAmountPerAreaData CountCrimesAmountPerAreaData

	value, ok := output.Load(CRIMES_AMOUNT_PER_AREA_OUTPUT_KEY)

	mu.Lock()

	if !ok {
		crimesAmountPerAreaData = make(CountCrimesAmountPerAreaData)
	} else {
		crimesAmountPerAreaData = value.(CountCrimesAmountPerAreaData)
	}

	crimesAmountPerAreaData[record.AREANAME]++

	output.Store(CRIMES_AMOUNT_PER_AREA_OUTPUT_KEY, crimesAmountPerAreaData)

	mu.Unlock()
}

func NewCountCrimesAmountPerAreaService() Service {
	return &CountCrimesAmountPerAreaService{}
}
