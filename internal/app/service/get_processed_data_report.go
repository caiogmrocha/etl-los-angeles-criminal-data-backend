package service

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type GetProcessedDataReportService struct {
}

type GetProcessedDataReportResponseDTO struct {
	CrimesAmountPerAge    CountCrimesAmountPerAgeData    `json:"crimes_amount_per_age"`
	CrimesAmountPerArea   CountCrimesAmountPerAreaData   `json:"crimes_amount_per_area"`
	CrimesAmountPerPeriod CountCrimesAmountPerPeriodData `json:"crimes_amount_per_period"`
	CrimesAmountPerSex    CountCrimesAmountPerSexData    `json:"crimes_amount_per_sex"`
}

func (g *GetProcessedDataReportService) Execute() (*GetProcessedDataReportResponseDTO, error) {
	reportFilePath := filepath.Join("..", "..", "assets", "output.json")

	reportFile, err := os.Open(reportFilePath)

	if err != nil {
		return nil, err
	}

	defer reportFile.Close()

	report := &GetProcessedDataReportResponseDTO{}

	err = json.NewDecoder(reportFile).Decode(report)

	if err != nil {
		return nil, err
	}

	return report, nil
}

func NewGetProcessedDataReportService() *GetProcessedDataReportService {
	return &GetProcessedDataReportService{}
}
