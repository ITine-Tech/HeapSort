package main

import (
	"encoding/json"
	"fmt"

	"net/http"

	"heapsort/internal"
)

type Request struct {
	Numbers []int `json:"numbers"`
}

func postHeapSort(w http.ResponseWriter, r *http.Request) {
	var request Request

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if len(request.Numbers) == 0 {
		http.Error(w, "numbers array cannot be empty", http.StatusBadRequest)
		return
	}

	if len(request.Numbers) == 10000 {
		http.Error(w, "numbers array cannot exceed 10000 elements", http.StatusBadRequest)
		return
	}

	if err := internal.HeapSort(request.Numbers); err != nil {
		http.Error(w, fmt.Errorf("heap sort failed: %v", err).Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"sorted_numbers": request.Numbers})
}
