CREATE TABLE IF NOT EXISTS telemetry_data (
    id SERIAL PRIMARY KEY,
    packet_id NUMERIC NOT NULL,
    packet_seq_ctrl NUMERIC NOT NULL,
    packet_length NUMERIC NOT NULL,
    "timestamp" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    subsystem_id NUMERIC NOT NULL,
    temperature NUMERIC,
    battery NUMERIC,
    altitude NUMERIC,
    signal NUMERIC,
    has_anomaly BOOLEAN
);
