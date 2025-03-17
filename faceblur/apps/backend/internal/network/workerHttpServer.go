package network

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"backend/interfaces"
	"backend/internal/util/syserr"

	"backend/internal/domain"
)

type WorkerHTTPServer struct {
	server         *http.Server

	configService   interfaces.ConfigService
	monitoringService   interfaces.MonitoringService
}

func NewWorkerHTTPServer(
	configService interfaces.ConfigService,
	monitoringService interfaces.MonitoringService,
) *WorkerHTTPServer {
	return &WorkerHTTPServer{
		configService:   configService,
		monitoringService: monitoringService,
	}
}

func (s *WorkerHTTPServer) GetMux(ctx context.Context) (http.Handler, error) {
	mainRouter := mux.NewRouter()

	mainRouter.Handle("/metrics", s.monitoringService.GetHandler())

	config, err := s.configService.GetConfig()
	if err != nil {
		return nil, syserr.Wrap(err, "could not get config")
	}

	externalMux := withCorsMiddleware(mainRouter, config)

	return externalMux, nil
}

func (s *WorkerHTTPServer) Start(ctx context.Context, config *domain.Config) error {
	serverMux, err := s.GetMux(ctx)
	if err != nil {
		return err
	}

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Backend.Worker.HTTP.Port),
		Handler: serverMux,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	return s.server.ListenAndServe()
}

func (s *WorkerHTTPServer) Stop(ctx context.Context) error {
	var errors []string

	shutdownErr := s.server.Shutdown(ctx)
	if shutdownErr != nil {
		errors = append(errors, shutdownErr.Error())
	}

	if len(errors) > 0 {
		return syserr.NewInternal(fmt.Sprintf("could not shut down HTTP server: %s", strings.Join(errors, ", ")))
	}

	return nil
}
