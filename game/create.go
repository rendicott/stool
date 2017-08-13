package game

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gapi/db"
)

func CreateGame(w http.ResponseWriter, r *http.Request) {
	data := db.GetDb()
	var g Game
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if errr := r.Body.Close(); errr != nil {
		panic(errr)
	}
	if errrr := json.Unmarshal(body, &g); errrr != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if errrrr := json.NewEncoder(w).Encode(errrr); errrrr != nil {
			panic(errrrr)
		}
	}
	g.CreateGame(data)
	/*if errrrrr := g.CreateGame(data); errrrrr != nil {
		util.RespondWithError(w, http.StatusInternalServerError, errrrrr.Error())
		return
	}*/
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if errrrrrr := json.NewEncoder(w).Encode(g); errrrrrr != nil {
		//respond with error here
		panic(errrrrrr)
	}
}

func (g *Game) CreateGame(db *db.Database) {
	// fmt.Printf("getting here in createGame\n")
	db.DB.Create(&g)
}
