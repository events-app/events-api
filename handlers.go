package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/events-app/events-api/internal/user"

	"github.com/events-app/events-api/internal/card"
	"github.com/gorilla/mux"

	jwt "github.com/dgrijalva/jwt-go"
)

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET: https://%s/api/v1/cards/main\n", r.Host)
	fmt.Fprintf(w, "GET: https://%s/api/v1/cards/{card-name}\n", r.Host)
	fmt.Fprintf(w, "GET: https://%s/api/v1/cards/secured, Bearer authorization\n", r.Host)
	fmt.Fprintf(w, "POST: https://%s/api/v1/login, Body: {\"username\":\"...\", \"password\":\"...\"}\n", r.Host)
	fmt.Fprintf(w, "POST: https://%s/api/v1/cards, Bearer authorization\n", r.Host)
	fmt.Fprintf(w, "PUT: https://%s/api/v1/cards/{name}, Bearer authorization\n", r.Host)
}

func SecuredContent(w http.ResponseWriter, r *http.Request) {
	content := card.Find("secured")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewEncoder(w).Encode(content); err != nil {
		log.Printf("error: encoding response: %s", err)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var u user.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		RespondJSON(true, "Could not decode response", w)
		return
	}
	if !user.ValidateUsername(u.Username) {
		RespondJSON(true, "Username is invalid", w)
		return
	}
	if u.Username != "admin" || u.Password != "admin" {
		RespondJSON(true, "Username or password is invalid", w)
		return
	}
	// set token expiration to 15 minutes
	expireToken := time.Now().Add(time.Minute * 15).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
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

func GetCard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	content := card.Find(name)
	if content == nil {
		RespondJSON(true, name+" does not exist", w)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewEncoder(w).Encode(&content); err != nil {
		log.Printf("error: encoding response: %s", err)
	}
}

func GetCards(w http.ResponseWriter, r *http.Request) {
	cards := card.GetAll()
	if cards == nil {
		RespondJSON(true, "no cards in database", w)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewEncoder(w).Encode(&cards); err != nil {
		log.Printf("error: encoding response: %s", err)
	}
}

func AddCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var c card.Card
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err = card.Add(c.Name, c.Text); err != nil {
		RespondJSON(true, err.Error(), w)
	}
}

func UpdateCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	name := params["name"]
	var c card.Card
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err = card.Update(name, c.Text); err != nil {
		http.Error(w, err.Error(), 404)
	}
}
