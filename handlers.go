package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/events-app/events-api/internal/card"
	"github.com/events-app/events-api/internal/file"
	"github.com/events-app/events-api/internal/user"
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
	fmt.Fprintf(w, "POST: https://%s/api/v1/upload, Body: \"file\": somefile\n", r.Host)
	fmt.Fprintf(w, "GET: https://%s/files\n", r.Host)
	fmt.Fprintf(w, "GET: https://%s/files/{filename}\n", r.Host)
}

// headerMiddlaware makes every handler use headers CORS and JSON
func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func SecuredContent(w http.ResponseWriter, r *http.Request) {
	content := card.Find("secured")
	if err := json.NewEncoder(w).Encode(content); err != nil {
		log.Printf("error: encoding response: %s", err)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var u user.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		ErrorJSON(w, "Could not decode response")
		return
	}
	if !user.ValidateUsername(u.Username) {
		ErrorJSON(w, "Username is invalid")
		return
	}
	if u.Username != "admin" || u.Password != "admin" {
		ErrorJSON(w, "Username or password is invalid")
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
		ErrorJSON(w, err.Error())
		return
	}
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
		ErrorJSON(w, name+" does not exist")
		return
	}
	if err := json.NewEncoder(w).Encode(&content); err != nil {
		log.Printf("error: encoding response: %s", err)
	}
}

func GetCards(w http.ResponseWriter, r *http.Request) {
	cards := card.GetAll()
	if cards == nil {
		ErrorJSON(w, "no cards in database")
		return
	}
	if err := json.NewEncoder(w).Encode(&cards); err != nil {
		log.Printf("error: encoding response: %s", err)
	}
}

func AddCard(w http.ResponseWriter, r *http.Request) {
	var c card.Card
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err = card.Add(c.Name, c.Text); err != nil {
		ErrorJSON(w, err.Error())
	}
}

func UpdateCard(w http.ResponseWriter, r *http.Request) {
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

func UploadFile(uploadPath string, maxUploadSize int64) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// validate file size
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			ErrorJSON(w, "file is too largre")
			// renderError(w, "file is too largre", http.StatusBadRequest)
			return
		}

		// parse and validate file
		f, _, err := r.FormFile("file")
		if err != nil {
			ErrorJSON(w, "invalid file")
			return
		}
		defer f.Close()
		filename, err := file.Upload(f, uploadPath)
		if err != nil {
			ErrorJSON(w, err.Error())
		}
		fl := file.New(fmt.Sprintf("https://%s/files/%s", r.Host, filename))
		// f := file.File{Path: fmt.Sprintf("https://%s/files/%s", r.Host, filename)}
		if err := json.NewEncoder(w).Encode(&fl); err != nil {
			log.Printf("error: encoding response: %s", err)
		}
	})
}
