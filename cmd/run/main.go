package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Oleja123/dcaa-property/internal/config"
	propertydb "github.com/Oleja123/dcaa-property/internal/repository/property/db"
	"github.com/Oleja123/dcaa-property/pkg/client/postgresql"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadConfig("config.yaml")
	fmt.Println(config)

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	client, err := postgresql.NewClient(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	repository := propertydb.NewRepository(client)

	all, err := repository.FindAll(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range all {
		fmt.Println(val)
	}

	val, err := repository.FindOne(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(val)
}
