package main

import (
	"log"
	"net/http"

	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/app/service"
	http_controller "github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/presentation"
)

func main() {
	// CORS
	httpMux := http.NewServeMux()

	getProcessedDataReportController := http_controller.NewGetProcessedDataReportController(
		service.NewGetProcessedDataReportService(),
	)
	httpMux.HandleFunc("/reports/criminal-data", getProcessedDataReportController.Handle)

	log.Println("Server running on port 8080")

	http.ListenAndServe(":8080", func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", http.MethodGet+", "+http.MethodOptions+", "+http.MethodPost+", "+http.MethodPut+", "+http.MethodDelete+", "+http.MethodPatch)
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			h.ServeHTTP(w, r)
		})
	}(httpMux))

	log.Println("Server stopped")
}
