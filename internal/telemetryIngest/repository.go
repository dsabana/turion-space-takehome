package telemetryIngest

import (
	_ "embed"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"time"

	// Package for DB connections.
	_ "github.com/lib/pq"
)

// Repository is a struct that contains the sqlx.DB
type Repository struct {
	DB *sqlx.DB
}

// Constants used to create a Postgres connection.
const (
	DatasourceNameFormat = "host=%s port=%d dbname=%s user=%s password=%s search_path=%s sslmode=%s"
	PQDriver             = "postgres"
)

//go:embed queries/insert-telemetry-entry-query.sql
var insertTelemetryEntryQuery string

// NewRepository Function used to create a repository that connects to Postgres
func NewRepository(cfg ClientConfiguration) (*Repository, error) {
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

	return &Repository{
		DB: db,
	}, nil
}

// SavePacketEntry function used to store a packet to the DB
func (r Repository) SavePacketEntry(packet *TelemetryPacket) {
	_, err := r.DB.Exec(insertTelemetryEntryQuery,
		packet.PrimaryHeader.PacketID,
		packet.PrimaryHeader.PacketSeqCtrl,
		packet.PrimaryHeader.PacketLength,
		time.Unix(int64(packet.SecondaryHeader.Timestamp), 0),
		packet.SecondaryHeader.SubsystemID,
		packet.Payload.Temperature,
		packet.Payload.Battery,
		packet.Payload.Altitude,
		packet.Payload.Signal,
		packet.HasAnomaly,
	)
	if err != nil {
		log.Printf("[ERROR] couldn't save packet data: %v", err)
	}
}
