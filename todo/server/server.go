package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		MaxHeaderBytes: 1 << 20, // 1MB
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s Server) Close(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

// Some new comment

// Something new

// Another amazing commits
