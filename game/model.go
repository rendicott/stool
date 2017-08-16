package game

import (
	"github.com/gapi/db"
)

// note id is provided by default, its the primary key
// for all of our records
type Game struct {
	Id   int    `gorm:"primary_key;"`
	Name string `json:"name"`
}

func (g *Game) CreateGame() error {
	data, err := db.NewDB()
	if err != nil {
		return err
	}
	data.Create(&g)
	return nil
}

func (g *Game) DeleteGame() error {
	data, err := db.NewDB()
	if err != nil {
		return err
	}
	data.Delete(&g)
	return nil
}

// todo: one of these things is not like the other
func GetGames() ([]Game, error) {
	data, err := db.NewDB()
	if err != nil {
		return nil, err
	}
	games := []Game{}
	data.Find(&games)
	return games, nil
}

func (g *Game) GetGame() (Game, error) {
	var game Game
	data, err := db.NewDB()
	if err != nil {
		return game, err
	}
	data.Find(&game, g.Id)

	return game, nil
}
