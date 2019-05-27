package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/events-app/events-api/internal/menu"
	"github.com/events-app/events-api/internal/platform/web"
	"github.com/gorilla/mux"
)

func GetMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "invalid menu ID")
		return
	}
	m, err := menu.Get(id)
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
	m.ID, err = menu.Add(m.Name, m.CardId)
	if err != nil {
		web.RespondWithError(w, http.StatusConflict, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusCreated, m)
}

func UpdateMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "invalid menu ID")
		return
	}
	var m menu.Menu
	err = json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		web.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()
	m.ID = id
	if err = menu.Update(id, m.Name, m.CardId); err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, m)
}

func DeleteMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "invalid menu ID")
		return
	}
	if err := menu.Delete(id); err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func GetCardOfMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "invalid menu ID")
		return
	}
	m, err := menu.GetCardOfMenu(id)
	if err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, m)
}
