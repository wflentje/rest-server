package middleware

import (
	"log/slog"
	"net/http"
)

// Recoverer catches panics from downstream handlers and returns
// an HTTP 500 response instead of terminating the server.
func Recoverer(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(
					"handler panic",
					"panic", err,
					"method", r.Method,
					"path", r.URL.Path,
				)

				http.Error(
					w,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError,
				)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
