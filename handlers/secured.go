package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/events-app/events-api/internal/card"
)

func SecuredContent(w http.ResponseWriter, r *http.Request) {
	content := card.Find("secured")
	if err := json.NewEncoder(w).Encode(content); err != nil {
		log.Printf("error: encoding response: %s", err)
	}
}
