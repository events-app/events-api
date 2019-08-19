package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/events-app/events-api/internal/menu"
	"github.com/events-app/events-api/internal/platform/web"
	"github.com/gorilla/mux"

	"github.com/jmoiron/sqlx"
)

// Menus defines all of the handlers related to menus.
// It holds the apllication state needed by the handler method.
type Menus struct {
	DB *sqlx.DB
}

// GetMenu gets all menus from the service layer encodes them for the
// client response
func (m *Menus) GetMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "invalid menu ID")
		return
	}
	me, err := menu.Get(m.DB, id)
	if err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, me)
}

func (m *Menus) GetMenus(w http.ResponseWriter, r *http.Request) {
	menus, err := menu.GetAll(m.DB)
	if err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, menus)
}

func (m *Menus) AddMenu(w http.ResponseWriter, r *http.Request) {
	var me *menu.Menu
	err := json.NewDecoder(r.Body).Decode(&me)
	if err != nil {
		web.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()
	me, err = menu.Add(m.DB, me.Name, me.CardID, time.Now())
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusCreated, me)
}

func (m *Menus) UpdateMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "invalid menu ID")
		return
	}
	var me *menu.Menu
	err = json.NewDecoder(r.Body).Decode(&me)
	if err != nil {
		web.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()
	me.ID = id
	if err = menu.Update(m.DB, id, me.Name, me.CardID, time.Now()); err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, me)
}

func (m *Menus) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "invalid menu ID")
		return
	}
	if err := menu.Delete(m.DB, id); err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (m *Menus) GetCardOfMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		web.RespondWithError(w, http.StatusBadRequest, "invalid menu ID")
		return
	}
	me, err := menu.GetCardOfMenu(m.DB, id)
	if err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, me)
}
