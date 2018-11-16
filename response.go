package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	IsError bool   `json:"error"`
	Text    string `json:"message"`
}

// RespondJSON returns message and information if message is error
func RespondJSON(isError bool, text string, w http.ResponseWriter) {
	response := Response{isError, text}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jsonResponse)
}
