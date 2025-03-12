import { Injectable, OnModuleInit } from "@nestjs/common";
import * as promClient from "prom-client";
import * as si from "systeminformation";

@Injectable()
export class MetricsService implements OnModuleInit {
  private register: promClient.Registry;
  private cpuUsageGauge: promClient.Gauge;
  private memoryUsageGauge: promClient.Gauge;
  private memoryTotalGauge: promClient.Gauge;
  private gpuUsageGauge: promClient.Gauge;
  private gpuMemoryGauge: promClient.Gauge;
  private updateInterval = 15000; // 15 seconds

  constructor() {
    // Create a new registry
    this.register = new promClient.Registry();

    // Add default metrics (process CPU, memory usage)
    promClient.collectDefaultMetrics({ register: this.register });

    // Create custom gauges for system metrics
    this.cpuUsageGauge = new promClient.Gauge({
      name: "system_cpu_usage_percent",
      help: "Current CPU usage in percentage",
      registers: [this.register],
    });

    this.memoryUsageGauge = new promClient.Gauge({
      name: "system_memory_usage_bytes",
      help: "Current memory usage in bytes",
      registers: [this.register],
    });

    this.memoryTotalGauge = new promClient.Gauge({
      name: "system_memory_total_bytes",
      help: "Total system memory in bytes",
      registers: [this.register],
    });

    this.gpuUsageGauge = new promClient.Gauge({
      name: "system_gpu_usage_percent",
      help: "Current GPU usage in percentage",
      labelNames: ["gpu"],
      registers: [this.register],
    });

    this.gpuMemoryGauge = new promClient.Gauge({
      name: "system_gpu_memory_usage_bytes",
      help: "Current GPU memory usage in bytes",
      labelNames: ["gpu"],
      registers: [this.register],
    });
  }

  onModuleInit() {
    // Start collecting metrics
    this.collectMetrics();
    setInterval(() => this.collectMetrics(), this.updateInterval);
  }

  async collectMetrics() {
    try {
      // Check if we're running on macOS
      //const isMacOS = process.env.HOST_PLATFORM === "macos";

      // Collect CPU metrics - systeminformation works on macOS
      const cpuData = await si.currentLoad();
      this.cpuUsageGauge.set(cpuData.currentLoad);

      // Collect memory metrics - systeminformation works on macOS
      const memData = await si.mem();
      this.memoryUsageGauge.set(memData.active);
      this.memoryTotalGauge.set(memData.total);

      // Collect GPU metrics if available
      try {
        const gpuData = await si.graphics();
        if (gpuData && gpuData.controllers && gpuData.controllers.length > 0) {
          gpuData.controllers.forEach((controller, index) => {
            // Not all GPUs report utilization, so check if the value exists
            if (controller.utilizationGpu !== undefined) {
              this.gpuUsageGauge.set(
                { gpu: `${controller.vendor}-${controller.model}` },
                controller.utilizationGpu
              );
            }

            // Check if memory metrics are available
            if (
              controller.memoryUsed !== undefined &&
              controller.memoryTotal !== undefined
            ) {
              this.gpuMemoryGauge.set(
                { gpu: `${controller.vendor}-${controller.model}` },
                controller.memoryUsed * 1024 * 1024 // Convert to bytes
              );
            }
          });
        }
      } catch (error) {
        console.warn("Unable to collect GPU metrics:", error.message);
      }
    } catch (error) {
      console.error("Error collecting metrics:", error);
    }
  }

  async getMetrics(): Promise<string> {
    return await this.register.metrics();
  }
}
