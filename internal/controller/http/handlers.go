package http

import (
	"net/http"
)

func broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "any message",
	}

	_ = writeJSON(w, http.StatusOK, payload)

}
