package player

import "github.com/jinzhu/gorm"

type Players []Player

type Player struct {
	gorm.Model
	Id   int
	Name string
}
