package webhooks

import (
	"net/http"

	webhooksV1 "gateway/internal/http/v1"
	"gateway/internal/interfaces"

	"github.com/labstack/echo/v4"
)

// WebhooksHandler implements the ServerInterface for webhooks operations
type WebhooksHandler struct {
	webhookDeduplicator interfaces.WebhookDeduplicator
}

// NewWebhooksHandler creates a new WebhooksHandler instance
func NewWebhooksHandler(webhookDeduplicator interfaces.WebhookDeduplicator) *WebhooksHandler {
	return &WebhooksHandler{
		webhookDeduplicator: webhookDeduplicator,
	}
}

func (h *WebhooksHandler) AcceptAcmeCorporationWebhook(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, webhooksV1.AcceptAcmeCorporationWebhookResponse{
		Message: "Webhook accepted successfully",
	})
}

// // GetStats handles GET /v1/stats - Get completion statistics
// func (h *TaskHandler) GetStats(ctx echo.Context) error {
// 	statsCount, err := h.taskService.GetStats(ctx.Request().Context())
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "failed to get stats"))
// 	}

// 	response := tasksV1.StatsResponse{
// 		Data: tasksV1.Stats{
// 			Completed: int(statsCount),
// 		},
// 	}

// 	return ctx.JSON(http.StatusOK, response)
// }
