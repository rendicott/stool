package outcome

import (
	"github.com/gapi/db"
)

// todo: foreign keys
// todo: include a int score
// todo: include a status win/lose/tie

type Outcome struct {
	Id       int  `gorm:"primary_key;"`
	GameId   int  `json:"gameid"`
	PlayerId int  `json:"playerid"`
	Win      bool `json:"win"`
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

func (o *Outcome) CreateOutcome() error {
	data, err := db.NewDB()
	if err != nil {
		return err
	}
	data.Create(&o)
	return nil
}

func (o *Outcome) DeleteOutcome() error {
	data, err := db.NewDB()
	if err != nil {
		return err
	}
	data.Delete(&o)
	return nil
}

func (o *Outcome) GetOutcome() (Outcome, error) {
	var outcome Outcome
	data, err := db.NewDB()
	if err != nil {
		return outcome, err
	}
	data.Find(&outcome, o.Id)

	return outcome, nil
}
