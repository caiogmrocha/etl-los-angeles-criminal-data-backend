package service

import (
	"sync"

	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/domain/entity"
)

type CountCrimesAmountPerSexService struct{}

func (s *CountCrimesAmountPerSexService) Execute(record *entity.Record, output *sync.Map) {
	if record.VictSex == "" {
		return
	}

	if record.VictSex == "" {
		record.VictSex = "U"
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
