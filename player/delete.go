package player

import (
	"net/http"
	"strconv"

	"github.com/gapi/db"
	"github.com/gapi/util"

	"github.com/gorilla/mux"
)

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	data := db.GetDb()
	vars := mux.Vars(r)
	playerId, err := strconv.Atoi(vars["playerId"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	p := Player{Id: playerId}
	if err := p.DeletePlayer(data); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (p *Player) DeletePlayer(db *db.Database) error {
	db.DB.Delete(p)
	return nil
}
