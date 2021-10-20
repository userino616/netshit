package server

import (
	"context"
	"net/http"
	"netflix-auth/internal/config"
	"netflix-auth/pkg/logger"
)

type APIServer struct {
	htppserver *http.Server
}

func New(cfg *config.Config, r http.Handler) *APIServer {
	return &APIServer{
		htppserver: &http.Server{
			Addr:    cfg.Server.Addr,
			Handler: r,
		},
	}
}

func (s *APIServer) Start() error {
	l := logger.GetLogger()
	l.Info("starting server")

	return s.htppserver.ListenAndServe()
}

func (s *APIServer) Shutdown(ctx context.Context) error {
	return s.htppserver.Shutdown(ctx)
}
