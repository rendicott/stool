package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
