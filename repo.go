// mock database (not thread safe!)
package main

import (
    "fmt"
    "database/sql"
    // "strconv"
)

// Create some seed data
func init() {

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

func (g *Game) CreateGame(db *sql.DB) error {
    // fmt.Printf("getting here in createGame\n")
    err := db.QueryRow("INSERT INTO games VALUES(DEFAULT, $1) RETURNING id", g.Name).Scan(&g.Id)
    if err != nil {
        return err
    }
    return nil
}

func (g *Game) DeleteGame(db *sql.DB) error {
    fmt.Printf("Id is $1", g.Id)
    _, err := db.Exec("DELETE FROM games WHERE id=$1", g.Id)
    return err
}

func GetPlayers(db *sql.DB) []Player {
    rows, err := db.Query("SELECT * FROM players")

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

    return players
}

func (p *Player) GetPlayer(db *sql.DB) error {
    return db.QueryRow("SELECT * FROM players WHERE id=$1", p.Id).Scan(&p.Id, &p.Name)
}

func (p *Player) CreatePlayer(db *sql.DB) error {
    // fmt.Printf("getting here in createGame\n")
    err := db.QueryRow("INSERT INTO players VALUES(DEFAULT, $1) RETURNING id", p.Name).Scan(&p.Id)
    if err != nil {
        return err
    }
    return nil
}

func (p *Player) DeletePlayer(db *sql.DB) error {
    fmt.Printf("Id is $1", p.Id)
    _, err := db.Exec("DELETE FROM players WHERE id=$1", p.Id)

    return err
}



