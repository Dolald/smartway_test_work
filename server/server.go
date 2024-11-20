package server

import (
	"context"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           port,
		Handler:        handler,
		MaxHeaderBytes: viper.GetInt("server.maxHeaderBytes"),
		ReadTimeout:    viper.GetDuration("server.readTimeout"),
		WriteTimeout:   viper.GetDuration("server.writeTimeout"),
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
