package http

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "any message",
	}

	encoder := json.NewEncoder(w)
	w.WriteHeader(http.StatusOK)
	encoder.Encode(payload)

}
