package temporal

import (
	"context"
	"lib/logger"
	"log/slog"

	"go.temporal.io/sdk/interceptor"
	"go.temporal.io/sdk/workflow"
)

// LoggingInterceptor implements interceptor.WorkerInterceptor
type LoggingInterceptor struct {
	interceptor.WorkerInterceptorBase
	ctx context.Context
	log *slog.Logger
}

// NewLoggingInterceptor creates a new logging interceptor
func NewLoggingInterceptor(ctx context.Context, log *slog.Logger) *LoggingInterceptor {
	return &LoggingInterceptor{
		ctx: ctx,
		log: log,
	}
}

// InterceptWorkflow intercepts workflow execution
func (l *LoggingInterceptor) InterceptWorkflow(ctx workflow.Context, next interceptor.WorkflowInboundInterceptor) interceptor.WorkflowInboundInterceptor {
	return &workflowInterceptor{
		WorkflowInboundInterceptorBase: interceptor.WorkflowInboundInterceptorBase{
			Next: next,
		},
		ctx: l.ctx,
		log: l.log,
	}
}

// workflowInterceptor implements interceptor.WorkflowInboundInterceptor
type workflowInterceptor struct {
	interceptor.WorkflowInboundInterceptorBase
	ctx context.Context
	log *slog.Logger
}

// ExecuteWorkflow intercepts workflow execution and logs failures
func (w *workflowInterceptor) ExecuteWorkflow(ctx workflow.Context, in *interceptor.ExecuteWorkflowInput) (interface{}, error) {
	result, err := w.Next.ExecuteWorkflow(ctx, in)

	if err != nil {
		workflowInfo := workflow.GetInfo(ctx)
		logger.Error(
			w.ctx,
			w.log,
			"Workflow failed",
			logger.F("type", workflowInfo.WorkflowType.Name),
			logger.F("workflow_id", workflowInfo.WorkflowExecution.ID),
			logger.F("run_id", workflowInfo.WorkflowExecution.RunID),
			logger.F("attempt", workflowInfo.Attempt),
			logger.F("error", err.Error()),
		)
	}

	return result, err
}
