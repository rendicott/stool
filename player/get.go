package player

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gapi/db"
	"github.com/gapi/util"
	"github.com/gorilla/mux"
)

func PlayerIndex(w http.ResponseWriter, r *http.Request) {
	data := db.GetDb()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(GetPlayers(data)); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func ShowPlayer(w http.ResponseWriter, r *http.Request) {
	data := db.GetDb()
	vars := mux.Vars(r)
	playerId, err := strconv.Atoi(vars["playerId"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	p := Player{Id: playerId}
	p.GetPlayer(data)
	/*if errr := p.GetPlayer(data); errr != nil {
		util.RespondWithError(w, http.StatusInternalServerError, errr.Error())
		return
	}*/
	if errrr := json.NewEncoder(w).Encode(p); errrr != nil {
		panic(errrr)
	}
}

func GetPlayers(db *db.Database) Players {
	players := Players{}
	db.DB.Find(&players)
	/*rows, err := db.Query("SELECT * FROM players")

	if err != nil {
		panic(err)
		return nil
	}

	// defer statement call executed after whole getGames function returns
	defer rows.Close()

	players := []Player{}

	for rows.Next() {
		var p Player
		if err := rows.Scan(&p.Id, &p.Name); err != nil { //http://piotrzurek.net/2013/09/20/pointers-in-go.html
			panic(err)
			return nil
		}
		players = append(players, p)
	}
	*/
	return players
}

func (p *Player) GetPlayer(db *db.Database) Player {
	var player Player
	db.DB.Find(&player, p.Id)

	return player

}
