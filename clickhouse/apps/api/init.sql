-- Create the sensor_data table
CREATE TABLE IF NOT EXISTS sensor_data (
  timestamp DateTime,
  device_id String,
  temperature Float32,
  humidity Float32
)
ENGINE = MergeTree()
PARTITION BY toYYYYMM(timestamp)
ORDER BY (device_id, timestamp); 
