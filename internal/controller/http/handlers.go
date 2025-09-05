package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:auth,omitempty`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func handleBroker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "any message",
	}

	_ = writeJSON(w, http.StatusOK, payload)

}

func handleAction(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := readJSON(w, r, &requestPayload)
	if err != nil {
		errorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "auth":
		authenticate(w, requestPayload.Auth)
	default:
		errorJSON(w, errors.New("unknown action"))
	}
}

func authenticate(w http.ResponseWriter, payload AuthPayload) {
	jsonData, _ := json.Marshal(payload)

	request, err := http.NewRequest("POST", "http://auth-service/authenticate", bytes.NewReader(jsonData))
	if err != nil {
		errorJSON(w, err)
		return
	}

	client := &http.Client{}

	responce, err := client.Do(request)
	if err != nil {
		errorJSON(w, err)
		return
	}
	defer responce.Body.Close()

	if responce.StatusCode == http.StatusUnauthorized {
		errorJSON(w, errors.New("invalid credentials"))
		return
	} else if responce.StatusCode != http.StatusOK {
		errorJSON(w, errors.New("error calling auth service"))
		return
	}

	var jsonResp jsonResponse
	dec := json.NewDecoder(responce.Body)
	err = dec.Decode(&jsonResp)
	if err != nil {
		errorJSON(w, err)
		return
	}

	if jsonResp.Error {
		errorJSON(w, errors.New("status unauthorized"), http.StatusUnauthorized)
		return
	}

	var payloadResponse jsonResponse
	payloadResponse.Error = false
	payloadResponse.Message = "Authenticated!"
	payloadResponse.Data = jsonResp.Data

	writeJSON(w, http.StatusAccepted, payloadResponse)

}
