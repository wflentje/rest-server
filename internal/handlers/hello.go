package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-server/internal/api"
)

func (h *Handler) GetHello(w http.ResponseWriter, r *http.Request) {
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
