package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type wrappedWritter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWritter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func Logging(logger *slog.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			wrapped := &wrappedWritter{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
			}

			next.ServeHTTP(wrapped, r)

			logger.Info(
				"handled request",
				slog.Int("statusCode", wrapped.statusCode),
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Duration("duration", time.Since(start)),
			)
		})
	}
}
