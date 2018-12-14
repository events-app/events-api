package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/events-app/events-api/internal/card"
	"github.com/events-app/events-api/internal/platform/web"
	"github.com/gorilla/mux"
)

func GetCard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	content := card.Find(name)
	if content == nil {
		web.ErrorJSON(w, name+" does not exist", http.StatusNoContent)
		return
	}
	if err := json.NewEncoder(w).Encode(&content); err != nil {
		log.Printf("error: encoding response: %s", err)
	}
}

func GetCards(w http.ResponseWriter, r *http.Request) {
	cards := card.GetAll()
	if cards == nil {
		web.ErrorJSON(w, "no cards in database", http.StatusNoContent)
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
		web.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = card.Add(c.Name, c.Text); err != nil {
		web.ErrorJSON(w, err.Error(), http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func UpdateCard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	var c card.Card
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		web.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = card.Update(name, c.Text); err != nil {
		web.ErrorJSON(w, err.Error(), http.StatusNotFound)
		return
		// http.Error(w, err.Error(), 404)
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteCard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	if err := card.Delete(name); err != nil {
		web.ErrorJSON(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
