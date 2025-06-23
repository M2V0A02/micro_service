// internal/infrastructure/server/handler.go
package server

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"notification/internal/application/notification"
	"notification/internal/generated"
	"notification/pkg/logger"
	"time"
)

type Server struct {
	service *notification.Service
	logger  *logger.Logger
	srv     *http.Server
	address string
}

func NewServer(service *notification.Service, logger *logger.Logger) *Server {
	return &Server{service: service, logger: logger}
}

func (s *Server) Run(ctx context.Context) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	h := generated.HandlerWithOptions(s, generated.StdHTTPServerOptions{
		BaseRouter: mux,
	})

	s.srv = &http.Server{
		Addr:    ":8080", // Добавь явно порт
		Handler: h,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	err := s.srv.ListenAndServe()
	if err != nil {
		s.logger.Error(ctx, err)
	}
}

func (h *Server) PostSendPush(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token string `json:"token"`
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error(r.Context(), err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	cmd := notification.SendNotificationCommand{
		Token: req.Token,
		Title: req.Title,
		Body:  req.Body,
	}
	if err := h.service.SendNotification(r.Context(), cmd); err != nil {
		h.logger.Error(r.Context(), err)
		http.Error(w, "Failed to send push", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Push sent successfully"))
}

func (h *Server) PostSchedulePush(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token  string    `json:"token"`
		Title  string    `json:"title"`
		Body   string    `json:"body"`
		SentAt time.Time `json:"sent_at"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error(r.Context(), err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	cmd := notification.SendScheduleNotificationCommand{
		Token:  req.Token,
		Title:  req.Title,
		Body:   req.Body,
		SentAt: req.SentAt,
	}

	if err := h.service.SendScheduleNotification(r.Context(), cmd); err != nil {
		h.logger.Error(r.Context(), err)
		http.Error(w, "Failed to send push", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Push sent successfully"))
}
