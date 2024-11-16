package main

import (
	"log"
	"os"

	_ "github.com/caiogmrocha/etl-los-angeles-criminal-data-backend/configs"
)

func main() {
	log.Printf("%s\n", os.Getenv("HELLO_WORLD"))
}
