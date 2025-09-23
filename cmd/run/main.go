package main

import (
	"context"
	"log"

	"github.com/Oleja123/dcaa-property/internal/config"
	"github.com/Oleja123/dcaa-property/pkg/client/postgresql"
	_ "github.com/lib/pq"
)

func main() {
	config := config.DatabaseConfig{
		Username:         "root",
		Password:         "root",
		Database:         "agency",
		Port:             "5432",
		Host:             "127.0.0.1",
		MaxAttempts:      5,
		SecondsToConnect: 5,
	}

	ctx := context.Background()
	_, err := postgresql.NewClient(ctx, config)
	if err != nil {
		log.Fatal(err)
	}
}
