package db

import (
	"sync"

	"github.com/jinzhu/gorm"
)

type Database struct {
	DB *gorm.DB
}

var instance *Database
var once sync.Once

func GetDb() *Database {
	once.Do(func() {
		db, err := gorm.Open("sqlite3", "test.db")
		if err != nil {
			panic("failed to connect database")
		}
		instance.DB = db
	})
	return instance
}
