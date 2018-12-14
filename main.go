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
const maxUploadSize = 2 * 1048576 // bytes = 2 mb
const uploadPath = "./uploaded-files"
const serverPort = "8000"

func main() {
	r := mux.NewRouter()
	// use middleware handler
	r.Use(HeaderMiddleware)
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err string) {
			ErrorJSON(w, err, http.StatusInternalServerError)
		},
	})

	r.HandleFunc("/", Info).Methods("GET")
	r.HandleFunc("/api/v1/health", HealthCheck).Methods("GET")
	r.Handle("/api/v1/cards/secured", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(SecuredContent)),
	)).Methods("GET")
	r.HandleFunc("/api/v1/cards/{name}", GetCard).Methods("GET")
	r.HandleFunc("/api/v1/cards", GetCards).Methods("GET")
	r.HandleFunc("/api/v1/login", Login).Methods("POST")
	r.HandleFunc("/api/v1/cards", AddCard).Methods("POST")
	r.HandleFunc("/api/v1/cards/{name}", UpdateCard).Methods("PUT")
	r.HandleFunc("/api/v1/cards/{name}", DeleteCard).Methods("DELETE")
	r.HandleFunc("/api/v1/upload", UploadFile(uploadPath, maxUploadSize)).Methods("POST")
	// r.PathPrefix("/files/").Handler(http.FileServer(http.Dir(uploadPath)))
	fs := http.FileServer(http.Dir(uploadPath))
	// r.PathPrefix("/files/").Handler(http.StripPrefix("files/", fs))
	r.Handle("/files", http.StripPrefix("/files", fs)).Methods("GET")
	r.Handle("/files/{file}", http.StripPrefix("/files", fs)).Methods("GET")
	// http.Handle("/files/", http.StripPrefix("/files", fs))

	// temporary handlers for backward compatibility with frontend
	r.Handle("/api/v1/content/secured", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(SecuredContent)),
	)).Methods("GET")
	r.HandleFunc("/api/v1/content/{name}", GetCard).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = serverPort
	}
	log.Println("Listening on port", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		log.Printf("error: listing and serving: %s", err)
		return
	}
}
