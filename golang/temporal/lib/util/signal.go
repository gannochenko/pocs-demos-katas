package util

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context, quitChan <-chan struct{}, start, stop func() error) error {
	if err := start(); err != nil {
		return err
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-sigChan:
		if err := stop(); err != nil {
			return err
		}
	case <-quitChan:
		if err := stop(); err != nil {
			return err
		}
	case <-ctx.Done():
		if err := stop(); err != nil {
			return err
		}
	}

	return nil
}
