INSERT INTO "telemetry_data" (
    "packet_id",
    "packet_seq_ctrl",
    "packet_length",
    "timestamp",
    "subsystem_id",
    "temperature",
    "battery",
    "altitude",
    "signal",
    "has_anomaly"
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;
