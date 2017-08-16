package player

import "github.com/gapi/db"

type Player struct {
	Id   int `gorm:"primary_key;"`
	Name string
}

func (p *Player) CreatePlayer() error {
	data, err := db.NewDB()
	if err != nil {
		return err
	}
	data.Create(p)
	return nil
}

func (p *Player) DeletePlayer() error {
	data, err := db.NewDB()
	if err != nil {
		return err
	}
	data.Delete(p)
	return nil
}

func GetPlayers() ([]Player, error) {
	var players []Player
	data, err := db.NewDB()
	if err != nil {
		return players, err
	}
	data.Find(&players)
	return players, nil
}

func (p *Player) GetPlayer() (Player, error) {
	var player Player
	data, err := db.NewDB()
	if err != nil {
		return player, err
	}
	data.Find(&player, p.Id)

	return player, nil

}
