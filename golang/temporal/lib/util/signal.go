package util

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
)

func Run(ctx context.Context, quitChan <-chan struct{}, start func(chan os.Signal) error, stop func()) error {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	if err := start(sigChan); err != nil {
		stop()
		return errors.Wrap(err, "could not start application")
	}

	select {
	case <-sigChan:
		stop()
	case <-quitChan:
		stop()
	case <-ctx.Done():
		stop()
	}

	return nil
}
