package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	sentry "github.com/getsentry/sentry-go"

	"requiems-api/app"
	"requiems-api/platform/config"
)

// shutdownTimeout bounds how long the server waits for in-flight requests to
// drain after receiving a shutdown signal before forcing an exit.
const shutdownTimeout = 15 * time.Second

func main() {
	os.Exit(run())
}

func run() int {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg := config.Load()

	if cfg.Environment != "development" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:              cfg.SentryDSN,
			Environment:      cfg.Environment,
			TracesSampleRate: 0.01,
		}); err != nil {
			logger.Error("sentry.Init failed", "error", err)
		}
	}

	appInstance, err := app.New(ctx, cfg)

	if err != nil {
		logger.Error("failed to initialise app", "error", err)
		sentry.Flush(2 * time.Second)
		return 1
	}

	addr := fmt.Sprintf(":%s", cfg.Port)

	server := &http.Server{
		Addr:              addr,
		Handler:           appInstance.Handler(),
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	serveErr := make(chan error, 1)

	go func() {
		logger.Info("API server listening", "addr", addr)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serveErr <- err
			return
		}

		serveErr <- nil
	}()

	select {
	case err := <-serveErr:
		if err != nil {
			logger.Error("server error", "error", err)
			sentry.Flush(2 * time.Second)
			appInstance.Close()
			return 1
		}
	case <-ctx.Done():
		logger.Info("shutdown signal received, draining in-flight requests")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.Error("graceful shutdown failed", "error", err)
		}
	}

	appInstance.Close()
	sentry.Flush(2 * time.Second)
	logger.Info("shutdown complete")
	return 0
}
