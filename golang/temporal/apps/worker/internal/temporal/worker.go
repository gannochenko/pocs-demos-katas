package temporal

import (
	"context"
	"errors"
	"lib/logger"
	libTemporal "lib/temporal"
	"log/slog"
	"worker/internal/domain"
	"worker/internal/interfaces"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func CreateWorker(ctx context.Context, client client.Client, log *slog.Logger, config *domain.Config, workflows []interfaces.TemporalWorkflowGroup, quitCh chan<- struct{}) (func() error, func(), error) {
	if client == nil {
		return nil, nil, errors.New("client is not initialized")
	}

	w := worker.New(client, libTemporal.DemoWorkflowsTaskQueue, worker.Options{
		MaxConcurrentWorkflowTaskPollers: config.Temporal.Worker.MaxConcurrentWorkflowTaskPollers,
		MaxConcurrentWorkflowTaskExecutionSize: config.Temporal.Worker.MaxConcurrentWorkflowTaskExecutionSize,
		MaxConcurrentActivityTaskPollers: config.Temporal.Worker.MaxConcurrentActivityTaskPollers,
		MaxConcurrentActivityExecutionSize: config.Temporal.Worker.MaxConcurrentActivityExecutionSize,
		OnFatalError: func(err error) {
			logger.Error(ctx, log, "Worker fatal error", logger.F("error", err.Error()))
			quitCh <- struct{}{}
		},
	})

	for _, wf := range workflows {
		for workflowName, workflowFunc := range wf.GetWorkflows() {
			w.RegisterWorkflowWithOptions(workflowFunc, workflow.RegisterOptions{
				Name: workflowName,
			})
		}
	}

	return func() error {
		logger.Info(ctx, log, "Worker started")
		return w.Start()
	}, w.Stop, nil
}
