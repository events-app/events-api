package handlers

import "net/http"

// GetProtocol returns type of HTTP protocol
func GetProtocol(r *http.Request) string {
	if r.TLS == nil {
		return "http"
	}
	return "https"
}
