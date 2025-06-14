package server

import (
	"context"
	"net"
	"net/http"

	"notification/internal/deps"
	"notification/internal/generated"
)

type Server struct {
	logger  deps.Logger
	address string
}

func NewServer(
	log deps.Logger,
	address string,
) *Server {
	return &Server{
		address: address,
		logger:  log,
	}
}

func (s *Server) Run(ctx context.Context) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	h := generated.HandlerFromMux(s, mux)

	srv := &http.Server{
		Handler: h,
		Addr:    s.address,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	err := srv.ListenAndServe()
	if err != nil {
		s.logger.Error(ctx, err)
	}
}
