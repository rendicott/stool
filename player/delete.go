package player

import (
	"net/http"
	"strconv"

	"github.com/gapi/util"

	"github.com/gorilla/mux"
)

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerId, err := strconv.Atoi(vars["playerId"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	p := Player{Id: playerId}
	if err := p.DeletePlayer(); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
