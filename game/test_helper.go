package game

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type GameDLFake struct {
}

func FakeGameDataContextMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		gameDl := &GameDLFake{}
		c.Set("Db", gameDl)
		c.Next()
	}
}

func (g *GameDLFake) CreateGame(gameName string) (Game, error) {
	game := Game{Id: 1, Name: "Busen memo"}
	return game, nil
}

func (g *GameDLFake) RetrieveAllGames() ([]Game, error) {
	games := []Game{{Id: 1, Name: "(("}, {Id: 2, Name: "))"}, {Id: 3, Name: "(( ))"}}
	return games, nil
}

func (g *GameDLFake) RetrieveSingleGame(gameId int) (Game, error) {
	game := Game{}
	if gameId == 1 {
		game = Game{Id: 1, Name: "The Ungame"}
	} else {
		return game, errors.New("Not found")
	}

	return game, nil
}

func (g *GameDLFake) DeleteGame(gameId int) error {
	if gameId != 1 {
		return errors.New("Not found")
	}
	return nil
}
