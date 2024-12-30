package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"

	"github.com/glotovs198/golang-htmx-templ/internal/app"
)

func main() {
	if err := runApp(); err != nil {
		slog.Error("failed to run app", "err", err)
		log.Fatal(err)
	}
}

func runApp() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := app.Run(ctx); err != nil {
		return err
	}

	return nil
}
