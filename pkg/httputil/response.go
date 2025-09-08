package httputil

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func OK(w http.ResponseWriter, data any) {
	WriteJSON(w, 200, data)
}

func Created(w http.ResponseWriter, data any) {
	WriteJSON(w, 201, data)
}
