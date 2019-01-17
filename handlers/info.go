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
	fmt.Fprintf(w, "GET: %s://%s/api/v1/cards/{card-name}\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "POST: %s://%s/api/v1/cards, Bearer authorization\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "PUT: %s://%s/api/v1/cards/{name}, Bearer authorization\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "DELETE: %s://%s/api/v1/cards/{name}\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "\n==================== menus =====================================\n")
	fmt.Fprintf(w, "GET: %s://%s/api/v1/menus/{name}\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "GET: %s://%s/api/v1/menus\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "POST: %s://%s/api/v1/menus\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "PUT: %s://%s/api/v1/menus/{name}\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "DELETE: %s://%s/api/v1/menus/{name}\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "\n==================== files =====================================\n")
	fmt.Fprintf(w, "POST: %s://%s/api/v1/upload, Body: \"file\": somefile\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "GET: %s://%s/files\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "GET: %s://%s/files/{filename}\n", GetProtocol(r), r.Host)
	fmt.Fprintf(w, "GET: %s://%s/health\n", GetProtocol(r), r.Host)
}
