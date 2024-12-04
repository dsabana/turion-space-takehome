SELECT *
FROM telemetry_data
WHERE ("timestamp" >= COALESCE($1, "timestamp"))
  AND ("timestamp" <= COALESCE($2, "timestamp"))
  AND has_anomaly = true;
