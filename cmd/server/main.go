package main

import (
	"encoding/json"
	"log"
	"net/http"
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

	log.Printf("Listening on %s", httpServer.Addr)

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
