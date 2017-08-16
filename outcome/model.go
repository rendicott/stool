package outcome

import (
	"github.com/gapi/db"
	"github.com/gapi/game"
	"github.com/gapi/player"
)

// todo: foreign keys
type Outcome struct {
	Id     int           `gorm:"primary_key;"`
	Game   game.Game     `json:"game"`
	Player player.Player `json:"player"`
	Win    bool          `json:"win"`
}

func GetOutcomes() ([]Outcome, error) {
	data, err := db.NewDB()
	if err != nil {
		return nil, err
	}
	outcomes := []Outcome{}
	data.Find(&outcomes)
	return outcomes, nil
}
