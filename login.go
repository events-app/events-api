package main

import (
	"encoding/json"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/events-app/events-api/internal/platform/auth"
	"github.com/events-app/events-api/internal/platform/web"
	"github.com/events-app/events-api/internal/user"
)

// TODO: move Login to handlers and JWT to auth

func Login(w http.ResponseWriter, r *http.Request) {
	var u user.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		web.ErrorJSON(w, "Could not decode response", http.StatusInternalServerError)
		return
	}
	if !user.ValidateUsername(u.Username) {
		web.ErrorJSON(w, "Username is invalid", http.StatusBadRequest)
		return
	}
	if u.Username != "admin" || u.Password != "admin" {
		web.ErrorJSON(w, "Username or password is invalid", http.StatusBadRequest)
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
		web.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(
		auth.JwtToken{
			Token:   tokenString,
			Expires: expireToken,
		})
}