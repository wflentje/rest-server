package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/wflentje/rest-server/internal/api"
	"github.com/wflentje/rest-server/internal/handlers"
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
	server := handlers.NewServer()
	handler := api.Handler(server)
	handler = middleware.RequestLogger(handler)
	handler = middleware.Recoverer(handler)

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
		log.Printf("Listening on %s", httpServer.Addr)

		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed: %v", err)
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

	log.Println("Shutdown signal received")

	// Allow up to 10 seconds for active requests to complete.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}

	log.Println("Server stopped")
}
