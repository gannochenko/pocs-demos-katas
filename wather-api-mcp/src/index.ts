#!/usr/bin/env node

import { Server } from "@modelcontextprotocol/sdk/server/index.js";
import { StdioServerTransport } from "@modelcontextprotocol/sdk/server/stdio.js";
import {
  CallToolRequestSchema,
  ListToolsRequestSchema,
} from "@modelcontextprotocol/sdk/types.js";
import { z } from "zod";
import { getWeatherData } from "./tools/getWeather.js";

const server = new Server(
  {
    name: "Weather API MCP Server",
    version: "0.1.0",
  },
  {
    capabilities: {
      tools: {},
    },
  }
);

/**
 * Handler that lists available tools.
 * Exposes a single "get_weather" tool that lets clients get the weather for a given location.
 */
server.setRequestHandler(ListToolsRequestSchema, async () => {
  return {
    tools: [
      {
        name: "get_weather",
        description: "Get the weather for a given location",
        inputSchema: {
          type: "object",
          properties: {
            location: {
              type: "string",
              description: "The location to get the weather for",
            },
          },
          required: ["location"],
        },
      },
    ],
  };
});

/**
 * Handler for the get_weather tool.
 * Gets the weather for a given location.
 */
server.setRequestHandler(CallToolRequestSchema, async (request) => {
  switch (request.params.name) {
    case "get_weather": {
      const locationSchema = z.object({
        location: z.string().min(1, "Location is required"),
      });

      const parseResult = locationSchema.safeParse(request.params.arguments);
      if (!parseResult.success) {
        throw new Error(parseResult.error.errors[0].message);
      }
      const { location } = parseResult.data;

      return {
        content: [
          {
            type: "text",
            text: await getWeatherData(location),
          },
        ],
      };
    }

    default:
      throw new Error("Unknown tool");
  }
});

/**
 * Start the server using stdio transport.
 * This allows the server to communicate via standard input/output streams.
 */
async function main() {
  const transport = new StdioServerTransport();
  await server.connect(transport);
}

main().catch((error) => {
  console.error("Server error:", error);
  process.exit(1);
});
