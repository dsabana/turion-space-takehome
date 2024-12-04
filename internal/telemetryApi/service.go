package telemetryApi

import (
	"context"
	"github.com/dsabana/turion-space-takehome/pkg/openapi"
)

type Service interface {
	GetTelemetryData(ctx context.Context, startTime string, endTime string) (*[]openapi.TelemetryPacket, error)
	GetTelemetryCurrentData(ctx context.Context) (*openapi.TelemetryPacket, error)
	GetTelemetryAnomaliesData(ctx context.Context, startTime string, endTime string) (*[]openapi.TelemetryPacket, error)
}

type Repository interface {
	RetrieveData(ctx context.Context, startTime string, endTime string) (*[]openapi.TelemetryPacket, error)
	RetrieveCurrentData(ctx context.Context) (*openapi.TelemetryPacket, error)
	RetrieveAnomaliesData(ctx context.Context, startTime string, endTime string) (*[]openapi.TelemetryPacket, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{
		r,
	}
}

func (s service) GetTelemetryData(ctx context.Context, startTime string, endTime string) (*[]openapi.TelemetryPacket, error) {
	return s.r.RetrieveData(ctx, startTime, endTime)
}

func (s service) GetTelemetryCurrentData(ctx context.Context) (*openapi.TelemetryPacket, error) {
	return s.r.RetrieveCurrentData(ctx)
}

func (s service) GetTelemetryAnomaliesData(ctx context.Context, startTime string, endTime string) (*[]openapi.TelemetryPacket, error) {
	return s.r.RetrieveAnomaliesData(ctx, startTime, endTime)
}
