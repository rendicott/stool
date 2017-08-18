package main

import (
	"net/http"

	"github.com/gapi/game"
	"github.com/gapi/outcome"
	"github.com/gapi/player"
	"github.com/gapi/util"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GameIndex",
		"GET",
		"/games",
		game.GameIndex,
	},
	Route{
		"ShowGame",
		"GET",
		"/games/{gameId}",
		game.ShowGame,
	},
	Route{
		"GameCreate",
		"POST",
		"/games",
		game.CreateGame,
	},
	Route{
		"PlayerIndex",
		"GET",
		"/players",
		player.PlayerIndex,
	},
	Route{
		"ShowPlayer",
		"GET",
		"/players/{playerId}",
		player.ShowPlayer,
	},
	Route{
		"PlayerCreate",
		"POST",
		"/players",
		player.CreatePlayer,
	},
	Route{
		"GameDelete",
		"DELETE",
		"/games/{gameId}",
		game.DeleteGame,
	},
	Route{
		"PlayerDelete",
		"DELETE",
		"/players/{playerId}",
		player.DeletePlayer,
	},
	Route{
		Name:        "OutcomeIndex",
		Method:      "GET",
		Pattern:     "/outcomes",
		HandlerFunc: outcome.OutcomeIndex,
	},
	Route{
		Name:        "ShowOutcome",
		Method:      "GET",
		Pattern:     "/outcomes/{outcomeId}",
		HandlerFunc: outcome.ShowOutcome,
	},
	Route{
		Name:        "OutcomeCreate",
		Method:      "POST",
		Pattern:     "/outcomes",
		HandlerFunc: outcome.CreateOutcome,
	},
	Route{
		Name:        "OutcomeDelete",
		Method:      "DELETE",
		Pattern:     "/outcomes/{outcomeId}",
		HandlerFunc: outcome.DeleteOutcome,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = util.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	router.Handle("/", http.FileServer(http.Dir("templates")))

	return router
}
