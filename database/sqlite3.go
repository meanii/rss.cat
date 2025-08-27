package database

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var SqlDB *gorm.DB

func NewSqlConn(name string) (*gorm.DB, error) {
	cwd, _ := os.Getwd()
	dbfile := fmt.Sprintf("%s/data/%s.db", cwd, name)

	// create data directory if not exists
	if err := os.MkdirAll(fmt.Sprintf("%s/data", cwd), os.ModePerm); err != nil {
		return nil, err
	}

	db, err := gorm.Open(sqlite.Open(dbfile), &gorm.Config{})
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
