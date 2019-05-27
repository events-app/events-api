package handlers

import (
	"net/http"
	"strings"
)

// HeaderMiddleware makes every handler use headers CORS and JSON
func HeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// workaround
		// if response returns file then do not set content type
		// TODO: fix it
		// fmt.Println(r.RequestURI + ": " + strconv.FormatBool(strings.HasPrefix(r.RequestURI, "/files/")))
		if !strings.HasPrefix(r.RequestURI, "/files/") {
			w.Header().Set("Content-Type", "application/json")
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
