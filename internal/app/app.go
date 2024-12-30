package app

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/glotovs198/golang-htmx-templ/internal/config"
	"github.com/glotovs198/golang-htmx-templ/internal/handler"
	"github.com/go-chi/chi/v5"
)

func Run(ctx context.Context) error {
	cfg := config.NewConfig()

	r := chi.NewRouter()
	handler.RegisterRoutes(r)

	s := &http.Server{
		Addr:    cfg.AppHost,
		Handler: r,
	}

	go func() {
		<-ctx.Done()
		slog.Info("shutting down server")
		s.Shutdown(ctx)
	}()

	slog.Info("starting server", "addr", cfg.AppHost)

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
