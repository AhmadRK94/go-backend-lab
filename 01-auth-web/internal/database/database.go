package database

import (
	"os/user"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var once sync.Once

func InitDB(dsn string) (*gorm.DB, error) {
	once.Do(func() {
		connection, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		db = connection

	})
	err := db.AutoMigrate(user.User{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
