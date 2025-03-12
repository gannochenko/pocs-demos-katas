import { OnModuleInit } from "@nestjs/common";
export declare class MetricsService implements OnModuleInit {
    private register;
    private cpuUsageGauge;
    private memoryUsageGauge;
    private memoryTotalGauge;
    private gpuUsageGauge;
    private gpuMemoryGauge;
    private updateInterval;
    constructor();
    onModuleInit(): void;
    collectMetrics(): Promise<void>;
    getMetrics(): Promise<string>;
}
