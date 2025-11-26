package workflows

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	v1 "api/internal/http/v1"
	"api/internal/interfaces"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// WorkflowsHandler implements the ServerInterface for webhooks operations
type WorkflowsHandler struct {
	temporalService interfaces.TemporalService
}

// NewWorkflowsHandler creates a new WebhooksHandler instance
func NewWorkflowsHandler(temporalService interfaces.TemporalService) *WorkflowsHandler {
	return &WorkflowsHandler{
		temporalService: temporalService,
	}
}

func (h *WorkflowsHandler) ManageGithubWorkflow(ctx echo.Context) error {
	var request v1.GithubWorkflowRequest
	if err := ctx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "failed to bind request"))
	}

	workflowID := h.createWorkflowID(request.Parameters.Repository)

	err := h.temporalService.ExecuteWorkflow(ctx.Request().Context(), string(request.WorkflowType), workflowID, request.Parameters, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "failed to execute workflow"))
	}

	return ctx.JSON(http.StatusOK, v1.GithubWorkflowResponse{
		WorkflowType: request.WorkflowType,
		WorkflowId:   workflowID,
		Action:       &request.Action,
	})
}

func (h *WorkflowsHandler) createWorkflowID(repository string) string {
	hash := sha256.Sum256([]byte(repository))
	return fmt.Sprintf("github-workflow-%x", hash[:6])
}
