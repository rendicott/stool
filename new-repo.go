// mock database (not thread safe!)
package main

import (
    // "fmt"
    "database/sql"
    // "strconv"
)

// Create some seed data
func init() {
    //load in current db stuff?

}

func GetGames(db *sql.DB) []Game {
    rows, err := db.Query("SELECT * FROM games")

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

    return games
}

func (g *Game) GetGame(db *sql.DB) error {
    return db.QueryRow("SELECT * FROM games WHERE id=$1", g.Id).Scan(&g.Id, &g.Name)
}



