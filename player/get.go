package player

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gapi/util"
	"github.com/gorilla/mux"
)

func PlayerIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	players, err := GetPlayers()
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err = json.NewEncoder(w).Encode(players); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func ShowPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerId, err := strconv.Atoi(vars["playerId"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	p := Player{Id: playerId}
	// todo : pass this by reference
	p, err = p.GetPlayer()
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(p); err != nil {
		panic(err)
	}
}
