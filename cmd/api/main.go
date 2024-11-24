package main

import (
	"log"
	"net/http"

	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/app/service"
	http_controller "github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/presentation"
)

func main() {

	getProcessedDataReportController := http_controller.NewGetProcessedDataReportController(
		service.NewGetProcessedDataReportService(),
	)
	http.HandleFunc("/reports/criminal-data", getProcessedDataReportController.Handle)

	log.Println("Server running on port 8080")

	http.ListenAndServe(":8080", nil)

	log.Println("Server stopped")
}
