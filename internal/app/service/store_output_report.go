package service

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/pkg/utils"
)

type StoreOutputReportService struct{}

func (s *StoreOutputReportService) Execute(outputDataMap *sync.Map, outputPath string) {

	//if file exists, open it, if not, create it
	_, err := os.Stat(outputPath)

	if err == nil {
		err = os.Remove(outputPath)

		if err != nil {
			utils.FailOnError(err, "Failed to remove output file")
		}
	}

	file, err := os.Create(outputPath)

	if err != nil {
		utils.FailOnError(err, "Failed to create output file")
	}

	defer file.Close()

	outputDataMapStd := map[string]interface{}{}

	outputDataMap.Range(func(key, value interface{}) bool {
		outputDataMapStd[key.(string)] = value
		return true
	})

	outputDataMapJSON, err := json.Marshal(outputDataMapStd)

	if err != nil {
		utils.FailOnError(err, "Failed to marshal output data map to JSON")
	}

	_, err = file.WriteString(string(outputDataMapJSON))

	if err != nil {
		utils.FailOnError(err, "Failed to write to output file")
	}
}

func NewStoreOutputReportService() *StoreOutputReportService {
	return &StoreOutputReportService{}
}
