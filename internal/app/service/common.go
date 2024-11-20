package service

import (
	"sync"

	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/internal/domain/entity"
)

type Service interface {
	// Execute the service
	Execute(output *sync.Map, args *entity.Record)
}
