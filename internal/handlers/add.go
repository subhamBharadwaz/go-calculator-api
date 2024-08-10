package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/subhamBharadwaz/go-calculator-api/internal/models"
)

func HandleAdd(w http.ResponseWriter, r *http.Request) {
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

	// Check if the inputs are valid numbers
	if req.A == 0 && req.B == 0 {
		if r.ContentLength > 0 {
			http.Error(w, "Invalid input: both A and B must be numbers", http.StatusBadRequest)
			return
		}
	}

	result := req.A + req.B
	resp := models.Response{
		Result: result,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
