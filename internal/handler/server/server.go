package server

import (
	"context"
	"encoding/json"
	"net"
	"net/http"

	"notification/internal/deps"
	"notification/internal/generated"
)

type Server struct {
	logger      deps.Logger
	pushService deps.PushService
	address     string
}

func NewServer(
	log deps.Logger,
	pushService deps.PushService,
	address string,
) *Server {
	return &Server{
		address:     address,
		pushService: pushService,
		logger:      log,
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

func (s *Server) PostSendPush(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token string `json:"token"`
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.logger.Error(r.Context(), err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := s.pushService.SendPush(r.Context(), req.Token, req.Title, req.Body)
	if err != nil {
		s.logger.Error(r.Context(), err)
		http.Error(w, "Failed to send push", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Push sent successfully"))
}
