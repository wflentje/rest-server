package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rest-server/internal/api"
	"rest-server/internal/config"
	"rest-server/internal/handlers"
	"rest-server/internal/logging"
	"rest-server/internal/middleware"
	"rest-server/internal/server"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger, err := logging.NewLogger(cfg.Logging)
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	apiHandler := handlers.NewHandler()

	handler := api.Handler(apiHandler)
	handler = middleware.RequestLogger(logger, handler)
	handler = middleware.Recoverer(logger, handler)

	httpServer := server.New(cfg.Server, handler)

	// Start the HTTP server in a separate goroutine so the main goroutine
	// can wait for an OS shutdown signal.
	go func() {
		logger.Info("server starting", "address", httpServer.Addr)

		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("server failed", "error", err)
			os.Exit(1)
		}
	}()

	// Create a context that is canceled when Ctrl+C (SIGINT) or SIGTERM is received.
	shutdownCtx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	// Block until a shutdown signal is received.
	<-shutdownCtx.Done()

	logger.Info("shutdown signal received")

	// Allow up to 10 seconds for active requests to complete.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Error("server shutdown failed", "error", err)
		os.Exit(1)
	}

	logger.Info("server stopped")
}
