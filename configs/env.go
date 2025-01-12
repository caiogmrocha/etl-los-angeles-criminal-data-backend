package configs

import (
	"github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/pkg/utils"
	"github.com/joho/godotenv"
)

func ConfigEnv() {
	err := godotenv.Load()

	utils.FailOnError(err, "Error loading .env file")
}
