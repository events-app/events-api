package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET: https://%s/api/v1/content/main\n", r.Host)
	fmt.Fprintf(w, "GET: https://%s/api/v1/content/secured, Bearer authorization\n", r.Host)
	fmt.Fprintf(w, "POST: https://%s/api/v1/content/login, Body: {\"username\":\"...\", \"password\":\"...\"}\n", r.Host)
}

func MainContent(w http.ResponseWriter, r *http.Request) {
	content := Find("main")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewEncoder(w).Encode(content); err != nil {
		log.Printf("error: encoding response: %s", err)
	}
}

func SecuredContent(w http.ResponseWriter, r *http.Request) {
	content := Find("secured")
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
	if !ValidateUsername(user.Username) {
		RespondJSON(true, "Username is invalid", w)
		return
	}
	if user.Username != "admin" || user.Password != "admin" {
		RespondJSON(true, "Username or password is invalid", w)
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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(
		JwtToken{
			Token:   tokenString,
			Expires: expireToken,
		})
}
