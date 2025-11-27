package temporal

import (
	"context"
	"log/slog"
	"time"

	"github.com/cenkalti/backoff/v4"
	"go.temporal.io/sdk/client"
)

func GetTemporalClient(ctx context.Context, log *slog.Logger, options client.Options) (client.Client, error) {
	// Add logger to options
	if log != nil {
		options.Logger = NewTemporalLogger(ctx, log)
	}

	var temporalClient client.Client

	operation := func() error {
		var err error
		temporalClient, err = client.Dial(options)
		return err
	}

	// Configure exponential backoff
	backoffStrategy := backoff.NewExponentialBackOff()
	backoffStrategy.MaxElapsedTime = 30 * time.Second
	backoffStrategy.InitialInterval = 1 * time.Second
	backoffStrategy.MaxInterval = 10 * time.Second

	// Retry with max 3 attempts
	err := backoff.Retry(operation, backoff.WithMaxRetries(backoffStrategy, 3))

	return temporalClient, err
}
