package web

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"error"`
}

// ErrorJSON returns error message in JSON format
func ErrorJSON(w http.ResponseWriter, message string, statusCode int) {
	response := Response{Message: message}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jsonResponse)
}

// func renderError(w http.ResponseWriter, message string, statusCode int) {
// 	w.WriteHeader(http.StatusBadRequest)
// 	w.Write([]byte(message))
// }
