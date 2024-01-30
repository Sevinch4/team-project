package main

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"teamProject/api"
	"teamProject/config"
	"teamProject/storage/postgres"
)

func main() {
	cfg := config.Load()

	store, err := postgres.New(context.Background(), cfg)
	if err != nil {
		log.Fatalln("error is while connecting to db", err.Error())
		return
	}

	defer store.Close()

	server := api.New(store)

	if err = server.Run("localhost:8080"); err != nil {
		fmt.Println("error is while running", err.Error())
	}
}
