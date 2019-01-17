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

func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	// if err := json.NewEncoder(w).Encode(&payload); err != nil {
	// 	log.Printf("error: encoding response: %s", err)
	// }
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	RespondWithJSON(w, statusCode, map[string]string{"error": message})
}
