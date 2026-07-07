package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/wflentje/rest-server/internal/api"
)

const (
	serverAddr        = ":8080"
	readTimeout       = 15 * time.Second
	readHeaderTimeout = 5 * time.Second
	writeTimeout      = 15 * time.Second
	idleTimeout       = 60 * time.Second
)

type Server struct{}

func (s *Server) GetHello(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, api.HelloResponse{
		Message: "Hello, World!",
	})
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func main() {
	server := &Server{}

	handler := api.Handler(server)

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
