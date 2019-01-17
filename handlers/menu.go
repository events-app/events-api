package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/events-app/events-api/internal/menu"
	"github.com/events-app/events-api/internal/platform/web"
	"github.com/gorilla/mux"
)

func GetMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	m, err := menu.Find(name)
	if err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, m)
}

func GetMenus(w http.ResponseWriter, r *http.Request) {
	menus, err := menu.GetAll()
	if err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, menus)
}

func AddMenu(w http.ResponseWriter, r *http.Request) {
	var m menu.Menu
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		web.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()
	if err = menu.Add(m.Name, m.Card); err != nil {
		web.RespondWithError(w, http.StatusConflict, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusCreated, m)
}

func UpdateMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	var m menu.Menu
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		web.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()
	m.Name = name
	if err = menu.Update(name, m.Card); err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, m)
}

func DeleteMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	if err := menu.Delete(name); err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
