package game

import (
	"net/http"
	"strconv"

	"github.com/gapi/db"
	"github.com/gapi/util"
	"github.com/gorilla/mux"
)

func DeleteGame(w http.ResponseWriter, r *http.Request) {
	data := db.GetDb()
	vars := mux.Vars(r)
	gameId, err := strconv.Atoi(vars["gameId"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	g := Game{Id: gameId}
	if err := g.DeleteGame(data); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (g *Game) DeleteGame(db *db.Database) error {
	db.DB.Delete(&g)
	return nil
}
