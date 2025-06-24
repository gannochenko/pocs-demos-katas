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
    timestamp: getCurrentTimestamp(),
    device_id,
    temperature,
    humidity,
  });

  res
    .status(200)
    .send("Sensor data inserted successfully: " + JSON.stringify(req.body));
});

app.get("/sensor-data-average", async (req, res) => {
  const { device_id } = req.query;
  const resultSet = await clickhouse.query({
    query: `
      SELECT 
        AVG(temperature) AS average_temperature,
        AVG(humidity) AS average_humidity
      FROM sensor_data
      WHERE device_id = {device_id:String}
        AND timestamp >= toStartOfDay(now())
        AND timestamp <= now()
    `,
    query_params: {
      device_id,
    },
  });

  const rows = await resultSet.json();

  console.log("sensor data average", rows);

  res.status(200).send(rows);
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

function getCurrentTimestamp() {
  const now = new Date();
  return (
    now.getFullYear() +
    "-" +
    String(now.getMonth() + 1).padStart(2, "0") +
    "-" +
    String(now.getDate()).padStart(2, "0") +
    " " +
    String(now.getHours()).padStart(2, "0") +
    ":" +
    String(now.getMinutes()).padStart(2, "0") +
    ":" +
    String(now.getSeconds()).padStart(2, "0")
  );
}
