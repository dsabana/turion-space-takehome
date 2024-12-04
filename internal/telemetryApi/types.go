package telemetryApi

import (
	"fmt"
	"time"
)

var ErrEncodingResponse = fmt.Errorf("unable to encode request")
var ErrRetrievingObject = fmt.Errorf("unable to retrieve object")

type TelemetryFull struct {
	ID            int        `db:"id"`
	PacketId      int        `db:"packet_id"`
	PacketLength  int        `db:"packet_length"`
	PacketSeqCtrl int        `db:"packet_seq_ctrl"`
	SubsystemId   int        `db:"subsystem_id"`
	Timestamp     *time.Time `db:"timestamp"`
	Altitude      float64    `db:"altitude"`
	Battery       float64    `db:"battery"`
	Signal        float64    `db:"signal"`
	Temperature   float64    `db:"temperature"`
	HasAnomaly    bool       `db:"has_anomaly"`
}
