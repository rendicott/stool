package butt

import (
	"github.com/gapi/db"
)

type Butt struct {
	Id   int    `gorm:"primary_key;"`
	Name string `json:"name"`
}

type ButtDLInterface interface {
	RetrieveAllButts() ([]Butt, error)
	RetrieveSingleButt(int) (Butt, error)
}

type ButtDLGorm struct {
}

func (b *ButtDLGorm) RetrieveAllButts() ([]Butt, error) {
	data, err := db.NewDB()
	if err != nil {
		return nil, err
	}
	butts := []Butt{}
	data.Find(&butts)
	return butts, nil
}

func (b *ButtDLGorm) RetrieveSingleButt(buttId int) (Butt, error) {
	var butt Butt
	data, err := db.NewDB()
	if err != nil {
		return butt, err
	}
	data.Find(&butt, buttId)

	return butt, nil
}
