package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/events-app/events-api/internal/platform/auth"
	"github.com/events-app/events-api/internal/platform/web"
	"github.com/events-app/events-api/internal/user"
	"github.com/spf13/viper"
)

// TODO: move Login to handlers and JWT to auth

func Login(w http.ResponseWriter, r *http.Request) {
	var u user.User
	err := json.NewDecoder(r.Body).Decode(&u)
	r.Body.Close()
	if err != nil {
		web.RespondWithError(w, http.StatusInternalServerError, "Could not decode response")
		return
	}
	if !user.ValidateUsername(u.Username) {
		web.RespondWithError(w, http.StatusUnauthorized, "Username is invalid")
		return
	}
	if u.Username != "admin" || u.Password != "admin" {
		web.RespondWithError(w, http.StatusUnauthorized, "Username or password is invalid")
		return
	}
	// set token expiration to 15 minutes
	expireToken := time.Now().Add(time.Minute * 15).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"exp":      expireToken,
	})
	tokenString, err := token.SignedString([]byte(viper.GetString("jwt-key")))
	if err != nil {
		web.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK,
		auth.JwtToken{
			Token:   tokenString,
			Expires: expireToken,
		})
}
