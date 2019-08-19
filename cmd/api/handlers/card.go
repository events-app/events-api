package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/events-app/events-api/internal/card"
	"github.com/events-app/events-api/internal/platform/web"
	"github.com/gorilla/mux"

	"github.com/jmoiron/sqlx"
)

// Cards defines all of the handlers related to cards.
// It holds the apllication state needed by the handler method.
type Cards struct {
	DB *sqlx.DB
}
// GetCard gets all cards from the service layer encodes them for the
// client response
func (c *Cards) GetCard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "invalid card ID")
		return
	}
	ca, err := card.Get(c.DB, id)
	if err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, ca)
}

func (c *Cards) GetCards(w http.ResponseWriter, r *http.Request) {
	cards, err := card.GetAll(c.DB)
	if err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, cards)
}

func (c *Cards) AddCard(w http.ResponseWriter, r *http.Request) {
	var ca *card.Card
	err := json.NewDecoder(r.Body).Decode(&ca)
	if err != nil {
		web.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()
	ca, err = card.Add(c.DB, ca.Name, ca.Text, time.Now())
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusCreated, ca)
}

func (c *Cards)UpdateCard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "invalid card ID")
		return
	}
	var ca *card.Card
	err = json.NewDecoder(r.Body).Decode(&ca)
	if err != nil {
		web.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()
	ca.ID = id
	if err = card.Update(c.DB, id, ca.Name, ca.Text, time.Now()); err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
		// http.Error(w, err.Error(), 404)
	}
	web.RespondWithJSON(w, http.StatusOK, ca)
}

func (c *Cards)DeleteCard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "invalid card ID")
		return
	}

	if err := card.Delete(c.DB, id); err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
