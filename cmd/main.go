package main

import (
	"fmt"
	"log"
	"teamProject/api"
	"teamProject/config"
	"teamProject/storage/postgres"
)

func main() {
	cfg := config.Load()

	store, err := postgres.New(cfg)
	if err != nil {
		log.Fatalf("error while connecting to db: %v", err)
	}
	defer store.Close()

	server := api.New(store)

	if err := server.Run("localhost:8080"); err != nil {
		fmt.Printf("error while running server: %v\n", err)
	}
}
