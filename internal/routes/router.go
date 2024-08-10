package routes

import (
	"log/slog"
	"net/http"

	"github.com/subhamBharadwaz/go-calculator-api/internal/handlers"
)

func NewRouter(logger *slog.Logger) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/add", handlers.HandleAdd)
	router.HandleFunc("/subtract", handlers.HandleSubtract)
	router.HandleFunc("/multiply", handlers.HandleMultiply)
	router.HandleFunc("/divide", handlers.HandleDivide)
	router.HandleFunc("/sum", handlers.HandleSum)
	return router
}
