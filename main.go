package main

//to add: curl -H "Content-Type: application/json" -d '{"name": "Camel Up"}' http://localhost:8080/games
//to delete: curl -X "DELETE" http://localhost:8080/players/4

import (
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/negroni"
	"github.com/gapi/db"
	"github.com/gapi/game"
	"github.com/gapi/outcome"
	"github.com/gapi/player"
	"path/filepath"
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

	// try this!! https://stackoverflow.com/questions/39899257/angular-2-frontend-golang-backend
	n := negroni.Classic()
	//n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.UseHandler(r)

	abspath, err := filepath.Abs("./frontend")

	if err != nil {
		fmt.Print(err)
	}

	fs := http.Dir(abspath)

	r.PathPrefix("/").Handler(http.FileServer(fs))


	log.Fatal(http.ListenAndServe(":8080", n))
}


