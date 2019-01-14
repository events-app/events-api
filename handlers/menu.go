package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/events-app/events-api/internal/card"
	"github.com/events-app/events-api/internal/platform/web"
	"github.com/gorilla/mux"
)

func GetMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	menu := card.FindMenu(name)
	if menu == nil {
		web.ErrorJSON(w, name+" does not exist", http.StatusNoContent)
		return
	}
	if err := json.NewEncoder(w).Encode(&menu); err != nil {
		log.Printf("error: encoding response: %s", err)
	}
}

func GetMenus(w http.ResponseWriter, r *http.Request) {
	menus := card.GetAllMenus()
	if menus == nil {
		web.ErrorJSON(w, "no menu objects in database", http.StatusNoContent)
		return
	}
	if err := json.NewEncoder(w).Encode(&menus); err != nil {
		log.Printf("error: encoding response: %s", err)
	}
}

func AddMenu(w http.ResponseWriter, r *http.Request) {
	var m card.Menu
	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		web.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = card.Add(m.Name, m.Card); err != nil {
		web.ErrorJSON(w, err.Error(), http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func UpdateMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	var m card.Menu
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		web.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = card.UpdateMenu(name, m.Card); err != nil {
		web.ErrorJSON(w, err.Error(), http.StatusNotFound)
		return
		// http.Error(w, err.Error(), 404)
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	if err := card.DeleteMenu(name); err != nil {
		web.ErrorJSON(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
