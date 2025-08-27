package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var SqlDB *gorm.DB

func NewSqlConn(name string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", name)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	SqlDB = db

	// doing auto migration
	db.AutoMigrate(
		&User{},
		&Rss{},
	)

	return db, nil
}
