package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/subhamBharadwaz/go-calculator-api/models"
)

func HandleAdd(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var req models.AddRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		result := req.A + req.B
		resp := models.AddResponse{
			Result: result,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
