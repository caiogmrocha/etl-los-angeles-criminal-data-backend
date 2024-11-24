package http_controller

import (
	"encoding/json"
	"net/http"

	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/app/service"
)

type GetProcessedDataReportController struct {
	GetProcessedDataReportService *service.GetProcessedDataReportService
}

func (c *GetProcessedDataReportController) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	report, err := c.GetProcessedDataReportService.Execute()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error getting report"}`))
		return
	}

	parsedReport, err := json.Marshal(report)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error parsing report"}`))
		return
	}

	w.Write(parsedReport)
	w.WriteHeader(http.StatusOK)
}

func NewGetProcessedDataReportController(service *service.GetProcessedDataReportService) GetProcessedDataReportController {
	return GetProcessedDataReportController{
		GetProcessedDataReportService: service,
	}
}
