package webhooks

import (
	"net/http"
	"time"

	"gateway/internal/domain"
	webhooksV1 "gateway/internal/http/v1"
	"gateway/internal/interfaces"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// WebhooksHandler implements the ServerInterface for webhooks operations
type WebhooksHandler struct {
	webhookService interfaces.WebhookService
}

// NewWebhooksHandler creates a new WebhooksHandler instance
func NewWebhooksHandler(webhookService interfaces.WebhookService) *WebhooksHandler {
	return &WebhooksHandler{
		webhookService: webhookService,
	}
}

func (h *WebhooksHandler) AcceptAcmeCorporationWebhook(ctx echo.Context) error {
	var request webhooksV1.AcceptAcmeCorporationWebhookRequest
	if err := ctx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "failed to bind request"))
	}

	webhook := &domain.WebhookEvent{
		EventID: request.EventId.String(),
		EventTimestamp: request.EventTimestamp.Format(time.RFC3339),
		EventType: string(request.EventType),
		Payload: request.Payload,
	}
	err := h.webhookService.HandleWebhook(ctx.Request().Context(), webhook)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "failed to handle webhook"))
	}

	return ctx.JSON(http.StatusOK, webhooksV1.AcceptAcmeCorporationWebhookResponse{
		Message: "Webhook accepted successfully",
	})
}
