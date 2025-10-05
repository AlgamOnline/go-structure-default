package utils

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

// JSON mengirimkan response JSON standar
func JSON(w http.ResponseWriter, status int, success bool, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(APIResponse{
		Success: success,
		Data:    data,
		Message: msg,
	})
}
