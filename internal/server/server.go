package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"

	"go-backend/internal/config"
	"go-backend/internal/handler"
	"go-backend/internal/service"
	"go-backend/internal/store"
)

// Server holds the HTTP server and dependencies.
type Server struct {
	httpServer *http.Server
	logger     *slog.Logger
}

// New builds the application and returns a Server.
func New(cfg *config.Config, logger *slog.Logger) *Server {
	if logger == nil {
		logger = slog.Default()
	}
	noteStore := store.NewInMemoryNoteStore()
	noteSvc := service.NewNotesService(noteStore)
	notesHandler := handler.NewNotesHandler(noteSvc)

	r := chi.NewRouter()
	r.Use(handler.Recover)
	r.Use(handler.Logging(logger))

	r.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	r.Route("/notes", func(r chi.Router) {
		r.Get("/", notesHandler.List)
		r.Post("/", notesHandler.Create)
		r.Get("/{id}", notesHandler.GetByID)
		r.Put("/{id}", notesHandler.Update)
		r.Delete("/{id}", notesHandler.Delete)
	})

	addr := fmt.Sprintf(":%d", cfg.ServerPort)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	return &Server{httpServer: srv, logger: logger}
}

// Start starts the HTTP server (blocking).
func (s *Server) Start() error {
	s.logger.Info("server starting", "addr", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down the server.
func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("server shutting down")
	return s.httpServer.Shutdown(ctx)
}
