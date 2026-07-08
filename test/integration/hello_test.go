package integration

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wflentje/rest-server/internal/api"
	"github.com/wflentje/rest-server/internal/handlers"
	"github.com/wflentje/rest-server/internal/middleware"
)

func TestHelloEndpoint(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	apiHandler := handlers.NewHandler()

	handler := api.Handler(apiHandler)
	handler = middleware.RequestLogger(logger, handler)
	handler = middleware.Recoverer(logger, handler)

	ts := httptest.NewServer(handler)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/hello")
	if err != nil {
		t.Fatalf("failed to call /hello: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var body api.HelloResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if body.Message != "Hello, World!" {
		t.Fatalf("expected message %q, got %q", "Hello, World!", body.Message)
	}
}
