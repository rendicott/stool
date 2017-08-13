package game

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gapi/db"
	"github.com/gapi/util"

	"github.com/gorilla/mux"
)

func GameIndex(w http.ResponseWriter, r *http.Request) {
	data := db.GetDb()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(GetGames(data)); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func ShowGame(w http.ResponseWriter, r *http.Request) {
	data := db.GetDb()
	vars := mux.Vars(r)
	gameId, err := strconv.Atoi(vars["gameId"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	g := Game{Id: gameId}
	g.GetGame(data)
	/*	if errr := g.GetGame(data); errr != nil {
		util.RespondWithError(w, http.StatusInternalServerError, errr.Error())
		return
	}*/
	if errrr := json.NewEncoder(w).Encode(g); errrr != nil {
		panic(errrr)
	}
}

func GetGames(db *db.Database) Games {
	games := Games{}
	db.DB.Find(&games)
	/*rows, err := db.Query("SELECT * FROM games")

	if err != nil {
		panic(err)
		return nil
	}

	// defer statement call executed after whole getGames function returns
	defer rows.Close()

	games := []Game{}

	for rows.Next() {
		var g Game
		if err := rows.Scan(&g.Id, &g.Name); err != nil { //http://piotrzurek.net/2013/09/20/pointers-in-go.html
			panic(err)
			return nil
		}
		games = append(games, g)
	}

	*/
	return games
}

func (g *Game) GetGame(db *db.Database) Game {
	var game Game
	db.DB.Find(&game, g.Id)

	return game
}
