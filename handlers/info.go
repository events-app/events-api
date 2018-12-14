package handlers

import (
	"fmt"
	"net/http"
)

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET: https://%s/api/v1/cards/main\n", r.Host)
	fmt.Fprintf(w, "GET: https://%s/api/v1/cards/{card-name}\n", r.Host)
	fmt.Fprintf(w, "GET: https://%s/api/v1/cards/secured, Bearer authorization\n", r.Host)
	fmt.Fprintf(w, "POST: https://%s/api/v1/login, Body: {\"username\":\"...\", \"password\":\"...\"}\n", r.Host)
	fmt.Fprintf(w, "POST: https://%s/api/v1/cards, Bearer authorization\n", r.Host)
	fmt.Fprintf(w, "PUT: https://%s/api/v1/cards/{name}, Bearer authorization\n", r.Host)
	fmt.Fprintf(w, "DELETE: https://%s/api/v1/cards/{name}\n", r.Host)
	fmt.Fprintf(w, "POST: https://%s/api/v1/upload, Body: \"file\": somefile\n", r.Host)
	fmt.Fprintf(w, "GET: https://%s/files\n", r.Host)
	fmt.Fprintf(w, "GET: https://%s/files/{filename}\n", r.Host)
	fmt.Fprintf(w, "GET: https://%s/health\n", r.Host)
}
