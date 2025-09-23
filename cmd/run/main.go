package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Oleja123/dcaa-property/internal/config"
	propertyhandler "github.com/Oleja123/dcaa-property/internal/handler/property"
	propertydb "github.com/Oleja123/dcaa-property/internal/repository/property/db"
	propertyservice "github.com/Oleja123/dcaa-property/internal/service/property"
	"github.com/Oleja123/dcaa-property/pkg/client/postgresql"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadConfig("config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	client, err := postgresql.NewClient(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	repository := propertydb.NewRepository(client)

	service := propertyservice.NewService(repository)

	handler := propertyhandler.NewHandler(service)

	mux := http.NewServeMux()

	mux.HandleFunc("/properties", handler.FindAll)
	mux.HandleFunc("/properties/{id}", handler.FindOne)
	mux.HandleFunc("/properties/create", handler.Create)
	mux.HandleFunc("/properties/update", handler.Update)
	mux.HandleFunc("/properties/delete/{id}", handler.Delete)

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	s.ListenAndServe()
}
