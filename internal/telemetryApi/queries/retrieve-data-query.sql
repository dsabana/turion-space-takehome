SELECT *
FROM telemetry_data
WHERE ("timestamp" >= COALESCE($1, "timestamp"))
  AND ("timestamp" <= COALESCE($2, "timestamp"))
ORDER BY "timestamp" DESC
LIMIT 200;
