package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/subhamBharadwaz/go-calculator-api/internal/models"
)

func HandleSubtract(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req models.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {

		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return

	}

	result := req.A - req.B
	resp := models.Response{
		Result: result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
