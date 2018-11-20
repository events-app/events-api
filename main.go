package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

const key = "KLHkjhsd*h67r3gJhjuds"

func main() {
	r := mux.NewRouter()
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err string) {
			RespondJSON(true, err, w)
		},
	})

	r.HandleFunc("/", Info).Methods("GET")
	r.HandleFunc("/api/v1/content/main", MainContent).Methods("GET")
	r.Handle("/api/v1/content/secured", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(SecuredContent)),
	)).Methods("GET")
	r.HandleFunc("/api/v1/login", Login).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Println("Listening on port", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		log.Printf("error: listing and serving: %s", err)
		return
	}
}
