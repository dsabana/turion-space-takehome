package telemetryApi

import (
	"context"
	"fmt"
	"github.com/dsabana/turion-space-takehome/pkg/openapi"
	"github.com/jmoiron/sqlx"
	"log"

	// Package for DB connections.
	_ "embed"
	_ "github.com/lib/pq"
)

// Storage is a struct that contains the sqlx.DB
type Storage struct {
	DB *sqlx.DB
}

// Constants used to create a Postgres connection.
const (
	DatasourceNameFormat = "host=%s port=%d dbname=%s user=%s password=%s search_path=%s sslmode=%s"
	PQDriver             = "postgres"
)

var (
	//go:embed queries/retrieve-anomalies-data-query.sql
	retrieveAnomaliesDataQuery string
	//go:embed queries/retrieve-current-data-query.sql
	retrieveCurrentDataQuery string
	//go:embed queries/retrieve-data-query.sql
	retrieveDataQuery string
)

// NewStorage Function used to create a repository that connects to Postgres
func NewStorage(cfg APIConfiguration) (*Storage, error) {
	db, err := sqlx.Open(PQDriver, fmt.Sprintf(DatasourceNameFormat, cfg.PGHost, cfg.PGPort, cfg.PGDatabase, cfg.PGUser, cfg.PGPassword, cfg.PGSchema, cfg.PGSSLMode))
	if err != nil {
		log.Printf("[ERROR] couldn't open a db connection: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		log.Printf("[ERROR] couldn't ping a db connection: %v", err)
		return nil, err
	}

	return &Storage{
		DB: db,
	}, nil
}

func (s *Storage) RetrieveData(ctx context.Context, startTime string, endTime string) (*[]openapi.TelemetryPacket, error) {
	rows, err := s.DB.QueryxContext(ctx,
		retrieveDataQuery,
		startTime,
		endTime,
	)
	if err != nil {
		log.Printf("error retrieving data")
		return nil, err
	}
	defer rows.Close()

	dataEntries := make([]openapi.TelemetryPacket, 0)
	var t TelemetryFull
	for rows.Next() {
		rows.StructScan(&t)
		dataEntries = append(dataEntries, *mapDBDataToOpenAPI(t))
	}

	return &dataEntries, nil
}

func (s *Storage) RetrieveCurrentData(ctx context.Context) (*openapi.TelemetryPacket, error) {
	var t TelemetryFull
	err := s.DB.QueryRowxContext(ctx, retrieveCurrentDataQuery).StructScan(&t)
	if err != nil {
		log.Printf("error retrieving data")
		return nil, err
	}

	return mapDBDataToOpenAPI(t), nil
}

func (s *Storage) RetrieveAnomaliesData(ctx context.Context, startTime string, endTime string) (*[]openapi.TelemetryPacket, error) {
	rows, err := s.DB.QueryxContext(ctx,
		retrieveAnomaliesDataQuery,
		startTime,
		endTime,
	)
	if err != nil {
		log.Printf("error retrieving data")
		return nil, err
	}
	defer rows.Close()

	dataEntries := make([]openapi.TelemetryPacket, 0)
	var t TelemetryFull
	for rows.Next() {
		rows.StructScan(&t)
		dataEntries = append(dataEntries, *mapDBDataToOpenAPI(t))
	}

	return &dataEntries, nil
}
