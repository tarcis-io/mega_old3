package main

import (
	"log"

	"mega/internal/config"
)

func main() {
	_, err := config.New()
	if err != nil {
		log.Fatalf("failed to create config: %v", err)
	}
}
