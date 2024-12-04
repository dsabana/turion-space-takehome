package telemetryIngest

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

// CCSDSPrimaryHeader (6 bytes)
type CCSDSPrimaryHeader struct {
	PacketID      uint16 // Version(3 bits), Type(1 bit), SecHdrFlag(1 bit), APID(11 bits)
	PacketSeqCtrl uint16 // SeqFlags(2 bits), SeqCount(14 bits)
	PacketLength  uint16 // Total packet length minus 7
}

// CCSDSSecondaryHeader (10 bytes)
type CCSDSSecondaryHeader struct {
	Timestamp   uint64 // Unix timestamp
	SubsystemID uint16 // Identifies the subsystem (e.g., power, thermal)
}

// TelemetryPayload represents the payload
type TelemetryPayload struct {
	Temperature float32 // Temperature in Celsius
	Battery     float32 // Battery percentage
	Altitude    float32 // Altitude in kilometers
	Signal      float32 // Signal strength in dB
}

// TelemetryPacket aggregates the entire packet structure
type TelemetryPacket struct {
	PrimaryHeader   CCSDSPrimaryHeader
	SecondaryHeader CCSDSSecondaryHeader
	Payload         TelemetryPayload
	HasAnomaly      bool
}

// ParseTelemetryPacket reads the raw data from the received packets and sets them in a struct
func ParseTelemetryPacket(data []byte) (*TelemetryPacket, error) {
	reader := bytes.NewReader(data)

	// Parse primary header
	var primaryHeader CCSDSPrimaryHeader
	if err := binary.Read(reader, binary.BigEndian, &primaryHeader); err != nil {
		return nil, fmt.Errorf("failed to parse primary header: %w", err)
	}

	// Parse secondary header
	var secondaryHeader CCSDSSecondaryHeader
	if err := binary.Read(reader, binary.BigEndian, &secondaryHeader); err != nil {
		return nil, fmt.Errorf("failed to parse secondary header: %w", err)
	}

	// Parse payload
	var payload TelemetryPayload
	if err := binary.Read(reader, binary.BigEndian, &payload); err != nil {
		return nil, fmt.Errorf("failed to parse payload: %w", err)
	}

	return &TelemetryPacket{
		PrimaryHeader:   primaryHeader,
		SecondaryHeader: secondaryHeader,
		Payload:         payload,
		HasAnomaly:      false,
	}, nil
}

// ValidatePacket checks the packet for anomalies and logs any findings
func ValidatePacket(packet *TelemetryPacket) {
	if packet.Payload.Temperature > 35 {
		packet.HasAnomaly = true
		log.Printf("[ALERT] temperature anomaly: [PacketSeqCtrl: %v] [Temperature: %v > 35]", packet.PrimaryHeader.PacketSeqCtrl, packet.Payload.Temperature)
	}

	if packet.Payload.Battery < 40 {
		packet.HasAnomaly = true
		log.Printf("[ALERT] battery anomaly: [PacketSeqCtrl: %v] [Battery: %v < 40]", packet.PrimaryHeader.PacketSeqCtrl, packet.Payload.Battery)
	}

	if packet.Payload.Altitude < 400 {
		packet.HasAnomaly = true
		log.Printf("[ALERT] altitude anomaly: [PacketSeqCtrl: %v] [Altitude: %v < 400]", packet.PrimaryHeader.PacketSeqCtrl, packet.Payload.Altitude)
	}

	if packet.Payload.Signal < -80 {
		packet.HasAnomaly = true
		log.Printf("[ALERT] signal strength anomaly: [PacketSeqCtrl: %v] [Signal: %v < -80]", packet.PrimaryHeader.PacketSeqCtrl, packet.Payload.Signal)
	}

	if packet.HasAnomaly {
		log.Printf("[INFO] packet with anomalies: %+v", packet)
	}
}
