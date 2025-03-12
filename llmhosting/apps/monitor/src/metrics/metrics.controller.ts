import { Controller, Get, Header } from "@nestjs/common";
import { MetricsService } from "./metrics.service";

@Controller()
export class MetricsController {
  constructor(private readonly metricsService: MetricsService) {}

  @Get("metrics")
  @Header("Content-Type", "text/plain")
  async getMetrics(): Promise<string> {
    return await this.metricsService.getMetrics();
  }

  @Get()
  getInfo(): string {
    return "System Metrics Exporter - Access /metrics for Prometheus metrics";
  }
}
