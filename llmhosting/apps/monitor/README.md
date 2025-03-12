# NestJS System Metrics Exporter

A simple NestJS application that collects system metrics (CPU, Memory, GPU) and exposes them in Prometheus format.

## Features

- Collects CPU utilization
- Collects memory usage
- Collects GPU utilization (when available)
- Exposes metrics in Prometheus format at `/metrics`

## Prerequisites

- Node.js (v16 or later)
- npm or yarn

## Installation

```bash
# Install dependencies
npm install

# Build the application
npm run build
```

## Running the app

```bash
# Development mode
npm run start:dev

# Production mode
npm run start:prod
```

The application will be available at http://localhost:3000, and metrics will be exposed at http://localhost:3000/metrics.

## Notes

- GPU metrics may not be available on all systems, especially in virtualized environments
- The application uses the `systeminformation` library to collect system metrics
- Metrics are collected every 15 seconds by default
