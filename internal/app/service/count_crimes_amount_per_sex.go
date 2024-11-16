package service

import (
	"sync"

	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/domain/entity"
	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/domain/value_objects"
)

type CountCrimesAmountPerSexService struct{}

func (s *CountCrimesAmountPerSexService) Execute(record *entity.Record, output *sync.Map) {
	if record.VictSex == "" {
		return
	}

	if record.VictSex == "" {
		record.VictSex = value_objects.Unknown
	}

	if value, ok := output.Load(record.VictSex); ok {
		output.Store(record.VictSex, value.(int)+1)
		return
	} else {
		output.Store(record.VictSex, 1)
	}
}

func NewCountCrimesAmountPerSexService() *CountCrimesAmountPerSexService {
	return &CountCrimesAmountPerSexService{}
}
