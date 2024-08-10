package main

import (
	"net/http"

	"github.com/rs/cors"
	"github.com/subhamBharadwaz/go-calculator-api/internal/logger"
	"github.com/subhamBharadwaz/go-calculator-api/internal/middleware"
	"github.com/subhamBharadwaz/go-calculator-api/internal/routes"
	"golang.org/x/time/rate"
)

func main() {
	logger := logger.NewLogger()
	router := routes.NewRouter(logger)

	corsHandler := cors.Default().Handler(router)
	limiter := rate.NewLimiter(5, 10)

	stack := middleware.CreateStack(
		middleware.Logging(logger),
		middleware.RateLimiter(limiter),
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(corsHandler),
	}

	logger.Info("Server listening on port :8080")

	server.ListenAndServe()
}
