package main

import (
	configs "github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/configs"
)

func main() {
	defer configs.Close()
}
