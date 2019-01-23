package handlers

import (
	"fmt"
	"net/http"
)

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "\n==================== login =====================================\n")
	fmt.Fprintf(w, "POST: %s://%s/api/v1/login, Body: {\"username\":\"...\", \"password\":\"...\"}\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "\n==================== cards =====================================\n")
	fmt.Fprintf(w, "GET: %s://%s/api/v1/cards/main\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "GET: %s://%s/api/v1/cards/secured, Bearer authorization\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "GET: %s://%s/api/v1/cards/{id}\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "POST: %s://%s/api/v1/cards, Body: \"name\": text, \"text\": text\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "PUT: %s://%s/api/v1/cards/{id}, Body: \"name\": text, \"text\": text\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "DELETE: %s://%s/api/v1/cards/{id}\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "\n==================== menus =====================================\n")
	fmt.Fprintf(w, "GET: %s://%s/api/v1/menus/{id}\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "GET: %s://%s/api/v1/menus\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "GET: %s://%s/api/v1/menus/{id}/cards\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "POST: %s://%s/api/v1/menus, Body: \"name\": text, \"cardId\": number\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "PUT: %s://%s/api/v1/menus/{id}, Body: \"name\": text, \"cardId\": number\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "DELETE: %s://%s/api/v1/menus/{id}\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "\n==================== files =====================================\n")
	fmt.Fprintf(w, "POST: %s://%s/api/v1/upload, Body/form-data: \"file\": file\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "GET: %s://%s/files\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "GET: %s://%s/files/{filename}\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "GET: %s://%s/health\n", GetProtocol(r), r.Host)
}
