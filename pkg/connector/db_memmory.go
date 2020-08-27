package connector

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// ConnectToMemmory will create connection instance to sqlite memory
// and migrate models to this connection
func ConnectToMemmory(models []interface{}) (*gorm.DB, error) {
	db, err := gorm.Open(
		"sqlite3",
		"file::memory:?cache=shared",
	)
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(models...).Error
	if err != nil {
		return nil, err
	}
	return db, nil
}
