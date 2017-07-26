// mock database (not thread safe!)
package main

import (
	"fmt"
	"strconv"
)

var currentGameId int
var currentPlayerId int

var games Games
var players Players

// Create some seed data
func init() {
	RepoCreateGame(Game{Name: "Splendor", Players: 4})
	RepoCreateGame(Game{Name: "Love Letter", Players: 4})
	RepoCreatePlayer(Player{Name: "Jessica"})
}

func RepoCreateGame(game Game) Game {
	currentGameId += 1
	game.Id = currentGameId
	games = append(games, game)
	return game
}

func RepoCreatePlayer(player Player) Player {
	currentPlayerId +=1
	player.Id = currentPlayerId
	players = append(players, player)
	return player
}

// make it expect a string here!!
func RepoFindGame(id string) Game {
	for _, game := range games {
		if strconv.Itoa(game.Id) == id {
			return game
		}
	}
	// return empty Game if not found
	return Game{}	
}

func RepoFindPlayer(id string) Player {
	for _, player := range players {
		if strconv.Itoa(player.Id) == id {
			return player
		}
	}
	// return empty player obj if not found
	return Player{}
}

func RepoDeleteGame(id string) error {
	for i, game := range games {
		if strconv.Itoa(game.Id) == id {
			games = append(games[:i], games[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not delete: can't find Game with id of %d", id)
}

func RepoDeletePlayer(id string) error {
	for i, player := range players {
		if strconv.Itoa(player.Id) == id {
			players = append(players[:i], players[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not delete: can't find Player with id of %d", id)
}