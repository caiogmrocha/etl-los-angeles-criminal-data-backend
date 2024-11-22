package service

import (
	"sync"
	"testing"

	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/domain/entity"
)

func TestCountCrimesAmountPerPeriodService_Execute_Success(t *testing.T) {
	// Arrange
	service := NewCountCrimesAmountPerPeriodService()

	syncMap := &sync.Map{}

	// Act
	service.Execute(syncMap, recordWithPeriodMockFactory("01/01/2021 00:00:00 AM"))
	service.Execute(syncMap, recordWithPeriodMockFactory("01/01/2021 00:00:00 AM"))
	service.Execute(syncMap, recordWithPeriodMockFactory("01/01/2022 00:00:00 AM"))

	value, ok := syncMap.Load(CRIMES_AMOUNT_PER_PERIOD_OUTPUT_KEY)

	// Assert
	if !ok {
		t.Errorf("Expected to have a value in key %s of syncMap", CRIMES_AMOUNT_PER_PERIOD_OUTPUT_KEY)
	}

	crimesAmountPerPeriodData := value.(*CountCrimesAmountPerPeriodData)

	assertionsMap := map[string]map[string]int{
		"2020": {
			"1": 0, "2": 0, "3": 0, "4": 0, "5": 0, "6": 0, "7": 0, "8": 0, "9": 0, "10": 0, "11": 0, "12": 0,
		},
		"2021": {
			"1": 2, "2": 0, "3": 0, "4": 0, "5": 0, "6": 0, "7": 0, "8": 0, "9": 0, "10": 0, "11": 0, "12": 0,
		},
		"2022": {
			"1": 1, "2": 0, "3": 0, "4": 0, "5": 0, "6": 0, "7": 0, "8": 0, "9": 0, "10": 0, "11": 0, "12": 0,
		},
		"2023": {
			"1": 0, "2": 0, "3": 0, "4": 0, "5": 0, "6": 0, "7": 0, "8": 0, "9": 0, "10": 0, "11": 0, "12": 0,
		},
		"2024": {
			"1": 0, "2": 0, "3": 0, "4": 0, "5": 0, "6": 0, "7": 0, "8": 0, "9": 0, "10": 0, "11": 0, "12": 0,
		},
	}

	for year, months := range *crimesAmountPerPeriodData {
		for month, amount := range months {
			expectedAmount := assertionsMap[year][month]

			if amount != expectedAmount {
				t.Errorf("Expected crimesAmountPerPeriodData[%s][%s] to be %d, got %d", year, month, expectedAmount, amount)
			}
		}
	}
}

func recordWithPeriodMockFactory(period string) *entity.Record {
	recordMock := &entity.Record{
		DR_NO:        "123",
		DateRptd:     "2021-01-01",
		DATEOCC:      period,
		TIMEOCC:      "00:00",
		AREA:         "01",
		AREANAME:     "Central",
		RptDistNo:    "0100",
		Part12:       "12",
		CrmCd:        "123",
		CrmCdDesc:    "Theft",
		Mocodes:      "123",
		VictAge:      "30",
		VictSex:      "M",
		VictDescent:  "H",
		PremisCd:     "123",
		PremisDesc:   "House",
		WeaponUsedCd: "123",
		WeaponDesc:   "Gun",
		Status:       "123",
		StatusDesc:   "123",
		CrmCd1:       "123",
		CrmCd2:       "123",
		CrmCd3:       "123",
		CrmCd4:       "123",
		LOCATION:     "123",
		CrossStreet:  "123",
		LAT:          "123",
		LON:          "123",
	}

	return recordMock
}
