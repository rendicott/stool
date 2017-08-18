package main

//to add: curl -H "Content-Type: application/json" -d '{"name": "Camel Up"}' http://localhost:8080/games
//to delete: curl -X "DELETE" http://localhost:8080/players/4

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gapi/db"
	"github.com/gapi/game"
	"github.com/gapi/outcome"
	"github.com/gapi/player"
)

func main() {

	db, err := db.NewDB()
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(player.Player{})
	db.AutoMigrate(game.Game{})
	db.AutoMigrate(outcome.Outcome{})
	r := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}


