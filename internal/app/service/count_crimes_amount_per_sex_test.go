package service

import (
	"sync"
	"testing"

	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/domain/entity"
)

func TestCountCrimesAmountPerSexService_Execute_Success(t *testing.T) {
	// Arrange
	service := NewCountCrimesAmountPerSexService()

	syncMap := &sync.Map{}

	// Act
	service.Execute(recordMockFactory("M"), syncMap)
	service.Execute(recordMockFactory("M"), syncMap)
	service.Execute(recordMockFactory("F"), syncMap)

	value, ok := syncMap.Load(CRIMES_AMOUNT_PER_AGE_OUTPUT_KEY)

	// Assert
	if !ok {
		t.Errorf("Expected to have a value in key %s of syncMap", CRIMES_AMOUNT_PER_AGE_OUTPUT_KEY)
	}

	crimesAmountPerSexData := value.(*CountCrimesAmountPerSexData)

	if crimesAmountPerSexData.Male != 2 {
		t.Errorf("Expected crimesAmountPerSex to be 2, got %d", crimesAmountPerSexData.Male)
	}

	if crimesAmountPerSexData.Female != 1 {
		t.Errorf("Expected crimesAmountPerSex to be 1, got %d", crimesAmountPerSexData.Female)
	}

	if crimesAmountPerSexData.NonBinary != 0 {
		t.Errorf("Expected crimesAmountPerSex to be 0, got %d", crimesAmountPerSexData.NonBinary)
	}

	if crimesAmountPerSexData.Unknown != 0 {
		t.Errorf("Expected crimesAmountPerSex to be 0, got %d", crimesAmountPerSexData.Unknown)
	}
}

func recordMockFactory(sex string) *entity.Record {
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
		VictAge:      "30",
		VictSex:      sex,
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
