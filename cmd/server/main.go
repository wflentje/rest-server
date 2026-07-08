package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/wflentje/rest-server/internal/api"
	"github.com/wflentje/rest-server/internal/handlers"
	"github.com/wflentje/rest-server/internal/logging"
	"github.com/wflentje/rest-server/internal/middleware"
)

const (
	serverAddr        = ":8080"
	readTimeout       = 15 * time.Second
	readHeaderTimeout = 5 * time.Second
	writeTimeout      = 15 * time.Second
	idleTimeout       = 60 * time.Second
)

func main() {
	logger := logging.NewLogger()

	server := handlers.NewServer()

	handler := api.Handler(server)
	handler = middleware.RequestLogger(logger, handler)
	handler = middleware.Recoverer(logger, handler)

	httpServer := &http.Server{
		Addr:              serverAddr,
		Handler:           handler,
		ReadTimeout:       readTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
	}

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
