package main

import (
	"github.com/dsabana/turion-space-takehome/internal/telemetryIngest"
	"log"
)

func main() {
	// Setup configuration
	telemetryIngest.LoadConfig(".")
	cfg := telemetryIngest.ClientConfig

	// Create new telemetry client
	client, err := telemetryIngest.NewTelemetryIngestClient(cfg)
	if err != nil {
		panic(err)
	}

	// Start listening for packets
	log.Println("[INFO] Telemetry Ingest client running...")
	client.ListenAndServe()
}
