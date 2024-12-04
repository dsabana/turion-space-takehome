package telemetryApi

import "github.com/dsabana/turion-space-takehome/pkg/openapi"

func mapDBDataToOpenAPI(dbData TelemetryFull) *openapi.TelemetryPacket {
	return &openapi.TelemetryPacket{
		PrimaryHeader: openapi.PrimaryHeader{
			PacketId:      &dbData.PacketId,
			PacketLength:  &dbData.PacketLength,
			PacketSeqCtrl: &dbData.PacketSeqCtrl,
		},
		SecondaryHeader: openapi.SecondaryHeader{
			SubsystemId: &dbData.SubsystemId,
			Timestamp:   dbData.Timestamp,
		},
		Payload: openapi.TelemetryPayload{
			Altitude:    dbData.Altitude,
			Battery:     dbData.Battery,
			Signal:      dbData.Signal,
			Temperature: dbData.Temperature,
		},
		HasAnomaly: dbData.HasAnomaly,
	}
}
