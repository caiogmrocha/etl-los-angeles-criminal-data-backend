package service

import (
	"sync"
	"testing"

	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/domain/entity"
)

func TestCountCrimesAmountPerAreaService_Execute_Success(t *testing.T) {
	// Arrange
	service := NewCountCrimesAmountPerAreaService()

	syncMap := &sync.Map{}

	// Act
	service.Execute(syncMap, recordWithAreaMockFactory("North"))
	service.Execute(syncMap, recordWithAreaMockFactory("South"))
	service.Execute(syncMap, recordWithAreaMockFactory("South"))

	value, ok := syncMap.Load(CRIMES_AMOUNT_PER_AREA_OUTPUT_KEY)

	// Assert
	if !ok {
		t.Errorf("Expected to have a value in key %s of syncMap", CRIMES_AMOUNT_PER_AREA_OUTPUT_KEY)
	}

	crimesAmountPerAreaData := value.(CountCrimesAmountPerAreaData)

	if crimesAmountPerAreaData["North"] != 1 {
		t.Errorf("Expected crimesAmountPerAreaData[\"North\"] to be 1, got %d", crimesAmountPerAreaData["North"])
	}

	if crimesAmountPerAreaData["South"] != 2 {
		t.Errorf("Expected crimesAmountPerAreaData[\"South\"] to be 2, got %d", crimesAmountPerAreaData["South"])
	}
}

func recordWithAreaMockFactory(area string) *entity.Record {
	recordMock := &entity.Record{
		DR_NO:        "123",
		DateRptd:     "2021-01-01",
		DATEOCC:      "2021-01-01",
		TIMEOCC:      "00:00",
		AREA:         "01",
		AREANAME:     area,
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
