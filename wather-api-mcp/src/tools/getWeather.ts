import { z } from "zod";
import { loadConfig } from "../lib/config.js";
import { BASE_URL, customFetch } from "../lib/customFetch.js";

const getWeatherDataResultSchema = z.object({
  location: z.object({
    name: z.string(),
    region: z.string(),
    country: z.string(),
    lat: z.number(),
    lon: z.number(),
    tz_id: z.string(),
    localtime_epoch: z.number(),
    localtime: z.string(),
  }),
  current: z.object({
    last_updated_epoch: z.number(),
    last_updated: z.string(),
    temp_c: z.number(),
    temp_f: z.number(),
    is_day: z.number(),
    condition: z.object({
      text: z.string(),
      icon: z.string(),
      code: z.number(),
    }),
    wind_mph: z.number(),
    wind_kph: z.number(),
    wind_degree: z.number(),
    wind_dir: z.string(),
    pressure_mb: z.number(),
    pressure_in: z.number(),
    precip_mm: z.number(),
    precip_in: z.number(),
    humidity: z.number(),
    cloud: z.number(),
    feelslike_c: z.number(),
    feelslike_f: z.number(),
    windchill_c: z.number(),
    windchill_f: z.number(),
    heatindex_c: z.number(),
    heatindex_f: z.number(),
    dewpoint_c: z.number(),
    dewpoint_f: z.number(),
    vis_km: z.number(),
    vis_miles: z.number(),
    uv: z.number(),
    gust_mph: z.number(),
    gust_kph: z.number(),
  }),
});

export async function getWeatherData(
  requestedLocation: string
): Promise<string> {
  const config = loadConfig();
  const API_KEY = config.weather_api_key;

  const url = `${BASE_URL}/current.json?key=${API_KEY}&q=${encodeURIComponent(
    requestedLocation
  )}&aqi=no`;

  const data = await customFetch(url);

  const parsed = getWeatherDataResultSchema.parse(data);
  const location = parsed.location;
  const current = parsed.current;

  return `Weather for ${location.name}, ${location.country} is ${Math.round(
    current.temp_c
  )} degrees Celsius, feels like ${Math.round(
    current.feelslike_c
  )} degrees Celsius, humidity is ${current.humidity}%, the wind speed is ${
    current.wind_kph
  } km/h, the pressure is ${current.pressure_mb} hPa, cloud cover is ${
    current.cloud
  }%, visibility is ${current.vis_km} km, condition: ${current.condition.text}`;
}
