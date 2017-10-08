package player

import "github.com/gapi/db"

type Player struct {
	Id   int    `gorm:"primary_key;"`
	Name string `json:"name"`
}

type PlayerDLInterface interface {
	RetrieveAllPlayers() ([]Player, error)
	RetrieveSinglePlayer(int) (Player, error)
	CreatePlayer(string) (Player, error)
	DeletePlayer(int) error
}

type PlayerDLGorm struct {
}

func (g *PlayerDLGorm) RetrieveSinglePlayer(playerId int) (Player, error) {
	var player Player
	data, err := db.NewDB()
	if err != nil {
		return player, err
	}
	data.Find(&player, playerId)

	return player, nil
}

func (g *PlayerDLGorm) RetrieveAllPlayers() ([]Player, error) {
	data, err := db.NewDB()
	if err != nil {
		return nil, err
	}
	players := []Player{}
	data.Find(&players)
	return players, nil
}

func (p *PlayerDLGorm) CreatePlayer(playerName string) (Player, error) {
	var player Player
	data, err := db.NewDB()
	if err != nil {
		return player, err
	}
	player = Player{Name: playerName}
	data.Create(&player)
	return player, err
}

func (p *PlayerDLGorm) DeletePlayer(playerId int) error {
	data, err := db.NewDB()
	player := Player{Id: playerId}
	if err != nil {
		return err
	}
	data.Delete(&player)
	return err
}
