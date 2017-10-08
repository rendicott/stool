package player

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type PlayerDLFake struct {
}

func FakePlayerDataContextMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		playerDl := &PlayerDLFake{}
		c.Set("Db", playerDl)
		c.Next()
	}
}

func (g *PlayerDLFake) CreatePlayer(playerName string) (Player, error) {
	player := Player{Id: 1, Name: "Vlaada Chvatil"}
	return player, nil
}

func (g *PlayerDLFake) RetrieveAllPlayers() ([]Player, error) {
	players := []Player{{Id: 1, Name: "Vlaada Chvatil"}, {Id: 2, Name: "Uwe Rosenberg"}, {Id: 3, Name: "Stefan Feld"}}
	return players, nil
}

func (g *PlayerDLFake) RetrieveSinglePlayer(playerId int) (Player, error) {
	player := Player{}
	if playerId == 1 {
		player = Player{Id: 1, Name: "Klaus Teuber"}
	} else {
		return player, errors.New("Not found")
	}

	return player, nil
}

func (g *PlayerDLFake) DeletePlayer(playerId int) error {
	if playerId != 1 {
		return errors.New("Not found")
	}
	return nil
}
