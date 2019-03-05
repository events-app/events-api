package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"

	"github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/events-app/events-api/handlers"
	"github.com/events-app/events-api/internal/platform/web"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s\n", err))
	}
	r := mux.NewRouter()
	// use middleware handler
	r.Use(handlers.HeaderMiddleware)
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("jwt-key")), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, errMessage string) {
			web.RespondWithError(w, http.StatusInternalServerError, errMessage)
		},
	})

	r.HandleFunc("/", handlers.Info).Methods("GET")
	r.HandleFunc("/api/v1/health", handlers.HealthCheck).Methods("GET")
	r.Handle("/api/v1/cards/secured", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(handlers.SecuredContent)),
	)).Methods("GET")

	r.HandleFunc("/api/v1/cards/{id}", handlers.GetCard).Methods("GET")
	r.HandleFunc("/api/v1/cards", handlers.GetCards).Methods("GET")
	r.HandleFunc("/api/v1/cards", handlers.AddCard).Methods("POST")
	r.HandleFunc("/api/v1/cards/{id}", handlers.UpdateCard).Methods("PUT")
	r.HandleFunc("/api/v1/cards/{id}", handlers.DeleteCard).Methods("DELETE")

	r.HandleFunc("/api/v1/login", handlers.Login).Methods("POST")

	r.HandleFunc("/api/v1/menus/{id}", handlers.GetMenu).Methods("GET")
	r.HandleFunc("/api/v1/menus", handlers.GetMenus).Methods("GET")
	r.HandleFunc("/api/v1/menus", handlers.AddMenu).Methods("POST")
	r.HandleFunc("/api/v1/menus/{id}", handlers.UpdateMenu).Methods("PUT")
	r.HandleFunc("/api/v1/menus/{id}", handlers.DeleteMenu).Methods("DELETE")
	r.HandleFunc("/api/v1/menus/{id}/cards", handlers.GetCardOfMenu).Methods("GET")

	r.HandleFunc("/api/v1/upload", handlers.UploadFile(
		viper.GetString("upload.path"),
		viper.GetInt64("upload.max-file-size"))).Methods("POST")

	// r.PathPrefix("/files/").Handler(http.FileServer(http.Dir(uploadPath)))
	fs := http.FileServer(http.Dir(viper.GetString("upload.path")))
	// --- r.PathPrefix("/files/").Handler(http.StripPrefix("files/", fs))
	// r.Handle("/files", http.StripPrefix("/files", fs)).Methods("GET")
	r.HandleFunc("/files", handlers.GetFiles(viper.GetString("upload.path"))).Methods("GET")
	r.Handle("/files/{file}", http.StripPrefix("/files", fs)).Methods("GET")
	// http.Handle("/files/", http.StripPrefix("/files", fs))

	// temporary handlers for backward compatibility with frontend
	r.Handle("/api/v1/content/secured", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(handlers.SecuredContent)),
	)).Methods("GET")
	r.HandleFunc("/api/v1/content/{name}", handlers.GetCard).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = viper.GetString("server-port")
	}
	log.Println("Listening on port", port)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowedMethods: []string{"GET", "POST", "PATCH"},
		// AllowedHeaders: []string{"Bearer", "Content-Type"}
	})
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), c.Handler(r)); err != nil {
		log.Printf("error: listing and serving: %s", err)
		return
	}
}
