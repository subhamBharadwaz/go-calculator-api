package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/subhamBharadwaz/go-calculator-api/internal/models"
)

func HandleSum(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid req method", http.StatusMethodNotAllowed)
		return
	}

	var numbers []float64
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	// sum the numbers
	var sum float64
	for _, num := range numbers {
		sum += num
	}

	resp := models.Response{
		Result: sum,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
