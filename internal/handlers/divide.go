package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/subhamBharadwaz/go-calculator-api/internal/models"
)

func HandleDivide(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid req method", http.StatusMethodNotAllowed)
		return
	}

	var req models.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid req body", http.StatusBadRequest)
		return
	}

	if req.B == 0 {
		http.Error(w, `{"error": "Divisor value cannot be 0"}`, http.StatusBadRequest)
		return
	}

	result := req.A / req.B
	resp := models.Response{
		Result: result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
