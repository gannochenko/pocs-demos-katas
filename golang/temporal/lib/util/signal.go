package util

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
)

func Run(ctx context.Context, quitChan <-chan struct{}, start func(chan os.Signal) error, stop func() error) error {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	if err := start(sigChan); err != nil {
		// couldn't start, let's try to stop
		if err := stop(); err != nil {
			return errors.Wrap(err, "could not stop application")
		}
		return errors.Wrap(err, "could not start application")
	}

	select {
	case <-sigChan:
		if err := stop(); err != nil {
			return errors.Wrap(err, "could not stop application")
		}
	case <-quitChan:
		if err := stop(); err != nil {
			return errors.Wrap(err, "could not stop application")
		}
	case <-ctx.Done():
		if err := stop(); err != nil {
			return errors.Wrap(err, "could not stop application")
		}
	}

	return nil
}
