package workflows

import (
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

	fmt.Printf("%+v\n", request)

	// // Validate request fields
	// if err := h.validateWebhookRequest(&request); err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "validation failed"))
	// }

	// webhook := &domain.WebhookEvent{
	// 	EventID: request.EventId.String(),
	// 	EventTimestamp: request.EventTimestamp.Format(time.RFC3339),
	// 	EventType: string(request.EventType),
	// 	Payload: request.Payload,
	// }
	// err := h.webhookService.HandleWebhook(ctx.Request().Context(), webhook)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "failed to handle webhook"))
	// }

	return ctx.JSON(http.StatusOK, v1.GithubWorkflowResponse{
		WorkflowType: request.WorkflowType,
		WorkflowId:   "1234567890",
		Action:       &request.Action,
	})
}

// func (h *WebhooksHandler) validateWebhookRequest(request *webhooksV1.AcceptAcmeCorporationWebhookRequest) error {
// 	// Validate event_id is not empty
// 	if request.EventId.String() == "" || request.EventId.String() == "00000000-0000-0000-0000-000000000000" {
// 		return errors.New("event_id is required and must be a valid UUID")
// 	}

// 	// Validate event_timestamp is not zero
// 	if request.EventTimestamp.IsZero() {
// 		return errors.New("event_timestamp is required")
// 	}

// 	// Validate event_timestamp is not in the future
// 	if request.EventTimestamp.After(time.Now().Add(5 * time.Minute)) {
// 		return errors.New("event_timestamp cannot be in the future")
// 	}

// 	// Validate event_type is not empty
// 	if request.EventType == "" {
// 		return errors.New("event_type is required")
// 	}

// 	// Validate event_type is one of the allowed values
// 	validEventTypes := map[webhooksV1.AcceptAcmeCorporationWebhookRequestEventType]bool{
// 		webhooksV1.OrderCancelled: true,
// 		webhooksV1.OrderCompleted: true,
// 		webhooksV1.OrderCreated:   true,
// 		webhooksV1.UserCreated:    true,
// 		webhooksV1.UserDeleted:    true,
// 		webhooksV1.UserUpdated:    true,
// 	}

// 	if !validEventTypes[request.EventType] {
// 		return errors.New("event_type must be one of: order.created, order.completed, order.cancelled, user.created, user.updated, user.deleted")
// 	}

// 	// Validate payload is not nil
// 	if request.Payload == nil {
// 		return errors.New("payload is required")
// 	}

// 	return nil
// }
