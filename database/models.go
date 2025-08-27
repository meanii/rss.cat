package database

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// User for storing users whoever intract with the bot
type User struct {
	gorm.Model
	ID                int   `gorm:"primaryKey"`
	UserId            int64 `gorm:"unique"`
	FirstName         sql.NullString
	LastName          sql.NullString
	Username          sql.NullString
	FirstSeen         time.Time `gorm:"autoCreateTime"`
	NotificationCount int64     `gorm:"default:0"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

// Rss for links, and subscribers
type Rss struct {
	gorm.Model
	ID           int `gorm:"primaryKey"`
	OwnerId      int64
	Link         string
	LastItemGUID string       `gorm:"size:255"`
	Active       sql.NullBool `gorm:"default:true"`
	Subscribers  []User       `gorm:"foreignKey:UserId"`
	CreatedAt    time.Time    `gorm:"autoCreateTime"`
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
