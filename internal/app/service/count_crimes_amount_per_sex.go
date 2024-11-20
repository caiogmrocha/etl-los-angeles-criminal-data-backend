package service

import (
	"sync"

	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/domain/entity"
)

type CountCrimesAmountPerSexService struct{}

type CountCrimesAmountPerSexData struct {
	Male      int `json:"male"`
	Female    int `json:"female"`
	Unknown   int `json:"unknown"`
	NonBinary int `json:"non_binary"`
}

const (
	CRIMES_AMOUNT_PER_SEX_OUTPUT_KEY         = "crimes_amount_per_sex"
	PROCESS_CRIMES_AMOUNT_PER_SEX_QUEUE_NAME = "process.crimes-amount-per-sex"
)

func (s *CountCrimesAmountPerSexService) Execute(record *entity.Record, output *sync.Map) {
	var crimesAmountPerSexData *CountCrimesAmountPerSexData

	value, ok := output.Load(CRIMES_AMOUNT_PER_AGE_OUTPUT_KEY)

	if !ok {
		crimesAmountPerSexData = &CountCrimesAmountPerSexData{
			Male:      0,
			Female:    0,
			Unknown:   0,
			NonBinary: 0,
		}
	} else {
		crimesAmountPerSexData = value.(*CountCrimesAmountPerSexData)
	}

	switch record.VictSex {
	case "M":
		crimesAmountPerSexData.Male++
	case "F":
		crimesAmountPerSexData.Female++
	case "X":
		crimesAmountPerSexData.NonBinary++
	default:
		crimesAmountPerSexData.Unknown++
	}

	output.Store(CRIMES_AMOUNT_PER_AGE_OUTPUT_KEY, crimesAmountPerSexData)
}

func NewCountCrimesAmountPerSexService() *CountCrimesAmountPerSexService {
	return &CountCrimesAmountPerSexService{}
}
