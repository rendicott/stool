// mock database (not thread safe!)
package main

import "fmt"

var currentId int

var games Games

// Create some seed data
func init() {
	RepoCreateGame(Game{Name: "Splendor", Players: 4})
	RepoCreateGame(Game{Name: "Love Letter", Players: 4})
}

func RepoCreateGame(game Game) Game {
	currentId += 1
	game.Id = currentId
	games = append(games, game)
	return game
}

func RepoFindGame(id int) Game {
	for _, game := range games {
		if game.Id == id {
			return game
		}
	}
	// return empty Game if not found
	return Game{}	
}

func RepoDeleteGame(id int) error {
	for i, game := range games {
		if game.Id == id {
			games = append(games[:i], games[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not delete: can't find Game with id of %d", id)
}