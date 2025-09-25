package main

import (
	"context"
	"log"
	"net/http"
	"time"

	categoryservice "github.com/Oleja123/dcaa-property/internal/application/category"
	propertyservice "github.com/Oleja123/dcaa-property/internal/application/property"
	propertyhandler "github.com/Oleja123/dcaa-property/internal/handler/property"
	categoryhttpclient "github.com/Oleja123/dcaa-property/internal/infrastructure/category/http"
	propertydb "github.com/Oleja123/dcaa-property/internal/infrastructure/property/db"
	"github.com/Oleja123/dcaa-property/pkg/client/postgresql"
	configpkg "github.com/Oleja123/dcaa-property/pkg/config"
)

func main() {
	config, err := configpkg.LoadConfig("config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	apiurl, err := configpkg.LoadAPIUrl("apiurl.yaml")

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	client, err := postgresql.NewClient(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	repository := propertydb.NewRepository(client)
	categoryClient := categoryhttpclient.NewClient(apiurl.Url)

	categoryService := categoryservice.NewService(categoryClient)
	propertyService := propertyservice.NewService(repository, categoryService)

	handler := propertyhandler.NewHandler(propertyService)

	mux := http.NewServeMux()

	mux.HandleFunc("/properties", handler.Handle)
	mux.HandleFunc("/properties/{id}", handler.HandleWithId)

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	s.ListenAndServe()
}
