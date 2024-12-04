package main

import (
	"context"
	"errors"
	"github.com/dsabana/turion-space-takehome/internal/telemetryApi"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	telemetryApi.LoadConfig(".")
	cfg := telemetryApi.APIConfig

	storage, err := telemetryApi.NewStorage(cfg)
	if err != nil {
		panic(err)
	}

	service := telemetryApi.NewService(storage)

	router, err := telemetryApi.SetupRouter(service)
	if err != nil {
		panic(err)
	}

	log.Printf("Starting...")

	gracefulShutdownChan := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdownChan, syscall.SIGINT, syscall.SIGTERM)

	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: telemetryApi.RegisterCorsHandler(router),
	}

	go func() {
		log.Printf("Running!")
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}

			close(gracefulShutdownChan)
		}
	}()

	<-gracefulShutdownChan

	log.Printf("Shutting down")

	if err := server.Shutdown(context.Background()); err != nil {
		panic(err)
	}

	log.Println("Stopped")

	os.Exit(0)
}
