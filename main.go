package main

import (
	"net/http"

	"github.com/subhamBharadwaz/go-calculator-api/logger"
	"github.com/subhamBharadwaz/go-calculator-api/middleware"
	"github.com/subhamBharadwaz/go-calculator-api/routes"
)

func main() {
	logger := logger.NewLogger()
	router := routes.NewRouter(logger)

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logging(logger, router),
	}

	logger.Info("Server listening on port :8080")
	server.ListenAndServe()
}
