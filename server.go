package todoapp

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Server struct {
	HttpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.HttpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1048576,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,

		Handler: handler,
	}

	logrus.Info("Run Server, port:8080")

	return s.HttpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	logrus.Info("Shutdown Server")
	return s.HttpServer.Shutdown(ctx)
}
