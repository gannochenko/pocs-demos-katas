"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.MetricsService = void 0;
const common_1 = require("@nestjs/common");
const promClient = require("prom-client");
const si = require("systeminformation");
let MetricsService = class MetricsService {
    constructor() {
        this.updateInterval = 15000;
        this.register = new promClient.Registry();
        promClient.collectDefaultMetrics({ register: this.register });
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
        this.collectMetrics();
        setInterval(() => this.collectMetrics(), this.updateInterval);
    }
    async collectMetrics() {
        try {
            const cpuData = await si.currentLoad();
            this.cpuUsageGauge.set(cpuData.currentLoad);
            const memData = await si.mem();
            this.memoryUsageGauge.set(memData.active);
            this.memoryTotalGauge.set(memData.total);
            try {
                const gpuData = await si.graphics();
                if (gpuData && gpuData.controllers && gpuData.controllers.length > 0) {
                    gpuData.controllers.forEach((controller, index) => {
                        if (controller.utilizationGpu !== undefined) {
                            this.gpuUsageGauge.set({ gpu: `${controller.vendor}-${controller.model}` }, controller.utilizationGpu);
                        }
                        if (controller.memoryUsed !== undefined &&
                            controller.memoryTotal !== undefined) {
                            this.gpuMemoryGauge.set({ gpu: `${controller.vendor}-${controller.model}` }, controller.memoryUsed * 1024 * 1024);
                        }
                    });
                }
            }
            catch (error) {
                console.warn("Unable to collect GPU metrics:", error.message);
            }
        }
        catch (error) {
            console.error("Error collecting metrics:", error);
        }
    }
    async getMetrics() {
        return await this.register.metrics();
    }
};
exports.MetricsService = MetricsService;
exports.MetricsService = MetricsService = __decorate([
    (0, common_1.Injectable)(),
    __metadata("design:paramtypes", [])
], MetricsService);
//# sourceMappingURL=metrics.service.js.map