package handlers

import (
	"net/http"

	"github.com/events-app/events-api/internal/card"
	"github.com/events-app/events-api/internal/platform/web"
)

func (c *Cards) SecuredContent(w http.ResponseWriter, r *http.Request) {
	ca, err := card.Find(c.DB, "secured")
	if err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, ca)
}
