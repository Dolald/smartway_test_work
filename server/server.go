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

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           port,
		Handler:        handler,
		MaxHeaderBytes: configs.MaxHeaderBytes,
		ReadTimeout:    configs.ReadTimeout,
		WriteTimeout:   configs.WriteTimeout,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
