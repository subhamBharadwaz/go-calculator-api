package main

import (
	"net/http"

	"github.com/rs/cors"
	"github.com/subhamBharadwaz/go-calculator-api/internal/logger"
	"github.com/subhamBharadwaz/go-calculator-api/internal/middleware"
	"github.com/subhamBharadwaz/go-calculator-api/internal/routes"
)

func main() {
	logger := logger.NewLogger()
	router := routes.NewRouter(logger)

	corsHandler := cors.Default().Handler(router)

	handler := middleware.Logging(logger, corsHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	logger.Info("Server listening on port :8080")

	server.ListenAndServe()
}
