package game

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gapi/util"

	"github.com/gorilla/mux"
)

func GameIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	games, err := GetGames()
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err = json.NewEncoder(w).Encode(games); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func ShowGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameId, err := strconv.Atoi(vars["gameId"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	g := Game{Id: gameId}
	game, err := g.GetGame()
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(game); err != nil {
		panic(err)
	}
}
