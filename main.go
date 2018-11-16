package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

type Content struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET: https://%s/api/v1/content/main\n", r.Host)
	fmt.Fprintf(w, "GET: https://%s/api/v1/content/secured, Bearer authorization\n", r.Host)
	fmt.Fprintf(w, "POST: https://%s/api/v1/content/login, Body: {\"username\":\"...\", \"password\":\"...\"}\n", r.Host)
}

func MainContent(w http.ResponseWriter, r *http.Request) {
	content := Content{
		Name: "main",
		Text: `# The New Event

The New Event is the best event ever.
You should definitelly attend!

+ Register at [The New Event](http://thenewevent.com/).
+ Come
+ Have fun

We are waiting for you!`,
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewEncoder(w).Encode(content); err != nil {
		log.Printf("error: encoding response: %s", err)
	}
}

func SecuredContent(w http.ResponseWriter, r *http.Request) {
	content := Content{
		Name: "secured",
		Text: "You are allowed to see this.",
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewEncoder(w).Encode(content); err != nil {
		log.Printf("error: encoding response: %s", err)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		RespondJSON(true, "Could not decode response", w)
		return
	}
	if user.Username != "admin" || user.Password != "admin" {
		RespondJSON(true, "User or password is invalid", w)
		return
	}
	// set token expiration to 15 minutes
	expireToken := time.Now().Add(time.Minute * 15).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      expireToken,
	})
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		RespondJSON(true, err.Error(), w)
		return
	}
	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})

}
