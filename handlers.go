package main

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the gAPI")
}


func DeleteGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameId := vars["gameId"]
	t := RepoDeleteGame(gameId)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func StatIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(stats); err != nil {
		panic(err)
	}
}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerId := vars["playerId"]
	t := RepoDeletePlayer(playerId)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}