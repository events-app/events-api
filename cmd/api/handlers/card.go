package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/events-app/events-api/internal/card"
	"github.com/events-app/events-api/internal/platform/web"
	"github.com/gorilla/mux"
)

func GetCard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "invalid card ID")
		return
	}
	c, err := card.Get(id)
	if err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, c)
}

func GetCards(w http.ResponseWriter, r *http.Request) {
	cards, err := card.GetAll()
	if err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, cards)
}

func AddCard(w http.ResponseWriter, r *http.Request) {
	var c card.Card
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		web.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()
	c.ID, err = card.Add(c.Name, c.Text)
	if err != nil {
		web.RespondWithError(w, http.StatusConflict, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusCreated, c)
}

func UpdateCard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "invalid card ID")
		return
	}
	var c card.Card
	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		web.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()
	c.ID = id
	if err = card.Update(id, c.Name, c.Text); err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
		// http.Error(w, err.Error(), 404)
	}
	web.RespondWithJSON(w, http.StatusOK, c)
}

func DeleteCard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "invalid card ID")
		return
	}

	if err := card.Delete(id); err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
