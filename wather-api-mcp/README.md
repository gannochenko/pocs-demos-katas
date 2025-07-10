# Weather API MCP Server

An MCP server that provides current weather data and forecasts using the WeatherAPI.com service.

This is a TypeScript-based MCP server that provides real-time weather information for any location worldwide. It demonstrates core MCP concepts by providing weather data tools that can be used by AI assistants and other MCP clients.

## Features

### Tools

- `get_weather` - Get current weather data for any location
  - Takes a location (city name, coordinates, etc.) as a required parameter
  - Returns comprehensive weather information including:
    - Current temperature (Celsius and Fahrenheit)
    - Feels-like temperature
    - Humidity percentage
    - Wind speed and direction
    - Atmospheric pressure
    - Cloud coverage
    - Visibility
    - Weather conditions

## Prerequisites

1. **WeatherAPI.com API Key**: Sign up for a free account at [WeatherAPI.com](https://www.weatherapi.com/) to get your API key
2. **Configuration File**: Create a YAML configuration file in your home directory

## Setup

1. **Get your API key** from [WeatherAPI.com](https://www.weatherapi.com/) (free tier available)

2. **Create configuration file**:

   ```bash
   cp .weather-api-mcp.yml.example ~/.weather-api-mcp.yml
   ```

3. **Edit the configuration file** `~/.weather-api-mcp.yml` and add your API key:
   ```yaml
   # Weather API MCP Configuration File
   weather_api_key: "your_actual_api_key_here"
   ```

## Development

Install dependencies:

```bash
npm install
```

Build the server:

```bash
npm run build
```

For development with auto-rebuild:

```bash
npm run watch
```

## Installation

To use with Claude Desktop, add the server config:

On MacOS: `~/Library/Application Support/Claude/claude_desktop_config.json`
On Windows: `%APPDATA%/Claude/claude_desktop_config.json`

```json
{
  "mcpServers": {
    "weather-api": {
      "command": "node",
      "args": ["/path/to/weather-api-mcp/build/index.js"]
    }
  }
}
```

### Debugging

Since MCP servers communicate over stdio, debugging can be challenging. We recommend using the [MCP Inspector](https://github.com/modelcontextprotocol/inspector), which is available as a package script:

```bash
npm run inspector
```

The Inspector will provide a URL to access debugging tools in your browser.

## Usage Example

Once installed, you can ask Claude (or any MCP client) questions like:

- "What's the weather like in Tokyo?"
- "Get me the current weather for New York City"
- "How's the weather in latitude 40.7128, longitude -74.0060?"

The server will return formatted weather information with current conditions, temperature, humidity, wind speed, and more.
