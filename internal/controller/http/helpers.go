package http

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 104856

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(data)
	if err != nil {
		return err
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have a single JSON value")
	}

	return nil
}

func writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {

	out, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.WriteHeader(status)

	if len(headers) > 0 {
		for _, header := range headers {
			for k, values := range header {
				for _, v := range values {
					w.Header().Set(k, v)
				}
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return writeJSON(w, statusCode, payload)

}
