package server

import (
	"context"
	"net/http"

	"github.com/Dolald/smartway_test_work/configs"
	_ "github.com/lib/pq"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(handler http.Handler, cfg *configs.ServerConfig) error {
	s.httpServer = &http.Server{
		Addr:           cfg.Port,
		Handler:        handler,
		MaxHeaderBytes: cfg.MaxHeaderBytes,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
