import express from "express";
import { createClient } from "@clickhouse/client";

const clickhouse = createClient({
  url: "http://localhost:8123",
  username: process.env.CLICKHOUSE_USER,
  password: process.env.CLICKHOUSE_PASSWORD,
  database: process.env.CLICKHOUSE_DB,
});

const app = express();
const port = 3000;

app.use(express.json());

app.get("/", (req, res) => {
  res.send("Hello, TypeScript + Express!");
});

app.post("/sensor-data", async (req, res) => {
  const { device_id, temperature, humidity } = req.body;

  // todo: validate data

  await insertSensorData({
    timestamp: prepareTimestamp(new Date()),
    device_id,
    temperature,
    humidity,
  });

  res
    .status(200)
    .send("Sensor data inserted successfully: " + JSON.stringify(req.body));
});

app.get("/sensor-data-average", async (req, res) => {
  const { device_id, timestamp } = req.query;

  if (!device_id) {
    res.status(400).send("device_id is required");
    return;
  }

  // Convert timestamp to Date object and subtract 1 day
  const endTimestamp = timestamp ? new Date(timestamp as string) : new Date();
  const startTimestamp = new Date(endTimestamp.getTime() - 24 * 60 * 60 * 1000); // Subtract 1 day in milliseconds

  const resultSet = await clickhouse.query({
    query: `
      SELECT 
        AVG(temperature) AS average_temperature,
        AVG(humidity) AS average_humidity
      FROM sensor_data
      WHERE device_id = {device_id:String}
        AND timestamp >= {start_timestamp:DateTime}
        AND timestamp <= {end_timestamp:DateTime}
    `,
    query_params: {
      device_id,
      start_timestamp: prepareTimestamp(startTimestamp),
      end_timestamp: prepareTimestamp(endTimestamp),
    },
  });

  const rows = await resultSet.json();

  console.log("sensor data average", rows);

  res.status(200).send(rows);
});

app.get("/sensor-data", async (req, res) => {
  const { device_id, timestamp } = req.query;

  if (!device_id) {
    res.status(400).send("device_id is required");
    return;
  }

  // Convert timestamp to Date object and subtract 1 day
  const endTimestamp = timestamp ? new Date(timestamp as string) : new Date();
  const startTimestamp = new Date(endTimestamp.getTime() - 24 * 60 * 60 * 1000); // Subtract 1 day in milliseconds

  console.log("startTimestamp", prepareTimestamp(startTimestamp));
  console.log("endTimestamp", prepareTimestamp(endTimestamp));
  console.log("device_id", device_id);

  const resultSet = await clickhouse.query({
    query: `
        SELECT 
          *
        FROM sensor_data
        WHERE device_id = {device_id:String}
          AND timestamp >= {start_timestamp:DateTime}
          AND timestamp <= {end_timestamp:DateTime}
      `,
    query_params: {
      device_id,
      start_timestamp: prepareTimestamp(startTimestamp),
      end_timestamp: prepareTimestamp(endTimestamp),
    },
  });

  const result = await resultSet.json();

  console.log("sensor data", result.data, result.rows);

  res.status(200).send(result);
});

app.listen(port, () => {
  console.log(`Server is running at http://localhost:${port}`);
});

/// clickhouse

type SensorData = {
  timestamp: string;
  device_id: string;
  temperature: number;
  humidity: number;
};

async function insertSensorData(data: SensorData) {
  console.log("inserting sensor data", data);
  const result = await clickhouse.insert({
    table: "sensor_data",
    values: [data],
    format: "JSONEachRow",
  });

  console.log("insert result", result);
}

function formatTimestampToDB(timestamp: Date) {
  return timestamp.toISOString().slice(0, 19).replace("T", " ");
}

function prepareTimestamp(date: Date) {
  return (
    date.getUTCFullYear() +
    "-" +
    String(date.getUTCMonth() + 1).padStart(2, "0") +
    "-" +
    String(date.getUTCDate()).padStart(2, "0") +
    " " +
    String(date.getUTCHours()).padStart(2, "0") +
    ":" +
    String(date.getUTCMinutes()).padStart(2, "0") +
    ":" +
    String(date.getUTCSeconds()).padStart(2, "0")
  );
}
