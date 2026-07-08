package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecovererAllowsSuccessfulRequest(t *testing.T) {
	handler := Recoverer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestRecovererReturnsInternalServerErrorOnPanic(t *testing.T) {
	panicHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("test panic")
	})

	handler := Recoverer(panicHandler)

	req := httptest.NewRequest(http.MethodGet, "/panic", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected status %d, got %d", http.StatusInternalServerError, rec.Code)
	}
}

func TestRecovererKeepsServerUsableAfterPanic(t *testing.T) {
	var panicNext bool

	handler := Recoverer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if panicNext {
			panic("test panic")
		}

		w.WriteHeader(http.StatusOK)
	}))

	panicNext = true

	req := httptest.NewRequest(http.MethodGet, "/panic", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected status %d, got %d", http.StatusInternalServerError, rec.Code)
	}

	panicNext = false

	req = httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec = httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected server to remain usable, got status %d", rec.Code)
	}
}
