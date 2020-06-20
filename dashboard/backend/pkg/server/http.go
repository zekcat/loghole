package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Option func(s *http.Server)

func WithReadTimeout(timeout time.Duration) Option {
	return func(s *http.Server) {
		s.ReadTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) Option {
	return func(s *http.Server) {
		s.WriteTimeout = timeout
	}
}

func WithIdleTimeout(timeout time.Duration) Option {
	return func(s *http.Server) {
		s.IdleTimeout = timeout
	}
}

type HTTP struct {
	addr   string
	server *http.Server
	router *mux.Router
}

func NewHTTP(addr string, options ...Option) *HTTP {
	router := mux.NewRouter()
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	for _, option := range options {
		option(server)
	}

	return &HTTP{
		addr:   addr,
		server: server,
		router: router,
	}
}

func (h *HTTP) ListenAndServe() error {
	err := h.server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}

func (h *HTTP) Router() *mux.Router {
	return h.router
}

func (h *HTTP) Addr() string {
	return h.addr
}

func (h *HTTP) Shutdown(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}
