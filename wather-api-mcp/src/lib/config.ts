import { readFileSync } from "fs";
import { homedir } from "os";
import { join } from "path";
import YAML from "yaml";
import { z } from "zod";

const configSchema = z.object({
  weather_api_key: z.string().min(1, "weather_api_key cannot be empty"),
});

export type Config = z.infer<typeof configSchema>;

const CONFIG_FILE_PATH = join(homedir(), ".weather-api-mcp.yml");

export function loadConfig(): Config {
  try {
    const fileContent = readFileSync(CONFIG_FILE_PATH, "utf8");
    const rawConfig = YAML.parse(fileContent);

    if (!rawConfig || typeof rawConfig !== "object") {
      throw new Error("Config file must contain a valid YAML object");
    }

    return configSchema.parse(rawConfig);
  } catch (error) {
    if (error instanceof Error) {
      if (error.message.includes("ENOENT")) {
        throw new Error(`Config file not found: ${CONFIG_FILE_PATH}.`);
      }

      if (error.name === "YAMLParseError") {
        throw new Error(
          `Invalid YAML syntax in config file: ${CONFIG_FILE_PATH}. ${error.message}`
        );
      }

      if (error.name === "ZodError") {
        throw new Error(`Invalid config format: ${error.message}`);
      }
    }

    throw error;
  }
}
