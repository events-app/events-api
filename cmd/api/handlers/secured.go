package handlers

import (
	"net/http"

	"github.com/events-app/events-api/internal/card"
	"github.com/events-app/events-api/internal/platform/web"
)

func SecuredContent(w http.ResponseWriter, r *http.Request) {
	c, err := card.Find("secured")
	if err != nil {
		web.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	web.RespondWithJSON(w, http.StatusOK, c)
}
