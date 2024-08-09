package routes

import (
	"log/slog"
	"net/http"

	"github.com/subhamBharadwaz/go-calculator-api/handlers"
)

func NewRouter(logger *slog.Logger) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/add", handlers.HandleAdd(logger))

	return router
}
