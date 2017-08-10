// mock database (not thread safe!)
package main

import (
	"fmt"
	"strconv"
	"time"
)

var currentGameId int
var currentPlayerId int
var currentStatId int

var games Games
var players Players
var stats Stats

// Create some seed data
func init() {
}

func RepoCreateStat(stat Stat) Stat {
	currentStatId += 1
	stat.Id = currentStatId
	stat.Date = time.Now()
	// fmt.Printf(stat.Date.Format("20060101"))
	stats = append(stats, stat)
	return stat
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