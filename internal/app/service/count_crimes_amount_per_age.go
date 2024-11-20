package service

import (
	"strconv"
	"sync"

	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/domain/entity"
)

type CountCrimesAmountPerAgeService struct{}

type CountCrimesAmountPerAgeData struct {
	From0To9   int `json:"from_0_to_9"`
	From10To19 int `json:"from_10_to_19"`
	From20To29 int `json:"from_20_to_29"`
	From30To39 int `json:"from_30_to_39"`
	From40To49 int `json:"from_40_to_49"`
	From50To59 int `json:"from_50_to_59"`
	From60To69 int `json:"from_60_to_69"`
	From70To79 int `json:"from_70_to_79"`
	From80To89 int `json:"from_80_to_89"`
	From90To99 int `json:"from_90_to_99"`
	Unknown    int `json:"unknown"`
}

const ()

const (
	CRIMES_AMOUNT_PER_AGE_OUTPUT_KEY         = "crimes_amount_per_age"
	PROCESS_CRIMES_AMOUNT_PER_AGE_QUEUE_NAME = "process.crimes-amount-per-age"
)

func (s *CountCrimesAmountPerAgeService) Execute(record *entity.Record, output *sync.Map) {
	var crimesAmountPerAgeData *CountCrimesAmountPerAgeData

	value, ok := output.Load(CRIMES_AMOUNT_PER_AGE_OUTPUT_KEY)

	if !ok {
		crimesAmountPerAgeData = &CountCrimesAmountPerAgeData{
			From0To9:   0,
			From10To19: 0,
			From20To29: 0,
			From30To39: 0,
			From40To49: 0,
			From50To59: 0,
			From60To69: 0,
			From70To79: 0,
			From80To89: 0,
			From90To99: 0,
			Unknown:    0,
		}
	} else {
		crimesAmountPerAgeData = value.(*CountCrimesAmountPerAgeData)
	}

	age, err := strconv.Atoi(record.VictAge)

	if err != nil {
		crimesAmountPerAgeData.Unknown++

		return
	}

	switch {
	case age >= 0 && age <= 9:
		crimesAmountPerAgeData.From0To9++
	case age >= 10 && age <= 19:
		crimesAmountPerAgeData.From10To19++
	case age >= 20 && age <= 29:
		crimesAmountPerAgeData.From20To29++
	case age >= 30 && age <= 39:
		crimesAmountPerAgeData.From30To39++
	case age >= 40 && age <= 49:
		crimesAmountPerAgeData.From40To49++
	case age >= 50 && age <= 59:
		crimesAmountPerAgeData.From50To59++
	case age >= 60 && age <= 69:
		crimesAmountPerAgeData.From60To69++
	case age >= 70 && age <= 79:
		crimesAmountPerAgeData.From70To79++
	case age >= 80 && age <= 89:
		crimesAmountPerAgeData.From80To89++
	case age >= 90 && age <= 99:
		crimesAmountPerAgeData.From90To99++
	default:
		crimesAmountPerAgeData.Unknown++
	}

	output.Store(CRIMES_AMOUNT_PER_AGE_OUTPUT_KEY, crimesAmountPerAgeData)
}

func NewCountCrimesAmountPerAgeService() *CountCrimesAmountPerAgeService {
	return &CountCrimesAmountPerAgeService{}
}
