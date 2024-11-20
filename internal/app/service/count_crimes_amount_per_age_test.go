package service

import (
	"strconv"
	"sync"
	"testing"

	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/domain/entity"
)

func TestCountCrimesAmountPerAgeService_Execute_Success(t *testing.T) {
	// Arrange
	service := NewCountCrimesAmountPerAgeService()

	syncMap := &sync.Map{}

	// Act
	service.Execute(syncMap, recordWithAgeMockFactory(20))
	service.Execute(syncMap, recordWithAgeMockFactory(30))
	service.Execute(syncMap, recordWithAgeMockFactory(30))

	value, ok := syncMap.Load(CRIMES_AMOUNT_PER_AGE_OUTPUT_KEY)

	// Assert
	if !ok {
		t.Errorf("Expected to have a value in key %s of syncMap", CRIMES_AMOUNT_PER_AGE_OUTPUT_KEY)
	}

	crimesAmountPerAgeData := value.(*CountCrimesAmountPerAgeData)

	if crimesAmountPerAgeData.From0To9 != 0 {
		t.Errorf("Expected crimesAmountPerAgeData.From0To9 to be 0, got %d", crimesAmountPerAgeData.From0To9)
	}

	if crimesAmountPerAgeData.From10To19 != 0 {
		t.Errorf("Expected crimesAmountPerAgeData.From10To19 to be 0, got %d", crimesAmountPerAgeData.From10To19)
	}

	if crimesAmountPerAgeData.From20To29 != 1 {
		t.Errorf("Expected crimesAmountPerAgeData.From20To29 to be 1, got %d", crimesAmountPerAgeData.From20To29)
	}

	if crimesAmountPerAgeData.From30To39 != 2 {
		t.Errorf("Expected crimesAmountPerAgeData.From30To39 to be 2, got %d", crimesAmountPerAgeData.From30To39)
	}

	if crimesAmountPerAgeData.From40To49 != 0 {
		t.Errorf("Expected crimesAmountPerAgeData.From40To49 to be 0, got %d", crimesAmountPerAgeData.From40To49)
	}

	if crimesAmountPerAgeData.From50To59 != 0 {
		t.Errorf("Expected crimesAmountPerAgeData.From50To59 to be 0, got %d", crimesAmountPerAgeData.From50To59)
	}

	if crimesAmountPerAgeData.From60To69 != 0 {
		t.Errorf("Expected crimesAmountPerAgeData.From60To69 to be 0, got %d", crimesAmountPerAgeData.From60To69)
	}

	if crimesAmountPerAgeData.From70To79 != 0 {
		t.Errorf("Expected crimesAmountPerAgeData.From70To79 to be 0, got %d", crimesAmountPerAgeData.From70To79)
	}

	if crimesAmountPerAgeData.From80To89 != 0 {
		t.Errorf("Expected crimesAmountPerAgeData.From80To89 to be 0, got %d", crimesAmountPerAgeData.From80To89)
	}

	if crimesAmountPerAgeData.From90To99 != 0 {
		t.Errorf("Expected crimesAmountPerAgeData.From90To99 to be 0, got %d", crimesAmountPerAgeData.From90To99)
	}

	if crimesAmountPerAgeData.Unknown != 0 {
		t.Errorf("Expected crimesAmountPerAgeData.Unknown to be 0, got %d", crimesAmountPerAgeData.Unknown)
	}
}

func recordWithAgeMockFactory(age int) *entity.Record {
	recordMock := &entity.Record{
		DR_NO:        "123",
		DateRptd:     "2021-01-01",
		DATEOCC:      "2021-01-01",
		TIMEOCC:      "00:00",
		AREA:         "01",
		AREANAME:     "Central",
		RptDistNo:    "0100",
		Part12:       "12",
		CrmCd:        "123",
		CrmCdDesc:    "Theft",
		Mocodes:      "123",
		VictAge:      strconv.Itoa(age),
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
