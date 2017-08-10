package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the gAPI")
}

// func GameIndex(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)
// 	if err := json.NewEncoder(w).Encode(GetGames(GiveMeCurrentApp().DB)); err != nil {
// 		panic(err)
// 	}
// }

// func ShowGame(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	gameId := vars["gameId"]
// 	// fmt.Fprintln(w, "Game:", gameId)
// 	if err := json.NewEncoder(w).Encode(RepoFindGame(gameId)); err != nil {
// 		panic(err)
// 	}
// }

func CreateGame(w http.ResponseWriter, r *http.Request) {
	var game Game
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &game); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateGame(game)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func DeleteGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameId := vars["gameId"]
	t := RepoDeleteGame(gameId)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func PlayerIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(players); err != nil {
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

func ShowPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerId := vars["playerId"]
	// fmt.Fprintln(w, "Game:", gameId)
	if err := json.NewEncoder(w).Encode(RepoFindPlayer(playerId)); err != nil {
		panic(err)
	}
}

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var player Player
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &player); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreatePlayer(player)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
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